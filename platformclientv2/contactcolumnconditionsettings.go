package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactcolumnconditionsettings
type Contactcolumnconditionsettings struct { 
	// ColumnName - The name of the contact list column to evaluate.
	ColumnName *string `json:"columnName,omitempty"`


	// Operator - The operator to use when comparing values.
	Operator *string `json:"operator,omitempty"`


	// Value - The value to compare against the contact's data.
	Value *string `json:"value,omitempty"`


	// ValueType - The data type the value should be treated as.
	ValueType *string `json:"valueType,omitempty"`

}

func (o *Contactcolumnconditionsettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactcolumnconditionsettings
	
	return json.Marshal(&struct { 
		ColumnName *string `json:"columnName,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		ValueType *string `json:"valueType,omitempty"`
		*Alias
	}{ 
		ColumnName: o.ColumnName,
		
		Operator: o.Operator,
		
		Value: o.Value,
		
		ValueType: o.ValueType,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactcolumnconditionsettings) UnmarshalJSON(b []byte) error {
	var ContactcolumnconditionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &ContactcolumnconditionsettingsMap)
	if err != nil {
		return err
	}
	
	if ColumnName, ok := ContactcolumnconditionsettingsMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
    
	if Operator, ok := ContactcolumnconditionsettingsMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := ContactcolumnconditionsettingsMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if ValueType, ok := ContactcolumnconditionsettingsMap["valueType"].(string); ok {
		o.ValueType = &ValueType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactcolumnconditionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
