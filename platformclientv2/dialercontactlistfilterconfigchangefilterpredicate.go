package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistfilterconfigchangefilterpredicate
type Dialercontactlistfilterconfigchangefilterpredicate struct { 
	// Column - The contact list column
	Column *string `json:"column,omitempty"`


	// ColumnType - Whether a contact column is numeric or alphabetic
	ColumnType *string `json:"columnType,omitempty"`


	// Operator - The comparison operator
	Operator *string `json:"operator,omitempty"`


	// Value - The value the predicate applies to
	Value *string `json:"value,omitempty"`


	// VarRange
	VarRange *Dialercontactlistfilterconfigchangerange `json:"range,omitempty"`


	// Inverted - Whether or not to invert to result of evaluating the predicate
	Inverted *bool `json:"inverted,omitempty"`

}

func (o *Dialercontactlistfilterconfigchangefilterpredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistfilterconfigchangefilterpredicate
	
	return json.Marshal(&struct { 
		Column *string `json:"column,omitempty"`
		
		ColumnType *string `json:"columnType,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarRange *Dialercontactlistfilterconfigchangerange `json:"range,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		*Alias
	}{ 
		Column: o.Column,
		
		ColumnType: o.ColumnType,
		
		Operator: o.Operator,
		
		Value: o.Value,
		
		VarRange: o.VarRange,
		
		Inverted: o.Inverted,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercontactlistfilterconfigchangefilterpredicate) UnmarshalJSON(b []byte) error {
	var DialercontactlistfilterconfigchangefilterpredicateMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistfilterconfigchangefilterpredicateMap)
	if err != nil {
		return err
	}
	
	if Column, ok := DialercontactlistfilterconfigchangefilterpredicateMap["column"].(string); ok {
		o.Column = &Column
	}
    
	if ColumnType, ok := DialercontactlistfilterconfigchangefilterpredicateMap["columnType"].(string); ok {
		o.ColumnType = &ColumnType
	}
    
	if Operator, ok := DialercontactlistfilterconfigchangefilterpredicateMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := DialercontactlistfilterconfigchangefilterpredicateMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if VarRange, ok := DialercontactlistfilterconfigchangefilterpredicateMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	
	if Inverted, ok := DialercontactlistfilterconfigchangefilterpredicateMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangefilterpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
