package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflownotificationuser
type Architectflownotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// HomeOrg
	HomeOrg *Architectflownotificationhomeorganization `json:"homeOrg,omitempty"`

}

func (o *Architectflownotificationuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflownotificationuser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		HomeOrg *Architectflownotificationhomeorganization `json:"homeOrg,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		HomeOrg: o.HomeOrg,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectflownotificationuser) UnmarshalJSON(b []byte) error {
	var ArchitectflownotificationuserMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflownotificationuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectflownotificationuserMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ArchitectflownotificationuserMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if HomeOrg, ok := ArchitectflownotificationuserMap["homeOrg"].(map[string]interface{}); ok {
		HomeOrgString, _ := json.Marshal(HomeOrg)
		json.Unmarshal(HomeOrgString, &o.HomeOrg)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflownotificationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
