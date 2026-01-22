package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationsearchaggregationresponse
type Evaluationsearchaggregationresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Value - Simple aggregation value (for SUM, COUNT, AVERAGE, etc.)
	Value *float64 `json:"value,omitempty"`

	// Count - Count of documents
	Count *int `json:"count,omitempty"`

	// Minimum - Minimum value
	Minimum *float64 `json:"minimum,omitempty"`

	// Maximum - Maximum value
	Maximum *float64 `json:"maximum,omitempty"`

	// Average - Average value
	Average *float64 `json:"average,omitempty"`

	// Sum - Total Sum
	Sum *float64 `json:"sum,omitempty"`

	// DocumentCountErrorUpperBound - Upper bound estimate of the error in document count for this aggregation
	DocumentCountErrorUpperBound *int `json:"documentCountErrorUpperBound,omitempty"`

	// SumOtherDocumentCount - Total document count for buckets not included in the response due to size limits
	SumOtherDocumentCount *int `json:"sumOtherDocumentCount,omitempty"`

	// Buckets - List of aggregation buckets
	Buckets *[]Evaluationsearchaggregationbucket `json:"buckets,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Evaluationsearchaggregationresponse) SetField(field string, fieldValue interface{}) {
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

func (o Evaluationsearchaggregationresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Evaluationsearchaggregationresponse
	
	return json.Marshal(&struct { 
		Value *float64 `json:"value,omitempty"`
		
		Count *int `json:"count,omitempty"`
		
		Minimum *float64 `json:"minimum,omitempty"`
		
		Maximum *float64 `json:"maximum,omitempty"`
		
		Average *float64 `json:"average,omitempty"`
		
		Sum *float64 `json:"sum,omitempty"`
		
		DocumentCountErrorUpperBound *int `json:"documentCountErrorUpperBound,omitempty"`
		
		SumOtherDocumentCount *int `json:"sumOtherDocumentCount,omitempty"`
		
		Buckets *[]Evaluationsearchaggregationbucket `json:"buckets,omitempty"`
		Alias
	}{ 
		Value: o.Value,
		
		Count: o.Count,
		
		Minimum: o.Minimum,
		
		Maximum: o.Maximum,
		
		Average: o.Average,
		
		Sum: o.Sum,
		
		DocumentCountErrorUpperBound: o.DocumentCountErrorUpperBound,
		
		SumOtherDocumentCount: o.SumOtherDocumentCount,
		
		Buckets: o.Buckets,
		Alias:    (Alias)(o),
	})
}

func (o *Evaluationsearchaggregationresponse) UnmarshalJSON(b []byte) error {
	var EvaluationsearchaggregationresponseMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationsearchaggregationresponseMap)
	if err != nil {
		return err
	}
	
	if Value, ok := EvaluationsearchaggregationresponseMap["value"].(float64); ok {
		o.Value = &Value
	}
    
	if Count, ok := EvaluationsearchaggregationresponseMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if Minimum, ok := EvaluationsearchaggregationresponseMap["minimum"].(float64); ok {
		o.Minimum = &Minimum
	}
    
	if Maximum, ok := EvaluationsearchaggregationresponseMap["maximum"].(float64); ok {
		o.Maximum = &Maximum
	}
    
	if Average, ok := EvaluationsearchaggregationresponseMap["average"].(float64); ok {
		o.Average = &Average
	}
    
	if Sum, ok := EvaluationsearchaggregationresponseMap["sum"].(float64); ok {
		o.Sum = &Sum
	}
    
	if DocumentCountErrorUpperBound, ok := EvaluationsearchaggregationresponseMap["documentCountErrorUpperBound"].(float64); ok {
		DocumentCountErrorUpperBoundInt := int(DocumentCountErrorUpperBound)
		o.DocumentCountErrorUpperBound = &DocumentCountErrorUpperBoundInt
	}
	
	if SumOtherDocumentCount, ok := EvaluationsearchaggregationresponseMap["sumOtherDocumentCount"].(float64); ok {
		SumOtherDocumentCountInt := int(SumOtherDocumentCount)
		o.SumOtherDocumentCount = &SumOtherDocumentCountInt
	}
	
	if Buckets, ok := EvaluationsearchaggregationresponseMap["buckets"].([]interface{}); ok {
		BucketsString, _ := json.Marshal(Buckets)
		json.Unmarshal(BucketsString, &o.Buckets)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationsearchaggregationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
