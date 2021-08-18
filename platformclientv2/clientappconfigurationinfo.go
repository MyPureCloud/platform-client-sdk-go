package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Clientappconfigurationinfo - Configuration information for the integration
type Clientappconfigurationinfo struct { 
	// Current - The current, active configuration for the integration.
	Current *Integrationconfiguration `json:"current,omitempty"`


	// Effective - The effective configuration for the app, containing the integration specific configuration along with overrides specified in the integration type.
	Effective *Effectiveconfiguration `json:"effective,omitempty"`

}

func (u *Clientappconfigurationinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Clientappconfigurationinfo

	

	return json.Marshal(&struct { 
		Current *Integrationconfiguration `json:"current,omitempty"`
		
		Effective *Effectiveconfiguration `json:"effective,omitempty"`
		*Alias
	}{ 
		Current: u.Current,
		
		Effective: u.Effective,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Clientappconfigurationinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
