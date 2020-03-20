package platformclientv2
import (
	"encoding/json"
)

// Contentqueryrequest
type Contentqueryrequest struct { 
	// QueryPhrase
	QueryPhrase *string `json:"queryPhrase,omitempty"`


	// PageNumber
	PageNumber *int32 `json:"pageNumber,omitempty"`


	// PageSize
	PageSize *int32 `json:"pageSize,omitempty"`


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

// String returns a JSON representation of the model
func (o *Contentqueryrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
