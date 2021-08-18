package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workspacemember
type Workspacemember struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// MemberType - The workspace member type.
	MemberType *string `json:"memberType,omitempty"`


	// Member
	Member *Domainentityref `json:"member,omitempty"`


	// User
	User *User `json:"user,omitempty"`


	// Group
	Group *Group `json:"group,omitempty"`


	// SecurityProfile
	SecurityProfile *Securityprofile `json:"securityProfile,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Workspacemember) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workspacemember

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Workspace *Domainentityref `json:"workspace,omitempty"`
		
		MemberType *string `json:"memberType,omitempty"`
		
		Member *Domainentityref `json:"member,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Group *Group `json:"group,omitempty"`
		
		SecurityProfile *Securityprofile `json:"securityProfile,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Workspace: u.Workspace,
		
		MemberType: u.MemberType,
		
		Member: u.Member,
		
		User: u.User,
		
		Group: u.Group,
		
		SecurityProfile: u.SecurityProfile,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Workspacemember) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
