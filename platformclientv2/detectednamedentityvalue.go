package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Detectednamedentityvalue
type Detectednamedentityvalue struct { 
	// Raw - The raw value of the detected named entity.
	Raw *string `json:"raw,omitempty"`


	// Resolved - The resolved value of the detected named entity.
	Resolved *string `json:"resolved,omitempty"`

}

// String returns a JSON representation of the model
func (o *Detectednamedentityvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
