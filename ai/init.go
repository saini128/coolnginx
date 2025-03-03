package ai

import (
	"coolnginx/db"
	"coolnginx/models"
	"errors"
)

func AddAi(agent *models.AiAgent) error {
	var AgentNotAllowed error = errors.New("AgentNotAllowed")
	if !IsAgentInList(agent.Name) {
		return AgentNotAllowed

	}
	err := db.StoreOrUpdateAI(agent)
	if err != nil {
		return err
	}
	return nil
}
