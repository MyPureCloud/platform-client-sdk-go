package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicphonenumbercolumn
type Conversationeventtopicphonenumbercolumn struct { 
	// ColumnName
	ColumnName *string `json:"columnName,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Conversationeventtopicphonenumbercolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicphonenumbercolumn
	
	return json.Marshal(&struct { 
		ColumnName *string `json:"columnName,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		ColumnName: o.ColumnName,
		
		VarType: o.VarType,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationeventtopicphonenumbercolumn) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicphonenumbercolumnMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicphonenumbercolumnMap)
	if err != nil {
		return err
	}
	
	if ColumnName, ok := ConversationeventtopicphonenumbercolumnMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
	
	if VarType, ok := ConversationeventtopicphonenumbercolumnMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if AdditionalProperties, ok := ConversationeventtopicphonenumbercolumnMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicphonenumbercolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
