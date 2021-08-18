package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Queuemember) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queuemember

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		RingNumber *int `json:"ringNumber,omitempty"`
		
		Joined *bool `json:"joined,omitempty"`
		
		MemberBy *string `json:"memberBy,omitempty"`
		
		RoutingStatus *Routingstatus `json:"routingStatus,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		User: u.User,
		
		RingNumber: u.RingNumber,
		
		Joined: u.Joined,
		
		MemberBy: u.MemberBy,
		
		RoutingStatus: u.RoutingStatus,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queuemember) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
