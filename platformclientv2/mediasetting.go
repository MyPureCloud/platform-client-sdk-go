package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediasetting
type Mediasetting struct { 
	// AlertingTimeoutSeconds
	AlertingTimeoutSeconds *int `json:"alertingTimeoutSeconds,omitempty"`


	// ServiceLevel
	ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`

}

func (u *Mediasetting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediasetting

	

	return json.Marshal(&struct { 
		AlertingTimeoutSeconds *int `json:"alertingTimeoutSeconds,omitempty"`
		
		ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`
		*Alias
	}{ 
		AlertingTimeoutSeconds: u.AlertingTimeoutSeconds,
		
		ServiceLevel: u.ServiceLevel,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Mediasetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
