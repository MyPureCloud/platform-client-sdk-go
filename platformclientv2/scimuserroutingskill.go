package platformclientv2
import (
	"encoding/json"
)

// Scimuserroutingskill - Routing skill assigned to user.
type Scimuserroutingskill struct { 
	// Name - Case-sensitive name identifying a skill configured in routing skills.
	Name *string `json:"name,omitempty"`


	// Proficiency - Proficiency is a rating from 0.0 to 5.0 on how competent an agent is for a particular skill. It is used when a queue is set to \"Best available skills\" mode to allow acd interactions to target agents with higher proficiency ratings.
	Proficiency *float64 `json:"proficiency,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimuserroutingskill) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
