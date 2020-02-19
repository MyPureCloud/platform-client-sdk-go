package platformclientv2
import (
	"encoding/json"
)

// Voicemailssearchresponse
type Voicemailssearchresponse struct { 
	// Total - The total number of results found
	Total *int64 `json:"total,omitempty"`


	// PageCount - The total number of pages
	PageCount *int32 `json:"pageCount,omitempty"`


	// PageSize - The current page size
	PageSize *int32 `json:"pageSize,omitempty"`


	// PageNumber - The current page number
	PageNumber *int32 `json:"pageNumber,omitempty"`


	// PreviousPage - Q64 value for the previous page of results
	PreviousPage *string `json:"previousPage,omitempty"`


	// CurrentPage - Q64 value for the current page of results
	CurrentPage *string `json:"currentPage,omitempty"`


	// NextPage - Q64 value for the next page of results
	NextPage *string `json:"nextPage,omitempty"`


	// Types - Resource types the search was performed against
	Types *[]string `json:"types,omitempty"`


	// Results - Search results
	Results *[]Voicemailmessage `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Voicemailssearchresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
