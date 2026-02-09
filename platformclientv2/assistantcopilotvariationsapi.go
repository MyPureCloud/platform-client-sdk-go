package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
	)

// AssistantCopilotVariationsApi provides functions for API endpoints
type AssistantCopilotVariationsApi struct {
	Configuration *Configuration
}

// NewAssistantCopilotVariationsApi creates an API instance using the default configuration
func NewAssistantCopilotVariationsApi() *AssistantCopilotVariationsApi {
	fmt.Sprintf(strings.Title(""), "")
	config := GetDefaultConfiguration()
	return &AssistantCopilotVariationsApi{
		Configuration: config,
	}
}

// NewAssistantCopilotVariationsApiWithConfig creates an API instance using the provided configuration
func NewAssistantCopilotVariationsApiWithConfig(config *Configuration) *AssistantCopilotVariationsApi {
	return &AssistantCopilotVariationsApi{
		Configuration: config,
	}
}

// DeleteAssistantVariation invokes DELETE /api/v2/assistants/{assistantId}/variations/{variationId}
//
// Delete assistant copilot variation by id
func (a AssistantCopilotVariationsApi) DeleteAssistantVariation(assistantId string, variationId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/assistants/{assistantId}/variations/{variationId}"
	path = strings.Replace(path, "{assistantId}", url.PathEscape(fmt.Sprintf("%v", assistantId)), -1)
	path = strings.Replace(path, "{variationId}", url.PathEscape(fmt.Sprintf("%v", variationId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'assistantId' is set
	if &assistantId == nil {
		// false
		return nil, errors.New("Missing required parameter 'assistantId' when calling AssistantCopilotVariationsApi->DeleteAssistantVariation")
	}
	// verify the required parameter 'variationId' is set
	if &variationId == nil {
		// false
		return nil, errors.New("Missing required parameter 'variationId' when calling AssistantCopilotVariationsApi->DeleteAssistantVariation")
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

	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	}
	return response, err
}

// GetAssistantVariation invokes GET /api/v2/assistants/{assistantId}/variations/{variationId}
//
// Get assistant copilot variation by id
func (a AssistantCopilotVariationsApi) GetAssistantVariation(assistantId string, variationId string) (*Assistantcopilotvariation, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/assistants/{assistantId}/variations/{variationId}"
	path = strings.Replace(path, "{assistantId}", url.PathEscape(fmt.Sprintf("%v", assistantId)), -1)
	path = strings.Replace(path, "{variationId}", url.PathEscape(fmt.Sprintf("%v", variationId)), -1)
	defaultReturn := new(Assistantcopilotvariation)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'assistantId' is set
	if &assistantId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'assistantId' when calling AssistantCopilotVariationsApi->GetAssistantVariation")
	}
	// verify the required parameter 'variationId' is set
	if &variationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'variationId' when calling AssistantCopilotVariationsApi->GetAssistantVariation")
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
	var successPayload *Assistantcopilotvariation
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Assistantcopilotvariation" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetAssistantVariations invokes GET /api/v2/assistants/{assistantId}/variations
//
// Get variations of an assistant copilot
func (a AssistantCopilotVariationsApi) GetAssistantVariations(assistantId string) (*Assistantcopilotvariationlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/assistants/{assistantId}/variations"
	path = strings.Replace(path, "{assistantId}", url.PathEscape(fmt.Sprintf("%v", assistantId)), -1)
	defaultReturn := new(Assistantcopilotvariationlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'assistantId' is set
	if &assistantId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'assistantId' when calling AssistantCopilotVariationsApi->GetAssistantVariations")
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
	var successPayload *Assistantcopilotvariationlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Assistantcopilotvariationlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostAssistantVariations invokes POST /api/v2/assistants/{assistantId}/variations
//
// Create assistant copilot variation
func (a AssistantCopilotVariationsApi) PostAssistantVariations(assistantId string, body Assistantcopilotvariation) (*Assistantcopilotvariation, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/assistants/{assistantId}/variations"
	path = strings.Replace(path, "{assistantId}", url.PathEscape(fmt.Sprintf("%v", assistantId)), -1)
	defaultReturn := new(Assistantcopilotvariation)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'assistantId' is set
	if &assistantId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'assistantId' when calling AssistantCopilotVariationsApi->PostAssistantVariations")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling AssistantCopilotVariationsApi->PostAssistantVariations")
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
	// body params
	postBody = &body

	var successPayload *Assistantcopilotvariation
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Assistantcopilotvariation" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutAssistantVariation invokes PUT /api/v2/assistants/{assistantId}/variations/{variationId}
//
// Update assistant copilot variation by id
func (a AssistantCopilotVariationsApi) PutAssistantVariation(assistantId string, variationId string, body Assistantcopilotvariation) (*Assistantcopilotvariation, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/assistants/{assistantId}/variations/{variationId}"
	path = strings.Replace(path, "{assistantId}", url.PathEscape(fmt.Sprintf("%v", assistantId)), -1)
	path = strings.Replace(path, "{variationId}", url.PathEscape(fmt.Sprintf("%v", variationId)), -1)
	defaultReturn := new(Assistantcopilotvariation)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'assistantId' is set
	if &assistantId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'assistantId' when calling AssistantCopilotVariationsApi->PutAssistantVariation")
	}
	// verify the required parameter 'variationId' is set
	if &variationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'variationId' when calling AssistantCopilotVariationsApi->PutAssistantVariation")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling AssistantCopilotVariationsApi->PutAssistantVariation")
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
	// body params
	postBody = &body

	var successPayload *Assistantcopilotvariation
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Assistantcopilotvariation" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

