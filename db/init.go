package db

import (
	"fmt"

	"go.etcd.io/bbolt"
)

var dbPath = "./data.db"
var BoltClient *bbolt.DB

func InitBoltDB() error {
	var err error
	BoltClient, err = bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	return nil
}
