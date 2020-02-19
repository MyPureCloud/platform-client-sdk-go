package platformclientv2
import (
	"encoding/json"
)

// Queueusereventtopicqueuemember
type Queueusereventtopicqueuemember struct { 
	// MemberId
	MemberId *string `json:"memberId,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// RingNumber
	RingNumber *int32 `json:"ringNumber,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Joined
	Joined *bool `json:"joined,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueusereventtopicqueuemember) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
