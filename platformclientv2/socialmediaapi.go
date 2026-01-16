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

// DeleteSocialmediaEscalationrule invokes DELETE /api/v2/socialmedia/escalationrules/{escalationRuleId}
//
// Delete an escalation rule.
func (a SocialMediaApi) DeleteSocialmediaEscalationrule(escalationRuleId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/escalationrules/{escalationRuleId}"
	path = strings.Replace(path, "{escalationRuleId}", url.PathEscape(fmt.Sprintf("%v", escalationRuleId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'escalationRuleId' is set
	if &escalationRuleId == nil {
		// false
		return nil, errors.New("Missing required parameter 'escalationRuleId' when calling SocialMediaApi->DeleteSocialmediaEscalationrule")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteSocialmediaMessage invokes DELETE /api/v2/socialmedia/messages/{messageId}
//
// Delete a social media message.
//
// Preview: DeleteSocialmediaMessage is a preview method and is subject to both breaking and non-breaking changes at any time without notice
func (a SocialMediaApi) DeleteSocialmediaMessage(messageId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/messages/{messageId}"
	path = strings.Replace(path, "{messageId}", url.PathEscape(fmt.Sprintf("%v", messageId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'messageId' is set
	if &messageId == nil {
		// false
		return nil, errors.New("Missing required parameter 'messageId' when calling SocialMediaApi->DeleteSocialmediaMessage")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
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

// DeleteSocialmediaTopic invokes DELETE /api/v2/socialmedia/topics/{topicId}
//
// Delete a social topic.
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

// DeleteSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId invokes DELETE /api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}
//
// Delete a Instagram data ingestion rule.
func (a SocialMediaApi) DeleteSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId(topicId string, instagramIngestionRuleId string, hardDelete bool) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{instagramIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", instagramIngestionRuleId)), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->DeleteSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId")
	}
	// verify the required parameter 'instagramIngestionRuleId' is set
	if &instagramIngestionRuleId == nil {
		// false
		return nil, errors.New("Missing required parameter 'instagramIngestionRuleId' when calling SocialMediaApi->DeleteSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// GetSocialmediaAnalyticsAggregatesJob invokes GET /api/v2/socialmedia/analytics/aggregates/jobs/{jobId}
//
// Get status for async query for social media aggregates
func (a SocialMediaApi) GetSocialmediaAnalyticsAggregatesJob(jobId string) (*Asyncquerystatus, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/analytics/aggregates/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Asyncquerystatus)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling SocialMediaApi->GetSocialmediaAnalyticsAggregatesJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Asyncquerystatus
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Asyncquerystatus" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaAnalyticsAggregatesJobResults invokes GET /api/v2/socialmedia/analytics/aggregates/jobs/{jobId}/results
//
// Fetch a page of results for an async social media query
func (a SocialMediaApi) GetSocialmediaAnalyticsAggregatesJobResults(jobId string, cursor string) (*Socialmediaasyncaggregatequeryresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/analytics/aggregates/jobs/{jobId}/results"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Socialmediaasyncaggregatequeryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling SocialMediaApi->GetSocialmediaAnalyticsAggregatesJobResults")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["cursor"] = a.Configuration.APIClient.ParameterToString(cursor, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Socialmediaasyncaggregatequeryresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Socialmediaasyncaggregatequeryresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaAnalyticsMessagesJob invokes GET /api/v2/socialmedia/analytics/messages/jobs/{jobId}
//
// Get status for async query for social media messages job
func (a SocialMediaApi) GetSocialmediaAnalyticsMessagesJob(jobId string) (*Asyncquerystatus, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/analytics/messages/jobs/{jobId}"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Asyncquerystatus)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling SocialMediaApi->GetSocialmediaAnalyticsMessagesJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Asyncquerystatus
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Asyncquerystatus" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaAnalyticsMessagesJobResults invokes GET /api/v2/socialmedia/analytics/messages/jobs/{jobId}/results
//
// Fetch a page of results for an async social media messages query
func (a SocialMediaApi) GetSocialmediaAnalyticsMessagesJobResults(jobId string, cursor string) (*Socialmediaasyncdetailqueryresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/analytics/messages/jobs/{jobId}/results"
	path = strings.Replace(path, "{jobId}", url.PathEscape(fmt.Sprintf("%v", jobId)), -1)
	defaultReturn := new(Socialmediaasyncdetailqueryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling SocialMediaApi->GetSocialmediaAnalyticsMessagesJobResults")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["cursor"] = a.Configuration.APIClient.ParameterToString(cursor, "")
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Socialmediaasyncdetailqueryresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Socialmediaasyncdetailqueryresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaEscalationrule invokes GET /api/v2/socialmedia/escalationrules/{escalationRuleId}
//
// Get a single escalation rule.
func (a SocialMediaApi) GetSocialmediaEscalationrule(escalationRuleId string, expand string) (*Escalationruleresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/escalationrules/{escalationRuleId}"
	path = strings.Replace(path, "{escalationRuleId}", url.PathEscape(fmt.Sprintf("%v", escalationRuleId)), -1)
	defaultReturn := new(Escalationruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'escalationRuleId' is set
	if &escalationRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'escalationRuleId' when calling SocialMediaApi->GetSocialmediaEscalationrule")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Escalationruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Escalationruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaEscalationrules invokes GET /api/v2/socialmedia/escalationrules
//
// Retrieve all escalation rules for a division.
func (a SocialMediaApi) GetSocialmediaEscalationrules(divisionId string, pageNumber int, pageSize int) (*Socialescalationresponseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/escalationrules"
	defaultReturn := new(Socialescalationresponseentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'divisionId' is set
	if &divisionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'divisionId' when calling SocialMediaApi->GetSocialmediaEscalationrules")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Socialescalationresponseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Socialescalationresponseentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopic invokes GET /api/v2/socialmedia/topics/{topicId}
//
// Get a single social topic.
func (a SocialMediaApi) GetSocialmediaTopic(topicId string, includeDeleted bool) (*Socialtopicwithdataingestionrulemetadataresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	defaultReturn := new(Socialtopicwithdataingestionrulemetadataresponse)
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
	var successPayload *Socialtopicwithdataingestionrulemetadataresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Socialtopicwithdataingestionrulemetadataresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopicDataingestionrules invokes GET /api/v2/socialmedia/topics/{topicId}/dataingestionrules
//
// Retrieve all social topic data ingestion rules with pagination.
func (a SocialMediaApi) GetSocialmediaTopicDataingestionrules(topicId string, pageNumber int, pageSize int, includeDeleted bool) (*Dataingestionruleresponseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	defaultReturn := new(Dataingestionruleresponseentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrules")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Dataingestionruleresponseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Dataingestionruleresponseentitylisting" == "string" {
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

// GetSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId invokes GET /api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}
//
// Get a single Instagram data ingestion rule.
func (a SocialMediaApi) GetSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId(topicId string, instagramIngestionRuleId string, includeDeleted bool) (*Instagramdataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{instagramIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", instagramIngestionRuleId)), -1)
	defaultReturn := new(Instagramdataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId")
	}
	// verify the required parameter 'instagramIngestionRuleId' is set
	if &instagramIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'instagramIngestionRuleId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Instagramdataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Instagramdataingestionruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleIdVersion invokes GET /api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}/versions/{dataIngestionRuleVersion}
//
// Get a single Instagram data ingestion rule version.
func (a SocialMediaApi) GetSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleIdVersion(topicId string, instagramIngestionRuleId string, dataIngestionRuleVersion string, includeDeleted bool) (*Instagramdataingestionruleversionresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}/versions/{dataIngestionRuleVersion}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{instagramIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", instagramIngestionRuleId)), -1)
	path = strings.Replace(path, "{dataIngestionRuleVersion}", url.PathEscape(fmt.Sprintf("%v", dataIngestionRuleVersion)), -1)
	defaultReturn := new(Instagramdataingestionruleversionresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleIdVersion")
	}
	// verify the required parameter 'instagramIngestionRuleId' is set
	if &instagramIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'instagramIngestionRuleId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleIdVersion")
	}
	// verify the required parameter 'dataIngestionRuleVersion' is set
	if &dataIngestionRuleVersion == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'dataIngestionRuleVersion' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleIdVersion")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Instagramdataingestionruleversionresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Instagramdataingestionruleversionresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleIdVersions invokes GET /api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}/versions
//
// Get the Instagram data ingestion rule versions.
func (a SocialMediaApi) GetSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleIdVersions(topicId string, instagramIngestionRuleId string, pageNumber int, pageSize int, includeDeleted bool) (*Instagramdataingestionruleversionresponseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}/versions"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{instagramIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", instagramIngestionRuleId)), -1)
	defaultReturn := new(Instagramdataingestionruleversionresponseentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleIdVersions")
	}
	// verify the required parameter 'instagramIngestionRuleId' is set
	if &instagramIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'instagramIngestionRuleId' when calling SocialMediaApi->GetSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleIdVersions")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Instagramdataingestionruleversionresponseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Instagramdataingestionruleversionresponseentitylisting" == "string" {
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
func (a SocialMediaApi) GetSocialmediaTopics(pageNumber int, pageSize int, divisionIds []string, includeDeleted bool, name string, ids []string) (*Socialtopicwithdataingestionrulemetadataresponseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics"
	defaultReturn := new(Socialtopicwithdataingestionrulemetadataresponseentitylisting)
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
	
	queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, "")
	
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
	var successPayload *Socialtopicwithdataingestionrulemetadataresponseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Socialtopicwithdataingestionrulemetadataresponseentitylisting" == "string" {
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

// PatchSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId invokes PATCH /api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}
//
// Update the status of a Instagram data ingestion rule.
func (a SocialMediaApi) PatchSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId(topicId string, instagramIngestionRuleId string, body Dataingestionrulestatuspatchrequest) (*Instagramdataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{instagramIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", instagramIngestionRuleId)), -1)
	defaultReturn := new(Instagramdataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PatchSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId")
	}
	// verify the required parameter 'instagramIngestionRuleId' is set
	if &instagramIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'instagramIngestionRuleId' when calling SocialMediaApi->PatchSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Instagramdataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Instagramdataingestionruleresponse" == "string" {
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

// PostSocialmediaAnalyticsAggregatesJobs invokes POST /api/v2/socialmedia/analytics/aggregates/jobs
//
// Query for social media aggregates asynchronously
func (a SocialMediaApi) PostSocialmediaAnalyticsAggregatesJobs(body Socialmediaasyncaggregationquery) (*Asyncqueryresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/analytics/aggregates/jobs"
	defaultReturn := new(Asyncqueryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SocialMediaApi->PostSocialmediaAnalyticsAggregatesJobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Asyncqueryresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Asyncqueryresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSocialmediaAnalyticsMessagesJobs invokes POST /api/v2/socialmedia/analytics/messages/jobs
//
// Query for social media messages asynchronously
func (a SocialMediaApi) PostSocialmediaAnalyticsMessagesJobs(body Socialmediaasyncdetailquery) (*Asyncqueryresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/analytics/messages/jobs"
	defaultReturn := new(Asyncqueryresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SocialMediaApi->PostSocialmediaAnalyticsMessagesJobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Asyncqueryresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Asyncqueryresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSocialmediaEscalationrules invokes POST /api/v2/socialmedia/escalationrules
//
// Create an escalation rule.
func (a SocialMediaApi) PostSocialmediaEscalationrules(body Escalationrulerequest) (*Escalationruleresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/escalationrules"
	defaultReturn := new(Escalationruleresponse)
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

	var successPayload *Escalationruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Escalationruleresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSocialmediaEscalationsMessages invokes POST /api/v2/socialmedia/escalations/messages
//
// Escalate message to a conversation manually
func (a SocialMediaApi) PostSocialmediaEscalationsMessages(divisionId string, body Manualescalationrequest) (*Manualescalationresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/escalations/messages"
	defaultReturn := new(Manualescalationresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'divisionId' is set
	if &divisionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'divisionId' when calling SocialMediaApi->PostSocialmediaEscalationsMessages")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
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
	// body params
	postBody = &body

	var successPayload *Manualescalationresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Manualescalationresponse" == "string" {
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

// PostSocialmediaTopicDataingestionrulesInstagram invokes POST /api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram
//
// Create an Instagram data ingestion rule.
func (a SocialMediaApi) PostSocialmediaTopicDataingestionrulesInstagram(topicId string, body Instagramdataingestionrulerequest) (*Instagramdataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	defaultReturn := new(Instagramdataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PostSocialmediaTopicDataingestionrulesInstagram")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Instagramdataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Instagramdataingestionruleresponse" == "string" {
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

// PostSocialmediaTopicDataingestionrulesOpenRuleIdMessagesBulk invokes POST /api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{ruleId}/messages/bulk
//
// Ingest a list of Open Social Messages
//
// Ingest a list of open social messages to an ingestion rule on a topic. This endpoint will ingest and enrich these messages.  In order to call this endpoint you will need OAuth token generated using OAuth client credentials authorized with at least social scope.
func (a SocialMediaApi) PostSocialmediaTopicDataingestionrulesOpenRuleIdMessagesBulk(topicId string, ruleId string, body []Opensocialmedianormalizedmessage) (*Opensocialnormalizedmessageentitylisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{ruleId}/messages/bulk"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{ruleId}", url.PathEscape(fmt.Sprintf("%v", ruleId)), -1)
	defaultReturn := new(Opensocialnormalizedmessageentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PostSocialmediaTopicDataingestionrulesOpenRuleIdMessagesBulk")
	}
	// verify the required parameter 'ruleId' is set
	if &ruleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'ruleId' when calling SocialMediaApi->PostSocialmediaTopicDataingestionrulesOpenRuleIdMessagesBulk")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// true
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SocialMediaApi->PostSocialmediaTopicDataingestionrulesOpenRuleIdMessagesBulk")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Opensocialnormalizedmessageentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Opensocialnormalizedmessageentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostSocialmediaTopicDataingestionrulesOpenRuleIdReactionsBulk invokes POST /api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{ruleId}/reactions/bulk
//
// Ingest a list of Open Social Reactions
//
// Ingest a list of open social reactions to an ingestion rule on a topic. This endpoint will ingest these reactions.  In order to call this endpoint you will need OAuth token generated using OAuth client credentials authorized with at least social scope.
func (a SocialMediaApi) PostSocialmediaTopicDataingestionrulesOpenRuleIdReactionsBulk(topicId string, ruleId string, body Opensocialmediareactionsrequest) (*Opensocialreactionsnormalizedevententitylisting, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/open/{ruleId}/reactions/bulk"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{ruleId}", url.PathEscape(fmt.Sprintf("%v", ruleId)), -1)
	defaultReturn := new(Opensocialreactionsnormalizedevententitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PostSocialmediaTopicDataingestionrulesOpenRuleIdReactionsBulk")
	}
	// verify the required parameter 'ruleId' is set
	if &ruleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'ruleId' when calling SocialMediaApi->PostSocialmediaTopicDataingestionrulesOpenRuleIdReactionsBulk")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SocialMediaApi->PostSocialmediaTopicDataingestionrulesOpenRuleIdReactionsBulk")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Opensocialreactionsnormalizedevententitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Opensocialreactionsnormalizedevententitylisting" == "string" {
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

// PostSocialmediaTwitterHistoricalTweets invokes POST /api/v2/socialmedia/twitter/historical/tweets
//
// Retrieves historical tweet count for search terms, optional countries list and the current limit and usage for the organization.
func (a SocialMediaApi) PostSocialmediaTwitterHistoricalTweets(body Twitterdatahistoricaltweetrequest) (*Twitterdatahistoricaltweetresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/twitter/historical/tweets"
	defaultReturn := new(Twitterdatahistoricaltweetresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling SocialMediaApi->PostSocialmediaTwitterHistoricalTweets")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Twitterdatahistoricaltweetresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Twitterdatahistoricaltweetresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutSocialmediaEscalationrule invokes PUT /api/v2/socialmedia/escalationrules/{escalationRuleId}
//
// Update the escalation rule.
func (a SocialMediaApi) PutSocialmediaEscalationrule(escalationRuleId string, body Escalationrulerequest) (*Escalationruleresponse, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/escalationrules/{escalationRuleId}"
	path = strings.Replace(path, "{escalationRuleId}", url.PathEscape(fmt.Sprintf("%v", escalationRuleId)), -1)
	defaultReturn := new(Escalationruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'escalationRuleId' is set
	if &escalationRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'escalationRuleId' when calling SocialMediaApi->PutSocialmediaEscalationrule")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Escalationruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Escalationruleresponse" == "string" {
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

// PutSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId invokes PUT /api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}
//
// Update the Instagram data ingestion rule.
func (a SocialMediaApi) PutSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId(topicId string, instagramIngestionRuleId string, body Instagramdataingestionrulerequest) (*Instagramdataingestionruleresponse, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/socialmedia/topics/{topicId}/dataingestionrules/instagram/{instagramIngestionRuleId}"
	path = strings.Replace(path, "{topicId}", url.PathEscape(fmt.Sprintf("%v", topicId)), -1)
	path = strings.Replace(path, "{instagramIngestionRuleId}", url.PathEscape(fmt.Sprintf("%v", instagramIngestionRuleId)), -1)
	defaultReturn := new(Instagramdataingestionruleresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'topicId' is set
	if &topicId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'topicId' when calling SocialMediaApi->PutSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId")
	}
	// verify the required parameter 'instagramIngestionRuleId' is set
	if &instagramIngestionRuleId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'instagramIngestionRuleId' when calling SocialMediaApi->PutSocialmediaTopicDataingestionrulesInstagramInstagramIngestionRuleId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// Find an replace keys that were altered to avoid clashes with go keywords 
	correctedQueryParams := make(map[string]string)
	for k, v := range queryParams {
		if k == "varType" {
			correctedQueryParams["type"] = v
			continue
		}
		correctedQueryParams[k] = v
	}
	queryParams = correctedQueryParams

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &body

	var successPayload *Instagramdataingestionruleresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Instagramdataingestionruleresponse" == "string" {
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

