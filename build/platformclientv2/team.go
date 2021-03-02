package platformclientv2
import (
	"time"
	"encoding/json"
)

// Team
type Team struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The team name
	Name *string `json:"name,omitempty"`


	// Description - Team information.
	Description *string `json:"description,omitempty"`


	// DateModified - Last modified datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// MemberCount - Number of members in a team
	MemberCount *int `json:"memberCount,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Team) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
