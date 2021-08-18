package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Listwrapperinterval
type Listwrapperinterval struct { 
	// Values
	Values *[]string `json:"values,omitempty"`

}

func (u *Listwrapperinterval) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Listwrapperinterval

	

	return json.Marshal(&struct { 
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Values: u.Values,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Listwrapperinterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
