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

func (u *Securesession) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Flow: u.Flow,
		
		UserData: u.UserData,
		
		State: u.State,
		
		SourceParticipantId: u.SourceParticipantId,
		
		Disconnect: u.Disconnect,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Securesession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
