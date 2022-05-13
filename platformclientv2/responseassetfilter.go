package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Responseassetfilter
type Responseassetfilter struct { 
	// EndValue - The end value of the range. This field is used for range search types. Accepts numbers and date in ISO8601 format
	EndValue *string `json:"endValue,omitempty"`


	// Values - A list of values for the search to match against
	Values *[]string `json:"values,omitempty"`


	// StartValue - The start value of the range. This field is used for range search types. Accepts numbers and date in ISO8601 format
	StartValue *string `json:"startValue,omitempty"`


	// Fields - Field name to search against. Allowed Values: divisionId, name, contentLength, contentType, dateCreated
	Fields *[]string `json:"fields,omitempty"`


	// Value - A value for the search to match against
	Value *string `json:"value,omitempty"`


	// VarType - How to apply this search criteria against other criteria
	VarType *string `json:"type,omitempty"`

}

func (o *Responseassetfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responseassetfilter
	
	return json.Marshal(&struct { 
		EndValue *string `json:"endValue,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		StartValue *string `json:"startValue,omitempty"`
		
		Fields *[]string `json:"fields,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		EndValue: o.EndValue,
		
		Values: o.Values,
		
		StartValue: o.StartValue,
		
		Fields: o.Fields,
		
		Value: o.Value,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Responseassetfilter) UnmarshalJSON(b []byte) error {
	var ResponseassetfilterMap map[string]interface{}
	err := json.Unmarshal(b, &ResponseassetfilterMap)
	if err != nil {
		return err
	}
	
	if EndValue, ok := ResponseassetfilterMap["endValue"].(string); ok {
		o.EndValue = &EndValue
	}
    
	if Values, ok := ResponseassetfilterMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if StartValue, ok := ResponseassetfilterMap["startValue"].(string); ok {
		o.StartValue = &StartValue
	}
    
	if Fields, ok := ResponseassetfilterMap["fields"].([]interface{}); ok {
		FieldsString, _ := json.Marshal(Fields)
		json.Unmarshal(FieldsString, &o.Fields)
	}
	
	if Value, ok := ResponseassetfilterMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if VarType, ok := ResponseassetfilterMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Responseassetfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
