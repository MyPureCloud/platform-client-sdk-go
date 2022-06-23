package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
	)

// MessagingApi provides functions for API endpoints
type MessagingApi struct {
	Configuration *Configuration
}

// NewMessagingApi creates an API instance using the default configuration
func NewMessagingApi() *MessagingApi {
	fmt.Sprintf(strings.Title(""), "")
	config := GetDefaultConfiguration()
	return &MessagingApi{
		Configuration: config,
	}
}

// NewMessagingApiWithConfig creates an API instance using the provided configuration
func NewMessagingApiWithConfig(config *Configuration) *MessagingApi {
	return &MessagingApi{
		Configuration: config,
	}
}

// DeleteMessagingSupportedcontentSupportedContentId invokes DELETE /api/v2/messaging/supportedcontent/{supportedContentId}
//
// Delete a supported content profile
func (a MessagingApi) DeleteMessagingSupportedcontentSupportedContentId(supportedContentId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/messaging/supportedcontent/{supportedContentId}"
	path = strings.Replace(path, "{supportedContentId}", fmt.Sprintf("%v", supportedContentId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'supportedContentId' is set
	if &supportedContentId == nil {
		// false
		return nil, errors.New("Missing required parameter 'supportedContentId' when calling MessagingApi->DeleteMessagingSupportedcontentSupportedContentId")
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

// GetMessagingSupportedcontent invokes GET /api/v2/messaging/supportedcontent
//
// Get a list of Supported Content profiles
func (a MessagingApi) GetMessagingSupportedcontent(pageSize int, pageNumber int) (*Supportedcontentlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/messaging/supportedcontent"
	defaultReturn := new(Supportedcontentlisting)
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
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, "")
	

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
	var successPayload *Supportedcontentlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Supportedcontentlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetMessagingSupportedcontentSupportedContentId invokes GET /api/v2/messaging/supportedcontent/{supportedContentId}
//
// Get a supported content profile
func (a MessagingApi) GetMessagingSupportedcontentSupportedContentId(supportedContentId string) (*Supportedcontent, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/messaging/supportedcontent/{supportedContentId}"
	path = strings.Replace(path, "{supportedContentId}", fmt.Sprintf("%v", supportedContentId), -1)
	defaultReturn := new(Supportedcontent)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'supportedContentId' is set
	if &supportedContentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'supportedContentId' when calling MessagingApi->GetMessagingSupportedcontentSupportedContentId")
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
	var successPayload *Supportedcontent
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Supportedcontent" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchMessagingSupportedcontentSupportedContentId invokes PATCH /api/v2/messaging/supportedcontent/{supportedContentId}
//
// Update a supported content profile
func (a MessagingApi) PatchMessagingSupportedcontentSupportedContentId(supportedContentId string, body Supportedcontent) (*Supportedcontent, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/messaging/supportedcontent/{supportedContentId}"
	path = strings.Replace(path, "{supportedContentId}", fmt.Sprintf("%v", supportedContentId), -1)
	defaultReturn := new(Supportedcontent)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'supportedContentId' is set
	if &supportedContentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'supportedContentId' when calling MessagingApi->PatchMessagingSupportedcontentSupportedContentId")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling MessagingApi->PatchMessagingSupportedcontentSupportedContentId")
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

	var successPayload *Supportedcontent
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Supportedcontent" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostMessagingSupportedcontent invokes POST /api/v2/messaging/supportedcontent
//
// Create a Supported Content profile
func (a MessagingApi) PostMessagingSupportedcontent(body Supportedcontent) (*Supportedcontent, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/messaging/supportedcontent"
	defaultReturn := new(Supportedcontent)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling MessagingApi->PostMessagingSupportedcontent")
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

	var successPayload *Supportedcontent
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Supportedcontent" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

