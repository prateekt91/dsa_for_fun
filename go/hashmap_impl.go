package main

import (
	"fmt"
)

// Entry to hold key-value pairs in the map

type Entry struct {
	Key   string
	Value string
}

// MyHashMap ---> represents the custom implementation of a hash map
// Buckets ---> 2D slice (slice of slices) to hold the entries in the hash map
type MyHashMap struct {
	Buckets [][]Entry
	Size    int // Size of underlyig array
}

// NewMyHashMap ---> constructor function to initialize the hash map
func NewMyHashMap(size int) *MyHashMap {

	buckets := make([][]Entry, size) // Create a 2D slice with the specified size

	for i := range buckets {
		buckets[i] = []Entry{} // Initialize each bucket as an empty slice of Entry
	}

	return &MyHashMap{
		Buckets: buckets,
		Size:    size,
	}
}

// hashFunction ---> computes the hash value for a given key

func (m *MyHashMap) hashFunction(key string) uint32 {

	var hash uint32 = 0

	for i := 0; i < len(key); i++ {
		hash = hash*31 + uint32(key[i]) // Using a common hash function (multiplying by a prime number and adding the ASCII value of the character)	}
	}
	return hash
}

// hashToIndex converts the raw hash code into a valid index within our map's size.
// This prevents the hash code from exceeding the capacity of the bucket array.
func (m *MyHashMap) hashToIndex(hash uint32) int {
	return int(hash % uint32(m.Size))
}

func (m *MyHashMap) Put(key string, value string) {

	//Hash the key to get the index of the bucket where the entry should be stored
	rawHash := m.hashFunction(key)

	// Determine the index of the bucket using the hash value
	index := m.hashToIndex(rawHash)

	//Collision handling: Check if the key already exists in the bucket
	bucket := m.Buckets[index]

	//Check if the key already exists in the bucket
	for i, entry := range bucket {
		if entry.Key == key {
			// If the key already exists, update the value and return
			bucket[i].Value = value
			fmt.Printf(" [Update] Key '%s' updated at index %d.\n", key, index)
			return
		}
	}

	// If the key does not exist, add a new entry to the bucket
	bucket = append(bucket, Entry{Key: key, Value: value})
	m.Buckets[index] = bucket
	fmt.Printf(" [Insert] Key '%s' added at index %d.\n", key, index)
}

// Get retrieves the value associated with the given key from the hash map.

func (m *MyHashMap) Get(key string) (string, bool) {
	//Hashing the key to get the index of the bucket where the entry should be stored
	rawHash := m.hashFunction(key)

	// Determine the index of the bucket using the hash value
	index := m.hashToIndex(rawHash)

	//Traeverse the bucket to find the entry with the matching key
	bucket := m.Buckets[index]

	for _, entry := range bucket {
		if entry.Key == key {
			return entry.Value, true // Return the value and true if the key is found
		}
	}

	return "", false // Return an empty string and false if the key is not found
}

// Remove deletes the key-value pair associated with the given key from the hash map.

func (m *MyHashMap) Remove(key string) bool {

	//Hashing the key to get the index of the bucket where the entry should be stored
	rawHash := m.hashFunction(key)
	// Determine the index of the bucket using the hash value
	index := m.hashToIndex(rawHash)
	// Traverse the bucket to find the entry with the matching key
	bucket := m.Buckets[index]

	for i, entry := range bucket {
		if entry.Key == key {
			// Remove the entry from the bucket by slicing out the entry at index i
			m.Buckets[index] = append(bucket[:i], bucket[i+1:]...)
			return true // Return true if the key was found and removed
		}
	}
	return false // Return false if the key was not found

}

func main() {

	hashMap := NewMyHashMap(10)

	hashMap.Put("Prateek", "Have a Black Car")
	hashMap.Put("Rahul", "Have a Blue Car")
	hashMap.Put("Prayas", "Have a White Car")

	fmt.Println(hashMap)

	value, found := hashMap.Get("Prayas")

	if found {
		fmt.Println("Value for key 'Prayas': ", value)
	} else {
		fmt.Println("Key 'Prayas' not found in the hash map.")
	}

	removed := hashMap.Remove("Rahul")
	if removed {
		fmt.Println("Key 'Rahul' removed successfully.")
	} else {
		fmt.Println("Key 'Rahul' not found in the hash map.")
	}

	fmt.Println(hashMap)

	value, found = hashMap.Get("Prateek")
	if found {
		fmt.Println("Value for key 'Prateek': ", value)
	} else {
		fmt.Println("Key 'Prateek' not found in the hash map.")
	}
}
