package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Jsonsearchresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Jsonsearchresponse

	

	return json.Marshal(&struct { 
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Types *[]string `json:"types,omitempty"`
		
		Results *Arraynode `json:"results,omitempty"`
		
		Aggregations *Arraynode `json:"aggregations,omitempty"`
		*Alias
	}{ 
		Total: u.Total,
		
		PageCount: u.PageCount,
		
		PageSize: u.PageSize,
		
		PageNumber: u.PageNumber,
		
		Types: u.Types,
		
		Results: u.Results,
		
		Aggregations: u.Aggregations,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Jsonsearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
