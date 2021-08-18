package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgesearchresponse
type Knowledgesearchresponse struct { 
	// SearchId - Search Id
	SearchId *string `json:"searchId,omitempty"`


	// Total - Total number of records returned
	Total *int `json:"total,omitempty"`


	// PageCount - Number of pages returned in the result calculated according to the pageSize and the total
	PageCount *int `json:"pageCount,omitempty"`


	// PageSize - Number of records according to the page size
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - Current page number for the returned records
	PageNumber *int `json:"pageNumber,omitempty"`


	// Results - Results associated to the search response
	Results *[]Knowledgesearchdocument `json:"results,omitempty"`

}

func (u *Knowledgesearchresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgesearchresponse

	

	return json.Marshal(&struct { 
		SearchId *string `json:"searchId,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Results *[]Knowledgesearchdocument `json:"results,omitempty"`
		*Alias
	}{ 
		SearchId: u.SearchId,
		
		Total: u.Total,
		
		PageCount: u.PageCount,
		
		PageSize: u.PageSize,
		
		PageNumber: u.PageNumber,
		
		Results: u.Results,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Knowledgesearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
