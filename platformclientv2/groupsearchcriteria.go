package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Groupsearchcriteria
type Groupsearchcriteria struct { 
	// EndValue - The end value of the range. This field is used for range search types.
	EndValue *string `json:"endValue,omitempty"`


	// Values - A list of values for the search to match against
	Values *[]string `json:"values,omitempty"`


	// StartValue - The start value of the range. This field is used for range search types.
	StartValue *string `json:"startValue,omitempty"`


	// Fields - Field names to search against
	Fields *[]string `json:"fields,omitempty"`


	// Value - A value for the search to match against
	Value *string `json:"value,omitempty"`


	// Operator - How to apply this search criteria against other criteria
	Operator *string `json:"operator,omitempty"`


	// Group - Groups multiple conditions
	Group *[]Groupsearchcriteria `json:"group,omitempty"`


	// DateFormat - Set date format for criteria values when using date range search type.  Supports Java date format syntax, example yyyy-MM-dd'T'HH:mm:ss.SSSX.
	DateFormat *string `json:"dateFormat,omitempty"`


	// VarType - Search Type
	VarType *string `json:"type,omitempty"`

}

func (o *Groupsearchcriteria) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Groupsearchcriteria
	
	return json.Marshal(&struct { 
		EndValue *string `json:"endValue,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		StartValue *string `json:"startValue,omitempty"`
		
		Fields *[]string `json:"fields,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Group *[]Groupsearchcriteria `json:"group,omitempty"`
		
		DateFormat *string `json:"dateFormat,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		EndValue: o.EndValue,
		
		Values: o.Values,
		
		StartValue: o.StartValue,
		
		Fields: o.Fields,
		
		Value: o.Value,
		
		Operator: o.Operator,
		
		Group: o.Group,
		
		DateFormat: o.DateFormat,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Groupsearchcriteria) UnmarshalJSON(b []byte) error {
	var GroupsearchcriteriaMap map[string]interface{}
	err := json.Unmarshal(b, &GroupsearchcriteriaMap)
	if err != nil {
		return err
	}
	
	if EndValue, ok := GroupsearchcriteriaMap["endValue"].(string); ok {
		o.EndValue = &EndValue
	}
	
	if Values, ok := GroupsearchcriteriaMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if StartValue, ok := GroupsearchcriteriaMap["startValue"].(string); ok {
		o.StartValue = &StartValue
	}
	
	if Fields, ok := GroupsearchcriteriaMap["fields"].([]interface{}); ok {
		FieldsString, _ := json.Marshal(Fields)
		json.Unmarshal(FieldsString, &o.Fields)
	}
	
	if Value, ok := GroupsearchcriteriaMap["value"].(string); ok {
		o.Value = &Value
	}
	
	if Operator, ok := GroupsearchcriteriaMap["operator"].(string); ok {
		o.Operator = &Operator
	}
	
	if Group, ok := GroupsearchcriteriaMap["group"].([]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if DateFormat, ok := GroupsearchcriteriaMap["dateFormat"].(string); ok {
		o.DateFormat = &DateFormat
	}
	
	if VarType, ok := GroupsearchcriteriaMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Groupsearchcriteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
