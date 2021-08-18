package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Namedentitytypedefinition
type Namedentitytypedefinition struct { 
	// Name - The name of the entity type.
	Name *string `json:"name,omitempty"`


	// Description - Description of the of the named entity type.
	Description *string `json:"description,omitempty"`


	// Mechanism - The mechanism enabling detection of the named entity type.
	Mechanism *Namedentitytypemechanism `json:"mechanism,omitempty"`

}

func (u *Namedentitytypedefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Namedentitytypedefinition

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Mechanism *Namedentitytypemechanism `json:"mechanism,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Description: u.Description,
		
		Mechanism: u.Mechanism,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Namedentitytypedefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
