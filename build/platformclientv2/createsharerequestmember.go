package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Createsharerequestmember
type Createsharerequestmember struct { 
	// MemberType
	MemberType *string `json:"memberType,omitempty"`


	// Member
	Member *Memberentity `json:"member,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createsharerequestmember) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
