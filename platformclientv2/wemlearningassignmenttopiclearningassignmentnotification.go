package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Wemlearningassignmenttopiclearningassignmentnotification
type Wemlearningassignmenttopiclearningassignmentnotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// User
	User *Wemlearningassignmenttopicuserreference `json:"user,omitempty"`


	// Module
	Module *Wemlearningassignmenttopiclearningmodulereference `json:"module,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// DateRecommendedForCompletion
	DateRecommendedForCompletion *time.Time `json:"dateRecommendedForCompletion,omitempty"`


	// CreatedBy
	CreatedBy *Wemlearningassignmenttopicuserreference `json:"createdBy,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ModifiedBy
	ModifiedBy *Wemlearningassignmenttopicuserreference `json:"modifiedBy,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// IsOverdue
	IsOverdue *bool `json:"isOverdue,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wemlearningassignmenttopiclearningassignmentnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
