package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Flowexecutionlaunchresponse - Response object from launching a flow.
type Flowexecutionlaunchresponse struct { 
	// Id - The flow execution ID
	Id *string `json:"id,omitempty"`


	// Name - The flow execution name.
	Name *string `json:"name,omitempty"`


	// FlowVersion - The version of the flow that launched
	FlowVersion *Domainentityref `json:"flowVersion,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Flowexecutionlaunchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
