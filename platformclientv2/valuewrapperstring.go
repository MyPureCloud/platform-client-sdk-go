package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Valuewrapperstring - An object to provide context to nullable fields in PATCH requests
type Valuewrapperstring struct { 
	// Value - The value for the associated field
	Value *string `json:"value,omitempty"`

}

func (u *Valuewrapperstring) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Valuewrapperstring

	

	return json.Marshal(&struct { 
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Valuewrapperstring) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
