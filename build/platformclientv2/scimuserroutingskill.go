package platformclientv2
import (
	"encoding/json"
)

// Scimuserroutingskill - The routing skill assigned to a user.
type Scimuserroutingskill struct { 
	// Name - The case-sensitive name of a routing skill configured in Genesys Cloud.
	Name *string `json:"name,omitempty"`


	// Proficiency - A rating from 0.0 to 5.0 that indicates how adept an agent is at a particular skill. When \"Best available skills\" is enabled for a queue in Genesys Cloud, ACD interactions in that queue are routed to agents with higher proficiency ratings.
	Proficiency *float64 `json:"proficiency,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimuserroutingskill) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
