package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsqueryaggregation
type Analyticsqueryaggregation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - Optional type, can usually be inferred
	VarType *string `json:"type,omitempty"`

	// Dimension - For use with termFrequency aggregations
	Dimension *string `json:"dimension,omitempty"`

	// Metric - For use with numericRange aggregations
	Metric *string `json:"metric,omitempty"`

	// Size - For use with termFrequency aggregations
	Size *int `json:"size,omitempty"`

	// Ranges - For use with numericRange aggregations
	Ranges *[]Aggregationrange `json:"ranges,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsqueryaggregation) SetField(field string, fieldValue interface{}) {
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

func (o Analyticsqueryaggregation) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Analyticsqueryaggregation
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Dimension *string `json:"dimension,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		Ranges *[]Aggregationrange `json:"ranges,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Dimension: o.Dimension,
		
		Metric: o.Metric,
		
		Size: o.Size,
		
		Ranges: o.Ranges,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticsqueryaggregation) UnmarshalJSON(b []byte) error {
	var AnalyticsqueryaggregationMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsqueryaggregationMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := AnalyticsqueryaggregationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Dimension, ok := AnalyticsqueryaggregationMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
    
	if Metric, ok := AnalyticsqueryaggregationMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if Size, ok := AnalyticsqueryaggregationMap["size"].(float64); ok {
		SizeInt := int(Size)
		o.Size = &SizeInt
	}
	
	if Ranges, ok := AnalyticsqueryaggregationMap["ranges"].([]interface{}); ok {
		RangesString, _ := json.Marshal(Ranges)
		json.Unmarshal(RangesString, &o.Ranges)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsqueryaggregation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
