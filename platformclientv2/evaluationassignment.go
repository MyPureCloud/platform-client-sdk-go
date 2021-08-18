package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Evaluationassignment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationassignment

	

	return json.Marshal(&struct { 
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		User *User `json:"user,omitempty"`
		*Alias
	}{ 
		EvaluationForm: u.EvaluationForm,
		
		User: u.User,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Evaluationassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
