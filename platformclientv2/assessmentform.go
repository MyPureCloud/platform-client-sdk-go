package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentform
type Assessmentform struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// DateModified - Last modified date of the assessment form. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// ContextId - The unique Id for all versions of this assessment form
	ContextId *string `json:"contextId,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// Published - If true, assessment form is published
	Published *bool `json:"published,omitempty"`

	// PassPercent - The pass percent for the assessment form
	PassPercent *int `json:"passPercent,omitempty"`

	// QuestionGroups - A list of question groups
	QuestionGroups *[]Assessmentformquestiongroup `json:"questionGroups,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Assessmentform) SetField(field string, fieldValue interface{}) {
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

func (o Assessmentform) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateModified", }
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
	type Alias Assessmentform
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		PassPercent *int `json:"passPercent,omitempty"`
		
		QuestionGroups *[]Assessmentformquestiongroup `json:"questionGroups,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		DateModified: DateModified,
		
		ContextId: o.ContextId,
		
		SelfUri: o.SelfUri,
		
		Published: o.Published,
		
		PassPercent: o.PassPercent,
		
		QuestionGroups: o.QuestionGroups,
		Alias:    (Alias)(o),
	})
}

func (o *Assessmentform) UnmarshalJSON(b []byte) error {
	var AssessmentformMap map[string]interface{}
	err := json.Unmarshal(b, &AssessmentformMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AssessmentformMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if dateModifiedString, ok := AssessmentformMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ContextId, ok := AssessmentformMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if SelfUri, ok := AssessmentformMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if Published, ok := AssessmentformMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if PassPercent, ok := AssessmentformMap["passPercent"].(float64); ok {
		PassPercentInt := int(PassPercent)
		o.PassPercent = &PassPercentInt
	}
	
	if QuestionGroups, ok := AssessmentformMap["questionGroups"].([]interface{}); ok {
		QuestionGroupsString, _ := json.Marshal(QuestionGroups)
		json.Unmarshal(QuestionGroupsString, &o.QuestionGroups)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Assessmentform) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
