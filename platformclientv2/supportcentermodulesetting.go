package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportcentermodulesetting
type Supportcentermodulesetting struct { 
	// VarType - Screen module type
	VarType *string `json:"type,omitempty"`


	// Enabled - Whether or not support center screen module is enabled
	Enabled *bool `json:"enabled,omitempty"`

}

func (o *Supportcentermodulesetting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportcentermodulesetting
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Enabled: o.Enabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Supportcentermodulesetting) UnmarshalJSON(b []byte) error {
	var SupportcentermodulesettingMap map[string]interface{}
	err := json.Unmarshal(b, &SupportcentermodulesettingMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := SupportcentermodulesettingMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Enabled, ok := SupportcentermodulesettingMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Supportcentermodulesetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
