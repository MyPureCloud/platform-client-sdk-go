package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Pinconfiguration
type Pinconfiguration struct { 
	// MinimumLength
	MinimumLength *int `json:"minimumLength,omitempty"`


	// MaximumLength
	MaximumLength *int `json:"maximumLength,omitempty"`

}

func (u *Pinconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Pinconfiguration

	

	return json.Marshal(&struct { 
		MinimumLength *int `json:"minimumLength,omitempty"`
		
		MaximumLength *int `json:"maximumLength,omitempty"`
		*Alias
	}{ 
		MinimumLength: u.MinimumLength,
		
		MaximumLength: u.MaximumLength,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Pinconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
