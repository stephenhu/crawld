package main

import (
	//"errors"
)


func (rs RedisStore) Put(k string, v CrawldEntity) (error) {
	return nil
} // Put


func (rs RedisStore) Get(k string) (*CrawldEntity, error) {

	ce := CrawldEntity{}

	return &ce, nil

} // Get


func (rs RedisStore) GetAll() (map[string] CrawldEntity, error) {

	return nil, nil

} // GetAll


func (rs RedisStore) Exists(k string) bool {
	return false
} // Exists
