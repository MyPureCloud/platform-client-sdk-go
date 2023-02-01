package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Calibrationassignment
type Calibrationassignment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Calibrator
	Calibrator *User `json:"calibrator,omitempty"`

	// Evaluators
	Evaluators *[]User `json:"evaluators,omitempty"`

	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`

	// ExpertEvaluator
	ExpertEvaluator *User `json:"expertEvaluator,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Calibrationassignment) SetField(field string, fieldValue interface{}) {
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

func (o Calibrationassignment) MarshalJSON() ([]byte, error) {
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
	type Alias Calibrationassignment
	
	return json.Marshal(&struct { 
		Calibrator *User `json:"calibrator,omitempty"`
		
		Evaluators *[]User `json:"evaluators,omitempty"`
		
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		ExpertEvaluator *User `json:"expertEvaluator,omitempty"`
		Alias
	}{ 
		Calibrator: o.Calibrator,
		
		Evaluators: o.Evaluators,
		
		EvaluationForm: o.EvaluationForm,
		
		ExpertEvaluator: o.ExpertEvaluator,
		Alias:    (Alias)(o),
	})
}

func (o *Calibrationassignment) UnmarshalJSON(b []byte) error {
	var CalibrationassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &CalibrationassignmentMap)
	if err != nil {
		return err
	}
	
	if Calibrator, ok := CalibrationassignmentMap["calibrator"].(map[string]interface{}); ok {
		CalibratorString, _ := json.Marshal(Calibrator)
		json.Unmarshal(CalibratorString, &o.Calibrator)
	}
	
	if Evaluators, ok := CalibrationassignmentMap["evaluators"].([]interface{}); ok {
		EvaluatorsString, _ := json.Marshal(Evaluators)
		json.Unmarshal(EvaluatorsString, &o.Evaluators)
	}
	
	if EvaluationForm, ok := CalibrationassignmentMap["evaluationForm"].(map[string]interface{}); ok {
		EvaluationFormString, _ := json.Marshal(EvaluationForm)
		json.Unmarshal(EvaluationFormString, &o.EvaluationForm)
	}
	
	if ExpertEvaluator, ok := CalibrationassignmentMap["expertEvaluator"].(map[string]interface{}); ok {
		ExpertEvaluatorString, _ := json.Marshal(ExpertEvaluator)
		json.Unmarshal(ExpertEvaluatorString, &o.ExpertEvaluator)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Calibrationassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
