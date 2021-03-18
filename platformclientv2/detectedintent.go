package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Detectedintent
type Detectedintent struct { 
	// Name - The name of the detected intent.
	Name *string `json:"name,omitempty"`


	// Probability - The probability of the detected intent.
	Probability *float64 `json:"probability,omitempty"`


	// Entities - The collection of named entities detected.
	Entities *[]Detectednamedentity `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Detectedintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
