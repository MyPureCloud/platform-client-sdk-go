package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Itemvalidationlimits
type Itemvalidationlimits struct { 
	// MinLength - A structure denoting the system-imposed minimum string length (for text-based core types) or numeric values (for number-based) core types.  For example, the validationLimits for a text-based core type specify the min/max values for a minimum string length (minLength) constraint supplied by a schemaauthor on a text field.  Similarly, the maxLength's min/max specifies maximum string length constraint supplied by a schema author for the same field.
	MinLength *Minlength `json:"minLength,omitempty"`


	// MaxLength - A structure denoting the system-imposed minimum and maximum string length (for text-based core types) or numeric values (for number-based) core types.  For example, the validationLimits for a text-based core type specify the min/max values for a minimum string length (minLength) constraint supplied by a schemaauthor on a text field.  Similarly, the maxLength's min/max specifies maximum string length constraint supplied by a schema author for the same field.
	MaxLength *Maxlength `json:"maxLength,omitempty"`

}

func (u *Itemvalidationlimits) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Itemvalidationlimits

	

	return json.Marshal(&struct { 
		MinLength *Minlength `json:"minLength,omitempty"`
		
		MaxLength *Maxlength `json:"maxLength,omitempty"`
		*Alias
	}{ 
		MinLength: u.MinLength,
		
		MaxLength: u.MaxLength,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Itemvalidationlimits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
