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

func (u *Trustorauditqueryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trustorauditqueryrequest

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if u.EndDate != nil {
		
		*EndDate = timeutil.Strftime(u.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		TrustorOrganizationId: u.TrustorOrganizationId,
		
		TrusteeUserIds: u.TrusteeUserIds,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		QueryPhrase: u.QueryPhrase,
		
		Facets: u.Facets,
		
		Filters: u.Filters,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trustorauditqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
