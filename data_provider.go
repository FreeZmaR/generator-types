package main

import (
	"reflect"
)

type DataProvider struct {
	collection map[string]any
}

func NewDataProvider() *DataProvider {
	return &DataProvider{
		collection: make(map[string]any),
	}
}

func (d *DataProvider) Collect(val any) {
	t := reflect.TypeOf(val)
	d.collection[t.String()] = val
}

func (d *DataProvider) Get(val string) (any, bool) {
	data, ok := d.collection[val]

	return data, ok
}
