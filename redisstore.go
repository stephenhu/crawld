package main

import (
	"errors"
	//"log"
	//"time"

	//"github.com/gomodule/redigo/redis"
	//"github.com/go-redis/redis"
)


func (rs RedisStore) Put(k string, v map[string]interface{}) (error) {

	if k == "" {
		return errors.New("Empty key not allowed")
	}

	err := rediss.HMSet(k, v).Err()

	if err != nil {
		return err
	}

	return nil

} // Put


func (rs RedisStore) Get(k string) ([]interface{}, error) {

	r, err := rediss.HMGet(
		k, FIELD_CREATED, FIELD_REFERRAL, FIELD_RATING).Result()

	if err != nil {
		return nil, err
	} else {
		
		return r, nil

	}

} // Get


func (ms RedisStore) GetAll() (map[string]interface{}, error) {

	return nil, nil

} // GetAll


func (rs RedisStore) Exists(k string) bool {

	r, err := rediss.HGetAll(k).Result()

	if err != nil {
		appLog(err.Error(), "RedisStore.Exists")
	} else {

		if len(r) > 0 {
			return true
		}

	}
	return false

} // Exists
