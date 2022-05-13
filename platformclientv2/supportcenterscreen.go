package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportcenterscreen
type Supportcenterscreen struct { 
	// VarType - The type of the screen
	VarType *string `json:"type,omitempty"`


	// ModuleSettings - Module settings for the screen
	ModuleSettings *[]Supportcentermodulesetting `json:"moduleSettings,omitempty"`

}

func (o *Supportcenterscreen) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportcenterscreen
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		ModuleSettings *[]Supportcentermodulesetting `json:"moduleSettings,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		ModuleSettings: o.ModuleSettings,
		Alias:    (*Alias)(o),
	})
}

func (o *Supportcenterscreen) UnmarshalJSON(b []byte) error {
	var SupportcenterscreenMap map[string]interface{}
	err := json.Unmarshal(b, &SupportcenterscreenMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := SupportcenterscreenMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if ModuleSettings, ok := SupportcenterscreenMap["moduleSettings"].([]interface{}); ok {
		ModuleSettingsString, _ := json.Marshal(ModuleSettings)
		json.Unmarshal(ModuleSettingsString, &o.ModuleSettings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Supportcenterscreen) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
