package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflowfields
type Architectflowfields struct { 
	// ArchitectFlow - The architect flow.
	ArchitectFlow *Addressableentityref `json:"architectFlow,omitempty"`


	// FlowRequestMappings - Collection of Architect Flow Request Mappings to use.
	FlowRequestMappings *[]Requestmapping `json:"flowRequestMappings,omitempty"`

}

func (o *Architectflowfields) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflowfields
	
	return json.Marshal(&struct { 
		ArchitectFlow *Addressableentityref `json:"architectFlow,omitempty"`
		
		FlowRequestMappings *[]Requestmapping `json:"flowRequestMappings,omitempty"`
		*Alias
	}{ 
		ArchitectFlow: o.ArchitectFlow,
		
		FlowRequestMappings: o.FlowRequestMappings,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectflowfields) UnmarshalJSON(b []byte) error {
	var ArchitectflowfieldsMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflowfieldsMap)
	if err != nil {
		return err
	}
	
	if ArchitectFlow, ok := ArchitectflowfieldsMap["architectFlow"].(map[string]interface{}); ok {
		ArchitectFlowString, _ := json.Marshal(ArchitectFlow)
		json.Unmarshal(ArchitectFlowString, &o.ArchitectFlow)
	}
	
	if FlowRequestMappings, ok := ArchitectflowfieldsMap["flowRequestMappings"].([]interface{}); ok {
		FlowRequestMappingsString, _ := json.Marshal(FlowRequestMappings)
		json.Unmarshal(FlowRequestMappingsString, &o.FlowRequestMappings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflowfields) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
