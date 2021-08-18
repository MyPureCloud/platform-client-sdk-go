package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Detecteddialogact
type Detecteddialogact struct { 
	// Name - The name of the detected dialog act.
	Name *string `json:"name,omitempty"`


	// Probability - The probability of the detected dialog act.
	Probability *float64 `json:"probability,omitempty"`

}

func (u *Detecteddialogact) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Detecteddialogact

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Probability *float64 `json:"probability,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Probability: u.Probability,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Detecteddialogact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
