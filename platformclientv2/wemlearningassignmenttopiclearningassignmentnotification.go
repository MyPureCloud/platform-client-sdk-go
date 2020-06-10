package platformclientv2
import (
	"time"
	"encoding/json"
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
	Version *int32 `json:"version,omitempty"`


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

}

// String returns a JSON representation of the model
func (o *Wemlearningassignmenttopiclearningassignmentnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
