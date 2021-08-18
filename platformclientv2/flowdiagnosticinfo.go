package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowdiagnosticinfo
type Flowdiagnosticinfo struct { 
	// LastActionId - The step number of the survey invite flow where the error occurred.
	LastActionId *int `json:"lastActionId,omitempty"`

}

func (u *Flowdiagnosticinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowdiagnosticinfo

	

	return json.Marshal(&struct { 
		LastActionId *int `json:"lastActionId,omitempty"`
		*Alias
	}{ 
		LastActionId: u.LastActionId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Flowdiagnosticinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
