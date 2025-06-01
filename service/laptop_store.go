package service

import (
	"errors"
	"fmt"
	"pcbook/pb"
	"sync"

	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	//Save saves the laptop to the store.
	Save(laptop *pb.Laptop) error

	//  Find finds the laptop by ID.
	Find(id string) (*pb.Laptop, error)
}

// InMemoryLaptopStore stores laptop in memory.
type inMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

// NewInMemoryLaptopStore returns a new in-memory laptop store.
func NewInMemoryLaptopStore() *inMemoryLaptopStore {
	return &inMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// Save saves the laptop to the in-memory store.
func (store *inMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	// Check if laptop with the same ID already exists
	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	// deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	store.data[laptop.Id] = other
	return nil

}

// Find finds the laptop by ID in the in-memory store.
func (store *inMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	// deep copy
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}
