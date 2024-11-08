package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalf("OPENAI_API_KEY not set in environment")
	}

	client := openai.NewClient(apiKey)
	reader := bufio.NewReader(os.Stdin)

	// Initial user input
	fmt.Print("Enter a question: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)

	// Create thread based on user's initial input with context prompt
	thread, err := createThreadWithMessage(client, userInput)
	if err != nil {
		log.Fatalf("Failed to create thread: %v", err)
	}
	fmt.Printf("Created thread with ID: %s\n", thread.ID)

	// Retrieve and print the assistant's initial response
	assistantResponse, err := getAssistantResponseViaChatCompletion(client, thread.ID)
	if err != nil {
		log.Fatalf("Failed to get assistant response: %v", err)
	}
	fmt.Println("Assistant Response:\n", assistantResponse)

	// Save assistant response to the thread
	err = saveAssistantResponseToThread(client, thread.ID, assistantResponse)
	if err != nil {
		log.Fatalf("Failed to save assistant response to thread: %v", err)
	}

	// Print the full thread history
	// printThreadHistory(client, thread.ID)

	// Conversation loop for follow-up questions
	for {
		fmt.Print("\nEnter a follow-up question (or type 'exit' to quit): ")
		userInput, _ = reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if strings.ToLower(userInput) == "exit" {
			break
		}

		// Send follow-up message to thread with explicit context reference
		err := addUserMessageToThread(client, thread.ID, openai.ThreadMessageRoleUser, userInput)
		if err != nil {
			log.Fatalf("Failed to send user message: %v", err)
		}

		// Retrieve and print assistant's response to follow-up question
		assistantResponse, err = getAssistantResponseViaChatCompletion(client, thread.ID)
		if err != nil {
			log.Fatalf("Failed to get assistant response: %v", err)
		}
		fmt.Println("Assistant Response:\n", assistantResponse)

		// Save assistant response to the thread
		err = saveAssistantResponseToThread(client, thread.ID, assistantResponse)
		if err != nil {
			log.Fatalf("Failed to save assistant response to thread: %v", err)
		}

		// Print the full thread history after each message exchange
		// printThreadHistory(client, thread.ID)
	}
}

// Helper function to create a new thread with an initial user message and a guiding instruction
func createThreadWithMessage(client *openai.Client, initialMessage string) (openai.Thread, error) {
	threadRequest := openai.ThreadRequest{
		Messages: []openai.ThreadMessage{
			{
				Role:    openai.ThreadMessageRoleUser,
				Content: initialMessage,
			},
		},
		Metadata: map[string]any{"conversation_type": "chatbot"},
	}
	thread, err := client.CreateThread(context.Background(), threadRequest)
	if err != nil {
		return openai.Thread{}, fmt.Errorf("failed to create thread: %v", err)
	}
	return thread, nil
}

// Helper function to send a message to the thread
func addUserMessageToThread(client *openai.Client, threadID string, role openai.ThreadMessageRole, content string) error {
	message := openai.MessageRequest{
		Role:    string(role),
		Content: content,
	}

	_, err := client.CreateMessage(context.Background(), threadID, message)
	if err != nil {
		return fmt.Errorf("failed to send %s message: %v", role, err)
	}
	//fmt.Printf("%s: %s\n", role, content)
	return nil
}

// Helper function to get assistant's response using CreateChatCompletion
func getAssistantResponseViaChatCompletion(client *openai.Client, threadID string) (string, error) {
	// Retrieve all messages from the thread to get the context
	messages, err := client.ListMessage(context.Background(), threadID, nil, nil, nil, nil, nil)
	if err != nil {
		return "", fmt.Errorf("failed to list messages: %v", err)
	}

	// Convert thread messages to ChatCompletionMessage format for API request
	var chatMessages []openai.ChatCompletionMessage
	for i := len(messages.Messages) - 1; i >= 0; i-- {
		msg := messages.Messages[i]
		chatMessages = append(chatMessages, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content[0].Text.Value, // Asumsi Content[0] adalah tipe Text yang valid
		})
	}

	// fmt.Println("Chat Messages: ", chatMessages)

	// Use CreateChatCompletion to generate an assistant response
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT4o,
			Messages: chatMessages,
		},
	)
	if err != nil {
		return "", fmt.Errorf("ChatCompletion error: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}

// Helper function to save the assistant's response to the thread
func saveAssistantResponseToThread(client *openai.Client, threadID string, assistantResponse string) error {
	message := openai.MessageRequest{
		Role:    string(openai.ThreadMessageRoleAssistant),
		Content: assistantResponse,
	}

	_, err := client.CreateMessage(context.Background(), threadID, message)
	if err != nil {
		return fmt.Errorf("failed to save assistant response to thread: %v", err)
	}
	return nil
}

// Helper function to print the entire thread history
// func printThreadHistory(client *openai.Client, threadID string) {
// 	messages, err := client.ListMessage(context.Background(), threadID, nil, nil, nil, nil, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to retrieve thread history: %v", err)
// 	}

// 	fmt.Println("\n--- Full Thread History ---")
// 	for _, msg := range messages.Messages {
// 		role := cases.Title(language.English).String(string(msg.Role))
// 		content := msg.Content[0].Text.Value // Assuming Content[0] contains valid Text
// 		fmt.Printf("%s: %s\n", role, content)
// 	}
// 	fmt.Println("----------------------------")
// }
