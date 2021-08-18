package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Templateparameter
type Templateparameter struct { 
	// Id - Response substitution identifier
	Id *string `json:"id,omitempty"`


	// Value - Response substitution value
	Value *string `json:"value,omitempty"`

}

func (u *Templateparameter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Templateparameter

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Value *string `json:"value,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Value: u.Value,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Templateparameter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
