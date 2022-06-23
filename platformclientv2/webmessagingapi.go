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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
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

