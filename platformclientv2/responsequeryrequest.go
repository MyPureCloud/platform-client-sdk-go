package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Responsequeryrequest - Used to query for responses
type Responsequeryrequest struct { 
	// QueryPhrase - Query phrase to search response text and name. If not set will match all.
	QueryPhrase *string `json:"queryPhrase,omitempty"`


	// PageSize - The maximum number of hits to return. Default: 25, Maximum: 500.
	PageSize *int `json:"pageSize,omitempty"`


	// Filters - Filter the query results.
	Filters *[]Responsefilter `json:"filters,omitempty"`

}

func (u *Responsequeryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responsequeryrequest

	

	return json.Marshal(&struct { 
		QueryPhrase *string `json:"queryPhrase,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		Filters *[]Responsefilter `json:"filters,omitempty"`
		*Alias
	}{ 
		QueryPhrase: u.QueryPhrase,
		
		PageSize: u.PageSize,
		
		Filters: u.Filters,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Responsequeryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
