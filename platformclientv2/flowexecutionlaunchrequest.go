package platformclientv2
import (
	"encoding/json"
)

// Flowexecutionlaunchrequest - Parameters for launching a flow.
type Flowexecutionlaunchrequest struct { 
	// FlowId - ID of the flow to launch.
	FlowId *string `json:"flowId,omitempty"`


	// FlowVersion - The version of the flow to launch. Omit this value (or supply null/empty) to use the latest published version.
	FlowVersion *string `json:"flowVersion,omitempty"`


	// InputData - Input values to the flow. Valid values are defined by a flow's input JSON schema.
	InputData *map[string]interface{} `json:"inputData,omitempty"`


	// Name - A displayable name to assign to the new flow execution
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowexecutionlaunchrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
