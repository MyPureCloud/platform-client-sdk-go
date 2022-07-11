package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediasetting
type Mediasetting struct { 
	// EnableAutoAnswer - Indicates if auto-answer is enabled for the given media type or subtype (default is false).  Subtype settings take precedence over media type settings.
	EnableAutoAnswer *bool `json:"enableAutoAnswer,omitempty"`


	// AlertingTimeoutSeconds
	AlertingTimeoutSeconds *int `json:"alertingTimeoutSeconds,omitempty"`


	// ServiceLevel
	ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`


	// SubTypeSettings - Map of media subtype to media subtype specific settings.
	SubTypeSettings *map[string]Basemediasettings `json:"subTypeSettings,omitempty"`

}

func (o *Mediasetting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediasetting
	
	return json.Marshal(&struct { 
		EnableAutoAnswer *bool `json:"enableAutoAnswer,omitempty"`
		
		AlertingTimeoutSeconds *int `json:"alertingTimeoutSeconds,omitempty"`
		
		ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`
		
		SubTypeSettings *map[string]Basemediasettings `json:"subTypeSettings,omitempty"`
		*Alias
	}{ 
		EnableAutoAnswer: o.EnableAutoAnswer,
		
		AlertingTimeoutSeconds: o.AlertingTimeoutSeconds,
		
		ServiceLevel: o.ServiceLevel,
		
		SubTypeSettings: o.SubTypeSettings,
		Alias:    (*Alias)(o),
	})
}

func (o *Mediasetting) UnmarshalJSON(b []byte) error {
	var MediasettingMap map[string]interface{}
	err := json.Unmarshal(b, &MediasettingMap)
	if err != nil {
		return err
	}
	
	if EnableAutoAnswer, ok := MediasettingMap["enableAutoAnswer"].(bool); ok {
		o.EnableAutoAnswer = &EnableAutoAnswer
	}
    
	if AlertingTimeoutSeconds, ok := MediasettingMap["alertingTimeoutSeconds"].(float64); ok {
		AlertingTimeoutSecondsInt := int(AlertingTimeoutSeconds)
		o.AlertingTimeoutSeconds = &AlertingTimeoutSecondsInt
	}
	
	if ServiceLevel, ok := MediasettingMap["serviceLevel"].(map[string]interface{}); ok {
		ServiceLevelString, _ := json.Marshal(ServiceLevel)
		json.Unmarshal(ServiceLevelString, &o.ServiceLevel)
	}
	
	if SubTypeSettings, ok := MediasettingMap["subTypeSettings"].(map[string]interface{}); ok {
		SubTypeSettingsString, _ := json.Marshal(SubTypeSettings)
		json.Unmarshal(SubTypeSettingsString, &o.SubTypeSettings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Mediasetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
