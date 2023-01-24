package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmenttopiclearningassignmentnotification
type Learningassignmenttopiclearningassignmentnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// User
	User *Learningassignmenttopicuserreference `json:"user,omitempty"`

	// Module
	Module *Learningassignmenttopiclearningmodulereference `json:"module,omitempty"`

	// Version
	Version *int `json:"version,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// DateRecommendedForCompletion
	DateRecommendedForCompletion *time.Time `json:"dateRecommendedForCompletion,omitempty"`

	// CreatedBy
	CreatedBy *Learningassignmenttopicuserreference `json:"createdBy,omitempty"`

	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// ModifiedBy
	ModifiedBy *Learningassignmenttopicuserreference `json:"modifiedBy,omitempty"`

	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`

	// IsOverdue
	IsOverdue *bool `json:"isOverdue,omitempty"`

	// LengthInMinutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Learningassignmenttopiclearningassignmentnotification) SetField(field string, fieldValue interface{}) {
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

func (o Learningassignmenttopiclearningassignmentnotification) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateRecommendedForCompletion","DateCreated","DateModified", }
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
	type Alias Learningassignmenttopiclearningassignmentnotification
	
	DateRecommendedForCompletion := new(string)
	if o.DateRecommendedForCompletion != nil {
		
		*DateRecommendedForCompletion = timeutil.Strftime(o.DateRecommendedForCompletion, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateRecommendedForCompletion = nil
	}
	
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
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		User *Learningassignmenttopicuserreference `json:"user,omitempty"`
		
		Module *Learningassignmenttopiclearningmodulereference `json:"module,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateRecommendedForCompletion *string `json:"dateRecommendedForCompletion,omitempty"`
		
		CreatedBy *Learningassignmenttopicuserreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ModifiedBy *Learningassignmenttopicuserreference `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		IsOverdue *bool `json:"isOverdue,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		User: o.User,
		
		Module: o.Module,
		
		Version: o.Version,
		
		State: o.State,
		
		DateRecommendedForCompletion: DateRecommendedForCompletion,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		
		IsOverdue: o.IsOverdue,
		
		LengthInMinutes: o.LengthInMinutes,
		Alias:    (Alias)(o),
	})
}

func (o *Learningassignmenttopiclearningassignmentnotification) UnmarshalJSON(b []byte) error {
	var LearningassignmenttopiclearningassignmentnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmenttopiclearningassignmentnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LearningassignmenttopiclearningassignmentnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if User, ok := LearningassignmenttopiclearningassignmentnotificationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Module, ok := LearningassignmenttopiclearningassignmentnotificationMap["module"].(map[string]interface{}); ok {
		ModuleString, _ := json.Marshal(Module)
		json.Unmarshal(ModuleString, &o.Module)
	}
	
	if Version, ok := LearningassignmenttopiclearningassignmentnotificationMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if State, ok := LearningassignmenttopiclearningassignmentnotificationMap["state"].(string); ok {
		o.State = &State
	}
    
	if dateRecommendedForCompletionString, ok := LearningassignmenttopiclearningassignmentnotificationMap["dateRecommendedForCompletion"].(string); ok {
		DateRecommendedForCompletion, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateRecommendedForCompletionString)
		o.DateRecommendedForCompletion = &DateRecommendedForCompletion
	}
	
	if CreatedBy, ok := LearningassignmenttopiclearningassignmentnotificationMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := LearningassignmenttopiclearningassignmentnotificationMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ModifiedBy, ok := LearningassignmenttopiclearningassignmentnotificationMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := LearningassignmenttopiclearningassignmentnotificationMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if IsOverdue, ok := LearningassignmenttopiclearningassignmentnotificationMap["isOverdue"].(bool); ok {
		o.IsOverdue = &IsOverdue
	}
    
	if LengthInMinutes, ok := LearningassignmenttopiclearningassignmentnotificationMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmenttopiclearningassignmentnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
