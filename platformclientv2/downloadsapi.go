package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
	)

// DownloadsApi provides functions for API endpoints
type DownloadsApi struct {
	Configuration *Configuration
}

// NewDownloadsApi creates an API instance using the default configuration
func NewDownloadsApi() *DownloadsApi {
	fmt.Sprintf(strings.Title(""), "")
	config := GetDefaultConfiguration()
	return &DownloadsApi{
		Configuration: config,
	}
}

// NewDownloadsApiWithConfig creates an API instance using the provided configuration
func NewDownloadsApiWithConfig(config *Configuration) *DownloadsApi {
	return &DownloadsApi{
		Configuration: config,
	}
}

// GetDownload invokes GET /api/v2/downloads/{downloadId}
//
// Issues a redirect to a signed secure download URL for specified download
//
// this method will issue a redirect to the url to the content
func (a DownloadsApi) GetDownload(downloadId string, contentDisposition string, issueRedirect bool, redirectToAuth bool) (*Urlresponse, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/downloads/{downloadId}"
	path = strings.Replace(path, "{downloadId}", url.PathEscape(fmt.Sprintf("%v", downloadId)), -1)
	defaultReturn := new(Urlresponse)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'downloadId' is set
	if &downloadId == nil {
		// false
		return defaultReturn, nil, errors.New("Missing required parameter 'downloadId' when calling DownloadsApi->GetDownload")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
	// authentication (PureCloud OAuth) required

	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
	
	queryParams["contentDisposition"] = a.Configuration.APIClient.ParameterToString(contentDisposition, "")
	
	queryParams["issueRedirect"] = a.Configuration.APIClient.ParameterToString(issueRedirect, "")
	
	queryParams["redirectToAuth"] = a.Configuration.APIClient.ParameterToString(redirectToAuth, "")
	

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
	var successPayload *Urlresponse
	response, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, postFileName, fileBytes, "other")
	if err != nil {
		// Nothing special to do here, but do avoid processing the response
	} else if err == nil && response.Error != nil {
		err = errors.New(response.ErrorMessage)
	} else if response.HasBody {
		if "Urlresponse" == "string" {
			copy(response.RawBody, &successPayload)
		} else {
			err = json.Unmarshal(response.RawBody, &successPayload)
		}
	}
	return successPayload, response, err
}

