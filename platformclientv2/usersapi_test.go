package platformclientv2

import (
	"fmt"
	"os"
)

func ExampleUsersApi_GetUsers() {
	// Use the default config instance and retrieve settings from env vars
	config := GetDefaultConfiguration()
	config.LoggingConfiguration.LogLevel = LNone
	config.BasePath = "https://api." + os.Getenv("PURECLOUD_ENVIRONMENT") // e.g. PURECLOUD_ENVIRONMENT=mypurecloud.com
	clientID := os.Getenv("PURECLOUD_CLIENT_ID")
	clientSecret := os.Getenv("PURECLOUD_CLIENT_SECRET")
	config.GateWayConfiguration = nil

	// Authorize the configuration
	err := config.AuthorizeClientCredentials(clientID, clientSecret)
	if err != nil {
		panic(err)
	}

	// Create an API instance using the default config
	usersAPI := NewUsersApi()

	// Invoke API
	_, response, err := usersAPI.GetUsers(100, 1, make([]string, 0), make([]string, 0), "", make([]string, 0), "", "")
	if err != nil {
		fmt.Printf("Error calling GetUsers: %v\n", err)
	} else {
		fmt.Printf("Successfully retrieved user data with status code %v\n", response.StatusCode)
	}
	// Output: Successfully retrieved user data with status code 200
}
