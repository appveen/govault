package vault

import (
	"log"
	"os"

	"github.com/boltdb/bolt"
)

var dbFile *os.File

//DB  - Base truststore struct
type DB struct {
	bucketName string
	STORE      *bolt.DB
}

//InitDB - Initialize trustore
func InitDB(filePath string, bucketName string) *DB {
	db, err := bolt.Open(filePath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{
		STORE:      db,
		bucketName: bucketName,
	}
}

//Upsert - Transactional Addition/Updation to Bucket
func (store *DB) Upsert(key string, data []byte) error {
	return store.STORE.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(store.bucketName))
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), data)
		return err
	})
}

//Delete - Transactional Delete to Bucket
func (store *DB) Delete(key string) error {
	return store.STORE.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(store.bucketName))
		err := b.Delete([]byte(key))
		return err
	})
}

//Get - Transactional Get for a key in specific bucket
func (store *DB) Get(key string) ([]byte, error) {
	tx, err := store.STORE.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	bucket := tx.Bucket([]byte(store.bucketName))
	value := bucket.Get([]byte(key))
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return value, nil
}

//CloseDB - Close the current DB file
func (store *DB) CloseDB() error {
	return store.STORE.Close()
}
