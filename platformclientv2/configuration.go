package platformclientv2

import (
	"crypto/rand"
	"math/big"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

// Configuration has settings to configure the SDK
type Configuration struct {
	UserName                 string                `json:"userName,omitempty"`
	Password                 string                `json:"password,omitempty"`
	APIKeyPrefix             map[string]string     `json:"APIKeyPrefix,omitempty"`
	APIKey                   map[string]string     `json:"APIKey,omitempty"`
	OAuthToken               string                `json:"oAuthToken,omitempty"`
	Timeout                  int                   `json:"timeout,omitempty"`
	BasePath                 string                `json:"basePath,omitempty"`
	Host                     string                `json:"host,omitempty"`
	Scheme                   string                `json:"scheme,omitempty"`
	AccessToken              string                `json:"accessToken,omitempty"`
	RefreshToken             string                `json:"refreshToken,omitempty"`
	ClientID                 string                `json:"clientId,omitempty"`
	ClientSecret             string                `json:"clientSecret,omitempty"`
	ShouldRefreshAccessToken bool                  `json:"shouldRefreshAccessToken,omitempty"`
	RefreshInProgress        int64                 `json:"refreshInProgress,omitempty"`
	RefreshTokenWaitTime     int                   `json:"refreshTokenWaitTime,omitempty"`
	AccessTokenExpiresIn     int                   `json:"accessTokenExpiresIn,omitempty"`
    	AutomaticTokenRefresh    bool                  `json:"automaticTokenRefresh,omitempty"`
	DefaultHeader            map[string]string     `json:"defaultHeader,omitempty"`
	UserAgent                string                `json:"userAgent,omitempty"`
	APIClient                APIClient             `json:"APIClient,omitempty"`
	RetryConfiguration       *RetryConfiguration   `json:"retryConfiguration,omitempty"`
	LoggingConfiguration     *LoggingConfiguration `json:"loggingConfiguration,omitempty"`
	ConfigFilePath           string                `json:"configFilePath,omitempty"`
	AutoReloadConfig         bool                  `json:"autoReloadConfig,omitempty"`
	ProxyConfiguration       *ProxyConfiguration   `json:"proxyConfiguration,omitempty"`
}

const (
	us_east_1      	= "https://api.mypurecloud.com"
	eu_west_1      	= "https://api.mypurecloud.ie"
	ap_southeast_2 	= "https://api.mypurecloud.com.au"
	ap_northeast_1 	= "https://api.mypurecloud.jp"
	eu_central_1   	= "https://api.mypurecloud.de"
	us_west_2      	= "https://api.usw2.pure.cloud"
	ca_central_1   	= "https://api.cac1.pure.cloud"
	ap_northeast_2 	= "https://api.apne2.pure.cloud"
	eu_west_2      	= "https://api.euw2.pure.cloud"
	ap_south_1     	= "https://api.aps1.pure.cloud"
	us_east_2     	= "https://api.use2.us-gov-pure.cloud"
	sa_east_1	 	= "https://api.sae1.pure.cloud"
	me_central_1 	= "https://api.mec1.pure.cloud"
    ap_northeast_3 	= "https://api.apne3.pure.cloud"
    eu_central_2 	= "https://api.euc2.pure.cloud"
)

// RetryConfiguration has settings to configure the SDK retry logic
type RetryConfiguration struct {
	RetryWaitMin    time.Duration   `json:"retry_wait_min,omitempty"`
	RetryWaitMax    time.Duration   `json:"retry_wait_max,omitempty"`
	RetryMax        int             `json:"retry_max,omitempty"`
	RequestLogHook  RequestLogHook  `json:"request_log_hook,omitempty"`
  	ResponseLogHook ResponseLogHook `json:"response_log_hook,omitempty"`
}

// ProxyConfiguration has settings to configure the SDK Proxy logic
type ProxyConfiguration struct {
	Protocol    string   `json:"protocol,omitempty"`
	Host        string   `json:"host,omitempty"`
	Port        string   `json:"port,omitempty"`
	Auth  		*Auth     `json:"auth,omitempty"`
}

type Auth struct {
	UserName    string   `json:"username,omitempty"`
	Password    string   `json:"password,omitempty"`
}

type RequestLogHook func(*http.Request, int)

type ResponseLogHook func(*http.Response)

// LoggingConfiguration has settings to configure the SDK logging
type LoggingConfiguration struct {
	LogLevel        LoggingLevel `json:"logLevel,omitempty"`
	LogRequestBody  bool         `json:"logRequestBody,omitempty"`
	LogResponseBody bool         `json:"logResponseBody,omitempty"`
	logFormat       LoggingFormat
	logFilePath     string
	logToConsole    bool
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

var (
	once     sync.Once
	instance *Configuration
)

// GetDefaultConfiguration returns the shared default Configuration instance
func GetDefaultConfiguration() *Configuration {
	once.Do(func() {
		instance = NewConfiguration()
	})
	return instance
}

// NewConfiguration returns a new Configuration instance
func NewConfiguration() *Configuration {
	homeDir, _ := os.UserHomeDir()

	// Make initial Configuration instance with default values
	c := generateDefaultConfig(filepath.Join(homeDir, ".genesyscloudgo", "config"), true)

	_ = c.updateConfigFromFile()
	if c.AutoReloadConfig {
		go c.periodicConfigUpdater()
	}

	c.LoggingConfiguration.configureLogging()
	c.APIClient = NewAPIClient(c)
	return c
}

// GetDefaultConfigurationWithConfigFile returns the shared default Configuration instance with overrides provided by a config file
func GetDefaultConfigurationWithConfigFile(filePath string, autoReload bool) *Configuration {
	once.Do(func() {
		instance = NewConfigurationWithConfigFile(filePath, autoReload)
	})
	return instance
}

// NewConfigurationWithConfigFile returns a new Configuration instance with overrides provided by a config file
func NewConfigurationWithConfigFile(filePath string, autoReload bool) *Configuration {
	// Make initial Configuration instance with default values
	c := generateDefaultConfig(filePath, autoReload)

	_ = c.updateConfigFromFile()
	if c.AutoReloadConfig {
		go c.periodicConfigUpdater()
	}

	c.LoggingConfiguration.configureLogging()
	c.APIClient = NewAPIClient(c)
	return c
}

func generateDefaultConfig(filePath string, autoReload bool) *Configuration {
	return &Configuration{
		BasePath: "https://api.mypurecloud.com",
		UserName: "",
		LoggingConfiguration: &LoggingConfiguration{
			LogLevel:     LNone,
			logFormat:    Text,
			logToConsole: true,
		},
		RetryConfiguration: &RetryConfiguration{
			RetryMax:     0,
			RetryWaitMax: time.Duration(0),
		},
		ShouldRefreshAccessToken: true,
		RefreshTokenWaitTime:     10,
		DefaultHeader:            make(map[string]string),
		APIKey:                   make(map[string]string),
		APIKeyPrefix:             make(map[string]string),
		UserAgent:                "PureCloud SDK",
		ConfigFilePath:           filePath,
		AutoReloadConfig:         autoReload,
	}
}

func isJsonFile(filePath string) (bool, error) {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	var js map[string]interface{}
	return json.Unmarshal(s, &js) == nil, nil
}

func (c *Configuration) updateConfigFromFile() error {
	isJson, err := isJsonFile(c.ConfigFilePath)
	if err != nil {
		return err
	}

	if isJson {
		viper.SetConfigType("json")
	} else {
		viper.SetConfigType("ini")
	}
	viper.SetConfigFile(c.ConfigFilePath)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	//proxy
    if getConfigString("proxy", "host") != "" {
                if isJson {
                        proxyI := getObject("proxy")
                        jsonbody, err := json.Marshal(proxyI)
                        if err != nil {
                                return err
                        }
                        proxyconf := ProxyConfiguration{}
                        if err := json.Unmarshal(jsonbody, &proxyconf); err != nil {
                                return err
                        }
                        c.ProxyConfiguration = &proxyconf
                } else {
                        proxyconf := ProxyConfiguration{}
                        c.ProxyConfiguration = &proxyconf

                        hostProxy := getConfigString("proxy", "host")
                        if hostProxy != "" {
                                c.ProxyConfiguration.Host = hostProxy
                        }

                        port := getConfigString("proxy", "port")
                        if port != "" {
                                c.ProxyConfiguration.Port = port
                        }

                        protocol := getConfigString("proxy", "protocol")
                        if port != "" {
                                c.ProxyConfiguration.Protocol = protocol
                        }

                        userName := getConfigString("proxy", "auth-username")
                        password := getConfigString("proxy", "auth-password")
                        if userName != "" && password != "" {
                                auth := Auth{}
                                c.ProxyConfiguration.Auth = &auth
                                c.ProxyConfiguration.Auth.UserName = userName
                                c.ProxyConfiguration.Auth.Password = password
                        }
                }
        }


	// logging
	logLevel := getConfigString("logging", "log_level")
	if logLevel != "" {
		c.LoggingConfiguration.LogLevel = *(loggingLevelFromString(logLevel))
	}

	logFormat := getConfigString("logging", "log_format")
	if logFormat != "" {
		c.LoggingConfiguration.logFormat = *(loggingFormatFromString(logFormat))
	}

	// Only update if the value is present
	if getConfigString("logging", "log_to_console") != "" {
		c.LoggingConfiguration.logToConsole = getConfigBool("logging", "log_to_console")
	}

	logFilePath := getConfigString("logging", "log_file_path")
	if logFilePath != "" {
		c.LoggingConfiguration.logFilePath = logFilePath
	}

	if getConfigString("logging", "log_request_body") != "" {
		c.LoggingConfiguration.LogRequestBody = getConfigBool("logging", "log_request_body")
	}

	if getConfigString("logging", "log_response_body") != "" {
		c.LoggingConfiguration.LogResponseBody = getConfigBool("logging", "log_response_body")
	}

	c.LoggingConfiguration.configureLogging()

	// general
	host := getConfigString("general", "host")
	if host != "" {
		c.BasePath = host
	}

	if getConfigString("general", "live_reload_config") != "" {
		c.AutoReloadConfig = getConfigBool("general", "live_reload_config")
	}

	// reauthentication
	if getConfigString("reauthentication", "refresh_access_token") != "" {
		c.ShouldRefreshAccessToken = getConfigBool("reauthentication", "refresh_access_token")
	}

	if getConfigString("reauthentication", "refresh_token_wait_max") != "" {
		c.RefreshTokenWaitTime = getConfigInt("reauthentication", "refresh_token_wait_max")
	}

	// retry
	if getConfigString("retry", "retry_wait_min") != "" {
		c.RetryConfiguration.RetryWaitMin = time.Duration(getConfigInt("retry", "retry_wait_min")) * time.Second
	}

	if getConfigString("retry", "retry_wait_max") != "" {
		c.RetryConfiguration.RetryWaitMax = time.Duration(getConfigInt("retry", "retry_wait_max")) * time.Second
	}

	if getConfigString("retry", "retry_max") != "" {
		c.RetryConfiguration.RetryMax = getConfigInt("retry", "retry_max")
	}

	return nil
}

func getConfigString(section, key string) string {
	value := viper.GetString(fmt.Sprintf("%s.%s", section, key))

	return strings.Trim(fmt.Sprintf("%s", value), " ")
}

func getObject(section string) interface{} {
        return viper.Get(fmt.Sprintf("%s", section))
}

func getConfigBool(section, key string) bool {
	return viper.GetBool(fmt.Sprintf("%s.%s", section, key))
}

func getConfigInt(section, key string) int {
	return viper.GetInt(fmt.Sprintf("%s.%s", section, key))
}

func (c *Configuration) periodicConfigUpdater() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write ||
					event.Op&fsnotify.Create == fsnotify.Create {
						if !c.AutoReloadConfig {
							return
						}
						// Only act on updates to the config file.
						if event.Name == c.ConfigFilePath {
							_ = c.updateConfigFromFile()
						}
				}
			case _, ok := <-watcher.Errors:
				if !ok {
					return
				}
			}
		}
	}()

	watchedDirectory := filepath.Dir(c.ConfigFilePath)
	err = watcher.Add(watchedDirectory)
	if err != nil {
		return
	}

	<-done
}

