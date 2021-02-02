package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

// WebChatApi provides functions for API endpoints
type WebChatApi struct {
	Configuration *Configuration
}

// NewWebChatApi creates an API instance using the default configuration
func NewWebChatApi() *WebChatApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating WebChatApi with base path: %s", strings.ToLower(config.BasePath)))
	return &WebChatApi{
		Configuration: config,
	}
}

// NewWebChatApiWithConfig creates an API instance using the provided configuration
func NewWebChatApiWithConfig(config *Configuration) *WebChatApi {
	config.Debugf("Creating WebChatApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &WebChatApi{
		Configuration: config,
	}
}

// DeleteWebchatDeployment invokes DELETE /api/v2/webchat/deployments/{deploymentId}
//
// Delete a WebChat deployment
//
// 
func (a WebChatApi) DeleteWebchatDeployment(deploymentId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/deployments/{deploymentId}"
	path = strings.Replace(path, "{deploymentId}", fmt.Sprintf("%v", deploymentId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'deploymentId' is set
	if &deploymentId == nil {
		// 
		return nil, errors.New("Missing required parameter 'deploymentId' when calling WebChatApi->DeleteWebchatDeployment")
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

	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	}
	return response, err
}

// DeleteWebchatGuestConversationMember invokes DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}
//
// Remove a member from a chat conversation
//
// 
func (a WebChatApi) DeleteWebchatGuestConversationMember(conversationId string, memberId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return nil, errors.New("Missing required parameter 'conversationId' when calling WebChatApi->DeleteWebchatGuestConversationMember")
	}
	// verify the required parameter 'memberId' is set
	if &memberId == nil {
		// 
		return nil, errors.New("Missing required parameter 'memberId' when calling WebChatApi->DeleteWebchatGuestConversationMember")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (Guest Chat JWT) required

	// set key with prefix in header
	headerParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")

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

	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	}
	return response, err
}

// DeleteWebchatSettings invokes DELETE /api/v2/webchat/settings
//
// Remove WebChat deployment settings
//
// 
func (a WebChatApi) DeleteWebchatSettings() (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/settings"
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
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

	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	}
	return response, err
}

