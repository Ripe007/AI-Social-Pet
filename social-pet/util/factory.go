package util

import (
	"errors"
	"social-pet/util/singleton"
	"sync"
)

var Factory factory

func InitFactory() {
	Factory.instances = make(map[string]singleton.Singleton)
}

type factory struct {
	instances map[string]singleton.Singleton
	lock      sync.Mutex
}

func (f *factory) Set(name string, init singleton.InitFunc) bool {
	f.lock.Lock()
	defer f.lock.Unlock()
	if _, ok := f.instances[name]; !ok {
		f.instances[name] = singleton.NewSingleton(init)
		return true
	}
	return false
}

func (f *factory) Get(name string) (interface{}, error) {
	if _, ok := f.instances[name]; ok {
		return f.instances[name].Get()
	}
	return nil, errors.New("factory get error : " + name + " not exists.")
}
