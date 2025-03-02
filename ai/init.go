package ai

import "errors"

func AddAi(agent, api_key string) error {
	var AgentNotAllowed error = errors.New("AgentNotAllowed")
	if !IsAgentInList(agent) {
		return AgentNotAllowed

	}
	//use db package to add ai to db
	return nil
}
