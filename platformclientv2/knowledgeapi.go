package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

// KnowledgeApi provides functions for API endpoints
type KnowledgeApi struct {
	Configuration *Configuration
}

// NewKnowledgeApi creates an API instance using the default configuration
func NewKnowledgeApi() *KnowledgeApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating KnowledgeApi with base path: %s", strings.ToLower(config.BasePath)))
	return &KnowledgeApi{
		Configuration: config,
	}
}

// NewKnowledgeApiWithConfig creates an API instance using the provided configuration
func NewKnowledgeApiWithConfig(config *Configuration) *KnowledgeApi {
	config.Debugf("Creating KnowledgeApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &KnowledgeApi{
		Configuration: config,
	}
}

// DeleteKnowledgeKnowledgebase invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}
//
// Delete knowledge base
//
// 
func (a KnowledgeApi) DeleteKnowledgeKnowledgebase(knowledgeBaseId string) (*Knowledgebase, *APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Knowledgebase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebase")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgebase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// DeleteKnowledgeKnowledgebaseLanguageCategory invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}
//
// Delete category
//
// 
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseLanguageCategory(categoryId string, knowledgeBaseId string, languageCode string) (*Knowledgecategory, *APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}"
	path = strings.Replace(path, "{categoryId}", fmt.Sprintf("%v", categoryId), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Knowledgecategory)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'categoryId' is set
	if &categoryId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'categoryId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLanguageCategory")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgecategory
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// DeleteKnowledgeKnowledgebaseLanguageDocument invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}
//
// Delete document
//
// 
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseLanguageDocument(documentId string, knowledgeBaseId string, languageCode string) (*Knowledgedocument, *APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}"
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Knowledgedocument)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLanguageDocument")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgedocument
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebase invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}
//
// Get knowledge base
//
// 
func (a KnowledgeApi) GetKnowledgeKnowledgebase(knowledgeBaseId string) (*Knowledgebase, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Knowledgebase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebase")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgebase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLanguageCategories invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories
//
// Get categories
//
// 
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageCategories(knowledgeBaseId string, languageCode string, before string, after string, limit string, pageSize string) (*Categorylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Categorylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageCategories")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageCategories")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	if str, ok := interface{}(before).(string); ok {
		if str != "" {
			queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, collectionFormat)
		}
	} else {
		queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(after).(string); ok {
		if str != "" {
			queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
		}
	} else {
		queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(limit).(string); ok {
		if str != "" {
			queryParams["limit"] = a.Configuration.APIClient.ParameterToString(limit, collectionFormat)
		}
	} else {
		queryParams["limit"] = a.Configuration.APIClient.ParameterToString(limit, collectionFormat)
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
	var successPayload *Categorylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLanguageCategory invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}
//
// Get category
//
// 
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageCategory(categoryId string, knowledgeBaseId string, languageCode string) (*Knowledgeextendedcategory, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}"
	path = strings.Replace(path, "{categoryId}", fmt.Sprintf("%v", categoryId), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Knowledgeextendedcategory)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'categoryId' is set
	if &categoryId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'categoryId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageCategory")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgeextendedcategory
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLanguageDocument invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}
//
// Get document
//
// 
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageDocument(documentId string, knowledgeBaseId string, languageCode string) (*Knowledgedocument, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}"
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Knowledgedocument)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocument")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgedocument
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLanguageDocuments invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents
//
// Get documents
//
// 
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageDocuments(knowledgeBaseId string, languageCode string, before string, after string, limit string, pageSize string, categories string, title string) (*Documentlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Documentlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocuments")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocuments")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	if str, ok := interface{}(before).(string); ok {
		if str != "" {
			queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, collectionFormat)
		}
	} else {
		queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(after).(string); ok {
		if str != "" {
			queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
		}
	} else {
		queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(limit).(string); ok {
		if str != "" {
			queryParams["limit"] = a.Configuration.APIClient.ParameterToString(limit, collectionFormat)
		}
	} else {
		queryParams["limit"] = a.Configuration.APIClient.ParameterToString(limit, collectionFormat)
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
	if str, ok := interface{}(categories).(string); ok {
		if str != "" {
			queryParams["categories"] = a.Configuration.APIClient.ParameterToString(categories, collectionFormat)
		}
	} else {
		queryParams["categories"] = a.Configuration.APIClient.ParameterToString(categories, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(title).(string); ok {
		if str != "" {
			queryParams["title"] = a.Configuration.APIClient.ParameterToString(title, collectionFormat)
		}
	} else {
		queryParams["title"] = a.Configuration.APIClient.ParameterToString(title, collectionFormat)
	}
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Documentlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLanguageTraining invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings/{trainingId}
//
// Get training detail
//
// 
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageTraining(knowledgeBaseId string, languageCode string, trainingId string) (*Knowledgetraining, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings/{trainingId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	path = strings.Replace(path, "{trainingId}", fmt.Sprintf("%v", trainingId), -1)
	defaultReturn := new(Knowledgetraining)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageTraining")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageTraining")
	}
	// verify the required parameter 'trainingId' is set
	if &trainingId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'trainingId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageTraining")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgetraining
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLanguageTrainings invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings
//
// Get all trainings information for a knowledgebase
//
// 
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageTrainings(knowledgeBaseId string, languageCode string, before string, after string, limit string, pageSize string, knowledgeDocumentsState string) (*Traininglisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Traininglisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageTrainings")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageTrainings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	if str, ok := interface{}(before).(string); ok {
		if str != "" {
			queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, collectionFormat)
		}
	} else {
		queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(after).(string); ok {
		if str != "" {
			queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
		}
	} else {
		queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(limit).(string); ok {
		if str != "" {
			queryParams["limit"] = a.Configuration.APIClient.ParameterToString(limit, collectionFormat)
		}
	} else {
		queryParams["limit"] = a.Configuration.APIClient.ParameterToString(limit, collectionFormat)
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
	if str, ok := interface{}(knowledgeDocumentsState).(string); ok {
		if str != "" {
			queryParams["knowledgeDocumentsState"] = a.Configuration.APIClient.ParameterToString(knowledgeDocumentsState, collectionFormat)
		}
	} else {
		queryParams["knowledgeDocumentsState"] = a.Configuration.APIClient.ParameterToString(knowledgeDocumentsState, collectionFormat)
	}
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Traininglisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebases invokes GET /api/v2/knowledge/knowledgebases
//
// Get knowledge bases
//
// 
func (a KnowledgeApi) GetKnowledgeKnowledgebases(before string, after string, limit string, pageSize string, name string) (*Knowledgebaselisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases"
	defaultReturn := new(Knowledgebaselisting)
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
	if str, ok := interface{}(before).(string); ok {
		if str != "" {
			queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, collectionFormat)
		}
	} else {
		queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(after).(string); ok {
		if str != "" {
			queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
		}
	} else {
		queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(limit).(string); ok {
		if str != "" {
			queryParams["limit"] = a.Configuration.APIClient.ParameterToString(limit, collectionFormat)
		}
	} else {
		queryParams["limit"] = a.Configuration.APIClient.ParameterToString(limit, collectionFormat)
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
	if str, ok := interface{}(name).(string); ok {
		if str != "" {
			queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
		}
	} else {
		queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
	}
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Knowledgebaselisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebase invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}
//
// Update knowledge base
//
// 
func (a KnowledgeApi) PatchKnowledgeKnowledgebase(knowledgeBaseId string, body Knowledgebase) (*Knowledgebase, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Knowledgebase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebase")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeKnowledgebase")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgebase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebaseLanguageCategory invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}
//
// Update category
//
// 
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseLanguageCategory(categoryId string, knowledgeBaseId string, languageCode string, body Knowledgecategoryrequest) (*Knowledgeextendedcategory, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}"
	path = strings.Replace(path, "{categoryId}", fmt.Sprintf("%v", categoryId), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Knowledgeextendedcategory)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'categoryId' is set
	if &categoryId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'categoryId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageCategory")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgeextendedcategory
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebaseLanguageDocument invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}
//
// Update document
//
// 
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseLanguageDocument(documentId string, knowledgeBaseId string, languageCode string, body Knowledgedocumentrequest) (*Knowledgedocument, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}"
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Knowledgedocument)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocument")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgedocument
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebaseLanguageDocuments invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents
//
// Update documents collection
//
// 
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseLanguageDocuments(knowledgeBaseId string, languageCode string, body []Knowledgedocumentbulkrequest) (*Documentlisting, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Documentlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocuments")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocuments")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// true
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocuments")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Documentlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseLanguageCategories invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories
//
// Create new category
//
// 
func (a KnowledgeApi) PostKnowledgeKnowledgebaseLanguageCategories(knowledgeBaseId string, languageCode string, body Knowledgecategoryrequest) (*Knowledgeextendedcategory, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Knowledgeextendedcategory)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageCategories")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageCategories")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageCategories")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgeextendedcategory
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseLanguageDocuments invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents
//
// Create document
//
// 
func (a KnowledgeApi) PostKnowledgeKnowledgebaseLanguageDocuments(knowledgeBaseId string, languageCode string, body Knowledgedocumentrequest) (*Knowledgedocument, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Knowledgedocument)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageDocuments")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageDocuments")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageDocuments")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgedocument
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseLanguageTrainingPromote invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings/{trainingId}/promote
//
// Promote trained documents from draft state to active.
//
// 
func (a KnowledgeApi) PostKnowledgeKnowledgebaseLanguageTrainingPromote(knowledgeBaseId string, languageCode string, trainingId string) (*Knowledgetraining, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings/{trainingId}/promote"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	path = strings.Replace(path, "{trainingId}", fmt.Sprintf("%v", trainingId), -1)
	defaultReturn := new(Knowledgetraining)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageTrainingPromote")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageTrainingPromote")
	}
	// verify the required parameter 'trainingId' is set
	if &trainingId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'trainingId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageTrainingPromote")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgetraining
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseLanguageTrainings invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings
//
// Trigger training
//
// 
func (a KnowledgeApi) PostKnowledgeKnowledgebaseLanguageTrainings(knowledgeBaseId string, languageCode string) (*Knowledgetraining, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Knowledgetraining)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageTrainings")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageTrainings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgetraining
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseSearch invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/search
//
// Search Documents
//
// 
func (a KnowledgeApi) PostKnowledgeKnowledgebaseSearch(knowledgeBaseId string, body Knowledgesearchrequest) (*Knowledgesearchresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/search"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Knowledgesearchresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseSearch")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgesearchresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebases invokes POST /api/v2/knowledge/knowledgebases
//
// Create new knowledge base
//
// 
func (a KnowledgeApi) PostKnowledgeKnowledgebases(body Knowledgebase) (*Knowledgebase, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases"
	defaultReturn := new(Knowledgebase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebases")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgebase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

