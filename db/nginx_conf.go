package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.etcd.io/bbolt"
)

var (
	NginxBucket  = []byte("nginx_confs")
	MaxInstances = 5
)

// NginxConfig represents the structure to store the configuration with a timestamp
type NginxConfig struct {
	Timestamp string                 `json:"timestamp"`
	Config    map[string]interface{} `json:"config"`
}

// StoreNginxConfig stores a new Nginx config while maintaining a maximum of 5 instances
func StoreNginxConfig(config map[string]interface{}) error {
	return BoltClient.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(NginxBucket)
		if err != nil {
			return err
		}

		// Retrieve all stored keys
		keys := []string{}
		err = bucket.ForEach(func(k, v []byte) error {
			keys = append(keys, string(k))
			return nil
		})
		if err != nil {
			return err
		}

		// Get the latest stored config (if exists)
		if len(keys) > 0 {
			latestKey := keys[len(keys)-1] // Last inserted key
			latestData := bucket.Get([]byte(latestKey))

			var latestConfig NginxConfig
			if err := json.Unmarshal(latestData, &latestConfig); err == nil {
				// Compare new config with latest stored config
				newConfigJSON, _ := json.Marshal(config)
				latestConfigJSON, _ := json.Marshal(latestConfig.Config)

				if string(newConfigJSON) == string(latestConfigJSON) {
					return errors.New("No update: Configuration unchanged")
				}
			}
		}

		// Create a new config entry with timestamp
		nginxEntry := NginxConfig{
			Timestamp: time.Now().Format(time.RFC3339),
			Config:    config,
		}

		// Convert struct to JSON
		data, err := json.Marshal(nginxEntry)
		if err != nil {
			return err
		}

		// Assign new key (incrementing index)
		newKey := string(len(keys) + 1)

		// Store new config
		err = bucket.Put([]byte(newKey), data)
		if err != nil {
			return err
		}

		// Maintain only the last 5 entries (FIFO behavior)
		if len(keys)+1 > MaxInstances {
			oldestKey := keys[0]
			bucket.Delete([]byte(oldestKey))
		}

		return nil
	})
}

// GetAllNginxConfigs retrieves all stored Nginx configurations
func GetAllNginxConfigs() ([]NginxConfig, error) {
	var configs []NginxConfig

	err := BoltClient.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(NginxBucket)
		if bucket == nil {
			return nil // No configs found
		}

		return bucket.ForEach(func(k, v []byte) error {
			var config NginxConfig
			fmt.Println("Key is ", string(k))
			if err := json.Unmarshal(v, &config); err != nil {
				return err
			}
			configs = append(configs, config)
			return nil
		})
	})

	return configs, err
}
