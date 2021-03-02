package platformclientv2
import (
	"encoding/json"
)

// Searchrequest
type Searchrequest struct { 
	// SortOrder - The sort order for results
	SortOrder *string `json:"sortOrder,omitempty"`


	// SortBy - The field in the resource that you want to sort the results by
	SortBy *string `json:"sortBy,omitempty"`


	// PageSize - The number of results per page
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - The page of resources you want to retrieve
	PageNumber *int `json:"pageNumber,omitempty"`


	// Sort - Multi-value sort order, list of multiple sort values
	Sort *[]Searchsort `json:"sort,omitempty"`


	// ReturnFields - A List of strings.  Possible values are any field in the resource you are searching on.  The other option is to use ALL_FIELDS, when this is provided all fields in the resource will be returned in the search results.
	ReturnFields *[]string `json:"returnFields,omitempty"`


	// Expand - Provides more details about a specified resource
	Expand *[]string `json:"expand,omitempty"`


	// Types - Resource domain type to search
	Types *[]string `json:"types,omitempty"`


	// Query - The search criteria
	Query *[]Searchcriteria `json:"query,omitempty"`


	// Aggregations - Aggregation criteria
	Aggregations *[]Searchaggregation `json:"aggregations,omitempty"`

}

// String returns a JSON representation of the model
func (o *Searchrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
