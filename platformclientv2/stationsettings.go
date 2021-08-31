package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stationsettings - Organization settings for stations
type Stationsettings struct { 
	// FreeSeatingConfiguration - Configuration options for free-seating
	FreeSeatingConfiguration *Freeseatingconfiguration `json:"freeSeatingConfiguration,omitempty"`

}

func (o *Stationsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stationsettings
	
	return json.Marshal(&struct { 
		FreeSeatingConfiguration *Freeseatingconfiguration `json:"freeSeatingConfiguration,omitempty"`
		*Alias
	}{ 
		FreeSeatingConfiguration: o.FreeSeatingConfiguration,
		Alias:    (*Alias)(o),
	})
}

func (o *Stationsettings) UnmarshalJSON(b []byte) error {
	var StationsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &StationsettingsMap)
	if err != nil {
		return err
	}
	
	if FreeSeatingConfiguration, ok := StationsettingsMap["freeSeatingConfiguration"].(map[string]interface{}); ok {
		FreeSeatingConfigurationString, _ := json.Marshal(FreeSeatingConfiguration)
		json.Unmarshal(FreeSeatingConfigurationString, &o.FreeSeatingConfiguration)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stationsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
