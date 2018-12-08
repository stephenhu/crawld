package main

import (
	"errors"
	//"log"
	//"strings"
	//"time"
)


func (ms MemStore) Put(k string, v map[string]interface{}) (error) {

	ms.Entities[k] = v

	return nil

} // Put


func (ms MemStore) Get(k string) ([] interface{}, error) {

	_, ok := ms.Entities[SanitizeURL(k)]

	if ok {
		return nil, nil
	} else {
		return nil, errors.New("URL not found.")
	}

} // Get


func (ms MemStore) GetAll() (map[string] interface{}, error) {

	return nil, nil

} // GetAll


func (ms MemStore) Exists(k string) bool {

	_, ok := ms.Entities[SanitizeURL(k)]

	if !ok {
		return false
	} else {
		return true
	}

	} // Exists
