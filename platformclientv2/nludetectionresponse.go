package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Nludetectionresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
