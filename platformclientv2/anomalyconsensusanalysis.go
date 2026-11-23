package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Anomalyconsensusanalysis
type Anomalyconsensusanalysis struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ConsensusType - The name of the anomaly detection algorithm used for this analysis
	ConsensusType *string `json:"consensusType,omitempty"`

	// Anomalous - True if this method determined the detected value to be anomalous
	Anomalous *bool `json:"anomalous,omitempty"`

	// LowerBound - The lower bound of the normal range for this method. Populated when consensusType is Chronos or ZScore
	LowerBound *float64 `json:"lowerBound,omitempty"`

	// UpperBound - The upper bound of the normal range for this method. Populated when consensusType is Chronos or ZScore
	UpperBound *float64 `json:"upperBound,omitempty"`

	// ExpectedValue - The expected value according to this method. Populated when consensusType is Chronos or ZScore
	ExpectedValue *float64 `json:"expectedValue,omitempty"`

	// Deviation - Percentage value of the deviation from the expected value. Not specified when the expected value is 0
	Deviation *float64 `json:"deviation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Anomalyconsensusanalysis) SetField(field string, fieldValue interface{}) {
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

func (o Anomalyconsensusanalysis) MarshalJSON() ([]byte, error) {
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
	type Alias Anomalyconsensusanalysis
	
	return json.Marshal(&struct { 
		ConsensusType *string `json:"consensusType,omitempty"`
		
		Anomalous *bool `json:"anomalous,omitempty"`
		
		LowerBound *float64 `json:"lowerBound,omitempty"`
		
		UpperBound *float64 `json:"upperBound,omitempty"`
		
		ExpectedValue *float64 `json:"expectedValue,omitempty"`
		
		Deviation *float64 `json:"deviation,omitempty"`
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

func (o *Anomalyconsensusanalysis) UnmarshalJSON(b []byte) error {
	var AnomalyconsensusanalysisMap map[string]interface{}
	err := json.Unmarshal(b, &AnomalyconsensusanalysisMap)
	if err != nil {
		return err
	}
	
	if ConsensusType, ok := AnomalyconsensusanalysisMap["consensusType"].(string); ok {
		o.ConsensusType = &ConsensusType
	}
    
	if Anomalous, ok := AnomalyconsensusanalysisMap["anomalous"].(bool); ok {
		o.Anomalous = &Anomalous
	}
    
	if LowerBound, ok := AnomalyconsensusanalysisMap["lowerBound"].(float64); ok {
		o.LowerBound = &LowerBound
	}
    
	if UpperBound, ok := AnomalyconsensusanalysisMap["upperBound"].(float64); ok {
		o.UpperBound = &UpperBound
	}
    
	if ExpectedValue, ok := AnomalyconsensusanalysisMap["expectedValue"].(float64); ok {
		o.ExpectedValue = &ExpectedValue
	}
    
	if Deviation, ok := AnomalyconsensusanalysisMap["deviation"].(float64); ok {
		o.Deviation = &Deviation
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Anomalyconsensusanalysis) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
