package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentaggregatequeryresponsestats
type Learningassignmentaggregatequeryresponsestats struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Count - The count for this metric
	Count *int `json:"count,omitempty"`

	// Min - The minimum value in this metric
	Min *float32 `json:"min,omitempty"`

	// Max - The maximum value in this metric
	Max *float32 `json:"max,omitempty"`

	// Sum - The total of the values for this metric
	Sum *float32 `json:"sum,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningassignmentaggregatequeryresponsestats) SetField(field string, fieldValue interface{}) {
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

func (o Learningassignmentaggregatequeryresponsestats) MarshalJSON() ([]byte, error) {
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
	type Alias Learningassignmentaggregatequeryresponsestats
	
	return json.Marshal(&struct { 
		Count *int `json:"count,omitempty"`
		
		Min *float32 `json:"min,omitempty"`
		
		Max *float32 `json:"max,omitempty"`
		
		Sum *float32 `json:"sum,omitempty"`
		Alias
	}{ 
		Count: o.Count,
		
		Min: o.Min,
		
		Max: o.Max,
		
		Sum: o.Sum,
		Alias:    (Alias)(o),
	})
}

func (o *Learningassignmentaggregatequeryresponsestats) UnmarshalJSON(b []byte) error {
	var LearningassignmentaggregatequeryresponsestatsMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentaggregatequeryresponsestatsMap)
	if err != nil {
		return err
	}
	
	if Count, ok := LearningassignmentaggregatequeryresponsestatsMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if Min, ok := LearningassignmentaggregatequeryresponsestatsMap["min"].(float64); ok {
		MinFloat32 := float32(Min)
		o.Min = &MinFloat32
	}
    
	if Max, ok := LearningassignmentaggregatequeryresponsestatsMap["max"].(float64); ok {
		MaxFloat32 := float32(Max)
		o.Max = &MaxFloat32
	}
    
	if Sum, ok := LearningassignmentaggregatequeryresponsestatsMap["sum"].(float64); ok {
		SumFloat32 := float32(Sum)
		o.Sum = &SumFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryresponsestats) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
