package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
	)

// WidgetsApi provides functions for API endpoints
type WidgetsApi struct {
	Configuration *Configuration
}

// NewWidgetsApi creates an API instance using the default configuration
func NewWidgetsApi() *WidgetsApi {
	fmt.Sprintf(strings.Title(""), "")
	config := GetDefaultConfiguration()
	return &WidgetsApi{
		Configuration: config,
	}
}

// NewWidgetsApiWithConfig creates an API instance using the provided configuration
func NewWidgetsApiWithConfig(config *Configuration) *WidgetsApi {
	return &WidgetsApi{
		Configuration: config,
	}
}

// DeleteWidgetsDeployment invokes DELETE /api/v2/widgets/deployments/{deploymentId}
//
// Delete a Widget deployment
//
// This endpoint is deprecated. Please see the article https://help.mypurecloud.com/articles/deprecation-removal-of-acd-web-chat-version-2/. 
//
// Deprecated: DeleteWidgetsDeployment is deprecated
func (a WidgetsApi) DeleteWidgetsDeployment(deploymentId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/widgets/deployments/{deploymentId}"
	path = strings.Replace(path, "{deploymentId}", url.PathEscape(fmt.Sprintf("%v", deploymentId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'deploymentId' is set
	if &deploymentId == nil {
		// false
		return nil, errors.New("Missing required parameter 'deploymentId' when calling WidgetsApi->DeleteWidgetsDeployment")
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

// GetWidgetsDeployment invokes GET /api/v2/widgets/deployments/{deploymentId}
//
// Get a Widget deployment
//
// This endpoint is deprecated. Please see the article https://help.mypurecloud.com/articles/deprecation-removal-of-acd-web-chat-version-2/. 
//
// Deprecated: GetWidgetsDeployment is deprecated
func (a WidgetsApi) GetWidgetsDeployment(deploymentId string) (*Widgetdeployment, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/widgets/deployments/{deploymentId}"
	path = strings.Replace(path, "{deploymentId}", url.PathEscape(fmt.Sprintf("%v", deploymentId)), -1)
	defaultReturn := new(Widgetdeployment)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'deploymentId' is set
	if &deploymentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'deploymentId' when calling WidgetsApi->GetWidgetsDeployment")
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
	var successPayload *Widgetdeployment
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Widgetdeployment" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWidgetsDeployments invokes GET /api/v2/widgets/deployments
//
// List Widget deployments
//
// This endpoint is deprecated. Please see the article https://help.mypurecloud.com/articles/deprecation-removal-of-acd-web-chat-version-2/. 
//
// Deprecated: GetWidgetsDeployments is deprecated
func (a WidgetsApi) GetWidgetsDeployments() (*Widgetdeploymententitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/widgets/deployments"
	defaultReturn := new(Widgetdeploymententitylisting)
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
	var successPayload *Widgetdeploymententitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Widgetdeploymententitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWidgetsDeployments invokes POST /api/v2/widgets/deployments
//
// Create Widget deployment
//
// This endpoint is deprecated. Please see the article https://help.mypurecloud.com/articles/deprecation-removal-of-acd-web-chat-version-2/. 
//
// Deprecated: PostWidgetsDeployments is deprecated
func (a WidgetsApi) PostWidgetsDeployments(body Widgetdeployment) (*Widgetdeployment, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/widgets/deployments"
	defaultReturn := new(Widgetdeployment)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WidgetsApi->PostWidgetsDeployments")
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

	var successPayload *Widgetdeployment
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Widgetdeployment" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutWidgetsDeployment invokes PUT /api/v2/widgets/deployments/{deploymentId}
//
// Update a Widget deployment
//
// This endpoint is deprecated. Please see the article https://help.mypurecloud.com/articles/deprecation-removal-of-acd-web-chat-version-2/. 
//
// Deprecated: PutWidgetsDeployment is deprecated
func (a WidgetsApi) PutWidgetsDeployment(deploymentId string, body Widgetdeployment) (*Widgetdeployment, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/widgets/deployments/{deploymentId}"
	path = strings.Replace(path, "{deploymentId}", url.PathEscape(fmt.Sprintf("%v", deploymentId)), -1)
	defaultReturn := new(Widgetdeployment)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'deploymentId' is set
	if &deploymentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'deploymentId' when calling WidgetsApi->PutWidgetsDeployment")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WidgetsApi->PutWidgetsDeployment")
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

	var successPayload *Widgetdeployment
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Widgetdeployment" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

