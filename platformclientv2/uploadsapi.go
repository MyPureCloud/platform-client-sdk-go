package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

// UploadsApi provides functions for API endpoints
type UploadsApi struct {
	Configuration *Configuration
}

// NewUploadsApi creates an API instance using the default configuration
func NewUploadsApi() *UploadsApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating UploadsApi with base path: %s", strings.ToLower(config.BasePath)))
	return &UploadsApi{
		Configuration: config,
	}
}

// NewUploadsApiWithConfig creates an API instance using the provided configuration
func NewUploadsApiWithConfig(config *Configuration) *UploadsApi {
	config.Debugf("Creating UploadsApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &UploadsApi{
		Configuration: config,
	}
}

// PostUploadsPublicassetsImages invokes POST /api/v2/uploads/publicassets/images
//
// Creates presigned url for uploading a public asset image
//
// 
func (a UploadsApi) PostUploadsPublicassetsImages(body Uploadurlrequest) (*Uploadurlresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/uploads/publicassets/images"
	defaultReturn := new(Uploadurlresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling UploadsApi->PostUploadsPublicassetsImages")
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

	var successPayload *Uploadurlresponse
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

