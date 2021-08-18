package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutingstatususerparam
type Userroutingstatususerparam struct { 
	// Key
	Key *string `json:"key,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Userroutingstatususerparam) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userroutingstatususerparam

	

	return json.Marshal(&struct { 
		Key *string `json:"key,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Key: u.Key,
		
		Value: u.Value,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userroutingstatususerparam) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
