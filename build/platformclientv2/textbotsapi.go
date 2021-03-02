package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

// TextbotsApi provides functions for API endpoints
type TextbotsApi struct {
	Configuration *Configuration
}

// NewTextbotsApi creates an API instance using the default configuration
func NewTextbotsApi() *TextbotsApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating TextbotsApi with base path: %s", strings.ToLower(config.BasePath)))
	return &TextbotsApi{
		Configuration: config,
	}
}

// NewTextbotsApiWithConfig creates an API instance using the provided configuration
func NewTextbotsApiWithConfig(config *Configuration) *TextbotsApi {
	config.Debugf("Creating TextbotsApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &TextbotsApi{
		Configuration: config,
	}
}

// PostTextbotsBotsExecute invokes POST /api/v2/textbots/bots/execute
//
// Send an intent to a bot to start a dialog/interact with it via text
//
// This will either start a bot with the given id or relay a communication to an existing bot session.
func (a TextbotsApi) PostTextbotsBotsExecute(postTextRequest Posttextrequest) (*Posttextresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/textbots/bots/execute"
	defaultReturn := new(Posttextresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'postTextRequest' is set
	if &postTextRequest == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'postTextRequest' when calling TextbotsApi->PostTextbotsBotsExecute")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &postTextRequest

	var successPayload *Posttextresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

