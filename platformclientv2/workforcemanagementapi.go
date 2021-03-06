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

// DeleteWorkforcemanagementBusinessunitActivitycode invokes DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}
//
// Deletes an activity code
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementBusinessunitActivitycode(businessUnitId string, activityCodeId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{activityCodeId}", fmt.Sprintf("%v", activityCodeId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementBusinessunitActivitycode")
	}
	// verify the required parameter 'activityCodeId' is set
	if &activityCodeId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteWorkforcemanagementManagementunit invokes DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}
//
// Delete management unit
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunit(managementUnitId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteWorkforcemanagementManagementunitWorkplanrotation invokes DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}
//
// Delete a work plan rotation
//
// 
func (a WorkforceManagementApi) DeleteWorkforcemanagementManagementunitWorkplanrotation(managementUnitId string, workPlanRotationId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{workPlanRotationId}", fmt.Sprintf("%v", workPlanRotationId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->DeleteWorkforcemanagementManagementunitWorkplanrotation")
	}
	// verify the required parameter 'workPlanRotationId' is set
	if &workPlanRotationId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitActivitycode invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}
//
// Get an activity code
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitActivitycode(businessUnitId string, activityCodeId string) (*Businessunitactivitycode, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{activityCodeId}", fmt.Sprintf("%v", activityCodeId), -1)
	defaultReturn := new(Businessunitactivitycode)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitActivitycode")
	}
	// verify the required parameter 'activityCodeId' is set
	if &activityCodeId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementBusinessunitActivitycodes invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes
