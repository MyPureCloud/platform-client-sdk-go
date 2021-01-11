package platformclientv2

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"sync"
)

// Configuration has settings to configure the SDK
type Configuration struct {
	UserName                  string            `json:"userName,omitempty"`
	Password                  string            `json:"password,omitempty"`
	APIKeyPrefix              map[string]string `json:"APIKeyPrefix,omitempty"`
	APIKey                    map[string]string `json:"APIKey,omitempty"`
	debug                     bool              `json:"debug,omitempty"`
	DebugFile                 string            `json:"debugFile,omitempty"`
	OAuthToken                string            `json:"oAuthToken,omitempty"`
	Timeout                   int               `json:"timeout,omitempty"`
	BasePath                  string            `json:"basePath,omitempty"`
	Host                      string            `json:"host,omitempty"`
	Scheme                    string            `json:"scheme,omitempty"`
	AccessToken               string            `json:"accessToken,omitempty"`
	RefreshToken              string            `json:"refreshToken,omitempty"`
	ClientID                  string            `json:"clientId,omitempty"`
	ClientSecret              string            `json:"clientSecret,omitempty"`
	ShouldRefreshAccessToken  bool              `json:"shouldRefreshAccessToken,omitempty"`
	RefreshInProgress         int64             `json:"refreshInProgress,omitempty"`
	RefreshTokenWaitTime      int               `json:"refreshTokenWaitTime,omitempty"`
	DefaultHeader             map[string]string `json:"defaultHeader,omitempty"`
	UserAgent                 string            `json:"userAgent,omitempty"`
	APIClient                 APIClient         `json:"APIClient,omitempty"`
}

// AuthResponse contains the access token to use in future requests
type AuthResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
}

