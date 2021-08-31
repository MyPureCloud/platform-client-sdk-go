package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Fileuploadsettings - File upload settings for messenger
type Fileuploadsettings struct { 
	// Modes - The list of supported file upload modes
	Modes *[]Fileuploadmode `json:"modes,omitempty"`

}

func (o *Fileuploadsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Fileuploadsettings
	
	return json.Marshal(&struct { 
		Modes *[]Fileuploadmode `json:"modes,omitempty"`
		*Alias
	}{ 
		Modes: o.Modes,
		Alias:    (*Alias)(o),
	})
}

func (o *Fileuploadsettings) UnmarshalJSON(b []byte) error {
	var FileuploadsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &FileuploadsettingsMap)
	if err != nil {
		return err
	}
	
	if Modes, ok := FileuploadsettingsMap["modes"].([]interface{}); ok {
		ModesString, _ := json.Marshal(Modes)
		json.Unmarshal(ModesString, &o.Modes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Fileuploadsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
