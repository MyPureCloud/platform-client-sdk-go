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

// GetJourneyCustomerCustomerIdSegments invokes GET /api/v2/journey/customers/{customerIdType}/{customerId}/segments
//
// Retrieve segment assignments by customer ID.
//
// 
func (a JourneyApi) GetJourneyCustomerCustomerIdSegments(customerIdType string, customerId string, pageSize string, after string, segmentScope string, assignmentState string) (*Segmentassignmentlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/journey/customers/{customerIdType}/{customerId}/segments"
	path = strings.Replace(path, "{customerIdType}", fmt.Sprintf("%v", customerIdType), -1)
	path = strings.Replace(path, "{customerId}", fmt.Sprintf("%v", customerId), -1)
	defaultReturn := new(Segmentassignmentlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'customerIdType' is set
	if &customerIdType == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'customerIdType' when calling JourneyApi->GetJourneyCustomerCustomerIdSegments")
	}
	// verify the required parameter 'customerId' is set
	if &customerId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'customerId' when calling JourneyApi->GetJourneyCustomerCustomerIdSegments")
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
	if str, ok := interface{}(after).(string); ok {
		if str != "" {
			queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
		}
	} else {
		queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(segmentScope).(string); ok {
		if str != "" {
			queryParams["segmentScope"] = a.Configuration.APIClient.ParameterToString(segmentScope, collectionFormat)
		}
	} else {
		queryParams["segmentScope"] = a.Configuration.APIClient.ParameterToString(segmentScope, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(assignmentState).(string); ok {
		if str != "" {
			queryParams["assignmentState"] = a.Configuration.APIClient.ParameterToString(assignmentState, collectionFormat)
		}
	} else {
		queryParams["assignmentState"] = a.Configuration.APIClient.ParameterToString(assignmentState, collectionFormat)
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
	var successPayload *Segmentassignmentlisting
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

// GetJourneyExternalcontactSegments invokes GET /api/v2/journey/externalcontacts/{externalContactId}/segments
//
// Retrieve segment assignments by external contact ID.
//
// 
func (a JourneyApi) GetJourneyExternalcontactSegments(externalContactId string, pageSize string, after string, segmentScope string, assignmentState string) (*Segmentassignmentlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/journey/externalcontacts/{externalContactId}/segments"
	path = strings.Replace(path, "{externalContactId}", fmt.Sprintf("%v", externalContactId), -1)
	defaultReturn := new(Segmentassignmentlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'externalContactId' is set
	if &externalContactId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'externalContactId' when calling JourneyApi->GetJourneyExternalcontactSegments")
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
	if str, ok := interface{}(after).(string); ok {
		if str != "" {
			queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
		}
	} else {
		queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(segmentScope).(string); ok {
		if str != "" {
			queryParams["segmentScope"] = a.Configuration.APIClient.ParameterToString(segmentScope, collectionFormat)
		}
	} else {
		queryParams["segmentScope"] = a.Configuration.APIClient.ParameterToString(segmentScope, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(assignmentState).(string); ok {
		if str != "" {
			queryParams["assignmentState"] = a.Configuration.APIClient.ParameterToString(assignmentState, collectionFormat)
		}
	} else {
		queryParams["assignmentState"] = a.Configuration.APIClient.ParameterToString(assignmentState, collectionFormat)
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
	var successPayload *Segmentassignmentlisting
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

// GetJourneySessionSegments invokes GET /api/v2/journey/sessions/{sessionId}/segments
//
// Retrieve segment assignments by session ID.
//
// 
func (a JourneyApi) GetJourneySessionSegments(sessionId string, pageSize string, after string, segmentScope string, assignmentState string) (*Segmentassignmentlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/journey/sessions/{sessionId}/segments"
	path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)
	defaultReturn := new(Segmentassignmentlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sessionId' is set
	if &sessionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'sessionId' when calling JourneyApi->GetJourneySessionSegments")
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
	if str, ok := interface{}(after).(string); ok {
		if str != "" {
			queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
		}
	} else {
		queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(segmentScope).(string); ok {
		if str != "" {
			queryParams["segmentScope"] = a.Configuration.APIClient.ParameterToString(segmentScope, collectionFormat)
		}
	} else {
		queryParams["segmentScope"] = a.Configuration.APIClient.ParameterToString(segmentScope, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(assignmentState).(string); ok {
		if str != "" {
			queryParams["assignmentState"] = a.Configuration.APIClient.ParameterToString(assignmentState, collectionFormat)
		}
	} else {
		queryParams["assignmentState"] = a.Configuration.APIClient.ParameterToString(assignmentState, collectionFormat)
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
	var successPayload *Segmentassignmentlisting
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

// PostJourneyExternalcontactSegments invokes POST /api/v2/journey/externalcontacts/{externalContactId}/segments
//
// Assign/Unassign a segment to/from an external contact or, if a segment is already assigned, update the expiry date of the segment assignment.
//
// 
func (a JourneyApi) PostJourneyExternalcontactSegments(externalContactId string, body []Segmentassignmentsupdate) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/journey/externalcontacts/{externalContactId}/segments"
	path = strings.Replace(path, "{externalContactId}", fmt.Sprintf("%v", externalContactId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'externalContactId' is set
	if &externalContactId == nil {
		// 
		return nil, errors.New("Missing required parameter 'externalContactId' when calling JourneyApi->PostJourneyExternalcontactSegments")
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


	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	}
	return response, err
}

