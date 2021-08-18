package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Groupmembersupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Groupmembersupdate

	

	return json.Marshal(&struct { 
		MemberIds *[]string `json:"memberIds,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		MemberIds: u.MemberIds,
		
		Version: u.Version,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Groupmembersupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
