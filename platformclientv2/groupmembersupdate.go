package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Groupmembersupdate
type Groupmembersupdate struct { 
	// MemberIds - A list of the ids of the members to add.
	MemberIds *[]string `json:"memberIds,omitempty"`


	// Version - The current group version.
	Version *int `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Groupmembersupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
