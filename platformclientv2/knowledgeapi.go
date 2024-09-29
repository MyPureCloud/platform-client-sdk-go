package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
	"time"
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{categoryId}", url.PathEscape(fmt.Sprintf("%v", categoryId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteKnowledgeKnowledgebaseDocumentVariation invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}
//
// Delete a variation for a document.
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseDocumentVariation(documentVariationId string, documentId string, knowledgeBaseId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/variations/{documentVariationId}"
	path = strings.Replace(path, "{documentVariationId}", url.PathEscape(fmt.Sprintf("%v", documentVariationId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteKnowledgeKnowledgebaseExportJob invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}
//
// Delete export job
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseExportJob(knowledgeBaseId string, exportJobId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/export/jobs/{exportJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{exportJobId}", url.PathEscape(fmt.Sprintf("%v", exportJobId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteKnowledgeKnowledgebaseImportJob invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}
//
// Delete import job
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseImportJob(knowledgeBaseId string, importJobId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{importJobId}", url.PathEscape(fmt.Sprintf("%v", importJobId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteKnowledgeKnowledgebaseLabel invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}
//
// Delete label
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseLabel(knowledgeBaseId string, labelId string) (*Labelresponse, *APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{labelId}", url.PathEscape(fmt.Sprintf("%v", labelId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: DeleteKnowledgeKnowledgebaseLanguageCategory is deprecated
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseLanguageCategory(categoryId string, knowledgeBaseId string, languageCode string) (*Knowledgecategory, *APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}"
	path = strings.Replace(path, "{categoryId}", url.PathEscape(fmt.Sprintf("%v", categoryId)), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: DeleteKnowledgeKnowledgebaseLanguageDocument is deprecated
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseLanguageDocument(documentId string, knowledgeBaseId string, languageCode string) (*Knowledgedocument, *APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}"
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: DeleteKnowledgeKnowledgebaseLanguageDocumentsImport is deprecated
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseLanguageDocumentsImport(knowledgeBaseId string, languageCode string, importId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/imports/{importId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
	path = strings.Replace(path, "{importId}", url.PathEscape(fmt.Sprintf("%v", importId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteKnowledgeKnowledgebaseSourcesSalesforceSourceId invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/salesforce/{sourceId}
//
// Delete Salesforce Knowledge integration source
//
// Preview: DeleteKnowledgeKnowledgebaseSourcesSalesforceSourceId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseSourcesSalesforceSourceId(knowledgeBaseId string, sourceId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/salesforce/{sourceId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{sourceId}", url.PathEscape(fmt.Sprintf("%v", sourceId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseSourcesSalesforceSourceId")
	}
	// verify the required parameter 'sourceId' is set
	if &sourceId == nil {
		// false
		return nil, errors.New("Missing required parameter 'sourceId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseSourcesSalesforceSourceId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteKnowledgeKnowledgebaseSourcesServicenowSourceId invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/servicenow/{sourceId}
//
// Delete ServiceNow Knowledge integration source
//
// Preview: DeleteKnowledgeKnowledgebaseSourcesServicenowSourceId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseSourcesServicenowSourceId(knowledgeBaseId string, sourceId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/servicenow/{sourceId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{sourceId}", url.PathEscape(fmt.Sprintf("%v", sourceId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseSourcesServicenowSourceId")
	}
	// verify the required parameter 'sourceId' is set
	if &sourceId == nil {
		// false
		return nil, errors.New("Missing required parameter 'sourceId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseSourcesServicenowSourceId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteKnowledgeKnowledgebaseSynchronizeJob invokes DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/synchronize/jobs/{syncJobId}
//
// Delete synchronization job
//
// Preview: DeleteKnowledgeKnowledgebaseSynchronizeJob is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) DeleteKnowledgeKnowledgebaseSynchronizeJob(knowledgeBaseId string, syncJobId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/synchronize/jobs/{syncJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{syncJobId}", url.PathEscape(fmt.Sprintf("%v", syncJobId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseSynchronizeJob")
	}
	// verify the required parameter 'syncJobId' is set
	if &syncJobId == nil {
		// false
		return nil, errors.New("Missing required parameter 'syncJobId' when calling KnowledgeApi->DeleteKnowledgeKnowledgebaseSynchronizeJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetKnowledgeGuestSessionCategories invokes GET /api/v2/knowledge/guest/sessions/{sessionId}/categories
//
// Get categories
func (a KnowledgeApi) GetKnowledgeGuestSessionCategories(sessionId string, before string, after string, pageSize string, parentId string, isRoot bool, name string, sortBy string, expand string, includeDocumentCount bool) (*Guestcategoryresponselisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/categories"
	path = strings.Replace(path, "{sessionId}", url.PathEscape(fmt.Sprintf("%v", sessionId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
func (a KnowledgeApi) GetKnowledgeGuestSessionDocument(sessionId string, documentId string) (*Knowledgeguestdocumentresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/documents/{documentId}"
	path = strings.Replace(path, "{sessionId}", url.PathEscape(fmt.Sprintf("%v", sessionId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	defaultReturn := new(Knowledgeguestdocumentresponse)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Knowledgeguestdocumentresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeguestdocumentresponse" == "string" {
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
func (a KnowledgeApi) GetKnowledgeGuestSessionDocuments(sessionId string, categoryId []string, pageSize int) (*Knowledgeguestdocumentresponselisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/documents"
	path = strings.Replace(path, "{sessionId}", url.PathEscape(fmt.Sprintf("%v", sessionId)), -1)
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
	var successPayload *Knowledgeguestdocumentresponselisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// GetKnowledgeIntegrationOptions invokes GET /api/v2/knowledge/integrations/{integrationId}/options
//
// Get sync options available for a knowledge-connect integration
//
// Preview: GetKnowledgeIntegrationOptions is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) GetKnowledgeIntegrationOptions(integrationId string) (*Knowledgeintegrationoptionsresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/integrations/{integrationId}/options"
	path = strings.Replace(path, "{integrationId}", url.PathEscape(fmt.Sprintf("%v", integrationId)), -1)
	defaultReturn := new(Knowledgeintegrationoptionsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'integrationId' is set
	if &integrationId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'integrationId' when calling KnowledgeApi->GetKnowledgeIntegrationOptions")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Knowledgeintegrationoptionsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeintegrationoptionsresponse" == "string" {
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{categoryId}", url.PathEscape(fmt.Sprintf("%v", categoryId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// GetKnowledgeKnowledgebaseDocumentFeedback invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback
//
// Get a list of feedback records given on a document
func (a KnowledgeApi) GetKnowledgeKnowledgebaseDocumentFeedback(knowledgeBaseId string, documentId string, before string, after string, pageSize string, onlyCommented bool, documentVersionId string, documentVariationId string, appType string, queryType string, userId string, queueId string, state string) (*Knowledgedocumentfeedbackresponselisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	defaultReturn := new(Knowledgedocumentfeedbackresponselisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentFeedback")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentFeedback")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	
	queryParams["onlyCommented"] = a.Configuration.APIClient.ParameterToString(onlyCommented, "")
	
	queryParams["documentVersionId"] = a.Configuration.APIClient.ParameterToString(documentVersionId, "")
	
	queryParams["documentVariationId"] = a.Configuration.APIClient.ParameterToString(documentVariationId, "")
	
	queryParams["appType"] = a.Configuration.APIClient.ParameterToString(appType, "")
	
	queryParams["queryType"] = a.Configuration.APIClient.ParameterToString(queryType, "")
	
	queryParams["userId"] = a.Configuration.APIClient.ParameterToString(userId, "")
	
	queryParams["queueId"] = a.Configuration.APIClient.ParameterToString(queueId, "")
	
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
	var successPayload *Knowledgedocumentfeedbackresponselisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentfeedbackresponselisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseDocumentFeedbackFeedbackId invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback/{feedbackId}
//
// Get a single feedback record given on a document
func (a KnowledgeApi) GetKnowledgeKnowledgebaseDocumentFeedbackFeedbackId(knowledgeBaseId string, documentId string, feedbackId string) (*Knowledgedocumentfeedbackresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback/{feedbackId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	path = strings.Replace(path, "{feedbackId}", url.PathEscape(fmt.Sprintf("%v", feedbackId)), -1)
	defaultReturn := new(Knowledgedocumentfeedbackresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentFeedbackFeedbackId")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentFeedbackFeedbackId")
	}
	// verify the required parameter 'feedbackId' is set
	if &feedbackId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'feedbackId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseDocumentFeedbackFeedbackId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Knowledgedocumentfeedbackresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentfeedbackresponse" == "string" {
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
	path = strings.Replace(path, "{documentVariationId}", url.PathEscape(fmt.Sprintf("%v", documentVariationId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	path = strings.Replace(path, "{versionId}", url.PathEscape(fmt.Sprintf("%v", versionId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	path = strings.Replace(path, "{versionId}", url.PathEscape(fmt.Sprintf("%v", versionId)), -1)
	path = strings.Replace(path, "{variationId}", url.PathEscape(fmt.Sprintf("%v", variationId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	path = strings.Replace(path, "{versionId}", url.PathEscape(fmt.Sprintf("%v", versionId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
func (a KnowledgeApi) GetKnowledgeKnowledgebaseDocuments(knowledgeBaseId string, before string, after string, pageSize string, interval string, documentId []string, categoryId []string, includeSubcategories bool, includeDrafts bool, labelIds []string, expand []string, externalIds []string) (*Knowledgedocumentresponselisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	
	queryParams["externalIds"] = a.Configuration.APIClient.ParameterToString(externalIds, "multi")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{exportJobId}", url.PathEscape(fmt.Sprintf("%v", exportJobId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
func (a KnowledgeApi) GetKnowledgeKnowledgebaseImportJob(knowledgeBaseId string, importJobId string, expand []string) (*Knowledgeimportjobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{importJobId}", url.PathEscape(fmt.Sprintf("%v", importJobId)), -1)
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
	var successPayload *Knowledgeimportjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{labelId}", url.PathEscape(fmt.Sprintf("%v", labelId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: GetKnowledgeKnowledgebaseLanguageCategories is deprecated
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageCategories(knowledgeBaseId string, languageCode string, before string, after string, limit string, pageSize string, name string) (*Categorylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: GetKnowledgeKnowledgebaseLanguageCategory is deprecated
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageCategory(categoryId string, knowledgeBaseId string, languageCode string) (*Knowledgeextendedcategory, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}"
	path = strings.Replace(path, "{categoryId}", url.PathEscape(fmt.Sprintf("%v", categoryId)), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: GetKnowledgeKnowledgebaseLanguageDocument is deprecated
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageDocument(documentId string, knowledgeBaseId string, languageCode string) (*Knowledgedocument, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}"
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// GetKnowledgeKnowledgebaseLanguageDocumentUpload invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}/uploads/{uploadId}
//
// Get document content upload status
//
// Deprecated: GetKnowledgeKnowledgebaseLanguageDocumentUpload is deprecated
//
// Preview: GetKnowledgeKnowledgebaseLanguageDocumentUpload is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageDocumentUpload(documentId string, knowledgeBaseId string, languageCode string, uploadId string) (*Knowledgedocumentcontentupload, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}/uploads/{uploadId}"
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
	path = strings.Replace(path, "{uploadId}", url.PathEscape(fmt.Sprintf("%v", uploadId)), -1)
	defaultReturn := new(Knowledgedocumentcontentupload)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocumentUpload")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocumentUpload")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocumentUpload")
	}
	// verify the required parameter 'uploadId' is set
	if &uploadId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'uploadId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseLanguageDocumentUpload")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Knowledgedocumentcontentupload
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentcontentupload" == "string" {
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
//
// Deprecated: GetKnowledgeKnowledgebaseLanguageDocuments is deprecated
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageDocuments(knowledgeBaseId string, languageCode string, before string, after string, limit string, pageSize string, categories string, title string, sortBy string, sortOrder string, documentIds []string) (*Documentlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: GetKnowledgeKnowledgebaseLanguageDocumentsImport is deprecated
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageDocumentsImport(knowledgeBaseId string, languageCode string, importId string) (*Knowledgeimport, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/imports/{importId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
	path = strings.Replace(path, "{importId}", url.PathEscape(fmt.Sprintf("%v", importId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: GetKnowledgeKnowledgebaseLanguageTraining is deprecated
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageTraining(knowledgeBaseId string, languageCode string, trainingId string) (*Knowledgetraining, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings/{trainingId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
	path = strings.Replace(path, "{trainingId}", url.PathEscape(fmt.Sprintf("%v", trainingId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: GetKnowledgeKnowledgebaseLanguageTrainings is deprecated
func (a KnowledgeApi) GetKnowledgeKnowledgebaseLanguageTrainings(knowledgeBaseId string, languageCode string, before string, after string, limit string, pageSize string, knowledgeDocumentsState string) (*Traininglisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// GetKnowledgeKnowledgebaseOperations invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/operations
//
// Get operations
//
// Preview: GetKnowledgeKnowledgebaseOperations is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) GetKnowledgeKnowledgebaseOperations(knowledgeBaseId string, before string, after string, pageSize string, userId []string, varType []string, status []string, interval string, sourceId []string) (*Operationlisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/operations"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	defaultReturn := new(Operationlisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseOperations")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	
	queryParams["userId"] = a.Configuration.APIClient.ParameterToString(userId, "multi")
	
	queryParams["varType"] = a.Configuration.APIClient.ParameterToString(varType, "multi")
	
	queryParams["status"] = a.Configuration.APIClient.ParameterToString(status, "multi")
	
	queryParams["interval"] = a.Configuration.APIClient.ParameterToString(interval, "")
	
	queryParams["sourceId"] = a.Configuration.APIClient.ParameterToString(sourceId, "multi")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Operationlisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Operationlisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseOperationsUsersQuery invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/operations/users/query
//
// Get ids of operation creator users and oauth clients
//
// Preview: GetKnowledgeKnowledgebaseOperationsUsersQuery is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) GetKnowledgeKnowledgebaseOperationsUsersQuery(knowledgeBaseId string) (*Operationcreatoruserresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/operations/users/query"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	defaultReturn := new(Operationcreatoruserresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseOperationsUsersQuery")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Operationcreatoruserresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Operationcreatoruserresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseParseJob invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/parse/jobs/{parseJobId}
//
// Get parse job report
//
// Preview: GetKnowledgeKnowledgebaseParseJob is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) GetKnowledgeKnowledgebaseParseJob(knowledgeBaseId string, parseJobId string, expand []string) (*Knowledgeparsejobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/parse/jobs/{parseJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{parseJobId}", url.PathEscape(fmt.Sprintf("%v", parseJobId)), -1)
	defaultReturn := new(Knowledgeparsejobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseParseJob")
	}
	// verify the required parameter 'parseJobId' is set
	if &parseJobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'parseJobId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseParseJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Knowledgeparsejobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeparsejobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseSources invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources
//
// Get Knowledge integration sources
//
// Preview: GetKnowledgeKnowledgebaseSources is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) GetKnowledgeKnowledgebaseSources(knowledgeBaseId string, varType string, expand []string, ids []string) ([]Sourcebaseresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	defaultReturn := make([]Sourcebaseresponse, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseSources")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["varType"] = a.Configuration.APIClient.ParameterToString(varType, "")
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	
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
	var successPayload []Sourcebaseresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "[]Sourcebaseresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseSourcesSalesforceSourceId invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/salesforce/{sourceId}
//
// Get Salesforce Knowledge integration source
//
// Preview: GetKnowledgeKnowledgebaseSourcesSalesforceSourceId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) GetKnowledgeKnowledgebaseSourcesSalesforceSourceId(knowledgeBaseId string, sourceId string, expand []string) (*Salesforcesourceresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/salesforce/{sourceId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{sourceId}", url.PathEscape(fmt.Sprintf("%v", sourceId)), -1)
	defaultReturn := new(Salesforcesourceresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseSourcesSalesforceSourceId")
	}
	// verify the required parameter 'sourceId' is set
	if &sourceId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'sourceId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseSourcesSalesforceSourceId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Salesforcesourceresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Salesforcesourceresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseSourcesServicenowSourceId invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/servicenow/{sourceId}
//
// Get ServiceNow Knowledge integration source
//
// Preview: GetKnowledgeKnowledgebaseSourcesServicenowSourceId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) GetKnowledgeKnowledgebaseSourcesServicenowSourceId(knowledgeBaseId string, sourceId string, expand []string) (*Servicenowsourceresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/servicenow/{sourceId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{sourceId}", url.PathEscape(fmt.Sprintf("%v", sourceId)), -1)
	defaultReturn := new(Servicenowsourceresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseSourcesServicenowSourceId")
	}
	// verify the required parameter 'sourceId' is set
	if &sourceId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'sourceId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseSourcesServicenowSourceId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Servicenowsourceresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Servicenowsourceresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetKnowledgeKnowledgebaseSynchronizeJob invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/synchronize/jobs/{syncJobId}
//
// Get synchronization job report
//
// Preview: GetKnowledgeKnowledgebaseSynchronizeJob is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) GetKnowledgeKnowledgebaseSynchronizeJob(knowledgeBaseId string, syncJobId string) (*Knowledgesyncjobresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/synchronize/jobs/{syncJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{syncJobId}", url.PathEscape(fmt.Sprintf("%v", syncJobId)), -1)
	defaultReturn := new(Knowledgesyncjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseSynchronizeJob")
	}
	// verify the required parameter 'syncJobId' is set
	if &syncJobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'syncJobId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseSynchronizeJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Knowledgesyncjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgesyncjobresponse" == "string" {
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
func (a KnowledgeApi) GetKnowledgeKnowledgebaseUnansweredGroup(knowledgeBaseId string, groupId string, app string, dateStart time.Time, dateEnd time.Time) (*Unansweredgroup, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{groupId}", url.PathEscape(fmt.Sprintf("%v", groupId)), -1)
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
	
	queryParams["dateStart"] = a.Configuration.APIClient.ParameterToString(dateStart, "")
	
	queryParams["dateEnd"] = a.Configuration.APIClient.ParameterToString(dateEnd, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
func (a KnowledgeApi) GetKnowledgeKnowledgebaseUnansweredGroupPhrasegroup(knowledgeBaseId string, groupId string, phraseGroupId string, app string, dateStart time.Time, dateEnd time.Time) (*Unansweredphrasegroup, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups/{groupId}/phrasegroups/{phraseGroupId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{groupId}", url.PathEscape(fmt.Sprintf("%v", groupId)), -1)
	path = strings.Replace(path, "{phraseGroupId}", url.PathEscape(fmt.Sprintf("%v", phraseGroupId)), -1)
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
	
	queryParams["dateStart"] = a.Configuration.APIClient.ParameterToString(dateStart, "")
	
	queryParams["dateEnd"] = a.Configuration.APIClient.ParameterToString(dateEnd, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
func (a KnowledgeApi) GetKnowledgeKnowledgebaseUnansweredGroups(knowledgeBaseId string, app string, dateStart time.Time, dateEnd time.Time) (*Unansweredgroups, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/unanswered/groups"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	
	queryParams["dateStart"] = a.Configuration.APIClient.ParameterToString(dateStart, "")
	
	queryParams["dateEnd"] = a.Configuration.APIClient.ParameterToString(dateEnd, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// GetKnowledgeKnowledgebaseUploadsUrlsJob invokes GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/uploads/urls/jobs/{jobId}
//
// Get content upload from URL job status
//
// Preview: GetKnowledgeKnowledgebaseUploadsUrlsJob is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) GetKnowledgeKnowledgebaseUploadsUrlsJob(knowledgeBaseId string, jobId string) (*Getuploadsourceurljobstatusresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/uploads/urls/jobs/{jobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Getuploadsourceurljobstatusresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseUploadsUrlsJob")
	}
	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling KnowledgeApi->GetKnowledgeKnowledgebaseUploadsUrlsJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Getuploadsourceurljobstatusresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Getuploadsourceurljobstatusresponse" == "string" {
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{sessionId}", url.PathEscape(fmt.Sprintf("%v", sessionId)), -1)
	path = strings.Replace(path, "{searchId}", url.PathEscape(fmt.Sprintf("%v", searchId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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

// PatchKnowledgeKnowledgebase invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}
//
// Update knowledge base
func (a KnowledgeApi) PatchKnowledgeKnowledgebase(knowledgeBaseId string, body Knowledgebaseupdaterequest) (*Knowledgebase, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseCategory(knowledgeBaseId string, categoryId string, body Categoryupdaterequest) (*Categoryresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{categoryId}", url.PathEscape(fmt.Sprintf("%v", categoryId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// PatchKnowledgeKnowledgebaseDocumentFeedbackFeedbackId invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback/{feedbackId}
//
// Update feedback on a document
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseDocumentFeedbackFeedbackId(knowledgeBaseId string, documentId string, feedbackId string, body Knowledgedocumentfeedbackupdaterequest) (*Knowledgedocumentfeedbackresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback/{feedbackId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	path = strings.Replace(path, "{feedbackId}", url.PathEscape(fmt.Sprintf("%v", feedbackId)), -1)
	defaultReturn := new(Knowledgedocumentfeedbackresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseDocumentFeedbackFeedbackId")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseDocumentFeedbackFeedbackId")
	}
	// verify the required parameter 'feedbackId' is set
	if &feedbackId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'feedbackId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseDocumentFeedbackFeedbackId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgedocumentfeedbackresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentfeedbackresponse" == "string" {
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
	path = strings.Replace(path, "{documentVariationId}", url.PathEscape(fmt.Sprintf("%v", documentVariationId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{searchId}", url.PathEscape(fmt.Sprintf("%v", searchId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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

// PatchKnowledgeKnowledgebaseImportJob invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}
//
// Start import job
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseImportJob(knowledgeBaseId string, importJobId string, body Importstatusrequest) (*Knowledgeimportjobresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{importJobId}", url.PathEscape(fmt.Sprintf("%v", importJobId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{labelId}", url.PathEscape(fmt.Sprintf("%v", labelId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: PatchKnowledgeKnowledgebaseLanguageCategory is deprecated
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseLanguageCategory(categoryId string, knowledgeBaseId string, languageCode string, body Knowledgecategoryrequest) (*Knowledgeextendedcategory, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}"
	path = strings.Replace(path, "{categoryId}", url.PathEscape(fmt.Sprintf("%v", categoryId)), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: PatchKnowledgeKnowledgebaseLanguageDocument is deprecated
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseLanguageDocument(documentId string, knowledgeBaseId string, languageCode string, body Knowledgedocumentrequest) (*Knowledgedocument, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}"
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: PatchKnowledgeKnowledgebaseLanguageDocuments is deprecated
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseLanguageDocuments(knowledgeBaseId string, languageCode string, body []Knowledgedocumentbulkrequest) (*Documentlisting, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: PatchKnowledgeKnowledgebaseLanguageDocumentsImport is deprecated
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseLanguageDocumentsImport(knowledgeBaseId string, languageCode string, importId string, body Importstatusrequest) (*Knowledgeimport, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/imports/{importId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
	path = strings.Replace(path, "{importId}", url.PathEscape(fmt.Sprintf("%v", importId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// PatchKnowledgeKnowledgebaseParseJob invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/parse/jobs/{parseJobId}
//
// Send update to the parse operation
//
// Preview: PatchKnowledgeKnowledgebaseParseJob is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseParseJob(knowledgeBaseId string, parseJobId string, body Knowledgeparsejobrequestpatch) (*APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/parse/jobs/{parseJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{parseJobId}", url.PathEscape(fmt.Sprintf("%v", parseJobId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseParseJob")
	}
	// verify the required parameter 'parseJobId' is set
	if &parseJobId == nil {
		// false
		return nil, errors.New("Missing required parameter 'parseJobId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseParseJob")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseParseJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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

// PatchKnowledgeKnowledgebaseSynchronizeJob invokes PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/synchronize/jobs/{syncJobId}
//
// Update synchronization job
//
// Preview: PatchKnowledgeKnowledgebaseSynchronizeJob is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PatchKnowledgeKnowledgebaseSynchronizeJob(knowledgeBaseId string, syncJobId string, body Syncstatusrequest) (*Knowledgesyncjobresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/synchronize/jobs/{syncJobId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{syncJobId}", url.PathEscape(fmt.Sprintf("%v", syncJobId)), -1)
	defaultReturn := new(Knowledgesyncjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseSynchronizeJob")
	}
	// verify the required parameter 'syncJobId' is set
	if &syncJobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'syncJobId' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseSynchronizeJob")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PatchKnowledgeKnowledgebaseSynchronizeJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgesyncjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgesyncjobresponse" == "string" {
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{groupId}", url.PathEscape(fmt.Sprintf("%v", groupId)), -1)
	path = strings.Replace(path, "{phraseGroupId}", url.PathEscape(fmt.Sprintf("%v", phraseGroupId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// PostKnowledgeGuestSessionDocumentCopies invokes POST /api/v2/knowledge/guest/sessions/{sessionId}/documents/{documentId}/copies
//
// Indicate that the document was copied by the user.
func (a KnowledgeApi) PostKnowledgeGuestSessionDocumentCopies(sessionId string, documentId string, body Knowledgeguestdocumentcopy) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/documents/{documentId}/copies"
	path = strings.Replace(path, "{sessionId}", url.PathEscape(fmt.Sprintf("%v", sessionId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sessionId' is set
	if &sessionId == nil {
		// false
		return nil, errors.New("Missing required parameter 'sessionId' when calling KnowledgeApi->PostKnowledgeGuestSessionDocumentCopies")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PostKnowledgeGuestSessionDocumentCopies")
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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

// PostKnowledgeGuestSessionDocumentFeedback invokes POST /api/v2/knowledge/guest/sessions/{sessionId}/documents/{documentId}/feedback
//
// Give feedback on a document
func (a KnowledgeApi) PostKnowledgeGuestSessionDocumentFeedback(sessionId string, documentId string, body Knowledgeguestdocumentfeedback) (*Knowledgeguestdocumentfeedback, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/documents/{documentId}/feedback"
	path = strings.Replace(path, "{sessionId}", url.PathEscape(fmt.Sprintf("%v", sessionId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	defaultReturn := new(Knowledgeguestdocumentfeedback)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sessionId' is set
	if &sessionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'sessionId' when calling KnowledgeApi->PostKnowledgeGuestSessionDocumentFeedback")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PostKnowledgeGuestSessionDocumentFeedback")
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgeguestdocumentfeedback
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeguestdocumentfeedback" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeGuestSessionDocumentViews invokes POST /api/v2/knowledge/guest/sessions/{sessionId}/documents/{documentId}/views
//
// Create view event for a document.
func (a KnowledgeApi) PostKnowledgeGuestSessionDocumentViews(sessionId string, documentId string, body Knowledgeguestdocumentview) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/documents/{documentId}/views"
	path = strings.Replace(path, "{sessionId}", url.PathEscape(fmt.Sprintf("%v", sessionId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sessionId' is set
	if &sessionId == nil {
		// false
		return nil, errors.New("Missing required parameter 'sessionId' when calling KnowledgeApi->PostKnowledgeGuestSessionDocumentViews")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PostKnowledgeGuestSessionDocumentViews")
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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

// PostKnowledgeGuestSessionDocumentsAnswers invokes POST /api/v2/knowledge/guest/sessions/{sessionId}/documents/answers
//
// Answer documents.
func (a KnowledgeApi) PostKnowledgeGuestSessionDocumentsAnswers(sessionId string, body Knowledgedocumentsanswerfilter) (*Knowledgeguestanswerdocumentsresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/documents/answers"
	path = strings.Replace(path, "{sessionId}", url.PathEscape(fmt.Sprintf("%v", sessionId)), -1)
	defaultReturn := new(Knowledgeguestanswerdocumentsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sessionId' is set
	if &sessionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'sessionId' when calling KnowledgeApi->PostKnowledgeGuestSessionDocumentsAnswers")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeGuestSessionDocumentsAnswers")
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgeguestanswerdocumentsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeguestanswerdocumentsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeGuestSessionDocumentsPresentations invokes POST /api/v2/knowledge/guest/sessions/{sessionId}/documents/presentations
//
// Indicate that documents were presented to the user.
func (a KnowledgeApi) PostKnowledgeGuestSessionDocumentsPresentations(sessionId string, body Knowledgeguestdocumentpresentation) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/documents/presentations"
	path = strings.Replace(path, "{sessionId}", url.PathEscape(fmt.Sprintf("%v", sessionId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sessionId' is set
	if &sessionId == nil {
		// false
		return nil, errors.New("Missing required parameter 'sessionId' when calling KnowledgeApi->PostKnowledgeGuestSessionDocumentsPresentations")
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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

// PostKnowledgeGuestSessionDocumentsSearch invokes POST /api/v2/knowledge/guest/sessions/{sessionId}/documents/search
//
// Search the documents in a guest session.
func (a KnowledgeApi) PostKnowledgeGuestSessionDocumentsSearch(sessionId string, expand []string, body Knowledgedocumentguestsearchrequest) (*Knowledgedocumentguestsearch, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/guest/sessions/{sessionId}/documents/search"
	path = strings.Replace(path, "{sessionId}", url.PathEscape(fmt.Sprintf("%v", sessionId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{sessionId}", url.PathEscape(fmt.Sprintf("%v", sessionId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
func (a KnowledgeApi) PostKnowledgeKnowledgebaseCategories(knowledgeBaseId string, body Categorycreaterequest) (*Categoryresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// PostKnowledgeKnowledgebaseDocumentCopies invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/copies
//
// Indicate that the document was copied by the user.
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocumentCopies(knowledgeBaseId string, documentId string, body Knowledgedocumentcopy) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/copies"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentCopies")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentCopies")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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

// PostKnowledgeKnowledgebaseDocumentFeedback invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback
//
// Give feedback on a document
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocumentFeedback(knowledgeBaseId string, documentId string, body Knowledgedocumentfeedback) (*Knowledgedocumentfeedbackresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/feedback"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	defaultReturn := new(Knowledgedocumentfeedbackresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentFeedback")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentFeedback")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgedocumentfeedbackresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentfeedbackresponse" == "string" {
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// PostKnowledgeKnowledgebaseDocumentViews invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/views
//
// Create view for a document.
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocumentViews(knowledgeBaseId string, documentId string, body Knowledgedocumentview) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}/views"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentViews")
	}
	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentViews")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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

// PostKnowledgeKnowledgebaseDocuments invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents
//
// Create document.
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocuments(knowledgeBaseId string, body Knowledgedocumentreq) (*Knowledgedocumentresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// PostKnowledgeKnowledgebaseDocumentsAnswers invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/answers
//
// Answer documents.
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocumentsAnswers(knowledgeBaseId string, body Knowledgedocumentsanswerfilter) (*Knowledgeanswerdocumentsresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/answers"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	defaultReturn := new(Knowledgeanswerdocumentsresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentsAnswers")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentsAnswers")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgeanswerdocumentsresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeanswerdocumentsresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseDocumentsBulkRemove invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/bulk/remove
//
// Bulk remove documents.
//
// Preview: PostKnowledgeKnowledgebaseDocumentsBulkRemove is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocumentsBulkRemove(knowledgeBaseId string, body Knowledgedocumentbulkremoverequest) (*Bulkresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/bulk/remove"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	defaultReturn := new(Bulkresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentsBulkRemove")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentsBulkRemove")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Bulkresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulkresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseDocumentsBulkUpdate invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/bulk/update
//
// Bulk update documents.
//
// Preview: PostKnowledgeKnowledgebaseDocumentsBulkUpdate is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocumentsBulkUpdate(knowledgeBaseId string, body Knowledgedocumentbulkupdaterequest) (*Bulkresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/bulk/update"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	defaultReturn := new(Bulkresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentsBulkUpdate")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentsBulkUpdate")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Bulkresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulkresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseDocumentsPresentations invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/presentations
//
// Indicate that documents were presented to the user.
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocumentsPresentations(knowledgeBaseId string, body Knowledgedocumentpresentation) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/presentations"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentsPresentations")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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

// PostKnowledgeKnowledgebaseDocumentsQuery invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/query
//
// Query for knowledge documents.
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocumentsQuery(knowledgeBaseId string, expand []string, body Knowledgedocumentquery) (*Knowledgedocumentqueryresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/query"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	defaultReturn := new(Knowledgedocumentqueryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentsQuery")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Knowledgedocumentqueryresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentqueryresponse" == "string" {
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// PostKnowledgeKnowledgebaseDocumentsVersionsBulkAdd invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/versions/bulk/add
//
// Bulk add document versions.
//
// Preview: PostKnowledgeKnowledgebaseDocumentsVersionsBulkAdd is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PostKnowledgeKnowledgebaseDocumentsVersionsBulkAdd(knowledgeBaseId string, body Knowledgedocumentbulkversionaddrequest) (*Bulkresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/versions/bulk/add"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	defaultReturn := new(Bulkresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentsVersionsBulkAdd")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseDocumentsVersionsBulkAdd")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Bulkresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Bulkresponse" == "string" {
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: PostKnowledgeKnowledgebaseLanguageCategories is deprecated
func (a KnowledgeApi) PostKnowledgeKnowledgebaseLanguageCategories(knowledgeBaseId string, languageCode string, body Knowledgecategoryrequest) (*Knowledgeextendedcategory, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// PostKnowledgeKnowledgebaseLanguageDocumentUploads invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}/uploads
//
// Upload Article Content
//
// Deprecated: PostKnowledgeKnowledgebaseLanguageDocumentUploads is deprecated
//
// Preview: PostKnowledgeKnowledgebaseLanguageDocumentUploads is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PostKnowledgeKnowledgebaseLanguageDocumentUploads(documentId string, knowledgeBaseId string, languageCode string, body Knowledgedocumentcontentupload) (*Knowledgedocumentcontentupload, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}/uploads"
	path = strings.Replace(path, "{documentId}", url.PathEscape(fmt.Sprintf("%v", documentId)), -1)
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
	defaultReturn := new(Knowledgedocumentcontentupload)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'documentId' is set
	if &documentId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'documentId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageDocumentUploads")
	}
	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageDocumentUploads")
	}
	// verify the required parameter 'languageCode' is set
	if &languageCode == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'languageCode' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageDocumentUploads")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseLanguageDocumentUploads")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgedocumentcontentupload
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgedocumentcontentupload" == "string" {
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
//
// Deprecated: PostKnowledgeKnowledgebaseLanguageDocuments is deprecated
func (a KnowledgeApi) PostKnowledgeKnowledgebaseLanguageDocuments(knowledgeBaseId string, languageCode string, body Knowledgedocumentrequest) (*Knowledgedocument, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: PostKnowledgeKnowledgebaseLanguageDocumentsImports is deprecated
func (a KnowledgeApi) PostKnowledgeKnowledgebaseLanguageDocumentsImports(knowledgeBaseId string, languageCode string, body Knowledgeimport) (*Knowledgeimport, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/imports"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: PostKnowledgeKnowledgebaseLanguageTrainingPromote is deprecated
func (a KnowledgeApi) PostKnowledgeKnowledgebaseLanguageTrainingPromote(knowledgeBaseId string, languageCode string, trainingId string) (*Knowledgetraining, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings/{trainingId}/promote"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
	path = strings.Replace(path, "{trainingId}", url.PathEscape(fmt.Sprintf("%v", trainingId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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
//
// Deprecated: PostKnowledgeKnowledgebaseLanguageTrainings is deprecated
func (a KnowledgeApi) PostKnowledgeKnowledgebaseLanguageTrainings(knowledgeBaseId string, languageCode string) (*Knowledgetraining, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{languageCode}", url.PathEscape(fmt.Sprintf("%v", languageCode)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// PostKnowledgeKnowledgebaseParseJobImport invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/parse/jobs/{parseJobId}/import
//
// Import the parsed articles
//
// Preview: PostKnowledgeKnowledgebaseParseJobImport is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PostKnowledgeKnowledgebaseParseJobImport(knowledgeBaseId string, parseJobId string, body Knowledgeparsejobrequestimport) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/parse/jobs/{parseJobId}/import"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{parseJobId}", url.PathEscape(fmt.Sprintf("%v", parseJobId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseParseJobImport")
	}
	// verify the required parameter 'parseJobId' is set
	if &parseJobId == nil {
		// false
		return nil, errors.New("Missing required parameter 'parseJobId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseParseJobImport")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseParseJobImport")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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

// PostKnowledgeKnowledgebaseParseJobs invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/parse/jobs
//
// Create parse job
//
// Preview: PostKnowledgeKnowledgebaseParseJobs is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PostKnowledgeKnowledgebaseParseJobs(knowledgeBaseId string, body Knowledgeparsejobrequest) (*Knowledgeparsejobresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/parse/jobs"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	defaultReturn := new(Knowledgeparsejobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseParseJobs")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseParseJobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgeparsejobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgeparsejobresponse" == "string" {
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
//
// Deprecated: PostKnowledgeKnowledgebaseSearch is deprecated
func (a KnowledgeApi) PostKnowledgeKnowledgebaseSearch(knowledgeBaseId string, body Knowledgesearchrequest) (*Knowledgesearchresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/search"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// PostKnowledgeKnowledgebaseSourcesSalesforce invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/salesforce
//
// Create Salesforce Knowledge integration source
//
// Preview: PostKnowledgeKnowledgebaseSourcesSalesforce is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PostKnowledgeKnowledgebaseSourcesSalesforce(knowledgeBaseId string, body Salesforcesourcerequest) (*Knowledgesyncjobresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/salesforce"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	defaultReturn := new(Knowledgesyncjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseSourcesSalesforce")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseSourcesSalesforce")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgesyncjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgesyncjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseSourcesSalesforceSourceIdSync invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/salesforce/{sourceId}/sync
//
// Start sync on Salesforce Knowledge integration source
//
// Preview: PostKnowledgeKnowledgebaseSourcesSalesforceSourceIdSync is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PostKnowledgeKnowledgebaseSourcesSalesforceSourceIdSync(knowledgeBaseId string, sourceId string) (*Sourcesyncresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/salesforce/{sourceId}/sync"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{sourceId}", url.PathEscape(fmt.Sprintf("%v", sourceId)), -1)
	defaultReturn := new(Sourcesyncresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseSourcesSalesforceSourceIdSync")
	}
	// verify the required parameter 'sourceId' is set
	if &sourceId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'sourceId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseSourcesSalesforceSourceIdSync")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Sourcesyncresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Sourcesyncresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseSourcesServicenow invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/servicenow
//
// Create ServiceNow Knowledge integration source
//
// Preview: PostKnowledgeKnowledgebaseSourcesServicenow is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PostKnowledgeKnowledgebaseSourcesServicenow(knowledgeBaseId string, body Servicenowsourcerequest) (*Knowledgesyncjobresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/servicenow"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	defaultReturn := new(Knowledgesyncjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseSourcesServicenow")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseSourcesServicenow")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgesyncjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgesyncjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseSourcesServicenowSourceIdSync invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/servicenow/{sourceId}/sync
//
// Start synchronization on ServiceNow Knowledge integration source
//
// Preview: PostKnowledgeKnowledgebaseSourcesServicenowSourceIdSync is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PostKnowledgeKnowledgebaseSourcesServicenowSourceIdSync(knowledgeBaseId string, sourceId string) (*Sourcesyncresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/servicenow/{sourceId}/sync"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{sourceId}", url.PathEscape(fmt.Sprintf("%v", sourceId)), -1)
	defaultReturn := new(Sourcesyncresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseSourcesServicenowSourceIdSync")
	}
	// verify the required parameter 'sourceId' is set
	if &sourceId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'sourceId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseSourcesServicenowSourceIdSync")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Sourcesyncresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Sourcesyncresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseSynchronizeJobs invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/synchronize/jobs
//
// Create synchronization job
//
// Preview: PostKnowledgeKnowledgebaseSynchronizeJobs is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PostKnowledgeKnowledgebaseSynchronizeJobs(knowledgeBaseId string, body Knowledgesyncjobrequest) (*Knowledgesyncjobresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/synchronize/jobs"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	defaultReturn := new(Knowledgesyncjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseSynchronizeJobs")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseSynchronizeJobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Knowledgesyncjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Knowledgesyncjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostKnowledgeKnowledgebaseUploadsUrlsJobs invokes POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/uploads/urls/jobs
//
// Create content upload from URL job
//
// Preview: PostKnowledgeKnowledgebaseUploadsUrlsJobs is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PostKnowledgeKnowledgebaseUploadsUrlsJobs(knowledgeBaseId string, body Createuploadsourceurljobrequest) (*Createuploadsourceurljobresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/uploads/urls/jobs"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	defaultReturn := new(Createuploadsourceurljobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PostKnowledgeKnowledgebaseUploadsUrlsJobs")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PostKnowledgeKnowledgebaseUploadsUrlsJobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Createuploadsourceurljobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Createuploadsourceurljobresponse" == "string" {
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
func (a KnowledgeApi) PostKnowledgeKnowledgebases(body Knowledgebasecreaterequest) (*Knowledgebase, *APIResponse, error) {
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
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
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
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
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

// PutKnowledgeKnowledgebaseSourcesSalesforceSourceId invokes PUT /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/salesforce/{sourceId}
//
// Update Salesforce Knowledge integration source
//
// Preview: PutKnowledgeKnowledgebaseSourcesSalesforceSourceId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PutKnowledgeKnowledgebaseSourcesSalesforceSourceId(knowledgeBaseId string, sourceId string, body Salesforcesourcerequest) (*Salesforcesourceresponse, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/salesforce/{sourceId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{sourceId}", url.PathEscape(fmt.Sprintf("%v", sourceId)), -1)
	defaultReturn := new(Salesforcesourceresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PutKnowledgeKnowledgebaseSourcesSalesforceSourceId")
	}
	// verify the required parameter 'sourceId' is set
	if &sourceId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'sourceId' when calling KnowledgeApi->PutKnowledgeKnowledgebaseSourcesSalesforceSourceId")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PutKnowledgeKnowledgebaseSourcesSalesforceSourceId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Salesforcesourceresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Salesforcesourceresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutKnowledgeKnowledgebaseSourcesServicenowSourceId invokes PUT /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/servicenow/{sourceId}
//
// Update ServiceNow Knowledge integration source
//
// Preview: PutKnowledgeKnowledgebaseSourcesServicenowSourceId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a KnowledgeApi) PutKnowledgeKnowledgebaseSourcesServicenowSourceId(knowledgeBaseId string, sourceId string, body Servicenowsourcerequest) (*Servicenowsourceresponse, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/knowledge/knowledgebases/{knowledgeBaseId}/sources/servicenow/{sourceId}"
	path = strings.Replace(path, "{knowledgeBaseId}", url.PathEscape(fmt.Sprintf("%v", knowledgeBaseId)), -1)
	path = strings.Replace(path, "{sourceId}", url.PathEscape(fmt.Sprintf("%v", sourceId)), -1)
	defaultReturn := new(Servicenowsourceresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'knowledgeBaseId' is set
	if &knowledgeBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'knowledgeBaseId' when calling KnowledgeApi->PutKnowledgeKnowledgebaseSourcesServicenowSourceId")
	}
	// verify the required parameter 'sourceId' is set
	if &sourceId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'sourceId' when calling KnowledgeApi->PutKnowledgeKnowledgebaseSourcesServicenowSourceId")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling KnowledgeApi->PutKnowledgeKnowledgebaseSourcesServicenowSourceId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Servicenowsourceresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Servicenowsourceresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

