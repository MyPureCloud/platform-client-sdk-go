package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Suggestsearchrequest
type Suggestsearchrequest struct { 
	// Expand - Provides more details about a specified resource
	Expand *[]string `json:"expand,omitempty"`


	// Types - Resource domain type to search
	Types *[]string `json:"types,omitempty"`


	// Query - Suggest query
	Query *[]Suggestsearchcriteria `json:"query,omitempty"`

}

func (u *Suggestsearchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Suggestsearchrequest

	

	return json.Marshal(&struct { 
		Expand *[]string `json:"expand,omitempty"`
		
		Types *[]string `json:"types,omitempty"`
		
		Query *[]Suggestsearchcriteria `json:"query,omitempty"`
		*Alias
	}{ 
		Expand: u.Expand,
		
		Types: u.Types,
		
		Query: u.Query,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Suggestsearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
