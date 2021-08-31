package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkbaseassignment
type Trunkbaseassignment struct { 
	// Family - The address family to use with the trunk base settings. 2=IPv4, 23=IPv6
	Family *int `json:"family,omitempty"`


	// TrunkBase - A trunk base settings reference.
	TrunkBase *Trunkbase `json:"trunkBase,omitempty"`

}

func (o *Trunkbaseassignment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkbaseassignment
	
	return json.Marshal(&struct { 
		Family *int `json:"family,omitempty"`
		
		TrunkBase *Trunkbase `json:"trunkBase,omitempty"`
		*Alias
	}{ 
		Family: o.Family,
		
		TrunkBase: o.TrunkBase,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkbaseassignment) UnmarshalJSON(b []byte) error {
	var TrunkbaseassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkbaseassignmentMap)
	if err != nil {
		return err
	}
	
	if Family, ok := TrunkbaseassignmentMap["family"].(float64); ok {
		FamilyInt := int(Family)
		o.Family = &FamilyInt
	}
	
	if TrunkBase, ok := TrunkbaseassignmentMap["trunkBase"].(map[string]interface{}); ok {
		TrunkBaseString, _ := json.Marshal(TrunkBase)
		json.Unmarshal(TrunkBaseString, &o.TrunkBase)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkbaseassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
