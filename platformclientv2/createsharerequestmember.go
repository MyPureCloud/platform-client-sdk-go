package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Createsharerequestmember) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createsharerequestmember

	

	return json.Marshal(&struct { 
		MemberType *string `json:"memberType,omitempty"`
		
		Member *Memberentity `json:"member,omitempty"`
		*Alias
	}{ 
		MemberType: u.MemberType,
		
		Member: u.Member,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createsharerequestmember) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
