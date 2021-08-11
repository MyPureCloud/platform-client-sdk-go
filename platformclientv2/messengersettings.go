package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Messengersettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
