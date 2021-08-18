package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Tagqueryrequest
type Tagqueryrequest struct { 
	// Query
	Query *string `json:"query,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`

}

func (u *Tagqueryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Tagqueryrequest

	

	return json.Marshal(&struct { 
		Query *string `json:"query,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		*Alias
	}{ 
		Query: u.Query,
		
		PageNumber: u.PageNumber,
		
		PageSize: u.PageSize,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Tagqueryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
