package platformclientv2
import (
	"encoding/json"
)

// Queuemember
type Queuemember struct { 
	// Id - The queue member's id.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// User
	User *User `json:"user,omitempty"`


	// RingNumber
	RingNumber *int `json:"ringNumber,omitempty"`


	// Joined
	Joined *bool `json:"joined,omitempty"`


	// MemberBy
	MemberBy *string `json:"memberBy,omitempty"`


	// RoutingStatus
	RoutingStatus *Routingstatus `json:"routingStatus,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queuemember) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
