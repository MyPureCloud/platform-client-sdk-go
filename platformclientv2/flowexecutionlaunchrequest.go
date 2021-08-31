package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Flowexecutionlaunchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowexecutionlaunchrequest
	
	return json.Marshal(&struct { 
		FlowId *string `json:"flowId,omitempty"`
		
		FlowVersion *string `json:"flowVersion,omitempty"`
		
		InputData *map[string]interface{} `json:"inputData,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		FlowId: o.FlowId,
		
		FlowVersion: o.FlowVersion,
		
		InputData: o.InputData,
		
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowexecutionlaunchrequest) UnmarshalJSON(b []byte) error {
	var FlowexecutionlaunchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &FlowexecutionlaunchrequestMap)
	if err != nil {
		return err
	}
	
	if FlowId, ok := FlowexecutionlaunchrequestMap["flowId"].(string); ok {
		o.FlowId = &FlowId
	}
	
	if FlowVersion, ok := FlowexecutionlaunchrequestMap["flowVersion"].(string); ok {
		o.FlowVersion = &FlowVersion
	}
	
	if InputData, ok := FlowexecutionlaunchrequestMap["inputData"].(map[string]interface{}); ok {
		InputDataString, _ := json.Marshal(InputData)
		json.Unmarshal(InputDataString, &o.InputData)
	}
	
	if Name, ok := FlowexecutionlaunchrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowexecutionlaunchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
