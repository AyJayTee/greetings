package database

import (
	"errors"
	"strings"
)

type Database struct {
	data []string
}

// Starts the databse and returns a pointer for clients to interact with
func StartDatabase() *Database {
	db := Database{
		data: make([]string, 0),
	}
	return &db
}

// Adds a url to the database
func (db *Database) AddUrl(url string) (int, error) {
	// Validation
	if url == "" {
		return 0, errors.New("url cannot be empty")
	}
	if !strings.HasPrefix(url, "https://") {
		return 0, errors.New("url must begin with 'https://'")
	}

	// Append the url to the end of the array and return the id
	db.data = append(db.data, url)
	return len(db.data), nil
}

// Fetches a url from the database
func (db *Database) FetchUrl(id int) (string, error) {
	// Guard against out of range
	if id > len(db.data) {
		return "", errors.New("id is out of range")
	}
	if id < 1 {
		return "", errors.New("id is out of range")
	}

	// Fetch the url
	url := db.data[id-1]
	return url, nil
}
