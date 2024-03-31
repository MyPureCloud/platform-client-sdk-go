package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusalertstopicconditionrulepredicate
type V2mobiusalertstopicconditionrulepredicate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

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

	// Status
	Status *string `json:"status,omitempty"`

	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// ComparisonOperator
	ComparisonOperator *string `json:"comparisonOperator,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2mobiusalertstopicconditionrulepredicate) SetField(field string, fieldValue interface{}) {
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

func (o V2mobiusalertstopicconditionrulepredicate) MarshalJSON() ([]byte, error) {
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
	type Alias V2mobiusalertstopicconditionrulepredicate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Entity *V2mobiusalertstopicentityproperties `json:"entity,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		MetricType *string `json:"metricType,omitempty"`
		
		MetricValueType *string `json:"metricValueType,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		ComparisonOperator *string `json:"comparisonOperator,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Entity: o.Entity,
		
		Metric: o.Metric,
		
		MetricType: o.MetricType,
		
		MetricValueType: o.MetricValueType,
		
		Value: o.Value,
		
		Status: o.Status,
		
		MediaType: o.MediaType,
		
		ComparisonOperator: o.ComparisonOperator,
		Alias:    (Alias)(o),
	})
}

func (o *V2mobiusalertstopicconditionrulepredicate) UnmarshalJSON(b []byte) error {
	var V2mobiusalertstopicconditionrulepredicateMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusalertstopicconditionrulepredicateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2mobiusalertstopicconditionrulepredicateMap["id"].(map[string]interface{}); ok {
		IdString, _ := json.Marshal(Id)
		json.Unmarshal(IdString, &o.Id)
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
    
	if Status, ok := V2mobiusalertstopicconditionrulepredicateMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if MediaType, ok := V2mobiusalertstopicconditionrulepredicateMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
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
