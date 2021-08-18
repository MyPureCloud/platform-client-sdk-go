package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Callroute
type Callroute struct { 
	// Targets - A list of CallTargets to be called when the CallRoute is executed
	Targets *[]Calltarget `json:"targets,omitempty"`

}

func (u *Callroute) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callroute

	

	return json.Marshal(&struct { 
		Targets *[]Calltarget `json:"targets,omitempty"`
		*Alias
	}{ 
		Targets: u.Targets,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Callroute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
