package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Valuewrapperstring - An object to provide context to nullable fields in PATCH requests
type Valuewrapperstring struct { 
	// Value - The value for the associated field
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Valuewrapperstring) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
