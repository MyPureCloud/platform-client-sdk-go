package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
