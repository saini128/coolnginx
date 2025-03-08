package checks

import (
	"coolnginx/ai"
	"coolnginx/db"
	"coolnginx/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Init() {
	err := db.InitBoltDB()
	if err != nil {
		fmt.Printf("Failed to initialize BoltDB: %v\n", err)
		os.Exit(1)
		return
	}
	fmt.Println("BoltDB initialized successfully.")
	fmt.Println("Checking Nginx running status")
	err = CheckIfNginxIsRunning()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Nginx Running Properly")
	}

	//check if db is empty
	err = CheckAIModelExists()
	if err != nil {
		fmt.Println(err.Error())
		err = StoreAIAgent()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	//make a go routing that curls the ai api to check if it is working or not
	// 2 edge cases, api unauthorized || network issues
	// nginx.StoreNginxMainConfigFile()
	FetchNginxConfig()
	fmt.Println("Initialization completed successfully.")

}

func CheckIfNginxIsRunning() error {
	cmd := exec.Command("systemctl", "is-active", "nginx")
	output, _ := cmd.Output()
	status := string(output)
	if status != "active\n" {
		return fmt.Errorf("nginx is not running, status: %s", status)
	}
	return nil
}

func CheckAIModelExists() error {
	agent, err := db.FetchAI()
	if err != nil {
		return err
	}
	if agent == nil {
		return fmt.Errorf("AI model not found in database")
	}
	fmt.Println("AI model Present: ", agent)
	return nil
}
func StoreAIAgent() error {
	agent := &models.AiAgent{}
	agent.Name = "Groq"
	agent.ApiKey = os.Getenv("GROQ_API_KEY")
	if agent.ApiKey == "" {
		return fmt.Errorf("api key not found")
	}
	fmt.Println(agent)
	err := ai.AddAi(agent)
	if err != nil {
		return err
	}
	fmt.Println("AI Agent stored successfully")
	return nil
}

func FetchNginxConfig() {

	configs, _ := db.GetAllNginxConfigs()
	conf, _ := json.Marshal(configs)
	log.Println("Stored Configs:", string(conf))
}
