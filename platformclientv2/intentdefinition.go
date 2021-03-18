package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Intentdefinition
type Intentdefinition struct { 
	// Name - The name of the intent.
	Name *string `json:"name,omitempty"`


	// EntityTypeBindings - The bindings for the named entity types used in this intent.
	EntityTypeBindings *[]Namedentitytypebinding `json:"entityTypeBindings,omitempty"`


	// Utterances - The utterances that act as training phrases for the intent.
	Utterances *[]Nluutterance `json:"utterances,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intentdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
