package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassessment
type Learningassessment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AssessmentId - The Id of the assessment
	AssessmentId *string `json:"assessmentId,omitempty"`

	// ContextId - The context Id of the related assessment form
	ContextId *string `json:"contextId,omitempty"`

	// AssessmentFormId - The Id of the related assessment form
	AssessmentFormId *string `json:"assessmentFormId,omitempty"`

	// Status - Status of the assessment
	Status *string `json:"status,omitempty"`

	// Answers - Answers for the assessment
	Answers *Assessmentscoringset `json:"answers,omitempty"`

	// DateCreated - Date the assessment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Date the assessment was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// DateSubmitted - Date the assessment was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateSubmitted *time.Time `json:"dateSubmitted,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningassessment) SetField(field string, fieldValue interface{}) {
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

func (o Learningassessment) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified","DateSubmitted", }
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
	type Alias Learningassessment
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateSubmitted := new(string)
	if o.DateSubmitted != nil {
		
		*DateSubmitted = timeutil.Strftime(o.DateSubmitted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateSubmitted = nil
	}
	
	return json.Marshal(&struct { 
		AssessmentId *string `json:"assessmentId,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		AssessmentFormId *string `json:"assessmentFormId,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Answers *Assessmentscoringset `json:"answers,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateSubmitted *string `json:"dateSubmitted,omitempty"`
		Alias
	}{ 
		AssessmentId: o.AssessmentId,
		
		ContextId: o.ContextId,
		
		AssessmentFormId: o.AssessmentFormId,
		
		Status: o.Status,
		
		Answers: o.Answers,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DateSubmitted: DateSubmitted,
		Alias:    (Alias)(o),
	})
}

func (o *Learningassessment) UnmarshalJSON(b []byte) error {
	var LearningassessmentMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassessmentMap)
	if err != nil {
		return err
	}
	
	if AssessmentId, ok := LearningassessmentMap["assessmentId"].(string); ok {
		o.AssessmentId = &AssessmentId
	}
    
	if ContextId, ok := LearningassessmentMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if AssessmentFormId, ok := LearningassessmentMap["assessmentFormId"].(string); ok {
		o.AssessmentFormId = &AssessmentFormId
	}
    
	if Status, ok := LearningassessmentMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Answers, ok := LearningassessmentMap["answers"].(map[string]interface{}); ok {
		AnswersString, _ := json.Marshal(Answers)
		json.Unmarshal(AnswersString, &o.Answers)
	}
	
	if dateCreatedString, ok := LearningassessmentMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := LearningassessmentMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateSubmittedString, ok := LearningassessmentMap["dateSubmitted"].(string); ok {
		DateSubmitted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateSubmittedString)
		o.DateSubmitted = &DateSubmitted
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassessment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
