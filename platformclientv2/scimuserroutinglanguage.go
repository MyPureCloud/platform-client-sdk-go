package platformclientv2
import (
	"encoding/json"
)

// Scimuserroutinglanguage - Routing language assigned to user.
type Scimuserroutinglanguage struct { 
	// Name - Case-sensitive name identifying a language configured in routing languages.
	Name *string `json:"name,omitempty"`


	// Proficiency - Proficiency is a rating from 0.0 to 5.0 on how competent an agent is for a particular language. It is used when a queue is set to \"Best available language\" mode to allow acd interactions to target agents with higher proficiency ratings.
	Proficiency *float64 `json:"proficiency,omitempty"`

}

// String returns a JSON representation of the model
func (o *Scimuserroutinglanguage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
