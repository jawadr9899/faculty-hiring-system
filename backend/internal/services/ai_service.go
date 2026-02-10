package services

import (
	"encoding/json"
	"uhs/internal/models"

	"github.com/hedgeg0d/groq"
	"github.com/labstack/echo/v5"
)

type AIOperations interface {
	ProcessCV(c *echo.Context, prompt string) (models.Analytics, error)
}

type AIService struct {
	APIKey string
	Model  string
}

func NewAIService(apiKey string, model string) AIOperations {
	return &AIService{
		APIKey: apiKey,
		Model:  model,
	}
}

func (ai *AIService) ProcessCV(c *echo.Context, prompt string) (models.Analytics, error) {
	client := groq.GroqClient{
		ApiKey: ai.APIKey,
		Model:  ai.Model,
	}
	resp, err := client.Ask(prompt)
	if err != nil {
		return models.Analytics{}, err
	}

	var aiAnalytics models.Analytics
	err = json.Unmarshal([]byte(resp), &aiAnalytics)
	if err != nil {
		c.Logger().Error("Failed to convert the Ai Analytics to DB model " + err.Error())
		
		return models.Analytics{}, err
	}

	return aiAnalytics, nil
}
