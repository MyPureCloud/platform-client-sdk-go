package platformclientv2
import (
	"time"
	"encoding/json"
)

// Trusteeauditqueryrequest
type Trusteeauditqueryrequest struct { 
	// TrusteeOrganizationIds - Limit returned audits to these trustee organizationIds.
	TrusteeOrganizationIds *[]string `json:"trusteeOrganizationIds,omitempty"`


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

// String returns a JSON representation of the model
func (o *Trusteeauditqueryrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
