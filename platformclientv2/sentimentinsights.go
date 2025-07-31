package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Sentimentinsights
type Sentimentinsights struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PositiveSentimentReasons - The reasons for positive sentiment found in the conversation
	PositiveSentimentReasons *[]Sentimentinsightentry `json:"positiveSentimentReasons,omitempty"`

	// NegativeSentimentReasons - The reasons for negative sentiment found in the conversation
	NegativeSentimentReasons *[]Sentimentinsightentry `json:"negativeSentimentReasons,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Sentimentinsights) SetField(field string, fieldValue interface{}) {
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

func (o Sentimentinsights) MarshalJSON() ([]byte, error) {
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
	type Alias Sentimentinsights
	
	return json.Marshal(&struct { 
		PositiveSentimentReasons *[]Sentimentinsightentry `json:"positiveSentimentReasons,omitempty"`
		
		NegativeSentimentReasons *[]Sentimentinsightentry `json:"negativeSentimentReasons,omitempty"`
		Alias
	}{ 
		PositiveSentimentReasons: o.PositiveSentimentReasons,
		
		NegativeSentimentReasons: o.NegativeSentimentReasons,
		Alias:    (Alias)(o),
	})
}

func (o *Sentimentinsights) UnmarshalJSON(b []byte) error {
	var SentimentinsightsMap map[string]interface{}
	err := json.Unmarshal(b, &SentimentinsightsMap)
	if err != nil {
		return err
	}
	
	if PositiveSentimentReasons, ok := SentimentinsightsMap["positiveSentimentReasons"].([]interface{}); ok {
		PositiveSentimentReasonsString, _ := json.Marshal(PositiveSentimentReasons)
		json.Unmarshal(PositiveSentimentReasonsString, &o.PositiveSentimentReasons)
	}
	
	if NegativeSentimentReasons, ok := SentimentinsightsMap["negativeSentimentReasons"].([]interface{}); ok {
		NegativeSentimentReasonsString, _ := json.Marshal(NegativeSentimentReasons)
		json.Unmarshal(NegativeSentimentReasonsString, &o.NegativeSentimentReasons)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Sentimentinsights) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
