package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectjobstateresponse
type Architectjobstateresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Flow - Flow created from the Architect Job
	Flow *Addressableentityref `json:"flow,omitempty"`


	// Status - Status of the Architect Job
	Status *string `json:"status,omitempty"`


	// Command - The command executed by the Architect Job
	Command *string `json:"command,omitempty"`


	// Messages - Warnings and Errors messages of the Architect Job
	Messages *[]Architectjobmessage `json:"messages,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Architectjobstateresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectjobstateresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Flow *Addressableentityref `json:"flow,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Command *string `json:"command,omitempty"`
		
		Messages *[]Architectjobmessage `json:"messages,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Flow: o.Flow,
		
		Status: o.Status,
		
		Command: o.Command,
		
		Messages: o.Messages,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectjobstateresponse) UnmarshalJSON(b []byte) error {
	var ArchitectjobstateresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectjobstateresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectjobstateresponseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Flow, ok := ArchitectjobstateresponseMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if Status, ok := ArchitectjobstateresponseMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if Command, ok := ArchitectjobstateresponseMap["command"].(string); ok {
		o.Command = &Command
	}
	
	if Messages, ok := ArchitectjobstateresponseMap["messages"].([]interface{}); ok {
		MessagesString, _ := json.Marshal(Messages)
		json.Unmarshal(MessagesString, &o.Messages)
	}
	
	if SelfUri, ok := ArchitectjobstateresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectjobstateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
