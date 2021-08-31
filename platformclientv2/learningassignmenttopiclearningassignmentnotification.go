package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmenttopiclearningassignmentnotification
type Learningassignmenttopiclearningassignmentnotification struct { 
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

}

func (o *Learningassignmenttopiclearningassignmentnotification) MarshalJSON() ([]byte, error) {
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
		*Alias
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
		Alias:    (*Alias)(o),
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmenttopiclearningassignmentnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
