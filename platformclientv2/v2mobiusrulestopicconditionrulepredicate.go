package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusrulestopicconditionrulepredicate
type V2mobiusrulestopicconditionrulepredicate struct { 
	// Entity
	Entity *V2mobiusrulestopicentityproperties `json:"entity,omitempty"`


	// Metric
	Metric *string `json:"metric,omitempty"`


	// MetricType
	MetricType *string `json:"metricType,omitempty"`


	// MetricValueType
	MetricValueType *string `json:"metricValueType,omitempty"`


	// Value
	Value *float32 `json:"value,omitempty"`


	// ComparisonOperator
	ComparisonOperator *string `json:"comparisonOperator,omitempty"`

}

func (o *V2mobiusrulestopicconditionrulepredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2mobiusrulestopicconditionrulepredicate
	
	return json.Marshal(&struct { 
		Entity *V2mobiusrulestopicentityproperties `json:"entity,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		MetricType *string `json:"metricType,omitempty"`
		
		MetricValueType *string `json:"metricValueType,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		
		ComparisonOperator *string `json:"comparisonOperator,omitempty"`
		*Alias
	}{ 
		Entity: o.Entity,
		
		Metric: o.Metric,
		
		MetricType: o.MetricType,
		
		MetricValueType: o.MetricValueType,
		
		Value: o.Value,
		
		ComparisonOperator: o.ComparisonOperator,
		Alias:    (*Alias)(o),
	})
}

func (o *V2mobiusrulestopicconditionrulepredicate) UnmarshalJSON(b []byte) error {
	var V2mobiusrulestopicconditionrulepredicateMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusrulestopicconditionrulepredicateMap)
	if err != nil {
		return err
	}
	
	if Entity, ok := V2mobiusrulestopicconditionrulepredicateMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if Metric, ok := V2mobiusrulestopicconditionrulepredicateMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if MetricType, ok := V2mobiusrulestopicconditionrulepredicateMap["metricType"].(string); ok {
		o.MetricType = &MetricType
	}
    
	if MetricValueType, ok := V2mobiusrulestopicconditionrulepredicateMap["metricValueType"].(string); ok {
		o.MetricValueType = &MetricValueType
	}
    
	if Value, ok := V2mobiusrulestopicconditionrulepredicateMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
    
	if ComparisonOperator, ok := V2mobiusrulestopicconditionrulepredicateMap["comparisonOperator"].(string); ok {
		o.ComparisonOperator = &ComparisonOperator
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusrulestopicconditionrulepredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
