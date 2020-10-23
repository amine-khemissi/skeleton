package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type Config interface {
	Get(moduleName, keyName string) interface{}
	isExpired() bool
}

func init() {
	configFile = flag.String("config-file", "config.json", "path to config file")
	go reloadConf(&locker)
}

var configFile *string

var gConf Config
var locker sync.Mutex

type config struct {
	Content    map[string]map[string]interface{} `json:","`
	lastUpdate time.Time
	l          *sync.Mutex
}

func (c *config) Get(moduleName, keyName string) interface{} {
	c.l.Lock()
	defer c.l.Unlock()
	keys, moduleExist := c.Content[moduleName]
	if !moduleExist {
		panic(fmt.Sprintln("module", moduleName, "not found"))
	}
	val, keyExist := keys[keyName]
	if !keyExist {
		panic(fmt.Sprintln(moduleName, "::", keyName, "not found"))
	}
	return val
}

func Instance() Config {
	locker.Lock()
	defer locker.Unlock()
	if gConf == nil {
		gConf = loadConf(&locker)
	}
	return gConf
}

func (c config) isExpired() bool {

	fi, err := os.Stat(*configFile)
	if err != nil {
		panic(err)
	}
	if c.lastUpdate.After(fi.ModTime()) {
		return false
	}
	return true
}

func loadConf(l *sync.Mutex) Config {
	fmt.Println("need to reload conf")
	tmp := &config{
		l: l,
	}
	bts, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}
	if errUnmarshall := json.Unmarshal(bts, &tmp.Content); errUnmarshall != nil {
		panic(errUnmarshall)
	}
	tmp.lastUpdate = time.Now()
	fmt.Println("conf reloaded")
	return tmp
}

func reloadConf(l *sync.Mutex) {

	for {
		time.Sleep(3 * time.Second)
		if gConf != nil && !gConf.isExpired() {
			continue
		}
		l.Lock()
		gConf = loadConf(l)
		l.Unlock()
	}
}
