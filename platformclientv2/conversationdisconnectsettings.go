package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationdisconnectsettings
type Conversationdisconnectsettings struct { 
	// Enabled - whether or not conversation disconnect setting is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// VarType - Conversation disconnect type
	VarType *string `json:"type,omitempty"`

}

func (o *Conversationdisconnectsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationdisconnectsettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationdisconnectsettings) UnmarshalJSON(b []byte) error {
	var ConversationdisconnectsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationdisconnectsettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := ConversationdisconnectsettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if VarType, ok := ConversationdisconnectsettingsMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationdisconnectsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
