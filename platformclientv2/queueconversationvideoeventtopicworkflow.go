package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicworkflow - Information about the workflow.
type Queueconversationvideoeventtopicworkflow struct { 
	// WorkflowId - The id of the workflow
	WorkflowId *string `json:"workflowId,omitempty"`

}

func (o *Queueconversationvideoeventtopicworkflow) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicworkflow
	
	return json.Marshal(&struct { 
		WorkflowId *string `json:"workflowId,omitempty"`
		*Alias
	}{ 
		WorkflowId: o.WorkflowId,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationvideoeventtopicworkflow) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicworkflowMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicworkflowMap)
	if err != nil {
		return err
	}
	
	if WorkflowId, ok := QueueconversationvideoeventtopicworkflowMap["workflowId"].(string); ok {
		o.WorkflowId = &WorkflowId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicworkflow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
