package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatsettings
type Webchatsettings struct { 
	// RequireDeployment
	RequireDeployment *bool `json:"requireDeployment,omitempty"`

}

func (o *Webchatsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatsettings
	
	return json.Marshal(&struct { 
		RequireDeployment *bool `json:"requireDeployment,omitempty"`
		*Alias
	}{ 
		RequireDeployment: o.RequireDeployment,
		Alias:    (*Alias)(o),
	})
}

func (o *Webchatsettings) UnmarshalJSON(b []byte) error {
	var WebchatsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatsettingsMap)
	if err != nil {
		return err
	}
	
	if RequireDeployment, ok := WebchatsettingsMap["requireDeployment"].(bool); ok {
		o.RequireDeployment = &RequireDeployment
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
