package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Aggregationrange
type Aggregationrange struct { 
	// Gte - Greater than or equal to
	Gte *float32 `json:"gte,omitempty"`


	// Lt - Less than
	Lt *float32 `json:"lt,omitempty"`

}

func (u *Aggregationrange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Aggregationrange

	

	return json.Marshal(&struct { 
		Gte *float32 `json:"gte,omitempty"`
		
		Lt *float32 `json:"lt,omitempty"`
		*Alias
	}{ 
		Gte: u.Gte,
		
		Lt: u.Lt,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Aggregationrange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
