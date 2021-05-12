package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Namedentityannotation
type Namedentityannotation struct { 
	// Name - The name of the annotated named entity.
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Namedentityannotation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
