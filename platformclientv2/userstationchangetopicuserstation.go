package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userstationchangetopicuserstation
type Userstationchangetopicuserstation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// AssociatedUser
	AssociatedUser *Userstationchangetopicuser `json:"associatedUser,omitempty"`

}

func (o *Userstationchangetopicuserstation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userstationchangetopicuserstation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		AssociatedUser *Userstationchangetopicuser `json:"associatedUser,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		AssociatedUser: o.AssociatedUser,
		Alias:    (*Alias)(o),
	})
}

func (o *Userstationchangetopicuserstation) UnmarshalJSON(b []byte) error {
	var UserstationchangetopicuserstationMap map[string]interface{}
	err := json.Unmarshal(b, &UserstationchangetopicuserstationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserstationchangetopicuserstationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UserstationchangetopicuserstationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if AssociatedUser, ok := UserstationchangetopicuserstationMap["associatedUser"].(map[string]interface{}); ok {
		AssociatedUserString, _ := json.Marshal(AssociatedUser)
		json.Unmarshal(AssociatedUserString, &o.AssociatedUser)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userstationchangetopicuserstation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
