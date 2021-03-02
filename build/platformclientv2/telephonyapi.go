package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"time"
"encoding/json"
)

// TelephonyApi provides functions for API endpoints
type TelephonyApi struct {
	Configuration *Configuration
}

// NewTelephonyApi creates an API instance using the default configuration
func NewTelephonyApi() *TelephonyApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating TelephonyApi with base path: %s", strings.ToLower(config.BasePath)))
	return &TelephonyApi{
		Configuration: config,
	}
}

// NewTelephonyApiWithConfig creates an API instance using the provided configuration
func NewTelephonyApiWithConfig(config *Configuration) *TelephonyApi {
	config.Debugf("Creating TelephonyApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &TelephonyApi{
		Configuration: config,
	}
}

// GetTelephonySiptraces invokes GET /api/v2/telephony/siptraces
//
// Fetch SIP metadata
//
// Fetch SIP metadata that matches a given parameter. If exactMatch is passed as a parameter only sip records that have exactly that value will be returned. For example, some records contain conversationId but not all relevant records for that call may contain the conversationId so only a partial view of the call will be reflected
func (a TelephonyApi) GetTelephonySiptraces(dateStart time.Time, dateEnd time.Time, callId string, toUser string, fromUser string, conversationId string) (*Sipsearchresult, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/siptraces"
	defaultReturn := new(Sipsearchresult)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'dateStart' is set
	if &dateStart == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'dateStart' when calling TelephonyApi->GetTelephonySiptraces")
	}
	// verify the required parameter 'dateEnd' is set
	if &dateEnd == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'dateEnd' when calling TelephonyApi->GetTelephonySiptraces")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	if str, ok := interface{}(callId).(string); ok {
		if str != "" {
			queryParams["callId"] = a.Configuration.APIClient.ParameterToString(callId, collectionFormat)
		}
	} else {
		queryParams["callId"] = a.Configuration.APIClient.ParameterToString(callId, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(toUser).(string); ok {
		if str != "" {
			queryParams["toUser"] = a.Configuration.APIClient.ParameterToString(toUser, collectionFormat)
		}
	} else {
		queryParams["toUser"] = a.Configuration.APIClient.ParameterToString(toUser, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(fromUser).(string); ok {
		if str != "" {
			queryParams["fromUser"] = a.Configuration.APIClient.ParameterToString(fromUser, collectionFormat)
		}
	} else {
		queryParams["fromUser"] = a.Configuration.APIClient.ParameterToString(fromUser, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(conversationId).(string); ok {
		if str != "" {
			queryParams["conversationId"] = a.Configuration.APIClient.ParameterToString(conversationId, collectionFormat)
		}
	} else {
		queryParams["conversationId"] = a.Configuration.APIClient.ParameterToString(conversationId, collectionFormat)
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
	
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Sipsearchresult
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// GetTelephonySiptracesDownloadDownloadId invokes GET /api/v2/telephony/siptraces/download/{downloadId}
//
// Get signed S3 URL for a pcap download
//
// 
func (a TelephonyApi) GetTelephonySiptracesDownloadDownloadId(downloadId string) (*Signedurlresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/siptraces/download/{downloadId}"
	path = strings.Replace(path, "{downloadId}", fmt.Sprintf("%v", downloadId), -1)
	defaultReturn := new(Signedurlresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'downloadId' is set
	if &downloadId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'downloadId' when calling TelephonyApi->GetTelephonySiptracesDownloadDownloadId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Signedurlresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

// PostTelephonySiptracesDownload invokes POST /api/v2/telephony/siptraces/download
//
// Request a download of a pcap file to S3
//
// 
func (a TelephonyApi) PostTelephonySiptracesDownload(sIPSearchPublicRequest Sipsearchpublicrequest) (*Sipdownloadresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/siptraces/download"
	defaultReturn := new(Sipdownloadresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'sIPSearchPublicRequest' is set
	if &sIPSearchPublicRequest == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'sIPSearchPublicRequest' when calling TelephonyApi->PostTelephonySiptracesDownload")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &sIPSearchPublicRequest

	var successPayload *Sipdownloadresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else {
		err = json.Unmarshal([]byte(response.RawBody), &successPayload)
	}
	return successPayload, response, err
}

