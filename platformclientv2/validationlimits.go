package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Validationlimits) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Validationlimits

	

	return json.Marshal(&struct { 
		MinLength *Minlength `json:"minLength,omitempty"`
		
		MaxLength *Maxlength `json:"maxLength,omitempty"`
		
		MinItems *Minlength `json:"minItems,omitempty"`
		
		MaxItems *Maxlength `json:"maxItems,omitempty"`
		
		Minimum *Minlength `json:"minimum,omitempty"`
		
		Maximum *Maxlength `json:"maximum,omitempty"`
		*Alias
	}{ 
		MinLength: u.MinLength,
		
		MaxLength: u.MaxLength,
		
		MinItems: u.MinItems,
		
		MaxItems: u.MaxItems,
		
		Minimum: u.Minimum,
		
		Maximum: u.Maximum,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Validationlimits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
