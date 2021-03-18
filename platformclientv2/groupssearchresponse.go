package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Groupssearchresponse
type Groupssearchresponse struct { 
	// Total - The total number of results found
	Total *int `json:"total,omitempty"`


	// PageCount - The total number of pages
	PageCount *int `json:"pageCount,omitempty"`


	// PageSize - The current page size
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - The current page number
	PageNumber *int `json:"pageNumber,omitempty"`


	// PreviousPage - Q64 value for the previous page of results
	PreviousPage *string `json:"previousPage,omitempty"`


	// CurrentPage - Q64 value for the current page of results
	CurrentPage *string `json:"currentPage,omitempty"`


	// NextPage - Q64 value for the next page of results
	NextPage *string `json:"nextPage,omitempty"`


	// Types - Resource types the search was performed against
	Types *[]string `json:"types,omitempty"`


	// Results - Search results
	Results *[]Group `json:"results,omitempty"`

}

// String returns a JSON representation of the model
func (o *Groupssearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
