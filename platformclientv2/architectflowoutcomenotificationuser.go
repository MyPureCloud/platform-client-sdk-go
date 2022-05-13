package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflowoutcomenotificationuser - The user who initiated the change.
type Architectflowoutcomenotificationuser struct { 
	// Id - The ID of the user.
	Id *string `json:"id,omitempty"`


	// Name - The name of the user, if available.
	Name *string `json:"name,omitempty"`


	// HomeOrg
	HomeOrg *Architectflowoutcomenotificationhomeorganization `json:"homeOrg,omitempty"`

}

func (o *Architectflowoutcomenotificationuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflowoutcomenotificationuser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		HomeOrg *Architectflowoutcomenotificationhomeorganization `json:"homeOrg,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		HomeOrg: o.HomeOrg,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectflowoutcomenotificationuser) UnmarshalJSON(b []byte) error {
	var ArchitectflowoutcomenotificationuserMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflowoutcomenotificationuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectflowoutcomenotificationuserMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ArchitectflowoutcomenotificationuserMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if HomeOrg, ok := ArchitectflowoutcomenotificationuserMap["homeOrg"].(map[string]interface{}); ok {
		HomeOrgString, _ := json.Marshal(HomeOrg)
		json.Unmarshal(HomeOrgString, &o.HomeOrg)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
