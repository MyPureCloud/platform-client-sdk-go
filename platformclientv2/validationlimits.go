package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Validationlimits
type Validationlimits struct { 
	// MinLength
	MinLength *Minlength `json:"minLength,omitempty"`


	// MaxLength
	MaxLength *Maxlength `json:"maxLength,omitempty"`


	// MinItems
	MinItems *Minlength `json:"minItems,omitempty"`


	// MaxItems
	MaxItems *Maxlength `json:"maxItems,omitempty"`


	// Minimum
	Minimum *Minlength `json:"minimum,omitempty"`


	// Maximum
	Maximum *Maxlength `json:"maximum,omitempty"`

}

// String returns a JSON representation of the model
func (o *Validationlimits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
