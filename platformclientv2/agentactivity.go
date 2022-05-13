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

func (o *Agentactivity) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		Agent: o.Agent,
		
		NumEvaluations: o.NumEvaluations,
		
		AverageEvaluationScore: o.AverageEvaluationScore,
		
		NumCriticalEvaluations: o.NumCriticalEvaluations,
		
		AverageCriticalScore: o.AverageCriticalScore,
		
		HighestEvaluationScore: o.HighestEvaluationScore,
		
		LowestEvaluationScore: o.LowestEvaluationScore,
		
		HighestCriticalScore: o.HighestCriticalScore,
		
		LowestCriticalScore: o.LowestCriticalScore,
		
		AgentEvaluatorActivityList: o.AgentEvaluatorActivityList,
		
		NumEvaluationsWithoutViewPermission: o.NumEvaluationsWithoutViewPermission,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Agentactivity) UnmarshalJSON(b []byte) error {
	var AgentactivityMap map[string]interface{}
	err := json.Unmarshal(b, &AgentactivityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AgentactivityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AgentactivityMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Agent, ok := AgentactivityMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if NumEvaluations, ok := AgentactivityMap["numEvaluations"].(float64); ok {
		NumEvaluationsInt := int(NumEvaluations)
		o.NumEvaluations = &NumEvaluationsInt
	}
	
	if AverageEvaluationScore, ok := AgentactivityMap["averageEvaluationScore"].(float64); ok {
		AverageEvaluationScoreInt := int(AverageEvaluationScore)
		o.AverageEvaluationScore = &AverageEvaluationScoreInt
	}
	
	if NumCriticalEvaluations, ok := AgentactivityMap["numCriticalEvaluations"].(float64); ok {
		NumCriticalEvaluationsInt := int(NumCriticalEvaluations)
		o.NumCriticalEvaluations = &NumCriticalEvaluationsInt
	}
	
	if AverageCriticalScore, ok := AgentactivityMap["averageCriticalScore"].(float64); ok {
		AverageCriticalScoreFloat32 := float32(AverageCriticalScore)
		o.AverageCriticalScore = &AverageCriticalScoreFloat32
	}
	
	if HighestEvaluationScore, ok := AgentactivityMap["highestEvaluationScore"].(float64); ok {
		HighestEvaluationScoreFloat32 := float32(HighestEvaluationScore)
		o.HighestEvaluationScore = &HighestEvaluationScoreFloat32
	}
	
	if LowestEvaluationScore, ok := AgentactivityMap["lowestEvaluationScore"].(float64); ok {
		LowestEvaluationScoreFloat32 := float32(LowestEvaluationScore)
		o.LowestEvaluationScore = &LowestEvaluationScoreFloat32
	}
	
	if HighestCriticalScore, ok := AgentactivityMap["highestCriticalScore"].(float64); ok {
		HighestCriticalScoreFloat32 := float32(HighestCriticalScore)
		o.HighestCriticalScore = &HighestCriticalScoreFloat32
	}
	
	if LowestCriticalScore, ok := AgentactivityMap["lowestCriticalScore"].(float64); ok {
		LowestCriticalScoreFloat32 := float32(LowestCriticalScore)
		o.LowestCriticalScore = &LowestCriticalScoreFloat32
	}
	
	if AgentEvaluatorActivityList, ok := AgentactivityMap["agentEvaluatorActivityList"].([]interface{}); ok {
		AgentEvaluatorActivityListString, _ := json.Marshal(AgentEvaluatorActivityList)
		json.Unmarshal(AgentEvaluatorActivityListString, &o.AgentEvaluatorActivityList)
	}
	
	if NumEvaluationsWithoutViewPermission, ok := AgentactivityMap["numEvaluationsWithoutViewPermission"].(float64); ok {
		NumEvaluationsWithoutViewPermissionInt := int(NumEvaluationsWithoutViewPermission)
		o.NumEvaluationsWithoutViewPermission = &NumEvaluationsWithoutViewPermissionInt
	}
	
	if SelfUri, ok := AgentactivityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
