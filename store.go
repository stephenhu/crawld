package main

import (
	//"errors"
)


type Store interface {

	Put(s string, v map[string] interface{})	error
	Get(s string)														([] interface{}, error)
	GetAll()																(map[string] interface{}, error)
	Exists(s string) 												bool

}

type MemStore struct {
	Entities				map[string] map[string] interface{}
}

type RedisStore struct {
}
