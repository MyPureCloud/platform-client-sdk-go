package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Skillgroupmemberdivisions
type Skillgroupmemberdivisions struct { 
	// AddDivisionIds
	AddDivisionIds *[]string `json:"addDivisionIds,omitempty"`


	// RemoveDivisionIds
	RemoveDivisionIds *[]string `json:"removeDivisionIds,omitempty"`

}

func (o *Skillgroupmemberdivisions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Skillgroupmemberdivisions
	
	return json.Marshal(&struct { 
		AddDivisionIds *[]string `json:"addDivisionIds,omitempty"`
		
		RemoveDivisionIds *[]string `json:"removeDivisionIds,omitempty"`
		*Alias
	}{ 
		AddDivisionIds: o.AddDivisionIds,
		
		RemoveDivisionIds: o.RemoveDivisionIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Skillgroupmemberdivisions) UnmarshalJSON(b []byte) error {
	var SkillgroupmemberdivisionsMap map[string]interface{}
	err := json.Unmarshal(b, &SkillgroupmemberdivisionsMap)
	if err != nil {
		return err
	}
	
	if AddDivisionIds, ok := SkillgroupmemberdivisionsMap["addDivisionIds"].([]interface{}); ok {
		AddDivisionIdsString, _ := json.Marshal(AddDivisionIds)
		json.Unmarshal(AddDivisionIdsString, &o.AddDivisionIds)
	}
	
	if RemoveDivisionIds, ok := SkillgroupmemberdivisionsMap["removeDivisionIds"].([]interface{}); ok {
		RemoveDivisionIdsString, _ := json.Marshal(RemoveDivisionIds)
		json.Unmarshal(RemoveDivisionIdsString, &o.RemoveDivisionIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Skillgroupmemberdivisions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
