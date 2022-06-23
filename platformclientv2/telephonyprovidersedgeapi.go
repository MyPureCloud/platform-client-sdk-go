package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
	)

// TelephonyProvidersEdgeApi provides functions for API endpoints
type TelephonyProvidersEdgeApi struct {
	Configuration *Configuration
}

// NewTelephonyProvidersEdgeApi creates an API instance using the default configuration
func NewTelephonyProvidersEdgeApi() *TelephonyProvidersEdgeApi {
	fmt.Sprintf(strings.Title(""), "")
	config := GetDefaultConfiguration()
	return &TelephonyProvidersEdgeApi{
		Configuration: config,
	}
}

// NewTelephonyProvidersEdgeApiWithConfig creates an API instance using the provided configuration
func NewTelephonyProvidersEdgeApiWithConfig(config *Configuration) *TelephonyProvidersEdgeApi {
	return &TelephonyProvidersEdgeApi{
		Configuration: config,
	}
}

// DeleteTelephonyProvidersEdge invokes DELETE /api/v2/telephony/providers/edges/{edgeId}
//
// Delete a edge.
func (a TelephonyProvidersEdgeApi) DeleteTelephonyProvidersEdge(edgeId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdge")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteTelephonyProvidersEdgeLogicalinterface invokes DELETE /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}
//
// Delete an edge logical interface
func (a TelephonyProvidersEdgeApi) DeleteTelephonyProvidersEdgeLogicalinterface(edgeId string, interfaceId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgeLogicalinterface")
	}
	// verify the required parameter 'interfaceId' is set
	if &interfaceId == nil {
		// false
		return nil, errors.New("Missing required parameter 'interfaceId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgeLogicalinterface")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteTelephonyProvidersEdgeSoftwareupdate invokes DELETE /api/v2/telephony/providers/edges/{edgeId}/softwareupdate
//
// Cancels any in-progress update for this edge.
func (a TelephonyProvidersEdgeApi) DeleteTelephonyProvidersEdgeSoftwareupdate(edgeId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/softwareupdate"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgeSoftwareupdate")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteTelephonyProvidersEdgesCertificateauthority invokes DELETE /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}
//
// Delete a certificate authority.
func (a TelephonyProvidersEdgeApi) DeleteTelephonyProvidersEdgesCertificateauthority(certificateId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/certificateauthorities/{certificateId}"
	path = strings.Replace(path, "{certificateId}", fmt.Sprintf("%v", certificateId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'certificateId' is set
	if &certificateId == nil {
		// false
		return nil, errors.New("Missing required parameter 'certificateId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgesCertificateauthority")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteTelephonyProvidersEdgesDidpool invokes DELETE /api/v2/telephony/providers/edges/didpools/{didPoolId}
//
// Delete a DID Pool by ID.
func (a TelephonyProvidersEdgeApi) DeleteTelephonyProvidersEdgesDidpool(didPoolId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/didpools/{didPoolId}"
	path = strings.Replace(path, "{didPoolId}", fmt.Sprintf("%v", didPoolId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'didPoolId' is set
	if &didPoolId == nil {
		// false
		return nil, errors.New("Missing required parameter 'didPoolId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgesDidpool")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteTelephonyProvidersEdgesEdgegroup invokes DELETE /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}
//
// Delete an edge group.
func (a TelephonyProvidersEdgeApi) DeleteTelephonyProvidersEdgesEdgegroup(edgeGroupId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}"
	path = strings.Replace(path, "{edgeGroupId}", fmt.Sprintf("%v", edgeGroupId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeGroupId' is set
	if &edgeGroupId == nil {
		// false
		return nil, errors.New("Missing required parameter 'edgeGroupId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgesEdgegroup")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteTelephonyProvidersEdgesExtensionpool invokes DELETE /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}
//
// Delete an extension pool by ID
func (a TelephonyProvidersEdgeApi) DeleteTelephonyProvidersEdgesExtensionpool(extensionPoolId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}"
	path = strings.Replace(path, "{extensionPoolId}", fmt.Sprintf("%v", extensionPoolId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'extensionPoolId' is set
	if &extensionPoolId == nil {
		// false
		return nil, errors.New("Missing required parameter 'extensionPoolId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgesExtensionpool")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteTelephonyProvidersEdgesOutboundroute invokes DELETE /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}
//
// Delete Outbound Route
//
// This route is deprecated, use /telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId} instead.
func (a TelephonyProvidersEdgeApi) DeleteTelephonyProvidersEdgesOutboundroute(outboundRouteId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}"
	path = strings.Replace(path, "{outboundRouteId}", fmt.Sprintf("%v", outboundRouteId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'outboundRouteId' is set
	if &outboundRouteId == nil {
		// false
		return nil, errors.New("Missing required parameter 'outboundRouteId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgesOutboundroute")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteTelephonyProvidersEdgesPhone invokes DELETE /api/v2/telephony/providers/edges/phones/{phoneId}
//
// Delete a Phone by ID
func (a TelephonyProvidersEdgeApi) DeleteTelephonyProvidersEdgesPhone(phoneId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phones/{phoneId}"
	path = strings.Replace(path, "{phoneId}", fmt.Sprintf("%v", phoneId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'phoneId' is set
	if &phoneId == nil {
		// false
		return nil, errors.New("Missing required parameter 'phoneId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgesPhone")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteTelephonyProvidersEdgesPhonebasesetting invokes DELETE /api/v2/telephony/providers/edges/phonebasesettings/{phoneBaseId}
//
// Delete a Phone Base Settings by ID
func (a TelephonyProvidersEdgeApi) DeleteTelephonyProvidersEdgesPhonebasesetting(phoneBaseId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phonebasesettings/{phoneBaseId}"
	path = strings.Replace(path, "{phoneBaseId}", fmt.Sprintf("%v", phoneBaseId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'phoneBaseId' is set
	if &phoneBaseId == nil {
		// false
		return nil, errors.New("Missing required parameter 'phoneBaseId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgesPhonebasesetting")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteTelephonyProvidersEdgesSite invokes DELETE /api/v2/telephony/providers/edges/sites/{siteId}
//
// Delete a Site by ID
func (a TelephonyProvidersEdgeApi) DeleteTelephonyProvidersEdgesSite(siteId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'siteId' is set
	if &siteId == nil {
		// false
		return nil, errors.New("Missing required parameter 'siteId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgesSite")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteTelephonyProvidersEdgesSiteOutboundroute invokes DELETE /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}
//
// Delete Outbound Route
func (a TelephonyProvidersEdgeApi) DeleteTelephonyProvidersEdgesSiteOutboundroute(siteId string, outboundRouteId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
	path = strings.Replace(path, "{outboundRouteId}", fmt.Sprintf("%v", outboundRouteId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'siteId' is set
	if &siteId == nil {
		// false
		return nil, errors.New("Missing required parameter 'siteId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgesSiteOutboundroute")
	}
	// verify the required parameter 'outboundRouteId' is set
	if &outboundRouteId == nil {
		// false
		return nil, errors.New("Missing required parameter 'outboundRouteId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgesSiteOutboundroute")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// DeleteTelephonyProvidersEdgesTrunkbasesetting invokes DELETE /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}
//
// Delete a Trunk Base Settings object by ID
func (a TelephonyProvidersEdgeApi) DeleteTelephonyProvidersEdgesTrunkbasesetting(trunkBaseSettingsId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}"
	path = strings.Replace(path, "{trunkBaseSettingsId}", fmt.Sprintf("%v", trunkBaseSettingsId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'trunkBaseSettingsId' is set
	if &trunkBaseSettingsId == nil {
		// false
		return nil, errors.New("Missing required parameter 'trunkBaseSettingsId' when calling TelephonyProvidersEdgeApi->DeleteTelephonyProvidersEdgesTrunkbasesetting")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// GetConfigurationSchemasEdgesVnext invokes GET /api/v2/configuration/schemas/edges/vnext
//
// Lists available schema categories (Deprecated)
func (a TelephonyProvidersEdgeApi) GetConfigurationSchemasEdgesVnext(pageSize int, pageNumber int) (*Schemacategoryentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/configuration/schemas/edges/vnext"
	defaultReturn := new(Schemacategoryentitylisting)
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Schemacategoryentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Schemacategoryentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetConfigurationSchemasEdgesVnextSchemaCategory invokes GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}
//
// List schemas of a specific category (Deprecated)
func (a TelephonyProvidersEdgeApi) GetConfigurationSchemasEdgesVnextSchemaCategory(schemaCategory string, pageSize int, pageNumber int) (*Schemareferenceentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/configuration/schemas/edges/vnext/{schemaCategory}"
	path = strings.Replace(path, "{schemaCategory}", fmt.Sprintf("%v", schemaCategory), -1)
	defaultReturn := new(Schemareferenceentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'schemaCategory' is set
	if &schemaCategory == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'schemaCategory' when calling TelephonyProvidersEdgeApi->GetConfigurationSchemasEdgesVnextSchemaCategory")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Schemareferenceentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Schemareferenceentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetConfigurationSchemasEdgesVnextSchemaCategorySchemaType invokes GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}
//
// List schemas of a specific category (Deprecated)
func (a TelephonyProvidersEdgeApi) GetConfigurationSchemasEdgesVnextSchemaCategorySchemaType(schemaCategory string, schemaType string, pageSize int, pageNumber int) (*Schemareferenceentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}"
	path = strings.Replace(path, "{schemaCategory}", fmt.Sprintf("%v", schemaCategory), -1)
	path = strings.Replace(path, "{schemaType}", fmt.Sprintf("%v", schemaType), -1)
	defaultReturn := new(Schemareferenceentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'schemaCategory' is set
	if &schemaCategory == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'schemaCategory' when calling TelephonyProvidersEdgeApi->GetConfigurationSchemasEdgesVnextSchemaCategorySchemaType")
	}
	// verify the required parameter 'schemaType' is set
	if &schemaType == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'schemaType' when calling TelephonyProvidersEdgeApi->GetConfigurationSchemasEdgesVnextSchemaCategorySchemaType")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Schemareferenceentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Schemareferenceentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaId invokes GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}
//
// Get a json schema (Deprecated)
func (a TelephonyProvidersEdgeApi) GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaId(schemaCategory string, schemaType string, schemaId string) (*Organization, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}"
	path = strings.Replace(path, "{schemaCategory}", fmt.Sprintf("%v", schemaCategory), -1)
	path = strings.Replace(path, "{schemaType}", fmt.Sprintf("%v", schemaType), -1)
	path = strings.Replace(path, "{schemaId}", fmt.Sprintf("%v", schemaId), -1)
	defaultReturn := new(Organization)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'schemaCategory' is set
	if &schemaCategory == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'schemaCategory' when calling TelephonyProvidersEdgeApi->GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaId")
	}
	// verify the required parameter 'schemaType' is set
	if &schemaType == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'schemaType' when calling TelephonyProvidersEdgeApi->GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaId")
	}
	// verify the required parameter 'schemaId' is set
	if &schemaId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'schemaId' when calling TelephonyProvidersEdgeApi->GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Organization
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Organization" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataId invokes GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}
//
// Get metadata for a schema (Deprecated)
func (a TelephonyProvidersEdgeApi) GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataId(schemaCategory string, schemaType string, schemaId string, extensionType string, metadataId string, varType string) (*Organization, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}"
	path = strings.Replace(path, "{schemaCategory}", fmt.Sprintf("%v", schemaCategory), -1)
	path = strings.Replace(path, "{schemaType}", fmt.Sprintf("%v", schemaType), -1)
	path = strings.Replace(path, "{schemaId}", fmt.Sprintf("%v", schemaId), -1)
	path = strings.Replace(path, "{extensionType}", fmt.Sprintf("%v", extensionType), -1)
	path = strings.Replace(path, "{metadataId}", fmt.Sprintf("%v", metadataId), -1)
	defaultReturn := new(Organization)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'schemaCategory' is set
	if &schemaCategory == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'schemaCategory' when calling TelephonyProvidersEdgeApi->GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataId")
	}
	// verify the required parameter 'schemaType' is set
	if &schemaType == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'schemaType' when calling TelephonyProvidersEdgeApi->GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataId")
	}
	// verify the required parameter 'schemaId' is set
	if &schemaId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'schemaId' when calling TelephonyProvidersEdgeApi->GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataId")
	}
	// verify the required parameter 'extensionType' is set
	if &extensionType == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'extensionType' when calling TelephonyProvidersEdgeApi->GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataId")
	}
	// verify the required parameter 'metadataId' is set
	if &metadataId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'metadataId' when calling TelephonyProvidersEdgeApi->GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Organization
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Organization" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdge invokes GET /api/v2/telephony/providers/edges/{edgeId}
//
// Get edge.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdge(edgeId string, expand []string) (*Edge, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Edge)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdge")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Edge
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edge" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeDiagnosticNslookup invokes GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup
//
// Get networking-related information from an Edge for a target IP or host.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeDiagnosticNslookup(edgeId string) (*Edgenetworkdiagnosticresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Edgenetworkdiagnosticresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeDiagnosticNslookup")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Edgenetworkdiagnosticresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgenetworkdiagnosticresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeDiagnosticPing invokes GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping
//
// Get networking-related information from an Edge for a target IP or host.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeDiagnosticPing(edgeId string) (*Edgenetworkdiagnosticresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Edgenetworkdiagnosticresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeDiagnosticPing")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Edgenetworkdiagnosticresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgenetworkdiagnosticresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeDiagnosticRoute invokes GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route
//
// Get networking-related information from an Edge for a target IP or host.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeDiagnosticRoute(edgeId string) (*Edgenetworkdiagnosticresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/diagnostic/route"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Edgenetworkdiagnosticresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeDiagnosticRoute")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Edgenetworkdiagnosticresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgenetworkdiagnosticresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeDiagnosticTracepath invokes GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath
//
// Get networking-related information from an Edge for a target IP or host.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeDiagnosticTracepath(edgeId string) (*Edgenetworkdiagnosticresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Edgenetworkdiagnosticresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeDiagnosticTracepath")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Edgenetworkdiagnosticresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgenetworkdiagnosticresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeLine invokes GET /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}
//
// Get line
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeLine(edgeId string, lineId string) (*Edgeline, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	path = strings.Replace(path, "{lineId}", fmt.Sprintf("%v", lineId), -1)
	defaultReturn := new(Edgeline)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeLine")
	}
	// verify the required parameter 'lineId' is set
	if &lineId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'lineId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeLine")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Edgeline
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgeline" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeLines invokes GET /api/v2/telephony/providers/edges/{edgeId}/lines
//
// Get the list of lines.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeLines(edgeId string, pageSize int, pageNumber int) (*Edgelineentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/lines"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Edgelineentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeLines")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Edgelineentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgelineentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeLogicalinterface invokes GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}
//
// Get an edge logical interface
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeLogicalinterface(edgeId string, interfaceId string, expand []string) (*Domainlogicalinterface, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceId), -1)
	defaultReturn := new(Domainlogicalinterface)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeLogicalinterface")
	}
	// verify the required parameter 'interfaceId' is set
	if &interfaceId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'interfaceId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeLogicalinterface")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Domainlogicalinterface
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Domainlogicalinterface" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeLogicalinterfaces invokes GET /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces
//
// Get edge logical interfaces.
//
// Retrieve a list of all configured logical interfaces from a specific edge.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeLogicalinterfaces(edgeId string, expand []string) (*Logicalinterfaceentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Logicalinterfaceentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeLogicalinterfaces")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Logicalinterfaceentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Logicalinterfaceentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeLogsJob invokes GET /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}
//
// Get an Edge logs job.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeLogsJob(edgeId string, jobId string) (*Edgelogsjob, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	path = strings.Replace(path, "{jobId}", fmt.Sprintf("%v", jobId), -1)
	defaultReturn := new(Edgelogsjob)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeLogsJob")
	}
	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'jobId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeLogsJob")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Edgelogsjob
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgelogsjob" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeMetrics invokes GET /api/v2/telephony/providers/edges/{edgeId}/metrics
//
// Get the edge metrics.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeMetrics(edgeId string) (*Edgemetrics, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/metrics"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Edgemetrics)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeMetrics")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Edgemetrics
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgemetrics" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgePhysicalinterface invokes GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}
//
// Get edge physical interface.
//
// Retrieve a physical interface from a specific edge.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgePhysicalinterface(edgeId string, interfaceId string) (*Domainphysicalinterface, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces/{interfaceId}"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceId), -1)
	defaultReturn := new(Domainphysicalinterface)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgePhysicalinterface")
	}
	// verify the required parameter 'interfaceId' is set
	if &interfaceId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'interfaceId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgePhysicalinterface")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Domainphysicalinterface
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Domainphysicalinterface" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgePhysicalinterfaces invokes GET /api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces
//
// Retrieve a list of all configured physical interfaces from a specific edge.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgePhysicalinterfaces(edgeId string) (*Physicalinterfaceentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/physicalinterfaces"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Physicalinterfaceentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgePhysicalinterfaces")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Physicalinterfaceentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Physicalinterfaceentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeSetuppackage invokes GET /api/v2/telephony/providers/edges/{edgeId}/setuppackage
//
// Get the setup package for a locally deployed edge device. This is needed to complete the setup process for the virtual edge.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeSetuppackage(edgeId string) (*Vmpairinginfo, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/setuppackage"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Vmpairinginfo)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeSetuppackage")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Vmpairinginfo
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Vmpairinginfo" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeSoftwareupdate invokes GET /api/v2/telephony/providers/edges/{edgeId}/softwareupdate
//
// Gets software update status information about any edge.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeSoftwareupdate(edgeId string) (*Domainedgesoftwareupdatedto, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/softwareupdate"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Domainedgesoftwareupdatedto)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeSoftwareupdate")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Domainedgesoftwareupdatedto
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Domainedgesoftwareupdatedto" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeSoftwareversions invokes GET /api/v2/telephony/providers/edges/{edgeId}/softwareversions
//
// Gets all the available software versions for this edge.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeSoftwareversions(edgeId string) (*Domainedgesoftwareversiondtoentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/softwareversions"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Domainedgesoftwareversiondtoentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeSoftwareversions")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Domainedgesoftwareversiondtoentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Domainedgesoftwareversiondtoentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgeTrunks invokes GET /api/v2/telephony/providers/edges/{edgeId}/trunks
//
// Get the list of available trunks for the given Edge.
//
// Trunks are created by assigning trunk base settings to an Edge or Edge Group.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgeTrunks(edgeId string, pageNumber int, pageSize int, sortBy string, sortOrder string, trunkBaseId string, trunkType string) (*Trunkentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/trunks"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Trunkentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgeTrunks")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
	queryParams["trunkBaseId"] = a.Configuration.APIClient.ParameterToString(trunkBaseId, "")
	
	queryParams["trunkType"] = a.Configuration.APIClient.ParameterToString(trunkType, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Trunkentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Trunkentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdges invokes GET /api/v2/telephony/providers/edges
//
// Get the list of edges.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdges(pageSize int, pageNumber int, name string, siteId string, edgeGroupId string, sortBy string, managed bool) (*Edgeentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges"
	defaultReturn := new(Edgeentitylisting)
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
	
	queryParams["siteId"] = a.Configuration.APIClient.ParameterToString(siteId, "")
	
	queryParams["edgeGroupId"] = a.Configuration.APIClient.ParameterToString(edgeGroupId, "")
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["managed"] = a.Configuration.APIClient.ParameterToString(managed, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Edgeentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgeentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesAvailablelanguages invokes GET /api/v2/telephony/providers/edges/availablelanguages
//
// Get the list of available languages.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesAvailablelanguages() (*Availablelanguagelist, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/availablelanguages"
	defaultReturn := new(Availablelanguagelist)
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
	var successPayload *Availablelanguagelist
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Availablelanguagelist" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesCertificateauthorities invokes GET /api/v2/telephony/providers/edges/certificateauthorities
//
// Get the list of certificate authorities.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesCertificateauthorities() (*Certificateauthorityentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/certificateauthorities"
	defaultReturn := new(Certificateauthorityentitylisting)
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
	var successPayload *Certificateauthorityentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Certificateauthorityentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesCertificateauthority invokes GET /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}
//
// Get a certificate authority.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesCertificateauthority(certificateId string) (*Domaincertificateauthority, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/certificateauthorities/{certificateId}"
	path = strings.Replace(path, "{certificateId}", fmt.Sprintf("%v", certificateId), -1)
	defaultReturn := new(Domaincertificateauthority)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'certificateId' is set
	if &certificateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'certificateId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesCertificateauthority")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Domaincertificateauthority
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Domaincertificateauthority" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesDid invokes GET /api/v2/telephony/providers/edges/dids/{didId}
//
// Get a DID by ID.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesDid(didId string) (*Did, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/dids/{didId}"
	path = strings.Replace(path, "{didId}", fmt.Sprintf("%v", didId), -1)
	defaultReturn := new(Did)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'didId' is set
	if &didId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'didId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesDid")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Did
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Did" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesDidpool invokes GET /api/v2/telephony/providers/edges/didpools/{didPoolId}
//
// Get a DID Pool by ID.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesDidpool(didPoolId string) (*Didpool, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/didpools/{didPoolId}"
	path = strings.Replace(path, "{didPoolId}", fmt.Sprintf("%v", didPoolId), -1)
	defaultReturn := new(Didpool)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'didPoolId' is set
	if &didPoolId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'didPoolId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesDidpool")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Didpool
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Didpool" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesDidpools invokes GET /api/v2/telephony/providers/edges/didpools
//
// Get a listing of DID Pools
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesDidpools(pageSize int, pageNumber int, sortBy string, id []string) (*Didpoolentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/didpools"
	defaultReturn := new(Didpoolentitylisting)
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
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["id"] = a.Configuration.APIClient.ParameterToString(id, "multi")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Didpoolentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Didpoolentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesDidpoolsDids invokes GET /api/v2/telephony/providers/edges/didpools/dids
//
// Get a listing of unassigned and/or assigned numbers in a set of DID Pools.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesDidpoolsDids(varType string, id []string, numberMatch string, pageSize int, pageNumber int, sortOrder string) (*Didnumberentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/didpools/dids"
	defaultReturn := new(Didnumberentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'varType' is set
	if &varType == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'varType' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesDidpoolsDids")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	
	queryParams["id"] = a.Configuration.APIClient.ParameterToString(id, "multi")
	
	queryParams["numberMatch"] = a.Configuration.APIClient.ParameterToString(numberMatch, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, "")
	
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
	var successPayload *Didnumberentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Didnumberentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesDids invokes GET /api/v2/telephony/providers/edges/dids
//
// Get a listing of DIDs
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesDids(pageSize int, pageNumber int, sortBy string, sortOrder string, phoneNumber string, ownerId string, didPoolId string, id []string) (*Didentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/dids"
	defaultReturn := new(Didentitylisting)
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
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
	queryParams["phoneNumber"] = a.Configuration.APIClient.ParameterToString(phoneNumber, "")
	
	queryParams["ownerId"] = a.Configuration.APIClient.ParameterToString(ownerId, "")
	
	queryParams["didPoolId"] = a.Configuration.APIClient.ParameterToString(didPoolId, "")
	
	queryParams["id"] = a.Configuration.APIClient.ParameterToString(id, "multi")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Didentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Didentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesEdgegroup invokes GET /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}
//
// Get edge group.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesEdgegroup(edgeGroupId string, expand []string) (*Edgegroup, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}"
	path = strings.Replace(path, "{edgeGroupId}", fmt.Sprintf("%v", edgeGroupId), -1)
	defaultReturn := new(Edgegroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeGroupId' is set
	if &edgeGroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeGroupId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesEdgegroup")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Edgegroup
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgegroup" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesEdgegroupEdgetrunkbase invokes GET /api/v2/telephony/providers/edges/edgegroups/{edgegroupId}/edgetrunkbases/{edgetrunkbaseId}
//
// Gets the edge trunk base associated with the edge group
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesEdgegroupEdgetrunkbase(edgegroupId string, edgetrunkbaseId string) (*Edgetrunkbase, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/edgegroups/{edgegroupId}/edgetrunkbases/{edgetrunkbaseId}"
	path = strings.Replace(path, "{edgegroupId}", fmt.Sprintf("%v", edgegroupId), -1)
	path = strings.Replace(path, "{edgetrunkbaseId}", fmt.Sprintf("%v", edgetrunkbaseId), -1)
	defaultReturn := new(Edgetrunkbase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgegroupId' is set
	if &edgegroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgegroupId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesEdgegroupEdgetrunkbase")
	}
	// verify the required parameter 'edgetrunkbaseId' is set
	if &edgetrunkbaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgetrunkbaseId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesEdgegroupEdgetrunkbase")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Edgetrunkbase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgetrunkbase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesEdgegroups invokes GET /api/v2/telephony/providers/edges/edgegroups
//
// Get the list of edge groups.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesEdgegroups(pageSize int, pageNumber int, name string, sortBy string, managed bool) (*Edgegroupentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/edgegroups"
	defaultReturn := new(Edgegroupentitylisting)
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
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["managed"] = a.Configuration.APIClient.ParameterToString(managed, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Edgegroupentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgegroupentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesEdgeversionreport invokes GET /api/v2/telephony/providers/edges/edgeversionreport
//
// Get the edge version report.
//
// The report will not have consistent data about the edge version(s) until all edges have been reset.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesEdgeversionreport() (*Edgeversionreport, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/edgeversionreport"
	defaultReturn := new(Edgeversionreport)
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
	var successPayload *Edgeversionreport
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgeversionreport" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesExpired invokes GET /api/v2/telephony/providers/edges/expired
//
// List of edges more than 4 edge versions behind the latest software.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesExpired() (*Expirededgelisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/expired"
	defaultReturn := new(Expirededgelisting)
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
	var successPayload *Expirededgelisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Expirededgelisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesExtension invokes GET /api/v2/telephony/providers/edges/extensions/{extensionId}
//
// Get an extension by ID.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesExtension(extensionId string) (*Extension, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/extensions/{extensionId}"
	path = strings.Replace(path, "{extensionId}", fmt.Sprintf("%v", extensionId), -1)
	defaultReturn := new(Extension)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'extensionId' is set
	if &extensionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'extensionId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesExtension")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Extension
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Extension" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesExtensionpool invokes GET /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}
//
// Get an extension pool by ID
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesExtensionpool(extensionPoolId string) (*Extensionpool, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}"
	path = strings.Replace(path, "{extensionPoolId}", fmt.Sprintf("%v", extensionPoolId), -1)
	defaultReturn := new(Extensionpool)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'extensionPoolId' is set
	if &extensionPoolId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'extensionPoolId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesExtensionpool")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Extensionpool
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Extensionpool" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesExtensionpools invokes GET /api/v2/telephony/providers/edges/extensionpools
//
// Get a listing of extension pools
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesExtensionpools(pageSize int, pageNumber int, sortBy string, number string) (*Extensionpoolentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/extensionpools"
	defaultReturn := new(Extensionpoolentitylisting)
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
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["number"] = a.Configuration.APIClient.ParameterToString(number, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Extensionpoolentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Extensionpoolentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesExtensions invokes GET /api/v2/telephony/providers/edges/extensions
//
// Get a listing of extensions
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesExtensions(pageSize int, pageNumber int, sortBy string, sortOrder string, number string) (*Extensionentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/extensions"
	defaultReturn := new(Extensionentitylisting)
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
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
	queryParams["number"] = a.Configuration.APIClient.ParameterToString(number, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Extensionentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Extensionentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesLine invokes GET /api/v2/telephony/providers/edges/lines/{lineId}
//
// Get a Line by ID
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesLine(lineId string) (*Line, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/lines/{lineId}"
	path = strings.Replace(path, "{lineId}", fmt.Sprintf("%v", lineId), -1)
	defaultReturn := new(Line)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'lineId' is set
	if &lineId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'lineId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesLine")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Line
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Line" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesLinebasesetting invokes GET /api/v2/telephony/providers/edges/linebasesettings/{lineBaseId}
//
// Get a line base settings object by ID
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesLinebasesetting(lineBaseId string) (*Linebase, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/linebasesettings/{lineBaseId}"
	path = strings.Replace(path, "{lineBaseId}", fmt.Sprintf("%v", lineBaseId), -1)
	defaultReturn := new(Linebase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'lineBaseId' is set
	if &lineBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'lineBaseId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesLinebasesetting")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Linebase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Linebase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesLinebasesettings invokes GET /api/v2/telephony/providers/edges/linebasesettings
//
// Get a listing of line base settings objects
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesLinebasesettings(pageNumber int, pageSize int, sortBy string, sortOrder string, expand []string) (*Linebaseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/linebasesettings"
	defaultReturn := new(Linebaseentitylisting)
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
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
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
	var successPayload *Linebaseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Linebaseentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesLines invokes GET /api/v2/telephony/providers/edges/lines
//
// Get a list of Lines
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesLines(pageSize int, pageNumber int, name string, sortBy string, expand []string) (*Lineentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/lines"
	defaultReturn := new(Lineentitylisting)
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
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
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
	var successPayload *Lineentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Lineentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesLinesTemplate invokes GET /api/v2/telephony/providers/edges/lines/template
//
// Get a Line instance template based on a Line Base Settings object. This object can then be modified and saved as a new Line instance
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesLinesTemplate(lineBaseSettingsId string) (*Line, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/lines/template"
	defaultReturn := new(Line)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'lineBaseSettingsId' is set
	if &lineBaseSettingsId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'lineBaseSettingsId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesLinesTemplate")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["lineBaseSettingsId"] = a.Configuration.APIClient.ParameterToString(lineBaseSettingsId, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Line
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Line" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesLogicalinterfaces invokes GET /api/v2/telephony/providers/edges/logicalinterfaces
//
// Get edge logical interfaces.
//
// Retrieve the configured logical interfaces for a list edges. Only 100 edges can be requested at a time.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesLogicalinterfaces(edgeIds string, expand []string) (*Logicalinterfaceentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/logicalinterfaces"
	defaultReturn := new(Logicalinterfaceentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeIds' is set
	if &edgeIds == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeIds' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesLogicalinterfaces")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["edgeIds"] = a.Configuration.APIClient.ParameterToString(edgeIds, "")
	
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
	var successPayload *Logicalinterfaceentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Logicalinterfaceentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesMetrics invokes GET /api/v2/telephony/providers/edges/metrics
//
// Get the metrics for a list of edges.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesMetrics(edgeIds string) ([]Edgemetrics, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/metrics"
	defaultReturn := make([]Edgemetrics, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeIds' is set
	if &edgeIds == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeIds' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesMetrics")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["edgeIds"] = a.Configuration.APIClient.ParameterToString(edgeIds, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload []Edgemetrics
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "[]Edgemetrics" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesOutboundroute invokes GET /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}
//
// Get outbound route
//
// This route is deprecated, use /telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId} instead.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesOutboundroute(outboundRouteId string) (*Outboundroute, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}"
	path = strings.Replace(path, "{outboundRouteId}", fmt.Sprintf("%v", outboundRouteId), -1)
	defaultReturn := new(Outboundroute)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'outboundRouteId' is set
	if &outboundRouteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'outboundRouteId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesOutboundroute")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Outboundroute
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Outboundroute" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesOutboundroutes invokes GET /api/v2/telephony/providers/edges/outboundroutes
//
// Get outbound routes
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesOutboundroutes(pageSize int, pageNumber int, name string, siteId string, externalTrunkBasesIds string, sortBy string) (*Outboundrouteentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/outboundroutes"
	defaultReturn := new(Outboundrouteentitylisting)
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
	
	queryParams["siteId"] = a.Configuration.APIClient.ParameterToString(siteId, "")
	
	queryParams["externalTrunkBasesIds"] = a.Configuration.APIClient.ParameterToString(externalTrunkBasesIds, "")
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Outboundrouteentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Outboundrouteentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesPhone invokes GET /api/v2/telephony/providers/edges/phones/{phoneId}
//
// Get a Phone by ID
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesPhone(phoneId string) (*Phone, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phones/{phoneId}"
	path = strings.Replace(path, "{phoneId}", fmt.Sprintf("%v", phoneId), -1)
	defaultReturn := new(Phone)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'phoneId' is set
	if &phoneId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'phoneId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesPhone")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Phone
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Phone" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesPhonebasesetting invokes GET /api/v2/telephony/providers/edges/phonebasesettings/{phoneBaseId}
//
// Get a Phone Base Settings object by ID
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesPhonebasesetting(phoneBaseId string) (*Phonebase, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phonebasesettings/{phoneBaseId}"
	path = strings.Replace(path, "{phoneBaseId}", fmt.Sprintf("%v", phoneBaseId), -1)
	defaultReturn := new(Phonebase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'phoneBaseId' is set
	if &phoneBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'phoneBaseId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesPhonebasesetting")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Phonebase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Phonebase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesPhonebasesettings invokes GET /api/v2/telephony/providers/edges/phonebasesettings
//
// Get a list of Phone Base Settings objects
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesPhonebasesettings(pageSize int, pageNumber int, sortBy string, sortOrder string, expand []string, name string) (*Phonebaseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phonebasesettings"
	defaultReturn := new(Phonebaseentitylisting)
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
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	
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
	var successPayload *Phonebaseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Phonebaseentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesPhonebasesettingsAvailablemetabases invokes GET /api/v2/telephony/providers/edges/phonebasesettings/availablemetabases
//
// Get a list of available makes and models to create a new Phone Base Settings
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesPhonebasesettingsAvailablemetabases(pageSize int, pageNumber int) (*Phonemetabaseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phonebasesettings/availablemetabases"
	defaultReturn := new(Phonemetabaseentitylisting)
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Phonemetabaseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Phonemetabaseentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesPhonebasesettingsTemplate invokes GET /api/v2/telephony/providers/edges/phonebasesettings/template
//
// Get a Phone Base Settings instance template from a given make and model. This object can then be modified and saved as a new Phone Base Settings instance
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesPhonebasesettingsTemplate(phoneMetabaseId string) (*Phonebase, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phonebasesettings/template"
	defaultReturn := new(Phonebase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'phoneMetabaseId' is set
	if &phoneMetabaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'phoneMetabaseId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesPhonebasesettingsTemplate")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["phoneMetabaseId"] = a.Configuration.APIClient.ParameterToString(phoneMetabaseId, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Phonebase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Phonebase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesPhones invokes GET /api/v2/telephony/providers/edges/phones
//
// Get a list of Phone Instances
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesPhones(pageNumber int, pageSize int, sortBy string, sortOrder string, siteId string, webRtcUserId string, phoneBaseSettingsId string, linesLoggedInUserId string, linesDefaultForUserId string, phoneHardwareId string, linesId string, linesName string, name string, statusOperationalStatus string, secondaryStatusOperationalStatus string, expand []string, fields []string) (*Phoneentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phones"
	defaultReturn := new(Phoneentitylisting)
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
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
	queryParams["siteId"] = a.Configuration.APIClient.ParameterToString(siteId, "")
	
	queryParams["webRtcUserId"] = a.Configuration.APIClient.ParameterToString(webRtcUserId, "")
	
	queryParams["phoneBaseSettingsId"] = a.Configuration.APIClient.ParameterToString(phoneBaseSettingsId, "")
	
	queryParams["linesLoggedInUserId"] = a.Configuration.APIClient.ParameterToString(linesLoggedInUserId, "")
	
	queryParams["linesDefaultForUserId"] = a.Configuration.APIClient.ParameterToString(linesDefaultForUserId, "")
	
	queryParams["phoneHardwareId"] = a.Configuration.APIClient.ParameterToString(phoneHardwareId, "")
	
	queryParams["linesId"] = a.Configuration.APIClient.ParameterToString(linesId, "")
	
	queryParams["linesName"] = a.Configuration.APIClient.ParameterToString(linesName, "")
	
	queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, "")
	
	queryParams["statusOperationalStatus"] = a.Configuration.APIClient.ParameterToString(statusOperationalStatus, "")
	
	queryParams["secondaryStatusOperationalStatus"] = a.Configuration.APIClient.ParameterToString(secondaryStatusOperationalStatus, "")
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	
	queryParams["fields"] = a.Configuration.APIClient.ParameterToString(fields, "multi")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Phoneentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Phoneentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesPhonesTemplate invokes GET /api/v2/telephony/providers/edges/phones/template
//
// Get a Phone instance template based on a Phone Base Settings object. This object can then be modified and saved as a new Phone instance
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesPhonesTemplate(phoneBaseSettingsId string) (*Phone, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phones/template"
	defaultReturn := new(Phone)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'phoneBaseSettingsId' is set
	if &phoneBaseSettingsId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'phoneBaseSettingsId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesPhonesTemplate")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["phoneBaseSettingsId"] = a.Configuration.APIClient.ParameterToString(phoneBaseSettingsId, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Phone
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Phone" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesPhysicalinterfaces invokes GET /api/v2/telephony/providers/edges/physicalinterfaces
//
// Get physical interfaces for edges.
//
// Retrieves a list of all configured physical interfaces for a list of edges. Only 100 edges can be requested at a time.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesPhysicalinterfaces(edgeIds string) (*Physicalinterfaceentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/physicalinterfaces"
	defaultReturn := new(Physicalinterfaceentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeIds' is set
	if &edgeIds == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeIds' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesPhysicalinterfaces")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["edgeIds"] = a.Configuration.APIClient.ParameterToString(edgeIds, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Physicalinterfaceentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Physicalinterfaceentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesSite invokes GET /api/v2/telephony/providers/edges/sites/{siteId}
//
// Get a Site by ID.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesSite(siteId string) (*Site, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
	defaultReturn := new(Site)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'siteId' is set
	if &siteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'siteId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesSite")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Site
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Site" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesSiteNumberplan invokes GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}
//
// Get a Number Plan by ID.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesSiteNumberplan(siteId string, numberPlanId string) (*Numberplan, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites/{siteId}/numberplans/{numberPlanId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
	path = strings.Replace(path, "{numberPlanId}", fmt.Sprintf("%v", numberPlanId), -1)
	defaultReturn := new(Numberplan)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'siteId' is set
	if &siteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'siteId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesSiteNumberplan")
	}
	// verify the required parameter 'numberPlanId' is set
	if &numberPlanId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'numberPlanId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesSiteNumberplan")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Numberplan
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Numberplan" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesSiteNumberplans invokes GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans
//
// Get the list of Number Plans for this Site. Only fetches the first 200 records.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesSiteNumberplans(siteId string) ([]Numberplan, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites/{siteId}/numberplans"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
	defaultReturn := make([]Numberplan, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'siteId' is set
	if &siteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'siteId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesSiteNumberplans")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload []Numberplan
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "[]Numberplan" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesSiteNumberplansClassifications invokes GET /api/v2/telephony/providers/edges/sites/{siteId}/numberplans/classifications
//
// Get a list of Classifications for this Site
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesSiteNumberplansClassifications(siteId string, classification string) ([]string, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites/{siteId}/numberplans/classifications"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
	defaultReturn := make([]string, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'siteId' is set
	if &siteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'siteId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesSiteNumberplansClassifications")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["classification"] = a.Configuration.APIClient.ParameterToString(classification, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload []string
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "[]string" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesSiteOutboundroute invokes GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}
//
// Get an outbound route
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesSiteOutboundroute(siteId string, outboundRouteId string) (*Outboundroutebase, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
	path = strings.Replace(path, "{outboundRouteId}", fmt.Sprintf("%v", outboundRouteId), -1)
	defaultReturn := new(Outboundroutebase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'siteId' is set
	if &siteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'siteId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesSiteOutboundroute")
	}
	// verify the required parameter 'outboundRouteId' is set
	if &outboundRouteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'outboundRouteId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesSiteOutboundroute")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Outboundroutebase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Outboundroutebase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesSiteOutboundroutes invokes GET /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes
//
// Get outbound routes
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesSiteOutboundroutes(siteId string, pageSize int, pageNumber int, name string, externalTrunkBasesIds string, sortBy string) (*Outboundroutebaseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
	defaultReturn := new(Outboundroutebaseentitylisting)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'siteId' is set
	if &siteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'siteId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesSiteOutboundroutes")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	
	queryParams["externalTrunkBasesIds"] = a.Configuration.APIClient.ParameterToString(externalTrunkBasesIds, "")
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Outboundroutebaseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Outboundroutebaseentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesSites invokes GET /api/v2/telephony/providers/edges/sites
//
// Get the list of Sites.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesSites(pageSize int, pageNumber int, sortBy string, sortOrder string, name string, locationId string, managed bool) (*Siteentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites"
	defaultReturn := new(Siteentitylisting)
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
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
	queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, "")
	
	queryParams["locationId"] = a.Configuration.APIClient.ParameterToString(locationId, "")
	
	queryParams["managed"] = a.Configuration.APIClient.ParameterToString(managed, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Siteentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Siteentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesTimezones invokes GET /api/v2/telephony/providers/edges/timezones
//
// Get a list of Edge-compatible time zones
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesTimezones(pageSize int, pageNumber int) (*Timezoneentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/timezones"
	defaultReturn := new(Timezoneentitylisting)
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
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Timezoneentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Timezoneentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesTrunk invokes GET /api/v2/telephony/providers/edges/trunks/{trunkId}
//
// Get a Trunk by ID
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesTrunk(trunkId string) (*Trunk, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/trunks/{trunkId}"
	path = strings.Replace(path, "{trunkId}", fmt.Sprintf("%v", trunkId), -1)
	defaultReturn := new(Trunk)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'trunkId' is set
	if &trunkId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'trunkId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesTrunk")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Trunk
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Trunk" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesTrunkMetrics invokes GET /api/v2/telephony/providers/edges/trunks/{trunkId}/metrics
//
// Get the trunk metrics.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesTrunkMetrics(trunkId string) (*Trunkmetrics, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/trunks/{trunkId}/metrics"
	path = strings.Replace(path, "{trunkId}", fmt.Sprintf("%v", trunkId), -1)
	defaultReturn := new(Trunkmetrics)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'trunkId' is set
	if &trunkId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'trunkId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesTrunkMetrics")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Trunkmetrics
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Trunkmetrics" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesTrunkbasesetting invokes GET /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}
//
// Get a Trunk Base Settings object by ID
//
// Managed properties will not be returned unless the user is assigned the internal:trunk:edit permission.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesTrunkbasesetting(trunkBaseSettingsId string, ignoreHidden bool) (*Trunkbase, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}"
	path = strings.Replace(path, "{trunkBaseSettingsId}", fmt.Sprintf("%v", trunkBaseSettingsId), -1)
	defaultReturn := new(Trunkbase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'trunkBaseSettingsId' is set
	if &trunkBaseSettingsId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'trunkBaseSettingsId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesTrunkbasesetting")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["ignoreHidden"] = a.Configuration.APIClient.ParameterToString(ignoreHidden, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Trunkbase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Trunkbase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesTrunkbasesettings invokes GET /api/v2/telephony/providers/edges/trunkbasesettings
//
// Get Trunk Base Settings listing
//
// Managed properties will not be returned unless the user is assigned the internal:trunk:edit permission.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesTrunkbasesettings(pageNumber int, pageSize int, sortBy string, sortOrder string, recordingEnabled bool, ignoreHidden bool, managed bool, expand []string, name string) (*Trunkbaseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/trunkbasesettings"
	defaultReturn := new(Trunkbaseentitylisting)
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
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
	queryParams["recordingEnabled"] = a.Configuration.APIClient.ParameterToString(recordingEnabled, "")
	
	queryParams["ignoreHidden"] = a.Configuration.APIClient.ParameterToString(ignoreHidden, "")
	
	queryParams["managed"] = a.Configuration.APIClient.ParameterToString(managed, "")
	
	queryParams["expand"] = a.Configuration.APIClient.ParameterToString(expand, "multi")
	
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
	var successPayload *Trunkbaseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Trunkbaseentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesTrunkbasesettingsAvailablemetabases invokes GET /api/v2/telephony/providers/edges/trunkbasesettings/availablemetabases
//
// Get a list of available makes and models to create a new Trunk Base Settings
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesTrunkbasesettingsAvailablemetabases(varType string, pageSize int, pageNumber int) (*Trunkmetabaseentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/trunkbasesettings/availablemetabases"
	defaultReturn := new(Trunkmetabaseentitylisting)
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
	
	queryParams["varType"] = a.Configuration.APIClient.ParameterToString(varType, "")
	
	queryParams["pageSize"] = a.Configuration.APIClient.ParameterToString(pageSize, "")
	
	queryParams["pageNumber"] = a.Configuration.APIClient.ParameterToString(pageNumber, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Trunkmetabaseentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Trunkmetabaseentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesTrunkbasesettingsTemplate invokes GET /api/v2/telephony/providers/edges/trunkbasesettings/template
//
// Get a Trunk Base Settings instance template from a given make and model. This object can then be modified and saved as a new Trunk Base Settings instance
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesTrunkbasesettingsTemplate(trunkMetabaseId string) (*Trunkbase, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/trunkbasesettings/template"
	defaultReturn := new(Trunkbase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'trunkMetabaseId' is set
	if &trunkMetabaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'trunkMetabaseId' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesTrunkbasesettingsTemplate")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["trunkMetabaseId"] = a.Configuration.APIClient.ParameterToString(trunkMetabaseId, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Trunkbase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Trunkbase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesTrunks invokes GET /api/v2/telephony/providers/edges/trunks
//
// Get the list of available trunks.
//
// Trunks are created by assigning trunk base settings to an Edge or Edge Group.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesTrunks(pageNumber int, pageSize int, sortBy string, sortOrder string, edgeId string, trunkBaseId string, trunkType string) (*Trunkentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/trunks"
	defaultReturn := new(Trunkentitylisting)
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
	
	queryParams["sortBy"] = a.Configuration.APIClient.ParameterToString(sortBy, "")
	
	queryParams["sortOrder"] = a.Configuration.APIClient.ParameterToString(sortOrder, "")
	
	queryParams["edgeId"] = a.Configuration.APIClient.ParameterToString(edgeId, "")
	
	queryParams["trunkBaseId"] = a.Configuration.APIClient.ParameterToString(trunkBaseId, "")
	
	queryParams["trunkType"] = a.Configuration.APIClient.ParameterToString(trunkType, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Trunkentitylisting
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Trunkentitylisting" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesTrunksMetrics invokes GET /api/v2/telephony/providers/edges/trunks/metrics
//
// Get the metrics for a list of trunks.
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesTrunksMetrics(trunkIds string) ([]Trunkmetrics, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/trunks/metrics"
	defaultReturn := make([]Trunkmetrics, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'trunkIds' is set
	if &trunkIds == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'trunkIds' when calling TelephonyProvidersEdgeApi->GetTelephonyProvidersEdgesTrunksMetrics")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["trunkIds"] = a.Configuration.APIClient.ParameterToString(trunkIds, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload []Trunkmetrics
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "[]Trunkmetrics" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// GetTelephonyProvidersEdgesTrunkswithrecording invokes GET /api/v2/telephony/providers/edges/trunkswithrecording
//
// Get Counts of trunks that have recording disabled or enabled
func (a TelephonyProvidersEdgeApi) GetTelephonyProvidersEdgesTrunkswithrecording(trunkType string) (*Trunkrecordingenabledcount, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/trunkswithrecording"
	defaultReturn := new(Trunkrecordingenabledcount)
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
	
	queryParams["trunkType"] = a.Configuration.APIClient.ParameterToString(trunkType, "")
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload *Trunkrecordingenabledcount
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Trunkrecordingenabledcount" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgeDiagnosticNslookup invokes POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup
//
// Nslookup request command to collect networking-related information from an Edge for a target IP or host.
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgeDiagnosticNslookup(edgeId string, body Edgenetworkdiagnosticrequest) (*Edgenetworkdiagnostic, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/diagnostic/nslookup"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Edgenetworkdiagnostic)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeDiagnosticNslookup")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeDiagnosticNslookup")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Edgenetworkdiagnostic
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgenetworkdiagnostic" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgeDiagnosticPing invokes POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping
//
// Ping Request command to collect networking-related information from an Edge for a target IP or host.
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgeDiagnosticPing(edgeId string, body Edgenetworkdiagnosticrequest) (*Edgenetworkdiagnostic, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/diagnostic/ping"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Edgenetworkdiagnostic)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeDiagnosticPing")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeDiagnosticPing")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Edgenetworkdiagnostic
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgenetworkdiagnostic" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgeDiagnosticRoute invokes POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/route
//
// Route request command to collect networking-related information from an Edge for a target IP or host.
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgeDiagnosticRoute(edgeId string, body Edgenetworkdiagnosticrequest) (*Edgenetworkdiagnostic, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/diagnostic/route"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Edgenetworkdiagnostic)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeDiagnosticRoute")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeDiagnosticRoute")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Edgenetworkdiagnostic
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgenetworkdiagnostic" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgeDiagnosticTracepath invokes POST /api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath
//
// Tracepath request command to collect networking-related information from an Edge for a target IP or host.
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgeDiagnosticTracepath(edgeId string, body Edgenetworkdiagnosticrequest) (*Edgenetworkdiagnostic, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Edgenetworkdiagnostic)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeDiagnosticTracepath")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeDiagnosticTracepath")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Edgenetworkdiagnostic
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgenetworkdiagnostic" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgeLogicalinterfaces invokes POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces
//
// Create an edge logical interface.
//
// Create
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgeLogicalinterfaces(edgeId string, body Domainlogicalinterface) (*Domainlogicalinterface, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Domainlogicalinterface)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeLogicalinterfaces")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeLogicalinterfaces")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Domainlogicalinterface
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Domainlogicalinterface" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgeLogsJobUpload invokes POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload
//
// Request that the specified fileIds be uploaded from the Edge.
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgeLogsJobUpload(edgeId string, jobId string, body Edgelogsjobuploadrequest) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/logs/jobs/{jobId}/upload"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	path = strings.Replace(path, "{jobId}", fmt.Sprintf("%v", jobId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeLogsJobUpload")
	}
	// verify the required parameter 'jobId' is set
	if &jobId == nil {
		// false
		return nil, errors.New("Missing required parameter 'jobId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeLogsJobUpload")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeLogsJobUpload")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// PostTelephonyProvidersEdgeLogsJobs invokes POST /api/v2/telephony/providers/edges/{edgeId}/logs/jobs
//
// Create a job to upload a list of Edge logs.
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgeLogsJobs(edgeId string, body Edgelogsjobrequest) (*Edgelogsjobresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/logs/jobs"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Edgelogsjobresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeLogsJobs")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeLogsJobs")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Edgelogsjobresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgelogsjobresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgeReboot invokes POST /api/v2/telephony/providers/edges/{edgeId}/reboot
//
// Reboot an Edge
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgeReboot(edgeId string, body Edgerebootparameters) (*string, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/reboot"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(string)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeReboot")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *string
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "string" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgeSoftwareupdate invokes POST /api/v2/telephony/providers/edges/{edgeId}/softwareupdate
//
// Starts a software update for this edge.
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgeSoftwareupdate(edgeId string, body Domainedgesoftwareupdatedto) (*Domainedgesoftwareupdatedto, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/softwareupdate"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Domainedgesoftwareupdatedto)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeSoftwareupdate")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeSoftwareupdate")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Domainedgesoftwareupdatedto
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Domainedgesoftwareupdatedto" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgeStatuscode invokes POST /api/v2/telephony/providers/edges/{edgeId}/statuscode
//
// Take an Edge in or out of service
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgeStatuscode(edgeId string, body Edgeservicestaterequest) (*string, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/statuscode"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(string)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeStatuscode")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *string
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "string" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgeUnpair invokes POST /api/v2/telephony/providers/edges/{edgeId}/unpair
//
// Unpair an Edge
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgeUnpair(edgeId string) (*string, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/unpair"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(string)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgeUnpair")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *string
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "string" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdges invokes POST /api/v2/telephony/providers/edges
//
// Create an edge.
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdges(body Edge) (*Edge, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges"
	defaultReturn := new(Edge)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdges")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Edge
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edge" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgesAddressvalidation invokes POST /api/v2/telephony/providers/edges/addressvalidation
//
// Validates a street address
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesAddressvalidation(body Validateaddressrequest) (*Validateaddressresponse, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/addressvalidation"
	defaultReturn := new(Validateaddressresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesAddressvalidation")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Validateaddressresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Validateaddressresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgesCertificateauthorities invokes POST /api/v2/telephony/providers/edges/certificateauthorities
//
// Create a certificate authority.
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesCertificateauthorities(body Domaincertificateauthority) (*Domaincertificateauthority, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/certificateauthorities"
	defaultReturn := new(Domaincertificateauthority)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesCertificateauthorities")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Domaincertificateauthority
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Domaincertificateauthority" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgesDidpools invokes POST /api/v2/telephony/providers/edges/didpools
//
// Create a new DID pool
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesDidpools(body Didpool) (*Didpool, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/didpools"
	defaultReturn := new(Didpool)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesDidpools")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Didpool
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Didpool" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgesEdgegroups invokes POST /api/v2/telephony/providers/edges/edgegroups
//
// Create an edge group.
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesEdgegroups(body Edgegroup) (*Edgegroup, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/edgegroups"
	defaultReturn := new(Edgegroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesEdgegroups")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Edgegroup
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgegroup" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgesExtensionpools invokes POST /api/v2/telephony/providers/edges/extensionpools
//
// Create a new extension pool
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesExtensionpools(body Extensionpool) (*Extensionpool, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/extensionpools"
	defaultReturn := new(Extensionpool)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesExtensionpools")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Extensionpool
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Extensionpool" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgesOutboundroutes invokes POST /api/v2/telephony/providers/edges/outboundroutes
//
// Create outbound rule
//
// This route is deprecated, use /telephony/providers/edges/sites/{siteId}/outboundroutes instead.
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesOutboundroutes(body Outboundroute) (*Outboundroute, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/outboundroutes"
	defaultReturn := new(Outboundroute)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesOutboundroutes")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Outboundroute
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Outboundroute" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgesPhoneReboot invokes POST /api/v2/telephony/providers/edges/phones/{phoneId}/reboot
//
// Reboot a Phone
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesPhoneReboot(phoneId string) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phones/{phoneId}/reboot"
	path = strings.Replace(path, "{phoneId}", fmt.Sprintf("%v", phoneId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'phoneId' is set
	if &phoneId == nil {
		// false
		return nil, errors.New("Missing required parameter 'phoneId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesPhoneReboot")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// PostTelephonyProvidersEdgesPhonebasesettings invokes POST /api/v2/telephony/providers/edges/phonebasesettings
//
// Create a new Phone Base Settings object
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesPhonebasesettings(body Phonebase) (*Phonebase, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phonebasesettings"
	defaultReturn := new(Phonebase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesPhonebasesettings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Phonebase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Phonebase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgesPhones invokes POST /api/v2/telephony/providers/edges/phones
//
// Create a new Phone
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesPhones(body Phone) (*Phone, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phones"
	defaultReturn := new(Phone)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesPhones")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Phone
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Phone" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgesPhonesReboot invokes POST /api/v2/telephony/providers/edges/phones/reboot
//
// Reboot Multiple Phones
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesPhonesReboot(body Phonesreboot) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phones/reboot"
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesPhonesReboot")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// PostTelephonyProvidersEdgesSiteOutboundroutes invokes POST /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes
//
// Create outbound route
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesSiteOutboundroutes(siteId string, body Outboundroutebase) (*Outboundroutebase, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
	defaultReturn := new(Outboundroutebase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'siteId' is set
	if &siteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'siteId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesSiteOutboundroutes")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesSiteOutboundroutes")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Outboundroutebase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Outboundroutebase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgesSiteRebalance invokes POST /api/v2/telephony/providers/edges/sites/{siteId}/rebalance
//
// Triggers the rebalance operation.
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesSiteRebalance(siteId string) (*APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites/{siteId}/rebalance"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'siteId' is set
	if &siteId == nil {
		// false
		return nil, errors.New("Missing required parameter 'siteId' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesSiteRebalance")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// PostTelephonyProvidersEdgesSites invokes POST /api/v2/telephony/providers/edges/sites
//
// Create a Site.
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesSites(body Site) (*Site, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites"
	defaultReturn := new(Site)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesSites")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Site
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Site" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PostTelephonyProvidersEdgesTrunkbasesettings invokes POST /api/v2/telephony/providers/edges/trunkbasesettings
//
// Create a Trunk Base Settings object
func (a TelephonyProvidersEdgeApi) PostTelephonyProvidersEdgesTrunkbasesettings(body Trunkbase) (*Trunkbase, *APIResponse, error) {
	var httpMethod = "POST"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/trunkbasesettings"
	defaultReturn := new(Trunkbase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PostTelephonyProvidersEdgesTrunkbasesettings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Trunkbase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Trunkbase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdge invokes PUT /api/v2/telephony/providers/edges/{edgeId}
//
// Update a edge.
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdge(edgeId string, body Edge) (*Edge, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	defaultReturn := new(Edge)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdge")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdge")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Edge
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edge" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgeLine invokes PUT /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}
//
// Update a line.
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgeLine(edgeId string, lineId string, body Edgeline) (*Edgeline, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	path = strings.Replace(path, "{lineId}", fmt.Sprintf("%v", lineId), -1)
	defaultReturn := new(Edgeline)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgeLine")
	}
	// verify the required parameter 'lineId' is set
	if &lineId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'lineId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgeLine")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgeLine")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Edgeline
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgeline" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgeLogicalinterface invokes PUT /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}
//
// Update an edge logical interface.
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgeLogicalinterface(edgeId string, interfaceId string, body Domainlogicalinterface) (*Domainlogicalinterface, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces/{interfaceId}"
	path = strings.Replace(path, "{edgeId}", fmt.Sprintf("%v", edgeId), -1)
	path = strings.Replace(path, "{interfaceId}", fmt.Sprintf("%v", interfaceId), -1)
	defaultReturn := new(Domainlogicalinterface)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeId' is set
	if &edgeId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgeLogicalinterface")
	}
	// verify the required parameter 'interfaceId' is set
	if &interfaceId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'interfaceId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgeLogicalinterface")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgeLogicalinterface")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Domainlogicalinterface
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Domainlogicalinterface" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesCertificateauthority invokes PUT /api/v2/telephony/providers/edges/certificateauthorities/{certificateId}
//
// Update a certificate authority.
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesCertificateauthority(certificateId string, body Domaincertificateauthority) (*Domaincertificateauthority, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/certificateauthorities/{certificateId}"
	path = strings.Replace(path, "{certificateId}", fmt.Sprintf("%v", certificateId), -1)
	defaultReturn := new(Domaincertificateauthority)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'certificateId' is set
	if &certificateId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'certificateId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesCertificateauthority")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesCertificateauthority")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Domaincertificateauthority
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Domaincertificateauthority" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesDid invokes PUT /api/v2/telephony/providers/edges/dids/{didId}
//
// Update a DID by ID.
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesDid(didId string, body Did) (*Did, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/dids/{didId}"
	path = strings.Replace(path, "{didId}", fmt.Sprintf("%v", didId), -1)
	defaultReturn := new(Did)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'didId' is set
	if &didId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'didId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesDid")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesDid")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Did
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Did" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesDidpool invokes PUT /api/v2/telephony/providers/edges/didpools/{didPoolId}
//
// Update a DID Pool by ID.
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesDidpool(didPoolId string, body Didpool) (*Didpool, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/didpools/{didPoolId}"
	path = strings.Replace(path, "{didPoolId}", fmt.Sprintf("%v", didPoolId), -1)
	defaultReturn := new(Didpool)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'didPoolId' is set
	if &didPoolId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'didPoolId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesDidpool")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesDidpool")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Didpool
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Didpool" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesEdgegroup invokes PUT /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}
//
// Update an edge group.
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesEdgegroup(edgeGroupId string, body Edgegroup) (*Edgegroup, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}"
	path = strings.Replace(path, "{edgeGroupId}", fmt.Sprintf("%v", edgeGroupId), -1)
	defaultReturn := new(Edgegroup)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgeGroupId' is set
	if &edgeGroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgeGroupId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesEdgegroup")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesEdgegroup")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Edgegroup
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgegroup" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesEdgegroupEdgetrunkbase invokes PUT /api/v2/telephony/providers/edges/edgegroups/{edgegroupId}/edgetrunkbases/{edgetrunkbaseId}
//
// Update the edge trunk base associated with the edge group
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesEdgegroupEdgetrunkbase(edgegroupId string, edgetrunkbaseId string, body Edgetrunkbase) (*Edgetrunkbase, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/edgegroups/{edgegroupId}/edgetrunkbases/{edgetrunkbaseId}"
	path = strings.Replace(path, "{edgegroupId}", fmt.Sprintf("%v", edgegroupId), -1)
	path = strings.Replace(path, "{edgetrunkbaseId}", fmt.Sprintf("%v", edgetrunkbaseId), -1)
	defaultReturn := new(Edgetrunkbase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'edgegroupId' is set
	if &edgegroupId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgegroupId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesEdgegroupEdgetrunkbase")
	}
	// verify the required parameter 'edgetrunkbaseId' is set
	if &edgetrunkbaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'edgetrunkbaseId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesEdgegroupEdgetrunkbase")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesEdgegroupEdgetrunkbase")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Edgetrunkbase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Edgetrunkbase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesExtension invokes PUT /api/v2/telephony/providers/edges/extensions/{extensionId}
//
// Update an extension by ID.
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesExtension(extensionId string, body Extension) (*Extension, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/extensions/{extensionId}"
	path = strings.Replace(path, "{extensionId}", fmt.Sprintf("%v", extensionId), -1)
	defaultReturn := new(Extension)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'extensionId' is set
	if &extensionId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'extensionId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesExtension")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesExtension")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Extension
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Extension" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesExtensionpool invokes PUT /api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}
//
// Update an extension pool by ID
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesExtensionpool(extensionPoolId string, body Extensionpool) (*Extensionpool, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/extensionpools/{extensionPoolId}"
	path = strings.Replace(path, "{extensionPoolId}", fmt.Sprintf("%v", extensionPoolId), -1)
	defaultReturn := new(Extensionpool)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'extensionPoolId' is set
	if &extensionPoolId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'extensionPoolId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesExtensionpool")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesExtensionpool")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Extensionpool
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Extensionpool" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesOutboundroute invokes PUT /api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}
//
// Update outbound route
//
// This route is deprecated, use /telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId} instead.
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesOutboundroute(outboundRouteId string, body Outboundroute) (*Outboundroute, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/outboundroutes/{outboundRouteId}"
	path = strings.Replace(path, "{outboundRouteId}", fmt.Sprintf("%v", outboundRouteId), -1)
	defaultReturn := new(Outboundroute)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'outboundRouteId' is set
	if &outboundRouteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'outboundRouteId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesOutboundroute")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesOutboundroute")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Outboundroute
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Outboundroute" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesPhone invokes PUT /api/v2/telephony/providers/edges/phones/{phoneId}
//
// Update a Phone by ID
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesPhone(phoneId string, body Phone) (*Phone, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phones/{phoneId}"
	path = strings.Replace(path, "{phoneId}", fmt.Sprintf("%v", phoneId), -1)
	defaultReturn := new(Phone)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'phoneId' is set
	if &phoneId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'phoneId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesPhone")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesPhone")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Phone
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Phone" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesPhonebasesetting invokes PUT /api/v2/telephony/providers/edges/phonebasesettings/{phoneBaseId}
//
// Update a Phone Base Settings by ID
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesPhonebasesetting(phoneBaseId string, body Phonebase) (*Phonebase, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/phonebasesettings/{phoneBaseId}"
	path = strings.Replace(path, "{phoneBaseId}", fmt.Sprintf("%v", phoneBaseId), -1)
	defaultReturn := new(Phonebase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'phoneBaseId' is set
	if &phoneBaseId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'phoneBaseId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesPhonebasesetting")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesPhonebasesetting")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Phonebase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Phonebase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesSite invokes PUT /api/v2/telephony/providers/edges/sites/{siteId}
//
// Update a Site by ID.
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesSite(siteId string, body Site) (*Site, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
	defaultReturn := new(Site)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'siteId' is set
	if &siteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'siteId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesSite")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesSite")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Site
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Site" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesSiteNumberplans invokes PUT /api/v2/telephony/providers/edges/sites/{siteId}/numberplans
//
// Update the list of Number Plans. A user can update maximum 200 number plans at a time.
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesSiteNumberplans(siteId string, body []Numberplan) ([]Numberplan, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites/{siteId}/numberplans"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
	defaultReturn := make([]Numberplan, 0)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'siteId' is set
	if &siteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'siteId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesSiteNumberplans")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// true
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesSiteNumberplans")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload []Numberplan
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "[]Numberplan" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesSiteOutboundroute invokes PUT /api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}
//
// Update outbound route
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesSiteOutboundroute(siteId string, outboundRouteId string, body Outboundroutebase) (*Outboundroutebase, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/sites/{siteId}/outboundroutes/{outboundRouteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteId), -1)
	path = strings.Replace(path, "{outboundRouteId}", fmt.Sprintf("%v", outboundRouteId), -1)
	defaultReturn := new(Outboundroutebase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'siteId' is set
	if &siteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'siteId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesSiteOutboundroute")
	}
	// verify the required parameter 'outboundRouteId' is set
	if &outboundRouteId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'outboundRouteId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesSiteOutboundroute")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesSiteOutboundroute")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Outboundroutebase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Outboundroutebase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

// PutTelephonyProvidersEdgesTrunkbasesetting invokes PUT /api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}
//
// Update a Trunk Base Settings object by ID
func (a TelephonyProvidersEdgeApi) PutTelephonyProvidersEdgesTrunkbasesetting(trunkBaseSettingsId string, body Trunkbase) (*Trunkbase, *APIResponse, error) {
	var httpMethod = "PUT"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/telephony/providers/edges/trunkbasesettings/{trunkBaseSettingsId}"
	path = strings.Replace(path, "{trunkBaseSettingsId}", fmt.Sprintf("%v", trunkBaseSettingsId), -1)
	defaultReturn := new(Trunkbase)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'trunkBaseSettingsId' is set
	if &trunkBaseSettingsId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'trunkBaseSettingsId' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesTrunkbasesetting")
	}
	// verify the required parameter 'body' is set
	if &body == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling TelephonyProvidersEdgeApi->PutTelephonyProvidersEdgesTrunkbasesetting")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Trunkbase
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes)
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Trunkbase" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

