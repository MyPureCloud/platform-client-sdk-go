package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

// StationsApi provides functions for API endpoints
type StationsApi struct {
	Configuration *Configuration
}

// NewStationsApi creates an API instance using the default configuration
func NewStationsApi() *StationsApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating StationsApi with base path: %s", strings.ToLower(config.BasePath)))
	return &StationsApi{
		Configuration: config,
	}
}

// NewStationsApiWithConfig creates an API instance using the provided configuration
func NewStationsApiWithConfig(config *Configuration) *StationsApi {
	config.Debugf("Creating StationsApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &StationsApi{
		Configuration: config,
	}
}

// DeleteStationAssociateduser invokes DELETE /api/v2/stations/{stationId}/associateduser
//
// Unassigns the user assigned to this station
//
// 
func (a StationsApi) DeleteStationAssociateduser(stationId string) (*APIResponse, error) {
	var httpMethod = "DELETE"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/stations/{stationId}/associateduser"
	path = strings.Replace(path, "{stationId}", fmt.Sprintf("%v", stationId), -1)
	if true == false {
		return nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'stationId' is set
	if &stationId == nil {
		// 
		return nil, errors.New("Missing required parameter 'stationId' when calling StationsApi->DeleteStationAssociateduser")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

// GetStation invokes GET /api/v2/stations/{stationId}
//
// Get station.
//
// 
func (a StationsApi) GetStation(stationId string) (*Station, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/stations/{stationId}"
	path = strings.Replace(path, "{stationId}", fmt.Sprintf("%v", stationId), -1)
	defaultReturn := new(Station)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'stationId' is set
	if &stationId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'stationId' when calling StationsApi->GetStation")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	var successPayload *Station
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

// GetStations invokes GET /api/v2/stations
//
// Get the list of available stations.
//
// 
func (a StationsApi) GetStations(pageSize int32, pageNumber int32, sortBy string, name string, userSelectable string, webRtcUserId string, id string, lineAppearanceId string) (*Stationentitylisting, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/stations"
	defaultReturn := new(Stationentitylisting)
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
	if str, ok := interface{}(name).(string); ok {
		if str != "" {
			queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
		}
	} else {
		queryParams["name"] = a.Configuration.APIClient.ParameterToString(name, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(userSelectable).(string); ok {
		if str != "" {
			queryParams["userSelectable"] = a.Configuration.APIClient.ParameterToString(userSelectable, collectionFormat)
		}
	} else {
		queryParams["userSelectable"] = a.Configuration.APIClient.ParameterToString(userSelectable, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(webRtcUserId).(string); ok {
		if str != "" {
			queryParams["webRtcUserId"] = a.Configuration.APIClient.ParameterToString(webRtcUserId, collectionFormat)
		}
	} else {
		queryParams["webRtcUserId"] = a.Configuration.APIClient.ParameterToString(webRtcUserId, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(id).(string); ok {
		if str != "" {
			queryParams["id"] = a.Configuration.APIClient.ParameterToString(id, collectionFormat)
		}
	} else {
		queryParams["id"] = a.Configuration.APIClient.ParameterToString(id, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(lineAppearanceId).(string); ok {
		if str != "" {
			queryParams["lineAppearanceId"] = a.Configuration.APIClient.ParameterToString(lineAppearanceId, collectionFormat)
		}
	} else {
		queryParams["lineAppearanceId"] = a.Configuration.APIClient.ParameterToString(lineAppearanceId, collectionFormat)
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
	var successPayload *Stationentitylisting
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

// GetStationsSettings invokes GET /api/v2/stations/settings
//
// Get an organization&#39;s StationSettings
//
// 
func (a StationsApi) GetStationsSettings() (*Stationsettings, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/stations/settings"
	defaultReturn := new(Stationsettings)
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
	var successPayload *Stationsettings
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

// PatchStationsSettings invokes PATCH /api/v2/stations/settings
//
// Patch an organization&#39;s StationSettings
//
// 
func (a StationsApi) PatchStationsSettings(body Stationsettings) (*Stationsettings, *APIResponse, error) {
	var httpMethod = "PATCH"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/stations/settings"
	defaultReturn := new(Stationsettings)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'body' is set
	if &body == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'body' when calling StationsApi->PatchStationsSettings")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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

	var successPayload *Stationsettings
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

