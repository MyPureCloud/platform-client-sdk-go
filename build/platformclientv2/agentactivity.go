package platformclientv2
import (
	"encoding/json"
)

// Agentactivity
type Agentactivity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Agent
	Agent *User `json:"agent,omitempty"`


	// NumEvaluations
	NumEvaluations *int `json:"numEvaluations,omitempty"`


	// AverageEvaluationScore
	AverageEvaluationScore *int `json:"averageEvaluationScore,omitempty"`


	// NumCriticalEvaluations
	NumCriticalEvaluations *int `json:"numCriticalEvaluations,omitempty"`


	// AverageCriticalScore
	AverageCriticalScore *float32 `json:"averageCriticalScore,omitempty"`


	// HighestEvaluationScore
	HighestEvaluationScore *float32 `json:"highestEvaluationScore,omitempty"`


	// LowestEvaluationScore
	LowestEvaluationScore *float32 `json:"lowestEvaluationScore,omitempty"`


	// HighestCriticalScore
	HighestCriticalScore *float32 `json:"highestCriticalScore,omitempty"`


	// LowestCriticalScore
	LowestCriticalScore *float32 `json:"lowestCriticalScore,omitempty"`


	// AgentEvaluatorActivityList
	AgentEvaluatorActivityList *[]Agentevaluatoractivity `json:"agentEvaluatorActivityList,omitempty"`


	// NumEvaluationsWithoutViewPermission
	NumEvaluationsWithoutViewPermission *int `json:"numEvaluationsWithoutViewPermission,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Agentactivity) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
