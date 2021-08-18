package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Activation
type Activation struct { 
	// VarType - Type of activation.
	VarType *string `json:"type,omitempty"`


	// DelayInSeconds - Activation delay time amount.
	DelayInSeconds *int `json:"delayInSeconds,omitempty"`

}

func (u *Activation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Activation

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		DelayInSeconds *int `json:"delayInSeconds,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		DelayInSeconds: u.DelayInSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Activation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
