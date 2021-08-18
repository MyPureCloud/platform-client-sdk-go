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

func (u *Architectflowfields) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflowfields

	

	return json.Marshal(&struct { 
		ArchitectFlow *Addressableentityref `json:"architectFlow,omitempty"`
		
		FlowRequestMappings *[]Requestmapping `json:"flowRequestMappings,omitempty"`
		*Alias
	}{ 
		ArchitectFlow: u.ArchitectFlow,
		
		FlowRequestMappings: u.FlowRequestMappings,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Architectflowfields) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
