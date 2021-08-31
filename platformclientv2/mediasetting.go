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

func (o *Mediasetting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediasetting
	
	return json.Marshal(&struct { 
		AlertingTimeoutSeconds *int `json:"alertingTimeoutSeconds,omitempty"`
		
		ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`
		*Alias
	}{ 
		AlertingTimeoutSeconds: o.AlertingTimeoutSeconds,
		
		ServiceLevel: o.ServiceLevel,
		Alias:    (*Alias)(o),
	})
}

func (o *Mediasetting) UnmarshalJSON(b []byte) error {
	var MediasettingMap map[string]interface{}
	err := json.Unmarshal(b, &MediasettingMap)
	if err != nil {
		return err
	}
	
	if AlertingTimeoutSeconds, ok := MediasettingMap["alertingTimeoutSeconds"].(float64); ok {
		AlertingTimeoutSecondsInt := int(AlertingTimeoutSeconds)
		o.AlertingTimeoutSeconds = &AlertingTimeoutSecondsInt
	}
	
	if ServiceLevel, ok := MediasettingMap["serviceLevel"].(map[string]interface{}); ok {
		ServiceLevelString, _ := json.Marshal(ServiceLevel)
		json.Unmarshal(ServiceLevelString, &o.ServiceLevel)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Mediasetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
