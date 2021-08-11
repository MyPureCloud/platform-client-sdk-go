package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Launcherbuttonsettings - The settings for the launcher button
type Launcherbuttonsettings struct { 
	// Visibility - The visibility settings for the button
	Visibility *string `json:"visibility,omitempty"`

}

// String returns a JSON representation of the model
func (o *Launcherbuttonsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
