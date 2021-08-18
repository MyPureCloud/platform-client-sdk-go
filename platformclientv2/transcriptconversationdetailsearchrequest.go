package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptconversationdetailsearchrequest
type Transcriptconversationdetailsearchrequest struct { 
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


	// Types - Resource domain type to search
	Types *[]string `json:"types,omitempty"`


	// Query - The search criteria
	Query *[]Transcriptconversationdetailsearchcriteria `json:"query,omitempty"`

}

func (u *Transcriptconversationdetailsearchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptconversationdetailsearchrequest

	

	return json.Marshal(&struct { 
		SortOrder *string `json:"sortOrder,omitempty"`
		
		SortBy *string `json:"sortBy,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Sort *[]Searchsort `json:"sort,omitempty"`
		
		Types *[]string `json:"types,omitempty"`
		
		Query *[]Transcriptconversationdetailsearchcriteria `json:"query,omitempty"`
		*Alias
	}{ 
		SortOrder: u.SortOrder,
		
		SortBy: u.SortBy,
		
		PageSize: u.PageSize,
		
		PageNumber: u.PageNumber,
		
		Sort: u.Sort,
		
		Types: u.Types,
		
		Query: u.Query,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Transcriptconversationdetailsearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
