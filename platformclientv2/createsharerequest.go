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

func (o *Createsharerequest) MarshalJSON() ([]byte, error) {
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
		SharedEntityType: o.SharedEntityType,
		
		SharedEntity: o.SharedEntity,
		
		MemberType: o.MemberType,
		
		Member: o.Member,
		
		Members: o.Members,
		Alias:    (*Alias)(o),
	})
}

func (o *Createsharerequest) UnmarshalJSON(b []byte) error {
	var CreatesharerequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatesharerequestMap)
	if err != nil {
		return err
	}
	
	if SharedEntityType, ok := CreatesharerequestMap["sharedEntityType"].(string); ok {
		o.SharedEntityType = &SharedEntityType
	}
	
	if SharedEntity, ok := CreatesharerequestMap["sharedEntity"].(map[string]interface{}); ok {
		SharedEntityString, _ := json.Marshal(SharedEntity)
		json.Unmarshal(SharedEntityString, &o.SharedEntity)
	}
	
	if MemberType, ok := CreatesharerequestMap["memberType"].(string); ok {
		o.MemberType = &MemberType
	}
	
	if Member, ok := CreatesharerequestMap["member"].(map[string]interface{}); ok {
		MemberString, _ := json.Marshal(Member)
		json.Unmarshal(MemberString, &o.Member)
	}
	
	if Members, ok := CreatesharerequestMap["members"].([]interface{}); ok {
		MembersString, _ := json.Marshal(Members)
		json.Unmarshal(MembersString, &o.Members)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createsharerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
