package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatsettings
type Webchatsettings struct { 
	// RequireDeployment
	RequireDeployment *bool `json:"requireDeployment,omitempty"`

}

func (u *Webchatsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatsettings

	

	return json.Marshal(&struct { 
		RequireDeployment *bool `json:"requireDeployment,omitempty"`
		*Alias
	}{ 
		RequireDeployment: u.RequireDeployment,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Webchatsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
