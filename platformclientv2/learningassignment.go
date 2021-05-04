package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignment - Learning module assignment with user information
type Learningassignment struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


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


	// IsRule - True if this assignment was created by a Rule
	IsRule *bool `json:"isRule,omitempty"`


	// IsManual - True if this assignment was created manually
	IsManual *bool `json:"isManual,omitempty"`


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

}

// String returns a JSON representation of the model
func (o *Learningassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
