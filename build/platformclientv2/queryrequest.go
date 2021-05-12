package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Queryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
