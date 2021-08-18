package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Valuewrapperplanningperiodsettings - An object to provide context to nullable fields in PATCH requests
type Valuewrapperplanningperiodsettings struct { 
	// Value - The value for the associated field
	Value *Planningperiodsettings `json:"value,omitempty"`

}

func (u *Valuewrapperplanningperiodsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Valuewrapperplanningperiodsettings

	

	return json.Marshal(&struct { 
		Value *Planningperiodsettings `json:"value,omitempty"`
		*Alias
	}{ 
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Valuewrapperplanningperiodsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
