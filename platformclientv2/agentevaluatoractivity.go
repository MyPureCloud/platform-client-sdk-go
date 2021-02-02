package platformclientv2
import (
	"encoding/json"
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


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Agentevaluatoractivity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