// GetWebchatDeployment invokes GET /api/v2/webchat/deployments/{deploymentId}
//
// Get a WebChat deployment
//
// 
func (a WebChatApi) GetWebchatDeployment(deploymentId string) (*Webchatdeployment, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/deployments/{deploymentId}"
	path = strings.Replace(path, "{deploymentId}", fmt.Sprintf("%v", deploymentId), -1)
	defaultReturn := new(Webchatdeployment)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'deploymentId' is set
	if &deploymentId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'deploymentId' when calling WebChatApi->GetWebchatDeployment")
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
	var successPayload *Webchatdeployment
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

// GetWebchatDeployments invokes GET /api/v2/webchat/deployments
//
// List WebChat deployments
//
// 
func (a WebChatApi) GetWebchatDeployments() (*Webchatdeploymententitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/deployments"
	defaultReturn := new(Webchatdeploymententitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
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
	var successPayload *Webchatdeploymententitylisting
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

// GetWebchatGuestConversationMediarequest invokes GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}
//
// Get a media request in the conversation
//
// 
func (a WebChatApi) GetWebchatGuestConversationMediarequest(conversationId string, mediaRequestId string) (*Webchatguestmediarequest, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{mediaRequestId}", fmt.Sprintf("%v", mediaRequestId), -1)
	defaultReturn := new(Webchatguestmediarequest)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling WebChatApi->GetWebchatGuestConversationMediarequest")
	}
	// verify the required parameter 'mediaRequestId' is set
	if &mediaRequestId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'mediaRequestId' when calling WebChatApi->GetWebchatGuestConversationMediarequest")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (Guest Chat JWT) required

	// set key with prefix in header
	headerParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")

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
	var successPayload *Webchatguestmediarequest
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

// GetWebchatGuestConversationMediarequests invokes GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests
//
// Get all media requests to the guest in the conversation
//
// 
func (a WebChatApi) GetWebchatGuestConversationMediarequests(conversationId string) (*Webchatguestmediarequestentitylist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/guest/conversations/{conversationId}/mediarequests"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	defaultReturn := new(Webchatguestmediarequestentitylist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling WebChatApi->GetWebchatGuestConversationMediarequests")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (Guest Chat JWT) required

	// set key with prefix in header
	headerParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")

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
	var successPayload *Webchatguestmediarequestentitylist
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

// GetWebchatGuestConversationMember invokes GET /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}
//
// Get a web chat conversation member
//
// 
func (a WebChatApi) GetWebchatGuestConversationMember(conversationId string, memberId string) (*Webchatmemberinfo, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberId), -1)
	defaultReturn := new(Webchatmemberinfo)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling WebChatApi->GetWebchatGuestConversationMember")
	}
	// verify the required parameter 'memberId' is set
	if &memberId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'memberId' when calling WebChatApi->GetWebchatGuestConversationMember")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (Guest Chat JWT) required

	// set key with prefix in header
	headerParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")

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
	var successPayload *Webchatmemberinfo
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

// GetWebchatGuestConversationMembers invokes GET /api/v2/webchat/guest/conversations/{conversationId}/members
//
// Get the members of a chat conversation.
//
// 
func (a WebChatApi) GetWebchatGuestConversationMembers(conversationId string, pageSize int, pageNumber int, excludeDisconnectedMembers bool) (*Webchatmemberinfoentitylist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/guest/conversations/{conversationId}/members"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	defaultReturn := new(Webchatmemberinfoentitylist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling WebChatApi->GetWebchatGuestConversationMembers")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (Guest Chat JWT) required

	// set key with prefix in header
	headerParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	var collectionFormat string
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(pageSize).(string); ok {
		if str != "" {
			queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, collectionFormat)
		}
	} else {
		queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(pageNumber).(string); ok {
		if str != "" {
			queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, collectionFormat)
		}
	} else {
		queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(excludeDisconnectedMembers).(string); ok {
		if str != "" {
			queryParams["excludeDisconnectedMembers"] = a.Configuration.APIClient.ParameterToString(excludeDisconnectedMembers, collectionFormat)
		}
	} else {
		queryParams["excludeDisconnectedMembers"] = a.Configuration.APIClient.ParameterToString(excludeDisconnectedMembers, collectionFormat)
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
	var successPayload *Webchatmemberinfoentitylist
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

// GetWebchatGuestConversationMessage invokes GET /api/v2/webchat/guest/conversations/{conversationId}/messages/{messageId}
//
// Get a web chat conversation message
//
// 
func (a WebChatApi) GetWebchatGuestConversationMessage(conversationId string, messageId string) (*Webchatmessage, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/guest/conversations/{conversationId}/messages/{messageId}"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{messageId}", fmt.Sprintf("%v", messageId), -1)
	defaultReturn := new(Webchatmessage)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling WebChatApi->GetWebchatGuestConversationMessage")
	}
	// verify the required parameter 'messageId' is set
	if &messageId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'messageId' when calling WebChatApi->GetWebchatGuestConversationMessage")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (Guest Chat JWT) required

	// set key with prefix in header
	headerParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")

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
	var successPayload *Webchatmessage
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

// GetWebchatGuestConversationMessages invokes GET /api/v2/webchat/guest/conversations/{conversationId}/messages
//
// Get the messages of a chat conversation.
//
// 
func (a WebChatApi) GetWebchatGuestConversationMessages(conversationId string, after string, before string, sortOrder string, maxResults int) (*Webchatmessageentitylist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/guest/conversations/{conversationId}/messages"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	defaultReturn := new(Webchatmessageentitylist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling WebChatApi->GetWebchatGuestConversationMessages")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (Guest Chat JWT) required

	// set key with prefix in header
	headerParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	var collectionFormat string
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(after).(string); ok {
		if str != "" {
			queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
		}
	} else {
		queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(before).(string); ok {
		if str != "" {
			queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, collectionFormat)
		}
	} else {
		queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(sortOrder).(string); ok {
		if str != "" {
			queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
		}
	} else {
		queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(maxResults).(string); ok {
		if str != "" {
			queryParams["maxResults"] = a.Configuration.APIClient.ParameterToString(maxResults, collectionFormat)
		}
	} else {
		queryParams["maxResults"] = a.Configuration.APIClient.ParameterToString(maxResults, collectionFormat)
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
	var successPayload *Webchatmessageentitylist
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

// GetWebchatSettings invokes GET /api/v2/webchat/settings
//
// Get WebChat deployment settings
//
// 
func (a WebChatApi) GetWebchatSettings() (*Webchatsettings, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/settings"
	defaultReturn := new(Webchatsettings)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
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
	var successPayload *Webchatsettings
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

// PatchWebchatGuestConversationMediarequest invokes PATCH /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}
//
// Update a media request in the conversation, setting the state to ACCEPTED/DECLINED/ERRORED
//
// 
func (a WebChatApi) PatchWebchatGuestConversationMediarequest(conversationId string, mediaRequestId string, body Webchatguestmediarequest) (*Webchatguestmediarequest, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{mediaRequestId}", fmt.Sprintf("%v", mediaRequestId), -1)
	defaultReturn := new(Webchatguestmediarequest)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling WebChatApi->PatchWebchatGuestConversationMediarequest")
	}
	// verify the required parameter 'mediaRequestId' is set
	if &mediaRequestId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'mediaRequestId' when calling WebChatApi->PatchWebchatGuestConversationMediarequest")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WebChatApi->PatchWebchatGuestConversationMediarequest")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (Guest Chat JWT) required

	// set key with prefix in header
	headerParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")

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
	postBody = &body

	var successPayload *Webchatguestmediarequest
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

// PostWebchatDeployments invokes POST /api/v2/webchat/deployments
//
// Create WebChat deployment
//
// 
func (a WebChatApi) PostWebchatDeployments(body Webchatdeployment) (*Webchatdeployment, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/deployments"
	defaultReturn := new(Webchatdeployment)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WebChatApi->PostWebchatDeployments")
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
	postBody = &body

	var successPayload *Webchatdeployment
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

// PostWebchatGuestConversationMemberMessages invokes POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages
//
// Send a message in a chat conversation.
//
// 
func (a WebChatApi) PostWebchatGuestConversationMemberMessages(conversationId string, memberId string, body Createwebchatmessagerequest) (*Webchatmessage, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberId), -1)
	defaultReturn := new(Webchatmessage)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling WebChatApi->PostWebchatGuestConversationMemberMessages")
	}
	// verify the required parameter 'memberId' is set
	if &memberId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'memberId' when calling WebChatApi->PostWebchatGuestConversationMemberMessages")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WebChatApi->PostWebchatGuestConversationMemberMessages")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (Guest Chat JWT) required

	// set key with prefix in header
	headerParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")

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
	postBody = &body

	var successPayload *Webchatmessage
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

// PostWebchatGuestConversationMemberTyping invokes POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing
//
// Send a typing-indicator in a chat conversation.
//
// 
func (a WebChatApi) PostWebchatGuestConversationMemberTyping(conversationId string, memberId string) (*Webchattyping, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberId), -1)
	defaultReturn := new(Webchattyping)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling WebChatApi->PostWebchatGuestConversationMemberTyping")
	}
	// verify the required parameter 'memberId' is set
	if &memberId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'memberId' when calling WebChatApi->PostWebchatGuestConversationMemberTyping")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (Guest Chat JWT) required

	// set key with prefix in header
	headerParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization")

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
	var successPayload *Webchattyping
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

