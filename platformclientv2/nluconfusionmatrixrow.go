package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluconfusionmatrixrow
type Nluconfusionmatrixrow struct { 
	// Name - The name of the intent for the row.
	Name *string `json:"name,omitempty"`


	// Columns - The columns of confusion matrix for the intent
	Columns *[]Nluconfusionmatrixcolumn `json:"columns,omitempty"`

}

func (u *Nluconfusionmatrixrow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nluconfusionmatrixrow

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Columns *[]Nluconfusionmatrixcolumn `json:"columns,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Columns: u.Columns,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nluconfusionmatrixrow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
