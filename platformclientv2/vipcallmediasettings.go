package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Vipcallmediasettings
type Vipcallmediasettings struct { 
	// Enabled - Toggle that enables VIP experience for this feature.
	Enabled *bool `json:"enabled,omitempty"`


	// SkipOwnershipTime - Toggle that enables this media type to fallback immediately to the configured VIP Backup.
	SkipOwnershipTime *bool `json:"skipOwnershipTime,omitempty"`

}

func (o *Vipcallmediasettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Vipcallmediasettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		SkipOwnershipTime *bool `json:"skipOwnershipTime,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		SkipOwnershipTime: o.SkipOwnershipTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Vipcallmediasettings) UnmarshalJSON(b []byte) error {
	var VipcallmediasettingsMap map[string]interface{}
	err := json.Unmarshal(b, &VipcallmediasettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := VipcallmediasettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if SkipOwnershipTime, ok := VipcallmediasettingsMap["skipOwnershipTime"].(bool); ok {
		o.SkipOwnershipTime = &SkipOwnershipTime
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Vipcallmediasettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
