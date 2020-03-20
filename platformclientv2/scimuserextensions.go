package platformclientv2
import (
	"encoding/json"
)

// Scimuserextensions - SCIM PureCloud extensions of user.
type Scimuserextensions struct { 
	// RoutingSkills - Routing Skills assigned to user. No more than 50 skills may be assigned to a user.
	RoutingSkills *[]Scimuserroutingskill `json:"routingSkills,omitempty"`


	// RoutingLanguages - Routing Languages assigned to user. No more than 50 languages may be assigned to a user.
	RoutingLanguages *[]Scimuserroutinglanguage `json:"routingLanguages,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimuserextensions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
