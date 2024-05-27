package client

import (
	"REST-API/REST-API/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CallExternal3rdPatryApi() (models.CallExternalApiResponse, error) {
	// Create a new HTTP client
	client := http.Client{}

	var response models.CallExternalApiResponse

	// Create a GET request to the API endpoint
	req, err := http.NewRequest("GET", "https://dummyjson.com/posts", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return response, err
	}

	// Set headers if needed
	// req.Header.Set("Authorization", "Bearer <your-token>")
	// req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return response, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return response, err
	}

	// Print the response body
	fmt.Println(string(body))

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return response, err
	}

	return response, err
}
