package platformclientv2
import (
	"time"
	"encoding/json"
)

// Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification
type Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// User
	User *Architectdependencytrackingbuildnotificationuser `json:"user,omitempty"`


	// Client
	Client *Architectdependencytrackingbuildnotificationclient `json:"client,omitempty"`


	// StartTime
	StartTime *time.Time `json:"startTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
