package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Pinconfiguration
type Pinconfiguration struct { 
	// MinimumLength
	MinimumLength *int `json:"minimumLength,omitempty"`


	// MaximumLength
	MaximumLength *int `json:"maximumLength,omitempty"`

}

// String returns a JSON representation of the model
func (o *Pinconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
