package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userparam
type Userparam struct { 
	// Key
	Key *string `json:"key,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`

}

func (u *Userparam) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userparam

	

	return json.Marshal(&struct { 
		Key *string `json:"key,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Key: u.Key,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userparam) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
