package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Usage
type Usage struct { 
	// Types
	Types *[]Usageitem `json:"types,omitempty"`

}

func (u *Usage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Usage

	

	return json.Marshal(&struct { 
		Types *[]Usageitem `json:"types,omitempty"`
		*Alias
	}{ 
		Types: u.Types,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Usage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
