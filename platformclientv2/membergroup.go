package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Membergroup
type Membergroup struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`


	// VarType - The type of group, e.g. TEAM, etc.
	VarType *string `json:"type,omitempty"`


	// MemberCount - The number of members in this group
	MemberCount *int `json:"memberCount,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Membergroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Membergroup
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		MemberCount *int `json:"memberCount,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		VarType: o.VarType,
		
		MemberCount: o.MemberCount,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Membergroup) UnmarshalJSON(b []byte) error {
	var MembergroupMap map[string]interface{}
	err := json.Unmarshal(b, &MembergroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MembergroupMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := MembergroupMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := MembergroupMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if VarType, ok := MembergroupMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if MemberCount, ok := MembergroupMap["memberCount"].(float64); ok {
		MemberCountInt := int(MemberCount)
		o.MemberCount = &MemberCountInt
	}
	
	if SelfUri, ok := MembergroupMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Membergroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
