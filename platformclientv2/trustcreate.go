package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Trustcreate
type Trustcreate struct { 
	// PairingId - The pairing Id created by the trustee. This is required to prove that the trustee agrees to the relationship.
	PairingId *string `json:"pairingId,omitempty"`


	// Enabled - If disabled no trustee user will have access, even if they were previously added.
	Enabled *bool `json:"enabled,omitempty"`


	// Users - The list of users and their roles to which access will be granted. The users are from the trustee and the roles are from the trustor. If no users are specified, at least one group is required.
	Users *[]Trustmembercreate `json:"users,omitempty"`


	// Groups - The list of groups and their roles to which access will be granted. The groups are from the trustee and the roles are from the trustor. If no groups are specified, at least one user is required.
	Groups *[]Trustmembercreate `json:"groups,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trustcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
