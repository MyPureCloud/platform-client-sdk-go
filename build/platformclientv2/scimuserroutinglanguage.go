package platformclientv2
import (
	"encoding/json"
)

// Scimuserroutinglanguage - The routing language assigned to a user.
type Scimuserroutinglanguage struct { 
	// Name - The case-sensitive name of a routing language configured in Genesys Cloud.
	Name *string `json:"name,omitempty"`


	// Proficiency - A rating from 0.0 to 5.0 that indicates how fluent an agent is in a particular language. ACD interactions are routed to agents with higher proficiency ratings.
	Proficiency *float64 `json:"proficiency,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimuserroutinglanguage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
