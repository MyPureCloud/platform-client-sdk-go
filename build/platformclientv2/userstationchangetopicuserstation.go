package platformclientv2
import (
	"encoding/json"
)

// Userstationchangetopicuserstation
type Userstationchangetopicuserstation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// AssociatedUser
	AssociatedUser *Userstationchangetopicuser `json:"associatedUser,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userstationchangetopicuserstation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
