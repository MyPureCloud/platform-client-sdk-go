package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createsecuresession
type Createsecuresession struct { 
	// SourceParticipantId - requesting participant
	SourceParticipantId *string `json:"sourceParticipantId,omitempty"`


	// FlowId - the flow id to execute in the secure session
	FlowId *string `json:"flowId,omitempty"`


	// UserData - user data for the secure session
	UserData *string `json:"userData,omitempty"`


	// Disconnect - if true, disconnect the agent after creating the session
	Disconnect *bool `json:"disconnect,omitempty"`

}

func (o *Createsecuresession) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createsecuresession
	
	return json.Marshal(&struct { 
		SourceParticipantId *string `json:"sourceParticipantId,omitempty"`
		
		FlowId *string `json:"flowId,omitempty"`
		
		UserData *string `json:"userData,omitempty"`
		
		Disconnect *bool `json:"disconnect,omitempty"`
		*Alias
	}{ 
		SourceParticipantId: o.SourceParticipantId,
		
		FlowId: o.FlowId,
		
		UserData: o.UserData,
		
		Disconnect: o.Disconnect,
		Alias:    (*Alias)(o),
	})
}

func (o *Createsecuresession) UnmarshalJSON(b []byte) error {
	var CreatesecuresessionMap map[string]interface{}
	err := json.Unmarshal(b, &CreatesecuresessionMap)
	if err != nil {
		return err
	}
	
	if SourceParticipantId, ok := CreatesecuresessionMap["sourceParticipantId"].(string); ok {
		o.SourceParticipantId = &SourceParticipantId
	}
    
	if FlowId, ok := CreatesecuresessionMap["flowId"].(string); ok {
		o.FlowId = &FlowId
	}
    
	if UserData, ok := CreatesecuresessionMap["userData"].(string); ok {
		o.UserData = &UserData
	}
    
	if Disconnect, ok := CreatesecuresessionMap["disconnect"].(bool); ok {
		o.Disconnect = &Disconnect
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createsecuresession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
