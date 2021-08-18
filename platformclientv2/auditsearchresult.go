package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Auditsearchresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditsearchresult

	

	return json.Marshal(&struct { 
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		FacetInfo *[]Facetinfo `json:"facetInfo,omitempty"`
		
		AuditMessages *[]Auditmessage `json:"auditMessages,omitempty"`
		*Alias
	}{ 
		PageNumber: u.PageNumber,
		
		PageSize: u.PageSize,
		
		Total: u.Total,
		
		PageCount: u.PageCount,
		
		FacetInfo: u.FacetInfo,
		
		AuditMessages: u.AuditMessages,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Auditsearchresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
