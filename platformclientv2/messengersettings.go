package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messengersettings - Settings concerning messenger
type Messengersettings struct { 
	// Enabled - Whether or not messenger is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// Styles - The style settings for messenger
	Styles *Messengerstyles `json:"styles,omitempty"`


	// LauncherButton - The launcher button settings for messenger
	LauncherButton *Launcherbuttonsettings `json:"launcherButton,omitempty"`


	// FileUpload - The file upload settings for messenger
	FileUpload *Fileuploadsettings `json:"fileUpload,omitempty"`

}

func (u *Messengersettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messengersettings

	

	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		Styles *Messengerstyles `json:"styles,omitempty"`
		
		LauncherButton *Launcherbuttonsettings `json:"launcherButton,omitempty"`
		
		FileUpload *Fileuploadsettings `json:"fileUpload,omitempty"`
		*Alias
	}{ 
		Enabled: u.Enabled,
		
		Styles: u.Styles,
		
		LauncherButton: u.LauncherButton,
		
		FileUpload: u.FileUpload,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Messengersettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