//
// Get activity codes
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitActivitycodes(businessUnitId string) (*Businessunitactivitycodelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	defaultReturn := new(Businessunitactivitycodelisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	if str, ok := interface{}(date).(string); ok {
		if str != "" {
			queryParams["date"] = a.Configuration.APIClient.ParameterToString(date, collectionFormat)
		}
	} else {
		queryParams["date"] = a.Configuration.APIClient.ParameterToString(date, collectionFormat)
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
	if str, ok := interface{}(feature).(string); ok {
		if str != "" {
			queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, collectionFormat)
		}
	} else {
		queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(divisionId).(string); ok {
		if str != "" {
			queryParams["divisionId"] = a.Configuration.APIClient.ParameterToString(divisionId, collectionFormat)
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
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekSchedule(businessUnitId string, weekId time.Time, scheduleId string, expand string) (*Buschedulemetadata, *APIResponse, error) {
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
	
	var collectionFormat string
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(expand).(string); ok {
		if str != "" {
			queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
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
	if str, ok := interface{}(forceDownload).(string); ok {
		if str != "" {
			queryParams["forceDownload"] = a.Configuration.APIClient.ParameterToString(forceDownload, collectionFormat)
		}
	} else {
		queryParams["forceDownload"] = a.Configuration.APIClient.ParameterToString(forceDownload, collectionFormat)
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

// GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgent invokes GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}
//
// Loads agent&#39;s schedule history.
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgent(businessUnitId string, weekId time.Time, scheduleId string, agentId string) (*Buagentschedulehistoryresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/history/agents/{agentId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{weekId}", fmt.Sprintf("%v", weekId), -1)
	path = strings.Replace(path, "{scheduleId}", fmt.Sprintf("%v", scheduleId), -1)
	path = strings.Replace(path, "{agentId}", fmt.Sprintf("%v", agentId), -1)
	defaultReturn := new(Buagentschedulehistoryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgent")
	}
	// verify the required parameter 'weekId' is set
	if &weekId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'weekId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgent")
	}
	// verify the required parameter 'scheduleId' is set
	if &scheduleId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'scheduleId' when calling WorkforceManagementApi->GetWorkforcemanagementBusinessunitWeekScheduleHistoryAgent")
	}
	// verify the required parameter 'agentId' is set
	if &agentId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	if str, ok := interface{}(includeOnlyPublished).(string); ok {
		if str != "" {
			queryParams["includeOnlyPublished"] = a.Configuration.APIClient.ParameterToString(includeOnlyPublished, collectionFormat)
		}
	} else {
		queryParams["includeOnlyPublished"] = a.Configuration.APIClient.ParameterToString(includeOnlyPublished, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(expand).(string); ok {
		if str != "" {
			queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
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
func (a WorkforceManagementApi) GetWorkforcemanagementBusinessunitWeekShorttermforecastData(businessUnitId string, weekDateId time.Time, forecastId string, weekNumber int, forceDownloadService bool) (*Buforecastresultresponse, *APIResponse, error) {
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
	if str, ok := interface{}(weekNumber).(string); ok {
		if str != "" {
			queryParams["weekNumber"] = a.Configuration.APIClient.ParameterToString(weekNumber, collectionFormat)
		}
	} else {
		queryParams["weekNumber"] = a.Configuration.APIClient.ParameterToString(weekNumber, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(forceDownloadService).(string); ok {
		if str != "" {
			queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
		}
	} else {
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
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
	if str, ok := interface{}(feature).(string); ok {
		if str != "" {
			queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, collectionFormat)
		}
	} else {
		queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(divisionId).(string); ok {
		if str != "" {
			queryParams["divisionId"] = a.Configuration.APIClient.ParameterToString(divisionId, collectionFormat)
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

// GetWorkforcemanagementManagementunit invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}
//
// Get management unit
//
// settings.shortTermForecasting is deprecated and now lives on the business unit
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunit(managementUnitId string, expand []string) (*Managementunit, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Managementunit)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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

// GetWorkforcemanagementManagementunitActivitycodes invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/activitycodes
//
// Get activity codes
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitActivitycodes(managementUnitId string) (*Activitycodecontainer, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/activitycodes"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Activitycodecontainer)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	if str, ok := interface{}(forceDownloadService).(string); ok {
		if str != "" {
			queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
		}
	} else {
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
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
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitAgent(managementUnitId string, agentId string, excludeCapabilities bool) (*Wfmagent, *APIResponse, error) {
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
	
	var collectionFormat string
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(excludeCapabilities).(string); ok {
		if str != "" {
			queryParams["excludeCapabilities"] = a.Configuration.APIClient.ParameterToString(excludeCapabilities, collectionFormat)
		}
	} else {
		queryParams["excludeCapabilities"] = a.Configuration.APIClient.ParameterToString(excludeCapabilities, collectionFormat)
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

// GetWorkforcemanagementManagementunitShifttradesMatched invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/shifttrades/matched
//
// Gets a summary of all shift trades in the matched state
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitShifttradesMatched(managementUnitId string) (*Shifttradematchessummaryresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/shifttrades/matched"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Shifttradematchessummaryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementManagementunitShifttradesUsers invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/shifttrades/users
//
// Gets list of users available for whom you can send direct shift trade requests
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitShifttradesUsers(managementUnitId string) (*Wfmuserentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/shifttrades/users"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Wfmuserentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementManagementunitUserTimeoffrequest invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}
//
// Get a time off request
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitUserTimeoffrequest(managementUnitId string, userId string, timeOffRequestId string) (*Timeoffrequestresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)
	path = strings.Replace(path, "{timeOffRequestId}", fmt.Sprintf("%v", timeOffRequestId), -1)
	defaultReturn := new(Timeoffrequestresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUserTimeoffrequest")
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

// GetWorkforcemanagementManagementunitUserTimeoffrequests invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests
//
// Get a list of time off requests for a given user
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitUserTimeoffrequests(managementUnitId string, userId string, recentlyReviewed bool) (*Timeoffrequestlist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)
	defaultReturn := new(Timeoffrequestlist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitUserTimeoffrequests")
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
	if str, ok := interface{}(recentlyReviewed).(string); ok {
		if str != "" {
			queryParams["recentlyReviewed"] = a.Configuration.APIClient.ParameterToString(recentlyReviewed, collectionFormat)
		}
	} else {
		queryParams["recentlyReviewed"] = a.Configuration.APIClient.ParameterToString(recentlyReviewed, collectionFormat)
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

// GetWorkforcemanagementManagementunitUsers invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users
//
// Get users in the management unit
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitUsers(managementUnitId string) (*Wfmuserentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/users"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Wfmuserentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
// Deprecated.  Use the equivalent business unit resource instead. Get a week schedule
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
	if str, ok := interface{}(expand).(string); ok {
		if str != "" {
			queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
		}
	} else {
		queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(forceDownloadService).(string); ok {
		if str != "" {
			queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
		}
	} else {
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
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

// GetWorkforcemanagementManagementunitWeekSchedules invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules
//
// Deprecated.  Use the equivalent business unit resource instead. Get the list of schedules in a week in management unit
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
	if str, ok := interface{}(includeOnlyPublished).(string); ok {
		if str != "" {
			queryParams["includeOnlyPublished"] = a.Configuration.APIClient.ParameterToString(includeOnlyPublished, collectionFormat)
		}
	} else {
		queryParams["includeOnlyPublished"] = a.Configuration.APIClient.ParameterToString(includeOnlyPublished, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(earliestWeekDate).(string); ok {
		if str != "" {
			queryParams["earliestWeekDate"] = a.Configuration.APIClient.ParameterToString(earliestWeekDate, collectionFormat)
		}
	} else {
		queryParams["earliestWeekDate"] = a.Configuration.APIClient.ParameterToString(earliestWeekDate, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(latestWeekDate).(string); ok {
		if str != "" {
			queryParams["latestWeekDate"] = a.Configuration.APIClient.ParameterToString(latestWeekDate, collectionFormat)
		}
	} else {
		queryParams["latestWeekDate"] = a.Configuration.APIClient.ParameterToString(latestWeekDate, collectionFormat)
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
	if str, ok := interface{}(evaluateMatches).(string); ok {
		if str != "" {
			queryParams["evaluateMatches"] = a.Configuration.APIClient.ParameterToString(evaluateMatches, collectionFormat)
		}
	} else {
		queryParams["evaluateMatches"] = a.Configuration.APIClient.ParameterToString(evaluateMatches, collectionFormat)
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

// GetWorkforcemanagementManagementunitWorkplanrotation invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}
//
// Get a work plan rotation
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWorkplanrotation(managementUnitId string, workPlanRotationId string) (*Workplanrotationresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{workPlanRotationId}", fmt.Sprintf("%v", workPlanRotationId), -1)
	defaultReturn := new(Workplanrotationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->GetWorkforcemanagementManagementunitWorkplanrotation")
	}
	// verify the required parameter 'workPlanRotationId' is set
	if &workPlanRotationId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetWorkforcemanagementManagementunitWorkplanrotations invokes GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations
//
// Get work plan rotations
//
// 
func (a WorkforceManagementApi) GetWorkforcemanagementManagementunitWorkplanrotations(managementUnitId string, expand []string) (*Workplanrotationlistresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Workplanrotationlistresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	var successPayload *Workplanrotationlistresponse
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
	if str, ok := interface{}(feature).(string); ok {
		if str != "" {
			queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, collectionFormat)
		}
	} else {
		queryParams["feature"] = a.Configuration.APIClient.ParameterToString(feature, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(divisionId).(string); ok {
		if str != "" {
			queryParams["divisionId"] = a.Configuration.APIClient.ParameterToString(divisionId, collectionFormat)
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
	if str, ok := interface{}(recentlyReviewed).(string); ok {
		if str != "" {
			queryParams["recentlyReviewed"] = a.Configuration.APIClient.ParameterToString(recentlyReviewed, collectionFormat)
		}
	} else {
		queryParams["recentlyReviewed"] = a.Configuration.APIClient.ParameterToString(recentlyReviewed, collectionFormat)
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

// PatchWorkforcemanagementBusinessunitActivitycode invokes PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}
//
// Update an activity code
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementBusinessunitActivitycode(businessUnitId string, activityCodeId string, body Updateactivitycoderequest) (*Businessunitactivitycode, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	path = strings.Replace(path, "{activityCodeId}", fmt.Sprintf("%v", activityCodeId), -1)
	defaultReturn := new(Businessunitactivitycode)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'businessUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementBusinessunitActivitycode")
	}
	// verify the required parameter 'activityCodeId' is set
	if &activityCodeId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PatchWorkforcemanagementManagementunit invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}
//
// Update the requested management unit
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunit(managementUnitId string, body Updatemanagementunitrequest) (*Managementunit, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Managementunit)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PatchWorkforcemanagementManagementunitUserTimeoffrequest invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}
//
// Update a time off request
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitUserTimeoffrequest(managementUnitId string, userId string, timeOffRequestId string, body Admintimeoffrequestpatch) (*Timeoffrequestresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userId), -1)
	path = strings.Replace(path, "{timeOffRequestId}", fmt.Sprintf("%v", timeOffRequestId), -1)
	defaultReturn := new(Timeoffrequestresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitUserTimeoffrequest")
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
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitWorkplan(managementUnitId string, workPlanId string, body Workplan, validationMode string) (*Workplan, *APIResponse, error) {
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
	
	var collectionFormat string
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(validationMode).(string); ok {
		if str != "" {
			queryParams["validationMode"] = a.Configuration.APIClient.ParameterToString(validationMode, collectionFormat)
		}
	} else {
		queryParams["validationMode"] = a.Configuration.APIClient.ParameterToString(validationMode, collectionFormat)
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

// PatchWorkforcemanagementManagementunitWorkplanrotation invokes PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}
//
// Update a work plan rotation
//
// 
func (a WorkforceManagementApi) PatchWorkforcemanagementManagementunitWorkplanrotation(managementUnitId string, workPlanRotationId string, body Updateworkplanrotationrequest) (*Workplanrotationresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{workPlanRotationId}", fmt.Sprintf("%v", workPlanRotationId), -1)
	defaultReturn := new(Workplanrotationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PatchWorkforcemanagementManagementunitWorkplanrotation")
	}
	// verify the required parameter 'workPlanRotationId' is set
	if &workPlanRotationId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementBusinessunitActivitycodes invokes POST /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes
//
// Create a new activity code
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementBusinessunitActivitycodes(businessUnitId string, body Createactivitycoderequest) (*Businessunitactivitycode, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes"
	path = strings.Replace(path, "{businessUnitId}", fmt.Sprintf("%v", businessUnitId), -1)
	defaultReturn := new(Businessunitactivitycode)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'businessUnitId' is set
	if &businessUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	if str, ok := interface{}(forceAsync).(string); ok {
		if str != "" {
			queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
		}
	} else {
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(forceDownloadService).(string); ok {
		if str != "" {
			queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
		}
	} else {
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
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
	if str, ok := interface{}(forceAsync).(string); ok {
		if str != "" {
			queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
		}
	} else {
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
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
	if str, ok := interface{}(forceAsync).(string); ok {
		if str != "" {
			queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
		}
	} else {
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(forceDownloadService).(string); ok {
		if str != "" {
			queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
		}
	} else {
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
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
	if str, ok := interface{}(forceAsync).(string); ok {
		if str != "" {
			queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
		}
	} else {
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
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
	if str, ok := interface{}(forceAsync).(string); ok {
		if str != "" {
			queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
		}
	} else {
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
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

// PostWorkforcemanagementManagementunitAgentschedulesSearch invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/agentschedules/search
//
// Query published schedules for given given time range for set of users
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitAgentschedulesSearch(managementUnitId string, body Busearchagentschedulesrequest, forceAsync bool, forceDownloadService bool) (*Buasyncagentschedulessearchresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/agentschedules/search"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Buasyncagentschedulessearchresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	
	var collectionFormat string
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(forceAsync).(string); ok {
		if str != "" {
			queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
		}
	} else {
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(forceDownloadService).(string); ok {
		if str != "" {
			queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
		}
	} else {
		queryParams["forceDownloadService"] = a.Configuration.APIClient.ParameterToString(forceDownloadService, collectionFormat)
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

// PostWorkforcemanagementManagementunitHistoricaladherencequery invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/historicaladherencequery
//
// Request a historical adherence report
//
// The maximum supported range for historical adherence queries is 31 days, or 7 days with includeExceptions = true
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitHistoricaladherencequery(managementUnitId string, body Wfmhistoricaladherencequery) (*Wfmhistoricaladherenceresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/historicaladherencequery"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Wfmhistoricaladherenceresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementManagementunitMove invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/move
//
// Move the requested management unit to a new business unit
//
// Returns status 200 if the management unit is already in the requested business unit
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitMove(managementUnitId string, body Movemanagementunitrequest) (*Movemanagementunitresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/move"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Movemanagementunitresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementManagementunitSchedulesSearch invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/schedules/search
//
// Query published schedules for given given time range for set of users
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitSchedulesSearch(managementUnitId string, body Userlistschedulerequestbody) (*Userschedulecontainer, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/schedules/search"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Userschedulecontainer)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementManagementunitTimeoffrequests invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests
//
// Create a new time off request
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitTimeoffrequests(managementUnitId string, body Createadmintimeoffrequest) (*Timeoffrequestlist, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Timeoffrequestlist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementManagementunitTimeoffrequestsQuery invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests/query
//
// Fetches time off requests matching the conditions specified in the request body
//
// Request body requires one of the following: User ID is specified, statuses == [Pending] or date range to be specified and less than or equal to 33 days.  All other fields are filters
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitTimeoffrequestsQuery(managementUnitId string, body Timeoffrequestquerybody) (*Timeoffrequestlisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/timeoffrequests/query"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Timeoffrequestlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	if str, ok := interface{}(forceAsync).(string); ok {
		if str != "" {
			queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
		}
	} else {
		queryParams["forceAsync"] = a.Configuration.APIClient.ParameterToString(forceAsync, collectionFormat)
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

// PostWorkforcemanagementManagementunitWorkplanValidate invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate
//
// Validate Work Plan
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWorkplanValidate(managementUnitId string, workPlanId string, body Workplanvalidationrequest, expand []string) (*Validateworkplanresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}/validate"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{workPlanId}", fmt.Sprintf("%v", workPlanId), -1)
	defaultReturn := new(Validateworkplanresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWorkplanValidate")
	}
	// verify the required parameter 'workPlanId' is set
	if &workPlanId == nil {
		// 
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
	// body params
	postBody = &body

	var successPayload *Validateworkplanresponse
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

// PostWorkforcemanagementManagementunitWorkplanrotationCopy invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy
//
// Create a copy of work plan rotation
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWorkplanrotationCopy(managementUnitId string, workPlanRotationId string, body Copyworkplanrotationrequest) (*Workplanrotationresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}/copy"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	path = strings.Replace(path, "{workPlanRotationId}", fmt.Sprintf("%v", workPlanRotationId), -1)
	defaultReturn := new(Workplanrotationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'managementUnitId' when calling WorkforceManagementApi->PostWorkforcemanagementManagementunitWorkplanrotationCopy")
	}
	// verify the required parameter 'workPlanRotationId' is set
	if &workPlanRotationId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// PostWorkforcemanagementManagementunitWorkplanrotations invokes POST /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations
//
// Create a new work plan rotation
//
// 
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWorkplanrotations(managementUnitId string, body Addworkplanrotationrequest) (*Workplanrotationresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations"
	path = strings.Replace(path, "{managementUnitId}", fmt.Sprintf("%v", managementUnitId), -1)
	defaultReturn := new(Workplanrotationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'managementUnitId' is set
	if &managementUnitId == nil {
		// 
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
func (a WorkforceManagementApi) PostWorkforcemanagementManagementunitWorkplans(managementUnitId string, body Createworkplan, validationMode string) (*Workplan, *APIResponse, error) {
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
	
	var collectionFormat string
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(validationMode).(string); ok {
		if str != "" {
			queryParams["validationMode"] = a.Configuration.APIClient.ParameterToString(validationMode, collectionFormat)
		}
	} else {
		queryParams["validationMode"] = a.Configuration.APIClient.ParameterToString(validationMode, collectionFormat)
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

