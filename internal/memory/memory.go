package memory

import (
	"errors"
	"strconv"
	"strings"
)

type Database struct {
	data []string
}

// Starts the databse and returns a pointer for clients to interact with
func NewDatabase(seed []string) *Database {
	db := Database{
		data: make([]string, 0),
	}

	for _, s := range seed {
		db.AddUrl(s)
	}

	return &db
}

// Adds a url to the database
func (db *Database) AddUrl(url string) (string, error) {
	// Validation
	if url == "" {
		return "", errors.New("url cannot be empty")
	}
	if !strings.HasPrefix(url, "https://") {
		return "", errors.New("url must begin with 'https://'")
	}

	// Append the url to the end of the array and return the id
	db.data = append(db.data, url)
	return strconv.Itoa(len(db.data)), nil
}

// Fetches a url from the database
func (db *Database) FetchUrl(id string) (string, error) {
	intId, err := strconv.Atoi(id)

	if err != nil {
		return "", err
	}

	// Guard against out of range
	if intId > len(db.data) {
		return "", errors.New("id is out of range")
	}
	if intId < 1 {
		return "", errors.New("id is out of range")
	}

	// Fetch the url
	url := db.data[intId-1]
	return url, nil
}
