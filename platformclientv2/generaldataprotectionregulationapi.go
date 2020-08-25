package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

// GeneralDataProtectionRegulationApi provides functions for API endpoints
type GeneralDataProtectionRegulationApi struct {
	Configuration *Configuration
}

// NewGeneralDataProtectionRegulationApi creates an API instance using the default configuration
func NewGeneralDataProtectionRegulationApi() *GeneralDataProtectionRegulationApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating GeneralDataProtectionRegulationApi with base path: %s", strings.ToLower(config.BasePath)))
	return &GeneralDataProtectionRegulationApi{
		Configuration: config,
	}
}

// NewGeneralDataProtectionRegulationApiWithConfig creates an API instance using the provided configuration
func NewGeneralDataProtectionRegulationApiWithConfig(config *Configuration) *GeneralDataProtectionRegulationApi {
	config.Debugf("Creating GeneralDataProtectionRegulationApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &GeneralDataProtectionRegulationApi{
		Configuration: config,
	}
}

// GetGdprRequest invokes GET /api/v2/gdpr/requests/{requestId}
//
// Get an existing GDPR request
//
// 
func (a GeneralDataProtectionRegulationApi) GetGdprRequest(requestId string) (*Gdprrequest, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/gdpr/requests/{requestId}"
	path = strings.Replace(path, "{requestId}", fmt.Sprintf("%v", requestId), -1)
	defaultReturn := new(Gdprrequest)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'requestId' is set
	if &requestId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'requestId' when calling GeneralDataProtectionRegulationApi->GetGdprRequest")
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
	var successPayload *Gdprrequest
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

// GetGdprRequests invokes GET /api/v2/gdpr/requests
//
// Get all GDPR requests
//
// 
func (a GeneralDataProtectionRegulationApi) GetGdprRequests(pageSize int32, pageNumber int32) (*Gdprrequestentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/gdpr/requests"
	defaultReturn := new(Gdprrequestentitylisting)
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
	var successPayload *Gdprrequestentitylisting
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

// GetGdprSubjects invokes GET /api/v2/gdpr/subjects
//
// Get GDPR subjects
//
// 
func (a GeneralDataProtectionRegulationApi) GetGdprSubjects(searchType string, searchValue string) (*Gdprsubjectentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/gdpr/subjects"
	defaultReturn := new(Gdprsubjectentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'searchType' is set
	if &searchType == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'searchType' when calling GeneralDataProtectionRegulationApi->GetGdprSubjects")
	}
	// verify the required parameter 'searchValue' is set
	if &searchValue == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'searchValue' when calling GeneralDataProtectionRegulationApi->GetGdprSubjects")
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
	if str, ok := interface{}(searchType).(string); ok {
		if str != "" {
			queryParams["searchType"] = a.Configuration.APIClient.ParameterToString(searchType, collectionFormat)
		}
	} else {
		queryParams["searchType"] = a.Configuration.APIClient.ParameterToString(searchType, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(searchValue).(string); ok {
		if str != "" {
			queryParams["searchValue"] = a.Configuration.APIClient.ParameterToString(searchValue, collectionFormat)
		}
	} else {
		queryParams["searchValue"] = a.Configuration.APIClient.ParameterToString(searchValue, collectionFormat)
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
	var successPayload *Gdprsubjectentitylisting
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

// PostGdprRequests invokes POST /api/v2/gdpr/requests
//
// Submit a new GDPR request
//
// 
func (a GeneralDataProtectionRegulationApi) PostGdprRequests(body Gdprrequest, deleteConfirmed bool) (*Gdprrequest, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/gdpr/requests"
	defaultReturn := new(Gdprrequest)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling GeneralDataProtectionRegulationApi->PostGdprRequests")
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
	if str, ok := interface{}(deleteConfirmed).(string); ok {
		if str != "" {
			queryParams["deleteConfirmed"] = a.Configuration.APIClient.ParameterToString(deleteConfirmed, collectionFormat)
		}
	} else {
		queryParams["deleteConfirmed"] = a.Configuration.APIClient.ParameterToString(deleteConfirmed, collectionFormat)
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

	var successPayload *Gdprrequest
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

