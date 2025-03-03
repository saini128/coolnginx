package db

import (
	"coolnginx/models"
	"encoding/json"
	"errors"

	"go.etcd.io/bbolt"
)

var AiBucket = []byte("ai_agent")

func StoreOrUpdateAI(ai *models.AiAgent) error {
	return BoltClient.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(AiBucket)
		if err != nil {
			return err
		}

		data, err := json.Marshal(ai)
		if err != nil {
			return err
		}

		return b.Put([]byte("singleton"), data) // Only one AI agent stored
	})
}

func FetchAI() (*models.AiAgent, error) {
	var ai models.AiAgent

	err := BoltClient.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(AiBucket)
		if b == nil {
			return errors.New("AI agent not found")
		}

		data := b.Get([]byte("singleton"))
		if data == nil {
			return errors.New("AI agent not set")
		}

		return json.Unmarshal(data, &ai)
	})

	if err != nil {
		return nil, err
	}

	return &ai, nil
}
