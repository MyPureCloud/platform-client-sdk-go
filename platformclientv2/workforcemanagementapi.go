package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
	"time"
)

// WorkforceManagementApi provides functions for API endpoints
type WorkforceManagementApi struct {
	Configuration *Configuration
}

// NewWorkforceManagementApi creates an API instance using the default configuration
func NewWorkforceManagementApi() *WorkforceManagementApi {
	fmt.Sprintf(strings.Title(""), "")
	config := GetDefaultConfiguration()
	return &WorkforceManagementApi{
		Configuration: config,
	}
}

// NewWorkforceManagementApiWithConfig creates an API instance using the provided configuration
func NewWorkforceManagementApiWithConfig(config *Configuration) *WorkforceManagementApi {
	return &WorkforceManagementApi{
		Configuration: config,
	}
}

// DeleteWorkforcemanagementBusinessunit invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}
//
// Delete business unit
//
// A business unit cannot be deleted if it contains one or more management units
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunit(businessUnitId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunit")
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

// DeleteWorkforcemanagementBusinessunitActivitycode invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}
//
// Deletes an activity code
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitActivitycode(businessUnitId string, activityCodeId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{activityCodeId}", url.PathEscape(fmt.Sprintf("%v", activityCodeId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitActivitycode")
	}
	// verify the required parameter 'activityCodeId' is set
	if &activityCodeId == nil {
		// false
		return nil, errors.New("Missing required parameter 'activityCodeId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitActivitycode")
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

// DeleteWorkforcemanagementBusinessunitPlanninggroup invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}
//
// Deletes the planning group
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitPlanninggroup(businessUnitId string, planningGroupId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{planningGroupId}", url.PathEscape(fmt.Sprintf("%v", planningGroupId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitPlanninggroup")
	}
	// verify the required parameter 'planningGroupId' is set
	if &planningGroupId == nil {
		// false
		return nil, errors.New("Missing required parameter 'planningGroupId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitPlanninggroup")
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

// DeleteWorkforcemanagementBusinessunitSchedulingRun invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}
//
// Cancel a scheduling run
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitSchedulingRun(businessUnitId string, runId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{runId}", url.PathEscape(fmt.Sprintf("%v", runId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitSchedulingRun")
	}
	// verify the required parameter 'runId' is set
	if &runId == nil {
		// false
		return nil, errors.New("Missing required parameter 'runId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitSchedulingRun")
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

// DeleteWorkforcemanagementBusinessunitServicegoaltemplate invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates/{serviceGoalTemplateId}
//
// Delete a service goal template
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitServicegoaltemplate(businessUnitId string, serviceGoalTemplateId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates/{serviceGoalTemplateId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{serviceGoalTemplateId}", url.PathEscape(fmt.Sprintf("%v", serviceGoalTemplateId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitServicegoaltemplate")
	}
	// verify the required parameter 'serviceGoalTemplateId' is set
	if &serviceGoalTemplateId == nil {
		// false
		return nil, errors.New("Missing required parameter 'serviceGoalTemplateId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitServicegoaltemplate")
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

// DeleteWorkforcemanagementBusinessunitStaffinggroup invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups/{staffingGroupId}
//
// Deletes a staffing group
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitStaffinggroup(businessUnitId string, staffingGroupId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups/{staffingGroupId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{staffingGroupId}", url.PathEscape(fmt.Sprintf("%v", staffingGroupId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitStaffinggroup")
	}
	// verify the required parameter 'staffingGroupId' is set
	if &staffingGroupId == nil {
		// false
		return nil, errors.New("Missing required parameter 'staffingGroupId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitStaffinggroup")
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

// DeleteWorkforcemanagementBusinessunitTimeofflimit invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/timeofflimits/{timeOffLimitId}
//
// Deletes a time-off limit object
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitTimeofflimit(businessUnitId string, timeOffLimitId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/timeofflimits/{timeOffLimitId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{timeOffLimitId}", url.PathEscape(fmt.Sprintf("%v", timeOffLimitId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitTimeofflimit")
	}
	// verify the required parameter 'timeOffLimitId' is set
	if &timeOffLimitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'timeOffLimitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitTimeofflimit")
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

// DeleteWorkforcemanagementBusinessunitTimeoffplan invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/timeoffplans/{timeOffPlanId}
//
// Deletes a time-off plan
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitTimeoffplan(businessUnitId string, timeOffPlanId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/timeoffplans/{timeOffPlanId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{timeOffPlanId}", url.PathEscape(fmt.Sprintf("%v", timeOffPlanId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitTimeoffplan")
	}
	// verify the required parameter 'timeOffPlanId' is set
	if &timeOffPlanId == nil {
		// false
		return nil, errors.New("Missing required parameter 'timeOffPlanId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitTimeoffplan")
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

// DeleteWorkforcemanagementBusinessunitWeekSchedule invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}
//
// Delete a schedule
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitWeekSchedule(businessUnitId string, weekId time.Time, scheduleId string) (*Buasyncscheduleresponse, *APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	defaultReturn := new(Buasyncscheduleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWeekSchedule")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWeekSchedule")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWeekSchedule")
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
	var successPayload *Buasyncscheduleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buasyncscheduleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// DeleteWorkforcemanagementBusinessunitWeekShorttermforecast invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}
//
// Delete a short term forecast
//
// Must not be tied to any schedules
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitWeekShorttermforecast(businessUnitId string, weekDateId time.Time, forecastId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	path = strings.Replace(path, "{forecastId}", url.PathEscape(fmt.Sprintf("%v", forecastId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWeekShorttermforecast")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWeekShorttermforecast")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// false
		return nil, errors.New("Missing required parameter 'forecastId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWeekShorttermforecast")
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

// DeleteWorkforcemanagementBusinessunitWorkplanbid invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}
//
// Delete a work plan bid
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitWorkplanbid(businessUnitId string, bidId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWorkplanbid")
	}
	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWorkplanbid")
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

// DeleteWorkforcemanagementBusinessunitWorkplanbidGroup invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups/{bidGroupId}
//
// Delete a bid group by bid group Id
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitWorkplanbidGroup(businessUnitId string, bidId string, bidGroupId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups/{bidGroupId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	path = strings.Replace(path, "{bidGroupId}", url.PathEscape(fmt.Sprintf("%v", bidGroupId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWorkplanbidGroup")
	}
	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWorkplanbidGroup")
	}
	// verify the required parameter 'bidGroupId' is set
	if &bidGroupId == nil {
		// false
		return nil, errors.New("Missing required parameter 'bidGroupId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWorkplanbidGroup")
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

// DeleteWorkforcemanagementCalendarUrlIcs invokes DELETE /api/v2/workforcemanagement/calendar/url/ics
//
// Disable generated calendar link for the current user
func (a WorkforceManagementApi) DeleteWorkforcemanagementCalendarUrlIcs() (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/calendar/url/ics"
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
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

// DeleteWorkforcemanagementManagementunit invokes DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}
//
// Delete management unit
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunit(managementUnitId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunit")
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

// DeleteWorkforcemanagementManagementunitTimeofflimit invokes DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}
//
// Deletes a time off limit object
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunitTimeofflimit(managementUnitId string, timeOffLimitId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{timeOffLimitId}", url.PathEscape(fmt.Sprintf("%v", timeOffLimitId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitTimeofflimit")
	}
	// verify the required parameter 'timeOffLimitId' is set
	if &timeOffLimitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'timeOffLimitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitTimeofflimit")
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

// DeleteWorkforcemanagementManagementunitTimeoffplan invokes DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffplans/{timeOffPlanId}
//
// Deletes a time off plan
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunitTimeoffplan(managementUnitId string, timeOffPlanId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffplans/{timeOffPlanId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{timeOffPlanId}", url.PathEscape(fmt.Sprintf("%v", timeOffPlanId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitTimeoffplan")
	}
	// verify the required parameter 'timeOffPlanId' is set
	if &timeOffPlanId == nil {
		// false
		return nil, errors.New("Missing required parameter 'timeOffPlanId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitTimeoffplan")
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

// DeleteWorkforcemanagementManagementunitWorkplan invokes DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}
//
// Delete a work plan
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunitWorkplan(managementUnitId string, workPlanId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{workPlanId}", url.PathEscape(fmt.Sprintf("%v", workPlanId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitWorkplan")
	}
	// verify the required parameter 'workPlanId' is set
	if &workPlanId == nil {
		// false
		return nil, errors.New("Missing required parameter 'workPlanId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitWorkplan")
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

// DeleteWorkforcemanagementManagementunitWorkplanrotation invokes DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}
//
// Delete a work plan rotation
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunitWorkplanrotation(managementUnitId string, workPlanRotationId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{workPlanRotationId}", url.PathEscape(fmt.Sprintf("%v", workPlanRotationId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitWorkplanrotation")
	}
	// verify the required parameter 'workPlanRotationId' is set
	if &workPlanRotationId == nil {
		// false
		return nil, errors.New("Missing required parameter 'workPlanRotationId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitWorkplanrotation")
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

// GetWorkforcemanagementAdherence invokes GET /api/v2/workforcemanagement/adherence
//
// Get a list of UserScheduleAdherence records for the requested users
func (a WorkforceManagementApi) GetWorkforcemanagementAdherence(userId []string) ([]Userscheduleadherence, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/adherence"
	defaultReturn := make([]Userscheduleadherence, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'userId' is set
	if &userId == nil {
		// true
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling WorkforceManagementApi->GetWorkforcemanagementAdherence")
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
	
	queryParams["userId"] = a.Configuration.APIClient.ParameterToString(userId, "multi")
	

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
	var successPayload []Userscheduleadherence
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "[]Userscheduleadherence" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAdherenceExplanation invokes GET /api/v2/workforcemanagement/adherence/explanations/{explanationId}
//
// Get an adherence explanation for the current user
func (a WorkforceManagementApi) GetWorkforcemanagementAdherenceExplanation(explanationId string) (*Adherenceexplanationresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/adherence/explanations/{explanationId}"
	path = strings.Replace(path, "{explanationId}", url.PathEscape(fmt.Sprintf("%v", explanationId)), -1)
	defaultReturn := new(Adherenceexplanationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'explanationId' is set
	if &explanationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'explanationId' when calling WorkforceManagementApi->GetWorkforcemanagementAdherenceExplanation")
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
	var successPayload *Adherenceexplanationresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Adherenceexplanationresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAdherenceExplanationsJob invokes GET /api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}
//
// Query the status of an adherence explanation operation. Only the user who started the operation can query the status
//
// Job details are only retained if the initial request returned a 202 ACCEPTED response
func (a WorkforceManagementApi) GetWorkforcemanagementAdherenceExplanationsJob(jobId string) (*Adherenceexplanationjob, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/adherence/explanations/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Adherenceexplanationjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementAdherenceExplanationsJob")
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
	var successPayload *Adherenceexplanationjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Adherenceexplanationjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAdherenceHistoricalBulkJob invokes GET /api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}
//
// Request to fetch the status of the historical adherence bulk job. Only the user who started the operation can query the status
//
// Job details are only retained if the initial request returned a 202 ACCEPTED response
func (a WorkforceManagementApi) GetWorkforcemanagementAdherenceHistoricalBulkJob(jobId string) (*Wfmhistoricaladherencebulkresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/adherence/historical/bulk/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Wfmhistoricaladherencebulkresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementAdherenceHistoricalBulkJob")
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
	var successPayload *Wfmhistoricaladherencebulkresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmhistoricaladherencebulkresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAdherenceHistoricalJob invokes GET /api/v2/workforcemanagement/adherence/historical/jobs/{jobId}
//
// Query the status of a historical adherence request operation. Only the user who started the operation can query the status
//
// Job details are only retained if the initial request returned a 202 ACCEPTED response
func (a WorkforceManagementApi) GetWorkforcemanagementAdherenceHistoricalJob(jobId string) (*Wfmhistoricaladherenceresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/adherence/historical/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Wfmhistoricaladherenceresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementAdherenceHistoricalJob")
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
	var successPayload *Wfmhistoricaladherenceresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmhistoricaladherenceresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAgentAdherenceExplanation invokes GET /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}
//
// Get an adherence explanation
func (a WorkforceManagementApi) GetWorkforcemanagementAgentAdherenceExplanation(agentId string, explanationId string) (*Adherenceexplanationresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}"
	path = strings.Replace(path, "{agentId}", url.PathEscape(fmt.Sprintf("%v", agentId)), -1)
	path = strings.Replace(path, "{explanationId}", url.PathEscape(fmt.Sprintf("%v", explanationId)), -1)
	defaultReturn := new(Adherenceexplanationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'agentId' is set
	if &agentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'agentId' when calling WorkforceManagementApi->GetWorkforcemanagementAgentAdherenceExplanation")
	}
	// verify the required parameter 'explanationId' is set
	if &explanationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'explanationId' when calling WorkforceManagementApi->GetWorkforcemanagementAgentAdherenceExplanation")
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
	var successPayload *Adherenceexplanationresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Adherenceexplanationresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAgentManagementunit invokes GET /api/v2/workforcemanagement/agents/{agentId}/managementunit
//
// Get the management unit to which the agent belongs
func (a WorkforceManagementApi) GetWorkforcemanagementAgentManagementunit(agentId string) (*Agentmanagementunitreference, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/agents/{agentId}/managementunit"
	path = strings.Replace(path, "{agentId}", url.PathEscape(fmt.Sprintf("%v", agentId)), -1)
	defaultReturn := new(Agentmanagementunitreference)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'agentId' is set
	if &agentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'agentId' when calling WorkforceManagementApi->GetWorkforcemanagementAgentManagementunit")
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
	var successPayload *Agentmanagementunitreference
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Agentmanagementunitreference" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAgentsMeManagementunit invokes GET /api/v2/workforcemanagement/agents/me/managementunit
//
// Get the management unit to which the currently logged in agent belongs
func (a WorkforceManagementApi) GetWorkforcemanagementAgentsMeManagementunit() (*Agentmanagementunitreference, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/agents/me/managementunit"
	defaultReturn := new(Agentmanagementunitreference)
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
	var successPayload *Agentmanagementunitreference
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Agentmanagementunitreference" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAlternativeshiftsOffersJob invokes GET /api/v2/workforcemanagement/alternativeshifts/offers/jobs/{jobId}
//
// Query the status of an alternative shift offers operation. Only the user who started the operation can query the status
//
// Job details are only retained if the initial request returned a 202 ACCEPTED response
func (a WorkforceManagementApi) GetWorkforcemanagementAlternativeshiftsOffersJob(jobId string) (*Alternativeshiftjobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/alternativeshifts/offers/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Alternativeshiftjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementAlternativeshiftsOffersJob")
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
	var successPayload *Alternativeshiftjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshiftjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAlternativeshiftsOffersSearchJob invokes GET /api/v2/workforcemanagement/alternativeshifts/offers/search/jobs/{jobId}
//
// Query the status of an alternative shift search offers operation. Only the user who started the operation can query the status
//
// Job details are only retained if the initial request returned a 202 ACCEPTED response
func (a WorkforceManagementApi) GetWorkforcemanagementAlternativeshiftsOffersSearchJob(jobId string) (*Alternativeshiftjobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/alternativeshifts/offers/search/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Alternativeshiftjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementAlternativeshiftsOffersSearchJob")
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
	var successPayload *Alternativeshiftjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshiftjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAlternativeshiftsSettings invokes GET /api/v2/workforcemanagement/alternativeshifts/settings
//
// Get alternative shifts settings from the current logged in agent’s business unit
func (a WorkforceManagementApi) GetWorkforcemanagementAlternativeshiftsSettings() (*Alternativeshiftbusettingsresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/alternativeshifts/settings"
	defaultReturn := new(Alternativeshiftbusettingsresponse)
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
	var successPayload *Alternativeshiftbusettingsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshiftbusettingsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAlternativeshiftsTrade invokes GET /api/v2/workforcemanagement/alternativeshifts/trades/{tradeId}
//
// Get my alternative shift trade by trade ID
func (a WorkforceManagementApi) GetWorkforcemanagementAlternativeshiftsTrade(tradeId string) (*Alternativeshifttraderesponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/alternativeshifts/trades/{tradeId}"
	path = strings.Replace(path, "{tradeId}", url.PathEscape(fmt.Sprintf("%v", tradeId)), -1)
	defaultReturn := new(Alternativeshifttraderesponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'tradeId' is set
	if &tradeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'tradeId' when calling WorkforceManagementApi->GetWorkforcemanagementAlternativeshiftsTrade")
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
	var successPayload *Alternativeshifttraderesponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshifttraderesponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAlternativeshiftsTrades invokes GET /api/v2/workforcemanagement/alternativeshifts/trades
//
// Get a list of my alternative shifts trades
func (a WorkforceManagementApi) GetWorkforcemanagementAlternativeshiftsTrades(forceAsync bool) (*Listalternativeshifttradesresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/alternativeshifts/trades"
	defaultReturn := new(Listalternativeshifttradesresponse)
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
	
	queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, "")
	

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
	var successPayload *Listalternativeshifttradesresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Listalternativeshifttradesresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAlternativeshiftsTradesJob invokes GET /api/v2/workforcemanagement/alternativeshifts/trades/jobs/{jobId}
//
// Query the status of an alternative shift trades operation. Only the user who started the operation can query the status
//
// Job details are only retained if the initial request returned a 202 ACCEPTED response
func (a WorkforceManagementApi) GetWorkforcemanagementAlternativeshiftsTradesJob(jobId string) (*Alternativeshiftjobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/alternativeshifts/trades/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Alternativeshiftjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementAlternativeshiftsTradesJob")
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
	var successPayload *Alternativeshiftjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshiftjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementAlternativeshiftsTradesStateJob invokes GET /api/v2/workforcemanagement/alternativeshifts/trades/state/jobs/{jobId}
//
// Query the status of an alternative shift trade state operation. Only the user who started the operation can query the status
//
// Job details are only retained if the initial request returned a 202 ACCEPTED response
func (a WorkforceManagementApi) GetWorkforcemanagementAlternativeshiftsTradesStateJob(jobId string) (*Alternativeshiftjobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/alternativeshifts/trades/state/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Alternativeshiftjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementAlternativeshiftsTradesStateJob")
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
	var successPayload *Alternativeshiftjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshiftjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunit invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}
//
// Get business unit
//
// Expanding \&quot;settings\&quot; will retrieve all settings.  All other expands will retrieve only the requested settings field(s).
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunit(businessUnitId string, expand []string) (*Businessunitresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Businessunitresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunit")
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
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	

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
	var successPayload *Businessunitresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Businessunitresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitActivitycode invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}
//
// Get an activity code
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitActivitycode(businessUnitId string, activityCodeId string) (*Businessunitactivitycode, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{activityCodeId}", url.PathEscape(fmt.Sprintf("%v", activityCodeId)), -1)
	defaultReturn := new(Businessunitactivitycode)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivitycode")
	}
	// verify the required parameter 'activityCodeId' is set
	if &activityCodeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'activityCodeId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivitycode")
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
	var successPayload *Businessunitactivitycode
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Businessunitactivitycode" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitActivitycodes invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes
//
// Get activity codes
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitActivitycodes(businessUnitId string, forceDownloadService bool) (*Businessunitactivitycodelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Businessunitactivitycodelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivitycodes")
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
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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
	var successPayload *Businessunitactivitycodelisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Businessunitactivitycodelisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitActivityplan invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/{activityPlanId}
//
// Get an activity plan
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitActivityplan(businessUnitId string, activityPlanId string) (*Activityplanresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/{activityPlanId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{activityPlanId}", url.PathEscape(fmt.Sprintf("%v", activityPlanId)), -1)
	defaultReturn := new(Activityplanresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivityplan")
	}
	// verify the required parameter 'activityPlanId' is set
	if &activityPlanId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'activityPlanId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivityplan")
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
	var successPayload *Activityplanresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Activityplanresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitActivityplanRunsJob invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/{activityPlanId}/runs/jobs/{jobId}
//
// Gets an activity plan run job
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitActivityplanRunsJob(businessUnitId string, activityPlanId string, jobId string) (*Activityplanrunjobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/{activityPlanId}/runs/jobs/{jobId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{activityPlanId}", url.PathEscape(fmt.Sprintf("%v", activityPlanId)), -1)
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Activityplanrunjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivityplanRunsJob")
	}
	// verify the required parameter 'activityPlanId' is set
	if &activityPlanId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'activityPlanId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivityplanRunsJob")
	}
	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivityplanRunsJob")
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
	var successPayload *Activityplanrunjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Activityplanrunjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitActivityplans invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans
//
// Get activity plans
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitActivityplans(businessUnitId string, state string) (*Activityplanlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Activityplanlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivityplans")
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
	
	queryParams["state"] = a.Configuration.APIClient.ParameterToString(state, "")
	

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
	var successPayload *Activityplanlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Activityplanlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitActivityplansJobs invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/jobs
//
// Gets the latest job for all activity plans in the business unit
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitActivityplansJobs(businessUnitId string) (*Activityplanjoblisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/jobs"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Activityplanjoblisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivityplansJobs")
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
	var successPayload *Activityplanjoblisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Activityplanjoblisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitAlternativeshiftsSettings invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/alternativeshifts/settings
//
// Get alternative shifts settings for a business unit
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitAlternativeshiftsSettings(businessUnitId string) (*Alternativeshiftbusettingsresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/alternativeshifts/settings"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Alternativeshiftbusettingsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitAlternativeshiftsSettings")
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
	var successPayload *Alternativeshiftbusettingsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshiftbusettingsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitAlternativeshiftsTrade invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/alternativeshifts/trades/{tradeId}
//
// Get an alternative shifts trade in a business unit for a given trade ID
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitAlternativeshiftsTrade(businessUnitId string, tradeId string) (*Alternativeshifttraderesponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/alternativeshifts/trades/{tradeId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{tradeId}", url.PathEscape(fmt.Sprintf("%v", tradeId)), -1)
	defaultReturn := new(Alternativeshifttraderesponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitAlternativeshiftsTrade")
	}
	// verify the required parameter 'tradeId' is set
	if &tradeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'tradeId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitAlternativeshiftsTrade")
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
	var successPayload *Alternativeshifttraderesponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshifttraderesponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitAlternativeshiftsTradesSearchJob invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/alternativeshifts/trades/search/jobs/{jobId}
//
// Query the status of an alternative shift search trade operation. Only the user who started the operation can query the status
//
// Job details are only retained if the initial request returned a 202 ACCEPTED response
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitAlternativeshiftsTradesSearchJob(businessUnitId string, jobId string) (*Bualternativeshiftjobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/alternativeshifts/trades/search/jobs/{jobId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Bualternativeshiftjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitAlternativeshiftsTradesSearchJob")
	}
	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitAlternativeshiftsTradesSearchJob")
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
	var successPayload *Bualternativeshiftjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bualternativeshiftjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitIntradayPlanninggroups invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/intraday/planninggroups
//
// Get intraday planning groups for the given date
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitIntradayPlanninggroups(businessUnitId string, date time.Time) (*Wfmintradayplanninggrouplisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/intraday/planninggroups"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Wfmintradayplanninggrouplisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitIntradayPlanninggroups")
	}
	// verify the required parameter 'date' is set
	if &date == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'date' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitIntradayPlanninggroups")
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
	
	queryParams["date"] = a.Configuration.APIClient.ParameterToString(date, "")
	

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
	var successPayload *Wfmintradayplanninggrouplisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmintradayplanninggrouplisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitManagementunits invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits
//
// Get all authorized management units in the business unit
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitManagementunits(businessUnitId string, feature string, divisionId string) (*Managementunitlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Managementunitlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitManagementunits")
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
	
	queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, "")
	
	queryParams["divisionId"] = a.Configuration.APIClient.ParameterToString(divisionId, "")
	

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
	var successPayload *Managementunitlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Managementunitlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitPlanninggroup invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}
//
// Get a planning group
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitPlanninggroup(businessUnitId string, planningGroupId string) (*Planninggroup, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{planningGroupId}", url.PathEscape(fmt.Sprintf("%v", planningGroupId)), -1)
	defaultReturn := new(Planninggroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitPlanninggroup")
	}
	// verify the required parameter 'planningGroupId' is set
	if &planningGroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'planningGroupId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitPlanninggroup")
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
	var successPayload *Planninggroup
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Planninggroup" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitPlanninggroups invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups
//
// Gets list of planning groups
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitPlanninggroups(businessUnitId string) (*Planninggrouplist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Planninggrouplist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitPlanninggroups")
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
	var successPayload *Planninggrouplist
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Planninggrouplist" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitSchedulingRun invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}
//
// Get a scheduling run
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitSchedulingRun(businessUnitId string, runId string) (*Buschedulerun, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{runId}", url.PathEscape(fmt.Sprintf("%v", runId)), -1)
	defaultReturn := new(Buschedulerun)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitSchedulingRun")
	}
	// verify the required parameter 'runId' is set
	if &runId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'runId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitSchedulingRun")
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
	var successPayload *Buschedulerun
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buschedulerun" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitSchedulingRunResult invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}/result
//
// Get the result of a rescheduling operation
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitSchedulingRunResult(businessUnitId string, runId string, managementUnitIds []string, expand []string) (*Burescheduleresult, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}/result"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{runId}", url.PathEscape(fmt.Sprintf("%v", runId)), -1)
	defaultReturn := new(Burescheduleresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitSchedulingRunResult")
	}
	// verify the required parameter 'runId' is set
	if &runId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'runId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitSchedulingRunResult")
	}
	// verify the required parameter 'managementUnitIds' is set
	if &managementUnitIds == nil {
		// true
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitIds' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitSchedulingRunResult")
	}
	// verify the required parameter 'expand' is set
	if &expand == nil {
		// true
		return defaultReturn, nil, errors.New("Missing required parameter 'expand' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitSchedulingRunResult")
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
	
	queryParams["managementUnitIds"] = a.Configuration.APIClient.ParameterToString(managementUnitIds, "multi")
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	

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
	var successPayload *Burescheduleresult
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Burescheduleresult" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitSchedulingRuns invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs
//
// Get the list of scheduling runs
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitSchedulingRuns(businessUnitId string) (*Buschedulerunlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Buschedulerunlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitSchedulingRuns")
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
	var successPayload *Buschedulerunlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buschedulerunlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitServicegoaltemplate invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates/{serviceGoalTemplateId}
//
// Get a service goal template
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitServicegoaltemplate(businessUnitId string, serviceGoalTemplateId string, expand []string) (*Servicegoaltemplate, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates/{serviceGoalTemplateId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{serviceGoalTemplateId}", url.PathEscape(fmt.Sprintf("%v", serviceGoalTemplateId)), -1)
	defaultReturn := new(Servicegoaltemplate)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitServicegoaltemplate")
	}
	// verify the required parameter 'serviceGoalTemplateId' is set
	if &serviceGoalTemplateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'serviceGoalTemplateId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitServicegoaltemplate")
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
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	

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
	var successPayload *Servicegoaltemplate
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Servicegoaltemplate" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitServicegoaltemplates invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates
//
// Gets list of service goal templates
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitServicegoaltemplates(businessUnitId string, expand []string) (*Servicegoaltemplatelist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Servicegoaltemplatelist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitServicegoaltemplates")
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
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	

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
	var successPayload *Servicegoaltemplatelist
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Servicegoaltemplatelist" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitStaffinggroup invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups/{staffingGroupId}
//
// Gets a staffing group
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitStaffinggroup(businessUnitId string, staffingGroupId string) (*Staffinggroupresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups/{staffingGroupId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{staffingGroupId}", url.PathEscape(fmt.Sprintf("%v", staffingGroupId)), -1)
	defaultReturn := new(Staffinggroupresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitStaffinggroup")
	}
	// verify the required parameter 'staffingGroupId' is set
	if &staffingGroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'staffingGroupId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitStaffinggroup")
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
	var successPayload *Staffinggroupresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Staffinggroupresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitStaffinggroups invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups
//
// Gets a list of staffing groups
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitStaffinggroups(businessUnitId string, managementUnitId string) (*Staffinggrouplisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Staffinggrouplisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitStaffinggroups")
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
	
	queryParams["managementUnitId"] = a.Configuration.APIClient.ParameterToString(managementUnitId, "")
	

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
	var successPayload *Staffinggrouplisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Staffinggrouplisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitTimeofflimit invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/timeofflimits/{timeOffLimitId}
//
// Gets a time-off limit object
//
// Returns properties of time-off limit object, but not daily values
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitTimeofflimit(businessUnitId string, timeOffLimitId string) (*Butimeofflimitresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/timeofflimits/{timeOffLimitId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{timeOffLimitId}", url.PathEscape(fmt.Sprintf("%v", timeOffLimitId)), -1)
	defaultReturn := new(Butimeofflimitresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitTimeofflimit")
	}
	// verify the required parameter 'timeOffLimitId' is set
	if &timeOffLimitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffLimitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitTimeofflimit")
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
	var successPayload *Butimeofflimitresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Butimeofflimitresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitTimeofflimits invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/timeofflimits
//
// Gets a list of time-off limit objects
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitTimeofflimits(businessUnitId string, managementUnitId string) (*Butimeofflimitlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/timeofflimits"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Butimeofflimitlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitTimeofflimits")
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
	
	queryParams["managementUnitId"] = a.Configuration.APIClient.ParameterToString(managementUnitId, "")
	

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
	var successPayload *Butimeofflimitlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Butimeofflimitlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitTimeoffplan invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/timeoffplans/{timeOffPlanId}
//
// Gets a time-off plan
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitTimeoffplan(businessUnitId string, timeOffPlanId string) (*Butimeoffplanresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/timeoffplans/{timeOffPlanId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{timeOffPlanId}", url.PathEscape(fmt.Sprintf("%v", timeOffPlanId)), -1)
	defaultReturn := new(Butimeoffplanresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitTimeoffplan")
	}
	// verify the required parameter 'timeOffPlanId' is set
	if &timeOffPlanId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffPlanId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitTimeoffplan")
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
	var successPayload *Butimeoffplanresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Butimeoffplanresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitTimeoffplans invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/timeoffplans
//
// Gets a list of time-off plans
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitTimeoffplans(businessUnitId string, managementUnitId string, forceDownloadService bool) (*Butimeoffplanlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/timeoffplans"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Butimeoffplanlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitTimeoffplans")
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
	
	queryParams["managementUnitId"] = a.Configuration.APIClient.ParameterToString(managementUnitId, "")
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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
	var successPayload *Butimeoffplanlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Butimeoffplanlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekSchedule invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}
//
// Get the metadata for the schedule, describing which management units and agents are in the scheduleSchedule data can then be loaded with the query route
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekSchedule(businessUnitId string, weekId time.Time, scheduleId string, expand string) (*Buschedulemetadata, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	defaultReturn := new(Buschedulemetadata)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedule")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedule")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedule")
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
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "")
	

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
	var successPayload *Buschedulemetadata
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buschedulemetadata" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekScheduleGenerationresults invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/generationresults
//
// Get the generation results for a generated schedule
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekScheduleGenerationresults(businessUnitId string, weekId time.Time, scheduleId string) (*Schedulegenerationresult, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/generationresults"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	defaultReturn := new(Schedulegenerationresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleGenerationresults")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleGenerationresults")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleGenerationresults")
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
	var successPayload *Schedulegenerationresult
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Schedulegenerationresult" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekScheduleHeadcountforecast invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/headcountforecast
//
// Get the headcount forecast by planning group for the schedule
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekScheduleHeadcountforecast(businessUnitId string, weekId time.Time, scheduleId string, forceDownload bool) (*Buheadcountforecastresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/headcountforecast"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	defaultReturn := new(Buheadcountforecastresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleHeadcountforecast")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleHeadcountforecast")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleHeadcountforecast")
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
	
	queryParams["forceDownload"] = a.Configuration.APIClient.ParameterToString(forceDownload, "")
	

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
	var successPayload *Buheadcountforecastresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buheadcountforecastresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgent invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}
//
// Loads agent's schedule history.
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgent(businessUnitId string, weekId time.Time, scheduleId string, agentId string) (*Buagentschedulehistoryresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	path = strings.Replace(path, "{agentId}", url.PathEscape(fmt.Sprintf("%v", agentId)), -1)
	defaultReturn := new(Buagentschedulehistoryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgent")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgent")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgent")
	}
	// verify the required parameter 'agentId' is set
	if &agentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'agentId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgent")
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
	var successPayload *Buagentschedulehistoryresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buagentschedulehistoryresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekSchedulePerformancepredictions invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/performancepredictions
//
// Get the performance prediction for the associated schedule
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekSchedulePerformancepredictions(businessUnitId string, weekId string, scheduleId string) (*Performancepredictionresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/performancepredictions"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	defaultReturn := new(Performancepredictionresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedulePerformancepredictions")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedulePerformancepredictions")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedulePerformancepredictions")
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
	var successPayload *Performancepredictionresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Performancepredictionresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculation invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/performancepredictions/recalculations/{recalculationId}
//
// Get recalculated performance prediction result
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculation(businessUnitId string, weekId string, scheduleId string, recalculationId string) (*Performancepredictionrecalculationresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/performancepredictions/recalculations/{recalculationId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	path = strings.Replace(path, "{recalculationId}", url.PathEscape(fmt.Sprintf("%v", recalculationId)), -1)
	defaultReturn := new(Performancepredictionrecalculationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculation")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculation")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculation")
	}
	// verify the required parameter 'recalculationId' is set
	if &recalculationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'recalculationId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculation")
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
	var successPayload *Performancepredictionrecalculationresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Performancepredictionrecalculationresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekSchedules invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules
//
// Get the list of week schedules for the specified week
//
// Use \&quot;recent\&quot; (without quotes) for the &#x60;weekId&#x60; path parameter to fetch all forecasts for +/- 26 weeks from the current date. Response will include any schedule which spans the specified week
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekSchedules(businessUnitId string, weekId string, includeOnlyPublished bool, expand string) (*Buschedulelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	defaultReturn := new(Buschedulelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedules")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedules")
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
	
	queryParams["includeOnlyPublished"] = a.Configuration.APIClient.ParameterToString(includeOnlyPublished, "")
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "")
	

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
	var successPayload *Buschedulelisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buschedulelisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekShorttermforecast invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}
//
// Get a short term forecast
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekShorttermforecast(businessUnitId string, weekDateId time.Time, forecastId string, expand []string) (*Bushorttermforecast, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	path = strings.Replace(path, "{forecastId}", url.PathEscape(fmt.Sprintf("%v", forecastId)), -1)
	defaultReturn := new(Bushorttermforecast)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecast")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecast")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'forecastId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecast")
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
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	

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
	var successPayload *Bushorttermforecast
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bushorttermforecast" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekShorttermforecastData invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/data
//
// Get the result of a short term forecast calculation
//
// Includes modifications unless you pass the doNotApplyModifications query parameter
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekShorttermforecastData(businessUnitId string, weekDateId time.Time, forecastId string, weekNumber int, forceDownloadService bool) (*Buforecastresultresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/data"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	path = strings.Replace(path, "{forecastId}", url.PathEscape(fmt.Sprintf("%v", forecastId)), -1)
	defaultReturn := new(Buforecastresultresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastData")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastData")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'forecastId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastData")
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
	
	queryParams["weekNumber"] = a.Configuration.APIClient.ParameterToString(weekNumber, "")
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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
	var successPayload *Buforecastresultresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buforecastresultresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresults invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/generationresults
//
// Gets the forecast generation results
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresults(businessUnitId string, weekDateId time.Time, forecastId string) (*Buforecastgenerationresult, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/generationresults"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	path = strings.Replace(path, "{forecastId}", url.PathEscape(fmt.Sprintf("%v", forecastId)), -1)
	defaultReturn := new(Buforecastgenerationresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresults")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresults")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'forecastId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresults")
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
	var successPayload *Buforecastgenerationresult
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buforecastgenerationresult" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekShorttermforecastLongtermforecastdata invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/longtermforecastdata
//
// Get the result of a long term forecast calculation
//
// Includes modifications unless you pass the doNotApplyModifications query parameter
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekShorttermforecastLongtermforecastdata(businessUnitId string, weekDateId time.Time, forecastId string, forceDownloadService bool) (*Longtermforecastresultresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/longtermforecastdata"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	path = strings.Replace(path, "{forecastId}", url.PathEscape(fmt.Sprintf("%v", forecastId)), -1)
	defaultReturn := new(Longtermforecastresultresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastLongtermforecastdata")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastLongtermforecastdata")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'forecastId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastLongtermforecastdata")
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
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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
	var successPayload *Longtermforecastresultresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Longtermforecastresultresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroups invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups
//
// Gets the forecast planning group snapshot
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroups(businessUnitId string, weekDateId time.Time, forecastId string) (*Forecastplanninggroupsresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	path = strings.Replace(path, "{forecastId}", url.PathEscape(fmt.Sprintf("%v", forecastId)), -1)
	defaultReturn := new(Forecastplanninggroupsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroups")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroups")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'forecastId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroups")
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
	var successPayload *Forecastplanninggroupsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Forecastplanninggroupsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekShorttermforecastStaffingrequirement invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/staffingrequirement
//
// Get the staffing requirement by planning group for a forecast
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekShorttermforecastStaffingrequirement(businessUnitId string, weekDateId time.Time, forecastId string, weekNumbers []string) (*Buforecaststaffingrequirementsresultresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/staffingrequirement"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	path = strings.Replace(path, "{forecastId}", url.PathEscape(fmt.Sprintf("%v", forecastId)), -1)
	defaultReturn := new(Buforecaststaffingrequirementsresultresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastStaffingrequirement")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastStaffingrequirement")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'forecastId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastStaffingrequirement")
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
	
	queryParams["weekNumbers"] = a.Configuration.APIClient.ParameterToString(weekNumbers, "multi")
	

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
	var successPayload *Buforecaststaffingrequirementsresultresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buforecaststaffingrequirementsresultresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWeekShorttermforecasts invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts
//
// Get short term forecasts
//
// Use \&quot;recent\&quot; (without quotes) for the &#x60;weekDateId&#x60; path parameter to fetch all forecasts for +/- 26 weeks from the current date. Response will include any forecast which spans the specified week
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekShorttermforecasts(businessUnitId string, weekDateId string) (*Bushorttermforecastlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	defaultReturn := new(Bushorttermforecastlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecasts")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecasts")
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
	var successPayload *Bushorttermforecastlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bushorttermforecastlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWorkplanbid invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}
//
// Get a work plan bid
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWorkplanbid(businessUnitId string, bidId string) (*Workplanbid, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	defaultReturn := new(Workplanbid)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWorkplanbid")
	}
	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWorkplanbid")
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
	var successPayload *Workplanbid
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanbid" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWorkplanbidGroup invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups/{bidGroupId}
//
// Get a bid group by bid group Id
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWorkplanbidGroup(businessUnitId string, bidId string, bidGroupId string) (*Workplanbidgroupresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups/{bidGroupId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	path = strings.Replace(path, "{bidGroupId}", url.PathEscape(fmt.Sprintf("%v", bidGroupId)), -1)
	defaultReturn := new(Workplanbidgroupresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWorkplanbidGroup")
	}
	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWorkplanbidGroup")
	}
	// verify the required parameter 'bidGroupId' is set
	if &bidGroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidGroupId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWorkplanbidGroup")
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
	var successPayload *Workplanbidgroupresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanbidgroupresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWorkplanbidGroupPreferences invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups/{bidGroupId}/preferences
//
// Gets the work plan preferences of all the agents in the work plan bid group
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWorkplanbidGroupPreferences(businessUnitId string, bidId string, bidGroupId string) (*Adminagentworkplanpreferenceresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups/{bidGroupId}/preferences"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	path = strings.Replace(path, "{bidGroupId}", url.PathEscape(fmt.Sprintf("%v", bidGroupId)), -1)
	defaultReturn := new(Adminagentworkplanpreferenceresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWorkplanbidGroupPreferences")
	}
	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWorkplanbidGroupPreferences")
	}
	// verify the required parameter 'bidGroupId' is set
	if &bidGroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidGroupId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWorkplanbidGroupPreferences")
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
	var successPayload *Adminagentworkplanpreferenceresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Adminagentworkplanpreferenceresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWorkplanbidGroupsSummary invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups/summary
//
// Get summary of bid groups that belong to a work plan bid
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWorkplanbidGroupsSummary(businessUnitId string, bidId string) (*Workplanbidgroupsummarylist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups/summary"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	defaultReturn := new(Workplanbidgroupsummarylist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWorkplanbidGroupsSummary")
	}
	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWorkplanbidGroupsSummary")
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
	var successPayload *Workplanbidgroupsummarylist
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanbidgroupsummarylist" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitWorkplanbids invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids
//
// Get list of work plan bids
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWorkplanbids(businessUnitId string) (*Workplanbidlistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Workplanbidlistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWorkplanbids")
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
	var successPayload *Workplanbidlistresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanbidlistresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunits invokes GET /api/v2/workforcemanagement/businessunits
//
// Get business units
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunits(feature string, divisionId string) (*Businessunitlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits"
	defaultReturn := new(Businessunitlisting)
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
	
	queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, "")
	
	queryParams["divisionId"] = a.Configuration.APIClient.ParameterToString(divisionId, "")
	

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
	var successPayload *Businessunitlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Businessunitlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementBusinessunitsDivisionviews invokes GET /api/v2/workforcemanagement/businessunits/divisionviews
//
// Get business units across divisions
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitsDivisionviews(divisionId []string) (*Businessunitlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/divisionviews"
	defaultReturn := new(Businessunitlisting)
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
	
	queryParams["divisionId"] = a.Configuration.APIClient.ParameterToString(divisionId, "multi")
	

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
	var successPayload *Businessunitlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Businessunitlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementCalendarDataIcs invokes GET /api/v2/workforcemanagement/calendar/data/ics
//
// Get ics formatted calendar based on shareable link
func (a WorkforceManagementApi) GetWorkforcemanagementCalendarDataIcs(calendarId string) (*string, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/calendar/data/ics"
	defaultReturn := new(string)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'calendarId' is set
	if &calendarId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'calendarId' when calling WorkforceManagementApi->GetWorkforcemanagementCalendarDataIcs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["calendarId"] = a.Configuration.APIClient.ParameterToString(calendarId, "")
	

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
		"text/calendar",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *string
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "string" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementCalendarUrlIcs invokes GET /api/v2/workforcemanagement/calendar/url/ics
//
// Get existing calendar link for the current user
func (a WorkforceManagementApi) GetWorkforcemanagementCalendarUrlIcs() (*Calendarurlresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/calendar/url/ics"
	defaultReturn := new(Calendarurlresponse)
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
	var successPayload *Calendarurlresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Calendarurlresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementHistoricaldataDeletejob invokes GET /api/v2/workforcemanagement/historicaldata/deletejob
//
// Retrieves delete job status for historical data imports of the organization
func (a WorkforceManagementApi) GetWorkforcemanagementHistoricaldataDeletejob() (*Historicalimportdeletejobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/historicaldata/deletejob"
	defaultReturn := new(Historicalimportdeletejobresponse)
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
	var successPayload *Historicalimportdeletejobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Historicalimportdeletejobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementHistoricaldataImportstatus invokes GET /api/v2/workforcemanagement/historicaldata/importstatus
//
// Retrieves status of the historical data imports of the organization
func (a WorkforceManagementApi) GetWorkforcemanagementHistoricaldataImportstatus() (*Historicalimportstatuslisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/historicaldata/importstatus"
	defaultReturn := new(Historicalimportstatuslisting)
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
	var successPayload *Historicalimportstatuslisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Historicalimportstatuslisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementIntegrationsHris invokes GET /api/v2/workforcemanagement/integrations/hris
//
// Get integrations
func (a WorkforceManagementApi) GetWorkforcemanagementIntegrationsHris() (*Wfmintegrationlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/integrations/hris"
	defaultReturn := new(Wfmintegrationlisting)
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
	var successPayload *Wfmintegrationlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmintegrationlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementIntegrationsHrisTimeofftypesJob invokes GET /api/v2/workforcemanagement/integrations/hris/timeofftypes/jobs/{jobId}
//
// Query the results of time off types job
func (a WorkforceManagementApi) GetWorkforcemanagementIntegrationsHrisTimeofftypesJob(jobId string) (*Hristimeofftypesjobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/integrations/hris/timeofftypes/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Hristimeofftypesjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementIntegrationsHrisTimeofftypesJob")
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
	var successPayload *Hristimeofftypesjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Hristimeofftypesjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunit invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}
//
// Get management unit
//
// settings.shortTermForecasting is deprecated and now lives on the business unit
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunit(managementUnitId string, expand []string) (*Managementunit, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Managementunit)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunit")
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
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	

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
	var successPayload *Managementunit
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Managementunit" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitActivitycodes invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/activitycodes
//
// Deprecated: Instead use /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes. Get the list of activity codes
//
// Deprecated: GetWorkforcemanagementManagementunitActivitycodes is deprecated
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitActivitycodes(managementUnitId string) (*Activitycodecontainer, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/activitycodes"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Activitycodecontainer)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitActivitycodes")
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
	var successPayload *Activitycodecontainer
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Activitycodecontainer" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitAdherence invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/adherence
//
// Get a list of user schedule adherence records for the requested management unit
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitAdherence(managementUnitId string, forceDownloadService bool) (*Userscheduleadherencelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/adherence"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Userscheduleadherencelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitAdherence")
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
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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
	var successPayload *Userscheduleadherencelisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Userscheduleadherencelisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitAgent invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}
//
// Get data for agent in the management unit
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitAgent(managementUnitId string, agentId string, excludeCapabilities bool, expand []string) (*Wfmagent, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{agentId}", url.PathEscape(fmt.Sprintf("%v", agentId)), -1)
	defaultReturn := new(Wfmagent)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitAgent")
	}
	// verify the required parameter 'agentId' is set
	if &agentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'agentId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitAgent")
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
	
	queryParams["excludeCapabilities"] = a.Configuration.APIClient.ParameterToString(excludeCapabilities, "")
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	

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
	var successPayload *Wfmagent
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmagent" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitAgentShifttrades invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades
//
// Gets all the shift trades for a given agent
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitAgentShifttrades(managementUnitId string, agentId string) (*Shifttradelistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{agentId}", url.PathEscape(fmt.Sprintf("%v", agentId)), -1)
	defaultReturn := new(Shifttradelistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitAgentShifttrades")
	}
	// verify the required parameter 'agentId' is set
	if &agentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'agentId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitAgentShifttrades")
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
	var successPayload *Shifttradelistresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Shifttradelistresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitShifttradesMatched invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/shifttrades/matched
//
// Gets a summary of all shift trades in the matched state
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitShifttradesMatched(managementUnitId string) (*Shifttradematchessummaryresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/shifttrades/matched"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Shifttradematchessummaryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitShifttradesMatched")
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
	var successPayload *Shifttradematchessummaryresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Shifttradematchessummaryresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitShifttradesUsers invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/shifttrades/users
//
// Gets list of users available for whom you can send direct shift trade requests
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitShifttradesUsers(managementUnitId string) (*Wfmuserentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/shifttrades/users"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Wfmuserentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitShifttradesUsers")
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
	var successPayload *Wfmuserentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmuserentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitTimeofflimit invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}
//
// Gets a time off limit object
//
// Returns properties of time off limit object, but not daily values.
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitTimeofflimit(managementUnitId string, timeOffLimitId string) (*Timeofflimit, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{timeOffLimitId}", url.PathEscape(fmt.Sprintf("%v", timeOffLimitId)), -1)
	defaultReturn := new(Timeofflimit)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitTimeofflimit")
	}
	// verify the required parameter 'timeOffLimitId' is set
	if &timeOffLimitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffLimitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitTimeofflimit")
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
	var successPayload *Timeofflimit
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeofflimit" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitTimeofflimits invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits
//
// Gets a list of time off limit objects under management unit.
//
// Currently only one time off limit object is allowed under management unit, so the list contains either 0 or 1 element.
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitTimeofflimits(managementUnitId string) (*Timeofflimitlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Timeofflimitlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitTimeofflimits")
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
	var successPayload *Timeofflimitlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeofflimitlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitTimeoffplan invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffplans/{timeOffPlanId}
//
// Gets a time off plan
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitTimeoffplan(managementUnitId string, timeOffPlanId string) (*Timeoffplan, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffplans/{timeOffPlanId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{timeOffPlanId}", url.PathEscape(fmt.Sprintf("%v", timeOffPlanId)), -1)
	defaultReturn := new(Timeoffplan)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitTimeoffplan")
	}
	// verify the required parameter 'timeOffPlanId' is set
	if &timeOffPlanId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffPlanId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitTimeoffplan")
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
	var successPayload *Timeoffplan
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffplan" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitTimeoffplans invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffplans
//
// Gets a list of time off plans
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitTimeoffplans(managementUnitId string) (*Timeoffplanlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffplans"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Timeoffplanlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitTimeoffplans")
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
	var successPayload *Timeoffplanlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffplanlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitUserTimeoffrequest invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}
//
// Get a time off request
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitUserTimeoffrequest(managementUnitId string, userId string, timeOffRequestId string) (*Timeoffrequestresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{userId}", url.PathEscape(fmt.Sprintf("%v", userId)), -1)
	path = strings.Replace(path, "{timeOffRequestId}", url.PathEscape(fmt.Sprintf("%v", timeOffRequestId)), -1)
	defaultReturn := new(Timeoffrequestresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUserTimeoffrequest")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUserTimeoffrequest")
	}
	// verify the required parameter 'timeOffRequestId' is set
	if &timeOffRequestId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffRequestId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUserTimeoffrequest")
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
	var successPayload *Timeoffrequestresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffrequestresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimits invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits
//
// Retrieves time off limit, allocated and waitlisted values according to specific time off request
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimits(managementUnitId string, userId string, timeOffRequestId string) (*Querytimeofflimitvaluesresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{userId}", url.PathEscape(fmt.Sprintf("%v", userId)), -1)
	path = strings.Replace(path, "{timeOffRequestId}", url.PathEscape(fmt.Sprintf("%v", timeOffRequestId)), -1)
	defaultReturn := new(Querytimeofflimitvaluesresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimits")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimits")
	}
	// verify the required parameter 'timeOffRequestId' is set
	if &timeOffRequestId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffRequestId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimits")
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
	var successPayload *Querytimeofflimitvaluesresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Querytimeofflimitvaluesresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitUserTimeoffrequests invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests
//
// Get a list of time off requests for a given user
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitUserTimeoffrequests(managementUnitId string, userId string, recentlyReviewed bool) (*Timeoffrequestlist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{userId}", url.PathEscape(fmt.Sprintf("%v", userId)), -1)
	defaultReturn := new(Timeoffrequestlist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUserTimeoffrequests")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUserTimeoffrequests")
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
	
	queryParams["recentlyReviewed"] = a.Configuration.APIClient.ParameterToString(recentlyReviewed, "")
	

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
	var successPayload *Timeoffrequestlist
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffrequestlist" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitUsers invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users
//
// Get users in the management unit
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitUsers(managementUnitId string) (*Wfmuserentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/users"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Wfmuserentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUsers")
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
	var successPayload *Wfmuserentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmuserentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitWeekSchedule invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}
//
// Deprecated.  Use the equivalent business unit resource instead. Get a week schedule
//
// Deprecated: GetWorkforcemanagementManagementunitWeekSchedule is deprecated
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWeekSchedule(managementUnitId string, weekId string, scheduleId string, expand string, forceDownloadService bool) (*Weekscheduleresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	defaultReturn := new(Weekscheduleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekSchedule")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekSchedule")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekSchedule")
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
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "")
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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
	var successPayload *Weekscheduleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Weekscheduleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitWeekSchedules invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules
//
// Deprecated.  Use the equivalent business unit resource instead. Get the list of schedules in a week in management unit
//
// Deprecated: GetWorkforcemanagementManagementunitWeekSchedules is deprecated
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWeekSchedules(managementUnitId string, weekId string, includeOnlyPublished bool, earliestWeekDate string, latestWeekDate string) (*Weekschedulelistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	defaultReturn := new(Weekschedulelistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekSchedules")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekSchedules")
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
	
	queryParams["includeOnlyPublished"] = a.Configuration.APIClient.ParameterToString(includeOnlyPublished, "")
	
	queryParams["earliestWeekDate"] = a.Configuration.APIClient.ParameterToString(earliestWeekDate, "")
	
	queryParams["latestWeekDate"] = a.Configuration.APIClient.ParameterToString(latestWeekDate, "")
	

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
	var successPayload *Weekschedulelistresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Weekschedulelistresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitWeekShifttrades invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades
//
// Gets all the shift trades for a given week
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWeekShifttrades(managementUnitId string, weekDateId time.Time, evaluateMatches bool, forceDownloadService bool) (*Weekshifttradelistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	defaultReturn := new(Weekshifttradelistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekShifttrades")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekShifttrades")
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
	
	queryParams["evaluateMatches"] = a.Configuration.APIClient.ParameterToString(evaluateMatches, "")
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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
	var successPayload *Weekshifttradelistresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Weekshifttradelistresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitWorkplan invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}
//
// Get a work plan
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWorkplan(managementUnitId string, workPlanId string, includeOnly []string) (*Workplan, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{workPlanId}", url.PathEscape(fmt.Sprintf("%v", workPlanId)), -1)
	defaultReturn := new(Workplan)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWorkplan")
	}
	// verify the required parameter 'workPlanId' is set
	if &workPlanId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workPlanId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWorkplan")
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
	
	queryParams["includeOnly"] = a.Configuration.APIClient.ParameterToString(includeOnly, "multi")
	

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
	var successPayload *Workplan
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplan" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitWorkplanrotation invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}
//
// Get a work plan rotation
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWorkplanrotation(managementUnitId string, workPlanRotationId string) (*Workplanrotationresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{workPlanRotationId}", url.PathEscape(fmt.Sprintf("%v", workPlanRotationId)), -1)
	defaultReturn := new(Workplanrotationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWorkplanrotation")
	}
	// verify the required parameter 'workPlanRotationId' is set
	if &workPlanRotationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workPlanRotationId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWorkplanrotation")
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
	var successPayload *Workplanrotationresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanrotationresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitWorkplanrotations invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations
//
// Get work plan rotations
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWorkplanrotations(managementUnitId string, expand []string) (*Workplanrotationlistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Workplanrotationlistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWorkplanrotations")
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
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	

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
	var successPayload *Workplanrotationlistresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanrotationlistresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitWorkplans invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans
//
// Get work plans
//
// \&quot;expand&#x3D;details\&quot; is deprecated
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWorkplans(managementUnitId string, expand []string, exclude []string) (*Workplanlistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Workplanlistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWorkplans")
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
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	
	queryParams["exclude"] = a.Configuration.APIClient.ParameterToString(exclude, "multi")
	

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
	var successPayload *Workplanlistresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanlistresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunits invokes GET /api/v2/workforcemanagement/managementunits
//
// Get management units
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunits(pageSize int, pageNumber int, expand string, feature string, divisionId string) (*Managementunitlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits"
	defaultReturn := new(Managementunitlisting)
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
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, "")
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "")
	
	queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, "")
	
	queryParams["divisionId"] = a.Configuration.APIClient.ParameterToString(divisionId, "")
	

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
	var successPayload *Managementunitlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Managementunitlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementManagementunitsDivisionviews invokes GET /api/v2/workforcemanagement/managementunits/divisionviews
//
// Get management units across divisions
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitsDivisionviews(divisionId []string) (*Managementunitlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/divisionviews"
	defaultReturn := new(Managementunitlisting)
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
	
	queryParams["divisionId"] = a.Configuration.APIClient.ParameterToString(divisionId, "multi")
	

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
	var successPayload *Managementunitlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Managementunitlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementNotifications invokes GET /api/v2/workforcemanagement/notifications
//
// Get a list of notifications for the current user
//
// Notifications are only initially sent if you have the relevant Notify and Edit permissions
func (a WorkforceManagementApi) GetWorkforcemanagementNotifications() (*Notificationsresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/notifications"
	defaultReturn := new(Notificationsresponse)
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
	var successPayload *Notificationsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Notificationsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementSchedulingjob invokes GET /api/v2/workforcemanagement/schedulingjobs/{jobId}
//
// Get status of the scheduling job
func (a WorkforceManagementApi) GetWorkforcemanagementSchedulingjob(jobId string) (*Schedulingstatusresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/schedulingjobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Schedulingstatusresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementSchedulingjob")
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
	var successPayload *Schedulingstatusresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Schedulingstatusresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementShifttrades invokes GET /api/v2/workforcemanagement/shifttrades
//
// Gets all of my shift trades
func (a WorkforceManagementApi) GetWorkforcemanagementShifttrades() (*Shifttradelistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/shifttrades"
	defaultReturn := new(Shifttradelistresponse)
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
	var successPayload *Shifttradelistresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Shifttradelistresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementShrinkageJob invokes GET /api/v2/workforcemanagement/shrinkage/jobs/{jobId}
//
// Request to fetch the status of the historical shrinkage query
func (a WorkforceManagementApi) GetWorkforcemanagementShrinkageJob(jobId string) (*Wfmhistoricalshrinkageresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/shrinkage/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Wfmhistoricalshrinkageresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementShrinkageJob")
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
	var successPayload *Wfmhistoricalshrinkageresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmhistoricalshrinkageresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementTimeoffbalanceJob invokes GET /api/v2/workforcemanagement/timeoffbalance/jobs/{jobId}
//
// Query the results of time off types job
func (a WorkforceManagementApi) GetWorkforcemanagementTimeoffbalanceJob(jobId string) (*Timeoffbalancejobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/timeoffbalance/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Timeoffbalancejobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementTimeoffbalanceJob")
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
	var successPayload *Timeoffbalancejobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffbalancejobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementTimeoffrequest invokes GET /api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}
//
// Get a time off request for the current user
func (a WorkforceManagementApi) GetWorkforcemanagementTimeoffrequest(timeOffRequestId string) (*Timeoffrequestresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}"
	path = strings.Replace(path, "{timeOffRequestId}", url.PathEscape(fmt.Sprintf("%v", timeOffRequestId)), -1)
	defaultReturn := new(Timeoffrequestresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'timeOffRequestId' is set
	if &timeOffRequestId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffRequestId' when calling WorkforceManagementApi->GetWorkforcemanagementTimeoffrequest")
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
	var successPayload *Timeoffrequestresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffrequestresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementTimeoffrequestWaitlistpositions invokes GET /api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}/waitlistpositions
//
// Get the daily waitlist positions of a time off request for the current user
func (a WorkforceManagementApi) GetWorkforcemanagementTimeoffrequestWaitlistpositions(timeOffRequestId string) (*Waitlistpositionlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}/waitlistpositions"
	path = strings.Replace(path, "{timeOffRequestId}", url.PathEscape(fmt.Sprintf("%v", timeOffRequestId)), -1)
	defaultReturn := new(Waitlistpositionlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'timeOffRequestId' is set
	if &timeOffRequestId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffRequestId' when calling WorkforceManagementApi->GetWorkforcemanagementTimeoffrequestWaitlistpositions")
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
	var successPayload *Waitlistpositionlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Waitlistpositionlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementTimeoffrequests invokes GET /api/v2/workforcemanagement/timeoffrequests
//
// Get a list of time off requests for the current user
func (a WorkforceManagementApi) GetWorkforcemanagementTimeoffrequests(recentlyReviewed bool) (*Timeoffrequestlist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/timeoffrequests"
	defaultReturn := new(Timeoffrequestlist)
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
	
	queryParams["recentlyReviewed"] = a.Configuration.APIClient.ParameterToString(recentlyReviewed, "")
	

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
	var successPayload *Timeoffrequestlist
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffrequestlist" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementUserWorkplanbidranks invokes GET /api/v2/workforcemanagement/users/{userId}/workplanbidranks
//
// Get work plan bid ranks for a user
func (a WorkforceManagementApi) GetWorkforcemanagementUserWorkplanbidranks(userId string) (*Workplanbidranks, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/users/{userId}/workplanbidranks"
	path = strings.Replace(path, "{userId}", url.PathEscape(fmt.Sprintf("%v", userId)), -1)
	defaultReturn := new(Workplanbidranks)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'userId' is set
	if &userId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling WorkforceManagementApi->GetWorkforcemanagementUserWorkplanbidranks")
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
	var successPayload *Workplanbidranks
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanbidranks" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementWorkplanbidPreferences invokes GET /api/v2/workforcemanagement/workplanbids/{bidId}/preferences
//
// Gets an agent's work plan bidding preference
func (a WorkforceManagementApi) GetWorkforcemanagementWorkplanbidPreferences(bidId string) (*Agentworkplanbiddingpreferenceresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/workplanbids/{bidId}/preferences"
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	defaultReturn := new(Agentworkplanbiddingpreferenceresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->GetWorkforcemanagementWorkplanbidPreferences")
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
	var successPayload *Agentworkplanbiddingpreferenceresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Agentworkplanbiddingpreferenceresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementWorkplanbidWorkplans invokes GET /api/v2/workforcemanagement/workplanbids/{bidId}/workplans
//
// Gets an agent's work plans for a bid
func (a WorkforceManagementApi) GetWorkforcemanagementWorkplanbidWorkplans(bidId string) (*Agentworkplanlistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/workplanbids/{bidId}/workplans"
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	defaultReturn := new(Agentworkplanlistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->GetWorkforcemanagementWorkplanbidWorkplans")
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
	var successPayload *Agentworkplanlistresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Agentworkplanlistresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetWorkforcemanagementWorkplanbids invokes GET /api/v2/workforcemanagement/workplanbids
//
// Gets the list of work plan bids that belong to an agent
func (a WorkforceManagementApi) GetWorkforcemanagementWorkplanbids() (*Agentworkplanbids, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/workplanbids"
	defaultReturn := new(Agentworkplanbids)
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
	var successPayload *Agentworkplanbids
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Agentworkplanbids" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementAgentAdherenceExplanation invokes PATCH /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}
//
// Update an adherence explanation
func (a WorkforceManagementApi) PatchWorkforcemanagementAgentAdherenceExplanation(agentId string, explanationId string, body Updateadherenceexplanationstatusrequest) (*Adherenceexplanationasyncresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/{explanationId}"
	path = strings.Replace(path, "{agentId}", url.PathEscape(fmt.Sprintf("%v", agentId)), -1)
	path = strings.Replace(path, "{explanationId}", url.PathEscape(fmt.Sprintf("%v", explanationId)), -1)
	defaultReturn := new(Adherenceexplanationasyncresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'agentId' is set
	if &agentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'agentId' when calling WorkforceManagementApi->PatchWorkforcemanagementAgentAdherenceExplanation")
	}
	// verify the required parameter 'explanationId' is set
	if &explanationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'explanationId' when calling WorkforceManagementApi->PatchWorkforcemanagementAgentAdherenceExplanation")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PatchWorkforcemanagementAgentAdherenceExplanation")
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

	var successPayload *Adherenceexplanationasyncresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Adherenceexplanationasyncresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementAlternativeshiftsTrade invokes PATCH /api/v2/workforcemanagement/alternativeshifts/trades/{tradeId}
//
// Update my alternative shifts trade by trade ID
func (a WorkforceManagementApi) PatchWorkforcemanagementAlternativeshiftsTrade(tradeId string, body Agentupdatealternativeshifttraderequest) (*Alternativeshifttraderesponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/alternativeshifts/trades/{tradeId}"
	path = strings.Replace(path, "{tradeId}", url.PathEscape(fmt.Sprintf("%v", tradeId)), -1)
	defaultReturn := new(Alternativeshifttraderesponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'tradeId' is set
	if &tradeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'tradeId' when calling WorkforceManagementApi->PatchWorkforcemanagementAlternativeshiftsTrade")
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

	var successPayload *Alternativeshifttraderesponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshifttraderesponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementAlternativeshiftsTradesStateJobs invokes PATCH /api/v2/workforcemanagement/alternativeshifts/trades/state/jobs
//
// Bulk update alternative shift trade states
func (a WorkforceManagementApi) PatchWorkforcemanagementAlternativeshiftsTradesStateJobs(body Adminbulkupdatealternativeshifttradestaterequest) (*Alternativeshiftasyncresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/alternativeshifts/trades/state/jobs"
	defaultReturn := new(Alternativeshiftasyncresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PatchWorkforcemanagementAlternativeshiftsTradesStateJobs")
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

	var successPayload *Alternativeshiftasyncresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshiftasyncresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementBusinessunit invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}
//
// Update business unit
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunit(businessUnitId string, body Updatebusinessunitrequest) (*Businessunitresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Businessunitresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunit")
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

	var successPayload *Businessunitresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Businessunitresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementBusinessunitActivitycode invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}
//
// Update an activity code
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitActivitycode(businessUnitId string, activityCodeId string, body Updateactivitycoderequest) (*Businessunitactivitycode, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{activityCodeId}", url.PathEscape(fmt.Sprintf("%v", activityCodeId)), -1)
	defaultReturn := new(Businessunitactivitycode)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitActivitycode")
	}
	// verify the required parameter 'activityCodeId' is set
	if &activityCodeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'activityCodeId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitActivitycode")
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

	var successPayload *Businessunitactivitycode
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Businessunitactivitycode" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementBusinessunitActivityplan invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/{activityPlanId}
//
// Update an activity plan
//
// If a job associated with the activity plan is in &#39;Processing&#39; state the activity plan cannot be updated
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitActivityplan(businessUnitId string, activityPlanId string, body Updateactivityplanrequest) (*Activityplanresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/{activityPlanId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{activityPlanId}", url.PathEscape(fmt.Sprintf("%v", activityPlanId)), -1)
	defaultReturn := new(Activityplanresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitActivityplan")
	}
	// verify the required parameter 'activityPlanId' is set
	if &activityPlanId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'activityPlanId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitActivityplan")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitActivityplan")
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

	var successPayload *Activityplanresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Activityplanresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementBusinessunitAlternativeshiftsSettings invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/alternativeshifts/settings
//
// Update alternative shifts settings for a business unit
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitAlternativeshiftsSettings(businessUnitId string, body Updatealternativeshiftbusettingsrequest) (*Alternativeshiftbusettingsresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/alternativeshifts/settings"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Alternativeshiftbusettingsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitAlternativeshiftsSettings")
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

	var successPayload *Alternativeshiftbusettingsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshiftbusettingsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementBusinessunitPlanninggroup invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}
//
// Updates the planning group
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitPlanninggroup(businessUnitId string, planningGroupId string, body Updateplanninggrouprequest) (*Planninggroup, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{planningGroupId}", url.PathEscape(fmt.Sprintf("%v", planningGroupId)), -1)
	defaultReturn := new(Planninggroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitPlanninggroup")
	}
	// verify the required parameter 'planningGroupId' is set
	if &planningGroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'planningGroupId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitPlanninggroup")
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

	var successPayload *Planninggroup
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Planninggroup" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementBusinessunitSchedulingRun invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}
//
// Mark a schedule run as applied
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitSchedulingRun(businessUnitId string, runId string, body Patchbuschedulerunrequest) (*APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{runId}", url.PathEscape(fmt.Sprintf("%v", runId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitSchedulingRun")
	}
	// verify the required parameter 'runId' is set
	if &runId == nil {
		// false
		return nil, errors.New("Missing required parameter 'runId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitSchedulingRun")
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

// PatchWorkforcemanagementBusinessunitServicegoaltemplate invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates/{serviceGoalTemplateId}
//
// Updates a service goal template
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitServicegoaltemplate(businessUnitId string, serviceGoalTemplateId string, body Updateservicegoaltemplate) (*Servicegoaltemplate, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates/{serviceGoalTemplateId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{serviceGoalTemplateId}", url.PathEscape(fmt.Sprintf("%v", serviceGoalTemplateId)), -1)
	defaultReturn := new(Servicegoaltemplate)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitServicegoaltemplate")
	}
	// verify the required parameter 'serviceGoalTemplateId' is set
	if &serviceGoalTemplateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'serviceGoalTemplateId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitServicegoaltemplate")
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

	var successPayload *Servicegoaltemplate
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Servicegoaltemplate" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementBusinessunitStaffinggroup invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups/{staffingGroupId}
//
// Updates a staffing group
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitStaffinggroup(businessUnitId string, staffingGroupId string, body Updatestaffinggrouprequest) (*Staffinggroupresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups/{staffingGroupId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{staffingGroupId}", url.PathEscape(fmt.Sprintf("%v", staffingGroupId)), -1)
	defaultReturn := new(Staffinggroupresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitStaffinggroup")
	}
	// verify the required parameter 'staffingGroupId' is set
	if &staffingGroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'staffingGroupId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitStaffinggroup")
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

	var successPayload *Staffinggroupresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Staffinggroupresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementBusinessunitTimeoffplan invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/timeoffplans/{timeOffPlanId}
//
// Updates a time-off plan
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitTimeoffplan(businessUnitId string, timeOffPlanId string, body Buupdatetimeoffplanrequest) (*Butimeoffplanresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/timeoffplans/{timeOffPlanId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{timeOffPlanId}", url.PathEscape(fmt.Sprintf("%v", timeOffPlanId)), -1)
	defaultReturn := new(Butimeoffplanresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitTimeoffplan")
	}
	// verify the required parameter 'timeOffPlanId' is set
	if &timeOffPlanId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffPlanId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitTimeoffplan")
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

	var successPayload *Butimeoffplanresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Butimeoffplanresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementBusinessunitWorkplanbid invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}
//
// Update work plan bid
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitWorkplanbid(businessUnitId string, bidId string, body Updateworkplanbid) (*Workplanbid, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	defaultReturn := new(Workplanbid)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitWorkplanbid")
	}
	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitWorkplanbid")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitWorkplanbid")
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

	var successPayload *Workplanbid
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanbid" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementBusinessunitWorkplanbidGroup invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups/{bidGroupId}
//
// Update a bid group by bid group Id
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitWorkplanbidGroup(businessUnitId string, bidId string, bidGroupId string, body Workplanbidgroupupdate) (*Workplanbidgroupresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups/{bidGroupId}"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	path = strings.Replace(path, "{bidGroupId}", url.PathEscape(fmt.Sprintf("%v", bidGroupId)), -1)
	defaultReturn := new(Workplanbidgroupresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitWorkplanbidGroup")
	}
	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitWorkplanbidGroup")
	}
	// verify the required parameter 'bidGroupId' is set
	if &bidGroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidGroupId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitWorkplanbidGroup")
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

	var successPayload *Workplanbidgroupresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanbidgroupresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementBusinessunitWorkplanbidGroupPreferences invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups/{bidGroupId}/preferences
//
// Overrides the assigned work plan for the specified agents
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitWorkplanbidGroupPreferences(businessUnitId string, bidId string, bidGroupId string, body Agentsbidassignedworkplanoverriderequest) (*Adminagentworkplanpreferenceresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups/{bidGroupId}/preferences"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	path = strings.Replace(path, "{bidGroupId}", url.PathEscape(fmt.Sprintf("%v", bidGroupId)), -1)
	defaultReturn := new(Adminagentworkplanpreferenceresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitWorkplanbidGroupPreferences")
	}
	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitWorkplanbidGroupPreferences")
	}
	// verify the required parameter 'bidGroupId' is set
	if &bidGroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidGroupId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitWorkplanbidGroupPreferences")
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

	var successPayload *Adminagentworkplanpreferenceresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Adminagentworkplanpreferenceresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementManagementunit invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}
//
// Update the requested management unit
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunit(managementUnitId string, body Updatemanagementunitrequest) (*Managementunit, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Managementunit)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunit")
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

	var successPayload *Managementunit
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Managementunit" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementManagementunitAgents invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents
//
// Update agent configurations
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitAgents(managementUnitId string, body Updatemuagentsrequest) (*APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/agents"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitAgents")
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

// PatchWorkforcemanagementManagementunitAgentsWorkplansBulk invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/workplans/bulk
//
// Updates agent work plan configuration
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitAgentsWorkplansBulk(managementUnitId string, body Updatemuagentworkplansbatchrequest) (*Updatemuagentworkplansbatchresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/workplans/bulk"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Updatemuagentworkplansbatchresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitAgentsWorkplansBulk")
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

	var successPayload *Updatemuagentworkplansbatchresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Updatemuagentworkplansbatchresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementManagementunitTimeofflimit invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}
//
// Updates a time off limit object.
//
// Updates time off limit object properties, but not daily values.
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitTimeofflimit(managementUnitId string, timeOffLimitId string, body Updatetimeofflimitrequest) (*Timeofflimit, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{timeOffLimitId}", url.PathEscape(fmt.Sprintf("%v", timeOffLimitId)), -1)
	defaultReturn := new(Timeofflimit)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitTimeofflimit")
	}
	// verify the required parameter 'timeOffLimitId' is set
	if &timeOffLimitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffLimitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitTimeofflimit")
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

	var successPayload *Timeofflimit
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeofflimit" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementManagementunitTimeoffplan invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffplans/{timeOffPlanId}
//
// Updates a time off plan
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitTimeoffplan(managementUnitId string, timeOffPlanId string, body Updatetimeoffplanrequest) (*Timeoffplan, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffplans/{timeOffPlanId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{timeOffPlanId}", url.PathEscape(fmt.Sprintf("%v", timeOffPlanId)), -1)
	defaultReturn := new(Timeoffplan)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitTimeoffplan")
	}
	// verify the required parameter 'timeOffPlanId' is set
	if &timeOffPlanId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffPlanId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitTimeoffplan")
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

	var successPayload *Timeoffplan
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffplan" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementManagementunitTimeoffrequestUserIntegrationstatus invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests/{timeOffRequestId}/users/{userId}/integrationstatus
//
// Set integration status for a time off request.
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitTimeoffrequestUserIntegrationstatus(managementUnitId string, timeOffRequestId string, userId string, body Settimeoffintegrationstatusrequest) (*Usertimeoffintegrationstatusresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests/{timeOffRequestId}/users/{userId}/integrationstatus"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{timeOffRequestId}", url.PathEscape(fmt.Sprintf("%v", timeOffRequestId)), -1)
	path = strings.Replace(path, "{userId}", url.PathEscape(fmt.Sprintf("%v", userId)), -1)
	defaultReturn := new(Usertimeoffintegrationstatusresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitTimeoffrequestUserIntegrationstatus")
	}
	// verify the required parameter 'timeOffRequestId' is set
	if &timeOffRequestId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffRequestId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitTimeoffrequestUserIntegrationstatus")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitTimeoffrequestUserIntegrationstatus")
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

	var successPayload *Usertimeoffintegrationstatusresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Usertimeoffintegrationstatusresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementManagementunitUserTimeoffrequest invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}
//
// Update a time off request
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitUserTimeoffrequest(managementUnitId string, userId string, timeOffRequestId string, body Admintimeoffrequestpatch) (*Timeoffrequestresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{userId}", url.PathEscape(fmt.Sprintf("%v", userId)), -1)
	path = strings.Replace(path, "{timeOffRequestId}", url.PathEscape(fmt.Sprintf("%v", timeOffRequestId)), -1)
	defaultReturn := new(Timeoffrequestresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitUserTimeoffrequest")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitUserTimeoffrequest")
	}
	// verify the required parameter 'timeOffRequestId' is set
	if &timeOffRequestId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffRequestId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitUserTimeoffrequest")
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

	var successPayload *Timeoffrequestresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffrequestresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementManagementunitWeekShifttrade invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}
//
// Updates a shift trade. This route can only be called by the initiating agent
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitWeekShifttrade(managementUnitId string, weekDateId time.Time, tradeId string, body Patchshifttraderequest) (*Shifttraderesponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	path = strings.Replace(path, "{tradeId}", url.PathEscape(fmt.Sprintf("%v", tradeId)), -1)
	defaultReturn := new(Shifttraderesponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWeekShifttrade")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWeekShifttrade")
	}
	// verify the required parameter 'tradeId' is set
	if &tradeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'tradeId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWeekShifttrade")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWeekShifttrade")
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

	var successPayload *Shifttraderesponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Shifttraderesponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementManagementunitWorkplan invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}
//
// Update a work plan
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitWorkplan(managementUnitId string, workPlanId string, validationMode string, body Workplan) (*Workplan, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{workPlanId}", url.PathEscape(fmt.Sprintf("%v", workPlanId)), -1)
	defaultReturn := new(Workplan)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWorkplan")
	}
	// verify the required parameter 'workPlanId' is set
	if &workPlanId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workPlanId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWorkplan")
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
	
	queryParams["validationMode"] = a.Configuration.APIClient.ParameterToString(validationMode, "")
	

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

	var successPayload *Workplan
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplan" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementManagementunitWorkplanrotation invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}
//
// Update a work plan rotation
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitWorkplanrotation(managementUnitId string, workPlanRotationId string, body Updateworkplanrotationrequest) (*Workplanrotationresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{workPlanRotationId}", url.PathEscape(fmt.Sprintf("%v", workPlanRotationId)), -1)
	defaultReturn := new(Workplanrotationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWorkplanrotation")
	}
	// verify the required parameter 'workPlanRotationId' is set
	if &workPlanRotationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workPlanRotationId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWorkplanrotation")
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

	var successPayload *Workplanrotationresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanrotationresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementTimeoffrequest invokes PATCH /api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}
//
// Update a time off request for the current user
func (a WorkforceManagementApi) PatchWorkforcemanagementTimeoffrequest(timeOffRequestId string, body Agenttimeoffrequestpatch) (*Timeoffrequestresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}"
	path = strings.Replace(path, "{timeOffRequestId}", url.PathEscape(fmt.Sprintf("%v", timeOffRequestId)), -1)
	defaultReturn := new(Timeoffrequestresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'timeOffRequestId' is set
	if &timeOffRequestId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffRequestId' when calling WorkforceManagementApi->PatchWorkforcemanagementTimeoffrequest")
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

	var successPayload *Timeoffrequestresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffrequestresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementUserWorkplanbidranks invokes PATCH /api/v2/workforcemanagement/users/{userId}/workplanbidranks
//
// Update work plan bid ranks for a user
func (a WorkforceManagementApi) PatchWorkforcemanagementUserWorkplanbidranks(userId string, body Workplanbidranks) (*Workplanbidranks, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/users/{userId}/workplanbidranks"
	path = strings.Replace(path, "{userId}", url.PathEscape(fmt.Sprintf("%v", userId)), -1)
	defaultReturn := new(Workplanbidranks)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'userId' is set
	if &userId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling WorkforceManagementApi->PatchWorkforcemanagementUserWorkplanbidranks")
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

	var successPayload *Workplanbidranks
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanbidranks" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementUsersWorkplanbidranksBulk invokes PATCH /api/v2/workforcemanagement/users/workplanbidranks/bulk
//
// Update bulk work plan bid ranks on users. Max 50 users can be updated at a time.
func (a WorkforceManagementApi) PatchWorkforcemanagementUsersWorkplanbidranksBulk(body []Workplanbidranks) (*Entitylisting, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/users/workplanbidranks/bulk"
	defaultReturn := new(Entitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// true
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PatchWorkforcemanagementUsersWorkplanbidranksBulk")
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

	var successPayload *Entitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Entitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchWorkforcemanagementWorkplanbidPreferences invokes PATCH /api/v2/workforcemanagement/workplanbids/{bidId}/preferences
//
// Update an agent's work plan bidding preference
func (a WorkforceManagementApi) PatchWorkforcemanagementWorkplanbidPreferences(bidId string, body Updateagentworkplanbiddingpreference) (*Agentworkplanbiddingpreferenceresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/workplanbids/{bidId}/preferences"
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	defaultReturn := new(Agentworkplanbiddingpreferenceresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->PatchWorkforcemanagementWorkplanbidPreferences")
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

	var successPayload *Agentworkplanbiddingpreferenceresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Agentworkplanbiddingpreferenceresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementAdherenceExplanations invokes POST /api/v2/workforcemanagement/adherence/explanations
//
// Submit an adherence explanation for the current user
func (a WorkforceManagementApi) PostWorkforcemanagementAdherenceExplanations(body Addadherenceexplanationagentrequest) (*Adherenceexplanationasyncresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/adherence/explanations"
	defaultReturn := new(Adherenceexplanationasyncresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementAdherenceExplanations")
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

	var successPayload *Adherenceexplanationasyncresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Adherenceexplanationasyncresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementAdherenceExplanationsQuery invokes POST /api/v2/workforcemanagement/adherence/explanations/query
//
// Query adherence explanations for the current user
func (a WorkforceManagementApi) PostWorkforcemanagementAdherenceExplanationsQuery(body Agentqueryadherenceexplanationsrequest, forceAsync bool, forceDownloadService bool) (*Queryadherenceexplanationsresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/adherence/explanations/query"
	defaultReturn := new(Queryadherenceexplanationsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementAdherenceExplanationsQuery")
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
	
	queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, "")
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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

	var successPayload *Queryadherenceexplanationsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Queryadherenceexplanationsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementAdherenceHistorical invokes POST /api/v2/workforcemanagement/adherence/historical
//
// Deprecated. Use bulk routes instead (/adherence/historical/bulk)
//
// Deprecated: PostWorkforcemanagementAdherenceHistorical is deprecated
func (a WorkforceManagementApi) PostWorkforcemanagementAdherenceHistorical(body Wfmhistoricaladherencequeryforusers) (*Wfmhistoricaladherenceresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/adherence/historical"
	defaultReturn := new(Wfmhistoricaladherenceresponse)
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
	// body params
	postBody = &body

	var successPayload *Wfmhistoricaladherenceresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmhistoricaladherenceresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementAdherenceHistoricalBulk invokes POST /api/v2/workforcemanagement/adherence/historical/bulk
//
// Request a historical adherence report in bulk
func (a WorkforceManagementApi) PostWorkforcemanagementAdherenceHistoricalBulk(body Wfmhistoricaladherencebulkquery) (*Wfmhistoricaladherencebulkresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/adherence/historical/bulk"
	defaultReturn := new(Wfmhistoricaladherencebulkresponse)
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
	// body params
	postBody = &body

	var successPayload *Wfmhistoricaladherencebulkresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmhistoricaladherencebulkresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementAgentAdherenceExplanations invokes POST /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations
//
// Add an adherence explanation for the requested user
func (a WorkforceManagementApi) PostWorkforcemanagementAgentAdherenceExplanations(agentId string, body Addadherenceexplanationadminrequest) (*Adherenceexplanationasyncresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations"
	path = strings.Replace(path, "{agentId}", url.PathEscape(fmt.Sprintf("%v", agentId)), -1)
	defaultReturn := new(Adherenceexplanationasyncresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'agentId' is set
	if &agentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'agentId' when calling WorkforceManagementApi->PostWorkforcemanagementAgentAdherenceExplanations")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementAgentAdherenceExplanations")
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

	var successPayload *Adherenceexplanationasyncresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Adherenceexplanationasyncresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementAgentAdherenceExplanationsQuery invokes POST /api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/query
//
// Query adherence explanations for the given agent across a specified range
func (a WorkforceManagementApi) PostWorkforcemanagementAgentAdherenceExplanationsQuery(agentId string, body Agentqueryadherenceexplanationsrequest, forceAsync bool, forceDownloadService bool) (*Agentqueryadherenceexplanationsresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/agents/{agentId}/adherence/explanations/query"
	path = strings.Replace(path, "{agentId}", url.PathEscape(fmt.Sprintf("%v", agentId)), -1)
	defaultReturn := new(Agentqueryadherenceexplanationsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'agentId' is set
	if &agentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'agentId' when calling WorkforceManagementApi->PostWorkforcemanagementAgentAdherenceExplanationsQuery")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementAgentAdherenceExplanationsQuery")
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
	
	queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, "")
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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

	var successPayload *Agentqueryadherenceexplanationsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Agentqueryadherenceexplanationsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementAgents invokes POST /api/v2/workforcemanagement/agents
//
// Move agents in and out of management unit
func (a WorkforceManagementApi) PostWorkforcemanagementAgents(body Moveagentsrequest) (*Moveagentsresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/agents"
	defaultReturn := new(Moveagentsresponse)
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
	// body params
	postBody = &body

	var successPayload *Moveagentsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Moveagentsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementAgentsIntegrationsHrisQuery invokes POST /api/v2/workforcemanagement/agents/integrations/hris/query
//
// Query integrations for agents
func (a WorkforceManagementApi) PostWorkforcemanagementAgentsIntegrationsHrisQuery(body Queryagentsintegrationsrequest) (*Agentsintegrationslisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/agents/integrations/hris/query"
	defaultReturn := new(Agentsintegrationslisting)
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
	// body params
	postBody = &body

	var successPayload *Agentsintegrationslisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Agentsintegrationslisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementAgentsMePossibleworkshifts invokes POST /api/v2/workforcemanagement/agents/me/possibleworkshifts
//
// Get agent possible work shifts for requested time frame
func (a WorkforceManagementApi) PostWorkforcemanagementAgentsMePossibleworkshifts(body Agentpossibleworkshiftsrequest) (*Agentpossibleworkshiftsresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/agents/me/possibleworkshifts"
	defaultReturn := new(Agentpossibleworkshiftsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementAgentsMePossibleworkshifts")
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

	var successPayload *Agentpossibleworkshiftsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Agentpossibleworkshiftsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementAgentschedulesMine invokes POST /api/v2/workforcemanagement/agentschedules/mine
//
// Get published schedule for the current user
func (a WorkforceManagementApi) PostWorkforcemanagementAgentschedulesMine(body Bugetcurrentagentschedulerequest) (*Bucurrentagentschedulesearchresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/agentschedules/mine"
	defaultReturn := new(Bucurrentagentschedulesearchresponse)
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
	// body params
	postBody = &body

	var successPayload *Bucurrentagentschedulesearchresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bucurrentagentschedulesearchresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementAlternativeshiftsOffersJobs invokes POST /api/v2/workforcemanagement/alternativeshifts/offers/jobs
//
// Request a list of alternative shift offers for a given schedule
func (a WorkforceManagementApi) PostWorkforcemanagementAlternativeshiftsOffersJobs(body Alternativeshiftoffersrequest) (*Alternativeshiftasyncresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/alternativeshifts/offers/jobs"
	defaultReturn := new(Alternativeshiftasyncresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementAlternativeshiftsOffersJobs")
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

	var successPayload *Alternativeshiftasyncresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshiftasyncresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementAlternativeshiftsOffersSearchJobs invokes POST /api/v2/workforcemanagement/alternativeshifts/offers/search/jobs
//
// Request a search of alternative shift offers for a given shift
func (a WorkforceManagementApi) PostWorkforcemanagementAlternativeshiftsOffersSearchJobs(body Alternativeshiftsearchoffersrequest) (*Alternativeshiftasyncresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/alternativeshifts/offers/search/jobs"
	defaultReturn := new(Alternativeshiftasyncresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementAlternativeshiftsOffersSearchJobs")
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

	var successPayload *Alternativeshiftasyncresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshiftasyncresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementAlternativeshiftsTrades invokes POST /api/v2/workforcemanagement/alternativeshifts/trades
//
// Create my alternative shift trade using an existing offer's jobId
func (a WorkforceManagementApi) PostWorkforcemanagementAlternativeshiftsTrades(body Createalternativeshifttraderequest) (*Alternativeshifttraderesponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/alternativeshifts/trades"
	defaultReturn := new(Alternativeshifttraderesponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementAlternativeshiftsTrades")
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

	var successPayload *Alternativeshifttraderesponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Alternativeshifttraderesponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitActivitycodes invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes
//
// Create a new activity code
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitActivitycodes(businessUnitId string, body Createactivitycoderequest) (*Businessunitactivitycode, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Businessunitactivitycode)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitActivitycodes")
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

	var successPayload *Businessunitactivitycode
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Businessunitactivitycode" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitActivityplanRunsJobs invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/{activityPlanId}/runs/jobs
//
// Run an activity plan manually
//
// Triggers a job running the activity plan. The activity plan cannot be updated until the job completes
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitActivityplanRunsJobs(businessUnitId string, activityPlanId string) (*Activityplanjobresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans/{activityPlanId}/runs/jobs"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{activityPlanId}", url.PathEscape(fmt.Sprintf("%v", activityPlanId)), -1)
	defaultReturn := new(Activityplanjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitActivityplanRunsJobs")
	}
	// verify the required parameter 'activityPlanId' is set
	if &activityPlanId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'activityPlanId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitActivityplanRunsJobs")
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
	var successPayload *Activityplanjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Activityplanjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitActivityplans invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans
//
// Create an activity plan
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitActivityplans(businessUnitId string, body Createactivityplanrequest) (*Activityplanresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activityplans"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Activityplanresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitActivityplans")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitActivityplans")
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

	var successPayload *Activityplanresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Activityplanresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitAdherenceExplanationsQuery invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/adherence/explanations/query
//
// Query adherence explanations across an entire business unit for the requested period
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitAdherenceExplanationsQuery(businessUnitId string, body Buqueryadherenceexplanationsrequest, forceAsync bool, forceDownloadService bool) (*Buqueryadherenceexplanationsresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/adherence/explanations/query"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Buqueryadherenceexplanationsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitAdherenceExplanationsQuery")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitAdherenceExplanationsQuery")
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
	
	queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, "")
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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

	var successPayload *Buqueryadherenceexplanationsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buqueryadherenceexplanationsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitAgentschedulesSearch invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/agentschedules/search
//
// Search published schedules
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitAgentschedulesSearch(businessUnitId string, forceAsync bool, forceDownloadService bool, body Busearchagentschedulesrequest) (*Buasyncagentschedulessearchresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/agentschedules/search"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Buasyncagentschedulessearchresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitAgentschedulesSearch")
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
	
	queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, "")
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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

	var successPayload *Buasyncagentschedulessearchresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buasyncagentschedulessearchresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitAlternativeshiftsTradesSearch invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/alternativeshifts/trades/search
//
// List alternative shifts trades for a given management unit or agent
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitAlternativeshiftsTradesSearch(businessUnitId string, body Searchalternativeshifttradesrequest, forceAsync bool) (*Bulistalternativeshifttradesresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/alternativeshifts/trades/search"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Bulistalternativeshifttradesresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitAlternativeshiftsTradesSearch")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitAlternativeshiftsTradesSearch")
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
	
	queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, "")
	

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

	var successPayload *Bulistalternativeshifttradesresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulistalternativeshifttradesresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitIntraday invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/intraday
//
// Get intraday data for the given date for the requested planningGroupIds
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitIntraday(businessUnitId string, forceAsync bool, body Intradayplanninggrouprequest) (*Asyncintradayresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/intraday"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Asyncintradayresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitIntraday")
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
	
	queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, "")
	

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

	var successPayload *Asyncintradayresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Asyncintradayresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitPlanninggroups invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups
//
// Adds a new planning group
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitPlanninggroups(businessUnitId string, body Createplanninggrouprequest) (*Planninggroup, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Planninggroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitPlanninggroups")
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

	var successPayload *Planninggroup
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Planninggroup" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitServicegoaltemplates invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates
//
// Adds a new service goal template
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitServicegoaltemplates(businessUnitId string, body Createservicegoaltemplate) (*Servicegoaltemplate, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Servicegoaltemplate)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitServicegoaltemplates")
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

	var successPayload *Servicegoaltemplate
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Servicegoaltemplate" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitStaffinggroups invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups
//
// Creates a new staffing group
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitStaffinggroups(businessUnitId string, body Createstaffinggrouprequest) (*Staffinggroupresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Staffinggroupresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitStaffinggroups")
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

	var successPayload *Staffinggroupresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Staffinggroupresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitStaffinggroupsQuery invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups/query
//
// Gets staffing group associations for a list of user IDs
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitStaffinggroupsQuery(businessUnitId string, body Queryuserstaffinggrouplistrequest) (*Userstaffinggrouplisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/staffinggroups/query"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Userstaffinggrouplisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitStaffinggroupsQuery")
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

	var successPayload *Userstaffinggrouplisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Userstaffinggrouplisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitTimeofflimits invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/timeofflimits
//
// Creates a new time-off limit object
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitTimeofflimits(businessUnitId string, body Bucreatetimeofflimitrequest) (*Butimeofflimitresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/timeofflimits"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Butimeofflimitresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitTimeofflimits")
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

	var successPayload *Butimeofflimitresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Butimeofflimitresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitTimeofflimitsValuesQuery invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/timeofflimits/values/query
//
// Retrieves time-off limit related values based on a given set of filters.
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitTimeofflimitsValuesQuery(businessUnitId string, body Querytimeofflimitvaluesrequest) (*Butimeofflimitvaluesresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/timeofflimits/values/query"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Butimeofflimitvaluesresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitTimeofflimitsValuesQuery")
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

	var successPayload *Butimeofflimitvaluesresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Butimeofflimitvaluesresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitTimeoffplans invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/timeoffplans
//
// Creates a new time-off plan
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitTimeoffplans(businessUnitId string, body Bucreatetimeoffplanrequest) (*Butimeoffplanresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/timeoffplans"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Butimeoffplanresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitTimeoffplans")
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

	var successPayload *Butimeoffplanresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Butimeoffplanresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQuery invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query
//
// Loads agent schedule data from the schedule. Used in combination with the metadata route
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQuery(businessUnitId string, weekId time.Time, scheduleId string, body Buqueryagentschedulesrequest, forceAsync bool, forceDownloadService bool) (*Buasyncagentschedulesqueryresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	defaultReturn := new(Buasyncagentschedulesqueryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQuery")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQuery")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQuery")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQuery")
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
	
	queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, "")
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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

	var successPayload *Buasyncagentschedulesqueryresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buasyncagentschedulesqueryresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekScheduleCopy invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy
//
// Copy a schedule
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekScheduleCopy(businessUnitId string, weekId time.Time, scheduleId string, body Bucopyschedulerequest) (*Buasyncscheduleresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	defaultReturn := new(Buasyncscheduleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleCopy")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleCopy")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleCopy")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleCopy")
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

	var successPayload *Buasyncscheduleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buasyncscheduleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculations invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/performancepredictions/recalculations
//
// Request a daily recalculation of the performance prediction for the associated schedule
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculations(businessUnitId string, weekId string, scheduleId string, body Wfmprocessuploadrequest) (*Performancepredictionrecalculationresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/performancepredictions/recalculations"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	defaultReturn := new(Performancepredictionrecalculationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculations")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculations")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculations")
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

	var successPayload *Performancepredictionrecalculationresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Performancepredictionrecalculationresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculationsUploadurl invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/performancepredictions/recalculations/uploadurl
//
// Upload daily activity changes to be able to request a performance prediction recalculation
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculationsUploadurl(businessUnitId string, weekId string, scheduleId string, body Uploadurlrequestbody) (*Performancepredictionrecalculationuploadresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/performancepredictions/recalculations/uploadurl"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	defaultReturn := new(Performancepredictionrecalculationuploadresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculationsUploadurl")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculationsUploadurl")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulePerformancepredictionsRecalculationsUploadurl")
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

	var successPayload *Performancepredictionrecalculationuploadresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Performancepredictionrecalculationuploadresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekScheduleReschedule invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule
//
// Start a rescheduling run
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekScheduleReschedule(businessUnitId string, weekId time.Time, scheduleId string, body Bureschedulerequest) (*Buasyncschedulerunresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	defaultReturn := new(Buasyncschedulerunresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleReschedule")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleReschedule")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleReschedule")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleReschedule")
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

	var successPayload *Buasyncschedulerunresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buasyncschedulerunresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekScheduleUpdate invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update
//
// Starts processing a schedule update
//
// Call after uploading the schedule data to the url supplied by the /update/uploadurl route
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekScheduleUpdate(businessUnitId string, weekId time.Time, scheduleId string, body Processscheduleupdateuploadrequest) (*Buasyncscheduleresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	defaultReturn := new(Buasyncscheduleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleUpdate")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleUpdate")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleUpdate")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleUpdate")
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

	var successPayload *Buasyncscheduleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buasyncscheduleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurl invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl
//
// Creates a signed upload URL for updating a schedule
//
// Once the upload is complete, call the /{scheduleId}/update route to start the schedule update process
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurl(businessUnitId string, weekId time.Time, scheduleId string, body Uploadurlrequestbody) (*Updatescheduleuploadresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/update/uploadurl"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	path = strings.Replace(path, "{scheduleId}", url.PathEscape(fmt.Sprintf("%v", scheduleId)), -1)
	defaultReturn := new(Updatescheduleuploadresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurl")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurl")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurl")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleUpdateUploadurl")
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

	var successPayload *Updatescheduleuploadresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Updatescheduleuploadresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekSchedules invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules
//
// Create a blank schedule
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekSchedules(businessUnitId string, weekId time.Time, body Bucreateblankschedulerequest) (*Buschedulemetadata, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	defaultReturn := new(Buschedulemetadata)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedules")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedules")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedules")
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

	var successPayload *Buschedulemetadata
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buschedulemetadata" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekSchedulesGenerate invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate
//
// Generate a schedule
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekSchedulesGenerate(businessUnitId string, weekId time.Time, body Bugenerateschedulerequest) (*Buasyncschedulerunresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	defaultReturn := new(Buasyncschedulerunresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulesGenerate")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulesGenerate")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulesGenerate")
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

	var successPayload *Buasyncschedulerunresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buasyncschedulerunresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekSchedulesImport invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import
//
// Starts processing a schedule import
//
// Call after uploading the schedule data to the url supplied by the /import/uploadurl route
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekSchedulesImport(businessUnitId string, weekId time.Time, body Wfmprocessuploadrequest) (*Scheduleuploadprocessingresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	defaultReturn := new(Scheduleuploadprocessingresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulesImport")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulesImport")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulesImport")
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

	var successPayload *Scheduleuploadprocessingresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Scheduleuploadprocessingresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekSchedulesImportUploadurl invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import/uploadurl
//
// Creates a signed upload URL for importing a schedule
//
// Once the upload is complete, call the /import route to start the schedule import process
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekSchedulesImportUploadurl(businessUnitId string, weekId time.Time, body Uploadurlrequestbody) (*Importscheduleuploadresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/import/uploadurl"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekId}", url.PathEscape(fmt.Sprintf("%v", weekId)), -1)
	defaultReturn := new(Importscheduleuploadresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulesImportUploadurl")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulesImportUploadurl")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulesImportUploadurl")
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

	var successPayload *Importscheduleuploadresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Importscheduleuploadresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekShorttermforecastCopy invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/copy
//
// Copy a short term forecast
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekShorttermforecastCopy(businessUnitId string, weekDateId time.Time, forecastId string, body Copybuforecastrequest, forceAsync bool) (*Asyncforecastoperationresult, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/copy"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	path = strings.Replace(path, "{forecastId}", url.PathEscape(fmt.Sprintf("%v", forecastId)), -1)
	defaultReturn := new(Asyncforecastoperationresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastCopy")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastCopy")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'forecastId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastCopy")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastCopy")
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
	
	queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, "")
	

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

	var successPayload *Asyncforecastoperationresult
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Asyncforecastoperationresult" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerate invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate
//
// Generate a short term forecast
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerate(businessUnitId string, weekDateId time.Time, body Generatebuforecastrequest, forceAsync bool) (*Asyncforecastoperationresult, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	defaultReturn := new(Asyncforecastoperationresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerate")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerate")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerate")
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
	
	queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, "")
	

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

	var successPayload *Asyncforecastoperationresult
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Asyncforecastoperationresult" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekShorttermforecastsImport invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/import
//
// Starts importing the uploaded short term forecast
//
// Call after uploading the forecast data to the url supplied by the /import/uploadurl route
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekShorttermforecastsImport(businessUnitId string, weekDateId time.Time, body Wfmprocessuploadrequest) (*Importforecastresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/import"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	defaultReturn := new(Importforecastresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastsImport")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastsImport")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastsImport")
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

	var successPayload *Importforecastresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Importforecastresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWeekShorttermforecastsImportUploadurl invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/import/uploadurl
//
// Creates a signed upload URL for importing a short term forecast
//
// Once the upload is complete, call the /import route to start the short term forecast import process
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekShorttermforecastsImportUploadurl(businessUnitId string, weekDateId time.Time, body Uploadurlrequestbody) (*Importforecastuploadresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/import/uploadurl"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	defaultReturn := new(Importforecastuploadresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastsImportUploadurl")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastsImportUploadurl")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastsImportUploadurl")
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

	var successPayload *Importforecastuploadresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Importforecastuploadresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWorkplanbidCopy invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/copy
//
// Copy a work plan bid
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWorkplanbidCopy(businessUnitId string, bidId string, body Copyworkplanbid) (*Workplanbid, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/copy"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	defaultReturn := new(Workplanbid)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWorkplanbidCopy")
	}
	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWorkplanbidCopy")
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

	var successPayload *Workplanbid
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanbid" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWorkplanbidGroups invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups
//
// Add a bid group in a given work plan bid
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWorkplanbidGroups(businessUnitId string, bidId string, body Workplanbidgroupcreate) (*Workplanbidgroupresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids/{bidId}/groups"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{bidId}", url.PathEscape(fmt.Sprintf("%v", bidId)), -1)
	defaultReturn := new(Workplanbidgroupresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWorkplanbidGroups")
	}
	// verify the required parameter 'bidId' is set
	if &bidId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'bidId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWorkplanbidGroups")
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

	var successPayload *Workplanbidgroupresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanbidgroupresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunitWorkplanbids invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids
//
// Create a new work plan bid
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWorkplanbids(businessUnitId string, body Createworkplanbid) (*Workplanbid, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/workplanbids"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	defaultReturn := new(Workplanbid)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWorkplanbids")
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

	var successPayload *Workplanbid
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanbid" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementBusinessunits invokes POST /api/v2/workforcemanagement/businessunits
//
// Add a new business unit
//
// It may take a minute or two for a new business unit to be available for api operations
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunits(body Createbusinessunitrequest) (*Businessunitresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits"
	defaultReturn := new(Businessunitresponse)
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
	// body params
	postBody = &body

	var successPayload *Businessunitresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Businessunitresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementCalendarUrlIcs invokes POST /api/v2/workforcemanagement/calendar/url/ics
//
// Create a newly generated calendar link for the current user; if the current user has previously generated one, the generated link will be returned
func (a WorkforceManagementApi) PostWorkforcemanagementCalendarUrlIcs(language string) (*Calendarurlresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/calendar/url/ics"
	defaultReturn := new(Calendarurlresponse)
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
	
	queryParams["language"] = a.Configuration.APIClient.ParameterToString(language, "")
	

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
	var successPayload *Calendarurlresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Calendarurlresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementHistoricaldataDeletejob invokes POST /api/v2/workforcemanagement/historicaldata/deletejob
//
// Delete the entries of the historical data imports in the organization
func (a WorkforceManagementApi) PostWorkforcemanagementHistoricaldataDeletejob() (*Historicalimportdeletejobresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/historicaldata/deletejob"
	defaultReturn := new(Historicalimportdeletejobresponse)
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
	var successPayload *Historicalimportdeletejobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Historicalimportdeletejobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementHistoricaldataValidate invokes POST /api/v2/workforcemanagement/historicaldata/validate
//
// Trigger validation process for historical import
func (a WorkforceManagementApi) PostWorkforcemanagementHistoricaldataValidate(body Validationservicerequest) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/historicaldata/validate"
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
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

// PostWorkforcemanagementIntegrationsHriTimeofftypesJobs invokes POST /api/v2/workforcemanagement/integrations/hris/{hrisIntegrationId}/timeofftypes/jobs
//
// Get list of time off types configured in integration
func (a WorkforceManagementApi) PostWorkforcemanagementIntegrationsHriTimeofftypesJobs(hrisIntegrationId string) (*Hristimeofftypesresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/integrations/hris/{hrisIntegrationId}/timeofftypes/jobs"
	path = strings.Replace(path, "{hrisIntegrationId}", url.PathEscape(fmt.Sprintf("%v", hrisIntegrationId)), -1)
	defaultReturn := new(Hristimeofftypesresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'hrisIntegrationId' is set
	if &hrisIntegrationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'hrisIntegrationId' when calling WorkforceManagementApi->PostWorkforcemanagementIntegrationsHriTimeofftypesJobs")
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
	var successPayload *Hristimeofftypesresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Hristimeofftypesresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitAgentsWorkplansQuery invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/workplans/query
//
// Get agents work plans configuration
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitAgentsWorkplansQuery(managementUnitId string, forceDownloadService bool, body Getagentsworkplansrequest) (*Agentsworkplansresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/workplans/query"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Agentsworkplansresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitAgentsWorkplansQuery")
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
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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

	var successPayload *Agentsworkplansresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Agentsworkplansresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitAgentschedulesSearch invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/agentschedules/search
//
// Query published schedules for given given time range for set of users
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitAgentschedulesSearch(managementUnitId string, forceAsync bool, forceDownloadService bool, body Busearchagentschedulesrequest) (*Buasyncagentschedulessearchresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/agentschedules/search"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Buasyncagentschedulessearchresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitAgentschedulesSearch")
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
	
	queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, "")
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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

	var successPayload *Buasyncagentschedulessearchresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Buasyncagentschedulessearchresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitHistoricaladherencequery invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/historicaladherencequery
//
// Request a historical adherence report
//
// The maximum supported range for historical adherence queries is 31 days, or 7 days with includeExceptions &#x3D; true
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitHistoricaladherencequery(managementUnitId string, body Wfmhistoricaladherencequery) (*Wfmhistoricaladherenceresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/historicaladherencequery"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Wfmhistoricaladherenceresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitHistoricaladherencequery")
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

	var successPayload *Wfmhistoricaladherenceresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmhistoricaladherenceresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitMove invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/move
//
// Move the requested management unit to a new business unit
//
// Returns status 200 if the management unit is already in the requested business unit
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitMove(managementUnitId string, body Movemanagementunitrequest) (*Movemanagementunitresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/move"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Movemanagementunitresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitMove")
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

	var successPayload *Movemanagementunitresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Movemanagementunitresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitSchedulesSearch invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/schedules/search
//
// Query published schedules for given given time range for set of users
//
// Deprecated: PostWorkforcemanagementManagementunitSchedulesSearch is deprecated
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitSchedulesSearch(managementUnitId string, body Userlistschedulerequestbody) (*Userschedulecontainer, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/schedules/search"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Userschedulecontainer)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitSchedulesSearch")
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

	var successPayload *Userschedulecontainer
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Userschedulecontainer" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitShrinkageJobs invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/shrinkage/jobs
//
// Request a historical shrinkage report
//
// The maximum supported range for historical shrinkage queries is up to 32 days. Historical Shrinkage for a given date range can be queried in two modes - granular and aggregated. To see granular shrinkage information, provide granularity in the request body. 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitShrinkageJobs(managementUnitId string, body Wfmhistoricalshrinkagerequest) (*Wfmhistoricalshrinkageresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/shrinkage/jobs"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Wfmhistoricalshrinkageresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitShrinkageJobs")
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

	var successPayload *Wfmhistoricalshrinkageresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmhistoricalshrinkageresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitTimeofflimits invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits
//
// Creates a new time off limit object under management unit.
//
// Only one limit object is allowed under management unit, so an attempt to create second object will fail.
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitTimeofflimits(managementUnitId string, body Createtimeofflimitrequest) (*Timeofflimit, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Timeofflimit)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitTimeofflimits")
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

	var successPayload *Timeofflimit
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeofflimit" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitTimeofflimitsValuesQuery invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/values/query
//
// Retrieves time off limit related values based on a given set of filters.
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitTimeofflimitsValuesQuery(managementUnitId string, body Querytimeofflimitvaluesrequest) (*Querytimeofflimitvaluesresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/values/query"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Querytimeofflimitvaluesresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitTimeofflimitsValuesQuery")
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

	var successPayload *Querytimeofflimitvaluesresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Querytimeofflimitvaluesresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitTimeoffplans invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffplans
//
// Creates a new time off plan
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitTimeoffplans(managementUnitId string, body Createtimeoffplanrequest) (*Timeoffplan, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffplans"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Timeoffplan)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitTimeoffplans")
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

	var successPayload *Timeoffplan
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffplan" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitTimeoffrequests invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests
//
// Create a new time off request
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitTimeoffrequests(managementUnitId string, body Createadmintimeoffrequest) (*Timeoffrequestlist, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Timeoffrequestlist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitTimeoffrequests")
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

	var successPayload *Timeoffrequestlist
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffrequestlist" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitTimeoffrequestsIntegrationstatusQuery invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests/integrationstatus/query
//
// Retrieves integration statuses for a list of time off requests
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitTimeoffrequestsIntegrationstatusQuery(managementUnitId string, body Querytimeoffintegrationstatusrequest) (*Usertimeoffintegrationstatusresponselisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests/integrationstatus/query"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Usertimeoffintegrationstatusresponselisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitTimeoffrequestsIntegrationstatusQuery")
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

	var successPayload *Usertimeoffintegrationstatusresponselisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Usertimeoffintegrationstatusresponselisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitTimeoffrequestsQuery invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests/query
//
// Fetches time off requests matching the conditions specified in the request body
//
// Request body requires one of the following: User ID is specified, statuses &#x3D;&#x3D; [Pending] or date range to be specified and less than or equal to 33 days.  All other fields are filters
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitTimeoffrequestsQuery(managementUnitId string, forceDownloadService bool, body Timeoffrequestquerybody) (*Timeoffrequestlisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests/query"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Timeoffrequestlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitTimeoffrequestsQuery")
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
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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

	var successPayload *Timeoffrequestlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffrequestlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitTimeoffrequestsWaitlistpositionsQuery invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests/waitlistpositions/query
//
// Retrieves daily waitlist position for a list of time off requests
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitTimeoffrequestsWaitlistpositionsQuery(managementUnitId string, body Querywaitlistpositionsrequest) (*Waitlistpositionlisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests/waitlistpositions/query"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Waitlistpositionlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitTimeoffrequestsWaitlistpositionsQuery")
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

	var successPayload *Waitlistpositionlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Waitlistpositionlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitUserTimeoffbalanceJobs invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffbalance/jobs
//
// Query time off balances for a given user for specified activity code and dates
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitUserTimeoffbalanceJobs(managementUnitId string, userId string, body Timeoffbalancerequest) (*Timeoffbalancesresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffbalance/jobs"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{userId}", url.PathEscape(fmt.Sprintf("%v", userId)), -1)
	defaultReturn := new(Timeoffbalancesresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitUserTimeoffbalanceJobs")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitUserTimeoffbalanceJobs")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitUserTimeoffbalanceJobs")
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

	var successPayload *Timeoffbalancesresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffbalancesresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitUserTimeoffrequestTimeoffbalanceJobs invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeoffbalance/jobs
//
// Query time off balances for dates spanned by a given time off request
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitUserTimeoffrequestTimeoffbalanceJobs(managementUnitId string, userId string, timeOffRequestId string) (*Timeoffbalancesresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeoffbalance/jobs"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{userId}", url.PathEscape(fmt.Sprintf("%v", userId)), -1)
	path = strings.Replace(path, "{timeOffRequestId}", url.PathEscape(fmt.Sprintf("%v", timeOffRequestId)), -1)
	defaultReturn := new(Timeoffbalancesresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitUserTimeoffrequestTimeoffbalanceJobs")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitUserTimeoffrequestTimeoffbalanceJobs")
	}
	// verify the required parameter 'timeOffRequestId' is set
	if &timeOffRequestId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffRequestId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitUserTimeoffrequestTimeoffbalanceJobs")
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
	var successPayload *Timeoffbalancesresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffbalancesresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitUserTimeoffrequestsEstimate invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/estimate
//
// Estimates available time off for an agent
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitUserTimeoffrequestsEstimate(managementUnitId string, userId string, body Estimateavailabletimeoffrequest) (*Estimateavailabletimeoffresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/estimate"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{userId}", url.PathEscape(fmt.Sprintf("%v", userId)), -1)
	defaultReturn := new(Estimateavailabletimeoffresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitUserTimeoffrequestsEstimate")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitUserTimeoffrequestsEstimate")
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

	var successPayload *Estimateavailabletimeoffresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Estimateavailabletimeoffresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitWeekShifttradeMatch invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}/match
//
// Matches a shift trade. This route can only be called by the receiving agent
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekShifttradeMatch(managementUnitId string, weekDateId time.Time, tradeId string, body Matchshifttraderequest) (*Matchshifttraderesponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}/match"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	path = strings.Replace(path, "{tradeId}", url.PathEscape(fmt.Sprintf("%v", tradeId)), -1)
	defaultReturn := new(Matchshifttraderesponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradeMatch")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradeMatch")
	}
	// verify the required parameter 'tradeId' is set
	if &tradeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'tradeId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradeMatch")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradeMatch")
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

	var successPayload *Matchshifttraderesponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Matchshifttraderesponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitWeekShifttrades invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades
//
// Adds a shift trade
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekShifttrades(managementUnitId string, weekDateId time.Time, body Addshifttraderequest) (*Shifttraderesponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	defaultReturn := new(Shifttraderesponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttrades")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttrades")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttrades")
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

	var successPayload *Shifttraderesponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Shifttraderesponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitWeekShifttradesSearch invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/search
//
// Searches for potential shift trade matches for the current agent
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekShifttradesSearch(managementUnitId string, weekDateId time.Time, body Searchshifttradesrequest, forceDownloadService bool) (*Searchshifttradesresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/search"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	defaultReturn := new(Searchshifttradesresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradesSearch")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradesSearch")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradesSearch")
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
	
	queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, "")
	

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

	var successPayload *Searchshifttradesresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Searchshifttradesresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitWeekShifttradesStateBulk invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/state/bulk
//
// Updates the state of a batch of shift trades
//
// Admin functionality is not supported with \&quot;mine\&quot;.
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekShifttradesStateBulk(managementUnitId string, weekDateId time.Time, body Bulkshifttradestateupdaterequest, forceAsync bool) (*Bulkupdateshifttradestateresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/state/bulk"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{weekDateId}", url.PathEscape(fmt.Sprintf("%v", weekDateId)), -1)
	defaultReturn := new(Bulkupdateshifttradestateresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradesStateBulk")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradesStateBulk")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradesStateBulk")
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
	
	queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, "")
	

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

	var successPayload *Bulkupdateshifttradestateresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulkupdateshifttradestateresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitWorkplanCopy invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/copy
//
// Create a copy of work plan
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWorkplanCopy(managementUnitId string, workPlanId string, body Copyworkplan) (*Workplan, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/copy"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{workPlanId}", url.PathEscape(fmt.Sprintf("%v", workPlanId)), -1)
	defaultReturn := new(Workplan)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWorkplanCopy")
	}
	// verify the required parameter 'workPlanId' is set
	if &workPlanId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workPlanId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWorkplanCopy")
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

	var successPayload *Workplan
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplan" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitWorkplanValidate invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate
//
// Validate Work Plan
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWorkplanValidate(managementUnitId string, workPlanId string, expand []string, body Workplanvalidationrequest) (*Validateworkplanresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{workPlanId}", url.PathEscape(fmt.Sprintf("%v", workPlanId)), -1)
	defaultReturn := new(Validateworkplanresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWorkplanValidate")
	}
	// verify the required parameter 'workPlanId' is set
	if &workPlanId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workPlanId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWorkplanValidate")
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
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	

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

	var successPayload *Validateworkplanresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Validateworkplanresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitWorkplanrotationCopy invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy
//
// Create a copy of work plan rotation
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWorkplanrotationCopy(managementUnitId string, workPlanRotationId string, body Copyworkplanrotationrequest) (*Workplanrotationresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{workPlanRotationId}", url.PathEscape(fmt.Sprintf("%v", workPlanRotationId)), -1)
	defaultReturn := new(Workplanrotationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWorkplanrotationCopy")
	}
	// verify the required parameter 'workPlanRotationId' is set
	if &workPlanRotationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'workPlanRotationId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWorkplanrotationCopy")
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

	var successPayload *Workplanrotationresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanrotationresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitWorkplanrotations invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations
//
// Create a new work plan rotation
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWorkplanrotations(managementUnitId string, body Addworkplanrotationrequest) (*Workplanrotationresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Workplanrotationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWorkplanrotations")
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

	var successPayload *Workplanrotationresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplanrotationresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunitWorkplans invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans
//
// Create a new work plan
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWorkplans(managementUnitId string, validationMode string, body Createworkplan) (*Workplan, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	defaultReturn := new(Workplan)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWorkplans")
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
	
	queryParams["validationMode"] = a.Configuration.APIClient.ParameterToString(validationMode, "")
	

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

	var successPayload *Workplan
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Workplan" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementManagementunits invokes POST /api/v2/workforcemanagement/managementunits
//
// Add a management unit
//
// It may take a minute or two for a new management unit to be available for api operations
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunits(body Createmanagementunitapirequest) (*Managementunit, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits"
	defaultReturn := new(Managementunit)
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
	// body params
	postBody = &body

	var successPayload *Managementunit
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Managementunit" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementNotificationsUpdate invokes POST /api/v2/workforcemanagement/notifications/update
//
// Mark a list of notifications as read or unread
func (a WorkforceManagementApi) PostWorkforcemanagementNotificationsUpdate(body Updatenotificationsrequest) (*Updatenotificationsresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/notifications/update"
	defaultReturn := new(Updatenotificationsresponse)
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
	// body params
	postBody = &body

	var successPayload *Updatenotificationsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Updatenotificationsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementSchedules invokes POST /api/v2/workforcemanagement/schedules
//
// Get published schedule for the current user
//
// Deprecated: PostWorkforcemanagementSchedules is deprecated
func (a WorkforceManagementApi) PostWorkforcemanagementSchedules(body Currentuserschedulerequestbody) (*Userschedulecontainer, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/schedules"
	defaultReturn := new(Userschedulecontainer)
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
	// body params
	postBody = &body

	var successPayload *Userschedulecontainer
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Userschedulecontainer" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementTeamAdherenceHistorical invokes POST /api/v2/workforcemanagement/teams/{teamId}/adherence/historical
//
// Request a teams historical adherence report
//
// The maximum supported range for historical adherence queries is 31 days, or 7 days with includeExceptions &#x3D; true
func (a WorkforceManagementApi) PostWorkforcemanagementTeamAdherenceHistorical(teamId string, body Wfmhistoricaladherencequeryforteams) (*Wfmhistoricaladherenceresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/teams/{teamId}/adherence/historical"
	path = strings.Replace(path, "{teamId}", url.PathEscape(fmt.Sprintf("%v", teamId)), -1)
	defaultReturn := new(Wfmhistoricaladherenceresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'teamId' is set
	if &teamId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'teamId' when calling WorkforceManagementApi->PostWorkforcemanagementTeamAdherenceHistorical")
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

	var successPayload *Wfmhistoricaladherenceresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmhistoricaladherenceresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementTeamShrinkageJobs invokes POST /api/v2/workforcemanagement/teams/{teamId}/shrinkage/jobs
//
// Request a historical shrinkage report
//
// The maximum supported range for historical shrinkage queries is up to 32 days
func (a WorkforceManagementApi) PostWorkforcemanagementTeamShrinkageJobs(teamId string, body Wfmhistoricalshrinkageteamsrequest) (*Wfmhistoricalshrinkageresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/teams/{teamId}/shrinkage/jobs"
	path = strings.Replace(path, "{teamId}", url.PathEscape(fmt.Sprintf("%v", teamId)), -1)
	defaultReturn := new(Wfmhistoricalshrinkageresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'teamId' is set
	if &teamId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'teamId' when calling WorkforceManagementApi->PostWorkforcemanagementTeamShrinkageJobs")
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

	var successPayload *Wfmhistoricalshrinkageresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Wfmhistoricalshrinkageresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementTimeoffbalanceJobs invokes POST /api/v2/workforcemanagement/timeoffbalance/jobs
//
// Query time off balances for the current user for specified activity code and dates
func (a WorkforceManagementApi) PostWorkforcemanagementTimeoffbalanceJobs(body Timeoffbalancerequest) (*Timeoffbalancesresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/timeoffbalance/jobs"
	defaultReturn := new(Timeoffbalancesresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementTimeoffbalanceJobs")
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

	var successPayload *Timeoffbalancesresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffbalancesresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementTimeofflimitsAvailableQuery invokes POST /api/v2/workforcemanagement/timeofflimits/available/query
//
// Queries available time off for the current user
func (a WorkforceManagementApi) PostWorkforcemanagementTimeofflimitsAvailableQuery(body Availabletimeoffrequest) (*Availabletimeoffresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/timeofflimits/available/query"
	defaultReturn := new(Availabletimeoffresponse)
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
	// body params
	postBody = &body

	var successPayload *Availabletimeoffresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Availabletimeoffresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementTimeoffrequests invokes POST /api/v2/workforcemanagement/timeoffrequests
//
// Create a time off request for the current user
func (a WorkforceManagementApi) PostWorkforcemanagementTimeoffrequests(body Createagenttimeoffrequest) (*Timeoffrequestresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/timeoffrequests"
	defaultReturn := new(Timeoffrequestresponse)
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
	// body params
	postBody = &body

	var successPayload *Timeoffrequestresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffrequestresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementTimeoffrequestsEstimate invokes POST /api/v2/workforcemanagement/timeoffrequests/estimate
//
// Estimates available time off for current user
func (a WorkforceManagementApi) PostWorkforcemanagementTimeoffrequestsEstimate(body Estimateavailabletimeoffrequest) (*Estimateavailabletimeoffresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/timeoffrequests/estimate"
	defaultReturn := new(Estimateavailabletimeoffresponse)
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
	// body params
	postBody = &body

	var successPayload *Estimateavailabletimeoffresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Estimateavailabletimeoffresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostWorkforcemanagementTimeoffrequestsIntegrationstatusQuery invokes POST /api/v2/workforcemanagement/timeoffrequests/integrationstatus/query
//
// Retrieves integration statuses for a list of current user time off requests
func (a WorkforceManagementApi) PostWorkforcemanagementTimeoffrequestsIntegrationstatusQuery(body Currentusertimeoffintegrationstatusrequest) (*Timeoffintegrationstatusresponselisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/timeoffrequests/integrationstatus/query"
	defaultReturn := new(Timeoffintegrationstatusresponselisting)
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
	// body params
	postBody = &body

	var successPayload *Timeoffintegrationstatusresponselisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeoffintegrationstatusresponselisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutWorkforcemanagementAgentIntegrationsHris invokes PUT /api/v2/workforcemanagement/agents/{agentId}/integrations/hris
//
// Update integrations for agent
func (a WorkforceManagementApi) PutWorkforcemanagementAgentIntegrationsHris(agentId string, body Agentintegrationsrequest) (*Agentintegrationsresponse, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/agents/{agentId}/integrations/hris"
	path = strings.Replace(path, "{agentId}", url.PathEscape(fmt.Sprintf("%v", agentId)), -1)
	defaultReturn := new(Agentintegrationsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'agentId' is set
	if &agentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'agentId' when calling WorkforceManagementApi->PutWorkforcemanagementAgentIntegrationsHris")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PutWorkforcemanagementAgentIntegrationsHris")
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

	var successPayload *Agentintegrationsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Agentintegrationsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutWorkforcemanagementBusinessunitTimeofflimitValues invokes PUT /api/v2/workforcemanagement/businessunits/{businessUnitId}/timeofflimits/{timeOffLimitId}/values
//
// Sets daily values for a date range of time-off limit object
//
// Note that only limit daily values can be set through API, allocated and waitlisted values are read-only for time-off limit API
func (a WorkforceManagementApi) PutWorkforcemanagementBusinessunitTimeofflimitValues(businessUnitId string, timeOffLimitId string, body Busettimeofflimitvaluesrequest) (*Butimeofflimitresponse, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/timeofflimits/{timeOffLimitId}/values"
	path = strings.Replace(path, "{businessUnitId}", url.PathEscape(fmt.Sprintf("%v", businessUnitId)), -1)
	path = strings.Replace(path, "{timeOffLimitId}", url.PathEscape(fmt.Sprintf("%v", timeOffLimitId)), -1)
	defaultReturn := new(Butimeofflimitresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PutWorkforcemanagementBusinessunitTimeofflimitValues")
	}
	// verify the required parameter 'timeOffLimitId' is set
	if &timeOffLimitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffLimitId' when calling WorkforceManagementApi->PutWorkforcemanagementBusinessunitTimeofflimitValues")
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

	var successPayload *Butimeofflimitresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Butimeofflimitresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutWorkforcemanagementManagementunitTimeofflimitValues invokes PUT /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}/values
//
// Sets daily values for a date range of time off limit object
//
// Note that only limit daily values can be set through API, allocated and waitlisted values are read-only for time off limit API
func (a WorkforceManagementApi) PutWorkforcemanagementManagementunitTimeofflimitValues(managementUnitId string, timeOffLimitId string, body Settimeofflimitvaluesrequest) (*Timeofflimit, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits/{timeOffLimitId}/values"
	path = strings.Replace(path, "{managementUnitId}", url.PathEscape(fmt.Sprintf("%v", managementUnitId)), -1)
	path = strings.Replace(path, "{timeOffLimitId}", url.PathEscape(fmt.Sprintf("%v", timeOffLimitId)), -1)
	defaultReturn := new(Timeofflimit)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PutWorkforcemanagementManagementunitTimeofflimitValues")
	}
	// verify the required parameter 'timeOffLimitId' is set
	if &timeOffLimitId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'timeOffLimitId' when calling WorkforceManagementApi->PutWorkforcemanagementManagementunitTimeofflimitValues")
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

	var successPayload *Timeofflimit
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timeofflimit" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

