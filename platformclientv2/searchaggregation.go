package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Searchaggregation
type Searchaggregation struct { 
	// Field - The field used for aggregation
	Field *string `json:"field,omitempty"`


	// Name - The name of the aggregation. The response aggregation uses this name.
	Name *string `json:"name,omitempty"`


	// VarType - The type of aggregation to perform
	VarType *string `json:"type,omitempty"`


	// Value - A value to use for aggregation
	Value *string `json:"value,omitempty"`


	// Size - The number aggregations results to return out of the entire result set
	Size *int `json:"size,omitempty"`


	// Order - The order in which aggregation results are sorted
	Order *[]string `json:"order,omitempty"`

}

func (o *Searchaggregation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Searchaggregation
	
	return json.Marshal(&struct { 
		Field *string `json:"field,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		Order *[]string `json:"order,omitempty"`
		*Alias
	}{ 
		Field: o.Field,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		Value: o.Value,
		
		Size: o.Size,
		
		Order: o.Order,
		Alias:    (*Alias)(o),
	})
}

func (o *Searchaggregation) UnmarshalJSON(b []byte) error {
	var SearchaggregationMap map[string]interface{}
	err := json.Unmarshal(b, &SearchaggregationMap)
	if err != nil {
		return err
	}
	
	if Field, ok := SearchaggregationMap["field"].(string); ok {
		o.Field = &Field
	}
    
	if Name, ok := SearchaggregationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := SearchaggregationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Value, ok := SearchaggregationMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if Size, ok := SearchaggregationMap["size"].(float64); ok {
		SizeInt := int(Size)
		o.Size = &SizeInt
	}
	
	if Order, ok := SearchaggregationMap["order"].([]interface{}); ok {
		OrderString, _ := json.Marshal(Order)
		json.Unmarshal(OrderString, &o.Order)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Searchaggregation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
