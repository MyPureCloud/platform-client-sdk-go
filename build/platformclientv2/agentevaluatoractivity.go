package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Agentevaluatoractivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