// AuthErrorResponse gives you some intel when authorization goes boom
type AuthErrorResponse struct {
	Error            string `json:"error,omitempty"`
	Description      string `json:"description,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

var once sync.Once
var instance *Configuration

// GetDefaultConfiguration returns the shared default configuration instance
func GetDefaultConfiguration() *Configuration {
	once.Do(func() {
		instance = NewConfiguration()
	})
	return instance
}

// NewConfiguration returns a new Configuration instance
func NewConfiguration() *Configuration {
	c := &Configuration{
		BasePath:                 "https://api.mypurecloud.com",
		UserName:                 "",
		debug:                    false,
		ShouldRefreshAccessToken: true,
		RefreshTokenWaitTime:     10,
		DefaultHeader:            make(map[string]string),
		APIKey:                   make(map[string]string),
		APIKeyPrefix:             make(map[string]string),
		UserAgent:                "PureCloud SDK",
	}
	c.APIClient = NewAPIClient(c)
	return c
}

// AuthorizeClientCredentials authorizes this Configuration instance using client credentials.
// The access token will be set automatically and API instances using this configuration object can now make authorized requests.
func (c *Configuration) AuthorizeClientCredentials(clientID string, clientSecret string) error {
	authHostRegex := regexp.MustCompile(`(?i)\/\/api\.`)
	authHost := authHostRegex.ReplaceAllString(c.BasePath, "//login.")
	headerParams := make(map[string]string)
	headerParams["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret))
	formParams := url.Values{}
	formParams["grant_type"] = []string{"client_credentials"}
	response, err := c.APIClient.CallAPI(authHost+"/oauth/token", "POST", nil, headerParams, nil, formParams, "", nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if response.StatusCode != 200 {
		var authErrorResponse *AuthErrorResponse
		err = json.Unmarshal([]byte(response.RawBody), &authErrorResponse)
		if err != nil {
			return err
		}
		return fmt.Errorf("Auth Error: %v (%v - %v)", authErrorResponse.Description, authErrorResponse.Error, authErrorResponse.ErrorDescription)
	}

	var authResponse *AuthResponse
	err = json.Unmarshal([]byte(response.RawBody), &authResponse)
	if err != nil {
		return err
	}
	c.AccessToken = authResponse.AccessToken
	if c.AccessToken == "" {
		return fmt.Errorf("Auth Error: No access token found")
	}
	c.Debugf("Token exipres in %v seconds\n", authResponse.ExpiresIn)
	return nil
}

// AuthorizeCodeGrant authorizes this Configuration instance using an authorization code grant.
// The access and refresh tokens will be set automatically and API instances using this configuration object can now make authorized requests.
func (c *Configuration) AuthorizeCodeGrant(clientID string, clientSecret string, authCode string, redirectUri string) (*AuthResponse, error) {
	c.ClientID = clientID
	c.ClientSecret = clientSecret
	authHostRegex := regexp.MustCompile(`(?i)\/\/api\.`)
	authHost := authHostRegex.ReplaceAllString(c.BasePath, "//login.")
	headerParams := make(map[string]string)
	headerParams["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret))
	headerParams["Content-Type"] = "application/x-www-form-urlencoded"
	formParams := url.Values{}
	formParams["grant_type"] = []string{"authorization_code"}
	formParams["code"] = []string{authCode}
	formParams["redirect_uri"] = []string{redirectUri}
	response, err := c.APIClient.CallAPI(authHost+"/oauth/token", "POST", nil, headerParams, nil, formParams, "", nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if response.StatusCode != 200 {
		var authErrorResponse *AuthErrorResponse
		err = json.Unmarshal([]byte(response.RawBody), &authErrorResponse)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("Auth Error: %v (%v - %v)", authErrorResponse.Description, authErrorResponse.Error, authErrorResponse.ErrorDescription)
	}

	var authResponse *AuthResponse
	err = json.Unmarshal([]byte(response.RawBody), &authResponse)
	if err != nil {
		return nil, err
	}
	c.AccessToken = authResponse.AccessToken
	if c.AccessToken == "" {
		return nil, fmt.Errorf("Auth Error: No access token found")
	}
	c.RefreshToken = authResponse.RefreshToken
	if c.RefreshToken == "" {
		return nil, fmt.Errorf("Auth Error: No refresh token found")
	}
	c.Debugf("Token exipres in %v seconds\n", authResponse.ExpiresIn)

	return authResponse, nil
}

// RefreshAuthorizationCodeGrant requests a new access token for the authorization code grant.
// The access and refresh tokens will be set automatically and API instances using this configuration object can continue to make authorized requests.
func (c *Configuration) RefreshAuthorizationCodeGrant(clientID string, clientSecret string, refreshToken string) (*AuthResponse, error) {
	authHostRegex := regexp.MustCompile(`(?i)\/\/api\.`)
	authHost := authHostRegex.ReplaceAllString(c.BasePath, "//login.")
	headerParams := make(map[string]string)
	headerParams["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret))
	headerParams["Content-Type"] = "application/x-www-form-urlencoded"
	formParams := url.Values{}
	formParams["grant_type"] = []string{"refresh_token"}
	formParams["refresh_token"] = []string{refreshToken}
	response, err := c.APIClient.CallAPI(authHost+"/oauth/token", "POST", nil, headerParams, nil, formParams, "", nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if response.StatusCode != 200 {
		var authErrorResponse *AuthErrorResponse
		err = json.Unmarshal([]byte(response.RawBody), &authErrorResponse)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("Auth Error: %v (%v - %v)", authErrorResponse.Description, authErrorResponse.Error, authErrorResponse.ErrorDescription)
	}

	var authResponse *AuthResponse
	err = json.Unmarshal([]byte(response.RawBody), &authResponse)
	if err != nil {
		return nil, err
	}
	c.AccessToken = authResponse.AccessToken
	if c.AccessToken == "" {
		return nil, fmt.Errorf("Auth Error: No access token found")
	}
	c.RefreshToken = authResponse.RefreshToken
	if c.RefreshToken == "" {
		return nil, fmt.Errorf("Auth Error: No refresh token found")
	}
	c.Debugf("Token exipres in %v seconds\n", authResponse.ExpiresIn)

	return authResponse, nil
}

// AddDefaultHeader sets a header that will be set on every request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// GetAPIKeyWithPrefix appends a prefix to the API key
func (c *Configuration) GetAPIKeyWithPrefix(APIKeyIdentifier string) string {
	if c.APIKeyPrefix[APIKeyIdentifier] != "" {
		return c.APIKeyPrefix[APIKeyIdentifier] + " " + c.APIKey[APIKeyIdentifier]
	}

	return c.APIKey[APIKeyIdentifier]
}

// SetDebug enables debug tracing for HTTP requests, and probably some other stuff
func (c *Configuration) SetDebug(enable bool) {
	c.debug = enable
}

// GetDebug tells you the value of the debug setting in case you forgot
func (c *Configuration) GetDebug() bool {
	return c.debug
}

// Debug prints the provided message using Println if debug tracing is enabled
func (c *Configuration) Debug(msg interface{}) {
	if !c.debug {
		return
	}
	fmt.Println(msg)
}

// Debugf prints the provided formatted message using Printf if debug tracing is enabled
func (c *Configuration) Debugf(msg string, params ...interface{}) {
	if !c.debug {
		return
	}
	fmt.Printf(msg, params...)
}
