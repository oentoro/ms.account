package dbclient

import (
	"model"
	"log"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(accountId string) (model.Account, error)
	Seed()
}

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb(){
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0500, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) Seed(){
	initializeBucket()
	seedAccounts()
}

func(bc *BoltClient) initializeBucket(){
	bc.boltDB.update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}