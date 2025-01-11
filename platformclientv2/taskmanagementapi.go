package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
	)

// TaskManagementApi provides functions for API endpoints
type TaskManagementApi struct {
	Configuration *Configuration
}

// NewTaskManagementApi creates an API instance using the default configuration
func NewTaskManagementApi() *TaskManagementApi {
	fmt.Sprintf(strings.Title(""), "")
	config := GetDefaultConfiguration()
	return &TaskManagementApi{
		Configuration: config,
	}
}

// NewTaskManagementApiWithConfig creates an API instance using the provided configuration
func NewTaskManagementApiWithConfig(config *Configuration) *TaskManagementApi {
	return &TaskManagementApi{
		Configuration: config,
	}
}

// DeleteTaskmanagementWorkbin invokes DELETE /api/v2/taskmanagement/workbins/{workbinId}
//
// Delete a workbin
func (a TaskManagementApi) DeleteTaskmanagementWorkbin(workbinId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workbins/{workbinId}"
	path = strings.Replace(path, "{workbinId}", url.PathEscape(fmt.Sprintf("%v", workbinId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workbinId' is set
	if &workbinId == nil {
		// false
		return nil, errors.New("Missing required parameter 'workbinId' when calling TaskManagementApi->DeleteTaskmanagementWorkbin")
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

// DeleteTaskmanagementWorkitem invokes DELETE /api/v2/taskmanagement/workitems/{workitemId}
//
// Delete a workitem
func (a TaskManagementApi) DeleteTaskmanagementWorkitem(workitemId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->DeleteTaskmanagementWorkitem")
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

// DeleteTaskmanagementWorkitemsBulkAddJob invokes DELETE /api/v2/taskmanagement/workitems/bulk/add/jobs/{bulkJobId}
//
// Delete a bulk add job
func (a TaskManagementApi) DeleteTaskmanagementWorkitemsBulkAddJob(bulkJobId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/bulk/add/jobs/{bulkJobId}"
	path = strings.Replace(path, "{bulkJobId}", url.PathEscape(fmt.Sprintf("%v", bulkJobId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'bulkJobId' is set
	if &bulkJobId == nil {
		// false
		return nil, errors.New("Missing required parameter 'bulkJobId' when calling TaskManagementApi->DeleteTaskmanagementWorkitemsBulkAddJob")
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

// DeleteTaskmanagementWorkitemsBulkTerminateJob invokes DELETE /api/v2/taskmanagement/workitems/bulk/terminate/jobs/{bulkJobId}
//
// Delete a Bulk job
func (a TaskManagementApi) DeleteTaskmanagementWorkitemsBulkTerminateJob(bulkJobId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/bulk/terminate/jobs/{bulkJobId}"
	path = strings.Replace(path, "{bulkJobId}", url.PathEscape(fmt.Sprintf("%v", bulkJobId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'bulkJobId' is set
	if &bulkJobId == nil {
		// false
		return nil, errors.New("Missing required parameter 'bulkJobId' when calling TaskManagementApi->DeleteTaskmanagementWorkitemsBulkTerminateJob")
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

// DeleteTaskmanagementWorkitemsSchema invokes DELETE /api/v2/taskmanagement/workitems/schemas/{schemaId}
//
// Delete a schema
func (a TaskManagementApi) DeleteTaskmanagementWorkitemsSchema(schemaId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/schemas/{schemaId}"
	path = strings.Replace(path, "{schemaId}", url.PathEscape(fmt.Sprintf("%v", schemaId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'schemaId' is set
	if &schemaId == nil {
		// false
		return nil, errors.New("Missing required parameter 'schemaId' when calling TaskManagementApi->DeleteTaskmanagementWorkitemsSchema")
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

// DeleteTaskmanagementWorktype invokes DELETE /api/v2/taskmanagement/worktypes/{worktypeId}
//
// Delete a worktype
func (a TaskManagementApi) DeleteTaskmanagementWorktype(worktypeId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->DeleteTaskmanagementWorktype")
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

// DeleteTaskmanagementWorktypeFlowsDatebasedRule invokes DELETE /api/v2/taskmanagement/worktypes/{worktypeId}/flows/datebased/rules/{ruleId}
//
// Delete a date based rule
func (a TaskManagementApi) DeleteTaskmanagementWorktypeFlowsDatebasedRule(worktypeId string, ruleId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/datebased/rules/{ruleId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	path = strings.Replace(path, "{ruleId}", url.PathEscape(fmt.Sprintf("%v", ruleId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->DeleteTaskmanagementWorktypeFlowsDatebasedRule")
	}
	// verify the required parameter 'ruleId' is set
	if &ruleId == nil {
		// false
		return nil, errors.New("Missing required parameter 'ruleId' when calling TaskManagementApi->DeleteTaskmanagementWorktypeFlowsDatebasedRule")
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

// DeleteTaskmanagementWorktypeFlowsOnattributechangeRule invokes DELETE /api/v2/taskmanagement/worktypes/{worktypeId}/flows/onattributechange/rules/{ruleId}
//
// Delete a rule
func (a TaskManagementApi) DeleteTaskmanagementWorktypeFlowsOnattributechangeRule(worktypeId string, ruleId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/onattributechange/rules/{ruleId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	path = strings.Replace(path, "{ruleId}", url.PathEscape(fmt.Sprintf("%v", ruleId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->DeleteTaskmanagementWorktypeFlowsOnattributechangeRule")
	}
	// verify the required parameter 'ruleId' is set
	if &ruleId == nil {
		// false
		return nil, errors.New("Missing required parameter 'ruleId' when calling TaskManagementApi->DeleteTaskmanagementWorktypeFlowsOnattributechangeRule")
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

// DeleteTaskmanagementWorktypeFlowsOncreateRule invokes DELETE /api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}
//
// Delete a rule
func (a TaskManagementApi) DeleteTaskmanagementWorktypeFlowsOncreateRule(worktypeId string, ruleId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	path = strings.Replace(path, "{ruleId}", url.PathEscape(fmt.Sprintf("%v", ruleId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->DeleteTaskmanagementWorktypeFlowsOncreateRule")
	}
	// verify the required parameter 'ruleId' is set
	if &ruleId == nil {
		// false
		return nil, errors.New("Missing required parameter 'ruleId' when calling TaskManagementApi->DeleteTaskmanagementWorktypeFlowsOncreateRule")
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

// DeleteTaskmanagementWorktypeStatus invokes DELETE /api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}
//
// Delete a status
func (a TaskManagementApi) DeleteTaskmanagementWorktypeStatus(worktypeId string, statusId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	path = strings.Replace(path, "{statusId}", url.PathEscape(fmt.Sprintf("%v", statusId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->DeleteTaskmanagementWorktypeStatus")
	}
	// verify the required parameter 'statusId' is set
	if &statusId == nil {
		// false
		return nil, errors.New("Missing required parameter 'statusId' when calling TaskManagementApi->DeleteTaskmanagementWorktypeStatus")
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

// GetTaskmanagementWorkbin invokes GET /api/v2/taskmanagement/workbins/{workbinId}
//
// Get a workbin
func (a TaskManagementApi) GetTaskmanagementWorkbin(workbinId string) (*Workbin, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workbins/{workbinId}"
	path = strings.Replace(path, "{workbinId}", url.PathEscape(fmt.Sprintf("%v", workbinId)), -1)
	defaultReturn := new(Workbin)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workbinId' is set
	if &workbinId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workbinId' when calling TaskManagementApi->GetTaskmanagementWorkbin")
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
	var successPayload *Workbin
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workbin" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkbinHistory invokes GET /api/v2/taskmanagement/workbins/{workbinId}/history
//
// Get a listing of a workbin's attribute change history
//
// Preview: GetTaskmanagementWorkbinHistory is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a TaskManagementApi) GetTaskmanagementWorkbinHistory(workbinId string, after string, pageSize int, sortOrder string) (*Workbinchangelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workbins/{workbinId}/history"
	path = strings.Replace(path, "{workbinId}", url.PathEscape(fmt.Sprintf("%v", workbinId)), -1)
	defaultReturn := new(Workbinchangelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workbinId' is set
	if &workbinId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workbinId' when calling TaskManagementApi->GetTaskmanagementWorkbinHistory")
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
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	

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
	var successPayload *Workbinchangelisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workbinchangelisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkbinVersion invokes GET /api/v2/taskmanagement/workbins/{workbinId}/versions/{entityVersion}
//
// Get a version of a workbin
//
// Preview: GetTaskmanagementWorkbinVersion is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a TaskManagementApi) GetTaskmanagementWorkbinVersion(workbinId string, entityVersion int) (*Workbinversion, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workbins/{workbinId}/versions/{entityVersion}"
	path = strings.Replace(path, "{workbinId}", url.PathEscape(fmt.Sprintf("%v", workbinId)), -1)
	path = strings.Replace(path, "{entityVersion}", url.PathEscape(fmt.Sprintf("%v", entityVersion)), -1)
	defaultReturn := new(Workbinversion)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workbinId' is set
	if &workbinId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workbinId' when calling TaskManagementApi->GetTaskmanagementWorkbinVersion")
	}
	// verify the required parameter 'entityVersion' is set
	if &entityVersion == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'entityVersion' when calling TaskManagementApi->GetTaskmanagementWorkbinVersion")
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
	var successPayload *Workbinversion
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workbinversion" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkbinVersions invokes GET /api/v2/taskmanagement/workbins/{workbinId}/versions
//
// Get all versions of a workbin
//
// Preview: GetTaskmanagementWorkbinVersions is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a TaskManagementApi) GetTaskmanagementWorkbinVersions(workbinId string, after string, pageSize int, sortOrder string) (*Workbinversionlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workbins/{workbinId}/versions"
	path = strings.Replace(path, "{workbinId}", url.PathEscape(fmt.Sprintf("%v", workbinId)), -1)
	defaultReturn := new(Workbinversionlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workbinId' is set
	if &workbinId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workbinId' when calling TaskManagementApi->GetTaskmanagementWorkbinVersions")
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
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	

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
	var successPayload *Workbinversionlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workbinversionlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitem invokes GET /api/v2/taskmanagement/workitems/{workitemId}
//
// Get a workitem
func (a TaskManagementApi) GetTaskmanagementWorkitem(workitemId string, expands string) (*Workitem, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	defaultReturn := new(Workitem)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->GetTaskmanagementWorkitem")
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
	
	queryParams["expands"] = a.Configuration.APIClient.ParameterToString(expands, "")
	

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
	var successPayload *Workitem
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitem" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemHistory invokes GET /api/v2/taskmanagement/workitems/{workitemId}/history
//
// Get a listing of a workitem's attribute change history
//
// Preview: GetTaskmanagementWorkitemHistory is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a TaskManagementApi) GetTaskmanagementWorkitemHistory(workitemId string, after string, pageSize int, sortOrder string) (*Workitemchangelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}/history"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	defaultReturn := new(Workitemchangelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->GetTaskmanagementWorkitemHistory")
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
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	

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
	var successPayload *Workitemchangelisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemchangelisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemUserWrapups invokes GET /api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups
//
// Get all wrapup codes added for the given user for a workitem.
func (a TaskManagementApi) GetTaskmanagementWorkitemUserWrapups(workitemId string, userId string, expands string, after string, pageSize int, sortOrder string) (*Workitemwrapup, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	path = strings.Replace(path, "{userId}", url.PathEscape(fmt.Sprintf("%v", userId)), -1)
	defaultReturn := new(Workitemwrapup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->GetTaskmanagementWorkitemUserWrapups")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling TaskManagementApi->GetTaskmanagementWorkitemUserWrapups")
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
	
	queryParams["expands"] = a.Configuration.APIClient.ParameterToString(expands, "")
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	

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
	var successPayload *Workitemwrapup
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemwrapup" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemVersion invokes GET /api/v2/taskmanagement/workitems/{workitemId}/versions/{entityVersion}
//
// Get a version of a workitem
//
// Preview: GetTaskmanagementWorkitemVersion is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a TaskManagementApi) GetTaskmanagementWorkitemVersion(workitemId string, entityVersion int) (*Workitemversion, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}/versions/{entityVersion}"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	path = strings.Replace(path, "{entityVersion}", url.PathEscape(fmt.Sprintf("%v", entityVersion)), -1)
	defaultReturn := new(Workitemversion)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->GetTaskmanagementWorkitemVersion")
	}
	// verify the required parameter 'entityVersion' is set
	if &entityVersion == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'entityVersion' when calling TaskManagementApi->GetTaskmanagementWorkitemVersion")
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
	var successPayload *Workitemversion
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemversion" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemVersions invokes GET /api/v2/taskmanagement/workitems/{workitemId}/versions
//
// Get all versions of a workitem
//
// Preview: GetTaskmanagementWorkitemVersions is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a TaskManagementApi) GetTaskmanagementWorkitemVersions(workitemId string, after string, pageSize int, sortOrder string) (*Workitemversionlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}/versions"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	defaultReturn := new(Workitemversionlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->GetTaskmanagementWorkitemVersions")
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
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	

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
	var successPayload *Workitemversionlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemversionlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemWrapups invokes GET /api/v2/taskmanagement/workitems/{workitemId}/wrapups
//
// Get all wrapup codes added for all users for a workitem.
func (a TaskManagementApi) GetTaskmanagementWorkitemWrapups(workitemId string, expands string, after string, pageSize int, sortOrder string) (*Workitemwrapupentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}/wrapups"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	defaultReturn := new(Workitemwrapupentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->GetTaskmanagementWorkitemWrapups")
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
	
	queryParams["expands"] = a.Configuration.APIClient.ParameterToString(expands, "")
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	

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
	var successPayload *Workitemwrapupentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemwrapupentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemsBulkAddJob invokes GET /api/v2/taskmanagement/workitems/bulk/add/jobs/{bulkJobId}
//
// Get the bulk add job associated with the job id.
func (a TaskManagementApi) GetTaskmanagementWorkitemsBulkAddJob(bulkJobId string) (*Bulkjob, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/bulk/add/jobs/{bulkJobId}"
	path = strings.Replace(path, "{bulkJobId}", url.PathEscape(fmt.Sprintf("%v", bulkJobId)), -1)
	defaultReturn := new(Bulkjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'bulkJobId' is set
	if &bulkJobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bulkJobId' when calling TaskManagementApi->GetTaskmanagementWorkitemsBulkAddJob")
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
	var successPayload *Bulkjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulkjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemsBulkAddJobResults invokes GET /api/v2/taskmanagement/workitems/bulk/add/jobs/{bulkJobId}/results
//
// Get bulk add job results.
func (a TaskManagementApi) GetTaskmanagementWorkitemsBulkAddJobResults(bulkJobId string) (*Bulkjobaddresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/bulk/add/jobs/{bulkJobId}/results"
	path = strings.Replace(path, "{bulkJobId}", url.PathEscape(fmt.Sprintf("%v", bulkJobId)), -1)
	defaultReturn := new(Bulkjobaddresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'bulkJobId' is set
	if &bulkJobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bulkJobId' when calling TaskManagementApi->GetTaskmanagementWorkitemsBulkAddJobResults")
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
	var successPayload *Bulkjobaddresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulkjobaddresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemsBulkJobsUsersMe invokes GET /api/v2/taskmanagement/workitems/bulk/jobs/users/me
//
// Get bulk jobs created by the currently logged in user.
func (a TaskManagementApi) GetTaskmanagementWorkitemsBulkJobsUsersMe(after string, pageSize int, sortOrder string, action string) (*Bulkjobslisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/bulk/jobs/users/me"
	defaultReturn := new(Bulkjobslisting)
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
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
	queryParams["action"] = a.Configuration.APIClient.ParameterToString(action, "")
	

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
	var successPayload *Bulkjobslisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulkjobslisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemsBulkTerminateJob invokes GET /api/v2/taskmanagement/workitems/bulk/terminate/jobs/{bulkJobId}
//
// Get the bulk job associated with the job id.
func (a TaskManagementApi) GetTaskmanagementWorkitemsBulkTerminateJob(bulkJobId string) (*Bulkjob, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/bulk/terminate/jobs/{bulkJobId}"
	path = strings.Replace(path, "{bulkJobId}", url.PathEscape(fmt.Sprintf("%v", bulkJobId)), -1)
	defaultReturn := new(Bulkjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'bulkJobId' is set
	if &bulkJobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bulkJobId' when calling TaskManagementApi->GetTaskmanagementWorkitemsBulkTerminateJob")
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
	var successPayload *Bulkjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulkjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemsBulkTerminateJobResults invokes GET /api/v2/taskmanagement/workitems/bulk/terminate/jobs/{bulkJobId}/results
//
// Get bulk terminate job results.
func (a TaskManagementApi) GetTaskmanagementWorkitemsBulkTerminateJobResults(bulkJobId string) (*Bulkjobterminateresultsresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/bulk/terminate/jobs/{bulkJobId}/results"
	path = strings.Replace(path, "{bulkJobId}", url.PathEscape(fmt.Sprintf("%v", bulkJobId)), -1)
	defaultReturn := new(Bulkjobterminateresultsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'bulkJobId' is set
	if &bulkJobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bulkJobId' when calling TaskManagementApi->GetTaskmanagementWorkitemsBulkTerminateJobResults")
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
	var successPayload *Bulkjobterminateresultsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulkjobterminateresultsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemsQueryJob invokes GET /api/v2/taskmanagement/workitems/query/jobs/{jobId}
//
// Get the workitem query job associated with the job id.
func (a TaskManagementApi) GetTaskmanagementWorkitemsQueryJob(jobId string) (*Workitemqueryjobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/query/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Workitemqueryjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling TaskManagementApi->GetTaskmanagementWorkitemsQueryJob")
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
	var successPayload *Workitemqueryjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemqueryjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemsQueryJobResults invokes GET /api/v2/taskmanagement/workitems/query/jobs/{jobId}/results
//
// Get results from for workitem query job 
func (a TaskManagementApi) GetTaskmanagementWorkitemsQueryJobResults(jobId string) (*Workitempagedentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/query/jobs/{jobId}/results"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Workitempagedentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling TaskManagementApi->GetTaskmanagementWorkitemsQueryJobResults")
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
	var successPayload *Workitempagedentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitempagedentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemsSchema invokes GET /api/v2/taskmanagement/workitems/schemas/{schemaId}
//
// Get a schema
func (a TaskManagementApi) GetTaskmanagementWorkitemsSchema(schemaId string) (*Dataschema, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/schemas/{schemaId}"
	path = strings.Replace(path, "{schemaId}", url.PathEscape(fmt.Sprintf("%v", schemaId)), -1)
	defaultReturn := new(Dataschema)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'schemaId' is set
	if &schemaId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'schemaId' when calling TaskManagementApi->GetTaskmanagementWorkitemsSchema")
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
	var successPayload *Dataschema
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dataschema" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemsSchemaVersion invokes GET /api/v2/taskmanagement/workitems/schemas/{schemaId}/versions/{versionId}
//
// Get a specific version of a schema
func (a TaskManagementApi) GetTaskmanagementWorkitemsSchemaVersion(schemaId string, versionId string) (*Dataschema, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/schemas/{schemaId}/versions/{versionId}"
	path = strings.Replace(path, "{schemaId}", url.PathEscape(fmt.Sprintf("%v", schemaId)), -1)
	path = strings.Replace(path, "{versionId}", url.PathEscape(fmt.Sprintf("%v", versionId)), -1)
	defaultReturn := new(Dataschema)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'schemaId' is set
	if &schemaId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'schemaId' when calling TaskManagementApi->GetTaskmanagementWorkitemsSchemaVersion")
	}
	// verify the required parameter 'versionId' is set
	if &versionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'versionId' when calling TaskManagementApi->GetTaskmanagementWorkitemsSchemaVersion")
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
	var successPayload *Dataschema
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dataschema" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemsSchemaVersions invokes GET /api/v2/taskmanagement/workitems/schemas/{schemaId}/versions
//
// Get all versions of a schema
func (a TaskManagementApi) GetTaskmanagementWorkitemsSchemaVersions(schemaId string) (*Dataschema, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/schemas/{schemaId}/versions"
	path = strings.Replace(path, "{schemaId}", url.PathEscape(fmt.Sprintf("%v", schemaId)), -1)
	defaultReturn := new(Dataschema)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'schemaId' is set
	if &schemaId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'schemaId' when calling TaskManagementApi->GetTaskmanagementWorkitemsSchemaVersions")
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
	var successPayload *Dataschema
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dataschema" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorkitemsSchemas invokes GET /api/v2/taskmanagement/workitems/schemas
//
// Get a list of schemas.
func (a TaskManagementApi) GetTaskmanagementWorkitemsSchemas() (*Dataschemalisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/schemas"
	defaultReturn := new(Dataschemalisting)
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
	var successPayload *Dataschemalisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dataschemalisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorktype invokes GET /api/v2/taskmanagement/worktypes/{worktypeId}
//
// Get a worktype
func (a TaskManagementApi) GetTaskmanagementWorktype(worktypeId string, expands []string) (*Worktype, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	defaultReturn := new(Worktype)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->GetTaskmanagementWorktype")
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
	
	queryParams["expands"] = a.Configuration.APIClient.ParameterToString(expands, "multi")
	

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
	var successPayload *Worktype
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Worktype" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorktypeFlowsDatebasedRule invokes GET /api/v2/taskmanagement/worktypes/{worktypeId}/flows/datebased/rules/{ruleId}
//
// Get a date based rule
func (a TaskManagementApi) GetTaskmanagementWorktypeFlowsDatebasedRule(worktypeId string, ruleId string) (*Workitemdatebasedrule, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/datebased/rules/{ruleId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	path = strings.Replace(path, "{ruleId}", url.PathEscape(fmt.Sprintf("%v", ruleId)), -1)
	defaultReturn := new(Workitemdatebasedrule)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->GetTaskmanagementWorktypeFlowsDatebasedRule")
	}
	// verify the required parameter 'ruleId' is set
	if &ruleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'ruleId' when calling TaskManagementApi->GetTaskmanagementWorktypeFlowsDatebasedRule")
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
	var successPayload *Workitemdatebasedrule
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemdatebasedrule" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorktypeFlowsDatebasedRules invokes GET /api/v2/taskmanagement/worktypes/{worktypeId}/flows/datebased/rules
//
// Get all date based rules for a worktype
func (a TaskManagementApi) GetTaskmanagementWorktypeFlowsDatebasedRules(worktypeId string, after string, pageSize int) (*Workitemdatebasedrulelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/datebased/rules"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	defaultReturn := new(Workitemdatebasedrulelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->GetTaskmanagementWorktypeFlowsDatebasedRules")
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
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	

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
	var successPayload *Workitemdatebasedrulelisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemdatebasedrulelisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorktypeFlowsOnattributechangeRule invokes GET /api/v2/taskmanagement/worktypes/{worktypeId}/flows/onattributechange/rules/{ruleId}
//
// Get an attribute change rule
func (a TaskManagementApi) GetTaskmanagementWorktypeFlowsOnattributechangeRule(worktypeId string, ruleId string) (*Workitemonattributechangerule, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/onattributechange/rules/{ruleId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	path = strings.Replace(path, "{ruleId}", url.PathEscape(fmt.Sprintf("%v", ruleId)), -1)
	defaultReturn := new(Workitemonattributechangerule)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->GetTaskmanagementWorktypeFlowsOnattributechangeRule")
	}
	// verify the required parameter 'ruleId' is set
	if &ruleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'ruleId' when calling TaskManagementApi->GetTaskmanagementWorktypeFlowsOnattributechangeRule")
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
	var successPayload *Workitemonattributechangerule
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemonattributechangerule" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorktypeFlowsOnattributechangeRules invokes GET /api/v2/taskmanagement/worktypes/{worktypeId}/flows/onattributechange/rules
//
// Get all attribute-change rules for a worktype
func (a TaskManagementApi) GetTaskmanagementWorktypeFlowsOnattributechangeRules(worktypeId string, after string, pageSize int) (*Workitemonattributechangerulelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/onattributechange/rules"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	defaultReturn := new(Workitemonattributechangerulelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->GetTaskmanagementWorktypeFlowsOnattributechangeRules")
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
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	

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
	var successPayload *Workitemonattributechangerulelisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemonattributechangerulelisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorktypeFlowsOncreateRule invokes GET /api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}
//
// Get an on-create rule
func (a TaskManagementApi) GetTaskmanagementWorktypeFlowsOncreateRule(worktypeId string, ruleId string) (*Workitemoncreaterule, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	path = strings.Replace(path, "{ruleId}", url.PathEscape(fmt.Sprintf("%v", ruleId)), -1)
	defaultReturn := new(Workitemoncreaterule)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->GetTaskmanagementWorktypeFlowsOncreateRule")
	}
	// verify the required parameter 'ruleId' is set
	if &ruleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'ruleId' when calling TaskManagementApi->GetTaskmanagementWorktypeFlowsOncreateRule")
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
	var successPayload *Workitemoncreaterule
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemoncreaterule" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorktypeFlowsOncreateRules invokes GET /api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules
//
// Get all on-create rules for a worktype
func (a TaskManagementApi) GetTaskmanagementWorktypeFlowsOncreateRules(worktypeId string, after string, pageSize int) (*Workitemoncreaterulelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	defaultReturn := new(Workitemoncreaterulelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->GetTaskmanagementWorktypeFlowsOncreateRules")
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
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	

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
	var successPayload *Workitemoncreaterulelisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemoncreaterulelisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorktypeHistory invokes GET /api/v2/taskmanagement/worktypes/{worktypeId}/history
//
// Get a listing of a worktype's attribute change history
//
// Preview: GetTaskmanagementWorktypeHistory is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a TaskManagementApi) GetTaskmanagementWorktypeHistory(worktypeId string, after string, pageSize int, sortOrder string, fields []string) (*Worktypechangelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/history"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	defaultReturn := new(Worktypechangelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->GetTaskmanagementWorktypeHistory")
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
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
	queryParams["fields"] = a.Configuration.APIClient.ParameterToString(fields, "multi")
	

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
	var successPayload *Worktypechangelisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Worktypechangelisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorktypeStatus invokes GET /api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}
//
// Get a status
func (a TaskManagementApi) GetTaskmanagementWorktypeStatus(worktypeId string, statusId string) (*Workitemstatus, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	path = strings.Replace(path, "{statusId}", url.PathEscape(fmt.Sprintf("%v", statusId)), -1)
	defaultReturn := new(Workitemstatus)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->GetTaskmanagementWorktypeStatus")
	}
	// verify the required parameter 'statusId' is set
	if &statusId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'statusId' when calling TaskManagementApi->GetTaskmanagementWorktypeStatus")
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
	var successPayload *Workitemstatus
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemstatus" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorktypeStatuses invokes GET /api/v2/taskmanagement/worktypes/{worktypeId}/statuses
//
// Get list of statuses for this worktype.
func (a TaskManagementApi) GetTaskmanagementWorktypeStatuses(worktypeId string) (*Workitemstatuslisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	defaultReturn := new(Workitemstatuslisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->GetTaskmanagementWorktypeStatuses")
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
	var successPayload *Workitemstatuslisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemstatuslisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorktypeVersion invokes GET /api/v2/taskmanagement/worktypes/{worktypeId}/versions/{entityVersion}
//
// Get a version of a worktype
//
// Preview: GetTaskmanagementWorktypeVersion is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a TaskManagementApi) GetTaskmanagementWorktypeVersion(worktypeId string, entityVersion int) (*Worktypeversion, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/versions/{entityVersion}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	path = strings.Replace(path, "{entityVersion}", url.PathEscape(fmt.Sprintf("%v", entityVersion)), -1)
	defaultReturn := new(Worktypeversion)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->GetTaskmanagementWorktypeVersion")
	}
	// verify the required parameter 'entityVersion' is set
	if &entityVersion == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'entityVersion' when calling TaskManagementApi->GetTaskmanagementWorktypeVersion")
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
	var successPayload *Worktypeversion
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Worktypeversion" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTaskmanagementWorktypeVersions invokes GET /api/v2/taskmanagement/worktypes/{worktypeId}/versions
//
// Get all versions of a worktype
//
// Preview: GetTaskmanagementWorktypeVersions is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a TaskManagementApi) GetTaskmanagementWorktypeVersions(worktypeId string, after string, pageSize int, sortOrder string) (*Worktypeversionlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/versions"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	defaultReturn := new(Worktypeversionlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->GetTaskmanagementWorktypeVersions")
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
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	

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
	var successPayload *Worktypeversionlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Worktypeversionlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchTaskmanagementWorkbin invokes PATCH /api/v2/taskmanagement/workbins/{workbinId}
//
// Update the attributes of a workbin
func (a TaskManagementApi) PatchTaskmanagementWorkbin(workbinId string, body Workbinupdate) (*Workbin, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workbins/{workbinId}"
	path = strings.Replace(path, "{workbinId}", url.PathEscape(fmt.Sprintf("%v", workbinId)), -1)
	defaultReturn := new(Workbin)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workbinId' is set
	if &workbinId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workbinId' when calling TaskManagementApi->PatchTaskmanagementWorkbin")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PatchTaskmanagementWorkbin")
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

	var successPayload *Workbin
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workbin" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchTaskmanagementWorkitem invokes PATCH /api/v2/taskmanagement/workitems/{workitemId}
//
// Update the attributes of a workitem
func (a TaskManagementApi) PatchTaskmanagementWorkitem(workitemId string, body Workitemupdate) (*Workitem, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	defaultReturn := new(Workitem)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->PatchTaskmanagementWorkitem")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PatchTaskmanagementWorkitem")
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

	var successPayload *Workitem
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitem" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchTaskmanagementWorkitemAssignment invokes PATCH /api/v2/taskmanagement/workitems/{workitemId}/assignment
//
// Attempts to manually assign a specified workitem to a specified user.  Ignores bullseye ring, PAR score, skills, and languages.
func (a TaskManagementApi) PatchTaskmanagementWorkitemAssignment(workitemId string, body Workitemmanualassign) (*APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}/assignment"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->PatchTaskmanagementWorkitemAssignment")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PatchTaskmanagementWorkitemAssignment")
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


	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	}
	return response, err
}

// PatchTaskmanagementWorkitemUserWrapups invokes PATCH /api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups
//
// Add/Remove a wrapup code for a given user in a workitem.
func (a TaskManagementApi) PatchTaskmanagementWorkitemUserWrapups(workitemId string, userId string, body Workitemwrapupupdate) (*Workitemwrapup, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}/users/{userId}/wrapups"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	path = strings.Replace(path, "{userId}", url.PathEscape(fmt.Sprintf("%v", userId)), -1)
	defaultReturn := new(Workitemwrapup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->PatchTaskmanagementWorkitemUserWrapups")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling TaskManagementApi->PatchTaskmanagementWorkitemUserWrapups")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PatchTaskmanagementWorkitemUserWrapups")
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

	var successPayload *Workitemwrapup
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemwrapup" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchTaskmanagementWorkitemUsersMeWrapups invokes PATCH /api/v2/taskmanagement/workitems/{workitemId}/users/me/wrapups
//
// Add/Remove a wrapup code for the current user in a workitem.
func (a TaskManagementApi) PatchTaskmanagementWorkitemUsersMeWrapups(workitemId string, body Workitemwrapupupdate) (*Workitemwrapup, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}/users/me/wrapups"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	defaultReturn := new(Workitemwrapup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->PatchTaskmanagementWorkitemUsersMeWrapups")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PatchTaskmanagementWorkitemUsersMeWrapups")
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

	var successPayload *Workitemwrapup
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemwrapup" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchTaskmanagementWorkitemsBulkAddJob invokes PATCH /api/v2/taskmanagement/workitems/bulk/add/jobs/{bulkJobId}
//
// Update workitem bulk add job.
func (a TaskManagementApi) PatchTaskmanagementWorkitemsBulkAddJob(bulkJobId string, body Bulkjobupdate) (*Bulkjob, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/bulk/add/jobs/{bulkJobId}"
	path = strings.Replace(path, "{bulkJobId}", url.PathEscape(fmt.Sprintf("%v", bulkJobId)), -1)
	defaultReturn := new(Bulkjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'bulkJobId' is set
	if &bulkJobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bulkJobId' when calling TaskManagementApi->PatchTaskmanagementWorkitemsBulkAddJob")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PatchTaskmanagementWorkitemsBulkAddJob")
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

	var successPayload *Bulkjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulkjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchTaskmanagementWorkitemsBulkTerminateJob invokes PATCH /api/v2/taskmanagement/workitems/bulk/terminate/jobs/{bulkJobId}
//
// Update workitem bulk terminate job.
func (a TaskManagementApi) PatchTaskmanagementWorkitemsBulkTerminateJob(bulkJobId string, body Bulkjobupdate) (*Bulkjob, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/bulk/terminate/jobs/{bulkJobId}"
	path = strings.Replace(path, "{bulkJobId}", url.PathEscape(fmt.Sprintf("%v", bulkJobId)), -1)
	defaultReturn := new(Bulkjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'bulkJobId' is set
	if &bulkJobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bulkJobId' when calling TaskManagementApi->PatchTaskmanagementWorkitemsBulkTerminateJob")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PatchTaskmanagementWorkitemsBulkTerminateJob")
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

	var successPayload *Bulkjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulkjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchTaskmanagementWorktype invokes PATCH /api/v2/taskmanagement/worktypes/{worktypeId}
//
// Update the attributes of a worktype
func (a TaskManagementApi) PatchTaskmanagementWorktype(worktypeId string, body Worktypeupdate) (*Worktype, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	defaultReturn := new(Worktype)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->PatchTaskmanagementWorktype")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PatchTaskmanagementWorktype")
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

	var successPayload *Worktype
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Worktype" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchTaskmanagementWorktypeFlowsDatebasedRule invokes PATCH /api/v2/taskmanagement/worktypes/{worktypeId}/flows/datebased/rules/{ruleId}
//
// Update the attributes of a date based rule
func (a TaskManagementApi) PatchTaskmanagementWorktypeFlowsDatebasedRule(worktypeId string, ruleId string, body Workitemdatebasedruleupdate) (*Workitemdatebasedrule, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/datebased/rules/{ruleId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	path = strings.Replace(path, "{ruleId}", url.PathEscape(fmt.Sprintf("%v", ruleId)), -1)
	defaultReturn := new(Workitemdatebasedrule)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->PatchTaskmanagementWorktypeFlowsDatebasedRule")
	}
	// verify the required parameter 'ruleId' is set
	if &ruleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'ruleId' when calling TaskManagementApi->PatchTaskmanagementWorktypeFlowsDatebasedRule")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PatchTaskmanagementWorktypeFlowsDatebasedRule")
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

	var successPayload *Workitemdatebasedrule
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemdatebasedrule" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchTaskmanagementWorktypeFlowsOnattributechangeRule invokes PATCH /api/v2/taskmanagement/worktypes/{worktypeId}/flows/onattributechange/rules/{ruleId}
//
// Update the attributes of a rule
func (a TaskManagementApi) PatchTaskmanagementWorktypeFlowsOnattributechangeRule(worktypeId string, ruleId string, body Workitemonattributechangeruleupdate) (*Workitemonattributechangerule, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/onattributechange/rules/{ruleId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	path = strings.Replace(path, "{ruleId}", url.PathEscape(fmt.Sprintf("%v", ruleId)), -1)
	defaultReturn := new(Workitemonattributechangerule)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->PatchTaskmanagementWorktypeFlowsOnattributechangeRule")
	}
	// verify the required parameter 'ruleId' is set
	if &ruleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'ruleId' when calling TaskManagementApi->PatchTaskmanagementWorktypeFlowsOnattributechangeRule")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PatchTaskmanagementWorktypeFlowsOnattributechangeRule")
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

	var successPayload *Workitemonattributechangerule
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemonattributechangerule" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchTaskmanagementWorktypeFlowsOncreateRule invokes PATCH /api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}
//
// Update the attributes of a rule
func (a TaskManagementApi) PatchTaskmanagementWorktypeFlowsOncreateRule(worktypeId string, ruleId string, body Workitemoncreateruleupdate) (*Workitemoncreaterule, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules/{ruleId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	path = strings.Replace(path, "{ruleId}", url.PathEscape(fmt.Sprintf("%v", ruleId)), -1)
	defaultReturn := new(Workitemoncreaterule)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->PatchTaskmanagementWorktypeFlowsOncreateRule")
	}
	// verify the required parameter 'ruleId' is set
	if &ruleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'ruleId' when calling TaskManagementApi->PatchTaskmanagementWorktypeFlowsOncreateRule")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PatchTaskmanagementWorktypeFlowsOncreateRule")
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

	var successPayload *Workitemoncreaterule
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemoncreaterule" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchTaskmanagementWorktypeStatus invokes PATCH /api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}
//
// Update the attributes of a status
func (a TaskManagementApi) PatchTaskmanagementWorktypeStatus(worktypeId string, statusId string, body Workitemstatusupdate) (*Workitemstatus, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses/{statusId}"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	path = strings.Replace(path, "{statusId}", url.PathEscape(fmt.Sprintf("%v", statusId)), -1)
	defaultReturn := new(Workitemstatus)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->PatchTaskmanagementWorktypeStatus")
	}
	// verify the required parameter 'statusId' is set
	if &statusId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'statusId' when calling TaskManagementApi->PatchTaskmanagementWorktypeStatus")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PatchTaskmanagementWorktypeStatus")
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

	var successPayload *Workitemstatus
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemstatus" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorkbins invokes POST /api/v2/taskmanagement/workbins
//
// Create a workbin
func (a TaskManagementApi) PostTaskmanagementWorkbins(body Workbincreate) (*Workbin, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workbins"
	defaultReturn := new(Workbin)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorkbins")
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

	var successPayload *Workbin
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workbin" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorkbinsQuery invokes POST /api/v2/taskmanagement/workbins/query
//
// Query for workbins
func (a TaskManagementApi) PostTaskmanagementWorkbinsQuery(body Workbinqueryrequest) (*Workbinqueryentitylisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workbins/query"
	defaultReturn := new(Workbinqueryentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorkbinsQuery")
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

	var successPayload *Workbinqueryentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workbinqueryentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorkitemAcdCancel invokes POST /api/v2/taskmanagement/workitems/{workitemId}/acd/cancel
//
// Cancel the assignment process for a workitem that is currently queued for assignment through ACD.
func (a TaskManagementApi) PostTaskmanagementWorkitemAcdCancel(workitemId string) (*Workitem, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}/acd/cancel"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	defaultReturn := new(Workitem)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->PostTaskmanagementWorkitemAcdCancel")
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
	var successPayload *Workitem
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitem" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorkitemDisconnect invokes POST /api/v2/taskmanagement/workitems/{workitemId}/disconnect
//
// Disconnect the assignee of the workitem
func (a TaskManagementApi) PostTaskmanagementWorkitemDisconnect(workitemId string) (*Workitem, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}/disconnect"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	defaultReturn := new(Workitem)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->PostTaskmanagementWorkitemDisconnect")
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
	var successPayload *Workitem
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitem" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorkitemTerminate invokes POST /api/v2/taskmanagement/workitems/{workitemId}/terminate
//
// Terminate a workitem
func (a TaskManagementApi) PostTaskmanagementWorkitemTerminate(workitemId string, body Workitemterminate) (*Workitem, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/{workitemId}/terminate"
	path = strings.Replace(path, "{workitemId}", url.PathEscape(fmt.Sprintf("%v", workitemId)), -1)
	defaultReturn := new(Workitem)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'workitemId' is set
	if &workitemId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workitemId' when calling TaskManagementApi->PostTaskmanagementWorkitemTerminate")
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

	var successPayload *Workitem
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitem" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorkitems invokes POST /api/v2/taskmanagement/workitems
//
// Create a workitem
func (a TaskManagementApi) PostTaskmanagementWorkitems(body Workitemcreate) (*Workitem, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems"
	defaultReturn := new(Workitem)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorkitems")
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

	var successPayload *Workitem
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitem" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorkitemsBulkAddJobs invokes POST /api/v2/taskmanagement/workitems/bulk/add/jobs
//
// Create a workitem bulk add job.
func (a TaskManagementApi) PostTaskmanagementWorkitemsBulkAddJobs(body Bulkjobaddrequest) (*Bulkjob, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/bulk/add/jobs"
	defaultReturn := new(Bulkjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorkitemsBulkAddJobs")
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

	var successPayload *Bulkjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulkjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorkitemsBulkTerminateJobs invokes POST /api/v2/taskmanagement/workitems/bulk/terminate/jobs
//
// Create a workitem bulk terminate job.
func (a TaskManagementApi) PostTaskmanagementWorkitemsBulkTerminateJobs(body Bulkjobterminaterequest) (*Bulkjob, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/bulk/terminate/jobs"
	defaultReturn := new(Bulkjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorkitemsBulkTerminateJobs")
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

	var successPayload *Bulkjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulkjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorkitemsQuery invokes POST /api/v2/taskmanagement/workitems/query
//
// Query for workitems
//
// This query requires at least one EQ filter on the workbinId, assigneeId or typeId attributes.
//
// Preview: PostTaskmanagementWorkitemsQuery is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a TaskManagementApi) PostTaskmanagementWorkitemsQuery(body Workitemquerypostrequest) (*Workitempostqueryentitylisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/query"
	defaultReturn := new(Workitempostqueryentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorkitemsQuery")
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

	var successPayload *Workitempostqueryentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitempostqueryentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorkitemsQueryJobs invokes POST /api/v2/taskmanagement/workitems/query/jobs
//
// Create a workitem query job
func (a TaskManagementApi) PostTaskmanagementWorkitemsQueryJobs(body Workitemqueryjobcreate) (*Workitemqueryjobresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/query/jobs"
	defaultReturn := new(Workitemqueryjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorkitemsQueryJobs")
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

	var successPayload *Workitemqueryjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemqueryjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorkitemsSchemas invokes POST /api/v2/taskmanagement/workitems/schemas
//
// Create a schema
func (a TaskManagementApi) PostTaskmanagementWorkitemsSchemas(body Dataschema) (*Dataschema, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/schemas"
	defaultReturn := new(Dataschema)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorkitemsSchemas")
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

	var successPayload *Dataschema
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dataschema" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorktypeFlowsDatebasedRules invokes POST /api/v2/taskmanagement/worktypes/{worktypeId}/flows/datebased/rules
//
// Add a date based rule to a worktype
func (a TaskManagementApi) PostTaskmanagementWorktypeFlowsDatebasedRules(worktypeId string, body Workitemdatebasedrulecreate) (*Workitemdatebasedrule, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/datebased/rules"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	defaultReturn := new(Workitemdatebasedrule)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->PostTaskmanagementWorktypeFlowsDatebasedRules")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorktypeFlowsDatebasedRules")
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

	var successPayload *Workitemdatebasedrule
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemdatebasedrule" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorktypeFlowsOnattributechangeRules invokes POST /api/v2/taskmanagement/worktypes/{worktypeId}/flows/onattributechange/rules
//
// Add an attribute-change rule to a worktype
func (a TaskManagementApi) PostTaskmanagementWorktypeFlowsOnattributechangeRules(worktypeId string, body Workitemonattributechangerulecreate) (*Workitemonattributechangerule, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/onattributechange/rules"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	defaultReturn := new(Workitemonattributechangerule)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->PostTaskmanagementWorktypeFlowsOnattributechangeRules")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorktypeFlowsOnattributechangeRules")
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

	var successPayload *Workitemonattributechangerule
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemonattributechangerule" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorktypeFlowsOncreateRules invokes POST /api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules
//
// Add an on-create rule to a worktype
func (a TaskManagementApi) PostTaskmanagementWorktypeFlowsOncreateRules(worktypeId string, body Workitemoncreaterulecreate) (*Workitemoncreaterule, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/flows/oncreate/rules"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	defaultReturn := new(Workitemoncreaterule)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->PostTaskmanagementWorktypeFlowsOncreateRules")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorktypeFlowsOncreateRules")
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

	var successPayload *Workitemoncreaterule
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemoncreaterule" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorktypeStatuses invokes POST /api/v2/taskmanagement/worktypes/{worktypeId}/statuses
//
// Add a status to a worktype
func (a TaskManagementApi) PostTaskmanagementWorktypeStatuses(worktypeId string, body Workitemstatuscreate) (*Workitemstatus, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/{worktypeId}/statuses"
	path = strings.Replace(path, "{worktypeId}", url.PathEscape(fmt.Sprintf("%v", worktypeId)), -1)
	defaultReturn := new(Workitemstatus)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'worktypeId' is set
	if &worktypeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'worktypeId' when calling TaskManagementApi->PostTaskmanagementWorktypeStatuses")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorktypeStatuses")
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

	var successPayload *Workitemstatus
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workitemstatus" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorktypes invokes POST /api/v2/taskmanagement/worktypes
//
// Create a worktype
func (a TaskManagementApi) PostTaskmanagementWorktypes(body Worktypecreate) (*Worktype, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes"
	defaultReturn := new(Worktype)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorktypes")
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

	var successPayload *Worktype
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Worktype" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTaskmanagementWorktypesQuery invokes POST /api/v2/taskmanagement/worktypes/query
//
// Query for worktypes
func (a TaskManagementApi) PostTaskmanagementWorktypesQuery(body Worktypequeryrequest) (*Worktypequeryentitylisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/worktypes/query"
	defaultReturn := new(Worktypequeryentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PostTaskmanagementWorktypesQuery")
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

	var successPayload *Worktypequeryentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Worktypequeryentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTaskmanagementWorkitemsSchema invokes PUT /api/v2/taskmanagement/workitems/schemas/{schemaId}
//
// Update a schema
func (a TaskManagementApi) PutTaskmanagementWorkitemsSchema(schemaId string, body Dataschema) (*Dataschema, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/taskmanagement/workitems/schemas/{schemaId}"
	path = strings.Replace(path, "{schemaId}", url.PathEscape(fmt.Sprintf("%v", schemaId)), -1)
	defaultReturn := new(Dataschema)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'schemaId' is set
	if &schemaId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'schemaId' when calling TaskManagementApi->PutTaskmanagementWorkitemsSchema")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TaskManagementApi->PutTaskmanagementWorkitemsSchema")
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

	var successPayload *Dataschema
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dataschema" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

