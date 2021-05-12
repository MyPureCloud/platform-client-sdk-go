package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Templateparameter
type Templateparameter struct { 
	// Id - Response substitution identifier
	Id *string `json:"id,omitempty"`


	// Value - Response substitution value
	Value *string `json:"value,omitempty"`

}

// String returns a JSON representation of the model
func (o *Templateparameter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
