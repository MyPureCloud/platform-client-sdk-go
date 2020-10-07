package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

// JourneyApi provides functions for API endpoints
type JourneyApi struct {
	Configuration *Configuration
}

// NewJourneyApi creates an API instance using the default configuration
func NewJourneyApi() *JourneyApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating JourneyApi with base path: %s", strings.ToLower(config.BasePath)))
	return &JourneyApi{
		Configuration: config,
	}
}

// NewJourneyApiWithConfig creates an API instance using the provided configuration
func NewJourneyApiWithConfig(config *Configuration) *JourneyApi {
	config.Debugf("Creating JourneyApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &JourneyApi{
		Configuration: config,
	}
}

// GetJourneyActiontarget invokes GET /api/v2/journey/actiontargets/{actionTargetId}
//
// Retrieve a single action target.
//
// 
func (a JourneyApi) GetJourneyActiontarget(actionTargetId string) (*Actiontarget, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/journey/actiontargets/{actionTargetId}"
	path = strings.Replace(path, "{actionTargetId}", fmt.Sprintf("%v", actionTargetId), -1)
	defaultReturn := new(Actiontarget)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionTargetId' is set
	if &actionTargetId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionTargetId' when calling JourneyApi->GetJourneyActiontarget")
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
	var successPayload *Actiontarget
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

// GetJourneyActiontargets invokes GET /api/v2/journey/actiontargets
//
// Retrieve all action targets.
//
// 
func (a JourneyApi) GetJourneyActiontargets(pageNumber int32, pageSize int32) (*Actiontargetlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/journey/actiontargets"
	defaultReturn := new(Actiontargetlisting)
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
	if str, ok := interface{}(pageNumber).(string); ok {
		if str != "" {
			queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, collectionFormat)
		}
	} else {
		queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(pageSize).(string); ok {
		if str != "" {
			queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, collectionFormat)
		}
	} else {
		queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, collectionFormat)
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
	var successPayload *Actiontargetlisting
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

// PatchJourneyActiontarget invokes PATCH /api/v2/journey/actiontargets/{actionTargetId}
//
// Update a single action target.
//
// 
func (a JourneyApi) PatchJourneyActiontarget(actionTargetId string, body Patchactiontarget) (*Actiontarget, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/journey/actiontargets/{actionTargetId}"
	path = strings.Replace(path, "{actionTargetId}", fmt.Sprintf("%v", actionTargetId), -1)
	defaultReturn := new(Actiontarget)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionTargetId' is set
	if &actionTargetId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionTargetId' when calling JourneyApi->PatchJourneyActiontarget")
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

	var successPayload *Actiontarget
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

// PostAnalyticsJourneysAggregatesQuery invokes POST /api/v2/analytics/journeys/aggregates/query
//
// Query for journey aggregates
//
// 
func (a JourneyApi) PostAnalyticsJourneysAggregatesQuery(body Journeyaggregationquery) (*Journeyaggregatequeryresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/analytics/journeys/aggregates/query"
	defaultReturn := new(Journeyaggregatequeryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling JourneyApi->PostAnalyticsJourneysAggregatesQuery")
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

	var successPayload *Journeyaggregatequeryresponse
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

