package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Vipmediasettings
type Vipmediasettings struct { 
	// Enabled - Toggle that enables VIP experience for this feature.
	Enabled *bool `json:"enabled,omitempty"`


	// SkipOwnershipTime - Toggle that enables this media type to fallback immediately to the configured VIP Backup.
	SkipOwnershipTime *bool `json:"skipOwnershipTime,omitempty"`

}

func (o *Vipmediasettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Vipmediasettings
	
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

func (o *Vipmediasettings) UnmarshalJSON(b []byte) error {
	var VipmediasettingsMap map[string]interface{}
	err := json.Unmarshal(b, &VipmediasettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := VipmediasettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if SkipOwnershipTime, ok := VipmediasettingsMap["skipOwnershipTime"].(bool); ok {
		o.SkipOwnershipTime = &SkipOwnershipTime
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Vipmediasettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
