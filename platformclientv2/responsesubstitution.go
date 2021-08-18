package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Responsesubstitution - Contains information about the substitutions associated with a response.
type Responsesubstitution struct { 
	// Id - Response substitution identifier.
	Id *string `json:"id,omitempty"`


	// Description - Response substitution description.
	Description *string `json:"description,omitempty"`


	// DefaultValue - Response substitution default value.
	DefaultValue *string `json:"defaultValue,omitempty"`

}

func (u *Responsesubstitution) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responsesubstitution

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DefaultValue *string `json:"defaultValue,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Description: u.Description,
		
		DefaultValue: u.DefaultValue,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Responsesubstitution) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
