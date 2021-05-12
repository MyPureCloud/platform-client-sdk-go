package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Stationsettings - Organization settings for stations
type Stationsettings struct { 
	// FreeSeatingConfiguration - Configuration options for free-seating
	FreeSeatingConfiguration *Freeseatingconfiguration `json:"freeSeatingConfiguration,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stationsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
