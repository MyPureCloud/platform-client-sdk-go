package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"time"
"encoding/json"
)

// LanguageUnderstandingApi provides functions for API endpoints
type LanguageUnderstandingApi struct {
	Configuration *Configuration
}

// NewLanguageUnderstandingApi creates an API instance using the default configuration
func NewLanguageUnderstandingApi() *LanguageUnderstandingApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating LanguageUnderstandingApi with base path: %s", strings.ToLower(config.BasePath)))
	return &LanguageUnderstandingApi{
		Configuration: config,
	}
}

// NewLanguageUnderstandingApiWithConfig creates an API instance using the provided configuration
func NewLanguageUnderstandingApiWithConfig(config *Configuration) *LanguageUnderstandingApi {
	config.Debugf("Creating LanguageUnderstandingApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &LanguageUnderstandingApi{
		Configuration: config,
	}
}

// DeleteLanguageunderstandingDomain invokes DELETE /api/v2/languageunderstanding/domains/{domainId}
//
// Delete an NLU Domain.
//
// 
func (a LanguageUnderstandingApi) DeleteLanguageunderstandingDomain(domainId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->DeleteLanguageunderstandingDomain")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteLanguageunderstandingDomainFeedbackFeedbackId invokes DELETE /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}
//
// Delete the feedback on the NLU Domain Version.
//
// 
func (a LanguageUnderstandingApi) DeleteLanguageunderstandingDomainFeedbackFeedbackId(domainId string, feedbackId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	path = strings.Replace(path, "{feedbackId}", fmt.Sprintf("%v", feedbackId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->DeleteLanguageunderstandingDomainFeedbackFeedbackId")
	}
	// verify the required parameter 'feedbackId' is set
	if &feedbackId == nil {
		// 
		return nil, errors.New("Missing required parameter 'feedbackId' when calling LanguageUnderstandingApi->DeleteLanguageunderstandingDomainFeedbackFeedbackId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteLanguageunderstandingDomainVersion invokes DELETE /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}
//
// Delete an NLU Domain Version
//
// 
func (a LanguageUnderstandingApi) DeleteLanguageunderstandingDomainVersion(domainId string, domainVersionId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	path = strings.Replace(path, "{domainVersionId}", fmt.Sprintf("%v", domainVersionId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->DeleteLanguageunderstandingDomainVersion")
	}
	// verify the required parameter 'domainVersionId' is set
	if &domainVersionId == nil {
		// 
		return nil, errors.New("Missing required parameter 'domainVersionId' when calling LanguageUnderstandingApi->DeleteLanguageunderstandingDomainVersion")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// GetLanguageunderstandingDomain invokes GET /api/v2/languageunderstanding/domains/{domainId}
//
// Find an NLU Domain.
//
// 
func (a LanguageUnderstandingApi) GetLanguageunderstandingDomain(domainId string) (*Nludomain, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	defaultReturn := new(Nludomain)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->GetLanguageunderstandingDomain")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Nludomain
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetLanguageunderstandingDomainFeedback invokes GET /api/v2/languageunderstanding/domains/{domainId}/feedback
//
// Get all feedback in the given NLU Domain Version.
//
// 
func (a LanguageUnderstandingApi) GetLanguageunderstandingDomainFeedback(domainId string, intentName string, assessment string, dateStart time.Time, dateEnd time.Time, includeDeleted bool, pageNumber int, pageSize int, enableCursorPagination bool, after string, fields []string) (*Nlufeedbacklisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}/feedback"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	defaultReturn := new(Nlufeedbacklisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->GetLanguageunderstandingDomainFeedback")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	if str, ok := interface{}(intentName).(string); ok {
		if str != "" {
			queryParams["intentName"] = a.Configuration.APIClient.ParameterToString(intentName, collectionFormat)
		}
	} else {
		queryParams["intentName"] = a.Configuration.APIClient.ParameterToString(intentName, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(assessment).(string); ok {
		if str != "" {
			queryParams["assessment"] = a.Configuration.APIClient.ParameterToString(assessment, collectionFormat)
		}
	} else {
		queryParams["assessment"] = a.Configuration.APIClient.ParameterToString(assessment, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(dateStart).(string); ok {
		if str != "" {
			queryParams["dateStart"] = a.Configuration.APIClient.ParameterToString(dateStart, collectionFormat)
		}
	} else {
		queryParams["dateStart"] = a.Configuration.APIClient.ParameterToString(dateStart, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(dateEnd).(string); ok {
		if str != "" {
			queryParams["dateEnd"] = a.Configuration.APIClient.ParameterToString(dateEnd, collectionFormat)
		}
	} else {
		queryParams["dateEnd"] = a.Configuration.APIClient.ParameterToString(dateEnd, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(includeDeleted).(string); ok {
		if str != "" {
			queryParams["includeDeleted"] = a.Configuration.APIClient.ParameterToString(includeDeleted, collectionFormat)
		}
	} else {
		queryParams["includeDeleted"] = a.Configuration.APIClient.ParameterToString(includeDeleted, collectionFormat)
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
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(enableCursorPagination).(string); ok {
		if str != "" {
			queryParams["enableCursorPagination"] = a.Configuration.APIClient.ParameterToString(enableCursorPagination, collectionFormat)
		}
	} else {
		queryParams["enableCursorPagination"] = a.Configuration.APIClient.ParameterToString(enableCursorPagination, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(after).(string); ok {
		if str != "" {
			queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
		}
	} else {
		queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
	}
	
	
	
	collectionFormat = "multi"
	if collectionFormat == "multi" {
		for _, value := range fields {
			queryParams["fields"] = value
		}
	} else {
		queryParams["fields"] = a.Configuration.APIClient.ParameterToString(fields, collectionFormat)
	}
	
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Nlufeedbacklisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetLanguageunderstandingDomainFeedbackFeedbackId invokes GET /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}
//
// Find a Feedback
//
// 
func (a LanguageUnderstandingApi) GetLanguageunderstandingDomainFeedbackFeedbackId(domainId string, feedbackId string, fields []string) (*Nlufeedbackresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	path = strings.Replace(path, "{feedbackId}", fmt.Sprintf("%v", feedbackId), -1)
	defaultReturn := new(Nlufeedbackresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->GetLanguageunderstandingDomainFeedbackFeedbackId")
	}
	// verify the required parameter 'feedbackId' is set
	if &feedbackId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'feedbackId' when calling LanguageUnderstandingApi->GetLanguageunderstandingDomainFeedbackFeedbackId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
		for _, value := range fields {
			queryParams["fields"] = value
		}
	} else {
		queryParams["fields"] = a.Configuration.APIClient.ParameterToString(fields, collectionFormat)
	}
	
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Nlufeedbackresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetLanguageunderstandingDomainVersion invokes GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}
//
// Find an NLU Domain Version.
//
// 
func (a LanguageUnderstandingApi) GetLanguageunderstandingDomainVersion(domainId string, domainVersionId string, includeUtterances bool) (*Nludomainversion, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	path = strings.Replace(path, "{domainVersionId}", fmt.Sprintf("%v", domainVersionId), -1)
	defaultReturn := new(Nludomainversion)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->GetLanguageunderstandingDomainVersion")
	}
	// verify the required parameter 'domainVersionId' is set
	if &domainVersionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainVersionId' when calling LanguageUnderstandingApi->GetLanguageunderstandingDomainVersion")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	if str, ok := interface{}(includeUtterances).(string); ok {
		if str != "" {
			queryParams["includeUtterances"] = a.Configuration.APIClient.ParameterToString(includeUtterances, collectionFormat)
		}
	} else {
		queryParams["includeUtterances"] = a.Configuration.APIClient.ParameterToString(includeUtterances, collectionFormat)
	}
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Nludomainversion
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetLanguageunderstandingDomainVersionReport invokes GET /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report
//
// Retrieved quality report for the specified NLU Domain Version
//
// 
func (a LanguageUnderstandingApi) GetLanguageunderstandingDomainVersionReport(domainId string, domainVersionId string) (*Nludomainversionqualityreport, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/report"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	path = strings.Replace(path, "{domainVersionId}", fmt.Sprintf("%v", domainVersionId), -1)
	defaultReturn := new(Nludomainversionqualityreport)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->GetLanguageunderstandingDomainVersionReport")
	}
	// verify the required parameter 'domainVersionId' is set
	if &domainVersionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainVersionId' when calling LanguageUnderstandingApi->GetLanguageunderstandingDomainVersionReport")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Nludomainversionqualityreport
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetLanguageunderstandingDomainVersions invokes GET /api/v2/languageunderstanding/domains/{domainId}/versions
//
// Get all NLU Domain Versions for a given Domain.
//
// 
func (a LanguageUnderstandingApi) GetLanguageunderstandingDomainVersions(domainId string, includeUtterances bool, pageNumber int, pageSize int) (*Nludomainversionlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}/versions"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	defaultReturn := new(Nludomainversionlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->GetLanguageunderstandingDomainVersions")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	if str, ok := interface{}(includeUtterances).(string); ok {
		if str != "" {
			queryParams["includeUtterances"] = a.Configuration.APIClient.ParameterToString(includeUtterances, collectionFormat)
		}
	} else {
		queryParams["includeUtterances"] = a.Configuration.APIClient.ParameterToString(includeUtterances, collectionFormat)
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
	var successPayload *Nludomainversionlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetLanguageunderstandingDomains invokes GET /api/v2/languageunderstanding/domains
//
// Get all NLU Domains.
//
// 
func (a LanguageUnderstandingApi) GetLanguageunderstandingDomains(pageNumber int, pageSize int) (*Nludomainlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains"
	defaultReturn := new(Nludomainlisting)
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
	var successPayload *Nludomainlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PatchLanguageunderstandingDomain invokes PATCH /api/v2/languageunderstanding/domains/{domainId}
//
// Update an NLU Domain.
//
// 
func (a LanguageUnderstandingApi) PatchLanguageunderstandingDomain(domainId string, body Nludomain) (*Nludomain, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	defaultReturn := new(Nludomain)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->PatchLanguageunderstandingDomain")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling LanguageUnderstandingApi->PatchLanguageunderstandingDomain")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Nludomain
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostLanguageunderstandingDomainFeedback invokes POST /api/v2/languageunderstanding/domains/{domainId}/feedback
//
// Create feedback for the NLU Domain Version.
//
// 
func (a LanguageUnderstandingApi) PostLanguageunderstandingDomainFeedback(domainId string, body Nlufeedbackrequest) (*Nlufeedbackresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}/feedback"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	defaultReturn := new(Nlufeedbackresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->PostLanguageunderstandingDomainFeedback")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling LanguageUnderstandingApi->PostLanguageunderstandingDomainFeedback")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Nlufeedbackresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostLanguageunderstandingDomainVersionDetect invokes POST /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/detect
//
// Detect intent, entities, etc. in the submitted text using the specified NLU domain version.
//
// 
func (a LanguageUnderstandingApi) PostLanguageunderstandingDomainVersionDetect(domainId string, domainVersionId string, body Nludetectionrequest) (*Nludetectionresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/detect"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	path = strings.Replace(path, "{domainVersionId}", fmt.Sprintf("%v", domainVersionId), -1)
	defaultReturn := new(Nludetectionresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->PostLanguageunderstandingDomainVersionDetect")
	}
	// verify the required parameter 'domainVersionId' is set
	if &domainVersionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainVersionId' when calling LanguageUnderstandingApi->PostLanguageunderstandingDomainVersionDetect")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling LanguageUnderstandingApi->PostLanguageunderstandingDomainVersionDetect")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Nludetectionresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostLanguageunderstandingDomainVersionPublish invokes POST /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/publish
//
// Publish the draft NLU Domain Version.
//
// 
func (a LanguageUnderstandingApi) PostLanguageunderstandingDomainVersionPublish(domainId string, domainVersionId string) (*Nludomainversion, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/publish"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	path = strings.Replace(path, "{domainVersionId}", fmt.Sprintf("%v", domainVersionId), -1)
	defaultReturn := new(Nludomainversion)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->PostLanguageunderstandingDomainVersionPublish")
	}
	// verify the required parameter 'domainVersionId' is set
	if &domainVersionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainVersionId' when calling LanguageUnderstandingApi->PostLanguageunderstandingDomainVersionPublish")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Nludomainversion
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostLanguageunderstandingDomainVersionTrain invokes POST /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/train
//
// Train the draft NLU Domain Version.
//
// 
func (a LanguageUnderstandingApi) PostLanguageunderstandingDomainVersionTrain(domainId string, domainVersionId string) (*Nludomainversiontrainingresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}/train"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	path = strings.Replace(path, "{domainVersionId}", fmt.Sprintf("%v", domainVersionId), -1)
	defaultReturn := new(Nludomainversiontrainingresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->PostLanguageunderstandingDomainVersionTrain")
	}
	// verify the required parameter 'domainVersionId' is set
	if &domainVersionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainVersionId' when calling LanguageUnderstandingApi->PostLanguageunderstandingDomainVersionTrain")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Nludomainversiontrainingresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostLanguageunderstandingDomainVersions invokes POST /api/v2/languageunderstanding/domains/{domainId}/versions
//
// Create an NLU Domain Version.
//
// 
func (a LanguageUnderstandingApi) PostLanguageunderstandingDomainVersions(domainId string, body Nludomainversion) (*Nludomainversion, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}/versions"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	defaultReturn := new(Nludomainversion)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->PostLanguageunderstandingDomainVersions")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling LanguageUnderstandingApi->PostLanguageunderstandingDomainVersions")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Nludomainversion
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostLanguageunderstandingDomains invokes POST /api/v2/languageunderstanding/domains
//
// Create an NLU Domain.
//
// 
func (a LanguageUnderstandingApi) PostLanguageunderstandingDomains(body Nludomain) (*Nludomain, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains"
	defaultReturn := new(Nludomain)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling LanguageUnderstandingApi->PostLanguageunderstandingDomains")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Nludomain
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PutLanguageunderstandingDomainVersion invokes PUT /api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}
//
// Update an NLU Domain Version.
//
// 
func (a LanguageUnderstandingApi) PutLanguageunderstandingDomainVersion(domainId string, domainVersionId string, body Nludomainversion) (*Nludomainversion, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/languageunderstanding/domains/{domainId}/versions/{domainVersionId}"
	path = strings.Replace(path, "{domainId}", fmt.Sprintf("%v", domainId), -1)
	path = strings.Replace(path, "{domainVersionId}", fmt.Sprintf("%v", domainVersionId), -1)
	defaultReturn := new(Nludomainversion)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'domainId' is set
	if &domainId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainId' when calling LanguageUnderstandingApi->PutLanguageunderstandingDomainVersion")
	}
	// verify the required parameter 'domainVersionId' is set
	if &domainVersionId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'domainVersionId' when calling LanguageUnderstandingApi->PutLanguageunderstandingDomainVersion")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling LanguageUnderstandingApi->PutLanguageunderstandingDomainVersion")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Nludomainversion
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

