package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Roledivisionpair
type Roledivisionpair struct { 
	// RoleId - The ID of the role
	RoleId *string `json:"roleId,omitempty"`


	// DivisionId - The ID of the division
	DivisionId *string `json:"divisionId,omitempty"`

}

func (o *Roledivisionpair) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Roledivisionpair
	
	return json.Marshal(&struct { 
		RoleId *string `json:"roleId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		*Alias
	}{ 
		RoleId: o.RoleId,
		
		DivisionId: o.DivisionId,
		Alias:    (*Alias)(o),
	})
}

func (o *Roledivisionpair) UnmarshalJSON(b []byte) error {
	var RoledivisionpairMap map[string]interface{}
	err := json.Unmarshal(b, &RoledivisionpairMap)
	if err != nil {
		return err
	}
	
	if RoleId, ok := RoledivisionpairMap["roleId"].(string); ok {
		o.RoleId = &RoleId
	}
    
	if DivisionId, ok := RoledivisionpairMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Roledivisionpair) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
