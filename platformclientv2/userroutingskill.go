package platformclientv2
import (
	"encoding/json"
)

// Userroutingskill - Represents an organization skill assigned to a user. When assigning to a user specify the organization skill id as the id.
type Userroutingskill struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Proficiency - A rating from 0.0 to 5.0 that indicates how adept an agent is at a particular skill. When \"Best available skills\" is enabled for a queue in Genesys Cloud, ACD interactions in that queue are routed to agents with higher proficiency ratings.
	Proficiency *float64 `json:"proficiency,omitempty"`


	// State - Activate or deactivate this routing skill.
	State *string `json:"state,omitempty"`


	// SkillUri - URI to the organization skill used by this user skill.
	SkillUri *string `json:"skillUri,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userroutingskill) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
