package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Externalsegment
type Externalsegment struct { 
	// Id - Identifier for the external segment in the system where it originates from.
	Id *string `json:"id,omitempty"`


	// Name - Name for the external segment in the system where it originates from.
	Name *string `json:"name,omitempty"`


	// Source - The external system where the segment originates from.
	Source *string `json:"source,omitempty"`

}

// String returns a JSON representation of the model
func (o *Externalsegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
