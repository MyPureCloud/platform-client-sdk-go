package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanpatternresponse
type Workplanpatternresponse struct { 
	// WorkPlans - List of work plans in order of rotation on a weekly basis
	WorkPlans *[]Workplanreference `json:"workPlans,omitempty"`

}

func (u *Workplanpatternresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workplanpatternresponse

	

	return json.Marshal(&struct { 
		WorkPlans *[]Workplanreference `json:"workPlans,omitempty"`
		*Alias
	}{ 
		WorkPlans: u.WorkPlans,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Workplanpatternresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
