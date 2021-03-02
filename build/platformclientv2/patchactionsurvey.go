package platformclientv2
import (
	"encoding/json"
)

// Patchactionsurvey
type Patchactionsurvey struct { 
	// Questions - Questions shown to the user.
	Questions *[]Patchsurveyquestion `json:"questions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchactionsurvey) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
