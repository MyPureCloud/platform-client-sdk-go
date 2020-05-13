package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"time"
"encoding/json"
)

// WorkforceManagementApi provides functions for API endpoints
type WorkforceManagementApi struct {
	Configuration *Configuration
}

// NewWorkforceManagementApi creates an API instance using the default configuration
func NewWorkforceManagementApi() *WorkforceManagementApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating WorkforceManagementApi with base path: %s", strings.ToLower(config.BasePath)))
	return &WorkforceManagementApi{
		Configuration: config,
	}
}

// NewWorkforceManagementApiWithConfig creates an API instance using the provided configuration
func NewWorkforceManagementApiWithConfig(config *Configuration) *WorkforceManagementApi {
	config.Debugf("Creating WorkforceManagementApi with base path: %s\n", strings.ToLower(config.BasePath))
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
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteWorkforcemanagementBusinessunitActivitycode invokes DELETE /api/v2/workforcemanagement/businessunits/{buId}/activitycodes/{acId}
//
// Deletes an activity code
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitActivitycode(buId string, acId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{buId}/activitycodes/{acId}"
	path = strings.Replace(path, "{buId}", fmt.Sprintf("%v", buId), -1)
	path = strings.Replace(path, "{acId}", fmt.Sprintf("%v", acId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'buId' is set
	if &buId == nil {
		// 
		return nil, errors.New("Missing required parameter 'buId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitActivitycode")
	}
	// verify the required parameter 'acId' is set
	if &acId == nil {
		// 
		return nil, errors.New("Missing required parameter 'acId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitActivitycode")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteWorkforcemanagementBusinessunitPlanninggroup invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}
//
// Deletes the planning group
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitPlanninggroup(businessUnitId string, planningGroupId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{planningGroupId}", fmt.Sprintf("%v", planningGroupId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitPlanninggroup")
	}
	// verify the required parameter 'planningGroupId' is set
	if &planningGroupId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteWorkforcemanagementBusinessunitSchedulingRun invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}
//
// Cancel a scheduling run
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitSchedulingRun(businessUnitId string, runId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{runId}", fmt.Sprintf("%v", runId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitSchedulingRun")
	}
	// verify the required parameter 'runId' is set
	if &runId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteWorkforcemanagementBusinessunitServicegoaltemplate invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates/{serviceGoalTemplateId}
//
// Delete a service goal template
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitServicegoaltemplate(businessUnitId string, serviceGoalTemplateId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates/{serviceGoalTemplateId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{serviceGoalTemplateId}", fmt.Sprintf("%v", serviceGoalTemplateId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitServicegoaltemplate")
	}
	// verify the required parameter 'serviceGoalTemplateId' is set
	if &serviceGoalTemplateId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteWorkforcemanagementBusinessunitWeekSchedule invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}
