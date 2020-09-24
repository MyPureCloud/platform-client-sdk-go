---
title: Platform API Client SDK - Go
---

A Go package to interface with the Genesys Cloud Platform API. View the documentation on the [pkg.go.dev](https://pkg.go.dev/github.com/MyPureCloud/platform-client-sdk-go/platformclientv2). Browse the source code on [Github](https://github.com/MyPureCloud/platform-client-sdk-go).

Latest version: 20.0.0 [![GitHub release](https://img.shields.io/github/release/mypurecloud/platform-client-sdk-go.svg)]()

## Golang Version Dependency

:::danger
Some macOS users encounter the error "argument list too long" when building or installing this package if their golang version is less than 1.15. Visit [Go Downloads](https://golang.org/dl/) and install _go1.15_ or newer if this error is thrown.
:::

## Get SDK Package

Retrieve the package from https://github.com/MyPureCloud/platform-client-sdk-go using `go get`:

```go
go get github.com/mypurecloud/platform-client-sdk-go/platformclientv2
```

## Using the SDK

### Importing the package

```go
import (
	"github.com/MyPureCloud/platform-client-sdk-go/platformclientv2"
)
```

### Configuring the SDK

The SDK can be configured by setting properties on a `Configuration` instance. Applications that are using a single access token (most use cases) can use the default configuration.

The default configuration will be used when creating new API instances without supplying the config.

```go
// Get default config to set config options
config := platformclientv2.GetDefaultConfiguration()

// Create API instance using default config
usersAPI := platformclientv2.NewUsersApi()
```

Applications that will be making requests using multiple access tokens can create multiple configuration instances and authorize them individually. This is a common pattern when using the SDK in a backend web server where multiple Genesys Cloud users will authenticate using an auth code grant. 

When creating configuration instances, they must be used when creating new API instances.

```go
// Create new config
config := NewConfiguration()

// Create API instance using config
usersAPI := platformclientv2.NewUsersApiWithConfig(config)
```

#### Setting the environment

To connect to regional Genesys Cloud instances, provide the Platform API's base path:

```go
config.BasePath = "https://api.mypurecloud.jp"
```

#### Setting the access token

If using a grant other than client credentials, the access token can be set manually:

```go
config.AccessToken = "anaccesstokenobtainedmanuallyfromanoauthgrant"
```

#### Enable debug logging

Enabling debug logging will trace out information about all requests and responses:

```go
config.SetDebug(false)
```

### Authorization

#### Client Credentials

The SDK provides a helper to authorize using client credentials. Use the `AuthorizeClientCredentials` function on the client to make a request to Genesys Cloud to exchange the client id and secret for an access token. If no error was returned, the configuration instance is now authorized and can begin making API requests. 

```go
err := config.AuthorizeClientCredentials(os.Getenv("GENESYS_CLOUD_CLIENT_ID"), os.Getenv("GENESYS_CLOUD_CLIENT_SECRET"))
if err != nil {
    panic(err)
}
```

### Making Requests

Once the SDK is authorized, API requests can be made. Each function on the API instance returns three values, or just the latter two if the resource does not return any value (e.g. a DELETE operation):

* A struct representing the API response. This will be nil if any error occurred.
* An `APIResponse` instance providing extended information about the response. This is useful for debugging, custom logging, and advanced error handling. This may be nil under certain error circumstances. This will have a value when the API returns an error.
* An error, if any error occurred at any point during processing. 

```go
userID := "asdf1234-5678-90ab-cde1-123456789012"
usersAPI := platformclientv2.NewUsersApi()

user, response, err := usersAPI.GetUser(userID, make([]string, 0), "")
fmt.Printf("Response:\n  Success: %v\n  Status code: %v\n  Correlation ID: %v\n", response.IsSuccess, response.StatusCode, response.CorrelationID)
if err != nil {
    fmt.Printf("Error calling GetUser: %v\n", err)
} else {
    fmt.Printf("Hello, %v\n", *user.Name)
}
```


## Versioning

The SDK's version is incremented according to the [Semantic Versioning Specification](https://semver.org/). The decision to increment version numbers is determined by [diffing the Platform API's swagger](https://github.com/purecloudlabs/platform-client-sdk-common/blob/master/modules/swaggerDiff.js) for automated builds, and optionally forcing a version bump when a build is triggered manually (e.g. releasing a bugfix).


## Support

This package is intended to be forwards compatible with v2 of Genesys Cloud's Platform API. While the general policy for the API is not to introduce breaking changes, there are certain additions and changes to the API that cause breaking changes for the SDK, often due to the way the API is expressed in its swagger definition. Because of this, the SDK can have a major version bump while the API remains at major version 2. While the SDK is intended to be forward compatible, patches will only be released to the latest version. For these reasons, it is strongly recommended that all applications using this SDK are kept up to date and use the latest version of the SDK.

For any issues, questions, or suggestions for the SDK, visit the [Genesys Cloud Developer Forum](https://developer.mypurecloud.com/forum/).
