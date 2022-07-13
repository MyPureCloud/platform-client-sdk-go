package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsagentgroup
type Analyticsagentgroup struct { 
	// AgentGroupId - Conditional group routing agent group identifier
	AgentGroupId *string `json:"agentGroupId,omitempty"`


	// AgentGroupType - Conditional group routing agent group type
	AgentGroupType *string `json:"agentGroupType,omitempty"`

}

func (o *Analyticsagentgroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsagentgroup
	
	return json.Marshal(&struct { 
		AgentGroupId *string `json:"agentGroupId,omitempty"`
		
		AgentGroupType *string `json:"agentGroupType,omitempty"`
		*Alias
	}{ 
		AgentGroupId: o.AgentGroupId,
		
		AgentGroupType: o.AgentGroupType,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsagentgroup) UnmarshalJSON(b []byte) error {
	var AnalyticsagentgroupMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsagentgroupMap)
	if err != nil {
		return err
	}
	
	if AgentGroupId, ok := AnalyticsagentgroupMap["agentGroupId"].(string); ok {
		o.AgentGroupId = &AgentGroupId
	}
    
	if AgentGroupType, ok := AnalyticsagentgroupMap["agentGroupType"].(string); ok {
		o.AgentGroupType = &AgentGroupType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsagentgroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
