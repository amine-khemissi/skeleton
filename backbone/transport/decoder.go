package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"

	"github.com/amine-khemissi/skeleton/backbone/errorsklt"

	"github.com/gorilla/mux"
)

func DecodeRequest(request interface{}) func(ctx context.Context, r *http.Request) (interface{}, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		return decodeRequest(request, r)
	}
}

func decodeRequest(request interface{}, r *http.Request) (interface{}, error) {
	var err error
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil && err.Error() != "EOF" {
		return nil, err
	}
	request, err = decodePath(request, r)
	if err != nil {
		return nil, errorsklt.Stack(err, "failed to decode path")
	}
	request, err = decodeQueryString(request, r)
	if err != nil {
		return nil, errorsklt.Stack(err, "failed to decode query string")
	}
	request, err = decodeHeader(request, r)
	if err != nil {
		return nil, errorsklt.Stack(err, "failed to decode header")
	}
	return request, nil
}

func decodePath(request interface{}, r *http.Request) (interface{}, error) {
	for i := 0; i < reflect.TypeOf(request).Elem().NumField(); i++ {
		fld := reflect.TypeOf(request).Elem().Field(i)
		elt, found := fld.Tag.Lookup("path")
		if !found {
			continue
		}
		vars := mux.Vars(r)
		if err := decodeScalar(fld.Type.Kind(), vars[elt], reflect.ValueOf(request).Elem().Field(i).Addr()); err != nil {
			return nil, errorsklt.Stack(err, "failed to decode scalar", vars[elt])
		}
	}
	return request, nil
}
func decodeQueryString(request interface{}, r *http.Request) (interface{}, error) {
	for i := 0; i < reflect.TypeOf(request).Elem().NumField(); i++ {
		fld := reflect.TypeOf(request).Elem().Field(i)
		elt, found := fld.Tag.Lookup("qs")
		if !found {
			continue
		}
		qs := r.URL.Query().Get(elt)
		if err := decodeScalar(fld.Type.Kind(), qs, reflect.ValueOf(request).Elem().Field(i)); err != nil {
			return nil, errorsklt.Stack(err, "failed to decode scalar", qs)
		}
	}
	return request, nil
}

func decodeHeader(request interface{}, r *http.Request) (interface{}, error) {
	for i := 0; i < reflect.TypeOf(request).Elem().NumField(); i++ {
		fld := reflect.TypeOf(request).Elem().Field(i)
		tag, found := fld.Tag.Lookup("header")
		if !found {
			continue
		}
		if err := decodeScalar(fld.Type.Kind(), r.Header.Get(tag), reflect.ValueOf(request).Elem().Field(i)); err != nil {
			return nil, errorsklt.Stack(err, "failed to decode scalar", tag)
		}
	}
	return request, nil
}

func decodeScalar(kind reflect.Kind, input string, v reflect.Value) error {

	v = reflect.Indirect(v)
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return errorsklt.New(http.StatusBadRequest, "failed to parse int64", err.Error())
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i, err := strconv.ParseUint(input, 10, 64)
		if err != nil {
			return errorsklt.New(http.StatusBadRequest, "failed to parse int64", err.Error())
		}
		v.SetUint(i)
	case reflect.String:
		v.SetString(input)
	case reflect.Bool:
		b, err := strconv.ParseBool(input)
		if err != nil {
			return errorsklt.New(http.StatusBadRequest, "failed to parse bool", err.Error())
		}
		v.SetBool(b)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errorsklt.New(http.StatusBadRequest, "failed to parse float64", err.Error())
		}
		v.SetFloat(f)
	default:
		return errorsklt.New(http.StatusBadRequest, "unsupported type", kind.String())
	}
	return nil
}
