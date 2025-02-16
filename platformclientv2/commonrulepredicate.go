package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Commonrulepredicate
type Commonrulepredicate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MetricType - The type of metric being evaluated.
	MetricType *string `json:"metricType,omitempty"`

	// MetricValueType - The type of metric value being evaluated.
	MetricValueType *string `json:"metricValueType,omitempty"`

	// ComparisonOperator - The comparison operator being performed on the metric.
	ComparisonOperator *string `json:"comparisonOperator,omitempty"`

	// Value - The value the metric will be compared to.
	Value *float64 `json:"value,omitempty"`

	// Status - The status of the entity corresponding to the metric.
	Status *string `json:"status,omitempty"`

	// Topic - The operational console topic corresponding to the metric.
	Topic *string `json:"topic,omitempty"`

	// Entity - The entity whose metric is being represented.
	Entity *Commonrulepredicateentity `json:"entity,omitempty"`

	// MediaType - The media type of the conversation the metric describes.
	MediaType *string `json:"mediaType,omitempty"`

	// Metric - The metric being evaluated.
	Metric *string `json:"metric,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Commonrulepredicate) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Commonrulepredicate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Commonrulepredicate
	
	return json.Marshal(&struct { 
		MetricType *string `json:"metricType,omitempty"`
		
		MetricValueType *string `json:"metricValueType,omitempty"`
		
		ComparisonOperator *string `json:"comparisonOperator,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Topic *string `json:"topic,omitempty"`
		
		Entity *Commonrulepredicateentity `json:"entity,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		Alias
	}{ 
		MetricType: o.MetricType,
		
		MetricValueType: o.MetricValueType,
		
		ComparisonOperator: o.ComparisonOperator,
		
		Value: o.Value,
		
		Status: o.Status,
		
		Topic: o.Topic,
		
		Entity: o.Entity,
		
		MediaType: o.MediaType,
		
		Metric: o.Metric,
		Alias:    (Alias)(o),
	})
}

func (o *Commonrulepredicate) UnmarshalJSON(b []byte) error {
	var CommonrulepredicateMap map[string]interface{}
	err := json.Unmarshal(b, &CommonrulepredicateMap)
	if err != nil {
		return err
	}
	
	if MetricType, ok := CommonrulepredicateMap["metricType"].(string); ok {
		o.MetricType = &MetricType
	}
    
	if MetricValueType, ok := CommonrulepredicateMap["metricValueType"].(string); ok {
		o.MetricValueType = &MetricValueType
	}
    
	if ComparisonOperator, ok := CommonrulepredicateMap["comparisonOperator"].(string); ok {
		o.ComparisonOperator = &ComparisonOperator
	}
    
	if Value, ok := CommonrulepredicateMap["value"].(float64); ok {
		o.Value = &Value
	}
    
	if Status, ok := CommonrulepredicateMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Topic, ok := CommonrulepredicateMap["topic"].(string); ok {
		o.Topic = &Topic
	}
    
	if Entity, ok := CommonrulepredicateMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if MediaType, ok := CommonrulepredicateMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Metric, ok := CommonrulepredicateMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Commonrulepredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
