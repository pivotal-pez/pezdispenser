package pezauth

import (
	"fmt"
	"strings"

	"github.com/garyburd/redigo/redis"
)

//NewKeyGen - create a new implementation of a KeyGenerator interface
func NewKeyGen(doer Doer, guid GUIDMaker) KeyGenerator {
	return &KeyGen{
		store:     doer,
		guidMaker: guid,
	}
}

func parseKeysResponse(r interface{}) (key, username, hash string, err error) {

	if resArr := r.([]interface{}); len(resArr) > 0 {
		ba := resArr[0].([]byte)
		hash = string(ba[:])
		key, username, err = hashSplit(hash)

	} else {
		err = ErrEmptyKeyResponse
	}
	return
}

func (s *KeyGen) getHashMap(hash string) (res interface{}, err error) {
	var byteResponse interface{}
	byteResponse, err = s.store.Do("HMGET", redis.Args{hash}.Add(HMFieldActive).Add(HMFieldDetails)...)
	castedByteResponse := byteResponse.([]interface{})

	if len(castedByteResponse) == 2 {
		res = map[string]interface{}{
			HMFieldActive:  string(castedByteResponse[0].([]byte)[:]),
			HMFieldDetails: string(castedByteResponse[1].([]byte)[:]),
		}

	} else {
		res = byteResponse
	}
	return
}

//Get - gets a key for a user
func (s *KeyGen) Get(user string) (res string, err error) {
	var r interface{}
	search := fmt.Sprintf("%s:*", user)

	if r, err = s.store.Do("KEYS", search); r != nil && err == nil {
		res, _, _, err = parseKeysResponse(r)
	}
	return
}

//GetByKey - gets a user for a given key
func (s *KeyGen) GetByKey(key string) (hash string, val interface{}, err error) {
	var r interface{}
	search := fmt.Sprintf("*:%s", key)

	if r, err = s.store.Do("KEYS", search); r != nil && err == nil {

		if _, _, hash, err = parseKeysResponse(r); err == nil {
			val, err = s.getHashMap(hash)
		}
	}
	return
}

func (s *KeyGen) getHash(user string) (hash string, err error) {
	var r interface{}
	search := fmt.Sprintf("%s:*", user)

	if r, err = s.store.Do("KEYS", search); r != nil && err == nil {
		_, _, hash, err = parseKeysResponse(r)
	}
	return
}

func hashSplit(hash string) (key, username string, err error) {
	usernameIndex := 0
	keyIndex := 1
	hashSplitArrayLen := 2

	if splitHash := strings.Split(hash, ":"); len(splitHash) == hashSplitArrayLen {
		key = splitHash[keyIndex]
		username = splitHash[usernameIndex]

	} else {
		err = ErrUnparsableHash
	}
	return
}

func createHash(user, guid string) (hash string) {
	hash = fmt.Sprintf("%s:%s", user, guid)
	return
}

//Create - creates a new key for a user
func (s *KeyGen) Create(user string, details string) (err error) {
	guid := s.guidMaker.Create()
	hash := createHash(user, guid)
	row := map[string]string{
		HMFieldActive:  "true",
		HMFieldDetails: details,
	}
	_, err = s.store.Do("HMSET", redis.Args{hash}.AddFlat(row)...)
	return
}

//Delete - deletes a key for a user
func (s *KeyGen) Delete(user string) (err error) {
	var apikey string

	if apikey, err = s.Get(user); err == nil {
		_, err = s.store.Do("DEL", createHash(user, apikey))
	}
	return
}
