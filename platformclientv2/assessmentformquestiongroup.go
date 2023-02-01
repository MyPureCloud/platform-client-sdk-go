package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentformquestiongroup
type Assessmentformquestiongroup struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the question group,
	Id *string `json:"id,omitempty"`

	// Name - The question group name
	Name *string `json:"name,omitempty"`

	// VarType - The question group type
	VarType *string `json:"type,omitempty"`

	// DefaultAnswersToHighest
	DefaultAnswersToHighest *bool `json:"defaultAnswersToHighest,omitempty"`

	// DefaultAnswersToNA
	DefaultAnswersToNA *bool `json:"defaultAnswersToNA,omitempty"`

	// NaEnabled
	NaEnabled *bool `json:"naEnabled,omitempty"`

	// Weight
	Weight *float32 `json:"weight,omitempty"`

	// ManualWeight
	ManualWeight *bool `json:"manualWeight,omitempty"`

	// Questions - The list of questions for this question group
	Questions *[]Assessmentformquestion `json:"questions,omitempty"`

	// VisibilityCondition
	VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Assessmentformquestiongroup) SetField(field string, fieldValue interface{}) {
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

func (o Assessmentformquestiongroup) MarshalJSON() ([]byte, error) {
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
	type Alias Assessmentformquestiongroup
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		DefaultAnswersToHighest *bool `json:"defaultAnswersToHighest,omitempty"`
		
		DefaultAnswersToNA *bool `json:"defaultAnswersToNA,omitempty"`
		
		NaEnabled *bool `json:"naEnabled,omitempty"`
		
		Weight *float32 `json:"weight,omitempty"`
		
		ManualWeight *bool `json:"manualWeight,omitempty"`
		
		Questions *[]Assessmentformquestion `json:"questions,omitempty"`
		
		VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		DefaultAnswersToHighest: o.DefaultAnswersToHighest,
		
		DefaultAnswersToNA: o.DefaultAnswersToNA,
		
		NaEnabled: o.NaEnabled,
		
		Weight: o.Weight,
		
		ManualWeight: o.ManualWeight,
		
		Questions: o.Questions,
		
		VisibilityCondition: o.VisibilityCondition,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Assessmentformquestiongroup) UnmarshalJSON(b []byte) error {
	var AssessmentformquestiongroupMap map[string]interface{}
	err := json.Unmarshal(b, &AssessmentformquestiongroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AssessmentformquestiongroupMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AssessmentformquestiongroupMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := AssessmentformquestiongroupMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if DefaultAnswersToHighest, ok := AssessmentformquestiongroupMap["defaultAnswersToHighest"].(bool); ok {
		o.DefaultAnswersToHighest = &DefaultAnswersToHighest
	}
    
	if DefaultAnswersToNA, ok := AssessmentformquestiongroupMap["defaultAnswersToNA"].(bool); ok {
		o.DefaultAnswersToNA = &DefaultAnswersToNA
	}
    
	if NaEnabled, ok := AssessmentformquestiongroupMap["naEnabled"].(bool); ok {
		o.NaEnabled = &NaEnabled
	}
    
	if Weight, ok := AssessmentformquestiongroupMap["weight"].(float64); ok {
		WeightFloat32 := float32(Weight)
		o.Weight = &WeightFloat32
	}
	
	if ManualWeight, ok := AssessmentformquestiongroupMap["manualWeight"].(bool); ok {
		o.ManualWeight = &ManualWeight
	}
    
	if Questions, ok := AssessmentformquestiongroupMap["questions"].([]interface{}); ok {
		QuestionsString, _ := json.Marshal(Questions)
		json.Unmarshal(QuestionsString, &o.Questions)
	}
	
	if VisibilityCondition, ok := AssessmentformquestiongroupMap["visibilityCondition"].(map[string]interface{}); ok {
		VisibilityConditionString, _ := json.Marshal(VisibilityCondition)
		json.Unmarshal(VisibilityConditionString, &o.VisibilityCondition)
	}
	
	if SelfUri, ok := AssessmentformquestiongroupMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Assessmentformquestiongroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
