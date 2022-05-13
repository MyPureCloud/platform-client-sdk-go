package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustorauditqueryrequest
type Trustorauditqueryrequest struct { 
	// TrustorOrganizationId - Limit returned audits to this trustor organizationId.
	TrustorOrganizationId *string `json:"trustorOrganizationId,omitempty"`


	// TrusteeUserIds - Limit returned audits to these trustee userIds.
	TrusteeUserIds *[]string `json:"trusteeUserIds,omitempty"`


	// StartDate - Starting date/time for the audit search. ISO-8601 formatted date-time, UTC.
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - Ending date/time for the audit search. ISO-8601 formatted date-time, UTC.
	EndDate *time.Time `json:"endDate,omitempty"`


	// QueryPhrase - Word or phrase to look for in audit bodies.
	QueryPhrase *string `json:"queryPhrase,omitempty"`


	// Facets - Facet information to be returned with the query results.
	Facets *[]Facet `json:"facets,omitempty"`


	// Filters - Additional custom filters to be applied to the query.
	Filters *[]Filter `json:"filters,omitempty"`

}

func (o *Trustorauditqueryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trustorauditqueryrequest
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		TrustorOrganizationId *string `json:"trustorOrganizationId,omitempty"`
		
		TrusteeUserIds *[]string `json:"trusteeUserIds,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		QueryPhrase *string `json:"queryPhrase,omitempty"`
		
		Facets *[]Facet `json:"facets,omitempty"`
		
		Filters *[]Filter `json:"filters,omitempty"`
		*Alias
	}{ 
		TrustorOrganizationId: o.TrustorOrganizationId,
		
		TrusteeUserIds: o.TrusteeUserIds,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		QueryPhrase: o.QueryPhrase,
		
		Facets: o.Facets,
		
		Filters: o.Filters,
		Alias:    (*Alias)(o),
	})
}

func (o *Trustorauditqueryrequest) UnmarshalJSON(b []byte) error {
	var TrustorauditqueryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TrustorauditqueryrequestMap)
	if err != nil {
		return err
	}
	
	if TrustorOrganizationId, ok := TrustorauditqueryrequestMap["trustorOrganizationId"].(string); ok {
		o.TrustorOrganizationId = &TrustorOrganizationId
	}
    
	if TrusteeUserIds, ok := TrustorauditqueryrequestMap["trusteeUserIds"].([]interface{}); ok {
		TrusteeUserIdsString, _ := json.Marshal(TrusteeUserIds)
		json.Unmarshal(TrusteeUserIdsString, &o.TrusteeUserIds)
	}
	
	if startDateString, ok := TrustorauditqueryrequestMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := TrustorauditqueryrequestMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if QueryPhrase, ok := TrustorauditqueryrequestMap["queryPhrase"].(string); ok {
		o.QueryPhrase = &QueryPhrase
	}
    
	if Facets, ok := TrustorauditqueryrequestMap["facets"].([]interface{}); ok {
		FacetsString, _ := json.Marshal(Facets)
		json.Unmarshal(FacetsString, &o.Facets)
	}
	
	if Filters, ok := TrustorauditqueryrequestMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trustorauditqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
