package platformclientv2
import (
	"encoding/json"
)

// Documentationsearchrequest
type Documentationsearchrequest struct { 
	// SortOrder - The sort order for results
	SortOrder *string `json:"sortOrder,omitempty"`


	// SortBy - The field in the resource that you want to sort the results by
	SortBy *string `json:"sortBy,omitempty"`


	// PageSize - The number of results per page
	PageSize *int32 `json:"pageSize,omitempty"`


	// PageNumber - The page of resources you want to retrieve
	PageNumber *int32 `json:"pageNumber,omitempty"`


	// Sort - Multi-value sort order, list of multiple sort values
	Sort *[]Searchsort `json:"sort,omitempty"`


	// Query
	Query *[]Documentationsearchcriteria `json:"query,omitempty"`

}

// String returns a JSON representation of the model
func (o *Documentationsearchrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
