package platformclientv2
import (
	"encoding/json"
)

// Transcriptsearchrequest
type Transcriptsearchrequest struct { 
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


	// ReturnFields
	ReturnFields *[]string `json:"returnFields,omitempty"`


	// Types - Resource domain type to search
	Types *[]string `json:"types,omitempty"`


	// Query - The search criteria
	Query *[]Transcriptsearchcriteria `json:"query,omitempty"`

}

// String returns a JSON representation of the model
func (o *Transcriptsearchrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
