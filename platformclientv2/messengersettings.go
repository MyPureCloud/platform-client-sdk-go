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


	// Position - The position settings for messenger
	Position *Messengerpositionsettings `json:"position,omitempty"`

}

func (o *Messengersettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messengersettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		Styles *Messengerstyles `json:"styles,omitempty"`
		
		LauncherButton *Launcherbuttonsettings `json:"launcherButton,omitempty"`
		
		FileUpload *Fileuploadsettings `json:"fileUpload,omitempty"`
		
		Position *Messengerpositionsettings `json:"position,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		Styles: o.Styles,
		
		LauncherButton: o.LauncherButton,
		
		FileUpload: o.FileUpload,
		
		Position: o.Position,
		Alias:    (*Alias)(o),
	})
}

func (o *Messengersettings) UnmarshalJSON(b []byte) error {
	var MessengersettingsMap map[string]interface{}
	err := json.Unmarshal(b, &MessengersettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := MessengersettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if Styles, ok := MessengersettingsMap["styles"].(map[string]interface{}); ok {
		StylesString, _ := json.Marshal(Styles)
		json.Unmarshal(StylesString, &o.Styles)
	}
	
	if LauncherButton, ok := MessengersettingsMap["launcherButton"].(map[string]interface{}); ok {
		LauncherButtonString, _ := json.Marshal(LauncherButton)
		json.Unmarshal(LauncherButtonString, &o.LauncherButton)
	}
	
	if FileUpload, ok := MessengersettingsMap["fileUpload"].(map[string]interface{}); ok {
		FileUploadString, _ := json.Marshal(FileUpload)
		json.Unmarshal(FileUploadString, &o.FileUpload)
	}
	
	if Position, ok := MessengersettingsMap["position"].(map[string]interface{}); ok {
		PositionString, _ := json.Marshal(Position)
		json.Unmarshal(PositionString, &o.Position)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messengersettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
