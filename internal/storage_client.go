package internal

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type IStorageClient interface {
	GetInventory() (Inventory, error)
	SetInventory(i Inventory) error
	IncrementUserStock(userID string, isbn string) error
	DecrementUserStock(userID string, isbn string) error
	IncrementLibraryStock(isbn string) error
	DecrementLibraryStock(isbn string) error
	GetUserInventory(userID string) (Inventory, error)
	Close()
}

// StorageClient serializes library inventory into storage
type storageClient struct {
	redisConn redis.Conn
}

// UniversalStockKey represents the redis key for library's stock
const UniversalStockKey = "universal_stock"

func NewStorageClient() (IStorageClient, error) {
	// TODO: Init inventory
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return nil, err
	}

	return &storageClient{
		redisConn: conn,
	}, nil
}

func (s *storageClient) GetInventory() (Inventory, error) {
	reply, err := redis.Values(s.redisConn.Do("HGETALL", UniversalStockKey))
	if err != nil {
		return nil, fmt.Errorf("Failed to get universal inventory: %s", err.Error())
	}

	inv, err := redis.StringMap(reply, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to convert redis value: %s", err.Error())
	}

	return inv, nil
}

func (s *storageClient) SetInventory(i Inventory) error {
	// TODO: Make more efficient
	for key, value := range i {
		_, err := s.redisConn.Do("HSET", UniversalStockKey, key, value)
		if err != nil {
			return fmt.Errorf("Failed to set inventory: %s", err.Error())
		}
	}

	return nil
}

func (s *storageClient) IncrementUserStock(userID string, isbn string) error {
	_, err := s.redisConn.Do("HINCRBY", userID, isbn, 1)
	if err != nil {
		return fmt.Errorf("Failed to increment user stock: %s", err.Error())
	}

	return nil
}

func (s *storageClient) DecrementUserStock(userID string, isbn string) error {
	_, err := s.redisConn.Do("HINCRBY", userID, isbn, -1)
	if err != nil {
		return fmt.Errorf("Failed to decrement user stock: %s", err.Error())
	}

	return nil
}

func (s *storageClient) IncrementLibraryStock(isbn string) error {
	_, err := s.redisConn.Do("HINCRBY", UniversalStockKey, isbn, 1)
	if err != nil {
		return fmt.Errorf("Failed to increment library stock: %s", err.Error())
	}

	return nil
}

func (s *storageClient) DecrementLibraryStock(isbn string) error {
	_, err := s.redisConn.Do("HINCRBY", UniversalStockKey, isbn, -1)
	if err != nil {
		return fmt.Errorf("Failed to decrement library stock: %s", err.Error())
	}

	return nil
}

func (s *storageClient) GetUserInventory(userID string) (Inventory, error) {
	reply, err := redis.Values(s.redisConn.Do("HGETALL", userID))
	if err != nil {
		return nil, fmt.Errorf("Failed to get user inventory: %s", err.Error())
	}

	inv, err := redis.StringMap(reply, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to convert redis value: %s", err.Error())
	}

	return inv, nil
}

func (s *storageClient) Close() {
	s.redisConn.Close()
}
