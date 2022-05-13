package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Securesession
type Securesession struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Flow - The flow to execute securely
	Flow *Domainentityref `json:"flow,omitempty"`


	// UserData - Customer-provided data
	UserData *string `json:"userData,omitempty"`


	// State - The current state of a secure session
	State *string `json:"state,omitempty"`


	// SourceParticipantId - Unique identifier for the participant initiating the secure session.
	SourceParticipantId *string `json:"sourceParticipantId,omitempty"`


	// Disconnect - If true, disconnect the agent after creating the session
	Disconnect *bool `json:"disconnect,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Securesession) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Securesession
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Flow *Domainentityref `json:"flow,omitempty"`
		
		UserData *string `json:"userData,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		SourceParticipantId *string `json:"sourceParticipantId,omitempty"`
		
		Disconnect *bool `json:"disconnect,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Flow: o.Flow,
		
		UserData: o.UserData,
		
		State: o.State,
		
		SourceParticipantId: o.SourceParticipantId,
		
		Disconnect: o.Disconnect,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Securesession) UnmarshalJSON(b []byte) error {
	var SecuresessionMap map[string]interface{}
	err := json.Unmarshal(b, &SecuresessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SecuresessionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Flow, ok := SecuresessionMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if UserData, ok := SecuresessionMap["userData"].(string); ok {
		o.UserData = &UserData
	}
    
	if State, ok := SecuresessionMap["state"].(string); ok {
		o.State = &State
	}
    
	if SourceParticipantId, ok := SecuresessionMap["sourceParticipantId"].(string); ok {
		o.SourceParticipantId = &SourceParticipantId
	}
    
	if Disconnect, ok := SecuresessionMap["disconnect"].(bool); ok {
		o.Disconnect = &Disconnect
	}
    
	if SelfUri, ok := SecuresessionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Securesession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
