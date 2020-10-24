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


	// ExternalIds - The list of external identifiers assigned to user. Always includes an immutable SCIM authority prefixed with \"x-pc:scimv2:v1\".
	ExternalIds *[]Scimgenesysuserexternalid `json:"externalIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimuserextensions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
