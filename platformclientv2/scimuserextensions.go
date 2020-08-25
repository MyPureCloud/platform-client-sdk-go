package platformclientv2
import (
	"encoding/json"
)

// Scimuserextensions - Genesys Cloud user extensions to SCIM RFC.
type Scimuserextensions struct { 
	// RoutingSkills - The list of routing skills assigned to a user. Maximum 50 skills.
	RoutingSkills *[]Scimuserroutingskill `json:"routingSkills,omitempty"`


	// RoutingLanguages - The list of routing languages assigned to a user. Maximum 50 languages.
	RoutingLanguages *[]Scimuserroutinglanguage `json:"routingLanguages,omitempty"`


	// ExternalIds - External Identifiers assigned to user. SCIM External ID will be visible here with authority prefix 'x-pc:scimv2:v1' but will be immutable.
	ExternalIds *[]Scimgenesysuserexternalid `json:"externalIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimuserextensions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
