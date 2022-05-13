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

func (o *Agentevaluatoractivity) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		Agent: o.Agent,
		
		Evaluator: o.Evaluator,
		
		NumEvaluations: o.NumEvaluations,
		
		AverageEvaluationScore: o.AverageEvaluationScore,
		
		NumEvaluationsWithoutViewPermission: o.NumEvaluationsWithoutViewPermission,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Agentevaluatoractivity) UnmarshalJSON(b []byte) error {
	var AgentevaluatoractivityMap map[string]interface{}
	err := json.Unmarshal(b, &AgentevaluatoractivityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AgentevaluatoractivityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AgentevaluatoractivityMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Agent, ok := AgentevaluatoractivityMap["agent"].(map[string]interface{}); ok {
		AgentString, _ := json.Marshal(Agent)
		json.Unmarshal(AgentString, &o.Agent)
	}
	
	if Evaluator, ok := AgentevaluatoractivityMap["evaluator"].(map[string]interface{}); ok {
		EvaluatorString, _ := json.Marshal(Evaluator)
		json.Unmarshal(EvaluatorString, &o.Evaluator)
	}
	
	if NumEvaluations, ok := AgentevaluatoractivityMap["numEvaluations"].(float64); ok {
		NumEvaluationsInt := int(NumEvaluations)
		o.NumEvaluations = &NumEvaluationsInt
	}
	
	if AverageEvaluationScore, ok := AgentevaluatoractivityMap["averageEvaluationScore"].(float64); ok {
		AverageEvaluationScoreInt := int(AverageEvaluationScore)
		o.AverageEvaluationScore = &AverageEvaluationScoreInt
	}
	
	if NumEvaluationsWithoutViewPermission, ok := AgentevaluatoractivityMap["numEvaluationsWithoutViewPermission"].(float64); ok {
		NumEvaluationsWithoutViewPermissionInt := int(NumEvaluationsWithoutViewPermission)
		o.NumEvaluationsWithoutViewPermission = &NumEvaluationsWithoutViewPermissionInt
	}
	
	if SelfUri, ok := AgentevaluatoractivityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agentevaluatoractivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
