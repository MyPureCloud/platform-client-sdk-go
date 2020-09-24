package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

// UserDevelopmentApi provides functions for API endpoints
type UserDevelopmentApi struct {
	Configuration *Configuration
}

// NewUserDevelopmentApi creates an API instance using the default configuration
func NewUserDevelopmentApi() *UserDevelopmentApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating UserDevelopmentApi with base path: %s", strings.ToLower(config.BasePath)))
	return &UserDevelopmentApi{
		Configuration: config,
	}
}

// NewUserDevelopmentApiWithConfig creates an API instance using the provided configuration
func NewUserDevelopmentApiWithConfig(config *Configuration) *UserDevelopmentApi {
	config.Debugf("Creating UserDevelopmentApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &UserDevelopmentApi{
		Configuration: config,
	}
}

// GetUsersDevelopmentActivities invokes GET /api/v2/users/development/activities
//
// Get list of Development Activities
//
// Either moduleId or userId is required. Results are filtered based on the applicable permissions.
func (a UserDevelopmentApi) GetUsersDevelopmentActivities(userId []string, moduleId string, interval string, completionInterval string, overdue string, pageSize int32, pageNumber int32, sortOrder string, types []string, statuses []string, relationship []string) (*Developmentactivitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/users/development/activities"
	defaultReturn := new(Developmentactivitylisting)
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
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range userId {
			queryParams["userId"] = value
		}
	} else {
		queryParams["userId"] = a.Configuration.APIClient.ParameterToString(userId, collectionFormat)
	}
	
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(moduleId).(string); ok {
		if str != "" {
			queryParams["moduleId"] = a.Configuration.APIClient.ParameterToString(moduleId, collectionFormat)
		}
	} else {
		queryParams["moduleId"] = a.Configuration.APIClient.ParameterToString(moduleId, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(interval).(string); ok {
		if str != "" {
			queryParams["interval"] = a.Configuration.APIClient.ParameterToString(interval, collectionFormat)
		}
	} else {
		queryParams["interval"] = a.Configuration.APIClient.ParameterToString(interval, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(completionInterval).(string); ok {
		if str != "" {
			queryParams["completionInterval"] = a.Configuration.APIClient.ParameterToString(completionInterval, collectionFormat)
		}
	} else {
		queryParams["completionInterval"] = a.Configuration.APIClient.ParameterToString(completionInterval, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(overdue).(string); ok {
		if str != "" {
			queryParams["overdue"] = a.Configuration.APIClient.ParameterToString(overdue, collectionFormat)
		}
	} else {
		queryParams["overdue"] = a.Configuration.APIClient.ParameterToString(overdue, collectionFormat)
	}
	
	
	
	
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
	if str, ok := interface{}(sortOrder).(string); ok {
		if str != "" {
			queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
		}
	} else {
		queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
	}
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range types {
			queryParams["types"] = value
		}
	} else {
		queryParams["types"] = a.Configuration.APIClient.ParameterToString(types, collectionFormat)
	}
	
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range statuses {
			queryParams["statuses"] = value
		}
	} else {
		queryParams["statuses"] = a.Configuration.APIClient.ParameterToString(statuses, collectionFormat)
	}
	
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range relationship {
			queryParams["relationship"] = value
		}
	} else {
		queryParams["relationship"] = a.Configuration.APIClient.ParameterToString(relationship, collectionFormat)
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
	var successPayload *Developmentactivitylisting
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

// GetUsersDevelopmentActivitiesMe invokes GET /api/v2/users/development/activities/me
//
// Get list of Development Activities for current user
//
// Results are filtered based on the applicable permissions.
func (a UserDevelopmentApi) GetUsersDevelopmentActivitiesMe(moduleId string, interval string, completionInterval string, overdue string, pageSize int32, pageNumber int32, sortOrder string, types []string, statuses []string, relationship []string) (*Developmentactivitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/users/development/activities/me"
	defaultReturn := new(Developmentactivitylisting)
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
	if str, ok := interface{}(moduleId).(string); ok {
		if str != "" {
			queryParams["moduleId"] = a.Configuration.APIClient.ParameterToString(moduleId, collectionFormat)
		}
	} else {
		queryParams["moduleId"] = a.Configuration.APIClient.ParameterToString(moduleId, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(interval).(string); ok {
		if str != "" {
			queryParams["interval"] = a.Configuration.APIClient.ParameterToString(interval, collectionFormat)
		}
	} else {
		queryParams["interval"] = a.Configuration.APIClient.ParameterToString(interval, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(completionInterval).(string); ok {
		if str != "" {
			queryParams["completionInterval"] = a.Configuration.APIClient.ParameterToString(completionInterval, collectionFormat)
		}
	} else {
		queryParams["completionInterval"] = a.Configuration.APIClient.ParameterToString(completionInterval, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(overdue).(string); ok {
		if str != "" {
			queryParams["overdue"] = a.Configuration.APIClient.ParameterToString(overdue, collectionFormat)
		}
	} else {
		queryParams["overdue"] = a.Configuration.APIClient.ParameterToString(overdue, collectionFormat)
	}
	
	
	
	
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
	if str, ok := interface{}(sortOrder).(string); ok {
		if str != "" {
			queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
		}
	} else {
		queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
	}
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range types {
			queryParams["types"] = value
		}
	} else {
		queryParams["types"] = a.Configuration.APIClient.ParameterToString(types, collectionFormat)
	}
	
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range statuses {
			queryParams["statuses"] = value
		}
	} else {
		queryParams["statuses"] = a.Configuration.APIClient.ParameterToString(statuses, collectionFormat)
	}
	
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range relationship {
			queryParams["relationship"] = value
		}
	} else {
		queryParams["relationship"] = a.Configuration.APIClient.ParameterToString(relationship, collectionFormat)
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
	var successPayload *Developmentactivitylisting
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

// GetUsersDevelopmentActivity invokes GET /api/v2/users/development/activities/{activityId}
//
// Get a Development Activity
//
// 
func (a UserDevelopmentApi) GetUsersDevelopmentActivity(activityId string, varType string) (*Developmentactivity, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/users/development/activities/{activityId}"
	path = strings.Replace(path, "{activityId}", fmt.Sprintf("%v", activityId), -1)
	defaultReturn := new(Developmentactivity)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'activityId' is set
	if &activityId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'activityId' when calling UserDevelopmentApi->GetUsersDevelopmentActivity")
	}
	// verify the required parameter 'varType' is set
	if &varType == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'varType' when calling UserDevelopmentApi->GetUsersDevelopmentActivity")
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
	if str, ok := interface{}(varType).(string); ok {
		if str != "" {
			queryParams["varType"] = a.Configuration.APIClient.ParameterToString(varType, collectionFormat)
		}
	} else {
		queryParams["varType"] = a.Configuration.APIClient.ParameterToString(varType, collectionFormat)
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
	var successPayload *Developmentactivity
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

// PostUsersDevelopmentActivitiesAggregatesQuery invokes POST /api/v2/users/development/activities/aggregates/query
//
// Retrieve aggregated development activity data
//
// Results are filtered based on the applicable permissions.
func (a UserDevelopmentApi) PostUsersDevelopmentActivitiesAggregatesQuery(body Developmentactivityaggregateparam) (*Developmentactivityaggregateresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/users/development/activities/aggregates/query"
	defaultReturn := new(Developmentactivityaggregateresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling UserDevelopmentApi->PostUsersDevelopmentActivitiesAggregatesQuery")
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

	var successPayload *Developmentactivityaggregateresponse
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

