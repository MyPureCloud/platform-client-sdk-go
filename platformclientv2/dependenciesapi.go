package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
	)

// DependenciesApi provides functions for API endpoints
type DependenciesApi struct {
	Configuration *Configuration
}

// NewDependenciesApi creates an API instance using the default configuration
func NewDependenciesApi() *DependenciesApi {
	fmt.Sprintf(strings.Title(""), "")
	config := GetDefaultConfiguration()
	return &DependenciesApi{
		Configuration: config,
	}
}

// NewDependenciesApiWithConfig creates an API instance using the provided configuration
func NewDependenciesApiWithConfig(config *Configuration) *DependenciesApi {
	return &DependenciesApi{
		Configuration: config,
	}
}

// GetDependenciesTypeEntityTypeIdEntityIdConnectionsRequiredby invokes GET /api/v2/dependencies/type/{entityType}/id/{entityId}/connections/requiredby
//
// Get entities that require the given entity
func (a DependenciesApi) GetDependenciesTypeEntityTypeIdEntityIdConnectionsRequiredby(entityType string, entityId string, pageSize string, beforeSourceType string, beforeSourceId string, afterSourceType string, afterSourceId string) (*Dependencyentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/dependencies/type/{entityType}/id/{entityId}/connections/requiredby"
	path = strings.Replace(path, "{entityType}", url.PathEscape(fmt.Sprintf("%v", entityType)), -1)
	path = strings.Replace(path, "{entityId}", url.PathEscape(fmt.Sprintf("%v", entityId)), -1)
	defaultReturn := new(Dependencyentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'entityType' is set
	if &entityType == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'entityType' when calling DependenciesApi->GetDependenciesTypeEntityTypeIdEntityIdConnectionsRequiredby")
	}
	// verify the required parameter 'entityId' is set
	if &entityId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'entityId' when calling DependenciesApi->GetDependenciesTypeEntityTypeIdEntityIdConnectionsRequiredby")
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
	
	queryParams["beforeSourceType"] = a.Configuration.APIClient.ParameterToString(beforeSourceType, "")
	
	queryParams["beforeSourceId"] = a.Configuration.APIClient.ParameterToString(beforeSourceId, "")
	
	queryParams["afterSourceType"] = a.Configuration.APIClient.ParameterToString(afterSourceType, "")
	
	queryParams["afterSourceId"] = a.Configuration.APIClient.ParameterToString(afterSourceId, "")
	

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
	var successPayload *Dependencyentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dependencyentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetDependenciesTypeEntityTypeIdEntityIdConnectionsRequiredbycounts invokes GET /api/v2/dependencies/type/{entityType}/id/{entityId}/connections/requiredbycounts
//
// An estimated count of entities that depend on this entity, including indirect dependencies.
func (a DependenciesApi) GetDependenciesTypeEntityTypeIdEntityIdConnectionsRequiredbycounts(entityType string, entityId string) (*Dependencycount, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/dependencies/type/{entityType}/id/{entityId}/connections/requiredbycounts"
	path = strings.Replace(path, "{entityType}", url.PathEscape(fmt.Sprintf("%v", entityType)), -1)
	path = strings.Replace(path, "{entityId}", url.PathEscape(fmt.Sprintf("%v", entityId)), -1)
	defaultReturn := new(Dependencycount)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'entityType' is set
	if &entityType == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'entityType' when calling DependenciesApi->GetDependenciesTypeEntityTypeIdEntityIdConnectionsRequiredbycounts")
	}
	// verify the required parameter 'entityId' is set
	if &entityId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'entityId' when calling DependenciesApi->GetDependenciesTypeEntityTypeIdEntityIdConnectionsRequiredbycounts")
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
	var successPayload *Dependencycount
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dependencycount" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetDependenciesTypeEntityTypeIdEntityIdConnectionsRequires invokes GET /api/v2/dependencies/type/{entityType}/id/{entityId}/connections/requires
//
// Get entities that the given entity requires
func (a DependenciesApi) GetDependenciesTypeEntityTypeIdEntityIdConnectionsRequires(entityType string, entityId string, pageSize string, beforeSourceType string, beforeSourceId string, afterSourceType string, afterSourceId string) (*Dependencyentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/dependencies/type/{entityType}/id/{entityId}/connections/requires"
	path = strings.Replace(path, "{entityType}", url.PathEscape(fmt.Sprintf("%v", entityType)), -1)
	path = strings.Replace(path, "{entityId}", url.PathEscape(fmt.Sprintf("%v", entityId)), -1)
	defaultReturn := new(Dependencyentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'entityType' is set
	if &entityType == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'entityType' when calling DependenciesApi->GetDependenciesTypeEntityTypeIdEntityIdConnectionsRequires")
	}
	// verify the required parameter 'entityId' is set
	if &entityId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'entityId' when calling DependenciesApi->GetDependenciesTypeEntityTypeIdEntityIdConnectionsRequires")
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
	
	queryParams["beforeSourceType"] = a.Configuration.APIClient.ParameterToString(beforeSourceType, "")
	
	queryParams["beforeSourceId"] = a.Configuration.APIClient.ParameterToString(beforeSourceId, "")
	
	queryParams["afterSourceType"] = a.Configuration.APIClient.ParameterToString(afterSourceType, "")
	
	queryParams["afterSourceId"] = a.Configuration.APIClient.ParameterToString(afterSourceId, "")
	

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
	var successPayload *Dependencyentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dependencyentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

