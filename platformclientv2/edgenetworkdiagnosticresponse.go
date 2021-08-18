package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgenetworkdiagnosticresponse
type Edgenetworkdiagnosticresponse struct { 
	// CommandCorrelationId - UUID of each executed command on edge
	CommandCorrelationId *string `json:"commandCorrelationId,omitempty"`


	// Diagnostics - Response string of executed command from edge
	Diagnostics *string `json:"diagnostics,omitempty"`

}

func (u *Edgenetworkdiagnosticresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgenetworkdiagnosticresponse

	

	return json.Marshal(&struct { 
		CommandCorrelationId *string `json:"commandCorrelationId,omitempty"`
		
		Diagnostics *string `json:"diagnostics,omitempty"`
		*Alias
	}{ 
		CommandCorrelationId: u.CommandCorrelationId,
		
		Diagnostics: u.Diagnostics,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgenetworkdiagnosticresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
