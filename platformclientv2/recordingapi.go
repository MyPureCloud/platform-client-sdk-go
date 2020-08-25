package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

// RecordingApi provides functions for API endpoints
type RecordingApi struct {
	Configuration *Configuration
}

// NewRecordingApi creates an API instance using the default configuration
func NewRecordingApi() *RecordingApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating RecordingApi with base path: %s", strings.ToLower(config.BasePath)))
	return &RecordingApi{
		Configuration: config,
	}
}

// NewRecordingApiWithConfig creates an API instance using the provided configuration
func NewRecordingApiWithConfig(config *Configuration) *RecordingApi {
	config.Debugf("Creating RecordingApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &RecordingApi{
		Configuration: config,
	}
}

// DeleteConversationRecordingAnnotation invokes DELETE /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}
//
// Delete annotation
//
// 
func (a RecordingApi) DeleteConversationRecordingAnnotation(conversationId string, recordingId string, annotationId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{recordingId}", fmt.Sprintf("%v", recordingId), -1)
	path = strings.Replace(path, "{annotationId}", fmt.Sprintf("%v", annotationId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return nil, errors.New("Missing required parameter 'conversationId' when calling RecordingApi->DeleteConversationRecordingAnnotation")
	}
	// verify the required parameter 'recordingId' is set
	if &recordingId == nil {
		// 
		return nil, errors.New("Missing required parameter 'recordingId' when calling RecordingApi->DeleteConversationRecordingAnnotation")
	}
	// verify the required parameter 'annotationId' is set
	if &annotationId == nil {
		// 
		return nil, errors.New("Missing required parameter 'annotationId' when calling RecordingApi->DeleteConversationRecordingAnnotation")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteOrphanrecording invokes DELETE /api/v2/orphanrecordings/{orphanId}
//
// Deletes a single orphan recording
//
// 
func (a RecordingApi) DeleteOrphanrecording(orphanId string) (*Orphanrecording, *APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/orphanrecordings/{orphanId}"
	path = strings.Replace(path, "{orphanId}", fmt.Sprintf("%v", orphanId), -1)
	defaultReturn := new(Orphanrecording)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'orphanId' is set
	if &orphanId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'orphanId' when calling RecordingApi->DeleteOrphanrecording")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Orphanrecording
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// DeleteRecordingJob invokes DELETE /api/v2/recording/jobs/{jobId}
//
// Delete the recording bulk job
//
// 
func (a RecordingApi) DeleteRecordingJob(jobId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", fmt.Sprintf("%v", jobId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// 
		return nil, errors.New("Missing required parameter 'jobId' when calling RecordingApi->DeleteRecordingJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteRecordingMediaretentionpolicies invokes DELETE /api/v2/recording/mediaretentionpolicies
//
// Delete media retention policies
//
// Bulk delete of media retention policies, this will only delete the polices that match the ids specified in the query param.
func (a RecordingApi) DeleteRecordingMediaretentionpolicies(ids string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/mediaretentionpolicies"
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'ids' is set
	if &ids == nil {
		// 
		return nil, errors.New("Missing required parameter 'ids' when calling RecordingApi->DeleteRecordingMediaretentionpolicies")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	if str, ok := interface{}(ids).(string); ok {
		if str != "" {
			queryParams["ids"] = a.Configuration.APIClient.ParameterToString(ids, collectionFormat)
		}
	} else {
		queryParams["ids"] = a.Configuration.APIClient.ParameterToString(ids, collectionFormat)
	}
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteRecordingMediaretentionpolicy invokes DELETE /api/v2/recording/mediaretentionpolicies/{policyId}
//
// Delete a media retention policy
//
// 
func (a RecordingApi) DeleteRecordingMediaretentionpolicy(policyId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/mediaretentionpolicies/{policyId}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'policyId' is set
	if &policyId == nil {
		// 
		return nil, errors.New("Missing required parameter 'policyId' when calling RecordingApi->DeleteRecordingMediaretentionpolicy")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// GetConversationRecording invokes GET /api/v2/conversations/{conversationId}/recordings/{recordingId}
//
// Gets a specific recording.
//
// 
func (a RecordingApi) GetConversationRecording(conversationId string, recordingId string, formatId string, emailFormatId string, chatFormatId string, messageFormatId string, download bool, fileName string, locale string) (*Recording, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/conversations/{conversationId}/recordings/{recordingId}"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{recordingId}", fmt.Sprintf("%v", recordingId), -1)
	defaultReturn := new(Recording)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling RecordingApi->GetConversationRecording")
	}
	// verify the required parameter 'recordingId' is set
	if &recordingId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'recordingId' when calling RecordingApi->GetConversationRecording")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	if str, ok := interface{}(formatId).(string); ok {
		if str != "" {
			queryParams["formatId"] = a.Configuration.APIClient.ParameterToString(formatId, collectionFormat)
		}
	} else {
		queryParams["formatId"] = a.Configuration.APIClient.ParameterToString(formatId, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(emailFormatId).(string); ok {
		if str != "" {
			queryParams["emailFormatId"] = a.Configuration.APIClient.ParameterToString(emailFormatId, collectionFormat)
		}
	} else {
		queryParams["emailFormatId"] = a.Configuration.APIClient.ParameterToString(emailFormatId, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(chatFormatId).(string); ok {
		if str != "" {
			queryParams["chatFormatId"] = a.Configuration.APIClient.ParameterToString(chatFormatId, collectionFormat)
		}
	} else {
		queryParams["chatFormatId"] = a.Configuration.APIClient.ParameterToString(chatFormatId, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(messageFormatId).(string); ok {
		if str != "" {
			queryParams["messageFormatId"] = a.Configuration.APIClient.ParameterToString(messageFormatId, collectionFormat)
		}
	} else {
		queryParams["messageFormatId"] = a.Configuration.APIClient.ParameterToString(messageFormatId, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(download).(string); ok {
		if str != "" {
			queryParams["download"] = a.Configuration.APIClient.ParameterToString(download, collectionFormat)
		}
	} else {
		queryParams["download"] = a.Configuration.APIClient.ParameterToString(download, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(fileName).(string); ok {
		if str != "" {
			queryParams["fileName"] = a.Configuration.APIClient.ParameterToString(fileName, collectionFormat)
		}
	} else {
		queryParams["fileName"] = a.Configuration.APIClient.ParameterToString(fileName, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(locale).(string); ok {
		if str != "" {
			queryParams["locale"] = a.Configuration.APIClient.ParameterToString(locale, collectionFormat)
		}
	} else {
		queryParams["locale"] = a.Configuration.APIClient.ParameterToString(locale, collectionFormat)
	}
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Recording
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetConversationRecordingAnnotation invokes GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}
//
// Get annotation
//
// 
func (a RecordingApi) GetConversationRecordingAnnotation(conversationId string, recordingId string, annotationId string) (*Annotation, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{recordingId}", fmt.Sprintf("%v", recordingId), -1)
	path = strings.Replace(path, "{annotationId}", fmt.Sprintf("%v", annotationId), -1)
	defaultReturn := new(Annotation)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling RecordingApi->GetConversationRecordingAnnotation")
	}
	// verify the required parameter 'recordingId' is set
	if &recordingId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'recordingId' when calling RecordingApi->GetConversationRecordingAnnotation")
	}
	// verify the required parameter 'annotationId' is set
	if &annotationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'annotationId' when calling RecordingApi->GetConversationRecordingAnnotation")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Annotation
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetConversationRecordingAnnotations invokes GET /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations
//
// Get annotations for recording
//
// 
func (a RecordingApi) GetConversationRecordingAnnotations(conversationId string, recordingId string) ([]Annotation, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{recordingId}", fmt.Sprintf("%v", recordingId), -1)
	defaultReturn := make([]Annotation, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling RecordingApi->GetConversationRecordingAnnotations")
	}
	// verify the required parameter 'recordingId' is set
	if &recordingId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'recordingId' when calling RecordingApi->GetConversationRecordingAnnotations")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload []Annotation
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetConversationRecordingmetadata invokes GET /api/v2/conversations/{conversationId}/recordingmetadata
//
// Get recording metadata for a conversation. Does not return playable media.
//
// 
func (a RecordingApi) GetConversationRecordingmetadata(conversationId string) ([]Recordingmetadata, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/conversations/{conversationId}/recordingmetadata"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	defaultReturn := make([]Recordingmetadata, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling RecordingApi->GetConversationRecordingmetadata")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload []Recordingmetadata
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetConversationRecordingmetadataRecordingId invokes GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}
//
// Get metadata for a specific recording. Does not return playable media.
//
// 
func (a RecordingApi) GetConversationRecordingmetadataRecordingId(conversationId string, recordingId string) (*Recordingmetadata, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{recordingId}", fmt.Sprintf("%v", recordingId), -1)
	defaultReturn := new(Recordingmetadata)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling RecordingApi->GetConversationRecordingmetadataRecordingId")
	}
	// verify the required parameter 'recordingId' is set
	if &recordingId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'recordingId' when calling RecordingApi->GetConversationRecordingmetadataRecordingId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Recordingmetadata
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetConversationRecordings invokes GET /api/v2/conversations/{conversationId}/recordings
//
// Get all of a Conversation&#39;s Recordings.
//
// 
func (a RecordingApi) GetConversationRecordings(conversationId string, maxWaitMs int32, formatId string) ([]Recording, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/conversations/{conversationId}/recordings"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	defaultReturn := make([]Recording, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling RecordingApi->GetConversationRecordings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	if str, ok := interface{}(maxWaitMs).(string); ok {
		if str != "" {
			queryParams["maxWaitMs"] = a.Configuration.APIClient.ParameterToString(maxWaitMs, collectionFormat)
		}
	} else {
		queryParams["maxWaitMs"] = a.Configuration.APIClient.ParameterToString(maxWaitMs, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(formatId).(string); ok {
		if str != "" {
			queryParams["formatId"] = a.Configuration.APIClient.ParameterToString(formatId, collectionFormat)
		}
	} else {
		queryParams["formatId"] = a.Configuration.APIClient.ParameterToString(formatId, collectionFormat)
	}
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload []Recording
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetOrphanrecording invokes GET /api/v2/orphanrecordings/{orphanId}
//
// Gets a single orphan recording
//
// 
func (a RecordingApi) GetOrphanrecording(orphanId string) (*Orphanrecording, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/orphanrecordings/{orphanId}"
	path = strings.Replace(path, "{orphanId}", fmt.Sprintf("%v", orphanId), -1)
	defaultReturn := new(Orphanrecording)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'orphanId' is set
	if &orphanId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'orphanId' when calling RecordingApi->GetOrphanrecording")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Orphanrecording
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetOrphanrecordingMedia invokes GET /api/v2/orphanrecordings/{orphanId}/media
//
// Gets the media of a single orphan recording
//
// A 202 response means the orphaned media is currently transcoding and will be available shortly.A 200 response denotes the transcoded orphan media is available now and is contained in the response body.
func (a RecordingApi) GetOrphanrecordingMedia(orphanId string, formatId string, emailFormatId string, chatFormatId string, messageFormatId string, download bool, fileName string, locale string) (*Recording, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/orphanrecordings/{orphanId}/media"
	path = strings.Replace(path, "{orphanId}", fmt.Sprintf("%v", orphanId), -1)
	defaultReturn := new(Recording)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'orphanId' is set
	if &orphanId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'orphanId' when calling RecordingApi->GetOrphanrecordingMedia")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	if str, ok := interface{}(formatId).(string); ok {
		if str != "" {
			queryParams["formatId"] = a.Configuration.APIClient.ParameterToString(formatId, collectionFormat)
		}
	} else {
		queryParams["formatId"] = a.Configuration.APIClient.ParameterToString(formatId, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(emailFormatId).(string); ok {
		if str != "" {
			queryParams["emailFormatId"] = a.Configuration.APIClient.ParameterToString(emailFormatId, collectionFormat)
		}
	} else {
		queryParams["emailFormatId"] = a.Configuration.APIClient.ParameterToString(emailFormatId, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(chatFormatId).(string); ok {
		if str != "" {
			queryParams["chatFormatId"] = a.Configuration.APIClient.ParameterToString(chatFormatId, collectionFormat)
		}
	} else {
		queryParams["chatFormatId"] = a.Configuration.APIClient.ParameterToString(chatFormatId, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(messageFormatId).(string); ok {
		if str != "" {
			queryParams["messageFormatId"] = a.Configuration.APIClient.ParameterToString(messageFormatId, collectionFormat)
		}
	} else {
		queryParams["messageFormatId"] = a.Configuration.APIClient.ParameterToString(messageFormatId, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(download).(string); ok {
		if str != "" {
			queryParams["download"] = a.Configuration.APIClient.ParameterToString(download, collectionFormat)
		}
	} else {
		queryParams["download"] = a.Configuration.APIClient.ParameterToString(download, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(fileName).(string); ok {
		if str != "" {
			queryParams["fileName"] = a.Configuration.APIClient.ParameterToString(fileName, collectionFormat)
		}
	} else {
		queryParams["fileName"] = a.Configuration.APIClient.ParameterToString(fileName, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(locale).(string); ok {
		if str != "" {
			queryParams["locale"] = a.Configuration.APIClient.ParameterToString(locale, collectionFormat)
		}
	} else {
		queryParams["locale"] = a.Configuration.APIClient.ParameterToString(locale, collectionFormat)
	}
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Recording
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetOrphanrecordings invokes GET /api/v2/orphanrecordings
//
// Gets all orphan recordings
//
// 
func (a RecordingApi) GetOrphanrecordings(pageSize int32, pageNumber int32, sortBy string, expand []string, nextPage string, previousPage string, hasConversation bool, media string) (*Orphanrecordinglisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/orphanrecordings"
	defaultReturn := new(Orphanrecordinglisting)
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
	if str, ok := interface{}(hasConversation).(string); ok {
		if str != "" {
			queryParams["hasConversation"] = a.Configuration.APIClient.ParameterToString(hasConversation, collectionFormat)
		}
	} else {
		queryParams["hasConversation"] = a.Configuration.APIClient.ParameterToString(hasConversation, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(media).(string); ok {
		if str != "" {
			queryParams["media"] = a.Configuration.APIClient.ParameterToString(media, collectionFormat)
		}
	} else {
		queryParams["media"] = a.Configuration.APIClient.ParameterToString(media, collectionFormat)
	}
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Orphanrecordinglisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetRecordingBatchrequest invokes GET /api/v2/recording/batchrequests/{jobId}
//
// Get the status and results for a batch request job, only the user that submitted the job may retrieve results
//
// 
func (a RecordingApi) GetRecordingBatchrequest(jobId string) (*Batchdownloadjobstatusresult, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/batchrequests/{jobId}"
	path = strings.Replace(path, "{jobId}", fmt.Sprintf("%v", jobId), -1)
	defaultReturn := new(Batchdownloadjobstatusresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling RecordingApi->GetRecordingBatchrequest")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Batchdownloadjobstatusresult
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetRecordingJob invokes GET /api/v2/recording/jobs/{jobId}
//
// Get the status of the job associated with the job id.
//
// 
func (a RecordingApi) GetRecordingJob(jobId string) (*Recordingjob, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", fmt.Sprintf("%v", jobId), -1)
	defaultReturn := new(Recordingjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling RecordingApi->GetRecordingJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Recordingjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetRecordingJobs invokes GET /api/v2/recording/jobs
//
// Get the status of all jobs within the user&#39;s organization
//
// 
func (a RecordingApi) GetRecordingJobs(pageSize int32, pageNumber int32, sortBy string, state string, showOnlyMyJobs bool, jobType string) (*Recordingjobentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/jobs"
	defaultReturn := new(Recordingjobentitylisting)
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
	if str, ok := interface{}(state).(string); ok {
		if str != "" {
			queryParams["state"] = a.Configuration.APIClient.ParameterToString(state, collectionFormat)
		}
	} else {
		queryParams["state"] = a.Configuration.APIClient.ParameterToString(state, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(showOnlyMyJobs).(string); ok {
		if str != "" {
			queryParams["showOnlyMyJobs"] = a.Configuration.APIClient.ParameterToString(showOnlyMyJobs, collectionFormat)
		}
	} else {
		queryParams["showOnlyMyJobs"] = a.Configuration.APIClient.ParameterToString(showOnlyMyJobs, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(jobType).(string); ok {
		if str != "" {
			queryParams["jobType"] = a.Configuration.APIClient.ParameterToString(jobType, collectionFormat)
		}
	} else {
		queryParams["jobType"] = a.Configuration.APIClient.ParameterToString(jobType, collectionFormat)
	}
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Recordingjobentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetRecordingLocalkeysSetting invokes GET /api/v2/recording/localkeys/settings/{settingsId}
//
// Get the local encryption settings
//
// 
func (a RecordingApi) GetRecordingLocalkeysSetting(settingsId string) (*Localencryptionconfiguration, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/localkeys/settings/{settingsId}"
	path = strings.Replace(path, "{settingsId}", fmt.Sprintf("%v", settingsId), -1)
	defaultReturn := new(Localencryptionconfiguration)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'settingsId' is set
	if &settingsId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'settingsId' when calling RecordingApi->GetRecordingLocalkeysSetting")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Localencryptionconfiguration
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetRecordingLocalkeysSettings invokes GET /api/v2/recording/localkeys/settings
//
// gets a list local key settings data
//
// 
func (a RecordingApi) GetRecordingLocalkeysSettings() (*Localencryptionconfigurationlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/localkeys/settings"
	defaultReturn := new(Localencryptionconfigurationlisting)
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
	var successPayload *Localencryptionconfigurationlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetRecordingMediaretentionpolicies invokes GET /api/v2/recording/mediaretentionpolicies
//
// Gets media retention policy list with query options to filter on name and enabled.
//
// for a less verbose response, add summary=true to this endpoint
func (a RecordingApi) GetRecordingMediaretentionpolicies(pageSize int32, pageNumber int32, sortBy string, expand []string, nextPage string, previousPage string, name string, enabled bool, summary bool, hasErrors bool) (*Policyentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/mediaretentionpolicies"
	defaultReturn := new(Policyentitylisting)
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
	if str, ok := interface{}(name).(string); ok {
		if str != "" {
			queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
		}
	} else {
		queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(enabled).(string); ok {
		if str != "" {
			queryParams["enabled"] = a.Configuration.APIClient.ParameterToString(enabled, collectionFormat)
		}
	} else {
		queryParams["enabled"] = a.Configuration.APIClient.ParameterToString(enabled, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(summary).(string); ok {
		if str != "" {
			queryParams["summary"] = a.Configuration.APIClient.ParameterToString(summary, collectionFormat)
		}
	} else {
		queryParams["summary"] = a.Configuration.APIClient.ParameterToString(summary, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(hasErrors).(string); ok {
		if str != "" {
			queryParams["hasErrors"] = a.Configuration.APIClient.ParameterToString(hasErrors, collectionFormat)
		}
	} else {
		queryParams["hasErrors"] = a.Configuration.APIClient.ParameterToString(hasErrors, collectionFormat)
	}
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Policyentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetRecordingMediaretentionpolicy invokes GET /api/v2/recording/mediaretentionpolicies/{policyId}
//
// Get a media retention policy
//
// 
func (a RecordingApi) GetRecordingMediaretentionpolicy(policyId string) (*Policy, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/mediaretentionpolicies/{policyId}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyId), -1)
	defaultReturn := new(Policy)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'policyId' is set
	if &policyId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'policyId' when calling RecordingApi->GetRecordingMediaretentionpolicy")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Policy
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetRecordingRecordingkeys invokes GET /api/v2/recording/recordingkeys
//
// Get encryption key list
//
// 
func (a RecordingApi) GetRecordingRecordingkeys(pageSize int32, pageNumber int32) (*Encryptionkeyentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/recordingkeys"
	defaultReturn := new(Encryptionkeyentitylisting)
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
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Encryptionkeyentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetRecordingRecordingkeysRotationschedule invokes GET /api/v2/recording/recordingkeys/rotationschedule
//
// Get key rotation schedule
//
// 
func (a RecordingApi) GetRecordingRecordingkeysRotationschedule() (*Keyrotationschedule, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/recordingkeys/rotationschedule"
	defaultReturn := new(Keyrotationschedule)
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
	var successPayload *Keyrotationschedule
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetRecordingSettings invokes GET /api/v2/recording/settings
//
// Get the Recording Settings for the Organization
//
// 
func (a RecordingApi) GetRecordingSettings(createDefault bool) (*Recordingsettings, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/settings"
	defaultReturn := new(Recordingsettings)
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
	if str, ok := interface{}(createDefault).(string); ok {
		if str != "" {
			queryParams["createDefault"] = a.Configuration.APIClient.ParameterToString(createDefault, collectionFormat)
		}
	} else {
		queryParams["createDefault"] = a.Configuration.APIClient.ParameterToString(createDefault, collectionFormat)
	}
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Recordingsettings
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetRecordingsScreensessions invokes GET /api/v2/recordings/screensessions
//
// Retrieves a paged listing of screen recording sessions
//
// 
func (a RecordingApi) GetRecordingsScreensessions(pageSize int32, pageNumber int32) (*Screenrecordingsessionlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recordings/screensessions"
	defaultReturn := new(Screenrecordingsessionlisting)
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
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Screenrecordingsessionlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PatchRecordingMediaretentionpolicy invokes PATCH /api/v2/recording/mediaretentionpolicies/{policyId}
//
// Patch a media retention policy
//
// 
func (a RecordingApi) PatchRecordingMediaretentionpolicy(policyId string, body Policy) (*Policy, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/mediaretentionpolicies/{policyId}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyId), -1)
	defaultReturn := new(Policy)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'policyId' is set
	if &policyId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'policyId' when calling RecordingApi->PatchRecordingMediaretentionpolicy")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PatchRecordingMediaretentionpolicy")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Policy
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PatchRecordingsScreensession invokes PATCH /api/v2/recordings/screensessions/{recordingSessionId}
//
// Update a screen recording session
//
// 
func (a RecordingApi) PatchRecordingsScreensession(recordingSessionId string, body Screenrecordingsessionrequest) (*APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recordings/screensessions/{recordingSessionId}"
	path = strings.Replace(path, "{recordingSessionId}", fmt.Sprintf("%v", recordingSessionId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'recordingSessionId' is set
	if &recordingSessionId == nil {
		// 
		return nil, errors.New("Missing required parameter 'recordingSessionId' when calling RecordingApi->PatchRecordingsScreensession")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// PostConversationRecordingAnnotations invokes POST /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations
//
// Create annotation
//
// 
func (a RecordingApi) PostConversationRecordingAnnotations(conversationId string, recordingId string, body Annotation) (*Annotation, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{recordingId}", fmt.Sprintf("%v", recordingId), -1)
	defaultReturn := new(Annotation)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling RecordingApi->PostConversationRecordingAnnotations")
	}
	// verify the required parameter 'recordingId' is set
	if &recordingId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'recordingId' when calling RecordingApi->PostConversationRecordingAnnotations")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PostConversationRecordingAnnotations")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Annotation
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostRecordingBatchrequests invokes POST /api/v2/recording/batchrequests
//
// Submit a batch download request for recordings. Recordings in response will be in their original format/codec - configured in the Trunk configuration.
//
// 
func (a RecordingApi) PostRecordingBatchrequests(body Batchdownloadjobsubmission) (*Batchdownloadjobsubmissionresult, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/batchrequests"
	defaultReturn := new(Batchdownloadjobsubmissionresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PostRecordingBatchrequests")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Batchdownloadjobsubmissionresult
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostRecordingJobs invokes POST /api/v2/recording/jobs
//
// Create a recording bulk job
//
// 
func (a RecordingApi) PostRecordingJobs(body Recordingjobsquery) (*Recordingjob, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/jobs"
	defaultReturn := new(Recordingjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PostRecordingJobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Recordingjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostRecordingLocalkeys invokes POST /api/v2/recording/localkeys
//
// create a local recording key
//
// 
func (a RecordingApi) PostRecordingLocalkeys(body Localencryptionkeyrequest) (*Encryptionkey, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/localkeys"
	defaultReturn := new(Encryptionkey)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PostRecordingLocalkeys")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Encryptionkey
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostRecordingLocalkeysSettings invokes POST /api/v2/recording/localkeys/settings
//
// create settings for local key creation
//
// 
func (a RecordingApi) PostRecordingLocalkeysSettings(body Localencryptionconfiguration) (*Localencryptionconfiguration, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/localkeys/settings"
	defaultReturn := new(Localencryptionconfiguration)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PostRecordingLocalkeysSettings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Localencryptionconfiguration
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostRecordingMediaretentionpolicies invokes POST /api/v2/recording/mediaretentionpolicies
//
// Create media retention policy
//
// 
func (a RecordingApi) PostRecordingMediaretentionpolicies(body Policycreate) (*Policy, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/mediaretentionpolicies"
	defaultReturn := new(Policy)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PostRecordingMediaretentionpolicies")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Policy
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostRecordingRecordingkeys invokes POST /api/v2/recording/recordingkeys
//
// Create encryption key
//
// 
func (a RecordingApi) PostRecordingRecordingkeys() (*Encryptionkey, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/recordingkeys"
	defaultReturn := new(Encryptionkey)
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
	var successPayload *Encryptionkey
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostRecordingsDeletionprotection invokes POST /api/v2/recordings/deletionprotection
//
// Get a list of conversations with protected recordings
//
// 
func (a RecordingApi) PostRecordingsDeletionprotection(body Conversationdeletionprotectionquery) ([]Addressableentityref, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recordings/deletionprotection"
	defaultReturn := make([]Addressableentityref, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PostRecordingsDeletionprotection")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload []Addressableentityref
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PutConversationRecording invokes PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}
//
// Updates the retention records on a recording.
//
// Currently supports updating and removing both archive and delete dates for eligible recordings. A request to change the archival date of an archived recording will result in a restoration of the recording until the new date set. The recording:recording:view permission is required for the recording, as well as either the recording:recording:editRetention or recording:screenRecording:editRetention permissions depending on the type of recording.
func (a RecordingApi) PutConversationRecording(conversationId string, recordingId string, body Recording) (*Recording, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/conversations/{conversationId}/recordings/{recordingId}"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{recordingId}", fmt.Sprintf("%v", recordingId), -1)
	defaultReturn := new(Recording)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling RecordingApi->PutConversationRecording")
	}
	// verify the required parameter 'recordingId' is set
	if &recordingId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'recordingId' when calling RecordingApi->PutConversationRecording")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PutConversationRecording")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Recording
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PutConversationRecordingAnnotation invokes PUT /api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}
//
// Update annotation
//
// 
func (a RecordingApi) PutConversationRecordingAnnotation(conversationId string, recordingId string, annotationId string, body Annotation) (*Annotation, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/conversations/{conversationId}/recordings/{recordingId}/annotations/{annotationId}"
	path = strings.Replace(path, "{conversationId}", fmt.Sprintf("%v", conversationId), -1)
	path = strings.Replace(path, "{recordingId}", fmt.Sprintf("%v", recordingId), -1)
	path = strings.Replace(path, "{annotationId}", fmt.Sprintf("%v", annotationId), -1)
	defaultReturn := new(Annotation)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling RecordingApi->PutConversationRecordingAnnotation")
	}
	// verify the required parameter 'recordingId' is set
	if &recordingId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'recordingId' when calling RecordingApi->PutConversationRecordingAnnotation")
	}
	// verify the required parameter 'annotationId' is set
	if &annotationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'annotationId' when calling RecordingApi->PutConversationRecordingAnnotation")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PutConversationRecordingAnnotation")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Annotation
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PutOrphanrecording invokes PUT /api/v2/orphanrecordings/{orphanId}
//
// Updates an orphan recording to a regular recording with retention values
//
// If this operation is successful the orphan will no longer exist. It will be replaced by the resulting recording in the response. This replacement recording is accessible by the normal Recording api.
func (a RecordingApi) PutOrphanrecording(orphanId string, body Orphanupdaterequest) (*Recording, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/orphanrecordings/{orphanId}"
	path = strings.Replace(path, "{orphanId}", fmt.Sprintf("%v", orphanId), -1)
	defaultReturn := new(Recording)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'orphanId' is set
	if &orphanId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'orphanId' when calling RecordingApi->PutOrphanrecording")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Recording
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PutRecordingJob invokes PUT /api/v2/recording/jobs/{jobId}
//
// Execute the recording bulk job.
//
// A job must be executed by the same user whom originally created the job.  In addition, the user must have permission to update the recording&#39;s retention.
func (a RecordingApi) PutRecordingJob(jobId string, body Executerecordingjobsquery) (*Recordingjob, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", fmt.Sprintf("%v", jobId), -1)
	defaultReturn := new(Recordingjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling RecordingApi->PutRecordingJob")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PutRecordingJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Recordingjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PutRecordingLocalkeysSetting invokes PUT /api/v2/recording/localkeys/settings/{settingsId}
//
// Update the local encryption settings
//
// 
func (a RecordingApi) PutRecordingLocalkeysSetting(settingsId string, body Localencryptionconfiguration) (*Localencryptionconfiguration, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/localkeys/settings/{settingsId}"
	path = strings.Replace(path, "{settingsId}", fmt.Sprintf("%v", settingsId), -1)
	defaultReturn := new(Localencryptionconfiguration)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'settingsId' is set
	if &settingsId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'settingsId' when calling RecordingApi->PutRecordingLocalkeysSetting")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PutRecordingLocalkeysSetting")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Localencryptionconfiguration
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PutRecordingMediaretentionpolicy invokes PUT /api/v2/recording/mediaretentionpolicies/{policyId}
//
// Update a media retention policy
//
// 
func (a RecordingApi) PutRecordingMediaretentionpolicy(policyId string, body Policy) (*Policy, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/mediaretentionpolicies/{policyId}"
	path = strings.Replace(path, "{policyId}", fmt.Sprintf("%v", policyId), -1)
	defaultReturn := new(Policy)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'policyId' is set
	if &policyId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'policyId' when calling RecordingApi->PutRecordingMediaretentionpolicy")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PutRecordingMediaretentionpolicy")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Policy
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PutRecordingRecordingkeysRotationschedule invokes PUT /api/v2/recording/recordingkeys/rotationschedule
//
// Update key rotation schedule
//
// 
func (a RecordingApi) PutRecordingRecordingkeysRotationschedule(body Keyrotationschedule) (*Keyrotationschedule, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/recordingkeys/rotationschedule"
	defaultReturn := new(Keyrotationschedule)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PutRecordingRecordingkeysRotationschedule")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Keyrotationschedule
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PutRecordingSettings invokes PUT /api/v2/recording/settings
//
// Update the Recording Settings for the Organization
//
// 
func (a RecordingApi) PutRecordingSettings(body Recordingsettings) (*Recordingsettings, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recording/settings"
	defaultReturn := new(Recordingsettings)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling RecordingApi->PutRecordingSettings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Recordingsettings
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PutRecordingsDeletionprotection invokes PUT /api/v2/recordings/deletionprotection
//
// Apply or revoke recording protection for conversations
//
// 
func (a RecordingApi) PutRecordingsDeletionprotection(protect bool, body Conversationdeletionprotectionquery) (*APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/recordings/deletionprotection"
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
	
	var collectionFormat string
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(protect).(string); ok {
		if str != "" {
			queryParams["protect"] = a.Configuration.APIClient.ParameterToString(protect, collectionFormat)
		}
	} else {
		queryParams["protect"] = a.Configuration.APIClient.ParameterToString(protect, collectionFormat)
	}
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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

