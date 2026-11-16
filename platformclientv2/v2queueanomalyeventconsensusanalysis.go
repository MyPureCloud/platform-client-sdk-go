package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2queueanomalyeventconsensusanalysis
type V2queueanomalyeventconsensusanalysis struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ConsensusType
	ConsensusType *string `json:"consensusType,omitempty"`

	// Anomalous
	Anomalous *bool `json:"anomalous,omitempty"`

	// LowerBound
	LowerBound *float32 `json:"lowerBound,omitempty"`

	// UpperBound
	UpperBound *float32 `json:"upperBound,omitempty"`

	// ExpectedValue
	ExpectedValue *float32 `json:"expectedValue,omitempty"`

	// Deviation
	Deviation *float32 `json:"deviation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2queueanomalyeventconsensusanalysis) SetField(field string, fieldValue interface{}) {
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

func (o V2queueanomalyeventconsensusanalysis) MarshalJSON() ([]byte, error) {
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
	type Alias V2queueanomalyeventconsensusanalysis
	
	return json.Marshal(&struct { 
		ConsensusType *string `json:"consensusType,omitempty"`
		
		Anomalous *bool `json:"anomalous,omitempty"`
		
		LowerBound *float32 `json:"lowerBound,omitempty"`
		
		UpperBound *float32 `json:"upperBound,omitempty"`
		
		ExpectedValue *float32 `json:"expectedValue,omitempty"`
		
		Deviation *float32 `json:"deviation,omitempty"`
		Alias
	}{ 
		ConsensusType: o.ConsensusType,
		
		Anomalous: o.Anomalous,
		
		LowerBound: o.LowerBound,
		
		UpperBound: o.UpperBound,
		
		ExpectedValue: o.ExpectedValue,
		
		Deviation: o.Deviation,
		Alias:    (Alias)(o),
	})
}

func (o *V2queueanomalyeventconsensusanalysis) UnmarshalJSON(b []byte) error {
	var V2queueanomalyeventconsensusanalysisMap map[string]interface{}
	err := json.Unmarshal(b, &V2queueanomalyeventconsensusanalysisMap)
	if err != nil {
		return err
	}
	
	if ConsensusType, ok := V2queueanomalyeventconsensusanalysisMap["consensusType"].(string); ok {
		o.ConsensusType = &ConsensusType
	}
    
	if Anomalous, ok := V2queueanomalyeventconsensusanalysisMap["anomalous"].(bool); ok {
		o.Anomalous = &Anomalous
	}
    
	if LowerBound, ok := V2queueanomalyeventconsensusanalysisMap["lowerBound"].(float64); ok {
		LowerBoundFloat32 := float32(LowerBound)
		o.LowerBound = &LowerBoundFloat32
	}
    
	if UpperBound, ok := V2queueanomalyeventconsensusanalysisMap["upperBound"].(float64); ok {
		UpperBoundFloat32 := float32(UpperBound)
		o.UpperBound = &UpperBoundFloat32
	}
    
	if ExpectedValue, ok := V2queueanomalyeventconsensusanalysisMap["expectedValue"].(float64); ok {
		ExpectedValueFloat32 := float32(ExpectedValue)
		o.ExpectedValue = &ExpectedValueFloat32
	}
    
	if Deviation, ok := V2queueanomalyeventconsensusanalysisMap["deviation"].(float64); ok {
		DeviationFloat32 := float32(Deviation)
		o.Deviation = &DeviationFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2queueanomalyeventconsensusanalysis) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
