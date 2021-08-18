package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludetectionresponse
type Nludetectionresponse struct { 
	// Version - The NLU domain version which performed the detection.
	Version *Nludomainversion `json:"version,omitempty"`


	// Output
	Output *Nludetectionoutput `json:"output,omitempty"`


	// Input
	Input *Nludetectioninput `json:"input,omitempty"`

}

func (u *Nludetectionresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludetectionresponse

	

	return json.Marshal(&struct { 
		Version *Nludomainversion `json:"version,omitempty"`
		
		Output *Nludetectionoutput `json:"output,omitempty"`
		
		Input *Nludetectioninput `json:"input,omitempty"`
		*Alias
	}{ 
		Version: u.Version,
		
		Output: u.Output,
		
		Input: u.Input,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nludetectionresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
