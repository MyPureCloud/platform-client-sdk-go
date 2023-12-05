package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignment - Learning module assignment with user information
type Learningassignment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Assessment - The assessment associated with this assignment
	Assessment *Learningassessment `json:"assessment,omitempty"`

	// CreatedBy - The user who created the assignment
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// DateCreated - The date when the assignment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// ModifiedBy - The user who modified the assignment
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// DateModified - The date when the assignment was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// IsOverdue - True if the assignment is overdue
	IsOverdue *bool `json:"isOverdue,omitempty"`

	// PercentageScore - The user's percentage score for this assignment
	PercentageScore *float32 `json:"percentageScore,omitempty"`

	// IsRule - True if this assignment was created by a Rule
	IsRule *bool `json:"isRule,omitempty"`

	// IsManual - True if this assignment was created manually
	IsManual *bool `json:"isManual,omitempty"`

	// IsPassed - True if the assessment was passed
	IsPassed *bool `json:"isPassed,omitempty"`

	// IsLatest - True if the assignment is based on latest module
	IsLatest *bool `json:"isLatest,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// State - The Learning Assignment state
	State *string `json:"state,omitempty"`

	// DateRecommendedForCompletion - The recommended completion date of the assignment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateRecommendedForCompletion *time.Time `json:"dateRecommendedForCompletion,omitempty"`

	// Version - The version of Learning module assigned
	Version *int `json:"version,omitempty"`

	// Module - The Learning module object associated with this assignment
	Module *Learningmodule `json:"module,omitempty"`

	// User - The user to whom the assignment is assigned
	User *Userreference `json:"user,omitempty"`

	// AssessmentForm - The assessment form associated with this assignment
	AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`

	// LengthInMinutes - The length in minutes of the assignment
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningassignment) SetField(field string, fieldValue interface{}) {
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

func (o Learningassignment) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified","DateRecommendedForCompletion", }
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
	type Alias Learningassignment
	
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
	
	DateRecommendedForCompletion := new(string)
	if o.DateRecommendedForCompletion != nil {
		
		*DateRecommendedForCompletion = timeutil.Strftime(o.DateRecommendedForCompletion, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateRecommendedForCompletion = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Assessment *Learningassessment `json:"assessment,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		IsOverdue *bool `json:"isOverdue,omitempty"`
		
		PercentageScore *float32 `json:"percentageScore,omitempty"`
		
		IsRule *bool `json:"isRule,omitempty"`
		
		IsManual *bool `json:"isManual,omitempty"`
		
		IsPassed *bool `json:"isPassed,omitempty"`
		
		IsLatest *bool `json:"isLatest,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateRecommendedForCompletion *string `json:"dateRecommendedForCompletion,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Module *Learningmodule `json:"module,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Assessment: o.Assessment,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		
		IsOverdue: o.IsOverdue,
		
		PercentageScore: o.PercentageScore,
		
		IsRule: o.IsRule,
		
		IsManual: o.IsManual,
		
		IsPassed: o.IsPassed,
		
		IsLatest: o.IsLatest,
		
		SelfUri: o.SelfUri,
		
		State: o.State,
		
		DateRecommendedForCompletion: DateRecommendedForCompletion,
		
		Version: o.Version,
		
		Module: o.Module,
		
		User: o.User,
		
		AssessmentForm: o.AssessmentForm,
		
		LengthInMinutes: o.LengthInMinutes,
		Alias:    (Alias)(o),
	})
}

func (o *Learningassignment) UnmarshalJSON(b []byte) error {
	var LearningassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LearningassignmentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Assessment, ok := LearningassignmentMap["assessment"].(map[string]interface{}); ok {
		AssessmentString, _ := json.Marshal(Assessment)
		json.Unmarshal(AssessmentString, &o.Assessment)
	}
	
	if CreatedBy, ok := LearningassignmentMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := LearningassignmentMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ModifiedBy, ok := LearningassignmentMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := LearningassignmentMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if IsOverdue, ok := LearningassignmentMap["isOverdue"].(bool); ok {
		o.IsOverdue = &IsOverdue
	}
    
	if PercentageScore, ok := LearningassignmentMap["percentageScore"].(float64); ok {
		PercentageScoreFloat32 := float32(PercentageScore)
		o.PercentageScore = &PercentageScoreFloat32
	}
	
	if IsRule, ok := LearningassignmentMap["isRule"].(bool); ok {
		o.IsRule = &IsRule
	}
    
	if IsManual, ok := LearningassignmentMap["isManual"].(bool); ok {
		o.IsManual = &IsManual
	}
    
	if IsPassed, ok := LearningassignmentMap["isPassed"].(bool); ok {
		o.IsPassed = &IsPassed
	}
    
	if IsLatest, ok := LearningassignmentMap["isLatest"].(bool); ok {
		o.IsLatest = &IsLatest
	}
    
	if SelfUri, ok := LearningassignmentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if State, ok := LearningassignmentMap["state"].(string); ok {
		o.State = &State
	}
    
	if dateRecommendedForCompletionString, ok := LearningassignmentMap["dateRecommendedForCompletion"].(string); ok {
		DateRecommendedForCompletion, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateRecommendedForCompletionString)
		o.DateRecommendedForCompletion = &DateRecommendedForCompletion
	}
	
	if Version, ok := LearningassignmentMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Module, ok := LearningassignmentMap["module"].(map[string]interface{}); ok {
		ModuleString, _ := json.Marshal(Module)
		json.Unmarshal(ModuleString, &o.Module)
	}
	
	if User, ok := LearningassignmentMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if AssessmentForm, ok := LearningassignmentMap["assessmentForm"].(map[string]interface{}); ok {
		AssessmentFormString, _ := json.Marshal(AssessmentForm)
		json.Unmarshal(AssessmentFormString, &o.AssessmentForm)
	}
	
	if LengthInMinutes, ok := LearningassignmentMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
