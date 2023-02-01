package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Aggregationresultentry
type Aggregationresultentry struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Count
	Count *int `json:"count,omitempty"`

	// Value - For termFrequency aggregations
	Value *string `json:"value,omitempty"`

	// Gte - For numericRange aggregations
	Gte *float32 `json:"gte,omitempty"`

	// Lt - For numericRange aggregations
	Lt *float32 `json:"lt,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Aggregationresultentry) SetField(field string, fieldValue interface{}) {
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

func (o Aggregationresultentry) MarshalJSON() ([]byte, error) {
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
	type Alias Aggregationresultentry
	
	return json.Marshal(&struct { 
		Count *int `json:"count,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Gte *float32 `json:"gte,omitempty"`
		
		Lt *float32 `json:"lt,omitempty"`
		Alias
	}{ 
		Count: o.Count,
		
		Value: o.Value,
		
		Gte: o.Gte,
		
		Lt: o.Lt,
		Alias:    (Alias)(o),
	})
}

func (o *Aggregationresultentry) UnmarshalJSON(b []byte) error {
	var AggregationresultentryMap map[string]interface{}
	err := json.Unmarshal(b, &AggregationresultentryMap)
	if err != nil {
		return err
	}
	
	if Count, ok := AggregationresultentryMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if Value, ok := AggregationresultentryMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if Gte, ok := AggregationresultentryMap["gte"].(float64); ok {
		GteFloat32 := float32(Gte)
		o.Gte = &GteFloat32
	}
    
	if Lt, ok := AggregationresultentryMap["lt"].(float64); ok {
		LtFloat32 := float32(Lt)
		o.Lt = &LtFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Aggregationresultentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
