package main

import (
	//"errors"
)


type Store interface {

	Put(s string, p CrawldEntity)						error
	Get(s string)														(*CrawldEntity, error)
	GetAll()																(map[string]CrawldEntity, error)
	Exists(s string) 												bool

}

type MemStore struct {
	Entities				map[string] CrawldEntity
}

type RedisStore struct {
}
