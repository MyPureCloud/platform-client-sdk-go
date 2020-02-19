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
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating MessagingApi with base path: %s", strings.ToLower(config.BasePath)))
	return &MessagingApi{
		Configuration: config,
	}
}

// NewMessagingApiWithConfig creates an API instance using the provided configuration
func NewMessagingApiWithConfig(config *Configuration) *MessagingApi {
	config.Debugf("Creating MessagingApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &MessagingApi{
		Configuration: config,
	}
}

// GetMessagingSticker invokes GET /api/v2/messaging/stickers/{messengerType}
//
// Get a list of Messaging Stickers
//
// 
func (a MessagingApi) GetMessagingSticker(messengerType string, pageSize int32, pageNumber int32) (*Messagingstickerentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/messaging/stickers/{messengerType}"
	path = strings.Replace(path, "{messengerType}", fmt.Sprintf("%v", messengerType), -1)
	defaultReturn := new(Messagingstickerentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'messengerType' is set
	if &messengerType == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'messengerType' when calling MessagingApi->GetMessagingSticker")
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, collectionFormat)
	
	

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
	var successPayload *Messagingstickerentitylisting
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

