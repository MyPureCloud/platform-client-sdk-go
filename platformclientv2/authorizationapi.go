package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

// AuthorizationApi provides functions for API endpoints
type AuthorizationApi struct {
	Configuration *Configuration
}

// NewAuthorizationApi creates an API instance using the default configuration
func NewAuthorizationApi() *AuthorizationApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating AuthorizationApi with base path: %s", strings.ToLower(config.BasePath)))
	return &AuthorizationApi{
		Configuration: config,
	}
}

// NewAuthorizationApiWithConfig creates an API instance using the provided configuration
func NewAuthorizationApiWithConfig(config *Configuration) *AuthorizationApi {
	config.Debugf("Creating AuthorizationApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &AuthorizationApi{
		Configuration: config,
	}
}

// DeleteAuthorizationDivision invokes DELETE /api/v2/authorization/divisions/{divisionId}
//
// Delete a division.
//
// 
func (a AuthorizationApi) DeleteAuthorizationDivision(divisionId string, force bool) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/divisions/{divisionId}"
	path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'divisionId' is set
	if &divisionId == nil {
		// 
		return nil, errors.New("Missing required parameter 'divisionId' when calling AuthorizationApi->DeleteAuthorizationDivision")
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
	if str, ok := interface{}(force).(string); ok {
		if str != "" {
			queryParams["force"] = a.Configuration.APIClient.ParameterToString(force, collectionFormat)
		}
	} else {
		queryParams["force"] = a.Configuration.APIClient.ParameterToString(force, collectionFormat)
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

// DeleteAuthorizationRole invokes DELETE /api/v2/authorization/roles/{roleId}
//
// Delete an organization role.
//
// 
func (a AuthorizationApi) DeleteAuthorizationRole(roleId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles/{roleId}"
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'roleId' is set
	if &roleId == nil {
		// 
		return nil, errors.New("Missing required parameter 'roleId' when calling AuthorizationApi->DeleteAuthorizationRole")
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

// DeleteAuthorizationSubjectDivisionRole invokes DELETE /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}
//
// Delete a grant of a role in a division
//
// 
func (a AuthorizationApi) DeleteAuthorizationSubjectDivisionRole(subjectId string, divisionId string, roleId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}"
	path = strings.Replace(path, "{subjectId}", fmt.Sprintf("%v", subjectId), -1)
	path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'subjectId' is set
	if &subjectId == nil {
		// 
		return nil, errors.New("Missing required parameter 'subjectId' when calling AuthorizationApi->DeleteAuthorizationSubjectDivisionRole")
	}
	// verify the required parameter 'divisionId' is set
	if &divisionId == nil {
		// 
		return nil, errors.New("Missing required parameter 'divisionId' when calling AuthorizationApi->DeleteAuthorizationSubjectDivisionRole")
	}
	// verify the required parameter 'roleId' is set
	if &roleId == nil {
		// 
		return nil, errors.New("Missing required parameter 'roleId' when calling AuthorizationApi->DeleteAuthorizationSubjectDivisionRole")
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

// GetAuthorizationDivision invokes GET /api/v2/authorization/divisions/{divisionId}
//
// Returns an authorization division.
//
// 
func (a AuthorizationApi) GetAuthorizationDivision(divisionId string, objectCount bool) (*Authzdivision, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/divisions/{divisionId}"
	path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)
	defaultReturn := new(Authzdivision)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'divisionId' is set
	if &divisionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'divisionId' when calling AuthorizationApi->GetAuthorizationDivision")
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
	if str, ok := interface{}(objectCount).(string); ok {
		if str != "" {
			queryParams["objectCount"] = a.Configuration.APIClient.ParameterToString(objectCount, collectionFormat)
		}
	} else {
		queryParams["objectCount"] = a.Configuration.APIClient.ParameterToString(objectCount, collectionFormat)
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
	var successPayload *Authzdivision
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

// GetAuthorizationDivisionGrants invokes GET /api/v2/authorization/divisions/{divisionId}/grants
//
// Gets all grants for a given division.
//
// Returns all grants assigned to a given division. Maximum page size is 500.
func (a AuthorizationApi) GetAuthorizationDivisionGrants(divisionId string, pageNumber int, pageSize int) (*Authzdivisiongrantentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/divisions/{divisionId}/grants"
	path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)
	defaultReturn := new(Authzdivisiongrantentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'divisionId' is set
	if &divisionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'divisionId' when calling AuthorizationApi->GetAuthorizationDivisionGrants")
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
	var successPayload *Authzdivisiongrantentitylisting
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

// GetAuthorizationDivisions invokes GET /api/v2/authorization/divisions
//
// Retrieve a list of all divisions defined for the organization
//
// Request specific divisions by id using a query param \&quot;id\&quot;, e.g.  ?id=5f777167-63be-4c24-ad41-374155d9e28b&amp;id=72e9fb25-c484-488d-9312-7acba82435b3
func (a AuthorizationApi) GetAuthorizationDivisions(pageSize int, pageNumber int, sortBy string, expand []string, nextPage string, previousPage string, objectCount bool, id []string, name string) (*Authzdivisionentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/divisions"
	defaultReturn := new(Authzdivisionentitylisting)
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
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(sortBy).(string); ok {
		if str != "" {
			queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, collectionFormat)
		}
	} else {
		queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, collectionFormat)
	}
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range expand {
			queryParams["expand"] = value
		}
	} else {
		queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
	}
	
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(nextPage).(string); ok {
		if str != "" {
			queryParams["nextPage"] = a.Configuration.APIClient.ParameterToString(nextPage, collectionFormat)
		}
	} else {
		queryParams["nextPage"] = a.Configuration.APIClient.ParameterToString(nextPage, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(previousPage).(string); ok {
		if str != "" {
			queryParams["previousPage"] = a.Configuration.APIClient.ParameterToString(previousPage, collectionFormat)
		}
	} else {
		queryParams["previousPage"] = a.Configuration.APIClient.ParameterToString(previousPage, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(objectCount).(string); ok {
		if str != "" {
			queryParams["objectCount"] = a.Configuration.APIClient.ParameterToString(objectCount, collectionFormat)
		}
	} else {
		queryParams["objectCount"] = a.Configuration.APIClient.ParameterToString(objectCount, collectionFormat)
	}
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range id {
			queryParams["id"] = value
		}
	} else {
		queryParams["id"] = a.Configuration.APIClient.ParameterToString(id, collectionFormat)
	}
	
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(name).(string); ok {
		if str != "" {
			queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
		}
	} else {
		queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
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
	var successPayload *Authzdivisionentitylisting
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

// GetAuthorizationDivisionsHome invokes GET /api/v2/authorization/divisions/home
//
// Retrieve the home division for the organization.
//
// Will not include object counts.
func (a AuthorizationApi) GetAuthorizationDivisionsHome() (*Authzdivision, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/divisions/home"
	defaultReturn := new(Authzdivision)
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
	var successPayload *Authzdivision
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

// GetAuthorizationDivisionsLimit invokes GET /api/v2/authorization/divisions/limit
//
// Returns the maximum allowed number of divisions.
//
// 
func (a AuthorizationApi) GetAuthorizationDivisionsLimit() (*int, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/divisions/limit"
	defaultReturn := new(int)
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
	var successPayload *int
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

// GetAuthorizationDivisionspermittedMe invokes GET /api/v2/authorization/divisionspermitted/me
//
// Returns which divisions the current user has the given permission in.
//
// This route is deprecated, use authorization/divisionspermitted/paged/me instead.
func (a AuthorizationApi) GetAuthorizationDivisionspermittedMe(permission string, name string) ([]Authzdivision, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/divisionspermitted/me"
	defaultReturn := make([]Authzdivision, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'permission' is set
	if &permission == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'permission' when calling AuthorizationApi->GetAuthorizationDivisionspermittedMe")
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
	if str, ok := interface{}(name).(string); ok {
		if str != "" {
			queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
		}
	} else {
		queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(permission).(string); ok {
		if str != "" {
			queryParams["permission"] = a.Configuration.APIClient.ParameterToString(permission, collectionFormat)
		}
	} else {
		queryParams["permission"] = a.Configuration.APIClient.ParameterToString(permission, collectionFormat)
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
	var successPayload []Authzdivision
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

// GetAuthorizationDivisionspermittedPagedMe invokes GET /api/v2/authorization/divisionspermitted/paged/me
//
// Returns which divisions the current user has the given permission in.
//
// 
func (a AuthorizationApi) GetAuthorizationDivisionspermittedPagedMe(permission string, pageNumber int, pageSize int) (*Divspermittedentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/divisionspermitted/paged/me"
	defaultReturn := new(Divspermittedentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'permission' is set
	if &permission == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'permission' when calling AuthorizationApi->GetAuthorizationDivisionspermittedPagedMe")
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
	if str, ok := interface{}(permission).(string); ok {
		if str != "" {
			queryParams["permission"] = a.Configuration.APIClient.ParameterToString(permission, collectionFormat)
		}
	} else {
		queryParams["permission"] = a.Configuration.APIClient.ParameterToString(permission, collectionFormat)
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
	var successPayload *Divspermittedentitylisting
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

// GetAuthorizationDivisionspermittedPagedSubjectId invokes GET /api/v2/authorization/divisionspermitted/paged/{subjectId}
//
// Returns which divisions the specified user has the given permission in.
//
// This route is deprecated, use authorization/divisionspermitted/paged/me instead.
func (a AuthorizationApi) GetAuthorizationDivisionspermittedPagedSubjectId(subjectId string, permission string, pageNumber int, pageSize int) (*Divspermittedentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/divisionspermitted/paged/{subjectId}"
	path = strings.Replace(path, "{subjectId}", fmt.Sprintf("%v", subjectId), -1)
	defaultReturn := new(Divspermittedentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'subjectId' is set
	if &subjectId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'subjectId' when calling AuthorizationApi->GetAuthorizationDivisionspermittedPagedSubjectId")
	}
	// verify the required parameter 'permission' is set
	if &permission == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'permission' when calling AuthorizationApi->GetAuthorizationDivisionspermittedPagedSubjectId")
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
	if str, ok := interface{}(permission).(string); ok {
		if str != "" {
			queryParams["permission"] = a.Configuration.APIClient.ParameterToString(permission, collectionFormat)
		}
	} else {
		queryParams["permission"] = a.Configuration.APIClient.ParameterToString(permission, collectionFormat)
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
	var successPayload *Divspermittedentitylisting
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

// GetAuthorizationPermissions invokes GET /api/v2/authorization/permissions
//
// Get all permissions.
//
// Retrieve a list of all permission defined in the system.
func (a AuthorizationApi) GetAuthorizationPermissions(pageSize int, pageNumber int, queryType string, query string) (*Permissioncollectionentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/permissions"
	defaultReturn := new(Permissioncollectionentitylisting)
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
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(queryType).(string); ok {
		if str != "" {
			queryParams["queryType"] = a.Configuration.APIClient.ParameterToString(queryType, collectionFormat)
		}
	} else {
		queryParams["queryType"] = a.Configuration.APIClient.ParameterToString(queryType, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(query).(string); ok {
		if str != "" {
			queryParams["query"] = a.Configuration.APIClient.ParameterToString(query, collectionFormat)
		}
	} else {
		queryParams["query"] = a.Configuration.APIClient.ParameterToString(query, collectionFormat)
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
	var successPayload *Permissioncollectionentitylisting
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

// GetAuthorizationProducts invokes GET /api/v2/authorization/products
//
// Get the list of enabled products
//
// Gets the list of enabled products. Some example product names are: collaborateFree, collaboratePro, communicate, and engage.
func (a AuthorizationApi) GetAuthorizationProducts() (*Organizationproductentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/products"
	defaultReturn := new(Organizationproductentitylisting)
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
	var successPayload *Organizationproductentitylisting
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

// GetAuthorizationRole invokes GET /api/v2/authorization/roles/{roleId}
//
// Get a single organization role.
//
// Get the organization role specified by its ID.
func (a AuthorizationApi) GetAuthorizationRole(roleId string, expand []string) (*Domainorganizationrole, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles/{roleId}"
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)
	defaultReturn := new(Domainorganizationrole)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'roleId' is set
	if &roleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'roleId' when calling AuthorizationApi->GetAuthorizationRole")
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
		for _, value := range expand {
			queryParams["expand"] = value
		}
	} else {
		queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
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
	var successPayload *Domainorganizationrole
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

// GetAuthorizationRoleComparedefaultRightRoleId invokes GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}
//
// Get an org role to default role comparison
//
// Compares any organization role to a default role id and show differences
func (a AuthorizationApi) GetAuthorizationRoleComparedefaultRightRoleId(leftRoleId string, rightRoleId string) (*Domainorgroledifference, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}"
	path = strings.Replace(path, "{leftRoleId}", fmt.Sprintf("%v", leftRoleId), -1)
	path = strings.Replace(path, "{rightRoleId}", fmt.Sprintf("%v", rightRoleId), -1)
	defaultReturn := new(Domainorgroledifference)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'leftRoleId' is set
	if &leftRoleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'leftRoleId' when calling AuthorizationApi->GetAuthorizationRoleComparedefaultRightRoleId")
	}
	// verify the required parameter 'rightRoleId' is set
	if &rightRoleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'rightRoleId' when calling AuthorizationApi->GetAuthorizationRoleComparedefaultRightRoleId")
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
	var successPayload *Domainorgroledifference
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

// GetAuthorizationRoleSubjectgrants invokes GET /api/v2/authorization/roles/{roleId}/subjectgrants
//
// Get the subjects&#39; granted divisions in the specified role.
//
// Includes the divisions for which the subject has a grant.
func (a AuthorizationApi) GetAuthorizationRoleSubjectgrants(roleId string, pageSize int, pageNumber int, sortBy string, expand []string, nextPage string, previousPage string) (*Subjectdivisiongrantsentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles/{roleId}/subjectgrants"
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)
	defaultReturn := new(Subjectdivisiongrantsentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'roleId' is set
	if &roleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'roleId' when calling AuthorizationApi->GetAuthorizationRoleSubjectgrants")
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
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(sortBy).(string); ok {
		if str != "" {
			queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, collectionFormat)
		}
	} else {
		queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, collectionFormat)
	}
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range expand {
			queryParams["expand"] = value
		}
	} else {
		queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
	}
	
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(nextPage).(string); ok {
		if str != "" {
			queryParams["nextPage"] = a.Configuration.APIClient.ParameterToString(nextPage, collectionFormat)
		}
	} else {
		queryParams["nextPage"] = a.Configuration.APIClient.ParameterToString(nextPage, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(previousPage).(string); ok {
		if str != "" {
			queryParams["previousPage"] = a.Configuration.APIClient.ParameterToString(previousPage, collectionFormat)
		}
	} else {
		queryParams["previousPage"] = a.Configuration.APIClient.ParameterToString(previousPage, collectionFormat)
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
	var successPayload *Subjectdivisiongrantsentitylisting
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

// GetAuthorizationRoleUsers invokes GET /api/v2/authorization/roles/{roleId}/users
//
// Get a list of the users in a specified role.
//
// Get an array of the UUIDs of the users in the specified role.
func (a AuthorizationApi) GetAuthorizationRoleUsers(roleId string, pageSize int, pageNumber int) (*Userentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles/{roleId}/users"
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)
	defaultReturn := new(Userentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'roleId' is set
	if &roleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'roleId' when calling AuthorizationApi->GetAuthorizationRoleUsers")
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
	var successPayload *Userentitylisting
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

// GetAuthorizationRoles invokes GET /api/v2/authorization/roles
//
// Retrieve a list of all roles defined for the organization
//
// 
func (a AuthorizationApi) GetAuthorizationRoles(pageSize int, pageNumber int, sortBy string, expand []string, nextPage string, previousPage string, name string, permission []string, defaultRoleId []string, userCount bool, id []string) (*Organizationroleentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles"
	defaultReturn := new(Organizationroleentitylisting)
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
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(sortBy).(string); ok {
		if str != "" {
			queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, collectionFormat)
		}
	} else {
		queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, collectionFormat)
	}
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range expand {
			queryParams["expand"] = value
		}
	} else {
		queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
	}
	
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(nextPage).(string); ok {
		if str != "" {
			queryParams["nextPage"] = a.Configuration.APIClient.ParameterToString(nextPage, collectionFormat)
		}
	} else {
		queryParams["nextPage"] = a.Configuration.APIClient.ParameterToString(nextPage, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(previousPage).(string); ok {
		if str != "" {
			queryParams["previousPage"] = a.Configuration.APIClient.ParameterToString(previousPage, collectionFormat)
		}
	} else {
		queryParams["previousPage"] = a.Configuration.APIClient.ParameterToString(previousPage, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(name).(string); ok {
		if str != "" {
			queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
		}
	} else {
		queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
	}
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range permission {
			queryParams["permission"] = value
		}
	} else {
		queryParams["permission"] = a.Configuration.APIClient.ParameterToString(permission, collectionFormat)
	}
	
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range defaultRoleId {
			queryParams["defaultRoleId"] = value
		}
	} else {
		queryParams["defaultRoleId"] = a.Configuration.APIClient.ParameterToString(defaultRoleId, collectionFormat)
	}
	
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(userCount).(string); ok {
		if str != "" {
			queryParams["userCount"] = a.Configuration.APIClient.ParameterToString(userCount, collectionFormat)
		}
	} else {
		queryParams["userCount"] = a.Configuration.APIClient.ParameterToString(userCount, collectionFormat)
	}
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range id {
			queryParams["id"] = value
		}
	} else {
		queryParams["id"] = a.Configuration.APIClient.ParameterToString(id, collectionFormat)
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
	var successPayload *Organizationroleentitylisting
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

// GetAuthorizationSubject invokes GET /api/v2/authorization/subjects/{subjectId}
//
// Returns a listing of roles and permissions for a user.
//
// 
func (a AuthorizationApi) GetAuthorizationSubject(subjectId string) (*Authzsubject, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/subjects/{subjectId}"
	path = strings.Replace(path, "{subjectId}", fmt.Sprintf("%v", subjectId), -1)
	defaultReturn := new(Authzsubject)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'subjectId' is set
	if &subjectId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'subjectId' when calling AuthorizationApi->GetAuthorizationSubject")
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
	var successPayload *Authzsubject
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

// GetAuthorizationSubjectsMe invokes GET /api/v2/authorization/subjects/me
//
// Returns a listing of roles and permissions for the currently authenticated user.
//
// 
func (a AuthorizationApi) GetAuthorizationSubjectsMe() (*Authzsubject, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/subjects/me"
	defaultReturn := new(Authzsubject)
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
	var successPayload *Authzsubject
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

// GetAuthorizationSubjectsRolecounts invokes GET /api/v2/authorization/subjects/rolecounts
//
// Get the count of roles granted to a list of subjects
//
// 
func (a AuthorizationApi) GetAuthorizationSubjectsRolecounts(id []string) (*map[string]interface{}, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/subjects/rolecounts"
	defaultReturn := new(map[string]interface{})
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
		for _, value := range id {
			queryParams["id"] = value
		}
	} else {
		queryParams["id"] = a.Configuration.APIClient.ParameterToString(id, collectionFormat)
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
	var successPayload *map[string]interface{}
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

// GetUserRoles invokes GET /api/v2/users/{userId}/roles
//
// Returns a listing of roles and permissions for a user.
//
// 
func (a AuthorizationApi) GetUserRoles(userId string) (*Userauthorization, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/users/{userId}/roles"
	path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)
	defaultReturn := new(Userauthorization)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'userId' is set
	if &userId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling AuthorizationApi->GetUserRoles")
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
	var successPayload *Userauthorization
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

// PatchAuthorizationRole invokes PATCH /api/v2/authorization/roles/{roleId}
//
// Patch Organization Role for needsUpdate Field
//
// Patch Organization Role for needsUpdate Field
func (a AuthorizationApi) PatchAuthorizationRole(roleId string, body Domainorganizationrole) (*Domainorganizationrole, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles/{roleId}"
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)
	defaultReturn := new(Domainorganizationrole)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'roleId' is set
	if &roleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'roleId' when calling AuthorizationApi->PatchAuthorizationRole")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PatchAuthorizationRole")
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

	var successPayload *Domainorganizationrole
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

// PostAuthorizationDivisionObject invokes POST /api/v2/authorization/divisions/{divisionId}/objects/{objectType}
//
// Assign a list of objects to a division
//
// Set the division of a specified list of objects. The objects must all be of the same type, one of:  CAMPAIGN, MANAGEMENTUNIT, FLOW, QUEUE, or USER.  The body of the request is a list of object IDs, which are expected to be  GUIDs, e.g. [\&quot;206ce31f-61ec-40ed-a8b1-be6f06303998\&quot;,\&quot;250a754e-f5e4-4f51-800f-a92f09d3bf8c\&quot;]
func (a AuthorizationApi) PostAuthorizationDivisionObject(divisionId string, objectType string, body []string) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/divisions/{divisionId}/objects/{objectType}"
	path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)
	path = strings.Replace(path, "{objectType}", fmt.Sprintf("%v", objectType), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'divisionId' is set
	if &divisionId == nil {
		// 
		return nil, errors.New("Missing required parameter 'divisionId' when calling AuthorizationApi->PostAuthorizationDivisionObject")
	}
	// verify the required parameter 'objectType' is set
	if &objectType == nil {
		// 
		return nil, errors.New("Missing required parameter 'objectType' when calling AuthorizationApi->PostAuthorizationDivisionObject")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// true
		return nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PostAuthorizationDivisionObject")
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

// PostAuthorizationDivisions invokes POST /api/v2/authorization/divisions
//
// Create a division.
//
// 
func (a AuthorizationApi) PostAuthorizationDivisions(body Authzdivision) (*Authzdivision, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/divisions"
	defaultReturn := new(Authzdivision)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PostAuthorizationDivisions")
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

	var successPayload *Authzdivision
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

// PostAuthorizationRole invokes POST /api/v2/authorization/roles/{roleId}
//
// Bulk-grant subjects and divisions with an organization role.
//
// 
func (a AuthorizationApi) PostAuthorizationRole(roleId string, body Subjectdivisions, subjectType string) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles/{roleId}"
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'roleId' is set
	if &roleId == nil {
		// 
		return nil, errors.New("Missing required parameter 'roleId' when calling AuthorizationApi->PostAuthorizationRole")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PostAuthorizationRole")
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
	if str, ok := interface{}(subjectType).(string); ok {
		if str != "" {
			queryParams["subjectType"] = a.Configuration.APIClient.ParameterToString(subjectType, collectionFormat)
		}
	} else {
		queryParams["subjectType"] = a.Configuration.APIClient.ParameterToString(subjectType, collectionFormat)
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

// PostAuthorizationRoleComparedefaultRightRoleId invokes POST /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}
//
// Get an unsaved org role to default role comparison
//
// Allows users to compare their existing roles in an unsaved state to its default role
func (a AuthorizationApi) PostAuthorizationRoleComparedefaultRightRoleId(leftRoleId string, rightRoleId string, body Domainorganizationrole) (*Domainorgroledifference, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}"
	path = strings.Replace(path, "{leftRoleId}", fmt.Sprintf("%v", leftRoleId), -1)
	path = strings.Replace(path, "{rightRoleId}", fmt.Sprintf("%v", rightRoleId), -1)
	defaultReturn := new(Domainorgroledifference)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'leftRoleId' is set
	if &leftRoleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'leftRoleId' when calling AuthorizationApi->PostAuthorizationRoleComparedefaultRightRoleId")
	}
	// verify the required parameter 'rightRoleId' is set
	if &rightRoleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'rightRoleId' when calling AuthorizationApi->PostAuthorizationRoleComparedefaultRightRoleId")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PostAuthorizationRoleComparedefaultRightRoleId")
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

	var successPayload *Domainorgroledifference
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

// PostAuthorizationRoles invokes POST /api/v2/authorization/roles
//
// Create an organization role.
//
// 
func (a AuthorizationApi) PostAuthorizationRoles(body Domainorganizationrolecreate) (*Domainorganizationrole, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles"
	defaultReturn := new(Domainorganizationrole)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PostAuthorizationRoles")
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

	var successPayload *Domainorganizationrole
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

// PostAuthorizationRolesDefault invokes POST /api/v2/authorization/roles/default
//
// Restores all default roles
//
// This endpoint serves several purposes. 1. It provides the org with default roles. This is important for default roles that will be added after go-live (they can retroactively add the new default-role). Note: When not using a query param of force=true, it only adds the default roles not configured for the org; it does not overwrite roles. 2. Using the query param force=true, you can restore all default roles. Note: This does not have an effect on custom roles.
func (a AuthorizationApi) PostAuthorizationRolesDefault(force bool) (*Organizationroleentitylisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles/default"
	defaultReturn := new(Organizationroleentitylisting)
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
	if str, ok := interface{}(force).(string); ok {
		if str != "" {
			queryParams["force"] = a.Configuration.APIClient.ParameterToString(force, collectionFormat)
		}
	} else {
		queryParams["force"] = a.Configuration.APIClient.ParameterToString(force, collectionFormat)
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
	var successPayload *Organizationroleentitylisting
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

// PostAuthorizationSubjectBulkadd invokes POST /api/v2/authorization/subjects/{subjectId}/bulkadd
//
// Bulk-grant roles and divisions to a subject.
//
// 
func (a AuthorizationApi) PostAuthorizationSubjectBulkadd(subjectId string, body Roledivisiongrants, subjectType string) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/subjects/{subjectId}/bulkadd"
	path = strings.Replace(path, "{subjectId}", fmt.Sprintf("%v", subjectId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'subjectId' is set
	if &subjectId == nil {
		// 
		return nil, errors.New("Missing required parameter 'subjectId' when calling AuthorizationApi->PostAuthorizationSubjectBulkadd")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PostAuthorizationSubjectBulkadd")
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
	if str, ok := interface{}(subjectType).(string); ok {
		if str != "" {
			queryParams["subjectType"] = a.Configuration.APIClient.ParameterToString(subjectType, collectionFormat)
		}
	} else {
		queryParams["subjectType"] = a.Configuration.APIClient.ParameterToString(subjectType, collectionFormat)
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

// PostAuthorizationSubjectBulkremove invokes POST /api/v2/authorization/subjects/{subjectId}/bulkremove
//
// Bulk-remove grants from a subject.
//
// 
func (a AuthorizationApi) PostAuthorizationSubjectBulkremove(subjectId string, body Roledivisiongrants) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/subjects/{subjectId}/bulkremove"
	path = strings.Replace(path, "{subjectId}", fmt.Sprintf("%v", subjectId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'subjectId' is set
	if &subjectId == nil {
		// 
		return nil, errors.New("Missing required parameter 'subjectId' when calling AuthorizationApi->PostAuthorizationSubjectBulkremove")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PostAuthorizationSubjectBulkremove")
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

// PostAuthorizationSubjectDivisionRole invokes POST /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}
//
// Make a grant of a role in a division
//
// 
func (a AuthorizationApi) PostAuthorizationSubjectDivisionRole(subjectId string, divisionId string, roleId string, subjectType string) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}"
	path = strings.Replace(path, "{subjectId}", fmt.Sprintf("%v", subjectId), -1)
	path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'subjectId' is set
	if &subjectId == nil {
		// 
		return nil, errors.New("Missing required parameter 'subjectId' when calling AuthorizationApi->PostAuthorizationSubjectDivisionRole")
	}
	// verify the required parameter 'divisionId' is set
	if &divisionId == nil {
		// 
		return nil, errors.New("Missing required parameter 'divisionId' when calling AuthorizationApi->PostAuthorizationSubjectDivisionRole")
	}
	// verify the required parameter 'roleId' is set
	if &roleId == nil {
		// 
		return nil, errors.New("Missing required parameter 'roleId' when calling AuthorizationApi->PostAuthorizationSubjectDivisionRole")
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
	if str, ok := interface{}(subjectType).(string); ok {
		if str != "" {
			queryParams["subjectType"] = a.Configuration.APIClient.ParameterToString(subjectType, collectionFormat)
		}
	} else {
		queryParams["subjectType"] = a.Configuration.APIClient.ParameterToString(subjectType, collectionFormat)
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

// PutAuthorizationDivision invokes PUT /api/v2/authorization/divisions/{divisionId}
//
// Update a division.
//
// 
func (a AuthorizationApi) PutAuthorizationDivision(divisionId string, body Authzdivision) (*Authzdivision, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/divisions/{divisionId}"
	path = strings.Replace(path, "{divisionId}", fmt.Sprintf("%v", divisionId), -1)
	defaultReturn := new(Authzdivision)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'divisionId' is set
	if &divisionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'divisionId' when calling AuthorizationApi->PutAuthorizationDivision")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PutAuthorizationDivision")
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

	var successPayload *Authzdivision
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

// PutAuthorizationRole invokes PUT /api/v2/authorization/roles/{roleId}
//
// Update an organization role.
//
// Update
func (a AuthorizationApi) PutAuthorizationRole(roleId string, body Domainorganizationroleupdate) (*Domainorganizationrole, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles/{roleId}"
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)
	defaultReturn := new(Domainorganizationrole)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'roleId' is set
	if &roleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'roleId' when calling AuthorizationApi->PutAuthorizationRole")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PutAuthorizationRole")
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

	var successPayload *Domainorganizationrole
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

// PutAuthorizationRoleUsersAdd invokes PUT /api/v2/authorization/roles/{roleId}/users/add
//
// Sets the users for the role
//
// 
func (a AuthorizationApi) PutAuthorizationRoleUsersAdd(roleId string, body []string) ([]string, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles/{roleId}/users/add"
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)
	defaultReturn := make([]string, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'roleId' is set
	if &roleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'roleId' when calling AuthorizationApi->PutAuthorizationRoleUsersAdd")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// true
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PutAuthorizationRoleUsersAdd")
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

	var successPayload []string
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

// PutAuthorizationRoleUsersRemove invokes PUT /api/v2/authorization/roles/{roleId}/users/remove
//
// Removes the users from the role
//
// 
func (a AuthorizationApi) PutAuthorizationRoleUsersRemove(roleId string, body []string) ([]string, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles/{roleId}/users/remove"
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleId), -1)
	defaultReturn := make([]string, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'roleId' is set
	if &roleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'roleId' when calling AuthorizationApi->PutAuthorizationRoleUsersRemove")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// true
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PutAuthorizationRoleUsersRemove")
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

	var successPayload []string
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

// PutAuthorizationRolesDefault invokes PUT /api/v2/authorization/roles/default
//
// Restore specified default roles
//
// 
func (a AuthorizationApi) PutAuthorizationRolesDefault(body []Domainorganizationrole) (*Organizationroleentitylisting, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/authorization/roles/default"
	defaultReturn := new(Organizationroleentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// true
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PutAuthorizationRolesDefault")
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

	var successPayload *Organizationroleentitylisting
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

// PutUserRoles invokes PUT /api/v2/users/{userId}/roles
//
// Sets the user&#39;s roles
//
// 
func (a AuthorizationApi) PutUserRoles(userId string, body []string) (*Userauthorization, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/users/{userId}/roles"
	path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)
	defaultReturn := new(Userauthorization)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'userId' is set
	if &userId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling AuthorizationApi->PutUserRoles")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// true
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling AuthorizationApi->PutUserRoles")
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

	var successPayload *Userauthorization
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

