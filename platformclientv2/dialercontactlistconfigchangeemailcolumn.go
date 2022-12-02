package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistconfigchangeemailcolumn
type Dialercontactlistconfigchangeemailcolumn struct { 
	// ColumnName - The name of the email address column
	ColumnName *string `json:"columnName,omitempty"`


	// VarType - The type of the email address column, for example, 'work' or 'home'
	VarType *string `json:"type,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

func (o *Dialercontactlistconfigchangeemailcolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistconfigchangeemailcolumn
	
	return json.Marshal(&struct { 
		ColumnName *string `json:"columnName,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		ColumnName: o.ColumnName,
		
		VarType: o.VarType,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercontactlistconfigchangeemailcolumn) UnmarshalJSON(b []byte) error {
	var DialercontactlistconfigchangeemailcolumnMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistconfigchangeemailcolumnMap)
	if err != nil {
		return err
	}
	
	if ColumnName, ok := DialercontactlistconfigchangeemailcolumnMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
    
	if VarType, ok := DialercontactlistconfigchangeemailcolumnMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if AdditionalProperties, ok := DialercontactlistconfigchangeemailcolumnMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistconfigchangeemailcolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
