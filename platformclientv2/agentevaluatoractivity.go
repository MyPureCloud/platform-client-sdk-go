package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentevaluatoractivity
type Agentevaluatoractivity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Agent
	Agent *User `json:"agent,omitempty"`


	// Evaluator
	Evaluator *User `json:"evaluator,omitempty"`


	// NumEvaluations
	NumEvaluations *int `json:"numEvaluations,omitempty"`


	// AverageEvaluationScore
	AverageEvaluationScore *int `json:"averageEvaluationScore,omitempty"`


	// NumEvaluationsWithoutViewPermission
	NumEvaluationsWithoutViewPermission *int `json:"numEvaluationsWithoutViewPermission,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Agentevaluatoractivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentevaluatoractivity

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Agent *User `json:"agent,omitempty"`
		
		Evaluator *User `json:"evaluator,omitempty"`
		
		NumEvaluations *int `json:"numEvaluations,omitempty"`
		
		AverageEvaluationScore *int `json:"averageEvaluationScore,omitempty"`
		
		NumEvaluationsWithoutViewPermission *int `json:"numEvaluationsWithoutViewPermission,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Agent: u.Agent,
		
		Evaluator: u.Evaluator,
		
		NumEvaluations: u.NumEvaluations,
		
		AverageEvaluationScore: u.AverageEvaluationScore,
		
		NumEvaluationsWithoutViewPermission: u.NumEvaluationsWithoutViewPermission,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Agentevaluatoractivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
