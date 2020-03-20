package platformclientv2
import (
	"encoding/json"
)

// Userroutinglanguage - Represents an organization langauge assigned to a user. When assigning to a user specify the organization language id as the id.
type Userroutinglanguage struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Proficiency - Proficiency is a rating from 0.0 to 5.0 on how competent an agent is for a particular language. It is used when a queue is set to \"Best available language\" mode to allow acd interactions to target agents with higher proficiency ratings.
	Proficiency *float64 `json:"proficiency,omitempty"`


	// State - Activate or deactivate this routing langauge.
	State *string `json:"state,omitempty"`


	// LanguageUri - URI to the organization language used by this user langauge.
	LanguageUri *string `json:"languageUri,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userroutinglanguage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
