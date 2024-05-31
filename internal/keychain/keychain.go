package keychain

import (
	"errors"
	"go.etcd.io/bbolt"
	"log"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrorNotFound    = errors.New("not found")
)

type Keychain struct {
	db *bbolt.DB
}

func NewKeychain(db *bbolt.DB) *Keychain {
	return &Keychain{
		db: db,
	}
}

func (k *Keychain) SavePublicKey(username, publicKeyArmor string) error {
	return k.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("public-key"))
		if err != nil {
			return err
		}

		if bucket.Get([]byte(username)) != nil {
			return ErrAlreadyExists
		}

		return bucket.Put([]byte(username), []byte(publicKeyArmor))
	})
}

func (k *Keychain) SavePrivateKey(username, privateKeyArmor string) error {
	return k.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("private-key"))
		if err != nil {
			return err
		}

		if bucket.Get([]byte(username)) != nil {
			return ErrAlreadyExists
		}

		return bucket.Put([]byte(username), []byte(privateKeyArmor))
	})
}

func (k *Keychain) GetPublicKey(username string) (string, error) {
	var publicKey []byte
	err := k.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("public-key"))
		if bucket == nil {
			return ErrorNotFound
		}

		publicKey = bucket.Get([]byte(username))
		if publicKey == nil {
			log.Println("account not found for username", username)
			return ErrorNotFound
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return string(publicKey), nil
}

func (k *Keychain) GetPrivateKey(username string) (string, error) {
	var privateKey []byte
	err := k.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("private-key"))
		if bucket == nil {
			return ErrorNotFound
		}

		privateKey = bucket.Get([]byte(username))
		if privateKey == nil {
			log.Println("account not found for username", username)
			return ErrorNotFound
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return string(privateKey), nil
}
