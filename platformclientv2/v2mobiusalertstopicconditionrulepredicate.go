package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusalertstopicconditionrulepredicate
type V2mobiusalertstopicconditionrulepredicate struct { 
	// Entity
	Entity *V2mobiusalertstopicentityproperties `json:"entity,omitempty"`


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

func (o *V2mobiusalertstopicconditionrulepredicate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2mobiusalertstopicconditionrulepredicate
	
	return json.Marshal(&struct { 
		Entity *V2mobiusalertstopicentityproperties `json:"entity,omitempty"`
		
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

func (o *V2mobiusalertstopicconditionrulepredicate) UnmarshalJSON(b []byte) error {
	var V2mobiusalertstopicconditionrulepredicateMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusalertstopicconditionrulepredicateMap)
	if err != nil {
		return err
	}
	
	if Entity, ok := V2mobiusalertstopicconditionrulepredicateMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if Metric, ok := V2mobiusalertstopicconditionrulepredicateMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if MetricType, ok := V2mobiusalertstopicconditionrulepredicateMap["metricType"].(string); ok {
		o.MetricType = &MetricType
	}
    
	if MetricValueType, ok := V2mobiusalertstopicconditionrulepredicateMap["metricValueType"].(string); ok {
		o.MetricValueType = &MetricValueType
	}
    
	if Value, ok := V2mobiusalertstopicconditionrulepredicateMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
    
	if ComparisonOperator, ok := V2mobiusalertstopicconditionrulepredicateMap["comparisonOperator"].(string); ok {
		o.ComparisonOperator = &ComparisonOperator
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusalertstopicconditionrulepredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
