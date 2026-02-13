package coordinator

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Chutchev/coordinatorAgent/internal/http/server"
	"github.com/Chutchev/goagent/pkg/agent"
	"github.com/Chutchev/goagent/pkg/clients/llm"
	"github.com/Chutchev/goagent/pkg/config"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"
	//"time"
)

type Coordinator struct {
	systemPrompt string
	userPrompt   string
	*agent.Agent
	agents []string
}

func NewCoordinator(promptFile *string, name, mode string) *Coordinator {
	systemPrompt, err := os.ReadFile(*promptFile)
	if err != nil {
		log.Fatalf("system prompt file read failed: %v", err)
	}
	userPrompt, err := os.ReadFile("prompts/coordinator_prompt.txt")
	if err != nil {
		log.Fatalf("system prompt file read failed: %v", err)
	}
	return &Coordinator{
		Agent:  agent.NewAgent(string(systemPrompt), string(userPrompt), name, mode),
		agents: make([]string, 0),
	}
}

func (c *Coordinator) runInteractive() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите текст (пустая строка для завершения):")
		var lines []string
		for {
			fmt.Print("> ")
			scanner.Scan()
			text := scanner.Text()

			if text == "" {
				break
			}

			lines = append(lines, text)
		}
		prompt := strings.Join(lines, "\n")
		c.do(prompt)
	}
}

func (c *Coordinator) do(userText string) {
	cfg := config.GetConfig()

	client := llm.NewLLMClient(
		cfg.LLMBaseURL,
		cfg.LLMConfig.LLMToken,
	)
	replacer := strings.NewReplacer(
		"{AgentsList}", strings.Join(c.agents, ","),
		"{UserPrompt}", userText,
	)
	up := replacer.Replace(c.GetUserPrompt())

	req := llm.ChatRequest{
		Model:       cfg.LLMConfig.LLMModel,
		Temperature: cfg.LLMConfig.Temperature,
		TopP:        cfg.LLMConfig.TopP,
		Seed:        cfg.LLMConfig.Seed,
		Messages: []llm.Message{
			{
				Role:    "system",
				Content: c.GetSystemPrompt(),
			},
			{
				Role:    "user",
				Content: up,
			},
		},
	}
	r, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Choices[0].Message.Content)
}

func (c *Coordinator) RunHTTP() {
	s := server.NewServer("0.0.0.0", 8080)
	s.Start()
}

func (c *Coordinator) Run() {
	go c.survey()
	switch c.GetMode() {
	case "i":
		go c.runInteractive()
	//case "grpc":
	//	go c.runGRPC()
	case "http":
		go c.RunHTTP()
	default:
		log.Fatal("")
	}

	// Сюда вставить запуск HTTP Сервера
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Получен сигнал завершения, останавливаемся...")
}

func (c *Coordinator) survey() {
	c.agents = []string{"developer", "analytic", "architector"}
	//for {
	//	slog.Info("survey")
	//	time.Sleep(10 * time.Second)
	//}
}
