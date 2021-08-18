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

func (u *Learningassignmenttopiclearningassignmentnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmenttopiclearningassignmentnotification

	
	DateRecommendedForCompletion := new(string)
	if u.DateRecommendedForCompletion != nil {
		
		*DateRecommendedForCompletion = timeutil.Strftime(u.DateRecommendedForCompletion, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateRecommendedForCompletion = nil
	}
	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		User: u.User,
		
		Module: u.Module,
		
		Version: u.Version,
		
		State: u.State,
		
		DateRecommendedForCompletion: DateRecommendedForCompletion,
		
		CreatedBy: u.CreatedBy,
		
		DateCreated: DateCreated,
		
		ModifiedBy: u.ModifiedBy,
		
		DateModified: DateModified,
		
		IsOverdue: u.IsOverdue,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassignmenttopiclearningassignmentnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
