package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
	)

// SpeechTextAnalyticsApi provides functions for API endpoints
type SpeechTextAnalyticsApi struct {
	Configuration *Configuration
}

// NewSpeechTextAnalyticsApi creates an API instance using the default configuration
func NewSpeechTextAnalyticsApi() *SpeechTextAnalyticsApi {
	fmt.Sprintf(strings.Title(""), "")
	config := GetDefaultConfiguration()
	return &SpeechTextAnalyticsApi{
		Configuration: config,
	}
}

// NewSpeechTextAnalyticsApiWithConfig creates an API instance using the provided configuration
func NewSpeechTextAnalyticsApiWithConfig(config *Configuration) *SpeechTextAnalyticsApi {
	return &SpeechTextAnalyticsApi{
		Configuration: config,
	}
}

// DeleteSpeechandtextanalyticsCategory invokes DELETE /api/v2/speechandtextanalytics/categories/{categoryId}
//
// Delete a Speech & Text Analytics category by ID
func (a SpeechTextAnalyticsApi) DeleteSpeechandtextanalyticsCategory(categoryId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/categories/{categoryId}"
	path = strings.Replace(path, "{categoryId}", url.PathEscape(fmt.Sprintf("%v", categoryId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'categoryId' is set
	if &categoryId == nil {
		// false
		return nil, errors.New("Missing required parameter 'categoryId' when calling SpeechTextAnalyticsApi->DeleteSpeechandtextanalyticsCategory")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteSpeechandtextanalyticsDictionaryfeedbackDictionaryFeedbackId invokes DELETE /api/v2/speechandtextanalytics/dictionaryfeedback/{dictionaryFeedbackId}
//
// Delete a Speech & Text Analytics DictionaryFeedback by Id
func (a SpeechTextAnalyticsApi) DeleteSpeechandtextanalyticsDictionaryfeedbackDictionaryFeedbackId(dictionaryFeedbackId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/dictionaryfeedback/{dictionaryFeedbackId}"
	path = strings.Replace(path, "{dictionaryFeedbackId}", url.PathEscape(fmt.Sprintf("%v", dictionaryFeedbackId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'dictionaryFeedbackId' is set
	if &dictionaryFeedbackId == nil {
		// false
		return nil, errors.New("Missing required parameter 'dictionaryFeedbackId' when calling SpeechTextAnalyticsApi->DeleteSpeechandtextanalyticsDictionaryfeedbackDictionaryFeedbackId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteSpeechandtextanalyticsProgram invokes DELETE /api/v2/speechandtextanalytics/programs/{programId}
//
// Delete a Speech & Text Analytics program by id
func (a SpeechTextAnalyticsApi) DeleteSpeechandtextanalyticsProgram(programId string, forceDelete bool) (*Deleteprogramresponse, *APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/{programId}"
	path = strings.Replace(path, "{programId}", url.PathEscape(fmt.Sprintf("%v", programId)), -1)
	defaultReturn := new(Deleteprogramresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'programId' is set
	if &programId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'programId' when calling SpeechTextAnalyticsApi->DeleteSpeechandtextanalyticsProgram")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["forceDelete"] = a.Configuration.APIClient.ParameterToString(forceDelete, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Deleteprogramresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Deleteprogramresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// DeleteSpeechandtextanalyticsReprocessingJob invokes DELETE /api/v2/speechandtextanalytics/reprocessing/jobs/{jobId}
//
// Delete a Speech & Text Analytics Reprocessing job by Id
//
// Preview: DeleteSpeechandtextanalyticsReprocessingJob is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SpeechTextAnalyticsApi) DeleteSpeechandtextanalyticsReprocessingJob(jobId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/reprocessing/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return nil, errors.New("Missing required parameter 'jobId' when calling SpeechTextAnalyticsApi->DeleteSpeechandtextanalyticsReprocessingJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteSpeechandtextanalyticsSentimentfeedback invokes DELETE /api/v2/speechandtextanalytics/sentimentfeedback
//
// Delete All Speech & Text Analytics SentimentFeedback
func (a SpeechTextAnalyticsApi) DeleteSpeechandtextanalyticsSentimentfeedback() (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/sentimentfeedback"
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

// DeleteSpeechandtextanalyticsSentimentfeedbackSentimentFeedbackId invokes DELETE /api/v2/speechandtextanalytics/sentimentfeedback/{sentimentFeedbackId}
//
// Delete a Speech & Text Analytics SentimentFeedback by Id
func (a SpeechTextAnalyticsApi) DeleteSpeechandtextanalyticsSentimentfeedbackSentimentFeedbackId(sentimentFeedbackId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/sentimentfeedback/{sentimentFeedbackId}"
	path = strings.Replace(path, "{sentimentFeedbackId}", url.PathEscape(fmt.Sprintf("%v", sentimentFeedbackId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sentimentFeedbackId' is set
	if &sentimentFeedbackId == nil {
		// false
		return nil, errors.New("Missing required parameter 'sentimentFeedbackId' when calling SpeechTextAnalyticsApi->DeleteSpeechandtextanalyticsSentimentfeedbackSentimentFeedbackId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteSpeechandtextanalyticsTopic invokes DELETE /api/v2/speechandtextanalytics/topics/{topicId}
//
// Delete a Speech & Text Analytics topic by id
func (a SpeechTextAnalyticsApi) DeleteSpeechandtextanalyticsTopic(topicId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/topics/{topicId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return nil, errors.New("Missing required parameter 'topicId' when calling SpeechTextAnalyticsApi->DeleteSpeechandtextanalyticsTopic")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// GetSpeechandtextanalyticsCategories invokes GET /api/v2/speechandtextanalytics/categories
//
// Get the list of Speech and Text Analytics categories
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsCategories(pageSize int, pageNumber int, name string, sortOrder string, sortBy string, ids []string) (*Categoriesentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/categories"
	defaultReturn := new(Categoriesentitylisting)
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
	
	queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["ids"] = a.Configuration.APIClient.ParameterToString(ids, "multi")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Categoriesentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Categoriesentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsCategory invokes GET /api/v2/speechandtextanalytics/categories/{categoryId}
//
// Get a Speech & Text Analytics Category by ID
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsCategory(categoryId string) (*Stacategory, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/categories/{categoryId}"
	path = strings.Replace(path, "{categoryId}", url.PathEscape(fmt.Sprintf("%v", categoryId)), -1)
	defaultReturn := new(Stacategory)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'categoryId' is set
	if &categoryId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'categoryId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsCategory")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Stacategory
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Stacategory" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsConversation invokes GET /api/v2/speechandtextanalytics/conversations/{conversationId}
//
// Get Speech and Text Analytics for a specific conversation
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsConversation(conversationId string) (*Conversationmetrics, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/conversations/{conversationId}"
	path = strings.Replace(path, "{conversationId}", url.PathEscape(fmt.Sprintf("%v", conversationId)), -1)
	defaultReturn := new(Conversationmetrics)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsConversation")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Conversationmetrics
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Conversationmetrics" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsConversationCategories invokes GET /api/v2/speechandtextanalytics/conversations/{conversationId}/categories
//
// Get the list of detected Speech and Text Analytics categories of conversation
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsConversationCategories(conversationId string, pageSize int, pageNumber int) (*Conversationcategoriesentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/conversations/{conversationId}/categories"
	path = strings.Replace(path, "{conversationId}", url.PathEscape(fmt.Sprintf("%v", conversationId)), -1)
	defaultReturn := new(Conversationcategoriesentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsConversationCategories")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Conversationcategoriesentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Conversationcategoriesentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsConversationCommunicationTranscripturl invokes GET /api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl
//
// Get the pre-signed S3 URL for the transcript of a specific communication of a conversation
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsConversationCommunicationTranscripturl(conversationId string, communicationId string) (*Transcripturl, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl"
	path = strings.Replace(path, "{conversationId}", url.PathEscape(fmt.Sprintf("%v", conversationId)), -1)
	path = strings.Replace(path, "{communicationId}", url.PathEscape(fmt.Sprintf("%v", communicationId)), -1)
	defaultReturn := new(Transcripturl)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsConversationCommunicationTranscripturl")
	}
	// verify the required parameter 'communicationId' is set
	if &communicationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'communicationId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsConversationCommunicationTranscripturl")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Transcripturl
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Transcripturl" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsConversationCommunicationTranscripturls invokes GET /api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturls
//
// Get the list of pre-signed S3 URL for the transcripts of a specific communication of a conversation
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsConversationCommunicationTranscripturls(conversationId string, communicationId string) (*Transcripturls, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturls"
	path = strings.Replace(path, "{conversationId}", url.PathEscape(fmt.Sprintf("%v", conversationId)), -1)
	path = strings.Replace(path, "{communicationId}", url.PathEscape(fmt.Sprintf("%v", communicationId)), -1)
	defaultReturn := new(Transcripturls)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsConversationCommunicationTranscripturls")
	}
	// verify the required parameter 'communicationId' is set
	if &communicationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'communicationId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsConversationCommunicationTranscripturls")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Transcripturls
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Transcripturls" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsConversationSentiments invokes GET /api/v2/speechandtextanalytics/conversations/{conversationId}/sentiments
//
// Get sentiment data
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsConversationSentiments(conversationId string) (*Sentimentdata, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/conversations/{conversationId}/sentiments"
	path = strings.Replace(path, "{conversationId}", url.PathEscape(fmt.Sprintf("%v", conversationId)), -1)
	defaultReturn := new(Sentimentdata)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsConversationSentiments")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Sentimentdata
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Sentimentdata" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsConversationSummaries invokes GET /api/v2/speechandtextanalytics/conversations/{conversationId}/summaries
//
// Get conversation summaries by conversation id.
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsConversationSummaries(conversationId string) (*Speechtextanalyticsconversationsummarylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/conversations/{conversationId}/summaries"
	path = strings.Replace(path, "{conversationId}", url.PathEscape(fmt.Sprintf("%v", conversationId)), -1)
	defaultReturn := new(Speechtextanalyticsconversationsummarylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsConversationSummaries")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Speechtextanalyticsconversationsummarylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Speechtextanalyticsconversationsummarylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsDictionaryfeedback invokes GET /api/v2/speechandtextanalytics/dictionaryfeedback
//
// Get the list of Speech & Text Analytics dictionary feedbacks
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsDictionaryfeedback(dialect string, nextPage string, pageSize int) (*Dictionaryfeedbackentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/dictionaryfeedback"
	defaultReturn := new(Dictionaryfeedbackentitylisting)
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
	
	queryParams["dialect"] = a.Configuration.APIClient.ParameterToString(dialect, "")
	
	queryParams["nextPage"] = a.Configuration.APIClient.ParameterToString(nextPage, "")
	
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
	var successPayload *Dictionaryfeedbackentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dictionaryfeedbackentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsDictionaryfeedbackDictionaryFeedbackId invokes GET /api/v2/speechandtextanalytics/dictionaryfeedback/{dictionaryFeedbackId}
//
// Get a Speech & Text Analytics dictionary feedback by id
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsDictionaryfeedbackDictionaryFeedbackId(dictionaryFeedbackId string) (*Dictionaryfeedback, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/dictionaryfeedback/{dictionaryFeedbackId}"
	path = strings.Replace(path, "{dictionaryFeedbackId}", url.PathEscape(fmt.Sprintf("%v", dictionaryFeedbackId)), -1)
	defaultReturn := new(Dictionaryfeedback)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'dictionaryFeedbackId' is set
	if &dictionaryFeedbackId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'dictionaryFeedbackId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsDictionaryfeedbackDictionaryFeedbackId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Dictionaryfeedback
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dictionaryfeedback" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsProgram invokes GET /api/v2/speechandtextanalytics/programs/{programId}
//
// Get a Speech & Text Analytics program by id
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsProgram(programId string) (*Program, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/{programId}"
	path = strings.Replace(path, "{programId}", url.PathEscape(fmt.Sprintf("%v", programId)), -1)
	defaultReturn := new(Program)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'programId' is set
	if &programId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'programId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsProgram")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Program
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Program" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsProgramMappings invokes GET /api/v2/speechandtextanalytics/programs/{programId}/mappings
//
// Get Speech & Text Analytics program mappings to queues and flows by id
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsProgramMappings(programId string) (*Programmappings, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/{programId}/mappings"
	path = strings.Replace(path, "{programId}", url.PathEscape(fmt.Sprintf("%v", programId)), -1)
	defaultReturn := new(Programmappings)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'programId' is set
	if &programId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'programId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsProgramMappings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Programmappings
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Programmappings" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsProgramSettingsInsights invokes GET /api/v2/speechandtextanalytics/programs/{programId}/settings/insights
//
// Get AI Insights settings of a program
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsProgramSettingsInsights(programId string) (*Programinsightssettings, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/{programId}/settings/insights"
	path = strings.Replace(path, "{programId}", url.PathEscape(fmt.Sprintf("%v", programId)), -1)
	defaultReturn := new(Programinsightssettings)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'programId' is set
	if &programId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'programId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsProgramSettingsInsights")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Programinsightssettings
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Programinsightssettings" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsProgramTranscriptionengines invokes GET /api/v2/speechandtextanalytics/programs/{programId}/transcriptionengines
//
// Get transcription engine settings of a program
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsProgramTranscriptionengines(programId string) (*Programtranscriptionengines, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/{programId}/transcriptionengines"
	path = strings.Replace(path, "{programId}", url.PathEscape(fmt.Sprintf("%v", programId)), -1)
	defaultReturn := new(Programtranscriptionengines)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'programId' is set
	if &programId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'programId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsProgramTranscriptionengines")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Programtranscriptionengines
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Programtranscriptionengines" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsPrograms invokes GET /api/v2/speechandtextanalytics/programs
//
// Get the list of Speech & Text Analytics programs
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsPrograms(nextPage string, pageSize int, state string) (*Programsentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs"
	defaultReturn := new(Programsentitylisting)
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
	
	queryParams["nextPage"] = a.Configuration.APIClient.ParameterToString(nextPage, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
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
	var successPayload *Programsentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Programsentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsProgramsGeneralJob invokes GET /api/v2/speechandtextanalytics/programs/general/jobs/{jobId}
//
// Get a Speech & Text Analytics general program job by id
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsProgramsGeneralJob(jobId string) (*Generalprogramjob, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/general/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Generalprogramjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsProgramsGeneralJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Generalprogramjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Generalprogramjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsProgramsMappings invokes GET /api/v2/speechandtextanalytics/programs/mappings
//
// Get the list of Speech & Text Analytics programs mappings to queues and flows
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsProgramsMappings(nextPage string, pageSize int) (*Programsmappingsentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/mappings"
	defaultReturn := new(Programsmappingsentitylisting)
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
	
	queryParams["nextPage"] = a.Configuration.APIClient.ParameterToString(nextPage, "")
	
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
	var successPayload *Programsmappingsentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Programsmappingsentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsProgramsPublishjob invokes GET /api/v2/speechandtextanalytics/programs/publishjobs/{jobId}
//
// Get a Speech & Text Analytics publish programs job by id
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsProgramsPublishjob(jobId string) (*Programjob, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/publishjobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Programjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsProgramsPublishjob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Programjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Programjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsProgramsSettingsInsights invokes GET /api/v2/speechandtextanalytics/programs/settings/insights
//
// Get the list of program AI Insights settings for the organization
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsProgramsSettingsInsights(pageSize int, pageNumber int, programIds []string) (*Programinsightssettingsentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/settings/insights"
	defaultReturn := new(Programinsightssettingsentitylisting)
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
	
	queryParams["programIds"] = a.Configuration.APIClient.ParameterToString(programIds, "multi")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Programinsightssettingsentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Programinsightssettingsentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsProgramsTranscriptionenginesDialects invokes GET /api/v2/speechandtextanalytics/programs/transcriptionengines/dialects
//
// Get supported dialects for each transcription engine
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsProgramsTranscriptionenginesDialects() (*Supporteddialectsentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/transcriptionengines/dialects"
	defaultReturn := new(Supporteddialectsentitylisting)
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
	var successPayload *Supporteddialectsentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Supporteddialectsentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsProgramsUnpublished invokes GET /api/v2/speechandtextanalytics/programs/unpublished
//
// Get the list of Speech & Text Analytics unpublished programs
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsProgramsUnpublished(nextPage string, pageSize int) (*Unpublishedprogramsentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/unpublished"
	defaultReturn := new(Unpublishedprogramsentitylisting)
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
	
	queryParams["nextPage"] = a.Configuration.APIClient.ParameterToString(nextPage, "")
	
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
	var successPayload *Unpublishedprogramsentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Unpublishedprogramsentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsReprocessingJob invokes GET /api/v2/speechandtextanalytics/reprocessing/jobs/{jobId}
//
// Get a Speech & Text Analytics reprocess job by id
//
// Preview: GetSpeechandtextanalyticsReprocessingJob is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsReprocessingJob(jobId string) (*Reprocessjobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/reprocessing/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Reprocessjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsReprocessingJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Reprocessjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Reprocessjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsReprocessingJobInteractions invokes GET /api/v2/speechandtextanalytics/reprocessing/jobs/{jobId}/interactions
//
// Get a Speech & Text Analytics Reprocessing interactions statuses by job id
//
// Preview: GetSpeechandtextanalyticsReprocessingJobInteractions is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsReprocessingJobInteractions(jobId string) (*Reprocessinteractionsbyjobidresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/reprocessing/jobs/{jobId}/interactions"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Reprocessinteractionsbyjobidresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsReprocessingJobInteractions")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Reprocessinteractionsbyjobidresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Reprocessinteractionsbyjobidresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsReprocessingJobs invokes GET /api/v2/speechandtextanalytics/reprocessing/jobs
//
// Get the list of Speech & Text Analytics reprocess jobs
//
// Preview: GetSpeechandtextanalyticsReprocessingJobs is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsReprocessingJobs(pageSize int, pageNumber int, sortOrder string, name string) (*Reprocessjobentitylistingresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/reprocessing/jobs"
	defaultReturn := new(Reprocessjobentitylistingresponse)
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
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
	queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Reprocessjobentitylistingresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Reprocessjobentitylistingresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsSentimentDialects invokes GET /api/v2/speechandtextanalytics/sentiment/dialects
//
// Get the list of Speech & Text Analytics sentiment supported dialects
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsSentimentDialects() (*Entitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/sentiment/dialects"
	defaultReturn := new(Entitylisting)
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

// GetSpeechandtextanalyticsSentimentfeedback invokes GET /api/v2/speechandtextanalytics/sentimentfeedback
//
// Get the list of Speech & Text Analytics SentimentFeedback
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsSentimentfeedback(dialect string) (*Sentimentfeedbackentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/sentimentfeedback"
	defaultReturn := new(Sentimentfeedbackentitylisting)
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
	
	queryParams["dialect"] = a.Configuration.APIClient.ParameterToString(dialect, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Sentimentfeedbackentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Sentimentfeedbackentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsSettings invokes GET /api/v2/speechandtextanalytics/settings
//
// Get Speech And Text Analytics Settings
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsSettings() (*Speechtextanalyticssettingsresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/settings"
	defaultReturn := new(Speechtextanalyticssettingsresponse)
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
	var successPayload *Speechtextanalyticssettingsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Speechtextanalyticssettingsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsTopic invokes GET /api/v2/speechandtextanalytics/topics/{topicId}
//
// Get a Speech & Text Analytics topic by id
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsTopic(topicId string) (*Topic, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/topics/{topicId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	defaultReturn := new(Topic)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsTopic")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Topic
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Topic" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsTopics invokes GET /api/v2/speechandtextanalytics/topics
//
// Get the list of Speech & Text Analytics topics
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsTopics(nextPage string, pageSize int, state string, name string, ids []string, dialects []string, sortBy string, sortOrder string) (*Topicsentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/topics"
	defaultReturn := new(Topicsentitylisting)
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
	
	queryParams["nextPage"] = a.Configuration.APIClient.ParameterToString(nextPage, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["state"] = a.Configuration.APIClient.ParameterToString(state, "")
	
	queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, "")
	
	queryParams["ids"] = a.Configuration.APIClient.ParameterToString(ids, "multi")
	
	queryParams["dialects"] = a.Configuration.APIClient.ParameterToString(dialects, "multi")
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
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
	var successPayload *Topicsentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Topicsentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsTopicsDialects invokes GET /api/v2/speechandtextanalytics/topics/dialects
//
// Get list of supported Speech & Text Analytics topics dialects
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsTopicsDialects() (*Entitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/topics/dialects"
	defaultReturn := new(Entitylisting)
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

// GetSpeechandtextanalyticsTopicsGeneral invokes GET /api/v2/speechandtextanalytics/topics/general
//
// Get the Speech & Text Analytics general topics for a given dialect
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsTopicsGeneral(dialect string) (*Generaltopicsentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/topics/general"
	defaultReturn := new(Generaltopicsentitylisting)
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
	
	queryParams["dialect"] = a.Configuration.APIClient.ParameterToString(dialect, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Generaltopicsentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Generaltopicsentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsTopicsGeneralStatus invokes GET /api/v2/speechandtextanalytics/topics/general/status
//
// Get the list of general topics from the org and the system with their current status
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsTopicsGeneralStatus(dialect string) (*Unifiedgeneraltopicentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/topics/general/status"
	defaultReturn := new(Unifiedgeneraltopicentitylisting)
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
	
	queryParams["dialect"] = a.Configuration.APIClient.ParameterToString(dialect, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Unifiedgeneraltopicentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Unifiedgeneraltopicentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsTopicsPublishjob invokes GET /api/v2/speechandtextanalytics/topics/publishjobs/{jobId}
//
// Get a Speech & Text Analytics publish topics job by id
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsTopicsPublishjob(jobId string) (*Topicjob, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/topics/publishjobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Topicjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsTopicsPublishjob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Topicjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Topicjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsTopicsTestphraseJob invokes GET /api/v2/speechandtextanalytics/topics/testphrase/jobs/{jobId}
//
// Get a Speech & Text Analytics test topics phrase job by id
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsTopicsTestphraseJob(jobId string) (*Testtopicphrasejob, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/topics/testphrase/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Testtopicphrasejob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsTopicsTestphraseJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Testtopicphrasejob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Testtopicphrasejob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsTranslationsLanguageConversation invokes GET /api/v2/speechandtextanalytics/translations/languages/{languageId}/conversations/{conversationId}
//
// Translate a single interaction recording (or an email conversation)
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsTranslationsLanguageConversation(languageId string, conversationId string, communicationId string, recordingId string) (*Communicationtranslationlist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/translations/languages/{languageId}/conversations/{conversationId}"
	path = strings.Replace(path, "{languageId}", url.PathEscape(fmt.Sprintf("%v", languageId)), -1)
	path = strings.Replace(path, "{conversationId}", url.PathEscape(fmt.Sprintf("%v", conversationId)), -1)
	defaultReturn := new(Communicationtranslationlist)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'languageId' is set
	if &languageId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'languageId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsTranslationsLanguageConversation")
	}
	// verify the required parameter 'conversationId' is set
	if &conversationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'conversationId' when calling SpeechTextAnalyticsApi->GetSpeechandtextanalyticsTranslationsLanguageConversation")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["communicationId"] = a.Configuration.APIClient.ParameterToString(communicationId, "")
	
	queryParams["recordingId"] = a.Configuration.APIClient.ParameterToString(recordingId, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Communicationtranslationlist
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Communicationtranslationlist" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSpeechandtextanalyticsTranslationsLanguages invokes GET /api/v2/speechandtextanalytics/translations/languages
//
// Get supported translation languages
func (a SpeechTextAnalyticsApi) GetSpeechandtextanalyticsTranslationsLanguages() (*Translatesupportedlanguagelist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/translations/languages"
	defaultReturn := new(Translatesupportedlanguagelist)
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
	var successPayload *Translatesupportedlanguagelist
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Translatesupportedlanguagelist" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchSpeechandtextanalyticsSettings invokes PATCH /api/v2/speechandtextanalytics/settings
//
// Patch Speech And Text Analytics Settings
func (a SpeechTextAnalyticsApi) PatchSpeechandtextanalyticsSettings(body Speechtextanalyticssettingsrequest) (*Speechtextanalyticssettingsresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/settings"
	defaultReturn := new(Speechtextanalyticssettingsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PatchSpeechandtextanalyticsSettings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Speechtextanalyticssettingsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Speechtextanalyticssettingsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSpeechandtextanalyticsCategories invokes POST /api/v2/speechandtextanalytics/categories
//
// Create new Speech & Text Analytics category
func (a SpeechTextAnalyticsApi) PostSpeechandtextanalyticsCategories(body Categoryrequest) (*Stacategory, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/categories"
	defaultReturn := new(Stacategory)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PostSpeechandtextanalyticsCategories")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Stacategory
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Stacategory" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSpeechandtextanalyticsDictionaryfeedback invokes POST /api/v2/speechandtextanalytics/dictionaryfeedback
//
// Create a Speech & Text Analytics DictionaryFeedback
func (a SpeechTextAnalyticsApi) PostSpeechandtextanalyticsDictionaryfeedback(body Dictionaryfeedback) (*Dictionaryfeedback, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/dictionaryfeedback"
	defaultReturn := new(Dictionaryfeedback)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PostSpeechandtextanalyticsDictionaryfeedback")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Dictionaryfeedback
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dictionaryfeedback" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSpeechandtextanalyticsPrograms invokes POST /api/v2/speechandtextanalytics/programs
//
// Create new Speech & Text Analytics program
func (a SpeechTextAnalyticsApi) PostSpeechandtextanalyticsPrograms(body Programrequest) (*Program, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs"
	defaultReturn := new(Program)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PostSpeechandtextanalyticsPrograms")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Program
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Program" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSpeechandtextanalyticsProgramsGeneralJobs invokes POST /api/v2/speechandtextanalytics/programs/general/jobs
//
// Create new Speech & Text Analytics general program job
func (a SpeechTextAnalyticsApi) PostSpeechandtextanalyticsProgramsGeneralJobs(body Generalprogramjobrequest) (*Generalprogramjob, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/general/jobs"
	defaultReturn := new(Generalprogramjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PostSpeechandtextanalyticsProgramsGeneralJobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Generalprogramjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Generalprogramjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSpeechandtextanalyticsProgramsPublishjobs invokes POST /api/v2/speechandtextanalytics/programs/publishjobs
//
// Create new Speech & Text Analytics publish programs job
func (a SpeechTextAnalyticsApi) PostSpeechandtextanalyticsProgramsPublishjobs(body Programjobrequest) (*Programjob, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/publishjobs"
	defaultReturn := new(Programjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PostSpeechandtextanalyticsProgramsPublishjobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Programjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Programjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSpeechandtextanalyticsReprocessingJobs invokes POST /api/v2/speechandtextanalytics/reprocessing/jobs
//
// Create a Speech & Text Analytics reprocess job.
//
// Preview: PostSpeechandtextanalyticsReprocessingJobs is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SpeechTextAnalyticsApi) PostSpeechandtextanalyticsReprocessingJobs(body Createreprocessjobrequest) (*Reprocessjobresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/reprocessing/jobs"
	defaultReturn := new(Reprocessjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PostSpeechandtextanalyticsReprocessingJobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Reprocessjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Reprocessjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSpeechandtextanalyticsSentimentfeedback invokes POST /api/v2/speechandtextanalytics/sentimentfeedback
//
// Create a Speech & Text Analytics SentimentFeedback
func (a SpeechTextAnalyticsApi) PostSpeechandtextanalyticsSentimentfeedback(body Sentimentfeedback) (*Sentimentfeedback, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/sentimentfeedback"
	defaultReturn := new(Sentimentfeedback)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PostSpeechandtextanalyticsSentimentfeedback")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Sentimentfeedback
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Sentimentfeedback" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSpeechandtextanalyticsTopics invokes POST /api/v2/speechandtextanalytics/topics
//
// Create new Speech & Text Analytics topic
func (a SpeechTextAnalyticsApi) PostSpeechandtextanalyticsTopics(body Topicrequest) (*Topic, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/topics"
	defaultReturn := new(Topic)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PostSpeechandtextanalyticsTopics")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Topic
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Topic" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSpeechandtextanalyticsTopicsPublishjobs invokes POST /api/v2/speechandtextanalytics/topics/publishjobs
//
// Create new Speech & Text Analytics publish topics job
func (a SpeechTextAnalyticsApi) PostSpeechandtextanalyticsTopicsPublishjobs(body Topicjobrequest) (*Topicjob, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/topics/publishjobs"
	defaultReturn := new(Topicjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PostSpeechandtextanalyticsTopicsPublishjobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Topicjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Topicjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSpeechandtextanalyticsTopicsTestphraseJobs invokes POST /api/v2/speechandtextanalytics/topics/testphrase/jobs
//
// Create new Speech & Text Analytics publish topics job
func (a SpeechTextAnalyticsApi) PostSpeechandtextanalyticsTopicsTestphraseJobs(body Testtopicphrasejobrequest) (*Testtopicphrasejobs, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/topics/testphrase/jobs"
	defaultReturn := new(Testtopicphrasejobs)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PostSpeechandtextanalyticsTopicsTestphraseJobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Testtopicphrasejobs
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Testtopicphrasejobs" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSpeechandtextanalyticsTranscriptsSearch invokes POST /api/v2/speechandtextanalytics/transcripts/search
//
// Search resources.
func (a SpeechTextAnalyticsApi) PostSpeechandtextanalyticsTranscriptsSearch(body Transcriptsearchrequest) (*Jsonsearchresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/transcripts/search"
	defaultReturn := new(Jsonsearchresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PostSpeechandtextanalyticsTranscriptsSearch")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Jsonsearchresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Jsonsearchresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutSpeechandtextanalyticsCategory invokes PUT /api/v2/speechandtextanalytics/categories/{categoryId}
//
// Update a Speech & Text Analytics category by ID
func (a SpeechTextAnalyticsApi) PutSpeechandtextanalyticsCategory(categoryId string, body Categoryrequest) (*Stacategory, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/categories/{categoryId}"
	path = strings.Replace(path, "{categoryId}", url.PathEscape(fmt.Sprintf("%v", categoryId)), -1)
	defaultReturn := new(Stacategory)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'categoryId' is set
	if &categoryId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'categoryId' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsCategory")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsCategory")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Stacategory
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Stacategory" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutSpeechandtextanalyticsDictionaryfeedbackDictionaryFeedbackId invokes PUT /api/v2/speechandtextanalytics/dictionaryfeedback/{dictionaryFeedbackId}
//
// Update existing Speech & Text Analytics dictionary feedback by id
func (a SpeechTextAnalyticsApi) PutSpeechandtextanalyticsDictionaryfeedbackDictionaryFeedbackId(dictionaryFeedbackId string, body Dictionaryfeedback) (*Dictionaryfeedback, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/dictionaryfeedback/{dictionaryFeedbackId}"
	path = strings.Replace(path, "{dictionaryFeedbackId}", url.PathEscape(fmt.Sprintf("%v", dictionaryFeedbackId)), -1)
	defaultReturn := new(Dictionaryfeedback)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'dictionaryFeedbackId' is set
	if &dictionaryFeedbackId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'dictionaryFeedbackId' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsDictionaryfeedbackDictionaryFeedbackId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Dictionaryfeedback
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dictionaryfeedback" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutSpeechandtextanalyticsProgram invokes PUT /api/v2/speechandtextanalytics/programs/{programId}
//
// Update existing Speech & Text Analytics program
func (a SpeechTextAnalyticsApi) PutSpeechandtextanalyticsProgram(programId string, body Programrequest) (*Program, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/{programId}"
	path = strings.Replace(path, "{programId}", url.PathEscape(fmt.Sprintf("%v", programId)), -1)
	defaultReturn := new(Program)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'programId' is set
	if &programId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'programId' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsProgram")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsProgram")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Program
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Program" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutSpeechandtextanalyticsProgramMappings invokes PUT /api/v2/speechandtextanalytics/programs/{programId}/mappings
//
// Set Speech & Text Analytics program mappings to queues and flows
func (a SpeechTextAnalyticsApi) PutSpeechandtextanalyticsProgramMappings(programId string, body Programmappingsrequest) (*Programmappings, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/{programId}/mappings"
	path = strings.Replace(path, "{programId}", url.PathEscape(fmt.Sprintf("%v", programId)), -1)
	defaultReturn := new(Programmappings)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'programId' is set
	if &programId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'programId' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsProgramMappings")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsProgramMappings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Programmappings
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Programmappings" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutSpeechandtextanalyticsProgramSettingsInsights invokes PUT /api/v2/speechandtextanalytics/programs/{programId}/settings/insights
//
// Update AI Insights settings of a program
func (a SpeechTextAnalyticsApi) PutSpeechandtextanalyticsProgramSettingsInsights(programId string, body Insightssettingsrequest) (*Programinsightssettings, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/{programId}/settings/insights"
	path = strings.Replace(path, "{programId}", url.PathEscape(fmt.Sprintf("%v", programId)), -1)
	defaultReturn := new(Programinsightssettings)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'programId' is set
	if &programId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'programId' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsProgramSettingsInsights")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsProgramSettingsInsights")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Programinsightssettings
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Programinsightssettings" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutSpeechandtextanalyticsProgramTranscriptionengines invokes PUT /api/v2/speechandtextanalytics/programs/{programId}/transcriptionengines
//
// Update transcription engine settings of a program
func (a SpeechTextAnalyticsApi) PutSpeechandtextanalyticsProgramTranscriptionengines(programId string, body Transcriptionenginesrequest) (*Programtranscriptionengines, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/programs/{programId}/transcriptionengines"
	path = strings.Replace(path, "{programId}", url.PathEscape(fmt.Sprintf("%v", programId)), -1)
	defaultReturn := new(Programtranscriptionengines)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'programId' is set
	if &programId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'programId' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsProgramTranscriptionengines")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsProgramTranscriptionengines")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Programtranscriptionengines
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Programtranscriptionengines" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutSpeechandtextanalyticsSettings invokes PUT /api/v2/speechandtextanalytics/settings
//
// Update Speech And Text Analytics Settings
func (a SpeechTextAnalyticsApi) PutSpeechandtextanalyticsSettings(body Speechtextanalyticssettingsrequest) (*Speechtextanalyticssettingsresponse, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/settings"
	defaultReturn := new(Speechtextanalyticssettingsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsSettings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Speechtextanalyticssettingsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Speechtextanalyticssettingsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutSpeechandtextanalyticsTopic invokes PUT /api/v2/speechandtextanalytics/topics/{topicId}
//
// Update existing Speech & Text Analytics topic
func (a SpeechTextAnalyticsApi) PutSpeechandtextanalyticsTopic(topicId string, body Topicrequest) (*Topic, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/speechandtextanalytics/topics/{topicId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	defaultReturn := new(Topic)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsTopic")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SpeechTextAnalyticsApi->PutSpeechandtextanalyticsTopic")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Topic
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Topic" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

