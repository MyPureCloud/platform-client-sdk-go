package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulepreviewupdaterequest - Learning module preview update request
type Learningmodulepreviewupdaterequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// State - The assignment State
	State *string `json:"state,omitempty"`

	// CurrentStep - The assignment current step
	CurrentStep *Learningmodulepreviewupdaterequestcurrentstep `json:"currentStep,omitempty"`

	// Steps - The assignment Steps
	Steps *[]Learningmodulepreviewupdatestep `json:"steps,omitempty"`

	// Assessment - The assessment for learning module
	Assessment *Learningassessment `json:"assessment,omitempty"`

	// AssessmentForm - The assessment form for learning module
	AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningmodulepreviewupdaterequest) SetField(field string, fieldValue interface{}) {
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

func (o Learningmodulepreviewupdaterequest) MarshalJSON() ([]byte, error) {
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
	type Alias Learningmodulepreviewupdaterequest
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		CurrentStep *Learningmodulepreviewupdaterequestcurrentstep `json:"currentStep,omitempty"`
		
		Steps *[]Learningmodulepreviewupdatestep `json:"steps,omitempty"`
		
		Assessment *Learningassessment `json:"assessment,omitempty"`
		
		AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`
		Alias
	}{ 
		State: o.State,
		
		CurrentStep: o.CurrentStep,
		
		Steps: o.Steps,
		
		Assessment: o.Assessment,
		
		AssessmentForm: o.AssessmentForm,
		Alias:    (Alias)(o),
	})
}

func (o *Learningmodulepreviewupdaterequest) UnmarshalJSON(b []byte) error {
	var LearningmodulepreviewupdaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulepreviewupdaterequestMap)
	if err != nil {
		return err
	}
	
	if State, ok := LearningmodulepreviewupdaterequestMap["state"].(string); ok {
		o.State = &State
	}
    
	if CurrentStep, ok := LearningmodulepreviewupdaterequestMap["currentStep"].(map[string]interface{}); ok {
		CurrentStepString, _ := json.Marshal(CurrentStep)
		json.Unmarshal(CurrentStepString, &o.CurrentStep)
	}
	
	if Steps, ok := LearningmodulepreviewupdaterequestMap["steps"].([]interface{}); ok {
		StepsString, _ := json.Marshal(Steps)
		json.Unmarshal(StepsString, &o.Steps)
	}
	
	if Assessment, ok := LearningmodulepreviewupdaterequestMap["assessment"].(map[string]interface{}); ok {
		AssessmentString, _ := json.Marshal(Assessment)
		json.Unmarshal(AssessmentString, &o.Assessment)
	}
	
	if AssessmentForm, ok := LearningmodulepreviewupdaterequestMap["assessmentForm"].(map[string]interface{}); ok {
		AssessmentFormString, _ := json.Marshal(AssessmentForm)
		json.Unmarshal(AssessmentFormString, &o.AssessmentForm)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
