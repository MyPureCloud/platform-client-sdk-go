package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Numericrange
type Numericrange struct { 
	// Gt - Greater than
	Gt *float32 `json:"gt,omitempty"`


	// Gte - Greater than or equal to
	Gte *float32 `json:"gte,omitempty"`


	// Lt - Less than
	Lt *float32 `json:"lt,omitempty"`


	// Lte - Less than or equal to
	Lte *float32 `json:"lte,omitempty"`

}

func (u *Numericrange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Numericrange

	

	return json.Marshal(&struct { 
		Gt *float32 `json:"gt,omitempty"`
		
		Gte *float32 `json:"gte,omitempty"`
		
		Lt *float32 `json:"lt,omitempty"`
		
		Lte *float32 `json:"lte,omitempty"`
		*Alias
	}{ 
		Gt: u.Gt,
		
		Gte: u.Gte,
		
		Lt: u.Lt,
		
		Lte: u.Lte,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Numericrange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
