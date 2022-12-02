package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicworkflow - Information about the workflow.
type Queueconversationeventtopicworkflow struct { 
	// WorkflowId - The id of the workflow
	WorkflowId *string `json:"workflowId,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationeventtopicworkflow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationeventtopicworkflow
	
	return json.Marshal(&struct { 
		WorkflowId *string `json:"workflowId,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		WorkflowId: o.WorkflowId,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationeventtopicworkflow) UnmarshalJSON(b []byte) error {
	var QueueconversationeventtopicworkflowMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationeventtopicworkflowMap)
	if err != nil {
		return err
	}
	
	if WorkflowId, ok := QueueconversationeventtopicworkflowMap["workflowId"].(string); ok {
		o.WorkflowId = &WorkflowId
	}
    
	if AdditionalProperties, ok := QueueconversationeventtopicworkflowMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicworkflow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
