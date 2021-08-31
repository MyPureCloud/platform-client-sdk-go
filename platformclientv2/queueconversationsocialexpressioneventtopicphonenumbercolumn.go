package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicphonenumbercolumn
type Queueconversationsocialexpressioneventtopicphonenumbercolumn struct { 
	// ColumnName
	ColumnName *string `json:"columnName,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationsocialexpressioneventtopicphonenumbercolumn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicphonenumbercolumn
	
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

func (o *Queueconversationsocialexpressioneventtopicphonenumbercolumn) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicphonenumbercolumnMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicphonenumbercolumnMap)
	if err != nil {
		return err
	}
	
	if ColumnName, ok := QueueconversationsocialexpressioneventtopicphonenumbercolumnMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
	
	if VarType, ok := QueueconversationsocialexpressioneventtopicphonenumbercolumnMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if AdditionalProperties, ok := QueueconversationsocialexpressioneventtopicphonenumbercolumnMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicphonenumbercolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
