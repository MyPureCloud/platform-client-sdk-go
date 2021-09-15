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

}

func (o *Analyticsreportingsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsreportingsettings
	
	return json.Marshal(&struct { 
		PiiMaskingEnabled *bool `json:"piiMaskingEnabled,omitempty"`
		*Alias
	}{ 
		PiiMaskingEnabled: o.PiiMaskingEnabled,
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsreportingsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
