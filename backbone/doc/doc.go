package doc

import (
	"reflect"

	"github.com/amine-khemissi/skeleton/backbone/endpointimpl"
)

var gGenerator Generator

type Generator interface {
	Register(impl endpointimpl.EndpointImpl)
	NewEndpoint() endpointimpl.EndpointImpl
}

type generator struct {
	doc map[string]map[string]interface{}
}

func (g *generator) Register(impl endpointimpl.EndpointImpl) {
	g.doc[impl.GetPath()][impl.GetVerb()] = ""
	for i := 0; i < reflect.TypeOf(impl.GetRequest()).Elem().NumField(); i++ {
		fld := reflect.TypeOf(impl.GetRequest()).Elem().Field(i)
		//todo : remove _ and replace it by elt
		_, found := fld.Tag.Lookup("json")
		if !found {
			continue
		}
		// todo : create requestBody  object and populate it
	}
	// todo: populate the responses associated

}

func (g *generator) NewEndpoint() endpointimpl.EndpointImpl {
	return g
}

func NewGenerator() Generator {
	if gGenerator == nil {
		gGenerator = &generator{
			doc: map[string]map[string]interface{}{},
		}
	}
	return gGenerator
}
