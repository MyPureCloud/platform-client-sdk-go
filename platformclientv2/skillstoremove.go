package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Skillstoremove
type Skillstoremove struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Skillstoremove) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Skillstoremove
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Skillstoremove) UnmarshalJSON(b []byte) error {
	var SkillstoremoveMap map[string]interface{}
	err := json.Unmarshal(b, &SkillstoremoveMap)
	if err != nil {
		return err
	}
	
	if Name, ok := SkillstoremoveMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Id, ok := SkillstoremoveMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SelfUri, ok := SkillstoremoveMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Skillstoremove) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
