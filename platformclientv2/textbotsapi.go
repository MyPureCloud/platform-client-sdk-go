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
	fmt.Sprintf(strings.Title(""), "")
	config := GetDefaultConfiguration()
	return &TextbotsApi{
		Configuration: config,
	}
}

// NewTextbotsApiWithConfig creates an API instance using the provided configuration
func NewTextbotsApiWithConfig(config *Configuration) *TextbotsApi {
	return &TextbotsApi{
		Configuration: config,
	}
}

// PostTextbotsBotflowsSessionTurns invokes POST /api/v2/textbots/botflows/sessions/{sessionId}/turns
//
// Issue a bot flow turn event
//
// Send a turn event to an executing bot flow and produce the next action to take.
func (a TextbotsApi) PostTextbotsBotflowsSessionTurns(sessionId string, turnRequest Textbotflowturnrequest) (*Textbotflowturnresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/textbots/botflows/sessions/{sessionId}/turns"
	path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)
	defaultReturn := new(Textbotflowturnresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sessionId' is set
	if &sessionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'sessionId' when calling TextbotsApi->PostTextbotsBotflowsSessionTurns")
	}
	// verify the required parameter 'turnRequest' is set
	if &turnRequest == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'turnRequest' when calling TextbotsApi->PostTextbotsBotflowsSessionTurns")
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
	postBody = &turnRequest

	var successPayload *Textbotflowturnresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		if "Textbotflowturnresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTextbotsBotflowsSessions invokes POST /api/v2/textbots/botflows/sessions
//
// Create an execution instance of a bot flow definition.
//
// The launch is asynchronous; use the returned instance ID to post turns to it using &#39;POST /api/v2/textbots/botflows/sessions/{sessionId}/turns&#39;.
func (a TextbotsApi) PostTextbotsBotflowsSessions(launchRequest Textbotflowlaunchrequest) (*Textbotflowlaunchresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/textbots/botflows/sessions"
	defaultReturn := new(Textbotflowlaunchresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'launchRequest' is set
	if &launchRequest == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'launchRequest' when calling TextbotsApi->PostTextbotsBotflowsSessions")
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
	postBody = &launchRequest

	var successPayload *Textbotflowlaunchresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		if "Textbotflowlaunchresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
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
		if "Posttextresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

