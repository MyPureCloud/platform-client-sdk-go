package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectpromptnotificationuser
type Architectpromptnotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// HomeOrg
	HomeOrg *Architectpromptnotificationhomeorganization `json:"homeOrg,omitempty"`

}

func (o *Architectpromptnotificationuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectpromptnotificationuser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		HomeOrg *Architectpromptnotificationhomeorganization `json:"homeOrg,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		HomeOrg: o.HomeOrg,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectpromptnotificationuser) UnmarshalJSON(b []byte) error {
	var ArchitectpromptnotificationuserMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectpromptnotificationuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectpromptnotificationuserMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ArchitectpromptnotificationuserMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if HomeOrg, ok := ArchitectpromptnotificationuserMap["homeOrg"].(map[string]interface{}); ok {
		HomeOrgString, _ := json.Marshal(HomeOrg)
		json.Unmarshal(HomeOrgString, &o.HomeOrg)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectpromptnotificationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
