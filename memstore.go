package main

import (
	"errors"
	"log"
	//"strings"
	//"time"
)


func (ms MemStore) Put(k string, v CrawldEntity) (error) {

	log.Println(k)
	log.Println(v)
	log.Println(len(ms.Entities))

	ms.Entities[k] = v

	return nil

} // Put


func (ms MemStore) Get(k string) (*CrawldEntity, error) {

	v, ok := ms.Entities[SanitizeURL(k)]

	if ok {
		return &v, nil
	} else {
		return nil, errors.New("URL not found.")
	}

} // Get


func (ms MemStore) GetAll() (map[string] CrawldEntity, error) {

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
