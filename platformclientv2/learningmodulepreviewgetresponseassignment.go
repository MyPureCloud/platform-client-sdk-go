package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulepreviewgetresponseassignment - Learning module preview get response assignment
type Learningmodulepreviewgetresponseassignment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// State - The Learning Assignment state
	State *string `json:"state,omitempty"`

	// PercentageScore - The user's percentage score for this assignment
	PercentageScore *float32 `json:"percentageScore,omitempty"`

	// AssessmentPercentageScore - The user's percentage score for this assignment's assessment
	AssessmentPercentageScore *float32 `json:"assessmentPercentageScore,omitempty"`

	// IsPassed - True if the assessment was passed
	IsPassed *bool `json:"isPassed,omitempty"`

	// AssessmentCompletionPercentage - The assessment completion percentage of assignment
	AssessmentCompletionPercentage *float32 `json:"assessmentCompletionPercentage,omitempty"`

	// CompletionPercentage - The overall completion percentage of assignment
	CompletionPercentage *float32 `json:"completionPercentage,omitempty"`

	// Steps - List of assignment steps
	Steps *[]Learningmodulepreviewgetresponsestep `json:"steps,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningmodulepreviewgetresponseassignment) SetField(field string, fieldValue interface{}) {
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

func (o Learningmodulepreviewgetresponseassignment) MarshalJSON() ([]byte, error) {
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
	type Alias Learningmodulepreviewgetresponseassignment
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		PercentageScore *float32 `json:"percentageScore,omitempty"`
		
		AssessmentPercentageScore *float32 `json:"assessmentPercentageScore,omitempty"`
		
		IsPassed *bool `json:"isPassed,omitempty"`
		
		AssessmentCompletionPercentage *float32 `json:"assessmentCompletionPercentage,omitempty"`
		
		CompletionPercentage *float32 `json:"completionPercentage,omitempty"`
		
		Steps *[]Learningmodulepreviewgetresponsestep `json:"steps,omitempty"`
		Alias
	}{ 
		State: o.State,
		
		PercentageScore: o.PercentageScore,
		
		AssessmentPercentageScore: o.AssessmentPercentageScore,
		
		IsPassed: o.IsPassed,
		
		AssessmentCompletionPercentage: o.AssessmentCompletionPercentage,
		
		CompletionPercentage: o.CompletionPercentage,
		
		Steps: o.Steps,
		Alias:    (Alias)(o),
	})
}

func (o *Learningmodulepreviewgetresponseassignment) UnmarshalJSON(b []byte) error {
	var LearningmodulepreviewgetresponseassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulepreviewgetresponseassignmentMap)
	if err != nil {
		return err
	}
	
	if State, ok := LearningmodulepreviewgetresponseassignmentMap["state"].(string); ok {
		o.State = &State
	}
    
	if PercentageScore, ok := LearningmodulepreviewgetresponseassignmentMap["percentageScore"].(float64); ok {
		PercentageScoreFloat32 := float32(PercentageScore)
		o.PercentageScore = &PercentageScoreFloat32
	}
	
	if AssessmentPercentageScore, ok := LearningmodulepreviewgetresponseassignmentMap["assessmentPercentageScore"].(float64); ok {
		AssessmentPercentageScoreFloat32 := float32(AssessmentPercentageScore)
		o.AssessmentPercentageScore = &AssessmentPercentageScoreFloat32
	}
	
	if IsPassed, ok := LearningmodulepreviewgetresponseassignmentMap["isPassed"].(bool); ok {
		o.IsPassed = &IsPassed
	}
    
	if AssessmentCompletionPercentage, ok := LearningmodulepreviewgetresponseassignmentMap["assessmentCompletionPercentage"].(float64); ok {
		AssessmentCompletionPercentageFloat32 := float32(AssessmentCompletionPercentage)
		o.AssessmentCompletionPercentage = &AssessmentCompletionPercentageFloat32
	}
	
	if CompletionPercentage, ok := LearningmodulepreviewgetresponseassignmentMap["completionPercentage"].(float64); ok {
		CompletionPercentageFloat32 := float32(CompletionPercentage)
		o.CompletionPercentage = &CompletionPercentageFloat32
	}
	
	if Steps, ok := LearningmodulepreviewgetresponseassignmentMap["steps"].([]interface{}); ok {
		StepsString, _ := json.Marshal(Steps)
		json.Unmarshal(StepsString, &o.Steps)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewgetresponseassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
