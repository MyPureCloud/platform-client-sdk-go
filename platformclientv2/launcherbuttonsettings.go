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

func (u *Launcherbuttonsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Launcherbuttonsettings

	

	return json.Marshal(&struct { 
		Visibility *string `json:"visibility,omitempty"`
		*Alias
	}{ 
		Visibility: u.Visibility,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Launcherbuttonsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
