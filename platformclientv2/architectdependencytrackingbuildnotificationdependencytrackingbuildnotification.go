package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
