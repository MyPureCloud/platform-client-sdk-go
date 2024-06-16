package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulerequest - Learning module request
type Learningmodulerequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of learning module
	Name *string `json:"name,omitempty"`

	// Description - The description of learning module
	Description *string `json:"description,omitempty"`

	// CompletionTimeInDays - The completion time of learning module in days
	CompletionTimeInDays *int `json:"completionTimeInDays,omitempty"`

	// InformSteps - The list of inform steps in a learning module
	InformSteps *[]Learningmoduleinformsteprequest `json:"informSteps,omitempty"`

	// VarType - The type for the learning module
	VarType *string `json:"type,omitempty"`

	// AssessmentForm - The assessment form for learning module
	AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`

	// CoverArt - The cover art for the learning module
	CoverArt *Learningmodulecoverartrequest `json:"coverArt,omitempty"`

	// LengthInMinutes - The recommended time in minutes to complete the module
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`

	// ExcludedFromCatalog - If true, learning module is excluded when retrieving modules for manual assignment
	ExcludedFromCatalog *bool `json:"excludedFromCatalog,omitempty"`

	// ExternalId - The external ID of the learning module. Maximum length: 50 characters.
	ExternalId *string `json:"externalId,omitempty"`

	// EnforceContentOrder - If true, learning module content should be viewed one by one in order
	EnforceContentOrder *bool `json:"enforceContentOrder,omitempty"`

	// ReviewAssessmentResults - Allows to view Assessment results in detail
	ReviewAssessmentResults *Reviewassessmentresults `json:"reviewAssessmentResults,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningmodulerequest) SetField(field string, fieldValue interface{}) {
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

func (o Learningmodulerequest) MarshalJSON() ([]byte, error) {
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
	type Alias Learningmodulerequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CompletionTimeInDays *int `json:"completionTimeInDays,omitempty"`
		
		InformSteps *[]Learningmoduleinformsteprequest `json:"informSteps,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`
		
		CoverArt *Learningmodulecoverartrequest `json:"coverArt,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		ExcludedFromCatalog *bool `json:"excludedFromCatalog,omitempty"`
		
		ExternalId *string `json:"externalId,omitempty"`
		
		EnforceContentOrder *bool `json:"enforceContentOrder,omitempty"`
		
		ReviewAssessmentResults *Reviewassessmentresults `json:"reviewAssessmentResults,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		CompletionTimeInDays: o.CompletionTimeInDays,
		
		InformSteps: o.InformSteps,
		
		VarType: o.VarType,
		
		AssessmentForm: o.AssessmentForm,
		
		CoverArt: o.CoverArt,
		
		LengthInMinutes: o.LengthInMinutes,
		
		ExcludedFromCatalog: o.ExcludedFromCatalog,
		
		ExternalId: o.ExternalId,
		
		EnforceContentOrder: o.EnforceContentOrder,
		
		ReviewAssessmentResults: o.ReviewAssessmentResults,
		Alias:    (Alias)(o),
	})
}

func (o *Learningmodulerequest) UnmarshalJSON(b []byte) error {
	var LearningmodulerequestMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulerequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := LearningmodulerequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := LearningmodulerequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if CompletionTimeInDays, ok := LearningmodulerequestMap["completionTimeInDays"].(float64); ok {
		CompletionTimeInDaysInt := int(CompletionTimeInDays)
		o.CompletionTimeInDays = &CompletionTimeInDaysInt
	}
	
	if InformSteps, ok := LearningmodulerequestMap["informSteps"].([]interface{}); ok {
		InformStepsString, _ := json.Marshal(InformSteps)
		json.Unmarshal(InformStepsString, &o.InformSteps)
	}
	
	if VarType, ok := LearningmodulerequestMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if AssessmentForm, ok := LearningmodulerequestMap["assessmentForm"].(map[string]interface{}); ok {
		AssessmentFormString, _ := json.Marshal(AssessmentForm)
		json.Unmarshal(AssessmentFormString, &o.AssessmentForm)
	}
	
	if CoverArt, ok := LearningmodulerequestMap["coverArt"].(map[string]interface{}); ok {
		CoverArtString, _ := json.Marshal(CoverArt)
		json.Unmarshal(CoverArtString, &o.CoverArt)
	}
	
	if LengthInMinutes, ok := LearningmodulerequestMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if ExcludedFromCatalog, ok := LearningmodulerequestMap["excludedFromCatalog"].(bool); ok {
		o.ExcludedFromCatalog = &ExcludedFromCatalog
	}
    
	if ExternalId, ok := LearningmodulerequestMap["externalId"].(string); ok {
		o.ExternalId = &ExternalId
	}
    
	if EnforceContentOrder, ok := LearningmodulerequestMap["enforceContentOrder"].(bool); ok {
		o.EnforceContentOrder = &EnforceContentOrder
	}
    
	if ReviewAssessmentResults, ok := LearningmodulerequestMap["reviewAssessmentResults"].(map[string]interface{}); ok {
		ReviewAssessmentResultsString, _ := json.Marshal(ReviewAssessmentResults)
		json.Unmarshal(ReviewAssessmentResultsString, &o.ReviewAssessmentResults)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
