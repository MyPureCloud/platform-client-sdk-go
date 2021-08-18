package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Featurestate
type Featurestate struct { 
	// Enabled
	Enabled *bool `json:"enabled,omitempty"`

}

func (u *Featurestate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Featurestate

	

	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		Enabled: u.Enabled,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Featurestate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
