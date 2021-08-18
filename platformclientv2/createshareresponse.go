package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createshareresponse
type Createshareresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SharedEntityType
	SharedEntityType *string `json:"sharedEntityType,omitempty"`


	// SharedEntity
	SharedEntity *Domainentityref `json:"sharedEntity,omitempty"`


	// MemberType
	MemberType *string `json:"memberType,omitempty"`


	// Member
	Member *Domainentityref `json:"member,omitempty"`


	// SharedBy
	SharedBy *Domainentityref `json:"sharedBy,omitempty"`


	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// Succeeded
	Succeeded *[]Share `json:"succeeded,omitempty"`


	// Failed
	Failed *[]Share `json:"failed,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Createshareresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createshareresponse

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SharedEntityType *string `json:"sharedEntityType,omitempty"`
		
		SharedEntity *Domainentityref `json:"sharedEntity,omitempty"`
		
		MemberType *string `json:"memberType,omitempty"`
		
		Member *Domainentityref `json:"member,omitempty"`
		
		SharedBy *Domainentityref `json:"sharedBy,omitempty"`
		
		Workspace *Domainentityref `json:"workspace,omitempty"`
		
		Succeeded *[]Share `json:"succeeded,omitempty"`
		
		Failed *[]Share `json:"failed,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		SharedEntityType: u.SharedEntityType,
		
		SharedEntity: u.SharedEntity,
		
		MemberType: u.MemberType,
		
		Member: u.Member,
		
		SharedBy: u.SharedBy,
		
		Workspace: u.Workspace,
		
		Succeeded: u.Succeeded,
		
		Failed: u.Failed,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createshareresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
