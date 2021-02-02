package platformclientv2
import (
	"encoding/json"
)

// Jsonsearchresponse
type Jsonsearchresponse struct { 
	// Total - The total number of results found
	Total *int `json:"total,omitempty"`


	// PageCount - The total number of pages
	PageCount *int `json:"pageCount,omitempty"`


	// PageSize - The current page size
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - The current page number
	PageNumber *int `json:"pageNumber,omitempty"`


	// Types - Resource types the search was performed against
	Types *[]string `json:"types,omitempty"`


	// Results - Search results
	Results *Arraynode `json:"results,omitempty"`


	// Aggregations
	Aggregations *Arraynode `json:"aggregations,omitempty"`

}

// String returns a JSON representation of the model
func (o *Jsonsearchresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
