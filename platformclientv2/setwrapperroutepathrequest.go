package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Setwrapperroutepathrequest
type Setwrapperroutepathrequest struct { 
	// Values
	Values *[]Routepathrequest `json:"values,omitempty"`

}

func (u *Setwrapperroutepathrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Setwrapperroutepathrequest

	

	return json.Marshal(&struct { 
		Values *[]Routepathrequest `json:"values,omitempty"`
		*Alias
	}{ 
		Values: u.Values,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Setwrapperroutepathrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
