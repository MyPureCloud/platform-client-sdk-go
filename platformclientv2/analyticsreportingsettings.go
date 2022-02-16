package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsreportingsettings
type Analyticsreportingsettings struct { 
	// PiiMaskingEnabled - Indication of whether or not personal data is masked in data export and the Analytics/Reporting UI
	PiiMaskingEnabled *bool `json:"piiMaskingEnabled,omitempty"`


	// QueueAgentAccessObfuscation - Indication of whether or not to obfuscate export data from the Queue Agent Details view based on User ACL
	QueueAgentAccessObfuscation *bool `json:"queueAgentAccessObfuscation,omitempty"`

}

func (o *Analyticsreportingsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsreportingsettings
	
	return json.Marshal(&struct { 
		PiiMaskingEnabled *bool `json:"piiMaskingEnabled,omitempty"`
		
		QueueAgentAccessObfuscation *bool `json:"queueAgentAccessObfuscation,omitempty"`
		*Alias
	}{ 
		PiiMaskingEnabled: o.PiiMaskingEnabled,
		
		QueueAgentAccessObfuscation: o.QueueAgentAccessObfuscation,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsreportingsettings) UnmarshalJSON(b []byte) error {
	var AnalyticsreportingsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsreportingsettingsMap)
	if err != nil {
		return err
	}
	
	if PiiMaskingEnabled, ok := AnalyticsreportingsettingsMap["piiMaskingEnabled"].(bool); ok {
		o.PiiMaskingEnabled = &PiiMaskingEnabled
	}
	
	if QueueAgentAccessObfuscation, ok := AnalyticsreportingsettingsMap["queueAgentAccessObfuscation"].(bool); ok {
		o.QueueAgentAccessObfuscation = &QueueAgentAccessObfuscation
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsreportingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
