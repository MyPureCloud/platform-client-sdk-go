package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgesearchclientapplication
type Knowledgesearchclientapplication struct { 
	// VarType - Application type.
	VarType *string `json:"type,omitempty"`


	// Deployment - Application details when type is MessengerKnowledgeApp or SupportCenter.
	Deployment *Addressableentityref `json:"deployment,omitempty"`


	// BotFlow - Application details when type is BotFlow.
	BotFlow *Addressableentityref `json:"botFlow,omitempty"`


	// Assistant - Application details when type is Assistant.
	Assistant *Addressableentityref `json:"assistant,omitempty"`

}

func (o *Knowledgesearchclientapplication) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgesearchclientapplication
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Deployment *Addressableentityref `json:"deployment,omitempty"`
		
		BotFlow *Addressableentityref `json:"botFlow,omitempty"`
		
		Assistant *Addressableentityref `json:"assistant,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Deployment: o.Deployment,
		
		BotFlow: o.BotFlow,
		
		Assistant: o.Assistant,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgesearchclientapplication) UnmarshalJSON(b []byte) error {
	var KnowledgesearchclientapplicationMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgesearchclientapplicationMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := KnowledgesearchclientapplicationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Deployment, ok := KnowledgesearchclientapplicationMap["deployment"].(map[string]interface{}); ok {
		DeploymentString, _ := json.Marshal(Deployment)
		json.Unmarshal(DeploymentString, &o.Deployment)
	}
	
	if BotFlow, ok := KnowledgesearchclientapplicationMap["botFlow"].(map[string]interface{}); ok {
		BotFlowString, _ := json.Marshal(BotFlow)
		json.Unmarshal(BotFlowString, &o.BotFlow)
	}
	
	if Assistant, ok := KnowledgesearchclientapplicationMap["assistant"].(map[string]interface{}); ok {
		AssistantString, _ := json.Marshal(Assistant)
		json.Unmarshal(AssistantString, &o.Assistant)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgesearchclientapplication) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
