package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Nludomainversiontrainingresponse
type Nludomainversiontrainingresponse struct { 
	// Message - A message indicating result of the action.
	Message *string `json:"message,omitempty"`


	// Version
	Version *Nludomainversion `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nludomainversiontrainingresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
