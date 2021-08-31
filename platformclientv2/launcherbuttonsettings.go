package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Launcherbuttonsettings - The settings for the launcher button
type Launcherbuttonsettings struct { 
	// Visibility - The visibility settings for the button
	Visibility *string `json:"visibility,omitempty"`

}

func (o *Launcherbuttonsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Launcherbuttonsettings
	
	return json.Marshal(&struct { 
		Visibility *string `json:"visibility,omitempty"`
		*Alias
	}{ 
		Visibility: o.Visibility,
		Alias:    (*Alias)(o),
	})
}

func (o *Launcherbuttonsettings) UnmarshalJSON(b []byte) error {
	var LauncherbuttonsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &LauncherbuttonsettingsMap)
	if err != nil {
		return err
	}
	
	if Visibility, ok := LauncherbuttonsettingsMap["visibility"].(string); ok {
		o.Visibility = &Visibility
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Launcherbuttonsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
