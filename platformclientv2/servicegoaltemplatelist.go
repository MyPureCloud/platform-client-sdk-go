package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Servicegoaltemplatelist - List of service goal templates
type Servicegoaltemplatelist struct { 
	// Entities
	Entities *[]Servicegoaltemplate `json:"entities,omitempty"`

}

func (u *Servicegoaltemplatelist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Servicegoaltemplatelist

	

	return json.Marshal(&struct { 
		Entities *[]Servicegoaltemplate `json:"entities,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Servicegoaltemplatelist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
