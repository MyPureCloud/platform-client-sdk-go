package platformclientv2
import (
	"encoding/json"
)

// Auditsearchresult
type Auditsearchresult struct { 
	// PageNumber - Which page was returned.
	PageNumber *int `json:"pageNumber,omitempty"`


	// PageSize - The number of results in a page.
	PageSize *int `json:"pageSize,omitempty"`


	// Total - The total number of results.
	Total *int `json:"total,omitempty"`


	// PageCount - The number of pages of results.
	PageCount *int `json:"pageCount,omitempty"`


	// FacetInfo
	FacetInfo *[]Facetinfo `json:"facetInfo,omitempty"`


	// AuditMessages
	AuditMessages *[]Auditmessage `json:"auditMessages,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditsearchresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