// PostWebchatGuestConversations invokes POST /api/v2/webchat/guest/conversations
//
// Create an ACD chat conversation from an external customer.
//
// This endpoint will create a new ACD Chat conversation under the specified Chat Deployment.  The conversation will begin with a guest member in it (with a role=CUSTOMER) according to the customer information that is supplied. If the guest member is authenticated, the &#39;memberAuthToken&#39; field should include his JWT as generated by the &#39;POST /api/v2/signeddata&#39; resource; if the guest member is anonymous (and the Deployment permits it) this field can be omitted.  The returned data includes the IDs of the conversation created, along with a newly-create JWT token that you can supply to all future endpoints as authentication to perform operations against that conversation. After successfully creating a conversation, you should connect a websocket to the event stream named in the &#39;eventStreamUri&#39; field of the response; the conversation is not routed until the event stream is attached.
func (a WebChatApi) PostWebchatGuestConversations(body Createwebchatconversationrequest) (*Createwebchatconversationresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/guest/conversations"
	defaultReturn := new(Createwebchatconversationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WebChatApi->PostWebchatGuestConversations")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte

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
	postBody = &body

	var successPayload *Createwebchatconversationresponse
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

// PutWebchatDeployment invokes PUT /api/v2/webchat/deployments/{deploymentId}
//
// Update a WebChat deployment
//
// 
func (a WebChatApi) PutWebchatDeployment(deploymentId string, body Webchatdeployment) (*Webchatdeployment, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/deployments/{deploymentId}"
	path = strings.Replace(path, "{deploymentId}", fmt.Sprintf("%v", deploymentId), -1)
	defaultReturn := new(Webchatdeployment)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'deploymentId' is set
	if &deploymentId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'deploymentId' when calling WebChatApi->PutWebchatDeployment")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WebChatApi->PutWebchatDeployment")
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
	postBody = &body

	var successPayload *Webchatdeployment
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

// PutWebchatSettings invokes PUT /api/v2/webchat/settings
//
// Update WebChat deployment settings
//
// 
func (a WebChatApi) PutWebchatSettings(body Webchatsettings) (*Webchatsettings, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webchat/settings"
	defaultReturn := new(Webchatsettings)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WebChatApi->PutWebchatSettings")
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
	postBody = &body

	var successPayload *Webchatsettings
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

