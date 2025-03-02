package ai

var ListAI []string = []string{
	"Groq",
}

func IsAgentInList(agent string) bool {
	for _, ai := range ListAI {
		if ai == agent {
			return true
		}
	}
	return false
}
