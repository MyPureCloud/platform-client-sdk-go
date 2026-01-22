package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationsearchaggregationbucket
type Evaluationsearchaggregationbucket struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Key - The key for this bucket
	Key *string `json:"key,omitempty"`

	// KeyAsString - The key as a string representation
	KeyAsString *string `json:"keyAsString,omitempty"`

	// DocumentCount - Number of documents in this bucket
	DocumentCount *int `json:"documentCount,omitempty"`

	// KeyValue - Numeric key value for date histograms
	KeyValue *int `json:"keyValue,omitempty"`

	// From - Lower bound for range buckets
	From *float64 `json:"from,omitempty"`

	// To - Upper bound for range buckets
	To *float64 `json:"to,omitempty"`

	// Value - Simple aggregation value
	Value *float64 `json:"value,omitempty"`

	// Count - Count of documents
	Count *int `json:"count,omitempty"`

	// Minimum - Minimum value in the aggregation
	Minimum *float64 `json:"minimum,omitempty"`

	// Maximum - Maximum value in the aggregation
	Maximum *float64 `json:"maximum,omitempty"`

	// Average - Average value in the aggregation
	Average *float64 `json:"average,omitempty"`

	// Sum - Sum of values in the aggregation
	Sum *float64 `json:"sum,omitempty"`

	// SubAggregations - Nested sub-aggregations
	SubAggregations *map[string]Evaluationsearchaggregationresponse `json:"subAggregations,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluationsearchaggregationbucket) SetField(field string, fieldValue interface{}) {
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

func (o Evaluationsearchaggregationbucket) MarshalJSON() ([]byte, error) {
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
	type Alias Evaluationsearchaggregationbucket
	
	return json.Marshal(&struct { 
		Key *string `json:"key,omitempty"`
		
		KeyAsString *string `json:"keyAsString,omitempty"`
		
		DocumentCount *int `json:"documentCount,omitempty"`
		
		KeyValue *int `json:"keyValue,omitempty"`
		
		From *float64 `json:"from,omitempty"`
		
		To *float64 `json:"to,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		
		Count *int `json:"count,omitempty"`
		
		Minimum *float64 `json:"minimum,omitempty"`
		
		Maximum *float64 `json:"maximum,omitempty"`
		
		Average *float64 `json:"average,omitempty"`
		
		Sum *float64 `json:"sum,omitempty"`
		
		SubAggregations *map[string]Evaluationsearchaggregationresponse `json:"subAggregations,omitempty"`
		Alias
	}{ 
		Key: o.Key,
		
		KeyAsString: o.KeyAsString,
		
		DocumentCount: o.DocumentCount,
		
		KeyValue: o.KeyValue,
		
		From: o.From,
		
		To: o.To,
		
		Value: o.Value,
		
		Count: o.Count,
		
		Minimum: o.Minimum,
		
		Maximum: o.Maximum,
		
		Average: o.Average,
		
		Sum: o.Sum,
		
		SubAggregations: o.SubAggregations,
		Alias:    (Alias)(o),
	})
}

func (o *Evaluationsearchaggregationbucket) UnmarshalJSON(b []byte) error {
	var EvaluationsearchaggregationbucketMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationsearchaggregationbucketMap)
	if err != nil {
		return err
	}
	
	if Key, ok := EvaluationsearchaggregationbucketMap["key"].(string); ok {
		o.Key = &Key
	}
    
	if KeyAsString, ok := EvaluationsearchaggregationbucketMap["keyAsString"].(string); ok {
		o.KeyAsString = &KeyAsString
	}
    
	if DocumentCount, ok := EvaluationsearchaggregationbucketMap["documentCount"].(float64); ok {
		DocumentCountInt := int(DocumentCount)
		o.DocumentCount = &DocumentCountInt
	}
	
	if KeyValue, ok := EvaluationsearchaggregationbucketMap["keyValue"].(float64); ok {
		KeyValueInt := int(KeyValue)
		o.KeyValue = &KeyValueInt
	}
	
	if From, ok := EvaluationsearchaggregationbucketMap["from"].(float64); ok {
		o.From = &From
	}
    
	if To, ok := EvaluationsearchaggregationbucketMap["to"].(float64); ok {
		o.To = &To
	}
    
	if Value, ok := EvaluationsearchaggregationbucketMap["value"].(float64); ok {
		o.Value = &Value
	}
    
	if Count, ok := EvaluationsearchaggregationbucketMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if Minimum, ok := EvaluationsearchaggregationbucketMap["minimum"].(float64); ok {
		o.Minimum = &Minimum
	}
    
	if Maximum, ok := EvaluationsearchaggregationbucketMap["maximum"].(float64); ok {
		o.Maximum = &Maximum
	}
    
	if Average, ok := EvaluationsearchaggregationbucketMap["average"].(float64); ok {
		o.Average = &Average
	}
    
	if Sum, ok := EvaluationsearchaggregationbucketMap["sum"].(float64); ok {
		o.Sum = &Sum
	}
    
	if SubAggregations, ok := EvaluationsearchaggregationbucketMap["subAggregations"].(map[string]interface{}); ok {
		SubAggregationsString, _ := json.Marshal(SubAggregations)
		json.Unmarshal(SubAggregationsString, &o.SubAggregations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationsearchaggregationbucket) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
