package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Webchatguestmediarequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
