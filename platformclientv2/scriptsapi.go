package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

// ScriptsApi provides functions for API endpoints
type ScriptsApi struct {
	Configuration *Configuration
}

// NewScriptsApi creates an API instance using the default configuration
func NewScriptsApi() *ScriptsApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating ScriptsApi with base path: %s", strings.ToLower(config.BasePath)))
	return &ScriptsApi{
		Configuration: config,
	}
}

// NewScriptsApiWithConfig creates an API instance using the provided configuration
func NewScriptsApiWithConfig(config *Configuration) *ScriptsApi {
	config.Debugf("Creating ScriptsApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &ScriptsApi{
		Configuration: config,
	}
}

// GetScript invokes GET /api/v2/scripts/{scriptId}
//
// Get a script
//
// 
func (a ScriptsApi) GetScript(scriptId string) (*Script, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/scripts/{scriptId}"
	path = strings.Replace(path, "{scriptId}", fmt.Sprintf("%v", scriptId), -1)
	defaultReturn := new(Script)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'scriptId' is set
	if &scriptId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scriptId' when calling ScriptsApi->GetScript")
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
	var successPayload *Script
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

// GetScriptPage invokes GET /api/v2/scripts/{scriptId}/pages/{pageId}
//
// Get a page
//
// 
func (a ScriptsApi) GetScriptPage(scriptId string, pageId string, scriptDataVersion string) (*Page, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/scripts/{scriptId}/pages/{pageId}"
	path = strings.Replace(path, "{scriptId}", fmt.Sprintf("%v", scriptId), -1)
	path = strings.Replace(path, "{pageId}", fmt.Sprintf("%v", pageId), -1)
	defaultReturn := new(Page)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'scriptId' is set
	if &scriptId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scriptId' when calling ScriptsApi->GetScriptPage")
	}
	// verify the required parameter 'pageId' is set
	if &pageId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'pageId' when calling ScriptsApi->GetScriptPage")
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
	if str, ok := interface{}(scriptDataVersion).(string); ok {
		if str != "" {
			queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
		}
	} else {
		queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
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
	var successPayload *Page
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

// GetScriptPages invokes GET /api/v2/scripts/{scriptId}/pages
//
// Get the list of pages
//
// 
func (a ScriptsApi) GetScriptPages(scriptId string, scriptDataVersion string) ([]Page, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/scripts/{scriptId}/pages"
	path = strings.Replace(path, "{scriptId}", fmt.Sprintf("%v", scriptId), -1)
	defaultReturn := make([]Page, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'scriptId' is set
	if &scriptId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scriptId' when calling ScriptsApi->GetScriptPages")
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
	if str, ok := interface{}(scriptDataVersion).(string); ok {
		if str != "" {
			queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
		}
	} else {
		queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
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
	var successPayload []Page
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

// GetScripts invokes GET /api/v2/scripts
//
// Get the list of scripts
//
// 
func (a ScriptsApi) GetScripts(pageSize int32, pageNumber int32, expand string, name string, feature string, flowId string, sortBy string, sortOrder string, scriptDataVersion string) (*Scriptentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/scripts"
	defaultReturn := new(Scriptentitylisting)
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
	if str, ok := interface{}(expand).(string); ok {
		if str != "" {
			queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
		}
	} else {
		queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
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
	if str, ok := interface{}(feature).(string); ok {
		if str != "" {
			queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, collectionFormat)
		}
	} else {
		queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(flowId).(string); ok {
		if str != "" {
			queryParams["flowId"] = a.Configuration.APIClient.ParameterToString(flowId, collectionFormat)
		}
	} else {
		queryParams["flowId"] = a.Configuration.APIClient.ParameterToString(flowId, collectionFormat)
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
	if str, ok := interface{}(scriptDataVersion).(string); ok {
		if str != "" {
			queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
		}
	} else {
		queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
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
	var successPayload *Scriptentitylisting
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

// GetScriptsPublished invokes GET /api/v2/scripts/published
//
// Get the published scripts.
//
// 
func (a ScriptsApi) GetScriptsPublished(pageSize int32, pageNumber int32, expand string, name string, feature string, flowId string, scriptDataVersion string) (*Scriptentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/scripts/published"
	defaultReturn := new(Scriptentitylisting)
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
	if str, ok := interface{}(expand).(string); ok {
		if str != "" {
			queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
		}
	} else {
		queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
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
	if str, ok := interface{}(feature).(string); ok {
		if str != "" {
			queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, collectionFormat)
		}
	} else {
		queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(flowId).(string); ok {
		if str != "" {
			queryParams["flowId"] = a.Configuration.APIClient.ParameterToString(flowId, collectionFormat)
		}
	} else {
		queryParams["flowId"] = a.Configuration.APIClient.ParameterToString(flowId, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(scriptDataVersion).(string); ok {
		if str != "" {
			queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
		}
	} else {
		queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
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
	var successPayload *Scriptentitylisting
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

// GetScriptsPublishedScriptId invokes GET /api/v2/scripts/published/{scriptId}
//
// Get the published script.
//
// 
func (a ScriptsApi) GetScriptsPublishedScriptId(scriptId string, scriptDataVersion string) (*Script, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/scripts/published/{scriptId}"
	path = strings.Replace(path, "{scriptId}", fmt.Sprintf("%v", scriptId), -1)
	defaultReturn := new(Script)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'scriptId' is set
	if &scriptId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scriptId' when calling ScriptsApi->GetScriptsPublishedScriptId")
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
	if str, ok := interface{}(scriptDataVersion).(string); ok {
		if str != "" {
			queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
		}
	} else {
		queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
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
	var successPayload *Script
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

// GetScriptsPublishedScriptIdPage invokes GET /api/v2/scripts/published/{scriptId}/pages/{pageId}
//
// Get the published page.
//
// 
func (a ScriptsApi) GetScriptsPublishedScriptIdPage(scriptId string, pageId string, scriptDataVersion string) (*Page, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/scripts/published/{scriptId}/pages/{pageId}"
	path = strings.Replace(path, "{scriptId}", fmt.Sprintf("%v", scriptId), -1)
	path = strings.Replace(path, "{pageId}", fmt.Sprintf("%v", pageId), -1)
	defaultReturn := new(Page)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'scriptId' is set
	if &scriptId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scriptId' when calling ScriptsApi->GetScriptsPublishedScriptIdPage")
	}
	// verify the required parameter 'pageId' is set
	if &pageId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'pageId' when calling ScriptsApi->GetScriptsPublishedScriptIdPage")
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
	if str, ok := interface{}(scriptDataVersion).(string); ok {
		if str != "" {
			queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
		}
	} else {
		queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
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
	var successPayload *Page
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

// GetScriptsPublishedScriptIdPages invokes GET /api/v2/scripts/published/{scriptId}/pages
//
// Get the list of published pages
//
// 
func (a ScriptsApi) GetScriptsPublishedScriptIdPages(scriptId string, scriptDataVersion string) ([]Page, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/scripts/published/{scriptId}/pages"
	path = strings.Replace(path, "{scriptId}", fmt.Sprintf("%v", scriptId), -1)
	defaultReturn := make([]Page, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'scriptId' is set
	if &scriptId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scriptId' when calling ScriptsApi->GetScriptsPublishedScriptIdPages")
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
	if str, ok := interface{}(scriptDataVersion).(string); ok {
		if str != "" {
			queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
		}
	} else {
		queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
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
	var successPayload []Page
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

// GetScriptsPublishedScriptIdVariables invokes GET /api/v2/scripts/published/{scriptId}/variables
//
// Get the published variables
//
// 
func (a ScriptsApi) GetScriptsPublishedScriptIdVariables(scriptId string, input string, output string, varType string, scriptDataVersion string) (*map[string]interface{}, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/scripts/published/{scriptId}/variables"
	path = strings.Replace(path, "{scriptId}", fmt.Sprintf("%v", scriptId), -1)
	defaultReturn := new(map[string]interface{})
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'scriptId' is set
	if &scriptId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scriptId' when calling ScriptsApi->GetScriptsPublishedScriptIdVariables")
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
	if str, ok := interface{}(input).(string); ok {
		if str != "" {
			queryParams["input"] = a.Configuration.APIClient.ParameterToString(input, collectionFormat)
		}
	} else {
		queryParams["input"] = a.Configuration.APIClient.ParameterToString(input, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(output).(string); ok {
		if str != "" {
			queryParams["output"] = a.Configuration.APIClient.ParameterToString(output, collectionFormat)
		}
	} else {
		queryParams["output"] = a.Configuration.APIClient.ParameterToString(output, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(varType).(string); ok {
		if str != "" {
			queryParams["varType"] = a.Configuration.APIClient.ParameterToString(varType, collectionFormat)
		}
	} else {
		queryParams["varType"] = a.Configuration.APIClient.ParameterToString(varType, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(scriptDataVersion).(string); ok {
		if str != "" {
			queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
		}
	} else {
		queryParams["scriptDataVersion"] = a.Configuration.APIClient.ParameterToString(scriptDataVersion, collectionFormat)
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

// GetScriptsUploadStatus invokes GET /api/v2/scripts/uploads/{uploadId}/status
//
// Get the upload status of an imported script
//
// 
func (a ScriptsApi) GetScriptsUploadStatus(uploadId string, longPoll bool) (*Importscriptstatusresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/scripts/uploads/{uploadId}/status"
	path = strings.Replace(path, "{uploadId}", fmt.Sprintf("%v", uploadId), -1)
	defaultReturn := new(Importscriptstatusresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'uploadId' is set
	if &uploadId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'uploadId' when calling ScriptsApi->GetScriptsUploadStatus")
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
	if str, ok := interface{}(longPoll).(string); ok {
		if str != "" {
			queryParams["longPoll"] = a.Configuration.APIClient.ParameterToString(longPoll, collectionFormat)
		}
	} else {
		queryParams["longPoll"] = a.Configuration.APIClient.ParameterToString(longPoll, collectionFormat)
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
	var successPayload *Importscriptstatusresponse
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

// PostScriptExport invokes POST /api/v2/scripts/{scriptId}/export
//
// Export a script via download service.
//
// 
func (a ScriptsApi) PostScriptExport(scriptId string, body Exportscriptrequest) (*Exportscriptresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/scripts/{scriptId}/export"
	path = strings.Replace(path, "{scriptId}", fmt.Sprintf("%v", scriptId), -1)
	defaultReturn := new(Exportscriptresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'scriptId' is set
	if &scriptId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scriptId' when calling ScriptsApi->PostScriptExport")
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

	var successPayload *Exportscriptresponse
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

