package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userappconfigurationinfo - Configuration information for the integration
type Userappconfigurationinfo struct { 
	// Current - The current, active configuration for the integration.
	Current *Integrationconfiguration `json:"current,omitempty"`


	// Effective - The effective configuration for the app, containing the integration specific configuration along with overrides specified in the integration type.
	Effective *Effectiveconfiguration `json:"effective,omitempty"`

}

func (o *Userappconfigurationinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userappconfigurationinfo
	
	return json.Marshal(&struct { 
		Current *Integrationconfiguration `json:"current,omitempty"`
		
		Effective *Effectiveconfiguration `json:"effective,omitempty"`
		*Alias
	}{ 
		Current: o.Current,
		
		Effective: o.Effective,
		Alias:    (*Alias)(o),
	})
}

func (o *Userappconfigurationinfo) UnmarshalJSON(b []byte) error {
	var UserappconfigurationinfoMap map[string]interface{}
	err := json.Unmarshal(b, &UserappconfigurationinfoMap)
	if err != nil {
		return err
	}
	
	if Current, ok := UserappconfigurationinfoMap["current"].(map[string]interface{}); ok {
		CurrentString, _ := json.Marshal(Current)
		json.Unmarshal(CurrentString, &o.Current)
	}
	
	if Effective, ok := UserappconfigurationinfoMap["effective"].(map[string]interface{}); ok {
		EffectiveString, _ := json.Marshal(Effective)
		json.Unmarshal(EffectiveString, &o.Effective)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userappconfigurationinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