//
// Delete a schedule
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitWeekSchedule(businessUnitId string, weekId time.Time, scheduleId string) (*Buasyncscheduleresponse, *APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	defaultReturn := new(Buasyncscheduleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWeekSchedule")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWeekSchedule")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteWorkforcemanagementBusinessunitWeekShorttermforecast invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}
//
// Delete a short term forecast
//
// Must not be tied to any schedules
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitWeekShorttermforecast(businessUnitId string, weekDateId time.Time, forecastId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	path = strings.Replace(path, "{forecastId}", fmt.Sprintf("%v", forecastId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWeekShorttermforecast")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitWeekShorttermforecast")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteWorkforcemanagementManagementunit invokes DELETE /api/v2/workforcemanagement/managementunits/{muId}
//
// Delete management unit
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunit(muId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunit")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteWorkforcemanagementManagementunitActivitycode invokes DELETE /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}
//
// Deletes an activity code
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunitActivitycode(muId string, acId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	path = strings.Replace(path, "{acId}", fmt.Sprintf("%v", acId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitActivitycode")
	}
	// verify the required parameter 'acId' is set
	if &acId == nil {
		// 
		return nil, errors.New("Missing required parameter 'acId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitActivitycode")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteWorkforcemanagementManagementunitSchedulingRun invokes DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}
//
// Cancel a schedule run
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunitSchedulingRun(managementUnitId string, runId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{runId}", fmt.Sprintf("%v", runId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitSchedulingRun")
	}
	// verify the required parameter 'runId' is set
	if &runId == nil {
		// 
		return nil, errors.New("Missing required parameter 'runId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitSchedulingRun")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteWorkforcemanagementManagementunitServicegoalgroup invokes DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/servicegoalgroups/{serviceGoalGroupId}
//
// Delete a service goal group
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunitServicegoalgroup(managementUnitId string, serviceGoalGroupId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/servicegoalgroups/{serviceGoalGroupId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{serviceGoalGroupId}", fmt.Sprintf("%v", serviceGoalGroupId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitServicegoalgroup")
	}
	// verify the required parameter 'serviceGoalGroupId' is set
	if &serviceGoalGroupId == nil {
		// 
		return nil, errors.New("Missing required parameter 'serviceGoalGroupId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitServicegoalgroup")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteWorkforcemanagementManagementunitWeekSchedule invokes DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}
//
// Delete a schedule
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunitWeekSchedule(managementUnitId string, weekId string, scheduleId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitWeekSchedule")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitWeekSchedule")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
		return nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitWeekSchedule")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteWorkforcemanagementManagementunitWeekShorttermforecast invokes DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}
//
// Delete a short term forecast
//
// Must not be tied to any schedules
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunitWeekShorttermforecast(managementUnitId string, weekDateId string, forecastId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	path = strings.Replace(path, "{forecastId}", fmt.Sprintf("%v", forecastId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitWeekShorttermforecast")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitWeekShorttermforecast")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// 
		return nil, errors.New("Missing required parameter 'forecastId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitWeekShorttermforecast")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteWorkforcemanagementManagementunitWorkplan invokes DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}
//
// Delete a work plan
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunitWorkplan(managementUnitId string, workPlanId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{workPlanId}", fmt.Sprintf("%v", workPlanId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitWorkplan")
	}
	// verify the required parameter 'workPlanId' is set
	if &workPlanId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementAdherence invokes GET /api/v2/workforcemanagement/adherence
//
// Get a list of UserScheduleAdherence records for the requested users
//
// 
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
	
	var collectionFormat string
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range userId {
			queryParams["userId"] = value
		}
	} else {
		queryParams["userId"] = a.Configuration.APIClient.ParameterToString(userId, collectionFormat)
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
	var successPayload []Userscheduleadherence
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

// GetWorkforcemanagementAdhocmodelingjob invokes GET /api/v2/workforcemanagement/adhocmodelingjobs/{jobId}
//
// Get status of the modeling job
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementAdhocmodelingjob(jobId string) (*Modelingstatusresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/adhocmodelingjobs/{jobId}"
	path = strings.Replace(path, "{jobId}", fmt.Sprintf("%v", jobId), -1)
	defaultReturn := new(Modelingstatusresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling WorkforceManagementApi->GetWorkforcemanagementAdhocmodelingjob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Modelingstatusresponse
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

// GetWorkforcemanagementBusinessunit invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}
//
// Get business unit
//
// Expanding \&quot;settings\&quot; will retrieve all settings.  All other expands will retrieve only the requested settings field(s).
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunit(businessUnitId string, expand []string) (*Businessunit, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	defaultReturn := new(Businessunit)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
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
	var successPayload *Businessunit
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

// GetWorkforcemanagementBusinessunitActivitycode invokes GET /api/v2/workforcemanagement/businessunits/{buId}/activitycodes/{acId}
//
// Get an activity code
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitActivitycode(buId string, acId string) (*Businessunitactivitycode, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{buId}/activitycodes/{acId}"
	path = strings.Replace(path, "{buId}", fmt.Sprintf("%v", buId), -1)
	path = strings.Replace(path, "{acId}", fmt.Sprintf("%v", acId), -1)
	defaultReturn := new(Businessunitactivitycode)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'buId' is set
	if &buId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'buId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivitycode")
	}
	// verify the required parameter 'acId' is set
	if &acId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'acId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivitycode")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Businessunitactivitycode
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

// GetWorkforcemanagementBusinessunitActivitycodes invokes GET /api/v2/workforcemanagement/businessunits/{buId}/activitycodes
//
// Get activity codes
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitActivitycodes(buId string) (*Businessunitactivitycodelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{buId}/activitycodes"
	path = strings.Replace(path, "{buId}", fmt.Sprintf("%v", buId), -1)
	defaultReturn := new(Businessunitactivitycodelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'buId' is set
	if &buId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'buId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivitycodes")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Businessunitactivitycodelisting
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

// GetWorkforcemanagementBusinessunitIntradayPlanninggroups invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/intraday/planninggroups
//
// Get intraday planning groups for the given date
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitIntradayPlanninggroups(businessUnitId string, date time.Time) (*Wfmintradayplanninggrouplisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/intraday/planninggroups"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	defaultReturn := new(Wfmintradayplanninggrouplisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitIntradayPlanninggroups")
	}
	// verify the required parameter 'date' is set
	if &date == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["date"] = a.Configuration.APIClient.ParameterToString(date, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitManagementunits invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits
//
// Get all authorized management units in the business unit
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitManagementunits(businessUnitId string, feature string, divisionId string) (*Managementunitlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	defaultReturn := new(Managementunitlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["divisionId"] = a.Configuration.APIClient.ParameterToString(divisionId, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitPlanninggroup invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}
//
// Get a planning group
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitPlanninggroup(businessUnitId string, planningGroupId string) (*Planninggroup, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{planningGroupId}", fmt.Sprintf("%v", planningGroupId), -1)
	defaultReturn := new(Planninggroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitPlanninggroup")
	}
	// verify the required parameter 'planningGroupId' is set
	if &planningGroupId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitPlanninggroups invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups
//
// Gets list of planning groups
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitPlanninggroups(businessUnitId string) (*Planninggrouplist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	defaultReturn := new(Planninggrouplist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitSchedulingRun invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}
//
// Get a scheduling run
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitSchedulingRun(businessUnitId string, runId string) (*Buschedulerun, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{runId}", fmt.Sprintf("%v", runId), -1)
	defaultReturn := new(Buschedulerun)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitSchedulingRun")
	}
	// verify the required parameter 'runId' is set
	if &runId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitSchedulingRunResult invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}/result
//
// Get the result of a rescheduling operation
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitSchedulingRunResult(businessUnitId string, runId string, managementUnitIds []string, expand []string) (*Burescheduleresult, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}/result"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{runId}", fmt.Sprintf("%v", runId), -1)
	defaultReturn := new(Burescheduleresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitSchedulingRunResult")
	}
	// verify the required parameter 'runId' is set
	if &runId == nil {
		// 
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
	
	var collectionFormat string
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range managementUnitIds {
			queryParams["managementUnitIds"] = value
		}
	} else {
		queryParams["managementUnitIds"] = a.Configuration.APIClient.ParameterToString(managementUnitIds, collectionFormat)
	}
	
	
	
	
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
	var successPayload *Burescheduleresult
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

// GetWorkforcemanagementBusinessunitSchedulingRuns invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs
//
// Get the list of scheduling runs
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitSchedulingRuns(businessUnitId string) (*Buschedulerunlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	defaultReturn := new(Buschedulerunlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitServicegoaltemplate invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates/{serviceGoalTemplateId}
//
// Get a service goal template
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitServicegoaltemplate(businessUnitId string, serviceGoalTemplateId string) (*Servicegoaltemplate, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates/{serviceGoalTemplateId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{serviceGoalTemplateId}", fmt.Sprintf("%v", serviceGoalTemplateId), -1)
	defaultReturn := new(Servicegoaltemplate)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitServicegoaltemplate")
	}
	// verify the required parameter 'serviceGoalTemplateId' is set
	if &serviceGoalTemplateId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitServicegoaltemplates invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates
//
// Gets list of service goal templates
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitServicegoaltemplates(businessUnitId string) (*Servicegoaltemplatelist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	defaultReturn := new(Servicegoaltemplatelist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitWeekSchedule invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}
//
// Get the metadata for the schedule, describing which management units and agents are in the scheduleSchedule data can then be loaded with the query route
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekSchedule(businessUnitId string, weekId time.Time, scheduleId string) (*Buschedulemetadata, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	defaultReturn := new(Buschedulemetadata)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedule")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedule")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitWeekScheduleGenerationresults invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/generationresults
//
// Get the generation results for a generated schedule
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekScheduleGenerationresults(businessUnitId string, weekId time.Time, scheduleId string) (*Schedulegenerationresult, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/generationresults"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	defaultReturn := new(Schedulegenerationresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleGenerationresults")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleGenerationresults")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitWeekScheduleHeadcountforecast invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/headcountforecast
//
// Get the headcount forecast by planning group for the schedule
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekScheduleHeadcountforecast(businessUnitId string, weekId time.Time, scheduleId string, forceDownload bool) (*Buheadcountforecastresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/headcountforecast"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	defaultReturn := new(Buheadcountforecastresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleHeadcountforecast")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleHeadcountforecast")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["forceDownload"] = a.Configuration.APIClient.ParameterToString(forceDownload, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitWeekSchedules invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules
//
// Get the list of week schedules for the specified week
//
// Use \&quot;recent\&quot; for the `weekId` path parameter to fetch all forecasts for +/- 26 weeks from the current date. Response will include any schedule which spans the specified week
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekSchedules(businessUnitId string, weekId string, includeOnlyPublished bool, expand string) (*Buschedulelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	defaultReturn := new(Buschedulelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekSchedules")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["includeOnlyPublished"] = a.Configuration.APIClient.ParameterToString(includeOnlyPublished, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitWeekShorttermforecast invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}
//
// Get a short term forecast
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekShorttermforecast(businessUnitId string, weekDateId time.Time, forecastId string, expand []string) (*Bushorttermforecast, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	path = strings.Replace(path, "{forecastId}", fmt.Sprintf("%v", forecastId), -1)
	defaultReturn := new(Bushorttermforecast)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecast")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecast")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// 
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
	var successPayload *Bushorttermforecast
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

// GetWorkforcemanagementBusinessunitWeekShorttermforecastData invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/data
//
// Get the result of a short term forecast calculation
//
// Includes modifications unless you pass the doNotApplyModifications query parameter
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekShorttermforecastData(businessUnitId string, weekDateId time.Time, forecastId string, weekNumber int32, forceDownloadService bool) (*Buforecastresultresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/data"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	path = strings.Replace(path, "{forecastId}", fmt.Sprintf("%v", forecastId), -1)
	defaultReturn := new(Buforecastresultresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastData")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastData")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["weekNumber"] = a.Configuration.APIClient.ParameterToString(weekNumber, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresults invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/generationresults
//
// Gets the forecast generation results
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresults(businessUnitId string, weekDateId time.Time, forecastId string) (*Buforecastgenerationresult, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/generationresults"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	path = strings.Replace(path, "{forecastId}", fmt.Sprintf("%v", forecastId), -1)
	defaultReturn := new(Buforecastgenerationresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresults")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastGenerationresults")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroups invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups
//
// Gets the forecast planning group snapshot
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroups(businessUnitId string, weekDateId time.Time, forecastId string) (*Forecastplanninggroupsresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	path = strings.Replace(path, "{forecastId}", fmt.Sprintf("%v", forecastId), -1)
	defaultReturn := new(Forecastplanninggroupsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroups")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroups")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitWeekShorttermforecasts invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts
//
// Get short term forecasts
//
// Use \&quot;recent\&quot; for the `weekDateId` path parameter to fetch all forecasts for +/- 26 weeks from the current date. Response will include any forecast which spans the specified week
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekShorttermforecasts(businessUnitId string, weekDateId string) (*Bushorttermforecastlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	defaultReturn := new(Bushorttermforecastlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekShorttermforecasts")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunits invokes GET /api/v2/workforcemanagement/businessunits
//
// Get business units
//
// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["divisionId"] = a.Configuration.APIClient.ParameterToString(divisionId, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitsDivisionviews invokes GET /api/v2/workforcemanagement/businessunits/divisionviews
//
// Get business units across divisions
//
// 
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
	
	var collectionFormat string
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range divisionId {
			queryParams["divisionId"] = value
		}
	} else {
		queryParams["divisionId"] = a.Configuration.APIClient.ParameterToString(divisionId, collectionFormat)
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
	var successPayload *Businessunitlisting
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

// GetWorkforcemanagementManagementunit invokes GET /api/v2/workforcemanagement/managementunits/{muId}
//
// Get management unit
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunit(muId string, expand []string) (*Managementunit, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Managementunit)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunit")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Managementunit
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

// GetWorkforcemanagementManagementunitActivitycode invokes GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}
//
// Get an activity code
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitActivitycode(muId string, acId string) (*Activitycode, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	path = strings.Replace(path, "{acId}", fmt.Sprintf("%v", acId), -1)
	defaultReturn := new(Activitycode)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitActivitycode")
	}
	// verify the required parameter 'acId' is set
	if &acId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'acId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitActivitycode")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Activitycode
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

// GetWorkforcemanagementManagementunitActivitycodes invokes GET /api/v2/workforcemanagement/managementunits/{muId}/activitycodes
//
// Get activity codes
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitActivitycodes(muId string) (*Activitycodecontainer, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/activitycodes"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Activitycodecontainer)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitActivitycodes")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Activitycodecontainer
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

// GetWorkforcemanagementManagementunitAdherence invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/adherence
//
// Get a list of user schedule adherence records for the requested management unit
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitAdherence(managementUnitId string, forceDownloadService bool) (*Userscheduleadherencelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/adherence"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Userscheduleadherencelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementManagementunitAgent invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}
//
// Get data for agent in the management unit
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitAgent(managementUnitId string, agentId string) (*Wfmagent, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{agentId}", fmt.Sprintf("%v", agentId), -1)
	defaultReturn := new(Wfmagent)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitAgent")
	}
	// verify the required parameter 'agentId' is set
	if &agentId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementManagementunitAgentShifttrades invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades
//
// Gets all the shift trades for a given agent
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitAgentShifttrades(managementUnitId string, agentId string) (*Shifttradelistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{agentId}", fmt.Sprintf("%v", agentId), -1)
	defaultReturn := new(Shifttradelistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitAgentShifttrades")
	}
	// verify the required parameter 'agentId' is set
	if &agentId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementManagementunitIntradayQueues invokes GET /api/v2/workforcemanagement/managementunits/{muId}/intraday/queues
//
// Get intraday queues for the given date
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitIntradayQueues(muId string, date string) (*Wfmintradayqueuelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/intraday/queues"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Wfmintradayqueuelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitIntradayQueues")
	}
	// verify the required parameter 'date' is set
	if &date == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'date' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitIntradayQueues")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
		queryParams["date"] = a.Configuration.APIClient.ParameterToString(date, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Wfmintradayqueuelisting
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

// GetWorkforcemanagementManagementunitSchedulingRun invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}
//
// Gets the status for a specific scheduling run
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitSchedulingRun(managementUnitId string, runId string) (*Schedulingrunresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{runId}", fmt.Sprintf("%v", runId), -1)
	defaultReturn := new(Schedulingrunresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitSchedulingRun")
	}
	// verify the required parameter 'runId' is set
	if &runId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'runId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitSchedulingRun")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Schedulingrunresponse
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

// GetWorkforcemanagementManagementunitSchedulingRunResult invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}/result
//
// Gets the result of a specific scheduling run
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitSchedulingRunResult(managementUnitId string, runId string) (*Rescheduleresult, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}/result"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{runId}", fmt.Sprintf("%v", runId), -1)
	defaultReturn := new(Rescheduleresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitSchedulingRunResult")
	}
	// verify the required parameter 'runId' is set
	if &runId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'runId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitSchedulingRunResult")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Rescheduleresult
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

// GetWorkforcemanagementManagementunitSchedulingRuns invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs
//
// Get the status of all the ongoing schedule runs
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitSchedulingRuns(managementUnitId string) (*Schedulingrunlistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Schedulingrunlistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitSchedulingRuns")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Schedulingrunlistresponse
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

// GetWorkforcemanagementManagementunitServicegoalgroup invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/servicegoalgroups/{serviceGoalGroupId}
//
// Get a service goal group
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitServicegoalgroup(managementUnitId string, serviceGoalGroupId string) (*Servicegoalgroup, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/servicegoalgroups/{serviceGoalGroupId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{serviceGoalGroupId}", fmt.Sprintf("%v", serviceGoalGroupId), -1)
	defaultReturn := new(Servicegoalgroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitServicegoalgroup")
	}
	// verify the required parameter 'serviceGoalGroupId' is set
	if &serviceGoalGroupId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'serviceGoalGroupId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitServicegoalgroup")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Servicegoalgroup
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

// GetWorkforcemanagementManagementunitServicegoalgroups invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/servicegoalgroups
//
// Get service goal groups
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitServicegoalgroups(managementUnitId string) (*Servicegoalgrouplist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/servicegoalgroups"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Servicegoalgrouplist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitServicegoalgroups")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Servicegoalgrouplist
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

// GetWorkforcemanagementManagementunitSettings invokes GET /api/v2/workforcemanagement/managementunits/{muId}/settings
//
// Get the settings for the requested management unit. Deprecated, use the GET management unit route instead
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitSettings(muId string) (*Managementunitsettingsresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/settings"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Managementunitsettingsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitSettings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Managementunitsettingsresponse
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

// GetWorkforcemanagementManagementunitShifttradesMatched invokes GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/matched
//
// Gets a summary of all shift trades in the matched state
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitShifttradesMatched(muId string) (*Shifttradematchessummaryresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/shifttrades/matched"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Shifttradematchessummaryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitShifttradesMatched")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Shifttradematchessummaryresponse
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

// GetWorkforcemanagementManagementunitShifttradesUsers invokes GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/users
//
// Gets list of users available for whom you can send direct shift trade requests
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitShifttradesUsers(muId string) (*Wfmuserentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/shifttrades/users"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Wfmuserentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitShifttradesUsers")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Wfmuserentitylisting
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

// GetWorkforcemanagementManagementunitUserTimeoffrequest invokes GET /api/v2/workforcemanagement/managementunits/{muId}/users/{userId}/timeoffrequests/{timeOffRequestId}
//
// Get a time off request
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitUserTimeoffrequest(muId string, userId string, timeOffRequestId string) (*Timeoffrequestresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/users/{userId}/timeoffrequests/{timeOffRequestId}"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)
	path = strings.Replace(path, "{timeOffRequestId}", fmt.Sprintf("%v", timeOffRequestId), -1)
	defaultReturn := new(Timeoffrequestresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUserTimeoffrequest")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUserTimeoffrequest")
	}
	// verify the required parameter 'timeOffRequestId' is set
	if &timeOffRequestId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementManagementunitUserTimeoffrequests invokes GET /api/v2/workforcemanagement/managementunits/{muId}/users/{userId}/timeoffrequests
//
// Get a list of time off requests for a given user
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitUserTimeoffrequests(muId string, userId string, recentlyReviewed bool) (*Timeoffrequestlist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/users/{userId}/timeoffrequests"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)
	defaultReturn := new(Timeoffrequestlist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUserTimeoffrequests")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["recentlyReviewed"] = a.Configuration.APIClient.ParameterToString(recentlyReviewed, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementManagementunitUsers invokes GET /api/v2/workforcemanagement/managementunits/{muId}/users
//
// Get users in the management unit
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitUsers(muId string) (*Wfmuserentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/users"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Wfmuserentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUsers")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Wfmuserentitylisting
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

// GetWorkforcemanagementManagementunitWeekSchedule invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}
//
// Get a week schedule
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWeekSchedule(managementUnitId string, weekId string, scheduleId string, expand string, forceDownloadService bool) (*Weekscheduleresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	defaultReturn := new(Weekscheduleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekSchedule")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekSchedule")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementManagementunitWeekScheduleGenerationresults invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/generationresults
//
// Get week schedule generation results
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWeekScheduleGenerationresults(managementUnitId string, weekId string, scheduleId string) (*Weekschedulegenerationresult, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/generationresults"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	defaultReturn := new(Weekschedulegenerationresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekScheduleGenerationresults")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekScheduleGenerationresults")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekScheduleGenerationresults")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Weekschedulegenerationresult
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

// GetWorkforcemanagementManagementunitWeekSchedules invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules
//
// Get the list of schedules in a week in management unit
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWeekSchedules(managementUnitId string, weekId string, includeOnlyPublished bool, earliestWeekDate string, latestWeekDate string) (*Weekschedulelistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	defaultReturn := new(Weekschedulelistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekSchedules")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["includeOnlyPublished"] = a.Configuration.APIClient.ParameterToString(includeOnlyPublished, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["earliestWeekDate"] = a.Configuration.APIClient.ParameterToString(earliestWeekDate, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["latestWeekDate"] = a.Configuration.APIClient.ParameterToString(latestWeekDate, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementManagementunitWeekShifttrades invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades
//
// Gets all the shift trades for a given week
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWeekShifttrades(managementUnitId string, weekDateId time.Time, evaluateMatches bool) (*Weekshifttradelistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	defaultReturn := new(Weekshifttradelistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekShifttrades")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["evaluateMatches"] = a.Configuration.APIClient.ParameterToString(evaluateMatches, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementManagementunitWeekShorttermforecastFinal invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/final
//
// Get the final result of a short term forecast calculation with modifications applied
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWeekShorttermforecastFinal(managementUnitId string, weekDateId string, forecastId string, forceDownloadService bool) (*Forecastresultresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/final"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	path = strings.Replace(path, "{forecastId}", fmt.Sprintf("%v", forecastId), -1)
	defaultReturn := new(Forecastresultresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekShorttermforecastFinal")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekShorttermforecastFinal")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'forecastId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekShorttermforecastFinal")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Forecastresultresponse
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

// GetWorkforcemanagementManagementunitWeekShorttermforecasts invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts
//
// Get short term forecasts
//
// Use \&quot;recent\&quot; for the `weekDateId` path parameter to fetch all forecasts for +/- 26 weeks from the current date
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWeekShorttermforecasts(managementUnitId string, weekDateId string) (*Shorttermforecastlistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	defaultReturn := new(Shorttermforecastlistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekShorttermforecasts")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWeekShorttermforecasts")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Shorttermforecastlistresponse
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

// GetWorkforcemanagementManagementunitWorkplan invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}
//
// Get a work plan
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWorkplan(managementUnitId string, workPlanId string) (*Workplan, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{workPlanId}", fmt.Sprintf("%v", workPlanId), -1)
	defaultReturn := new(Workplan)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWorkplan")
	}
	// verify the required parameter 'workPlanId' is set
	if &workPlanId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementManagementunitWorkplans invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans
//
// Get work plans
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWorkplans(managementUnitId string, expand []string) (*Workplanlistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Workplanlistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	var successPayload *Workplanlistresponse
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

// GetWorkforcemanagementManagementunits invokes GET /api/v2/workforcemanagement/managementunits
//
// Get management units
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunits(pageSize int32, pageNumber int32, expand string, feature string, divisionId string) (*Managementunitlisting, *APIResponse, error) {
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["divisionId"] = a.Configuration.APIClient.ParameterToString(divisionId, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementManagementunitsDivisionviews invokes GET /api/v2/workforcemanagement/managementunits/divisionviews
//
// Get management units across divisions
//
// 
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
	
	var collectionFormat string
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range divisionId {
			queryParams["divisionId"] = value
		}
	} else {
		queryParams["divisionId"] = a.Configuration.APIClient.ParameterToString(divisionId, collectionFormat)
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
	var successPayload *Managementunitlisting
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

// GetWorkforcemanagementNotifications invokes GET /api/v2/workforcemanagement/notifications
//
// Get a list of notifications for the current user
//
// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementSchedulingjob invokes GET /api/v2/workforcemanagement/schedulingjobs/{jobId}
//
// Get status of the scheduling job
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementSchedulingjob(jobId string) (*Schedulingstatusresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/schedulingjobs/{jobId}"
	path = strings.Replace(path, "{jobId}", fmt.Sprintf("%v", jobId), -1)
	defaultReturn := new(Schedulingstatusresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementShifttrades invokes GET /api/v2/workforcemanagement/shifttrades
//
// Gets all of my shift trades
//
// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementTimeoffrequest invokes GET /api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}
//
// Get a time off request for the current user
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementTimeoffrequest(timeOffRequestId string) (*Timeoffrequestresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}"
	path = strings.Replace(path, "{timeOffRequestId}", fmt.Sprintf("%v", timeOffRequestId), -1)
	defaultReturn := new(Timeoffrequestresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'timeOffRequestId' is set
	if &timeOffRequestId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementTimeoffrequests invokes GET /api/v2/workforcemanagement/timeoffrequests
//
// Get a list of time off requests for the current user
//
// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["recentlyReviewed"] = a.Configuration.APIClient.ParameterToString(recentlyReviewed, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PatchWorkforcemanagementBusinessunit invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}
//
// Update business unit
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunit(businessUnitId string, body Updatebusinessunitrequest) (*Businessunit, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	defaultReturn := new(Businessunit)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

	var successPayload *Businessunit
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

// PatchWorkforcemanagementBusinessunitActivitycode invokes PATCH /api/v2/workforcemanagement/businessunits/{buId}/activitycodes/{acId}
//
// Update an activity code
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitActivitycode(buId string, acId string, body Updateactivitycoderequest) (*Businessunitactivitycode, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{buId}/activitycodes/{acId}"
	path = strings.Replace(path, "{buId}", fmt.Sprintf("%v", buId), -1)
	path = strings.Replace(path, "{acId}", fmt.Sprintf("%v", acId), -1)
	defaultReturn := new(Businessunitactivitycode)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'buId' is set
	if &buId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'buId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitActivitycode")
	}
	// verify the required parameter 'acId' is set
	if &acId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'acId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitActivitycode")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Businessunitactivitycode
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

// PatchWorkforcemanagementBusinessunitPlanninggroup invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}
//
// Updates the planning group
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitPlanninggroup(businessUnitId string, planningGroupId string, body Updateplanninggrouprequest) (*Planninggroup, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{planningGroupId}", fmt.Sprintf("%v", planningGroupId), -1)
	defaultReturn := new(Planninggroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitPlanninggroup")
	}
	// verify the required parameter 'planningGroupId' is set
	if &planningGroupId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PatchWorkforcemanagementBusinessunitSchedulingRun invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}
//
// Mark a schedule run as applied
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitSchedulingRun(businessUnitId string, runId string, body Patchbuschedulerunrequest) (*APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{runId}", fmt.Sprintf("%v", runId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitSchedulingRun")
	}
	// verify the required parameter 'runId' is set
	if &runId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PatchWorkforcemanagementBusinessunitServicegoaltemplate invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates/{serviceGoalTemplateId}
//
// Updates a service goal template
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitServicegoaltemplate(businessUnitId string, serviceGoalTemplateId string, body Updateservicegoaltemplate) (*Servicegoaltemplate, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates/{serviceGoalTemplateId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{serviceGoalTemplateId}", fmt.Sprintf("%v", serviceGoalTemplateId), -1)
	defaultReturn := new(Servicegoaltemplate)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitServicegoaltemplate")
	}
	// verify the required parameter 'serviceGoalTemplateId' is set
	if &serviceGoalTemplateId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PatchWorkforcemanagementManagementunit invokes PATCH /api/v2/workforcemanagement/managementunits/{muId}
//
// Update the requested management unit
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunit(muId string, body Updatemanagementunitrequest) (*Managementunit, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Managementunit)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunit")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Managementunit
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

// PatchWorkforcemanagementManagementunitActivitycode invokes PATCH /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}
//
// Update an activity code
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitActivitycode(muId string, acId string, body Updateactivitycoderequest) (*Activitycode, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	path = strings.Replace(path, "{acId}", fmt.Sprintf("%v", acId), -1)
	defaultReturn := new(Activitycode)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitActivitycode")
	}
	// verify the required parameter 'acId' is set
	if &acId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'acId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitActivitycode")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Activitycode
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

// PatchWorkforcemanagementManagementunitSchedulingRun invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}
//
// Marks a specific scheduling run as applied, allowing a new rescheduling run to be started
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitSchedulingRun(managementUnitId string, runId string, body Updateschedulingrunrequest) (*Rescheduleresult, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{runId}", fmt.Sprintf("%v", runId), -1)
	defaultReturn := new(Rescheduleresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitSchedulingRun")
	}
	// verify the required parameter 'runId' is set
	if &runId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'runId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitSchedulingRun")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Rescheduleresult
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

// PatchWorkforcemanagementManagementunitServicegoalgroup invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/servicegoalgroups/{serviceGoalGroupId}
//
// Update a service goal group
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitServicegoalgroup(managementUnitId string, serviceGoalGroupId string, body Servicegoalgroup) (*Servicegoalgroup, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/servicegoalgroups/{serviceGoalGroupId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{serviceGoalGroupId}", fmt.Sprintf("%v", serviceGoalGroupId), -1)
	defaultReturn := new(Servicegoalgroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitServicegoalgroup")
	}
	// verify the required parameter 'serviceGoalGroupId' is set
	if &serviceGoalGroupId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'serviceGoalGroupId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitServicegoalgroup")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Servicegoalgroup
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

// PatchWorkforcemanagementManagementunitSettings invokes PATCH /api/v2/workforcemanagement/managementunits/{muId}/settings
//
// Update the settings for the requested management unit
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitSettings(muId string, body Managementunitsettingsrequest) (*Managementunitsettingsresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/settings"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Managementunitsettingsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitSettings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Managementunitsettingsresponse
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

// PatchWorkforcemanagementManagementunitUserTimeoffrequest invokes PATCH /api/v2/workforcemanagement/managementunits/{muId}/users/{userId}/timeoffrequests/{timeOffRequestId}
//
// Update a time off request
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitUserTimeoffrequest(muId string, userId string, timeOffRequestId string, body Admintimeoffrequestpatch) (*Timeoffrequestresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/users/{userId}/timeoffrequests/{timeOffRequestId}"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)
	path = strings.Replace(path, "{timeOffRequestId}", fmt.Sprintf("%v", timeOffRequestId), -1)
	defaultReturn := new(Timeoffrequestresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitUserTimeoffrequest")
	}
	// verify the required parameter 'userId' is set
	if &userId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'userId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitUserTimeoffrequest")
	}
	// verify the required parameter 'timeOffRequestId' is set
	if &timeOffRequestId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PatchWorkforcemanagementManagementunitWeekSchedule invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}
//
// Update a week schedule
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitWeekSchedule(managementUnitId string, weekId string, scheduleId string, forceAsync bool, forceDownloadService bool, body Updateweekschedulerequest) (*Asyncweekscheduleresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	defaultReturn := new(Asyncweekscheduleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWeekSchedule")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWeekSchedule")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWeekSchedule")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

	var successPayload *Asyncweekscheduleresponse
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

// PatchWorkforcemanagementManagementunitWeekShifttrade invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}
//
// Updates a shift trade. This route can only be called by the initiating agent
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitWeekShifttrade(managementUnitId string, weekDateId time.Time, body Patchshifttraderequest, tradeId string) (*Shifttraderesponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	path = strings.Replace(path, "{tradeId}", fmt.Sprintf("%v", tradeId), -1)
	defaultReturn := new(Shifttraderesponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWeekShifttrade")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWeekShifttrade")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWeekShifttrade")
	}
	// verify the required parameter 'tradeId' is set
	if &tradeId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'tradeId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWeekShifttrade")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Shifttraderesponse
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

// PatchWorkforcemanagementManagementunitWorkplan invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}
//
// Update a work plan
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitWorkplan(managementUnitId string, workPlanId string, body Workplan) (*Workplan, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{workPlanId}", fmt.Sprintf("%v", workPlanId), -1)
	defaultReturn := new(Workplan)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWorkplan")
	}
	// verify the required parameter 'workPlanId' is set
	if &workPlanId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PatchWorkforcemanagementTimeoffrequest invokes PATCH /api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}
//
// Update a time off request for the current user
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementTimeoffrequest(timeOffRequestId string, body Agenttimeoffrequestpatch) (*Timeoffrequestresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/timeoffrequests/{timeOffRequestId}"
	path = strings.Replace(path, "{timeOffRequestId}", fmt.Sprintf("%v", timeOffRequestId), -1)
	defaultReturn := new(Timeoffrequestresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'timeOffRequestId' is set
	if &timeOffRequestId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementAdherenceHistorical invokes POST /api/v2/workforcemanagement/adherence/historical
//
// Request a historical adherence report for users across management units
//
// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementAgentschedulesMine invokes POST /api/v2/workforcemanagement/agentschedules/mine
//
// Get published schedule for the current user
//
// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementBusinessunitActivitycodes invokes POST /api/v2/workforcemanagement/businessunits/{buId}/activitycodes
//
// Create a new activity code
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitActivitycodes(buId string, body Createactivitycoderequest) (*Businessunitactivitycode, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{buId}/activitycodes"
	path = strings.Replace(path, "{buId}", fmt.Sprintf("%v", buId), -1)
	defaultReturn := new(Businessunitactivitycode)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'buId' is set
	if &buId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'buId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitActivitycodes")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Businessunitactivitycode
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

// PostWorkforcemanagementBusinessunitAgentschedulesSearch invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/agentschedules/search
//
// Search published schedules
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitAgentschedulesSearch(businessUnitId string, body Busearchagentschedulesrequest, forceAsync bool, forceDownloadService bool) (*Buasyncagentschedulessearchresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/agentschedules/search"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	defaultReturn := new(Buasyncagentschedulessearchresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementBusinessunitIntraday invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/intraday
//
// Get intraday data for the given date for the requested planningGroupIds
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitIntraday(businessUnitId string, forceAsync bool, body Intradayplanninggrouprequest) (*Asyncintradayresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/intraday"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	defaultReturn := new(Asyncintradayresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementBusinessunitPlanninggroups invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups
//
// Adds a new planning group
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitPlanninggroups(businessUnitId string, body Createplanninggrouprequest) (*Planninggroup, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	defaultReturn := new(Planninggroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementBusinessunitServicegoaltemplates invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates
//
// Adds a new service goal template
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitServicegoaltemplates(businessUnitId string, body Createservicegoaltemplate) (*Servicegoaltemplate, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/servicegoaltemplates"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	defaultReturn := new(Servicegoaltemplate)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQuery invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query
//
// Loads agent schedule data from the schedule. Used in combination with the metadata route
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQuery(businessUnitId string, weekId time.Time, scheduleId string, body Buqueryagentschedulesrequest, forceAsync bool, forceDownloadService bool) (*Buasyncagentschedulesqueryresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	defaultReturn := new(Buasyncagentschedulesqueryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQuery")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQuery")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleAgentschedulesQuery")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementBusinessunitWeekScheduleCopy invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy
//
// Copy a schedule
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekScheduleCopy(businessUnitId string, weekId time.Time, scheduleId string, body Bucopyschedulerequest) (*Buasyncscheduleresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	defaultReturn := new(Buasyncscheduleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleCopy")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleCopy")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleCopy")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementBusinessunitWeekScheduleReschedule invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule
//
// Start a rescheduling run
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekScheduleReschedule(businessUnitId string, weekId time.Time, scheduleId string, body Bureschedulerequest) (*Buasyncschedulerunresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	defaultReturn := new(Buasyncschedulerunresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleReschedule")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleReschedule")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekScheduleReschedule")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementBusinessunitWeekSchedules invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules
//
// Create a blank schedule
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekSchedules(businessUnitId string, weekId time.Time, body Bucreateblankschedulerequest) (*Buschedulemetadata, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	defaultReturn := new(Buschedulemetadata)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedules")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedules")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementBusinessunitWeekSchedulesGenerate invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate
//
// Generate a schedule
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekSchedulesGenerate(businessUnitId string, weekId time.Time, body Bugenerateschedulerequest) (*Buasyncschedulerunresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/generate"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	defaultReturn := new(Buasyncschedulerunresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulesGenerate")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekSchedulesGenerate")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementBusinessunitWeekShorttermforecastCopy invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/copy
//
// Copy a short term forecast
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekShorttermforecastCopy(businessUnitId string, weekDateId time.Time, forecastId string, body Copybuforecastrequest, forceAsync bool) (*Asyncforecastoperationresult, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/copy"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	path = strings.Replace(path, "{forecastId}", fmt.Sprintf("%v", forecastId), -1)
	defaultReturn := new(Asyncforecastoperationresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastCopy")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastCopy")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'forecastId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastCopy")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerate invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate
//
// Generate a short term forecast
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerate(businessUnitId string, weekDateId time.Time, body Generatebuforecastrequest, forceAsync bool) (*Asyncforecastoperationresult, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/generate"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	defaultReturn := new(Asyncforecastoperationresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerate")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementBusinessunitWeekShorttermforecastsGenerate")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementBusinessunits invokes POST /api/v2/workforcemanagement/businessunits
//
// Add a new business unit
//
// It may take a minute or two for a new business unit to be available for api operations
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunits(body Createbusinessunitrequest) (*Businessunit, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits"
	defaultReturn := new(Businessunit)
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

	var successPayload *Businessunit
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

// PostWorkforcemanagementManagementunitActivitycodes invokes POST /api/v2/workforcemanagement/managementunits/{muId}/activitycodes
//
// Create a new activity code
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitActivitycodes(muId string, body Createactivitycoderequest) (*Activitycode, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/activitycodes"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Activitycode)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitActivitycodes")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Activitycode
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

// PostWorkforcemanagementManagementunitAgentschedulesSearch invokes POST /api/v2/workforcemanagement/managementunits/{muId}/agentschedules/search
//
// Query published schedules for given given time range for set of users
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitAgentschedulesSearch(muId string, body Busearchagentschedulesrequest, forceAsync bool, forceDownloadService bool) (*Buasyncagentschedulessearchresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/agentschedules/search"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Buasyncagentschedulessearchresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitAgentschedulesSearch")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementManagementunitHistoricaladherencequery invokes POST /api/v2/workforcemanagement/managementunits/{muId}/historicaladherencequery
//
// Request a historical adherence report
//
// The maximum supported range for historical adherence queries is 31 days, or 7 days with includeExceptions = true
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitHistoricaladherencequery(muId string, body Wfmhistoricaladherencequery) (*Wfmhistoricaladherenceresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/historicaladherencequery"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Wfmhistoricaladherenceresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitHistoricaladherencequery")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Wfmhistoricaladherenceresponse
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

// PostWorkforcemanagementManagementunitIntraday invokes POST /api/v2/workforcemanagement/managementunits/{muId}/intraday
//
// Get intraday data for the given date for the requested queueIds
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitIntraday(muId string, body Intradayquerydatacommand) (*Intradayresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/intraday"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Intradayresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitIntraday")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Intradayresponse
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

// PostWorkforcemanagementManagementunitMove invokes POST /api/v2/workforcemanagement/managementunits/{muId}/move
//
// Move the requested management unit to a new business unit
//
// Returns status 200 if the management unit is already in the requested business unit
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitMove(muId string, body Movemanagementunitrequest) (*Movemanagementunitresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/move"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Movemanagementunitresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitMove")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Movemanagementunitresponse
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

// PostWorkforcemanagementManagementunitSchedulesSearch invokes POST /api/v2/workforcemanagement/managementunits/{muId}/schedules/search
//
// Query published schedules for given given time range for set of users
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitSchedulesSearch(muId string, body Userlistschedulerequestbody) (*Userschedulecontainer, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/schedules/search"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Userschedulecontainer)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitSchedulesSearch")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Userschedulecontainer
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

// PostWorkforcemanagementManagementunitServicegoalgroups invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/servicegoalgroups
//
// Create a new service goal group
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitServicegoalgroups(managementUnitId string, body Createservicegoalgrouprequest) (*Servicegoalgroup, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/servicegoalgroups"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Servicegoalgroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitServicegoalgroups")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Servicegoalgroup
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

// PostWorkforcemanagementManagementunitTimeoffrequests invokes POST /api/v2/workforcemanagement/managementunits/{muId}/timeoffrequests
//
// Create a new time off request
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitTimeoffrequests(muId string, body Createadmintimeoffrequest) (*Timeoffrequestlist, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/timeoffrequests"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Timeoffrequestlist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitTimeoffrequests")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Timeoffrequestlist
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

// PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetails invokes POST /api/v2/workforcemanagement/managementunits/{muId}/timeoffrequests/fetchdetails
//
// Gets a list of time off requests from lookup ids
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetails(muId string, body Timeoffrequestlookuplist) (*Timeoffrequestentitylist, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/timeoffrequests/fetchdetails"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Timeoffrequestentitylist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetails")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Timeoffrequestentitylist
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

// PostWorkforcemanagementManagementunitTimeoffrequestsQuery invokes POST /api/v2/workforcemanagement/managementunits/{muId}/timeoffrequests/query
//
// Gets the lookup ids to fetch the specified set of requests
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitTimeoffrequestsQuery(muId string, body Timeoffrequestquerybody) (*Timeoffrequestlookuplist, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{muId}/timeoffrequests/query"
	path = strings.Replace(path, "{muId}", fmt.Sprintf("%v", muId), -1)
	defaultReturn := new(Timeoffrequestlookuplist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'muId' is set
	if &muId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'muId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitTimeoffrequestsQuery")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Timeoffrequestlookuplist
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

// PostWorkforcemanagementManagementunitWeekScheduleCopy invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy
//
// Copy a week schedule
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekScheduleCopy(managementUnitId string, weekId string, scheduleId string, forceAsync bool, forceDownloadService bool, body Copyweekschedulerequest) (*Asyncweekscheduleresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/copy"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	defaultReturn := new(Asyncweekscheduleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekScheduleCopy")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekScheduleCopy")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekScheduleCopy")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

	var successPayload *Asyncweekscheduleresponse
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

// PostWorkforcemanagementManagementunitWeekScheduleReschedule invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule
//
// Start a scheduling run to compute the reschedule. When the scheduling run finishes, a client can get the reschedule changes and then the client can apply them to the schedule, save the schedule, and mark the scheduling run as applied
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekScheduleReschedule(managementUnitId string, weekId string, scheduleId string, body Reschedulerequest) (*Asyncweekscheduleresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/{scheduleId}/reschedule"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	defaultReturn := new(Asyncweekscheduleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekScheduleReschedule")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekScheduleReschedule")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekScheduleReschedule")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Asyncweekscheduleresponse
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

// PostWorkforcemanagementManagementunitWeekSchedules invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules
//
// Add a schedule for a week in management unit using imported data. Use partial uploads of user schedules if activity count in schedule is greater than 17500
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekSchedules(managementUnitId string, weekId string, forceAsync bool, forceDownloadService bool, body Importweekschedulerequest) (*Asyncweekscheduleresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	defaultReturn := new(Asyncweekscheduleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekSchedules")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekSchedules")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	
	
	
	
		collectionFormat = ""
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

	var successPayload *Asyncweekscheduleresponse
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

// PostWorkforcemanagementManagementunitWeekSchedulesGenerate invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/generate
//
// Generate a week schedule
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekSchedulesGenerate(managementUnitId string, weekId string, body Generateweekschedulerequest) (*Generateweekscheduleresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/generate"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	defaultReturn := new(Generateweekscheduleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekSchedulesGenerate")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekSchedulesGenerate")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Generateweekscheduleresponse
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

// PostWorkforcemanagementManagementunitWeekSchedulesPartialupload invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/partialupload
//
// Partial upload of user schedules where activity count is greater than 17500
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekSchedulesPartialupload(managementUnitId string, weekId string, body Userschedulespartialuploadrequest) (*Partialuploadresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules/partialupload"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	defaultReturn := new(Partialuploadresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekSchedulesPartialupload")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekSchedulesPartialupload")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Partialuploadresponse
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

// PostWorkforcemanagementManagementunitWeekShifttradeMatch invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}/match
//
// Matches a shift trade. This route can only be called by the receiving agent
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekShifttradeMatch(managementUnitId string, weekDateId time.Time, body Matchshifttraderequest, tradeId string) (*Matchshifttraderesponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}/match"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	path = strings.Replace(path, "{tradeId}", fmt.Sprintf("%v", tradeId), -1)
	defaultReturn := new(Matchshifttraderesponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradeMatch")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradeMatch")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradeMatch")
	}
	// verify the required parameter 'tradeId' is set
	if &tradeId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'tradeId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradeMatch")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Matchshifttraderesponse
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

// PostWorkforcemanagementManagementunitWeekShifttrades invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades
//
// Adds a shift trade
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekShifttrades(managementUnitId string, weekDateId time.Time, body Addshifttraderequest) (*Shifttraderesponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	defaultReturn := new(Shifttraderesponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttrades")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttrades")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementManagementunitWeekShifttradesSearch invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/search
//
// Searches for potential shift trade matches for the current agent
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekShifttradesSearch(managementUnitId string, weekDateId time.Time, body Searchshifttradesrequest) (*Searchshifttradesresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/search"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	defaultReturn := new(Searchshifttradesresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradesSearch")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradesSearch")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementManagementunitWeekShifttradesStateBulk invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/state/bulk
//
// Updates the state of a batch of shift trades
//
// Admin functionality is not supported with \&quot;mine\&quot;.
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekShifttradesStateBulk(managementUnitId string, weekDateId time.Time, body Bulkshifttradestateupdaterequest, forceAsync bool) (*Bulkupdateshifttradestateresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/state/bulk"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	defaultReturn := new(Bulkupdateshifttradestateresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradesStateBulk")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShifttradesStateBulk")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
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
	
	var collectionFormat string
	
	
	
		collectionFormat = ""
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementManagementunitWeekShorttermforecastCopy invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/copy
//
// Copy a short term forecast
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekShorttermforecastCopy(managementUnitId string, weekDateId string, forecastId string, body Copyshorttermforecastrequest, forceAsync bool) (*Shorttermforecastresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/copy"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	path = strings.Replace(path, "{forecastId}", fmt.Sprintf("%v", forecastId), -1)
	defaultReturn := new(Shorttermforecastresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShorttermforecastCopy")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShorttermforecastCopy")
	}
	// verify the required parameter 'forecastId' is set
	if &forecastId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'forecastId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShorttermforecastCopy")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShorttermforecastCopy")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

	var successPayload *Shorttermforecastresponse
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

// PostWorkforcemanagementManagementunitWeekShorttermforecasts invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts
//
// Import a short term forecast
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekShorttermforecasts(managementUnitId string, weekDateId string, body Importshorttermforecastrequest, forceAsync bool) (*Shorttermforecastresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	defaultReturn := new(Shorttermforecastresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShorttermforecasts")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShorttermforecasts")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShorttermforecasts")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

	var successPayload *Shorttermforecastresponse
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

// PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerate invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate
//
// Generate a short term forecast
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerate(managementUnitId string, weekDateId string, body Generateshorttermforecastrequest, forceAsync bool) (*Generateshorttermforecastresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/generate"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	defaultReturn := new(Generateshorttermforecastresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerate")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerate")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShorttermforecastsGenerate")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

	var successPayload *Generateshorttermforecastresponse
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

// PostWorkforcemanagementManagementunitWeekShorttermforecastsPartialupload invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/partialupload
//
// Import a short term forecast
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWeekShorttermforecastsPartialupload(managementUnitId string, weekDateId string, body Routegrouplist) (*Partialuploadresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shorttermforecasts/partialupload"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{weekDateId}", fmt.Sprintf("%v", weekDateId), -1)
	defaultReturn := new(Partialuploadresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShorttermforecastsPartialupload")
	}
	// verify the required parameter 'weekDateId' is set
	if &weekDateId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekDateId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShorttermforecastsPartialupload")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWeekShorttermforecastsPartialupload")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Partialuploadresponse
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

// PostWorkforcemanagementManagementunitWorkplanCopy invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/copy
//
// Create a copy of work plan
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWorkplanCopy(managementUnitId string, workPlanId string, body Copyworkplan) (*Workplan, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/copy"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{workPlanId}", fmt.Sprintf("%v", workPlanId), -1)
	defaultReturn := new(Workplan)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWorkplanCopy")
	}
	// verify the required parameter 'workPlanId' is set
	if &workPlanId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementManagementunitWorkplans invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans
//
// Create a new work plan
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWorkplans(managementUnitId string, body Createworkplan) (*Workplan, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Workplan)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementNotificationsUpdate invokes POST /api/v2/workforcemanagement/notifications/update
//
// Mark a list of notifications as read or unread
//
// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementSchedules invokes POST /api/v2/workforcemanagement/schedules
//
// Get published schedule for the current user
//
// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementTimeoffrequests invokes POST /api/v2/workforcemanagement/timeoffrequests
//
// Create a time off request for the current user
//
// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

