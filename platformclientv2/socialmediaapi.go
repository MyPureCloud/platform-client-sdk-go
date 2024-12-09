package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
	)

// SocialMediaApi provides functions for API endpoints
type SocialMediaApi struct {
	Configuration *Configuration
}

// NewSocialMediaApi creates an API instance using the default configuration
func NewSocialMediaApi() *SocialMediaApi {
	fmt.Sprintf(strings.Title(""), "")
	config := GetDefaultConfiguration()
	return &SocialMediaApi{
		Configuration: config,
	}
}

// NewSocialMediaApiWithConfig creates an API instance using the provided configuration
func NewSocialMediaApiWithConfig(config *Configuration) *SocialMediaApi {
	return &SocialMediaApi{
		Configuration: config,
	}
}

// DeleteSocialmediaTopic invokes DELETE /api/v2/socialmedia/topics/{topicId}
//
// Delete a social topic.
//
// Preview: DeleteSocialmediaTopic is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) DeleteSocialmediaTopic(topicId string, hardDelete bool) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->DeleteSocialmediaTopic")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["hardDelete"] = a.Configuration.APIClient.ParameterToString(hardDelete, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId invokes DELETE /api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}
//
// Delete a Facebook data ingestion rule.
//
// Preview: DeleteSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) DeleteSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId(topicId string, facebookIngestionRuleId string, hardDelete bool) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{facebookIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", facebookIngestionRuleId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->DeleteSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId")
	}
	// verify the required parameter 'facebookIngestionRuleId' is set
	if &facebookIngestionRuleId == nil {
		// false
		return nil, errors.New("Missing required parameter 'facebookIngestionRuleId' when calling SocialMediaApi->DeleteSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["hardDelete"] = a.Configuration.APIClient.ParameterToString(hardDelete, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteSocialmediaTopicDataingestionrulesOpenOpenId invokes DELETE /api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{openId}
//
// Delete a open data ingestion rule.
//
// Preview: DeleteSocialmediaTopicDataingestionrulesOpenOpenId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) DeleteSocialmediaTopicDataingestionrulesOpenOpenId(topicId string, openId string, hardDelete bool) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{openId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{openId}", url.PathEscape(fmt.Sprintf("%v", openId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->DeleteSocialmediaTopicDataingestionrulesOpenOpenId")
	}
	// verify the required parameter 'openId' is set
	if &openId == nil {
		// false
		return nil, errors.New("Missing required parameter 'openId' when calling SocialMediaApi->DeleteSocialmediaTopicDataingestionrulesOpenOpenId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["hardDelete"] = a.Configuration.APIClient.ParameterToString(hardDelete, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId invokes DELETE /api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}
//
// Delete a X (formally Twitter) data ingestion rule.
//
// Preview: DeleteSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) DeleteSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId(topicId string, twitterIngestionRuleId string, hardDelete bool) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{twitterIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", twitterIngestionRuleId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->DeleteSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId")
	}
	// verify the required parameter 'twitterIngestionRuleId' is set
	if &twitterIngestionRuleId == nil {
		// false
		return nil, errors.New("Missing required parameter 'twitterIngestionRuleId' when calling SocialMediaApi->DeleteSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["hardDelete"] = a.Configuration.APIClient.ParameterToString(hardDelete, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// GetSocialmediaTopic invokes GET /api/v2/socialmedia/topics/{topicId}
//
// Get a single social topic.
//
// Preview: GetSocialmediaTopic is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) GetSocialmediaTopic(topicId string, includeDeleted bool) (*Socialtopicresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	defaultReturn := new(Socialtopicresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopic")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["includeDeleted"] = a.Configuration.APIClient.ParameterToString(includeDeleted, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Socialtopicresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Socialtopicresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId invokes GET /api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}
//
// Get a single Facebook data ingestion rule.
//
// Preview: GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId(topicId string, facebookIngestionRuleId string, includeDeleted bool) (*Facebookdataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{facebookIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", facebookIngestionRuleId)), -1)
	defaultReturn := new(Facebookdataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId")
	}
	// verify the required parameter 'facebookIngestionRuleId' is set
	if &facebookIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'facebookIngestionRuleId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["includeDeleted"] = a.Configuration.APIClient.ParameterToString(includeDeleted, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Facebookdataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Facebookdataingestionruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleIdVersion invokes GET /api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}/versions/{dataIngestionRuleVersion}
//
// Get a single Facebook data ingestion rule version.
//
// Preview: GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleIdVersion is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleIdVersion(topicId string, facebookIngestionRuleId string, dataIngestionRuleVersion string, includeDeleted bool) (*Facebookdataingestionruleversionresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}/versions/{dataIngestionRuleVersion}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{facebookIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", facebookIngestionRuleId)), -1)
	path = strings.Replace(path, "{dataIngestionRuleVersion}", url.PathEscape(fmt.Sprintf("%v", dataIngestionRuleVersion)), -1)
	defaultReturn := new(Facebookdataingestionruleversionresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleIdVersion")
	}
	// verify the required parameter 'facebookIngestionRuleId' is set
	if &facebookIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'facebookIngestionRuleId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleIdVersion")
	}
	// verify the required parameter 'dataIngestionRuleVersion' is set
	if &dataIngestionRuleVersion == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'dataIngestionRuleVersion' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleIdVersion")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["includeDeleted"] = a.Configuration.APIClient.ParameterToString(includeDeleted, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Facebookdataingestionruleversionresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Facebookdataingestionruleversionresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleIdVersions invokes GET /api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}/versions
//
// Get the Facebook data ingestion rule versions.
//
// Preview: GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleIdVersions is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleIdVersions(topicId string, facebookIngestionRuleId string, pageNumber int, pageSize int, includeDeleted bool) (*Facebookdataingestionruleversionresponseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}/versions"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{facebookIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", facebookIngestionRuleId)), -1)
	defaultReturn := new(Facebookdataingestionruleversionresponseentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleIdVersions")
	}
	// verify the required parameter 'facebookIngestionRuleId' is set
	if &facebookIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'facebookIngestionRuleId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleIdVersions")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["includeDeleted"] = a.Configuration.APIClient.ParameterToString(includeDeleted, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Facebookdataingestionruleversionresponseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Facebookdataingestionruleversionresponseentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopicDataingestionrulesOpenOpenId invokes GET /api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{openId}
//
// Get a single open data ingestion rule.
//
// Preview: GetSocialmediaTopicDataingestionrulesOpenOpenId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) GetSocialmediaTopicDataingestionrulesOpenOpenId(topicId string, openId string, includeDeleted bool) (*Opendataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{openId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{openId}", url.PathEscape(fmt.Sprintf("%v", openId)), -1)
	defaultReturn := new(Opendataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesOpenOpenId")
	}
	// verify the required parameter 'openId' is set
	if &openId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'openId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesOpenOpenId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["includeDeleted"] = a.Configuration.APIClient.ParameterToString(includeDeleted, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Opendataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Opendataingestionruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopicDataingestionrulesOpenOpenIdVersion invokes GET /api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{openId}/versions/{dataIngestionRuleVersion}
//
// Get a single Open data ingestion rule version.
//
// Preview: GetSocialmediaTopicDataingestionrulesOpenOpenIdVersion is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) GetSocialmediaTopicDataingestionrulesOpenOpenIdVersion(topicId string, openId string, dataIngestionRuleVersion string, includeDeleted bool) (*Opendataingestionruleversionresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{openId}/versions/{dataIngestionRuleVersion}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{openId}", url.PathEscape(fmt.Sprintf("%v", openId)), -1)
	path = strings.Replace(path, "{dataIngestionRuleVersion}", url.PathEscape(fmt.Sprintf("%v", dataIngestionRuleVersion)), -1)
	defaultReturn := new(Opendataingestionruleversionresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesOpenOpenIdVersion")
	}
	// verify the required parameter 'openId' is set
	if &openId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'openId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesOpenOpenIdVersion")
	}
	// verify the required parameter 'dataIngestionRuleVersion' is set
	if &dataIngestionRuleVersion == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'dataIngestionRuleVersion' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesOpenOpenIdVersion")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["includeDeleted"] = a.Configuration.APIClient.ParameterToString(includeDeleted, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Opendataingestionruleversionresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Opendataingestionruleversionresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopicDataingestionrulesOpenOpenIdVersions invokes GET /api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{openId}/versions
//
// Get the Open data ingestion rule versions.
//
// Preview: GetSocialmediaTopicDataingestionrulesOpenOpenIdVersions is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) GetSocialmediaTopicDataingestionrulesOpenOpenIdVersions(topicId string, openId string, pageNumber int, pageSize int, includeDeleted bool) (*Opendataingestionruleversionresponseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{openId}/versions"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{openId}", url.PathEscape(fmt.Sprintf("%v", openId)), -1)
	defaultReturn := new(Opendataingestionruleversionresponseentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesOpenOpenIdVersions")
	}
	// verify the required parameter 'openId' is set
	if &openId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'openId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesOpenOpenIdVersions")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["includeDeleted"] = a.Configuration.APIClient.ParameterToString(includeDeleted, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Opendataingestionruleversionresponseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Opendataingestionruleversionresponseentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId invokes GET /api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}
//
// Get a single X (formally Twitter) data ingestion rule.
//
// Preview: GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId(topicId string, twitterIngestionRuleId string, includeDeleted bool) (*Twitterdataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{twitterIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", twitterIngestionRuleId)), -1)
	defaultReturn := new(Twitterdataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId")
	}
	// verify the required parameter 'twitterIngestionRuleId' is set
	if &twitterIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'twitterIngestionRuleId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["includeDeleted"] = a.Configuration.APIClient.ParameterToString(includeDeleted, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Twitterdataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Twitterdataingestionruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleIdVersion invokes GET /api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}/versions/{dataIngestionRuleVersion}
//
// Get a single X (formally Twitter) data ingestion rule version.
//
// Preview: GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleIdVersion is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleIdVersion(topicId string, twitterIngestionRuleId string, dataIngestionRuleVersion string, includeDeleted bool) (*Twitterdataingestionruleversionresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}/versions/{dataIngestionRuleVersion}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{twitterIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", twitterIngestionRuleId)), -1)
	path = strings.Replace(path, "{dataIngestionRuleVersion}", url.PathEscape(fmt.Sprintf("%v", dataIngestionRuleVersion)), -1)
	defaultReturn := new(Twitterdataingestionruleversionresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleIdVersion")
	}
	// verify the required parameter 'twitterIngestionRuleId' is set
	if &twitterIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'twitterIngestionRuleId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleIdVersion")
	}
	// verify the required parameter 'dataIngestionRuleVersion' is set
	if &dataIngestionRuleVersion == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'dataIngestionRuleVersion' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleIdVersion")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["includeDeleted"] = a.Configuration.APIClient.ParameterToString(includeDeleted, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Twitterdataingestionruleversionresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Twitterdataingestionruleversionresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleIdVersions invokes GET /api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}/versions
//
// Get the Open data ingestion rule versions.
//
// Preview: GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleIdVersions is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleIdVersions(topicId string, twitterIngestionRuleId string, pageNumber int, pageSize int, includeDeleted bool) (*Twitterdataingestionruleversionresponseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}/versions"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{twitterIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", twitterIngestionRuleId)), -1)
	defaultReturn := new(Twitterdataingestionruleversionresponseentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleIdVersions")
	}
	// verify the required parameter 'twitterIngestionRuleId' is set
	if &twitterIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'twitterIngestionRuleId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleIdVersions")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["includeDeleted"] = a.Configuration.APIClient.ParameterToString(includeDeleted, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Twitterdataingestionruleversionresponseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Twitterdataingestionruleversionresponseentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopics invokes GET /api/v2/socialmedia/topics
//
// Retrieve all social topics.
//
// Preview: GetSocialmediaTopics is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) GetSocialmediaTopics(pageNumber int, pageSize int, divisionIds []string, includeDeleted bool) (*Socialtopicresponseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics"
	defaultReturn := new(Socialtopicresponseentitylisting)
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
	
	queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["divisionIds"] = a.Configuration.APIClient.ParameterToString(divisionIds, "multi")
	
	queryParams["includeDeleted"] = a.Configuration.APIClient.ParameterToString(includeDeleted, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Socialtopicresponseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Socialtopicresponseentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchSocialmediaTopic invokes PATCH /api/v2/socialmedia/topics/{topicId}
//
// Update a social topic.
//
// Preview: PatchSocialmediaTopic is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) PatchSocialmediaTopic(topicId string, body Socialtopicpatchrequest) (*Socialtopicresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	defaultReturn := new(Socialtopicresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PatchSocialmediaTopic")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Socialtopicresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Socialtopicresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId invokes PATCH /api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}
//
// Update the status of a Facebook data ingestion rule.
//
// Preview: PatchSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) PatchSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId(topicId string, facebookIngestionRuleId string, body Dataingestionrulestatuspatchrequest) (*Facebookdataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{facebookIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", facebookIngestionRuleId)), -1)
	defaultReturn := new(Facebookdataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PatchSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId")
	}
	// verify the required parameter 'facebookIngestionRuleId' is set
	if &facebookIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'facebookIngestionRuleId' when calling SocialMediaApi->PatchSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Facebookdataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Facebookdataingestionruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchSocialmediaTopicDataingestionrulesOpenOpenId invokes PATCH /api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{openId}
//
// Update the status of a open data ingestion rule.
//
// Preview: PatchSocialmediaTopicDataingestionrulesOpenOpenId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) PatchSocialmediaTopicDataingestionrulesOpenOpenId(topicId string, openId string, body Dataingestionrulestatuspatchrequest) (*Opendataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{openId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{openId}", url.PathEscape(fmt.Sprintf("%v", openId)), -1)
	defaultReturn := new(Opendataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PatchSocialmediaTopicDataingestionrulesOpenOpenId")
	}
	// verify the required parameter 'openId' is set
	if &openId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'openId' when calling SocialMediaApi->PatchSocialmediaTopicDataingestionrulesOpenOpenId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Opendataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Opendataingestionruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PatchSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId invokes PATCH /api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}
//
// Update the status of a X (formally Twitter) data ingestion rule.
//
// Preview: PatchSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) PatchSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId(topicId string, twitterIngestionRuleId string, body Dataingestionrulestatuspatchrequest) (*Twitterdataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{twitterIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", twitterIngestionRuleId)), -1)
	defaultReturn := new(Twitterdataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PatchSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId")
	}
	// verify the required parameter 'twitterIngestionRuleId' is set
	if &twitterIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'twitterIngestionRuleId' when calling SocialMediaApi->PatchSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Twitterdataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Twitterdataingestionruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSocialmediaTopicDataingestionrulesFacebook invokes POST /api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook
//
// Create an Facebook data ingestion rule.
//
// Preview: PostSocialmediaTopicDataingestionrulesFacebook is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) PostSocialmediaTopicDataingestionrulesFacebook(topicId string, body Facebookdataingestionrulerequest) (*Facebookdataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	defaultReturn := new(Facebookdataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PostSocialmediaTopicDataingestionrulesFacebook")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Facebookdataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Facebookdataingestionruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSocialmediaTopicDataingestionrulesOpen invokes POST /api/v2/socialmedia/topics/{topicId}/dataingestionrules/open
//
// Create an open data ingestion rule.
//
// Preview: PostSocialmediaTopicDataingestionrulesOpen is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) PostSocialmediaTopicDataingestionrulesOpen(topicId string, body Opendataingestionrulerequest) (*Opendataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/open"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	defaultReturn := new(Opendataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PostSocialmediaTopicDataingestionrulesOpen")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Opendataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Opendataingestionruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSocialmediaTopicDataingestionrulesTwitter invokes POST /api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter
//
// Create an twitter data ingestion rule.
//
// Preview: PostSocialmediaTopicDataingestionrulesTwitter is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) PostSocialmediaTopicDataingestionrulesTwitter(topicId string, body Twitterdataingestionrulerequest) (*Twitterdataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	defaultReturn := new(Twitterdataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PostSocialmediaTopicDataingestionrulesTwitter")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Twitterdataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Twitterdataingestionruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSocialmediaTopics invokes POST /api/v2/socialmedia/topics
//
// Create a social topic.
//
// Preview: PostSocialmediaTopics is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) PostSocialmediaTopics(body Socialtopicrequest) (*Socialtopicresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics"
	defaultReturn := new(Socialtopicresponse)
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

	var successPayload *Socialtopicresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Socialtopicresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId invokes PUT /api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}
