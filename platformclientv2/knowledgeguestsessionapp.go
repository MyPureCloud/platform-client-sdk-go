package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeguestsessionapp
type Knowledgeguestsessionapp struct { 
	// DeploymentId - App deployment ID.
	DeploymentId *string `json:"deploymentId,omitempty"`


	// VarType - App type.
	VarType *string `json:"type,omitempty"`

}

func (o *Knowledgeguestsessionapp) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeguestsessionapp
	
	return json.Marshal(&struct { 
		DeploymentId *string `json:"deploymentId,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		DeploymentId: o.DeploymentId,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeguestsessionapp) UnmarshalJSON(b []byte) error {
	var KnowledgeguestsessionappMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeguestsessionappMap)
	if err != nil {
		return err
	}
	
	if DeploymentId, ok := KnowledgeguestsessionappMap["deploymentId"].(string); ok {
		o.DeploymentId = &DeploymentId
	}
    
	if VarType, ok := KnowledgeguestsessionappMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeguestsessionapp) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
