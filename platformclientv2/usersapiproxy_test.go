package platformclientv2

import (
        "testing"
        "fmt"
        "os"
)

func TestExampleUsersApi_GetUsers_OnProxy(t *testing.T) {

        config := GetDefaultConfiguration()
        config.LoggingConfiguration.LogLevel = LNone

        config.BasePath = "https://api." + os.Getenv("PURECLOUD_ENVIRONMENT") // e.g. PURECLOUD_ENVIRONMENT=mypurecloud.com
        clientID := os.Getenv("PURECLOUD_CLIENT_ID")
        clientSecret := os.Getenv("PURECLOUD_CLIENT_SECRET")
        config.GateWayConfiguration = nil

        err := config.AuthorizeClientCredentials(clientID, clientSecret)
        if err != nil {
                panic(err)
        }

        proxyconf := ProxyConfiguration{}
        config.ProxyConfiguration = &proxyconf
        config.ProxyConfiguration.Protocol = "http"
        config.ProxyConfiguration.Host = "localhost"
        config.ProxyConfiguration.Port = "4001"

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