//
// Update the Facebook data ingestion rule.
//
// Preview: PutSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) PutSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId(topicId string, facebookIngestionRuleId string, body Facebookdataingestionrulerequest) (*Facebookdataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/facebook/{facebookIngestionRuleId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{facebookIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", facebookIngestionRuleId)), -1)
	defaultReturn := new(Facebookdataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PutSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId")
	}
	// verify the required parameter 'facebookIngestionRuleId' is set
	if &facebookIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'facebookIngestionRuleId' when calling SocialMediaApi->PutSocialmediaTopicDataingestionrulesFacebookFacebookIngestionRuleId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Facebookdataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Facebookdataingestionruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutSocialmediaTopicDataingestionrulesOpenOpenId invokes PUT /api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{openId}
//
// Update the open data ingestion rule.
//
// Preview: PutSocialmediaTopicDataingestionrulesOpenOpenId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) PutSocialmediaTopicDataingestionrulesOpenOpenId(topicId string, openId string, body Opendataingestionrulerequest) (*Opendataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{openId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{openId}", url.PathEscape(fmt.Sprintf("%v", openId)), -1)
	defaultReturn := new(Opendataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PutSocialmediaTopicDataingestionrulesOpenOpenId")
	}
	// verify the required parameter 'openId' is set
	if &openId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'openId' when calling SocialMediaApi->PutSocialmediaTopicDataingestionrulesOpenOpenId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Opendataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Opendataingestionruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId invokes PUT /api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}
//
// Update the X (formally Twitter) data ingestion rule.
//
// Preview: PutSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) PutSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId(topicId string, twitterIngestionRuleId string, body Twitterdataingestionrulerequest) (*Twitterdataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/twitter/{twitterIngestionRuleId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{twitterIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", twitterIngestionRuleId)), -1)
	defaultReturn := new(Twitterdataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PutSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId")
	}
	// verify the required parameter 'twitterIngestionRuleId' is set
	if &twitterIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'twitterIngestionRuleId' when calling SocialMediaApi->PutSocialmediaTopicDataingestionrulesTwitterTwitterIngestionRuleId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Twitterdataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Twitterdataingestionruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