func getFileHash(filePath string) (string, error) {
	hasher := sha256.New()
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	hasher.Write(s)

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func (c *Configuration) setAutomaticTokenRefresh(value bool) {
    c.AutomaticTokenRefresh = value
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
	if err != nil && response == nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		var authErrorResponse *AuthErrorResponse
		err = json.Unmarshal([]byte(response.RawBody), &authErrorResponse)
		if err != nil {
			return err
		}
		return fmt.Errorf("Auth Error: %v - %v (%v)", response.StatusCode, authErrorResponse.Error, authErrorResponse.ErrorDescription)
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
	c.AccessTokenExpiresIn = authResponse.ExpiresIn

    if c.AutomaticTokenRefresh {
        go func(){
            select {
            case <-time.After(time.Second * time.Duration(c.AccessTokenExpiresIn - 60)):
		log.Println("Refreshing access token")
                err := c.AuthorizeClientCredentials(clientID, clientSecret)
                if err != nil {
                    panic(err)
                }
            }
        }()
    }

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
	if err != nil && response == nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		var authErrorResponse *AuthErrorResponse
		err = json.Unmarshal([]byte(response.RawBody), &authErrorResponse)
		if err != nil {
			return nil, err
		}
		return  nil, fmt.Errorf("Auth Error: %v - %v (%v)", response.StatusCode, authErrorResponse.Error, authErrorResponse.ErrorDescription)
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
	if err != nil && response == nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		var authErrorResponse *AuthErrorResponse
		err = json.Unmarshal([]byte(response.RawBody), &authErrorResponse)
		if err != nil {
			return nil, err
		}
		return  nil, fmt.Errorf("Auth Error: %v - %v (%v)", response.StatusCode, authErrorResponse.Error, authErrorResponse.ErrorDescription)
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

	return authResponse, nil
}

// Generate a random string used as PKCE Code Verifier.
func (c *Configuration) GeneratePKCECodeVerifier(n int) (string, error) {
	if n < 43 || n > 128 {
		return "", fmt.Errorf("Error: PKCE Code Verifier (length) must be between 43 and 128 characters")
	}
	const unreservedCharacters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-._~"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(unreservedCharacters))))
		if err != nil {
			return "", err
		}
		ret[i] = unreservedCharacters[num.Int64()]
	}
	
	return string(ret), nil
}

