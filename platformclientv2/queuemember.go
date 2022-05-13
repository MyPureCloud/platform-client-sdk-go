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

func (o *Queuemember) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		User: o.User,
		
		RingNumber: o.RingNumber,
		
		Joined: o.Joined,
		
		MemberBy: o.MemberBy,
		
		RoutingStatus: o.RoutingStatus,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Queuemember) UnmarshalJSON(b []byte) error {
	var QueuememberMap map[string]interface{}
	err := json.Unmarshal(b, &QueuememberMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueuememberMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := QueuememberMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if User, ok := QueuememberMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if RingNumber, ok := QueuememberMap["ringNumber"].(float64); ok {
		RingNumberInt := int(RingNumber)
		o.RingNumber = &RingNumberInt
	}
	
	if Joined, ok := QueuememberMap["joined"].(bool); ok {
		o.Joined = &Joined
	}
    
	if MemberBy, ok := QueuememberMap["memberBy"].(string); ok {
		o.MemberBy = &MemberBy
	}
    
	if RoutingStatus, ok := QueuememberMap["routingStatus"].(map[string]interface{}); ok {
		RoutingStatusString, _ := json.Marshal(RoutingStatus)
		json.Unmarshal(RoutingStatusString, &o.RoutingStatus)
	}
	
	if SelfUri, ok := QueuememberMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queuemember) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
