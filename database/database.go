package database

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/boltdb/bolt"
	"os"
	"time"
)

var bucketName = []byte("db")

type Database struct {
	db *bolt.DB
}

func Start(fileName string) (*Database, error) {
	userReadWritePermission := os.FileMode(0600)
	db, err := bolt.Open(fileName, userReadWritePermission, &bolt.Options{Timeout: 100 * time.Millisecond})
	if err != nil {
		return nil, err
	}
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	}); err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) Get(key string) (*string, error) {
	var value *string = nil
	err := d.db.View(func(tx *bolt.Tx) error {
		cursor := tx.Bucket(bucketName).Cursor()
		dbKey, dbValue := cursor.Seek([]byte(key))
		if dbKey == nil || string(dbKey) != key {
			return errors.New("key not found")
		} else if err := gob.NewDecoder(bytes.NewReader(dbValue)).Decode(&value); err != nil {
			return err
		}
		return nil
	})
	return value, err
}

func (d *Database) Set(key string, value string) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(value); err != nil {
		return err
	}
	return d.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bucketName).Put([]byte(key), buf.Bytes())
	})
}

func (d *Database) Delete(key string) error {
	return d.db.Update(func(tx *bolt.Tx) error {
		cursor := tx.Bucket(bucketName).Cursor()
		if dbKey, _ := cursor.Seek([]byte(key)); dbKey == nil || string(dbKey) != key {
			return errors.New("key not found")
		} else {
			return cursor.Delete()
		}
	})
}
