package platformclientv2


import (
	"testing"
	/*
	"fmt"
	"os"
	"path/filepath"
	*/
)

//commented,  will be uncommented once permanent proxy is made avaiable.

func TestExampleUsersApi_GetUsers_OnProxy(t *testing.T) {
	// Use the default config instance and retrieve settings from env vars
	/*
	homeDir, _ := os.UserHomeDir()
	config := NewConfigurationWithConfigFile(filepath.Join(homeDir, ".genesyscloudgo", "configproxy.ini"),true)
	

	config.BasePath = "https://api." + os.Getenv("PURECLOUD_ENVIRONMENT") // e.g. PURECLOUD_ENVIRONMENT=mypurecloud.com
	clientID := os.Getenv("PURECLOUD_CLIENT_ID")
	clientSecret := os.Getenv("PURECLOUD_CLIENT_SECRET")


	err := config.AuthorizeClientCredentials(clientID, clientSecret)
	if err != nil {
		panic(err)
	}

	// Create an API instance using the default config
	usersAPI := NewUsersApiWithConfig(config)

	// Invoke API
	_, response, err := usersAPI.GetUsers(100, 1, make([]string, 0), make([]string, 0), "", make([]string, 0), "", "")
	if err != nil {
		fmt.Printf("Error calling GetUsers: %v\n", err)
	} else {
		fmt.Printf("Successfully retrieved user data with status code %v\n", response.StatusCode)
	}
	// Output: Successfully retrieved user data with status code 200
	*/
}


/*
func TestExampleUsersApi_GetUsers_OnProxy_withConfigJson(t *testing.T) {
	// Use the default config instance and retrieve settings from env vars
	homeDir, _ := os.UserHomeDir()
	config := NewConfigurationWithConfigFile(filepath.Join(homeDir, "genesyscloudgo", "configproxy.json"),true)
	

	config.BasePath = "https://api." + os.Getenv("PURECLOUD_ENVIRONMENT") // e.g. PURECLOUD_ENVIRONMENT=mypurecloud.com
	clientID := os.Getenv("PURECLOUD_CLIENT_ID")
	clientSecret := os.Getenv("PURECLOUD_CLIENT_SECRET")


	err := config.AuthorizeClientCredentials(clientID, clientSecret)
	if err != nil {
		panic(err)
	}

	// Create an API instance using the default config
	usersAPI := NewUsersApiWithConfig(config)

	// Invoke API
	_, response, err := usersAPI.GetUsers(100, 1, make([]string, 0), make([]string, 0), "", make([]string, 0), "", "")
	if err != nil {
		fmt.Printf("Error calling GetUsers: %v\n", err)
	} else {
		fmt.Printf("Successfully retrieved user data with status code %v\n", response.StatusCode)
	}
	// Output: Successfully retrieved user data with status code 200
}
*/