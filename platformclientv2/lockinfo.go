package platformclientv2
import (
	"time"
	"encoding/json"
)

// Lockinfo
type Lockinfo struct { 
	// LockedBy
	LockedBy *Domainentityref `json:"lockedBy,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateExpires - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateExpires *time.Time `json:"dateExpires,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

// String returns a JSON representation of the model
func (o *Lockinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
