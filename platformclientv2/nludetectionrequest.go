package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludetectionrequest
type Nludetectionrequest struct { 
	// Input - The input subject to NLU detection.
	Input *Nludetectioninput `json:"input,omitempty"`


	// Context - The context for the input to NLU detection.
	Context *Nludetectioncontext `json:"context,omitempty"`

}

func (u *Nludetectionrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludetectionrequest

	

	return json.Marshal(&struct { 
		Input *Nludetectioninput `json:"input,omitempty"`
		
		Context *Nludetectioncontext `json:"context,omitempty"`
		*Alias
	}{ 
		Input: u.Input,
		
		Context: u.Context,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nludetectionrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
