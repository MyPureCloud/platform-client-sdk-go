package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Subjectdivisiongrants
type Subjectdivisiongrants struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Divisions
	Divisions *[]Division `json:"divisions,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Subjectdivisiongrants) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Subjectdivisiongrants
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Divisions *[]Division `json:"divisions,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Divisions: o.Divisions,
		
		VarType: o.VarType,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Subjectdivisiongrants) UnmarshalJSON(b []byte) error {
	var SubjectdivisiongrantsMap map[string]interface{}
	err := json.Unmarshal(b, &SubjectdivisiongrantsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SubjectdivisiongrantsMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SubjectdivisiongrantsMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Divisions, ok := SubjectdivisiongrantsMap["divisions"].([]interface{}); ok {
		DivisionsString, _ := json.Marshal(Divisions)
		json.Unmarshal(DivisionsString, &o.Divisions)
	}
	
	if VarType, ok := SubjectdivisiongrantsMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if SelfUri, ok := SubjectdivisiongrantsMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Subjectdivisiongrants) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
