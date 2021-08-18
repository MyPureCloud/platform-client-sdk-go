package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queryrequest
type Queryrequest struct { 
	// QueryPhrase
	QueryPhrase *string `json:"queryPhrase,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// FacetNameRequests
	FacetNameRequests *[]string `json:"facetNameRequests,omitempty"`


	// Sort
	Sort *[]Sortitem `json:"sort,omitempty"`


	// Filters
	Filters *[]Contentfilteritem `json:"filters,omitempty"`


	// AttributeFilters
	AttributeFilters *[]Attributefilteritem `json:"attributeFilters,omitempty"`


	// IncludeShares
	IncludeShares *bool `json:"includeShares,omitempty"`

}

func (u *Queryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queryrequest

	

	return json.Marshal(&struct { 
		QueryPhrase *string `json:"queryPhrase,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		FacetNameRequests *[]string `json:"facetNameRequests,omitempty"`
		
		Sort *[]Sortitem `json:"sort,omitempty"`
		
		Filters *[]Contentfilteritem `json:"filters,omitempty"`
		
		AttributeFilters *[]Attributefilteritem `json:"attributeFilters,omitempty"`
		
		IncludeShares *bool `json:"includeShares,omitempty"`
		*Alias
	}{ 
		QueryPhrase: u.QueryPhrase,
		
		PageNumber: u.PageNumber,
		
		PageSize: u.PageSize,
		
		FacetNameRequests: u.FacetNameRequests,
		
		Sort: u.Sort,
		
		Filters: u.Filters,
		
		AttributeFilters: u.AttributeFilters,
		
		IncludeShares: u.IncludeShares,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
