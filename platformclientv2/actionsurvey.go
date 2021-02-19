package platformclientv2
import (
	"encoding/json"
)

// Actionsurvey
type Actionsurvey struct { 
	// Questions - Questions shown to the user.
	Questions *[]Journeysurveyquestion `json:"questions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Actionsurvey) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
