package main

import (
	"flag"
	"github.com/Chutchev/coordinatorAgent/internal/coordinator"
)

func main() {
	systemPromptFile := flag.String("spf", "prompts/coordinator_agent_prompt.txt", "system prompt file")
	botName := flag.String("bot", "coordinator", "bot name")
	mode := flag.String("mode", "http", "agent mode. interactive/http/grpc")
	flag.Parse()

	newAgent := coordinator.NewCoordinator(systemPromptFile, *botName, *mode)
	newAgent.Run()

}
