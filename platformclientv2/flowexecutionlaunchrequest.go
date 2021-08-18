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

func (u *Flowexecutionlaunchrequest) MarshalJSON() ([]byte, error) {
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
		FlowId: u.FlowId,
		
		FlowVersion: u.FlowVersion,
		
		InputData: u.InputData,
		
		Name: u.Name,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Flowexecutionlaunchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
