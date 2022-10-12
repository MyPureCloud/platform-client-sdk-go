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
	fmt.Sprintf(strings.Title(""), "")
	config := GetDefaultConfiguration()
	return &KnowledgeApi{
		Configuration: config,
	}
}

// NewKnowledgeApiWithConfig creates an API instance using the provided configuration
func NewKnowledgeApiWithConfig(config *Configuration) *KnowledgeApi {
	return &KnowledgeApi{
		Configuration: config,
	}
}

// DeleteKnowledgeKnowledgebase invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}
//
// Delete knowledge base
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
		// false
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
	} else if response.HasBody {
		if "Knowledgebase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// DeleteKnowledgeKnowledgebaseCategory invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}
//
// Delete category
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseCategory(knowledgeBaseId string, categoryId string) (*Categoryresponse, *APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{categoryId}", fmt.Sprintf("%v", categoryId), -1)
	defaultReturn := new(Categoryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseCategory")
	}
	// verify the required parameter 'categoryId' is set
	if &categoryId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'categoryId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseCategory")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Categoryresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Categoryresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// DeleteKnowledgeKnowledgebaseDocument invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}
//
// Delete document.
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseDocument(knowledgeBaseId string, documentId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseDocument")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseDocument")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteKnowledgeKnowledgebaseDocumentVariation invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}
//
// Delete a variation for a document.
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseDocumentVariation(documentVariationId string, documentId string, knowledgeBaseId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}"
	path = strings.Replace(path, "{documentVariationId}", fmt.Sprintf("%v", documentVariationId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'documentVariationId' is set
	if &documentVariationId == nil {
		// false
		return nil, errors.New("Missing required parameter 'documentVariationId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseDocumentVariation")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseDocumentVariation")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseDocumentVariation")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteKnowledgeKnowledgebaseExportJob invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}
//
// Delete export job
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseExportJob(knowledgeBaseId string, exportJobId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{exportJobId}", fmt.Sprintf("%v", exportJobId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseExportJob")
	}
	// verify the required parameter 'exportJobId' is set
	if &exportJobId == nil {
		// false
		return nil, errors.New("Missing required parameter 'exportJobId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseExportJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteKnowledgeKnowledgebaseImportJob invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}
//
// Delete import job
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseImportJob(knowledgeBaseId string, importJobId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{importJobId}", fmt.Sprintf("%v", importJobId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseImportJob")
	}
	// verify the required parameter 'importJobId' is set
	if &importJobId == nil {
		// false
		return nil, errors.New("Missing required parameter 'importJobId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseImportJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteKnowledgeKnowledgebaseLabel invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}
//
// Delete label
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseLabel(knowledgeBaseId string, labelId string) (*Labelresponse, *APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{labelId}", fmt.Sprintf("%v", labelId), -1)
	defaultReturn := new(Labelresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLabel")
	}
	// verify the required parameter 'labelId' is set
	if &labelId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'labelId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLabel")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Labelresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Labelresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// DeleteKnowledgeKnowledgebaseLanguageCategory invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}
//
// Delete category
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'categoryId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
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
	} else if response.HasBody {
		if "Knowledgecategory" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// DeleteKnowledgeKnowledgebaseLanguageDocument invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}
//
// Delete document
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
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
	} else if response.HasBody {
		if "Knowledgedocument" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// DeleteKnowledgeKnowledgebaseLanguageDocumentsImport invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/imports/{importId}
//
// Delete import operation
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseLanguageDocumentsImport(knowledgeBaseId string, languageCode string, importId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/imports/{importId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	path = strings.Replace(path, "{importId}", fmt.Sprintf("%v", importId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLanguageDocumentsImport")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
		return nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLanguageDocumentsImport")
	}
	// verify the required parameter 'importId' is set
	if &importId == nil {
		// false
		return nil, errors.New("Missing required parameter 'importId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseLanguageDocumentsImport")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// GetKnowledgeGuestSessionCategories invokes GET /api/v2/knowledge/guest/sessions/{sessionId}/categories
//
// Get categories
func (a KnowledgeApi) GetKnowledgeGuestSessionCategories(sessionId string, before string, after string, pageSize string, parentId string, isRoot bool, name string, sortBy string, expand string, includeDocumentCount bool) (*Guestcategoryresponselisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/categories"
	path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)
	defaultReturn := new(Guestcategoryresponselisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sessionId' is set
	if &sessionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'sessionId' when calling KnowledgeApi->GetKnowledgeGuestSessionCategories")
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
	
	queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, "")
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["parentId"] = a.Configuration.APIClient.ParameterToString(parentId, "")
	
	queryParams["isRoot"] = a.Configuration.APIClient.ParameterToString(isRoot, "")
	
	queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, "")
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "")
	
	queryParams["includeDocumentCount"] = a.Configuration.APIClient.ParameterToString(includeDocumentCount, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Guestcategoryresponselisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Guestcategoryresponselisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeGuestSessionDocument invokes GET /api/v2/knowledge/guest/sessions/{sessionId}/documents/{documentId}
//
// Get a knowledge document by ID.
func (a KnowledgeApi) GetKnowledgeGuestSessionDocument(sessionId string, documentId string) (*Knowledgeguestdocument, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/documents/{documentId}"
	path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	defaultReturn := new(Knowledgeguestdocument)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sessionId' is set
	if &sessionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'sessionId' when calling KnowledgeApi->GetKnowledgeGuestSessionDocument")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->GetKnowledgeGuestSessionDocument")
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Knowledgeguestdocument
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeguestdocument" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeGuestSessionDocuments invokes GET /api/v2/knowledge/guest/sessions/{sessionId}/documents
//
// Get documents.
func (a KnowledgeApi) GetKnowledgeGuestSessionDocuments(sessionId string, categoryId []string, includeSubcategories bool, pageSize string) (*Knowledgeguestdocumentresponselisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/documents"
	path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)
	defaultReturn := new(Knowledgeguestdocumentresponselisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sessionId' is set
	if &sessionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'sessionId' when calling KnowledgeApi->GetKnowledgeGuestSessionDocuments")
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
	
	queryParams["categoryId"] = a.Configuration.APIClient.ParameterToString(categoryId, "multi")
	
	queryParams["includeSubcategories"] = a.Configuration.APIClient.ParameterToString(includeSubcategories, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Knowledgeguestdocumentresponselisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeguestdocumentresponselisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebase invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}
//
// Get knowledge base
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
		// false
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
	} else if response.HasBody {
		if "Knowledgebase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseCategories invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories
//
// Get categories
func (a KnowledgeApi) GetKnowledgeKnowledgebaseCategories(knowledgeBaseId string, before string, after string, pageSize string, parentId string, isRoot bool, name string, sortBy string, expand string, includeDocumentCount bool) (*Categoryresponselisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Categoryresponselisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseCategories")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, "")
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["parentId"] = a.Configuration.APIClient.ParameterToString(parentId, "")
	
	queryParams["isRoot"] = a.Configuration.APIClient.ParameterToString(isRoot, "")
	
	queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, "")
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "")
	
	queryParams["includeDocumentCount"] = a.Configuration.APIClient.ParameterToString(includeDocumentCount, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Categoryresponselisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Categoryresponselisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseCategory invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}
//
// Get category
func (a KnowledgeApi) GetKnowledgeKnowledgebaseCategory(knowledgeBaseId string, categoryId string) (*Categoryresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{categoryId}", fmt.Sprintf("%v", categoryId), -1)
	defaultReturn := new(Categoryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseCategory")
	}
	// verify the required parameter 'categoryId' is set
	if &categoryId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'categoryId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseCategory")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Categoryresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Categoryresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseDocument invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}
//
// Get document.
func (a KnowledgeApi) GetKnowledgeKnowledgebaseDocument(knowledgeBaseId string, documentId string, expand []string, state string) (*Knowledgedocumentresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	defaultReturn := new(Knowledgedocumentresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocument")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocument")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	
	queryParams["state"] = a.Configuration.APIClient.ParameterToString(state, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Knowledgedocumentresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseDocumentVariation invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}
//
// Get a variation for a document.
func (a KnowledgeApi) GetKnowledgeKnowledgebaseDocumentVariation(documentVariationId string, documentId string, knowledgeBaseId string, documentState string) (*Documentvariation, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}"
	path = strings.Replace(path, "{documentVariationId}", fmt.Sprintf("%v", documentVariationId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Documentvariation)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'documentVariationId' is set
	if &documentVariationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentVariationId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVariation")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVariation")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVariation")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["documentState"] = a.Configuration.APIClient.ParameterToString(documentState, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Documentvariation
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Documentvariation" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseDocumentVariations invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations
//
// Get variations for a document.
func (a KnowledgeApi) GetKnowledgeKnowledgebaseDocumentVariations(knowledgeBaseId string, documentId string, before string, after string, pageSize string, documentState string) (*Documentvariationlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	defaultReturn := new(Documentvariationlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVariations")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVariations")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, "")
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["documentState"] = a.Configuration.APIClient.ParameterToString(documentState, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Documentvariationlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Documentvariationlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseDocumentVersion invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}
//
// Get document version.
func (a KnowledgeApi) GetKnowledgeKnowledgebaseDocumentVersion(knowledgeBaseId string, documentId string, versionId string, expand []string) (*Knowledgedocumentversion, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	path = strings.Replace(path, "{versionId}", fmt.Sprintf("%v", versionId), -1)
	defaultReturn := new(Knowledgedocumentversion)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVersion")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVersion")
	}
	// verify the required parameter 'versionId' is set
	if &versionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'versionId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVersion")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Knowledgedocumentversion
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentversion" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseDocumentVersionVariation invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}/variations/{variationId}
//
// Get variation for the given document version.
func (a KnowledgeApi) GetKnowledgeKnowledgebaseDocumentVersionVariation(knowledgeBaseId string, documentId string, versionId string, variationId string) (*Knowledgedocumentversionvariation, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}/variations/{variationId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	path = strings.Replace(path, "{versionId}", fmt.Sprintf("%v", versionId), -1)
	path = strings.Replace(path, "{variationId}", fmt.Sprintf("%v", variationId), -1)
	defaultReturn := new(Knowledgedocumentversionvariation)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVersionVariation")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVersionVariation")
	}
	// verify the required parameter 'versionId' is set
	if &versionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'versionId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVersionVariation")
	}
	// verify the required parameter 'variationId' is set
	if &variationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'variationId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVersionVariation")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgedocumentversionvariation
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentversionvariation" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseDocumentVersionVariations invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}/variations
//
// Get variations for the given document version.
func (a KnowledgeApi) GetKnowledgeKnowledgebaseDocumentVersionVariations(knowledgeBaseId string, documentId string, versionId string, before string, after string, pageSize string) (*Knowledgedocumentversionvariationlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions/{versionId}/variations"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	path = strings.Replace(path, "{versionId}", fmt.Sprintf("%v", versionId), -1)
	defaultReturn := new(Knowledgedocumentversionvariationlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVersionVariations")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVersionVariations")
	}
	// verify the required parameter 'versionId' is set
	if &versionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'versionId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVersionVariations")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, "")
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Knowledgedocumentversionvariationlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentversionvariationlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseDocumentVersions invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions
//
// Get document versions.
func (a KnowledgeApi) GetKnowledgeKnowledgebaseDocumentVersions(knowledgeBaseId string, documentId string, before string, after string, pageSize string, expand []string) (*Knowledgedocumentversionlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	defaultReturn := new(Knowledgedocumentversionlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVersions")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentVersions")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, "")
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Knowledgedocumentversionlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentversionlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseDocuments invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents
//
// Get documents.
func (a KnowledgeApi) GetKnowledgeKnowledgebaseDocuments(knowledgeBaseId string, before string, after string, pageSize string, interval string, documentId []string, categoryId []string, includeSubcategories bool, includeDrafts bool, labelIds []string, expand []string) (*Knowledgedocumentresponselisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Knowledgedocumentresponselisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocuments")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, "")
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["interval"] = a.Configuration.APIClient.ParameterToString(interval, "")
	
	queryParams["documentId"] = a.Configuration.APIClient.ParameterToString(documentId, "multi")
	
	queryParams["categoryId"] = a.Configuration.APIClient.ParameterToString(categoryId, "multi")
	
	queryParams["includeSubcategories"] = a.Configuration.APIClient.ParameterToString(includeSubcategories, "")
	
	queryParams["includeDrafts"] = a.Configuration.APIClient.ParameterToString(includeDrafts, "")
	
	queryParams["labelIds"] = a.Configuration.APIClient.ParameterToString(labelIds, "multi")
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Knowledgedocumentresponselisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentresponselisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseExportJob invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}
//
// Get export job report
func (a KnowledgeApi) GetKnowledgeKnowledgebaseExportJob(knowledgeBaseId string, exportJobId string) (*Knowledgeexportjobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{exportJobId}", fmt.Sprintf("%v", exportJobId), -1)
	defaultReturn := new(Knowledgeexportjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseExportJob")
	}
	// verify the required parameter 'exportJobId' is set
	if &exportJobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'exportJobId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseExportJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgeexportjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeexportjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseImportJob invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}
//
// Get import job report
func (a KnowledgeApi) GetKnowledgeKnowledgebaseImportJob(knowledgeBaseId string, importJobId string) (*Knowledgeimportjobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{importJobId}", fmt.Sprintf("%v", importJobId), -1)
	defaultReturn := new(Knowledgeimportjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseImportJob")
	}
	// verify the required parameter 'importJobId' is set
	if &importJobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'importJobId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseImportJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgeimportjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeimportjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLabel invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}
//
// Get label
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLabel(knowledgeBaseId string, labelId string) (*Labelresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{labelId}", fmt.Sprintf("%v", labelId), -1)
	defaultReturn := new(Labelresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLabel")
	}
	// verify the required parameter 'labelId' is set
	if &labelId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'labelId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLabel")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Labelresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Labelresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLabels invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels
//
// Get labels
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLabels(knowledgeBaseId string, before string, after string, pageSize string, name string, includeDocumentCount bool) (*Labellisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Labellisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLabels")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, "")
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, "")
	
	queryParams["includeDocumentCount"] = a.Configuration.APIClient.ParameterToString(includeDocumentCount, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Labellisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Labellisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLanguageCategories invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories
//
// Get categories
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageCategories(knowledgeBaseId string, languageCode string, before string, after string, limit string, pageSize string, name string) (*Categorylisting, *APIResponse, error) {
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageCategories")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
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
	
	queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, "")
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["limit"] = a.Configuration.APIClient.ParameterToString(limit, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	} else if response.HasBody {
		if "Categorylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLanguageCategory invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}
//
// Get category
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'categoryId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
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
	} else if response.HasBody {
		if "Knowledgeextendedcategory" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLanguageDocument invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}
//
// Get document
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
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
	} else if response.HasBody {
		if "Knowledgedocument" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLanguageDocuments invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents
//
// Get documents
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageDocuments(knowledgeBaseId string, languageCode string, before string, after string, limit string, pageSize string, categories string, title string, sortBy string, sortOrder string, documentIds []string) (*Documentlisting, *APIResponse, error) {
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocuments")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
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
	
	queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, "")
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["limit"] = a.Configuration.APIClient.ParameterToString(limit, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["categories"] = a.Configuration.APIClient.ParameterToString(categories, "")
	
	queryParams["title"] = a.Configuration.APIClient.ParameterToString(title, "")
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
	queryParams["documentIds"] = a.Configuration.APIClient.ParameterToString(documentIds, "multi")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	} else if response.HasBody {
		if "Documentlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLanguageDocumentsImport invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/imports/{importId}
//
// Get import operation report
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageDocumentsImport(knowledgeBaseId string, languageCode string, importId string) (*Knowledgeimport, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/imports/{importId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	path = strings.Replace(path, "{importId}", fmt.Sprintf("%v", importId), -1)
	defaultReturn := new(Knowledgeimport)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocumentsImport")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocumentsImport")
	}
	// verify the required parameter 'importId' is set
	if &importId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'importId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocumentsImport")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgeimport
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeimport" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLanguageTraining invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings/{trainingId}
//
// Get training detail
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageTraining")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageTraining")
	}
	// verify the required parameter 'trainingId' is set
	if &trainingId == nil {
		// false
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
	} else if response.HasBody {
		if "Knowledgetraining" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseLanguageTrainings invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings
//
// Get all trainings information for a knowledgebase
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageTrainings")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
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
	
	queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, "")
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["limit"] = a.Configuration.APIClient.ParameterToString(limit, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["knowledgeDocumentsState"] = a.Configuration.APIClient.ParameterToString(knowledgeDocumentsState, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	} else if response.HasBody {
		if "Traininglisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseUnansweredGroup invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}
//
// Get knowledge base unanswered group for a particular groupId
func (a KnowledgeApi) GetKnowledgeKnowledgebaseUnansweredGroup(knowledgeBaseId string, groupId string, app string) (*Unansweredgroup, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupId), -1)
	defaultReturn := new(Unansweredgroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseUnansweredGroup")
	}
	// verify the required parameter 'groupId' is set
	if &groupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'groupId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseUnansweredGroup")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["app"] = a.Configuration.APIClient.ParameterToString(app, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Unansweredgroup
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Unansweredgroup" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseUnansweredGroupPhrasegroup invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}/phrasegroups/{phraseGroupId}
//
// Get knowledge base unanswered phrase group for a particular phraseGroupId
func (a KnowledgeApi) GetKnowledgeKnowledgebaseUnansweredGroupPhrasegroup(knowledgeBaseId string, groupId string, phraseGroupId string, app string) (*Unansweredphrasegroup, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}/phrasegroups/{phraseGroupId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupId), -1)
	path = strings.Replace(path, "{phraseGroupId}", fmt.Sprintf("%v", phraseGroupId), -1)
	defaultReturn := new(Unansweredphrasegroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseUnansweredGroupPhrasegroup")
	}
	// verify the required parameter 'groupId' is set
	if &groupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'groupId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseUnansweredGroupPhrasegroup")
	}
	// verify the required parameter 'phraseGroupId' is set
	if &phraseGroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'phraseGroupId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseUnansweredGroupPhrasegroup")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["app"] = a.Configuration.APIClient.ParameterToString(app, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Unansweredphrasegroup
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Unansweredphrasegroup" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseUnansweredGroups invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups
//
// Get knowledge base unanswered groups
func (a KnowledgeApi) GetKnowledgeKnowledgebaseUnansweredGroups(knowledgeBaseId string, app string) (*Unansweredgroups, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Unansweredgroups)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseUnansweredGroups")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["app"] = a.Configuration.APIClient.ParameterToString(app, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Unansweredgroups
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Unansweredgroups" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebases invokes GET /api/v2/knowledge/knowledgebases
//
// Get knowledge bases
func (a KnowledgeApi) GetKnowledgeKnowledgebases(before string, after string, limit string, pageSize string, name string, coreLanguage string, published bool, sortBy string, sortOrder string) (*Knowledgebaselisting, *APIResponse, error) {
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
	
	queryParams["before"] = a.Configuration.APIClient.ParameterToString(before, "")
	
	queryParams["after"] = a.Configuration.APIClient.ParameterToString(after, "")
	
	queryParams["limit"] = a.Configuration.APIClient.ParameterToString(limit, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, "")
	
	queryParams["coreLanguage"] = a.Configuration.APIClient.ParameterToString(coreLanguage, "")
	
	queryParams["published"] = a.Configuration.APIClient.ParameterToString(published, "")
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	} else if response.HasBody {
		if "Knowledgebaselisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchKnowledgeGuestSessionDocumentsSearchSearchId invokes PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}
//
// Update search result.
func (a KnowledgeApi) PatchKnowledgeGuestSessionDocumentsSearchSearchId(sessionId string, searchId string, body Searchupdaterequest) (*APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}"
	path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)
	path = strings.Replace(path, "{searchId}", fmt.Sprintf("%v", searchId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sessionId' is set
	if &sessionId == nil {
		// false
		return nil, errors.New("Missing required parameter 'sessionId' when calling KnowledgeApi->PatchKnowledgeGuestSessionDocumentsSearchSearchId")
	}
	// verify the required parameter 'searchId' is set
	if &searchId == nil {
		// false
		return nil, errors.New("Missing required parameter 'searchId' when calling KnowledgeApi->PatchKnowledgeGuestSessionDocumentsSearchSearchId")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeGuestSessionDocumentsSearchSearchId")
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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

// PatchKnowledgeKnowledgebase invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}
//
// Update knowledge base
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebase")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
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
	} else if response.HasBody {
		if "Knowledgebase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebaseCategory invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}
//
// Update category
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseCategory(knowledgeBaseId string, categoryId string, body Categoryrequest) (*Categoryresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{categoryId}", fmt.Sprintf("%v", categoryId), -1)
	defaultReturn := new(Categoryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseCategory")
	}
	// verify the required parameter 'categoryId' is set
	if &categoryId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'categoryId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseCategory")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseCategory")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Categoryresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Categoryresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebaseDocument invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}
//
// Update document.
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseDocument(knowledgeBaseId string, documentId string, body Knowledgedocumentreq) (*Knowledgedocumentresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	defaultReturn := new(Knowledgedocumentresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseDocument")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseDocument")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseDocument")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgedocumentresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebaseDocumentVariation invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}
//
// Update a variation for a document.
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseDocumentVariation(documentVariationId string, documentId string, knowledgeBaseId string, body Documentvariation) (*Documentvariation, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}"
	path = strings.Replace(path, "{documentVariationId}", fmt.Sprintf("%v", documentVariationId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Documentvariation)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'documentVariationId' is set
	if &documentVariationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentVariationId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseDocumentVariation")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseDocumentVariation")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseDocumentVariation")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseDocumentVariation")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Documentvariation
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Documentvariation" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebaseDocumentsSearchSearchId invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}
//
// Update search result.
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseDocumentsSearchSearchId(knowledgeBaseId string, searchId string, body Searchupdaterequest) (*APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{searchId}", fmt.Sprintf("%v", searchId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseDocumentsSearchSearchId")
	}
	// verify the required parameter 'searchId' is set
	if &searchId == nil {
		// false
		return nil, errors.New("Missing required parameter 'searchId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseDocumentsSearchSearchId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// PatchKnowledgeKnowledgebaseImportJob invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}
//
// Start import job
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseImportJob(knowledgeBaseId string, importJobId string, body Importstatusrequest) (*Knowledgeimportjobresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{importJobId}", fmt.Sprintf("%v", importJobId), -1)
	defaultReturn := new(Knowledgeimportjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseImportJob")
	}
	// verify the required parameter 'importJobId' is set
	if &importJobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'importJobId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseImportJob")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseImportJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgeimportjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeimportjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebaseLabel invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}
//
// Update label
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseLabel(knowledgeBaseId string, labelId string, body Labelupdaterequest) (*Labelresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{labelId}", fmt.Sprintf("%v", labelId), -1)
	defaultReturn := new(Labelresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLabel")
	}
	// verify the required parameter 'labelId' is set
	if &labelId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'labelId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLabel")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLabel")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Labelresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Labelresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebaseLanguageCategory invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}
//
// Update category
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'categoryId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageCategory")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
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
	} else if response.HasBody {
		if "Knowledgeextendedcategory" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebaseLanguageDocument invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}
//
// Update document
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocument")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
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
	} else if response.HasBody {
		if "Knowledgedocument" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebaseLanguageDocuments invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents
//
// Update documents collection
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocuments")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
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
	} else if response.HasBody {
		if "Documentlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebaseLanguageDocumentsImport invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/imports/{importId}
//
// Start import operation
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseLanguageDocumentsImport(knowledgeBaseId string, languageCode string, importId string, body Importstatusrequest) (*Knowledgeimport, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/imports/{importId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	path = strings.Replace(path, "{importId}", fmt.Sprintf("%v", importId), -1)
	defaultReturn := new(Knowledgeimport)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocumentsImport")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocumentsImport")
	}
	// verify the required parameter 'importId' is set
	if &importId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'importId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocumentsImport")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseLanguageDocumentsImport")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgeimport
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeimport" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroup invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}/phrasegroups/{phraseGroupId}
//
// Update a Knowledge base unanswered phrase group
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroup(knowledgeBaseId string, groupId string, phraseGroupId string, body Unansweredphrasegrouppatchrequestbody) (*Unansweredphrasegroupupdateresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}/phrasegroups/{phraseGroupId}"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{groupId}", fmt.Sprintf("%v", groupId), -1)
	path = strings.Replace(path, "{phraseGroupId}", fmt.Sprintf("%v", phraseGroupId), -1)
	defaultReturn := new(Unansweredphrasegroupupdateresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroup")
	}
	// verify the required parameter 'groupId' is set
	if &groupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'groupId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroup")
	}
	// verify the required parameter 'phraseGroupId' is set
	if &phraseGroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'phraseGroupId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroup")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseUnansweredGroupPhrasegroup")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Unansweredphrasegroupupdateresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Unansweredphrasegroupupdateresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeDocumentuploads invokes POST /api/v2/knowledge/documentuploads
//
// Creates a presigned URL for uploading a knowledge import file with a set of documents
func (a KnowledgeApi) PostKnowledgeDocumentuploads(body Uploadurlrequest) (*Uploadurlresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/documentuploads"
	defaultReturn := new(Uploadurlresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeDocumentuploads")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Uploadurlresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Uploadurlresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeGuestSessionDocumentsSearch invokes POST /api/v2/knowledge/guest/sessions/{sessionId}/documents/search
//
// Search the documents in a guest session.
func (a KnowledgeApi) PostKnowledgeGuestSessionDocumentsSearch(sessionId string, expand []string, body Knowledgedocumentguestsearchrequest) (*Knowledgedocumentguestsearch, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/documents/search"
	path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)
	defaultReturn := new(Knowledgedocumentguestsearch)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sessionId' is set
	if &sessionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'sessionId' when calling KnowledgeApi->PostKnowledgeGuestSessionDocumentsSearch")
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
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgedocumentguestsearch
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentguestsearch" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeGuestSessionDocumentsSearchSuggestions invokes POST /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/suggestions
//
// Query the knowledge documents to provide suggestions for auto completion.
func (a KnowledgeApi) PostKnowledgeGuestSessionDocumentsSearchSuggestions(sessionId string, body Knowledgeguestdocumentsuggestionrequest) (*Knowledgeguestdocumentsuggestion, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/documents/search/suggestions"
	path = strings.Replace(path, "{sessionId}", fmt.Sprintf("%v", sessionId), -1)
	defaultReturn := new(Knowledgeguestdocumentsuggestion)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sessionId' is set
	if &sessionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'sessionId' when calling KnowledgeApi->PostKnowledgeGuestSessionDocumentsSearchSuggestions")
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgeguestdocumentsuggestion
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeguestdocumentsuggestion" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeGuestSessions invokes POST /api/v2/knowledge/guest/sessions
//
// Create guest session
func (a KnowledgeApi) PostKnowledgeGuestSessions(body Knowledgeguestsession) (*Knowledgeguestsession, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions"
	defaultReturn := new(Knowledgeguestsession)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeGuestSessions")
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgeguestsession
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeguestsession" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseCategories invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories
//
// Create new category
func (a KnowledgeApi) PostKnowledgeKnowledgebaseCategories(knowledgeBaseId string, body Categoryrequest) (*Categoryresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Categoryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseCategories")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseCategories")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Categoryresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Categoryresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseDocumentVariations invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations
//
// Create a variation for a document.
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocumentVariations(knowledgeBaseId string, documentId string, body Documentvariation) (*Documentvariation, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	defaultReturn := new(Documentvariation)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentVariations")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentVariations")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentVariations")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Documentvariation
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Documentvariation" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseDocumentVersions invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions
//
// Creates or restores a document version.
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocumentVersions(knowledgeBaseId string, documentId string, body Knowledgedocumentversion) (*Knowledgedocumentversion, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/versions"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{documentId}", fmt.Sprintf("%v", documentId), -1)
	defaultReturn := new(Knowledgedocumentversion)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentVersions")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentVersions")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentVersions")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgedocumentversion
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentversion" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseDocuments invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents
//
// Create document.
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocuments(knowledgeBaseId string, body Knowledgedocumentreq) (*Knowledgedocumentresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Knowledgedocumentresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocuments")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocuments")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgedocumentresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseDocumentsSearch invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search
//
// Search the documents in a knowledge base.
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocumentsSearch(knowledgeBaseId string, expand []string, body Knowledgedocumentsearchrequest) (*Knowledgedocumentsearch, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Knowledgedocumentsearch)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentsSearch")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgedocumentsearch
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentsearch" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseDocumentsSearchSuggestions invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/suggestions
//
// Query the knowledge documents to provide suggestions for auto completion.
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocumentsSearchSuggestions(knowledgeBaseId string, body Knowledgedocumentsuggestionrequest) (*Knowledgedocumentsuggestion, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/suggestions"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Knowledgedocumentsuggestion)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentsSearchSuggestions")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgedocumentsuggestion
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentsuggestion" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseExportJobs invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs
//
// Create export job
func (a KnowledgeApi) PostKnowledgeKnowledgebaseExportJobs(knowledgeBaseId string, body Knowledgeexportjobrequest) (*Knowledgeexportjobresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Knowledgeexportjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseExportJobs")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseExportJobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgeexportjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeexportjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseImportJobs invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs
//
// Create import job
func (a KnowledgeApi) PostKnowledgeKnowledgebaseImportJobs(knowledgeBaseId string, body Knowledgeimportjobrequest) (*Knowledgeimportjobresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Knowledgeimportjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseImportJobs")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseImportJobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgeimportjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeimportjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseLabels invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels
//
// Create new label
func (a KnowledgeApi) PostKnowledgeKnowledgebaseLabels(knowledgeBaseId string, body Labelcreaterequest) (*Labelresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	defaultReturn := new(Labelresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLabels")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLabels")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Labelresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Labelresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseLanguageCategories invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories
//
// Create new category
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageCategories")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageCategories")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
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
	} else if response.HasBody {
		if "Knowledgeextendedcategory" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseLanguageDocuments invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents
//
// Create document
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageDocuments")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageDocuments")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
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
	} else if response.HasBody {
		if "Knowledgedocument" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseLanguageDocumentsImports invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/imports
//
// Create import operation
func (a KnowledgeApi) PostKnowledgeKnowledgebaseLanguageDocumentsImports(knowledgeBaseId string, languageCode string, body Knowledgeimport) (*Knowledgeimport, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/imports"
	path = strings.Replace(path, "{knowledgeBaseId}", fmt.Sprintf("%v", knowledgeBaseId), -1)
	path = strings.Replace(path, "{languageCode}", fmt.Sprintf("%v", languageCode), -1)
	defaultReturn := new(Knowledgeimport)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageDocumentsImports")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageDocumentsImports")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageDocumentsImports")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgeimport
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeimport" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseLanguageTrainingPromote invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings/{trainingId}/promote
//
// Promote trained documents from draft state to active.
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageTrainingPromote")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageTrainingPromote")
	}
	// verify the required parameter 'trainingId' is set
	if &trainingId == nil {
		// false
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
	} else if response.HasBody {
		if "Knowledgetraining" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseLanguageTrainings invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings
//
// Trigger training
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
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageTrainings")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
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
	} else if response.HasBody {
		if "Knowledgetraining" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseSearch invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/search
//
// Search Documents
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
		// false
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
	} else if response.HasBody {
		if "Knowledgesearchresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebases invokes POST /api/v2/knowledge/knowledgebases
//
// Create new knowledge base
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
		// false
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
	} else if response.HasBody {
		if "Knowledgebase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

