package platformclientv2
import (
	"encoding/json"
)

// Groupmembersupdate
type Groupmembersupdate struct { 
	// MemberIds - A list of the ids of the members to add.
	MemberIds *[]string `json:"memberIds,omitempty"`


	// Version - The current group version.
	Version *int32 `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Groupmembersupdate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
