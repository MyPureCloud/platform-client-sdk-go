package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Integrationconfigurationinfo - Configuration information for the integration
type Integrationconfigurationinfo struct { 
	// Current - The current, active configuration for the integration.
	Current *Integrationconfiguration `json:"current,omitempty"`

}

func (o *Integrationconfigurationinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Integrationconfigurationinfo
	
	return json.Marshal(&struct { 
		Current *Integrationconfiguration `json:"current,omitempty"`
		*Alias
	}{ 
		Current: o.Current,
		Alias:    (*Alias)(o),
	})
}

func (o *Integrationconfigurationinfo) UnmarshalJSON(b []byte) error {
	var IntegrationconfigurationinfoMap map[string]interface{}
	err := json.Unmarshal(b, &IntegrationconfigurationinfoMap)
	if err != nil {
		return err
	}
	
	if Current, ok := IntegrationconfigurationinfoMap["current"].(map[string]interface{}); ok {
		CurrentString, _ := json.Marshal(Current)
		json.Unmarshal(CurrentString, &o.Current)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Integrationconfigurationinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