// Compute Base64Url PKCE Code Challenge from Code Verifier.
func (c *Configuration) ComputePKCECodeChallenge(codeVerifier string) (string, error) {
	if len(codeVerifier) < 43 || len(codeVerifier) > 128 {
		return "", fmt.Errorf("Error: PKCE Code Verifier (length) must be between 43 and 128 characters")
	}
	h := sha256.New()
	h.Write([]byte(codeVerifier))
	codeVerifierHash := h.Sum(nil)
	codeChallenge := base64.StdEncoding.EncodeToString([]byte(codeVerifierHash))
	codeChallenge = strings.Replace(codeChallenge, "+", "-", -1) // 62nd char of encoding
	codeChallenge = strings.Replace(codeChallenge, "/", "_", -1) // 63rd char of encoding
	codeChallenge = (strings.Split(codeChallenge, "="))[0]
	return codeChallenge, nil
}

// AuthorizePKCEGrant authorizes this Configuration instance using a pkce grant.
// The access and refresh tokens will be set automatically and API instances using this configuration object can now make authorized requests.
func (c *Configuration) AuthorizePKCEGrant(clientID string, codeVerifier string, authCode string, redirectUri string) (*AuthResponse, error) {
	c.ClientID = clientID
	authHostRegex := regexp.MustCompile(`(?i)\/\/api\.`)
	authHost := authHostRegex.ReplaceAllString(c.BasePath, "//login.")
	headerParams := make(map[string]string)
	headerParams["Content-Type"] = "application/x-www-form-urlencoded"
	formParams := url.Values{}
	formParams["grant_type"] = []string{"authorization_code"}
	formParams["code"] = []string{authCode}
	formParams["code_verifier"] = []string{codeVerifier}
	formParams["client_id"] = []string{clientID}
	formParams["redirect_uri"] = []string{redirectUri}
	response, err := c.APIClient.CallAPI(authHost+"/oauth/token", "POST", nil, headerParams, nil, formParams, "", nil)
	if err != nil && response == nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		var authErrorResponse *AuthErrorResponse
		err = json.Unmarshal([]byte(response.RawBody), &authErrorResponse)
		if err != nil {
			return nil, err
		}
		return  nil, fmt.Errorf("Auth Error: %v - %v (%v)", response.StatusCode, authErrorResponse.Error, authErrorResponse.ErrorDescription)
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

func (c *LoggingConfiguration) SetLogFormat(logFormat LoggingFormat) {
	c.logFormat = logFormat
	c.configureLogging()
}

func (c *LoggingConfiguration) GetLogFormat() LoggingFormat {
	return c.logFormat
}

func (c *LoggingConfiguration) SetLogFilePath(logFilePath string) {
	c.logFilePath = logFilePath
	c.configureLogging()
}

func (c *LoggingConfiguration) GetLogFilePath() string {
	return c.logFilePath
}

func (c *LoggingConfiguration) SetLogToConsole(logToConsole bool) {
	c.logToConsole = logToConsole
	c.configureLogging()
}

func (c *LoggingConfiguration) GetLogToConsole() bool {
	return c.logToConsole
}
