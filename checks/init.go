package checks

import (
	"coolnginx/db"
	"fmt"
	"os"
)

func Init() {
	err := db.InitBoltDB()
	if err != nil {
		fmt.Printf("Failed to initialize BoltDB: %v\n", err)
		os.Exit(1)
		return
	}
	fmt.Println("BoltDB initialized successfully.")

	//check if nginx is running
	//check if nginx config is accessible
	//check if db is empty
	//check if ai model exists

}
