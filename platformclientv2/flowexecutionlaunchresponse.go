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

func (o *Flowexecutionlaunchresponse) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		FlowVersion: o.FlowVersion,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowexecutionlaunchresponse) UnmarshalJSON(b []byte) error {
	var FlowexecutionlaunchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &FlowexecutionlaunchresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FlowexecutionlaunchresponseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := FlowexecutionlaunchresponseMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if FlowVersion, ok := FlowexecutionlaunchresponseMap["flowVersion"].(map[string]interface{}); ok {
		FlowVersionString, _ := json.Marshal(FlowVersion)
		json.Unmarshal(FlowVersionString, &o.FlowVersion)
	}
	
	if SelfUri, ok := FlowexecutionlaunchresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowexecutionlaunchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
