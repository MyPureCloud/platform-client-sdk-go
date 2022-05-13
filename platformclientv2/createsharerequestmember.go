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

func (o *Createsharerequestmember) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createsharerequestmember
	
	return json.Marshal(&struct { 
		MemberType *string `json:"memberType,omitempty"`
		
		Member *Memberentity `json:"member,omitempty"`
		*Alias
	}{ 
		MemberType: o.MemberType,
		
		Member: o.Member,
		Alias:    (*Alias)(o),
	})
}

func (o *Createsharerequestmember) UnmarshalJSON(b []byte) error {
	var CreatesharerequestmemberMap map[string]interface{}
	err := json.Unmarshal(b, &CreatesharerequestmemberMap)
	if err != nil {
		return err
	}
	
	if MemberType, ok := CreatesharerequestmemberMap["memberType"].(string); ok {
		o.MemberType = &MemberType
	}
    
	if Member, ok := CreatesharerequestmemberMap["member"].(map[string]interface{}); ok {
		MemberString, _ := json.Marshal(Member)
		json.Unmarshal(MemberString, &o.Member)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createsharerequestmember) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
