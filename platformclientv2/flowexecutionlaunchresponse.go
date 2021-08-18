package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Flowexecutionlaunchresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowexecutionlaunchresponse

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		FlowVersion *Domainentityref `json:"flowVersion,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		FlowVersion: u.FlowVersion,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Flowexecutionlaunchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
