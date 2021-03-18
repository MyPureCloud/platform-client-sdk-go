package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Trustrequestcreate
type Trustrequestcreate struct { 
	// UserIds - The list of trustee users that are requesting access. If no users are specified, at least one group is required.
	UserIds *[]string `json:"userIds,omitempty"`


	// GroupIds - The list of trustee groups that are requesting access. If no groups are specified, at least one user is required.
	GroupIds *[]string `json:"groupIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trustrequestcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
