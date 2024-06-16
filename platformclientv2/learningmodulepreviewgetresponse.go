package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulepreviewgetresponse - Learning module preview get response
type Learningmodulepreviewgetresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of learning module
	Name *string `json:"name,omitempty"`

	// Description - The description of learning module
	Description *string `json:"description,omitempty"`

	// CoverArt - The cover art for the learning module
	CoverArt *Learningmodulecoverartresponse `json:"coverArt,omitempty"`

	// EnforceContentOrder - If true, learning module content should be viewed one by one in order
	EnforceContentOrder *bool `json:"enforceContentOrder,omitempty"`

	// ReviewAssessmentResults - Allows to view Assessment results in detail
	ReviewAssessmentResults *Reviewassessmentresults `json:"reviewAssessmentResults,omitempty"`

	// AssessmentForm - The assessment form for learning module
	AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`

	// Assignment - the assignment preview
	Assignment *Learningmodulepreviewgetresponseassignment `json:"assignment,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningmodulepreviewgetresponse) SetField(field string, fieldValue interface{}) {
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

func (o Learningmodulepreviewgetresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Learningmodulepreviewgetresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CoverArt *Learningmodulecoverartresponse `json:"coverArt,omitempty"`
		
		EnforceContentOrder *bool `json:"enforceContentOrder,omitempty"`
		
		ReviewAssessmentResults *Reviewassessmentresults `json:"reviewAssessmentResults,omitempty"`
		
		AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`
		
		Assignment *Learningmodulepreviewgetresponseassignment `json:"assignment,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		CoverArt: o.CoverArt,
		
		EnforceContentOrder: o.EnforceContentOrder,
		
		ReviewAssessmentResults: o.ReviewAssessmentResults,
		
		AssessmentForm: o.AssessmentForm,
		
		Assignment: o.Assignment,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Learningmodulepreviewgetresponse) UnmarshalJSON(b []byte) error {
	var LearningmodulepreviewgetresponseMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulepreviewgetresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LearningmodulepreviewgetresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := LearningmodulepreviewgetresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := LearningmodulepreviewgetresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if CoverArt, ok := LearningmodulepreviewgetresponseMap["coverArt"].(map[string]interface{}); ok {
		CoverArtString, _ := json.Marshal(CoverArt)
		json.Unmarshal(CoverArtString, &o.CoverArt)
	}
	
	if EnforceContentOrder, ok := LearningmodulepreviewgetresponseMap["enforceContentOrder"].(bool); ok {
		o.EnforceContentOrder = &EnforceContentOrder
	}
    
	if ReviewAssessmentResults, ok := LearningmodulepreviewgetresponseMap["reviewAssessmentResults"].(map[string]interface{}); ok {
		ReviewAssessmentResultsString, _ := json.Marshal(ReviewAssessmentResults)
		json.Unmarshal(ReviewAssessmentResultsString, &o.ReviewAssessmentResults)
	}
	
	if AssessmentForm, ok := LearningmodulepreviewgetresponseMap["assessmentForm"].(map[string]interface{}); ok {
		AssessmentFormString, _ := json.Marshal(AssessmentForm)
		json.Unmarshal(AssessmentFormString, &o.AssessmentForm)
	}
	
	if Assignment, ok := LearningmodulepreviewgetresponseMap["assignment"].(map[string]interface{}); ok {
		AssignmentString, _ := json.Marshal(Assignment)
		json.Unmarshal(AssignmentString, &o.Assignment)
	}
	
	if SelfUri, ok := LearningmodulepreviewgetresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulepreviewgetresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
