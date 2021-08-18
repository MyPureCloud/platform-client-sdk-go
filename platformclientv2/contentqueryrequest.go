package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentqueryrequest
type Contentqueryrequest struct { 
	// QueryPhrase
	QueryPhrase *string `json:"queryPhrase,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// FacetNameRequests
	FacetNameRequests *[]string `json:"facetNameRequests,omitempty"`


	// Sort
	Sort *[]Contentsortitem `json:"sort,omitempty"`


	// Filters
	Filters *[]Contentfacetfilteritem `json:"filters,omitempty"`


	// AttributeFilters
	AttributeFilters *[]Contentattributefilteritem `json:"attributeFilters,omitempty"`


	// IncludeShares
	IncludeShares *bool `json:"includeShares,omitempty"`

}

func (u *Contentqueryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentqueryrequest

	

	return json.Marshal(&struct { 
		QueryPhrase *string `json:"queryPhrase,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		FacetNameRequests *[]string `json:"facetNameRequests,omitempty"`
		
		Sort *[]Contentsortitem `json:"sort,omitempty"`
		
		Filters *[]Contentfacetfilteritem `json:"filters,omitempty"`
		
		AttributeFilters *[]Contentattributefilteritem `json:"attributeFilters,omitempty"`
		
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
func (o *Contentqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
