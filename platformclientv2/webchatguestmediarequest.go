package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatguestmediarequest - Object representing the guest model of a media request of a chat conversation.
type Webchatguestmediarequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Types - The types of media being requested.
	Types *[]string `json:"types,omitempty"`


	// State - The state of the media request, one of PENDING|ACCEPTED|DECLINED|TIMEDOUT|CANCELLED|ERRORED.
	State *string `json:"state,omitempty"`


	// CommunicationId - The ID of the new media communication, if applicable.
	CommunicationId *string `json:"communicationId,omitempty"`


	// SecurityKey - The security information related to a media request.
	SecurityKey *string `json:"securityKey,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Webchatguestmediarequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatguestmediarequest

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Types *[]string `json:"types,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		CommunicationId *string `json:"communicationId,omitempty"`
		
		SecurityKey *string `json:"securityKey,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Types: u.Types,
		
		State: u.State,
		
		CommunicationId: u.CommunicationId,
		
		SecurityKey: u.SecurityKey,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Webchatguestmediarequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
