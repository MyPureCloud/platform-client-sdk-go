package platformclientv2
import (
	"encoding/json"
)

// Scimuserextensions - Defines a SCIM Genesys Cloud user.
type Scimuserextensions struct { 
	// RoutingSkills - The list of routing skills assigned to a user. Maximum 50 skills.
	RoutingSkills *[]Scimuserroutingskill `json:"routingSkills,omitempty"`


	// RoutingLanguages - The list of routing languages assigned to a user. Maximum 50 languages.
	RoutingLanguages *[]Scimuserroutinglanguage `json:"routingLanguages,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimuserextensions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
