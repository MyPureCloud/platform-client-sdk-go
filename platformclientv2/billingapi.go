package platformclientv2

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"time"
"encoding/json"
)

// BillingApi provides functions for API endpoints
type BillingApi struct {
	Configuration *Configuration
}

// NewBillingApi creates an API instance using the default configuration
func NewBillingApi() *BillingApi {
	config := GetDefaultConfiguration()
	config.Debug(fmt.Sprintf("Creating BillingApi with base path: %s", strings.ToLower(config.BasePath)))
	return &BillingApi{
		Configuration: config,
	}
}

// NewBillingApiWithConfig creates an API instance using the provided configuration
func NewBillingApiWithConfig(config *Configuration) *BillingApi {
	config.Debugf("Creating BillingApi with base path: %s\n", strings.ToLower(config.BasePath))
	return &BillingApi{
		Configuration: config,
	}
}

// GetBillingReportsBillableusage invokes GET /api/v2/billing/reports/billableusage
//
// Get a report of the billable license usages
//
// Report is of the billable usages (e.g. licenses and devices utilized) for a given period. If response&#39;s status is InProgress, wait a few seconds, then try the same request again.
func (a BillingApi) GetBillingReportsBillableusage(startDate time.Time, endDate time.Time) (*Billingusagereport, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/billing/reports/billableusage"
	defaultReturn := new(Billingusagereport)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'startDate' is set
	if &startDate == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'startDate' when calling BillingApi->GetBillingReportsBillableusage")
	}
	// verify the required parameter 'endDate' is set
	if &endDate == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'endDate' when calling BillingApi->GetBillingReportsBillableusage")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	if str, ok := interface{}(startDate).(string); ok {
		if str != "" {
			queryParams["startDate"] = a.Configuration.APIClient.ParameterToString(startDate, collectionFormat)
		}
	} else {
		queryParams["startDate"] = a.Configuration.APIClient.ParameterToString(startDate, collectionFormat)
	}
	
	
	
	
	collectionFormat = ""
	if str, ok := interface{}(endDate).(string); ok {
		if str != "" {
			queryParams["endDate"] = a.Configuration.APIClient.ParameterToString(endDate, collectionFormat)
		}
	} else {
		queryParams["endDate"] = a.Configuration.APIClient.ParameterToString(endDate, collectionFormat)
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
	var successPayload *Billingusagereport
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

// GetBillingTrusteebillingoverviewTrustorOrgId invokes GET /api/v2/billing/trusteebillingoverview/{trustorOrgId}
//
// Get the billing overview for an organization that is managed by a partner.
//
// Tax Disclaimer: Prices returned by this API do not include applicable taxes. It is the responsibility of the customer to pay all taxes that are appropriate in their jurisdiction. See the PureCloud API Documentation in the Developer Center for more information about this API: https://developer.mypurecloud.com/api/rest/v2/
func (a BillingApi) GetBillingTrusteebillingoverviewTrustorOrgId(trustorOrgId string, billingPeriodIndex int) (*Trusteebillingoverview, *APIResponse, error) {
	var httpMethod = "GET"
	// create path and map variables
	path := a.Configuration.BasePath + "/api/v2/billing/trusteebillingoverview/{trustorOrgId}"
	path = strings.Replace(path, "{trustorOrgId}", fmt.Sprintf("%v", trustorOrgId), -1)
	defaultReturn := new(Trusteebillingoverview)
	if true == false {
		return defaultReturn, nil, errors.New("This message brought to you by the laws of physics being broken")
	}

	// verify the required parameter 'trustorOrgId' is set
	if &trustorOrgId == nil {
		// 
		return defaultReturn, nil, errors.New("Missing required parameter 'trustorOrgId' when calling BillingApi->GetBillingTrusteebillingoverviewTrustorOrgId")
	}

	headerParams := make(map[string]string)
	queryParams := make(map[string]string)
	formParams := url.Values{}
	var postBody interface{}
	var postFileName string
	var fileBytes []byte
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
	if str, ok := interface{}(billingPeriodIndex).(string); ok {
		if str != "" {
			queryParams["billingPeriodIndex"] = a.Configuration.APIClient.ParameterToString(billingPeriodIndex, collectionFormat)
		}
	} else {
		queryParams["billingPeriodIndex"] = a.Configuration.APIClient.ParameterToString(billingPeriodIndex, collectionFormat)
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
	var successPayload *Trusteebillingoverview
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

