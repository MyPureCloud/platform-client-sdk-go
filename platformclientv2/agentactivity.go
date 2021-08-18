package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Agentactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentactivity

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Agent *User `json:"agent,omitempty"`
		
		NumEvaluations *int `json:"numEvaluations,omitempty"`
		
		AverageEvaluationScore *int `json:"averageEvaluationScore,omitempty"`
		
		NumCriticalEvaluations *int `json:"numCriticalEvaluations,omitempty"`
		
		AverageCriticalScore *float32 `json:"averageCriticalScore,omitempty"`
		
		HighestEvaluationScore *float32 `json:"highestEvaluationScore,omitempty"`
		
		LowestEvaluationScore *float32 `json:"lowestEvaluationScore,omitempty"`
		
		HighestCriticalScore *float32 `json:"highestCriticalScore,omitempty"`
		
		LowestCriticalScore *float32 `json:"lowestCriticalScore,omitempty"`
		
		AgentEvaluatorActivityList *[]Agentevaluatoractivity `json:"agentEvaluatorActivityList,omitempty"`
		
		NumEvaluationsWithoutViewPermission *int `json:"numEvaluationsWithoutViewPermission,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Agent: u.Agent,
		
		NumEvaluations: u.NumEvaluations,
		
		AverageEvaluationScore: u.AverageEvaluationScore,
		
		NumCriticalEvaluations: u.NumCriticalEvaluations,
		
		AverageCriticalScore: u.AverageCriticalScore,
		
		HighestEvaluationScore: u.HighestEvaluationScore,
		
		LowestEvaluationScore: u.LowestEvaluationScore,
		
		HighestCriticalScore: u.HighestCriticalScore,
		
		LowestCriticalScore: u.LowestCriticalScore,
		
		AgentEvaluatorActivityList: u.AgentEvaluatorActivityList,
		
		NumEvaluationsWithoutViewPermission: u.NumEvaluationsWithoutViewPermission,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Agentactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
