package internal

import (
	"github.com/gomodule/redigo/redis"
)

type IStorageClient interface {
	GetInventory() Inventory
	SetInventory(i Inventory)
	IncrementUserStock(userID string, isbn string)
	DecrementUserStock(userID string, isbn string)
	IncrementLibraryStock(isbn string)
	DecrementLibraryStock(isbn string)
	GetUserInventory(userID string) Inventory
}

// StorageClient serializes library inventory into storage
type storageClient struct {
	redisConn redis.Conn
}

// UniversalStockKey represents the redis key for library's stock
const UniversalStockKey = "universal_stock"

func NewStorageClient() IStorageClient {
	// TODO: Close connection
	// TODO: Actually handle error
	// TODO: Init inventory
	conn, _ := redis.Dial("tcp", ":6379")

	return &storageClient{
		redisConn: conn,
	}
}

func (s *storageClient) GetInventory() Inventory {
	// TODO: Handle error
	reply, _ := redis.Values(s.redisConn.Do("HGETALL", UniversalStockKey))
	inv, _ := redis.StringMap(reply, nil)

	return inv
}

func (s *storageClient) SetInventory(i Inventory) {
	// TODO: Handle error
	// TODO: Make more efficient
	for key, value := range i {
		_, _ = s.redisConn.Do("HSET", UniversalStockKey, key, value)
	}
}

func (s *storageClient) IncrementUserStock(userID string, isbn string) {
	s.redisConn.Do("HINCRBY", userID, isbn, 1)
}

func (s *storageClient) DecrementUserStock(userID string, isbn string) {
	s.redisConn.Do("HINCRBY", userID, isbn, -1)
}

func (s *storageClient) IncrementLibraryStock(isbn string) {
	s.redisConn.Do("HINCRBY", UniversalStockKey, isbn, 1)
}

func (s *storageClient) DecrementLibraryStock(isbn string) {
	s.redisConn.Do("HINCRBY", UniversalStockKey, isbn, -1)
}

func (s *storageClient) GetUserInventory(userID string) Inventory {
	// TODO: Handle error
	reply, _ := redis.Values(s.redisConn.Do("HGETALL", userID))
	inv, _ := redis.StringMap(reply, nil)

	return inv
}
