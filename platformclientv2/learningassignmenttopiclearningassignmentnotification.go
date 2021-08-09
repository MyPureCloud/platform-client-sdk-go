package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Learningassignmenttopiclearningassignmentnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
