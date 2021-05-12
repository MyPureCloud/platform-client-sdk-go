package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Namedentity
type Namedentity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the object.
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Namedentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
