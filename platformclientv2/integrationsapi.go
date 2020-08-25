package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

// IntegrationsApi provides functions for API endpoints
type IntegrationsApi struct {
	Configuration *Configuration
}

// NewIntegrationsApi creates an API instance using the default configuration
func NewIntegrationsApi() *IntegrationsApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating IntegrationsApi with base path: %s", strings.ToLower(config.BasePath)))
	return &IntegrationsApi{
		Configuration: config,
	}
}

// NewIntegrationsApiWithConfig creates an API instance using the provided configuration
func NewIntegrationsApiWithConfig(config *Configuration) *IntegrationsApi {
	config.Debugf("Creating IntegrationsApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &IntegrationsApi{
		Configuration: config,
	}
}

// DeleteIntegration invokes DELETE /api/v2/integrations/{integrationId}
//
// Delete integration.
//
// 
func (a IntegrationsApi) DeleteIntegration(integrationId string) (*Integration, *APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/{integrationId}"
	path = strings.Replace(path, "{integrationId}", fmt.Sprintf("%v", integrationId), -1)
	defaultReturn := new(Integration)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'integrationId' is set
	if &integrationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'integrationId' when calling IntegrationsApi->DeleteIntegration")
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
	var successPayload *Integration
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

// DeleteIntegrationsAction invokes DELETE /api/v2/integrations/actions/{actionId}
//
// Delete an Action
//
// 
func (a IntegrationsApi) DeleteIntegrationsAction(actionId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->DeleteIntegrationsAction")
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

// DeleteIntegrationsActionDraft invokes DELETE /api/v2/integrations/actions/{actionId}/draft
//
// Delete a Draft
//
// 
func (a IntegrationsApi) DeleteIntegrationsActionDraft(actionId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}/draft"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->DeleteIntegrationsActionDraft")
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

// DeleteIntegrationsCredential invokes DELETE /api/v2/integrations/credentials/{credentialId}
//
// Delete a set of credentials
//
// 
func (a IntegrationsApi) DeleteIntegrationsCredential(credentialId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/credentials/{credentialId}"
	path = strings.Replace(path, "{credentialId}", fmt.Sprintf("%v", credentialId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'credentialId' is set
	if &credentialId == nil {
		// 
		return nil, errors.New("Missing required parameter 'credentialId' when calling IntegrationsApi->DeleteIntegrationsCredential")
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

// GetIntegration invokes GET /api/v2/integrations/{integrationId}
//
// Get integration.
//
// 
func (a IntegrationsApi) GetIntegration(integrationId string, pageSize int32, pageNumber int32, sortBy string, expand []string, nextPage string, previousPage string) (*Integration, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/{integrationId}"
	path = strings.Replace(path, "{integrationId}", fmt.Sprintf("%v", integrationId), -1)
	defaultReturn := new(Integration)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'integrationId' is set
	if &integrationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'integrationId' when calling IntegrationsApi->GetIntegration")
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
	var successPayload *Integration
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

// GetIntegrationConfigCurrent invokes GET /api/v2/integrations/{integrationId}/config/current
//
// Get integration configuration.
//
// 
func (a IntegrationsApi) GetIntegrationConfigCurrent(integrationId string) (*Integrationconfiguration, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/{integrationId}/config/current"
	path = strings.Replace(path, "{integrationId}", fmt.Sprintf("%v", integrationId), -1)
	defaultReturn := new(Integrationconfiguration)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'integrationId' is set
	if &integrationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'integrationId' when calling IntegrationsApi->GetIntegrationConfigCurrent")
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
	var successPayload *Integrationconfiguration
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

// GetIntegrations invokes GET /api/v2/integrations
//
// List integrations
//
// 
func (a IntegrationsApi) GetIntegrations(pageSize int32, pageNumber int32, sortBy string, expand []string, nextPage string, previousPage string) (*Integrationentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations"
	defaultReturn := new(Integrationentitylisting)
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
	var successPayload *Integrationentitylisting
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

// GetIntegrationsAction invokes GET /api/v2/integrations/actions/{actionId}
//
// Retrieves a single Action matching id.
//
// 
func (a IntegrationsApi) GetIntegrationsAction(actionId string, expand string, includeConfig bool) (*Action, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	defaultReturn := new(Action)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->GetIntegrationsAction")
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
	if str, ok := interface{}(expand).(string); ok {
		if str != "" {
			queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
		}
	} else {
		queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(includeConfig).(string); ok {
		if str != "" {
			queryParams["includeConfig"] = a.Configuration.APIClient.ParameterToString(includeConfig, collectionFormat)
		}
	} else {
		queryParams["includeConfig"] = a.Configuration.APIClient.ParameterToString(includeConfig, collectionFormat)
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
	var successPayload *Action
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

// GetIntegrationsActionDraft invokes GET /api/v2/integrations/actions/{actionId}/draft
//
// Retrieve a Draft
//
// 
func (a IntegrationsApi) GetIntegrationsActionDraft(actionId string, expand string, includeConfig bool) (*Action, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}/draft"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	defaultReturn := new(Action)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->GetIntegrationsActionDraft")
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
	if str, ok := interface{}(expand).(string); ok {
		if str != "" {
			queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
		}
	} else {
		queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(includeConfig).(string); ok {
		if str != "" {
			queryParams["includeConfig"] = a.Configuration.APIClient.ParameterToString(includeConfig, collectionFormat)
		}
	} else {
		queryParams["includeConfig"] = a.Configuration.APIClient.ParameterToString(includeConfig, collectionFormat)
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
	var successPayload *Action
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

// GetIntegrationsActionDraftSchema invokes GET /api/v2/integrations/actions/{actionId}/draft/schemas/{fileName}
//
// Retrieve schema for a Draft based on filename.
//
// 
func (a IntegrationsApi) GetIntegrationsActionDraftSchema(actionId string, fileName string) (*Jsonschemadocument, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}/draft/schemas/{fileName}"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	path = strings.Replace(path, "{fileName}", fmt.Sprintf("%v", fileName), -1)
	defaultReturn := new(Jsonschemadocument)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->GetIntegrationsActionDraftSchema")
	}
	// verify the required parameter 'fileName' is set
	if &fileName == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'fileName' when calling IntegrationsApi->GetIntegrationsActionDraftSchema")
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
	var successPayload *Jsonschemadocument
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

// GetIntegrationsActionDraftTemplate invokes GET /api/v2/integrations/actions/{actionId}/draft/templates/{fileName}
//
// Retrieve templates for a Draft based on filename.
//
// 
func (a IntegrationsApi) GetIntegrationsActionDraftTemplate(actionId string, fileName string) (*string, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}/draft/templates/{fileName}"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	path = strings.Replace(path, "{fileName}", fmt.Sprintf("%v", fileName), -1)
	defaultReturn := new(string)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->GetIntegrationsActionDraftTemplate")
	}
	// verify the required parameter 'fileName' is set
	if &fileName == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'fileName' when calling IntegrationsApi->GetIntegrationsActionDraftTemplate")
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
		"text/plain",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *string
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

// GetIntegrationsActionDraftValidation invokes GET /api/v2/integrations/actions/{actionId}/draft/validation
//
// Validate current Draft configuration.
//
// 
func (a IntegrationsApi) GetIntegrationsActionDraftValidation(actionId string) (*Draftvalidationresult, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}/draft/validation"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	defaultReturn := new(Draftvalidationresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->GetIntegrationsActionDraftValidation")
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
	var successPayload *Draftvalidationresult
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

// GetIntegrationsActionSchema invokes GET /api/v2/integrations/actions/{actionId}/schemas/{fileName}
//
// Retrieve schema for an action based on filename.
//
// 
func (a IntegrationsApi) GetIntegrationsActionSchema(actionId string, fileName string) (*Jsonschemadocument, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}/schemas/{fileName}"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	path = strings.Replace(path, "{fileName}", fmt.Sprintf("%v", fileName), -1)
	defaultReturn := new(Jsonschemadocument)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->GetIntegrationsActionSchema")
	}
	// verify the required parameter 'fileName' is set
	if &fileName == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'fileName' when calling IntegrationsApi->GetIntegrationsActionSchema")
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
	var successPayload *Jsonschemadocument
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

// GetIntegrationsActionTemplate invokes GET /api/v2/integrations/actions/{actionId}/templates/{fileName}
//
// Retrieve text of templates for an action based on filename.
//
// 
func (a IntegrationsApi) GetIntegrationsActionTemplate(actionId string, fileName string) (*string, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}/templates/{fileName}"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	path = strings.Replace(path, "{fileName}", fmt.Sprintf("%v", fileName), -1)
	defaultReturn := new(string)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->GetIntegrationsActionTemplate")
	}
	// verify the required parameter 'fileName' is set
	if &fileName == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'fileName' when calling IntegrationsApi->GetIntegrationsActionTemplate")
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
		"text/plain",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *string
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

// GetIntegrationsActions invokes GET /api/v2/integrations/actions
//
// Retrieves all actions associated with filters passed in via query param.
//
// 
func (a IntegrationsApi) GetIntegrationsActions(pageSize int32, pageNumber int32, nextPage string, previousPage string, sortBy string, sortOrder string, category string, name string, secure string, includeAuthActions string) (*Actionentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions"
	defaultReturn := new(Actionentitylisting)
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
	if str, ok := interface{}(sortBy).(string); ok {
		if str != "" {
			queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, collectionFormat)
		}
	} else {
		queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(sortOrder).(string); ok {
		if str != "" {
			queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
		}
	} else {
		queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(category).(string); ok {
		if str != "" {
			queryParams["category"] = a.Configuration.APIClient.ParameterToString(category, collectionFormat)
		}
	} else {
		queryParams["category"] = a.Configuration.APIClient.ParameterToString(category, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(name).(string); ok {
		if str != "" {
			queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
		}
	} else {
		queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(secure).(string); ok {
		if str != "" {
			queryParams["secure"] = a.Configuration.APIClient.ParameterToString(secure, collectionFormat)
		}
	} else {
		queryParams["secure"] = a.Configuration.APIClient.ParameterToString(secure, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(includeAuthActions).(string); ok {
		if str != "" {
			queryParams["includeAuthActions"] = a.Configuration.APIClient.ParameterToString(includeAuthActions, collectionFormat)
		}
	} else {
		queryParams["includeAuthActions"] = a.Configuration.APIClient.ParameterToString(includeAuthActions, collectionFormat)
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
	var successPayload *Actionentitylisting
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

// GetIntegrationsActionsCategories invokes GET /api/v2/integrations/actions/categories
//
// Retrieves all categories of available Actions
//
// 
func (a IntegrationsApi) GetIntegrationsActionsCategories(pageSize int32, pageNumber int32, nextPage string, previousPage string, sortBy string, sortOrder string, secure string) (*Categoryentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/categories"
	defaultReturn := new(Categoryentitylisting)
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
	if str, ok := interface{}(sortBy).(string); ok {
		if str != "" {
			queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, collectionFormat)
		}
	} else {
		queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(sortOrder).(string); ok {
		if str != "" {
			queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
		}
	} else {
		queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(secure).(string); ok {
		if str != "" {
			queryParams["secure"] = a.Configuration.APIClient.ParameterToString(secure, collectionFormat)
		}
	} else {
		queryParams["secure"] = a.Configuration.APIClient.ParameterToString(secure, collectionFormat)
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
	var successPayload *Categoryentitylisting
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

// GetIntegrationsActionsDrafts invokes GET /api/v2/integrations/actions/drafts
//
// Retrieves all action drafts associated with the filters passed in via query param.
//
// 
func (a IntegrationsApi) GetIntegrationsActionsDrafts(pageSize int32, pageNumber int32, nextPage string, previousPage string, sortBy string, sortOrder string, category string, name string, secure string, includeAuthActions string) (*Actionentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/drafts"
	defaultReturn := new(Actionentitylisting)
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
	if str, ok := interface{}(sortBy).(string); ok {
		if str != "" {
			queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, collectionFormat)
		}
	} else {
		queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(sortOrder).(string); ok {
		if str != "" {
			queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
		}
	} else {
		queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(category).(string); ok {
		if str != "" {
			queryParams["category"] = a.Configuration.APIClient.ParameterToString(category, collectionFormat)
		}
	} else {
		queryParams["category"] = a.Configuration.APIClient.ParameterToString(category, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(name).(string); ok {
		if str != "" {
			queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
		}
	} else {
		queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(secure).(string); ok {
		if str != "" {
			queryParams["secure"] = a.Configuration.APIClient.ParameterToString(secure, collectionFormat)
		}
	} else {
		queryParams["secure"] = a.Configuration.APIClient.ParameterToString(secure, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(includeAuthActions).(string); ok {
		if str != "" {
			queryParams["includeAuthActions"] = a.Configuration.APIClient.ParameterToString(includeAuthActions, collectionFormat)
		}
	} else {
		queryParams["includeAuthActions"] = a.Configuration.APIClient.ParameterToString(includeAuthActions, collectionFormat)
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
	var successPayload *Actionentitylisting
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

// GetIntegrationsClientapps invokes GET /api/v2/integrations/clientapps
//
// List permitted client app integrations for the logged in user
//
// 
func (a IntegrationsApi) GetIntegrationsClientapps(pageSize int32, pageNumber int32, sortBy string, expand []string, nextPage string, previousPage string) (*Clientappentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/clientapps"
	defaultReturn := new(Clientappentitylisting)
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
	var successPayload *Clientappentitylisting
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

// GetIntegrationsCredential invokes GET /api/v2/integrations/credentials/{credentialId}
//
// Get a single credential with sensitive fields redacted
//
// 
func (a IntegrationsApi) GetIntegrationsCredential(credentialId string) (*Credential, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/credentials/{credentialId}"
	path = strings.Replace(path, "{credentialId}", fmt.Sprintf("%v", credentialId), -1)
	defaultReturn := new(Credential)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'credentialId' is set
	if &credentialId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'credentialId' when calling IntegrationsApi->GetIntegrationsCredential")
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
	var successPayload *Credential
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

// GetIntegrationsCredentials invokes GET /api/v2/integrations/credentials
//
// List multiple sets of credentials
//
// 
func (a IntegrationsApi) GetIntegrationsCredentials(pageNumber int32, pageSize int32) (*Credentialinfolisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/credentials"
	defaultReturn := new(Credentialinfolisting)
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
	var successPayload *Credentialinfolisting
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

// GetIntegrationsCredentialsTypes invokes GET /api/v2/integrations/credentials/types
//
// List all credential types
//
// 
func (a IntegrationsApi) GetIntegrationsCredentialsTypes() (*Credentialtypelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/credentials/types"
	defaultReturn := new(Credentialtypelisting)
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
	var successPayload *Credentialtypelisting
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

// GetIntegrationsEventlog invokes GET /api/v2/integrations/eventlog
//
// List all events
//
// 
func (a IntegrationsApi) GetIntegrationsEventlog(pageSize int32, pageNumber int32, sortBy string, sortOrder string, entityId string) (*Integrationevententitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/eventlog"
	defaultReturn := new(Integrationevententitylisting)
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
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(sortOrder).(string); ok {
		if str != "" {
			queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
		}
	} else {
		queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(entityId).(string); ok {
		if str != "" {
			queryParams["entityId"] = a.Configuration.APIClient.ParameterToString(entityId, collectionFormat)
		}
	} else {
		queryParams["entityId"] = a.Configuration.APIClient.ParameterToString(entityId, collectionFormat)
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
	var successPayload *Integrationevententitylisting
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

// GetIntegrationsEventlogEventId invokes GET /api/v2/integrations/eventlog/{eventId}
//
// Get a single event
//
// 
func (a IntegrationsApi) GetIntegrationsEventlogEventId(eventId string) (*Integrationevent, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/eventlog/{eventId}"
	path = strings.Replace(path, "{eventId}", fmt.Sprintf("%v", eventId), -1)
	defaultReturn := new(Integrationevent)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'eventId' is set
	if &eventId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'eventId' when calling IntegrationsApi->GetIntegrationsEventlogEventId")
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
	var successPayload *Integrationevent
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

// GetIntegrationsSpeechDialogflowAgent invokes GET /api/v2/integrations/speech/dialogflow/agents/{agentId}
//
// Get details about a Dialogflow agent
//
// 
func (a IntegrationsApi) GetIntegrationsSpeechDialogflowAgent(agentId string) (*Dialogflowagent, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/speech/dialogflow/agents/{agentId}"
	path = strings.Replace(path, "{agentId}", fmt.Sprintf("%v", agentId), -1)
	defaultReturn := new(Dialogflowagent)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'agentId' is set
	if &agentId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'agentId' when calling IntegrationsApi->GetIntegrationsSpeechDialogflowAgent")
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
	var successPayload *Dialogflowagent
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

// GetIntegrationsSpeechDialogflowAgents invokes GET /api/v2/integrations/speech/dialogflow/agents
//
// Get a list of Dialogflow agents in the customers&#39; Google accounts
//
// 
func (a IntegrationsApi) GetIntegrationsSpeechDialogflowAgents(pageNumber int32, pageSize int32, name string) (*Dialogflowagentsummaryentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/speech/dialogflow/agents"
	defaultReturn := new(Dialogflowagentsummaryentitylisting)
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
	var successPayload *Dialogflowagentsummaryentitylisting
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

// GetIntegrationsSpeechLexBotAlias invokes GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}
//
// Get details about a Lex bot alias
//
// 
func (a IntegrationsApi) GetIntegrationsSpeechLexBotAlias(aliasId string) (*Lexbotalias, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/speech/lex/bot/alias/{aliasId}"
	path = strings.Replace(path, "{aliasId}", fmt.Sprintf("%v", aliasId), -1)
	defaultReturn := new(Lexbotalias)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'aliasId' is set
	if &aliasId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'aliasId' when calling IntegrationsApi->GetIntegrationsSpeechLexBotAlias")
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
	var successPayload *Lexbotalias
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

// GetIntegrationsSpeechLexBotBotIdAliases invokes GET /api/v2/integrations/speech/lex/bot/{botId}/aliases
//
// Get a list of aliases for a bot in the customer&#39;s AWS accounts
//
// 
func (a IntegrationsApi) GetIntegrationsSpeechLexBotBotIdAliases(botId string, pageNumber int32, pageSize int32, status string, name string) (*Lexbotaliasentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/speech/lex/bot/{botId}/aliases"
	path = strings.Replace(path, "{botId}", fmt.Sprintf("%v", botId), -1)
	defaultReturn := new(Lexbotaliasentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'botId' is set
	if &botId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'botId' when calling IntegrationsApi->GetIntegrationsSpeechLexBotBotIdAliases")
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
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(status).(string); ok {
		if str != "" {
			queryParams["status"] = a.Configuration.APIClient.ParameterToString(status, collectionFormat)
		}
	} else {
		queryParams["status"] = a.Configuration.APIClient.ParameterToString(status, collectionFormat)
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
	var successPayload *Lexbotaliasentitylisting
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

// GetIntegrationsSpeechLexBots invokes GET /api/v2/integrations/speech/lex/bots
//
// Get a list of Lex bots in the customers&#39; AWS accounts
//
// 
func (a IntegrationsApi) GetIntegrationsSpeechLexBots(pageNumber int32, pageSize int32, name string) (*Lexbotentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/speech/lex/bots"
	defaultReturn := new(Lexbotentitylisting)
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
	var successPayload *Lexbotentitylisting
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

// GetIntegrationsSpeechTtsEngine invokes GET /api/v2/integrations/speech/tts/engines/{engineId}
//
// Get details about a TTS engine
//
// 
func (a IntegrationsApi) GetIntegrationsSpeechTtsEngine(engineId string, includeVoices bool) (*Ttsengineentity, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/speech/tts/engines/{engineId}"
	path = strings.Replace(path, "{engineId}", fmt.Sprintf("%v", engineId), -1)
	defaultReturn := new(Ttsengineentity)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'engineId' is set
	if &engineId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'engineId' when calling IntegrationsApi->GetIntegrationsSpeechTtsEngine")
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
	if str, ok := interface{}(includeVoices).(string); ok {
		if str != "" {
			queryParams["includeVoices"] = a.Configuration.APIClient.ParameterToString(includeVoices, collectionFormat)
		}
	} else {
		queryParams["includeVoices"] = a.Configuration.APIClient.ParameterToString(includeVoices, collectionFormat)
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
	var successPayload *Ttsengineentity
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

// GetIntegrationsSpeechTtsEngineVoice invokes GET /api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}
//
// Get details about a specific voice for a TTS engine
//
// 
func (a IntegrationsApi) GetIntegrationsSpeechTtsEngineVoice(engineId string, voiceId string) (*Ttsvoiceentity, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}"
	path = strings.Replace(path, "{engineId}", fmt.Sprintf("%v", engineId), -1)
	path = strings.Replace(path, "{voiceId}", fmt.Sprintf("%v", voiceId), -1)
	defaultReturn := new(Ttsvoiceentity)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'engineId' is set
	if &engineId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'engineId' when calling IntegrationsApi->GetIntegrationsSpeechTtsEngineVoice")
	}
	// verify the required parameter 'voiceId' is set
	if &voiceId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'voiceId' when calling IntegrationsApi->GetIntegrationsSpeechTtsEngineVoice")
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
	var successPayload *Ttsvoiceentity
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

// GetIntegrationsSpeechTtsEngineVoices invokes GET /api/v2/integrations/speech/tts/engines/{engineId}/voices
//
// Get a list of voices for a TTS engine
//
// 
func (a IntegrationsApi) GetIntegrationsSpeechTtsEngineVoices(engineId string, pageNumber int32, pageSize int32) (*Ttsvoiceentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/speech/tts/engines/{engineId}/voices"
	path = strings.Replace(path, "{engineId}", fmt.Sprintf("%v", engineId), -1)
	defaultReturn := new(Ttsvoiceentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'engineId' is set
	if &engineId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'engineId' when calling IntegrationsApi->GetIntegrationsSpeechTtsEngineVoices")
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
	var successPayload *Ttsvoiceentitylisting
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

// GetIntegrationsSpeechTtsEngines invokes GET /api/v2/integrations/speech/tts/engines
//
// Get a list of TTS engines enabled for org
//
// 
func (a IntegrationsApi) GetIntegrationsSpeechTtsEngines(pageNumber int32, pageSize int32, includeVoices bool, name string, language string) (*Ttsengineentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/speech/tts/engines"
	defaultReturn := new(Ttsengineentitylisting)
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
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(includeVoices).(string); ok {
		if str != "" {
			queryParams["includeVoices"] = a.Configuration.APIClient.ParameterToString(includeVoices, collectionFormat)
		}
	} else {
		queryParams["includeVoices"] = a.Configuration.APIClient.ParameterToString(includeVoices, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(name).(string); ok {
		if str != "" {
			queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
		}
	} else {
		queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(language).(string); ok {
		if str != "" {
			queryParams["language"] = a.Configuration.APIClient.ParameterToString(language, collectionFormat)
		}
	} else {
		queryParams["language"] = a.Configuration.APIClient.ParameterToString(language, collectionFormat)
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
	var successPayload *Ttsengineentitylisting
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

// GetIntegrationsSpeechTtsSettings invokes GET /api/v2/integrations/speech/tts/settings
//
// Get TTS settings for an org
//
// 
func (a IntegrationsApi) GetIntegrationsSpeechTtsSettings() (*Ttssettings, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/speech/tts/settings"
	defaultReturn := new(Ttssettings)
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
	var successPayload *Ttssettings
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

// GetIntegrationsType invokes GET /api/v2/integrations/types/{typeId}
//
// Get integration type.
//
// 
func (a IntegrationsApi) GetIntegrationsType(typeId string) (*Integrationtype, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/types/{typeId}"
	path = strings.Replace(path, "{typeId}", fmt.Sprintf("%v", typeId), -1)
	defaultReturn := new(Integrationtype)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'typeId' is set
	if &typeId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'typeId' when calling IntegrationsApi->GetIntegrationsType")
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
	var successPayload *Integrationtype
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

// GetIntegrationsTypeConfigschema invokes GET /api/v2/integrations/types/{typeId}/configschemas/{configType}
//
// Get properties config schema for an integration type.
//
// 
func (a IntegrationsApi) GetIntegrationsTypeConfigschema(typeId string, configType string) (*Jsonschemadocument, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/types/{typeId}/configschemas/{configType}"
	path = strings.Replace(path, "{typeId}", fmt.Sprintf("%v", typeId), -1)
	path = strings.Replace(path, "{configType}", fmt.Sprintf("%v", configType), -1)
	defaultReturn := new(Jsonschemadocument)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'typeId' is set
	if &typeId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'typeId' when calling IntegrationsApi->GetIntegrationsTypeConfigschema")
	}
	// verify the required parameter 'configType' is set
	if &configType == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'configType' when calling IntegrationsApi->GetIntegrationsTypeConfigschema")
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
	var successPayload *Jsonschemadocument
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

// GetIntegrationsTypes invokes GET /api/v2/integrations/types
//
// List integration types
//
// 
func (a IntegrationsApi) GetIntegrationsTypes(pageSize int32, pageNumber int32, sortBy string, expand []string, nextPage string, previousPage string) (*Integrationtypeentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/types"
	defaultReturn := new(Integrationtypeentitylisting)
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
	var successPayload *Integrationtypeentitylisting
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

// GetIntegrationsUserapps invokes GET /api/v2/integrations/userapps
//
// List permitted user app integrations for the logged in user
//
// 
func (a IntegrationsApi) GetIntegrationsUserapps(pageSize int32, pageNumber int32, sortBy string, expand []string, nextPage string, previousPage string, appHost string) (*Userappentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/userapps"
	defaultReturn := new(Userappentitylisting)
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
	if str, ok := interface{}(appHost).(string); ok {
		if str != "" {
			queryParams["appHost"] = a.Configuration.APIClient.ParameterToString(appHost, collectionFormat)
		}
	} else {
		queryParams["appHost"] = a.Configuration.APIClient.ParameterToString(appHost, collectionFormat)
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
	var successPayload *Userappentitylisting
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

// PatchIntegration invokes PATCH /api/v2/integrations/{integrationId}
//
// Update an integration.
//
// 
func (a IntegrationsApi) PatchIntegration(integrationId string, body Integration, pageSize int32, pageNumber int32, sortBy string, expand []string, nextPage string, previousPage string) (*Integration, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/{integrationId}"
	path = strings.Replace(path, "{integrationId}", fmt.Sprintf("%v", integrationId), -1)
	defaultReturn := new(Integration)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'integrationId' is set
	if &integrationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'integrationId' when calling IntegrationsApi->PatchIntegration")
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
	// body params
	postBody = &body

	var successPayload *Integration
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

// PatchIntegrationsAction invokes PATCH /api/v2/integrations/actions/{actionId}
//
// Patch an Action
//
// 
func (a IntegrationsApi) PatchIntegrationsAction(actionId string, body Updateactioninput) (*Action, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	defaultReturn := new(Action)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->PatchIntegrationsAction")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling IntegrationsApi->PatchIntegrationsAction")
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

	var successPayload *Action
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

// PatchIntegrationsActionDraft invokes PATCH /api/v2/integrations/actions/{actionId}/draft
//
// Update an existing Draft
//
// 
func (a IntegrationsApi) PatchIntegrationsActionDraft(actionId string, body Updatedraftinput) (*Action, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}/draft"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	defaultReturn := new(Action)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->PatchIntegrationsActionDraft")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling IntegrationsApi->PatchIntegrationsActionDraft")
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

	var successPayload *Action
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

// PostIntegrations invokes POST /api/v2/integrations
//
// Create an integration.
//
// 
func (a IntegrationsApi) PostIntegrations(body Createintegrationrequest) (*Integration, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations"
	defaultReturn := new(Integration)
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
	// body params
	postBody = &body

	var successPayload *Integration
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

// PostIntegrationsActionDraft invokes POST /api/v2/integrations/actions/{actionId}/draft
//
// Create a new Draft from existing Action
//
// 
func (a IntegrationsApi) PostIntegrationsActionDraft(actionId string) (*Action, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}/draft"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	defaultReturn := new(Action)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->PostIntegrationsActionDraft")
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
	var successPayload *Action
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

// PostIntegrationsActionDraftPublish invokes POST /api/v2/integrations/actions/{actionId}/draft/publish
//
// Publish a Draft and make it the active Action configuration
//
// 
func (a IntegrationsApi) PostIntegrationsActionDraftPublish(actionId string, body Publishdraftinput) (*Action, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}/draft/publish"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	defaultReturn := new(Action)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->PostIntegrationsActionDraftPublish")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling IntegrationsApi->PostIntegrationsActionDraftPublish")
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

	var successPayload *Action
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

// PostIntegrationsActionDraftTest invokes POST /api/v2/integrations/actions/{actionId}/draft/test
//
// Test the execution of a draft. Responses will show execution steps broken out with intermediate results to help in debugging.
//
// 
func (a IntegrationsApi) PostIntegrationsActionDraftTest(actionId string, body map[string]interface{}) (*Testexecutionresult, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}/draft/test"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	defaultReturn := new(Testexecutionresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->PostIntegrationsActionDraftTest")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling IntegrationsApi->PostIntegrationsActionDraftTest")
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

	var successPayload *Testexecutionresult
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

// PostIntegrationsActionExecute invokes POST /api/v2/integrations/actions/{actionId}/execute
//
// Execute Action and return response from 3rd party.  Responses will follow the schemas defined on the Action for success and error.
//
// 
func (a IntegrationsApi) PostIntegrationsActionExecute(actionId string, body map[string]interface{}) (*map[string]interface{}, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}/execute"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	defaultReturn := new(map[string]interface{})
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->PostIntegrationsActionExecute")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling IntegrationsApi->PostIntegrationsActionExecute")
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

// PostIntegrationsActionTest invokes POST /api/v2/integrations/actions/{actionId}/test
//
// Test the execution of an action. Responses will show execution steps broken out with intermediate results to help in debugging.
//
// 
func (a IntegrationsApi) PostIntegrationsActionTest(actionId string, body map[string]interface{}) (*Testexecutionresult, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/{actionId}/test"
	path = strings.Replace(path, "{actionId}", fmt.Sprintf("%v", actionId), -1)
	defaultReturn := new(Testexecutionresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'actionId' is set
	if &actionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'actionId' when calling IntegrationsApi->PostIntegrationsActionTest")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling IntegrationsApi->PostIntegrationsActionTest")
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

	var successPayload *Testexecutionresult
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

// PostIntegrationsActions invokes POST /api/v2/integrations/actions
//
// Create a new Action
//
// 
func (a IntegrationsApi) PostIntegrationsActions(body Postactioninput) (*Action, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions"
	defaultReturn := new(Action)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling IntegrationsApi->PostIntegrationsActions")
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

	var successPayload *Action
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

// PostIntegrationsActionsDrafts invokes POST /api/v2/integrations/actions/drafts
//
// Create a new Draft
//
// 
func (a IntegrationsApi) PostIntegrationsActionsDrafts(body Postactioninput) (*Action, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/actions/drafts"
	defaultReturn := new(Action)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling IntegrationsApi->PostIntegrationsActionsDrafts")
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

	var successPayload *Action
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

// PostIntegrationsCredentials invokes POST /api/v2/integrations/credentials
//
// Create a set of credentials
//
// 
func (a IntegrationsApi) PostIntegrationsCredentials(body Credential) (*Credentialinfo, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/credentials"
	defaultReturn := new(Credentialinfo)
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
	// body params
	postBody = &body

	var successPayload *Credentialinfo
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

// PostIntegrationsWorkforcemanagementVendorconnection invokes POST /api/v2/integrations/workforcemanagement/vendorconnection
//
// Add a vendor connection
//
// 
func (a IntegrationsApi) PostIntegrationsWorkforcemanagementVendorconnection(body Vendorconnectionrequest) (*Useractioncategoryentitylisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/workforcemanagement/vendorconnection"
	defaultReturn := new(Useractioncategoryentitylisting)
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
	// body params
	postBody = &body

	var successPayload *Useractioncategoryentitylisting
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

// PutIntegrationConfigCurrent invokes PUT /api/v2/integrations/{integrationId}/config/current
//
// Update integration configuration.
//
// 
func (a IntegrationsApi) PutIntegrationConfigCurrent(integrationId string, body Integrationconfiguration) (*Integrationconfiguration, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/{integrationId}/config/current"
	path = strings.Replace(path, "{integrationId}", fmt.Sprintf("%v", integrationId), -1)
	defaultReturn := new(Integrationconfiguration)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'integrationId' is set
	if &integrationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'integrationId' when calling IntegrationsApi->PutIntegrationConfigCurrent")
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

	var successPayload *Integrationconfiguration
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

// PutIntegrationsCredential invokes PUT /api/v2/integrations/credentials/{credentialId}
//
// Update a set of credentials
//
// 
func (a IntegrationsApi) PutIntegrationsCredential(credentialId string, body Credential) (*Credentialinfo, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/credentials/{credentialId}"
	path = strings.Replace(path, "{credentialId}", fmt.Sprintf("%v", credentialId), -1)
	defaultReturn := new(Credentialinfo)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'credentialId' is set
	if &credentialId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'credentialId' when calling IntegrationsApi->PutIntegrationsCredential")
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

	var successPayload *Credentialinfo
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

// PutIntegrationsSpeechTtsSettings invokes PUT /api/v2/integrations/speech/tts/settings
//
// Update TTS settings for an org
//
// 
func (a IntegrationsApi) PutIntegrationsSpeechTtsSettings(body Ttssettings) (*Ttssettings, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/integrations/speech/tts/settings"
	defaultReturn := new(Ttssettings)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling IntegrationsApi->PutIntegrationsSpeechTtsSettings")
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

	var successPayload *Ttssettings
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

