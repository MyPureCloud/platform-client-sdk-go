package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

// SpeechTextAnalyticsApi provides functions for API endpoints
type SpeechTextAnalyticsApi struct {
	Configuration *Configuration
}

// NewSpeechTextAnalyticsApi creates an API instance using the default configuration
func NewSpeechTextAnalyticsApi() *SpeechTextAnalyticsApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating SpeechTextAnalyticsApi with base path: %s", strings.ToLower(config.BasePath)))
	return &SpeechTextAnalyticsApi{
		Configuration: config,
	}
}

// NewSpeechTextAnalyticsApiWithConfig creates an API instance using the provided configuration
func NewSpeechTextAnalyticsApiWithConfig(config *Configuration) *SpeechTextAnalyticsApi {
	config.Debugf("Creating SpeechTextAnalyticsApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &SpeechTextAnalyticsApi{
		Configuration: config,
	}
}

// GetConversationTranscriptproperty invokes GET /api/v2/conversations/{conversationId}/transcriptproperties/{communicationId}
//
// Get the pre-signed S3 URL for the transcript of a specific communication of a conversation
//
// 
func (a SpeechTextAnalyticsApi) GetConversationTranscriptproperty(conversationId string, communicationId string) (*Transcriptproperty, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/conversations/{conversationId}/transcriptproperties/{communicationId}"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{communicationId}", fmt.Sprintf("%v", communicationId), -1)
	defaultReturn := new(Transcriptproperty)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling SpeechTextAnalyticsApi->GetConversationTranscriptproperty")
	}
	// verify the required parameter 'communicationId' is set
	if &communicationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'communicationId' when calling SpeechTextAnalyticsApi->GetConversationTranscriptproperty")
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
	var successPayload *Transcriptproperty
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

