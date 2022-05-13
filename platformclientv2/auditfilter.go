package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditfilter
type Auditfilter struct { 
	// Name - The name of the field by which to filter.
	Name *string `json:"name,omitempty"`


	// VarType - The type of the filter, DATE or STRING.
	VarType *string `json:"type,omitempty"`


	// Operator - The operation that the filter performs.
	Operator *string `json:"operator,omitempty"`


	// Values - The values to make the filter comparison against.
	Values *[]string `json:"values,omitempty"`

}

func (o *Auditfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditfilter
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		VarType: o.VarType,
		
		Operator: o.Operator,
		
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditfilter) UnmarshalJSON(b []byte) error {
	var AuditfilterMap map[string]interface{}
	err := json.Unmarshal(b, &AuditfilterMap)
	if err != nil {
		return err
	}
	
	if Name, ok := AuditfilterMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := AuditfilterMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Operator, ok := AuditfilterMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Values, ok := AuditfilterMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
