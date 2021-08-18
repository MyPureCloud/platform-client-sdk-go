package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createsharerequest
type Createsharerequest struct { 
	// SharedEntityType - The share entity type
	SharedEntityType *string `json:"sharedEntityType,omitempty"`


	// SharedEntity - The entity that will be shared
	SharedEntity *Sharedentity `json:"sharedEntity,omitempty"`


	// MemberType
	MemberType *string `json:"memberType,omitempty"`


	// Member - The member that will have access to this share. Only required if a list of members is not provided.
	Member *Sharedentity `json:"member,omitempty"`


	// Members
	Members *[]Createsharerequestmember `json:"members,omitempty"`

}

func (u *Createsharerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createsharerequest

	

	return json.Marshal(&struct { 
		SharedEntityType *string `json:"sharedEntityType,omitempty"`
		
		SharedEntity *Sharedentity `json:"sharedEntity,omitempty"`
		
		MemberType *string `json:"memberType,omitempty"`
		
		Member *Sharedentity `json:"member,omitempty"`
		
		Members *[]Createsharerequestmember `json:"members,omitempty"`
		*Alias
	}{ 
		SharedEntityType: u.SharedEntityType,
		
		SharedEntity: u.SharedEntity,
		
		MemberType: u.MemberType,
		
		Member: u.Member,
		
		Members: u.Members,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createsharerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
