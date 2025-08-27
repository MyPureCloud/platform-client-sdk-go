package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
	)

// WebMessagingApi provides functions for API endpoints
type WebMessagingApi struct {
	Configuration *Configuration
}

// NewWebMessagingApi creates an API instance using the default configuration
func NewWebMessagingApi() *WebMessagingApi {
	fmt.Sprintf(strings.Title(""), "")
	config := GetDefaultConfiguration()
	return &WebMessagingApi{
		Configuration: config,
	}
}

// NewWebMessagingApiWithConfig creates an API instance using the provided configuration
func NewWebMessagingApiWithConfig(config *Configuration) *WebMessagingApi {
	return &WebMessagingApi{
		Configuration: config,
	}
}

// DeleteWebmessagingDeploymentPushdevice invokes DELETE /api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}
//
// Delete device information
func (a WebMessagingApi) DeleteWebmessagingDeploymentPushdevice(deploymentId string, tokenId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}"
	path = strings.Replace(path, "{deploymentId}", url.PathEscape(fmt.Sprintf("%v", deploymentId)), -1)
	path = strings.Replace(path, "{tokenId}", url.PathEscape(fmt.Sprintf("%v", tokenId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'deploymentId' is set
	if &deploymentId == nil {
		// false
		return nil, errors.New("Missing required parameter 'deploymentId' when calling WebMessagingApi->DeleteWebmessagingDeploymentPushdevice")
	}
	// verify the required parameter 'tokenId' is set
	if &tokenId == nil {
		// false
		return nil, errors.New("Missing required parameter 'tokenId' when calling WebMessagingApi->DeleteWebmessagingDeploymentPushdevice")
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

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

	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	}
	return response, err
}

// GetWebmessagingMessages invokes GET /api/v2/webmessaging/messages
//
// Get the messages for a web messaging session.
func (a WebMessagingApi) GetWebmessagingMessages(pageSize int, pageNumber int) (*Webmessagingmessageentitylist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webmessaging/messages"
	defaultReturn := new(Webmessagingmessageentitylist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
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
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

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
	var successPayload *Webmessagingmessageentitylist
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Webmessagingmessageentitylist" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWebmessagingDeploymentPushdevice invokes PATCH /api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}
//
// Edit device information
func (a WebMessagingApi) PatchWebmessagingDeploymentPushdevice(deploymentId string, tokenId string, body Pushdeviceupdaterequest) (*APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}"
	path = strings.Replace(path, "{deploymentId}", url.PathEscape(fmt.Sprintf("%v", deploymentId)), -1)
	path = strings.Replace(path, "{tokenId}", url.PathEscape(fmt.Sprintf("%v", tokenId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'deploymentId' is set
	if &deploymentId == nil {
		// false
		return nil, errors.New("Missing required parameter 'deploymentId' when calling WebMessagingApi->PatchWebmessagingDeploymentPushdevice")
	}
	// verify the required parameter 'tokenId' is set
	if &tokenId == nil {
		// false
		return nil, errors.New("Missing required parameter 'tokenId' when calling WebMessagingApi->PatchWebmessagingDeploymentPushdevice")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return nil, errors.New("Missing required parameter 'body' when calling WebMessagingApi->PatchWebmessagingDeploymentPushdevice")
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

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


	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	}
	return response, err
}

// PostWebmessagingDeploymentPushdevice invokes POST /api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}
//
// Add a new device information
func (a WebMessagingApi) PostWebmessagingDeploymentPushdevice(deploymentId string, tokenId string, body Pushdeviceinsertrequest) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/webmessaging/deployments/{deploymentId}/pushdevices/{tokenId}"
	path = strings.Replace(path, "{deploymentId}", url.PathEscape(fmt.Sprintf("%v", deploymentId)), -1)
	path = strings.Replace(path, "{tokenId}", url.PathEscape(fmt.Sprintf("%v", tokenId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'deploymentId' is set
	if &deploymentId == nil {
		// false
		return nil, errors.New("Missing required parameter 'deploymentId' when calling WebMessagingApi->PostWebmessagingDeploymentPushdevice")
	}
	// verify the required parameter 'tokenId' is set
	if &tokenId == nil {
		// false
		return nil, errors.New("Missing required parameter 'tokenId' when calling WebMessagingApi->PostWebmessagingDeploymentPushdevice")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return nil, errors.New("Missing required parameter 'body' when calling WebMessagingApi->PostWebmessagingDeploymentPushdevice")
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

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


	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	}
	return response, err
}

