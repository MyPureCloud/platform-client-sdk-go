package platformclientv2
import (
	"encoding/json"
)

// Evaluationassignment
type Evaluationassignment struct { 
	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`


	// User
	User *User `json:"user,omitempty"`

}

// String returns a JSON representation of the model
func (o *Evaluationassignment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
