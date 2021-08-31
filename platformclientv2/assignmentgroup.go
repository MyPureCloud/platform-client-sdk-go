package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assignmentgroup
type Assignmentgroup struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *Assignmentgroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assignmentgroup
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SelfUri: o.SelfUri,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Assignmentgroup) UnmarshalJSON(b []byte) error {
	var AssignmentgroupMap map[string]interface{}
	err := json.Unmarshal(b, &AssignmentgroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AssignmentgroupMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := AssignmentgroupMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if SelfUri, ok := AssignmentgroupMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if VarType, ok := AssignmentgroupMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Assignmentgroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
