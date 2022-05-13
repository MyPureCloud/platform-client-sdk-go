package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectdependencytrackingbuildnotificationuser - The user who initiated the change.
type Architectdependencytrackingbuildnotificationuser struct { 
	// Id - The ID of the user.
	Id *string `json:"id,omitempty"`


	// Name - The name of the user, if available.
	Name *string `json:"name,omitempty"`


	// HomeOrg
	HomeOrg *Architectdependencytrackingbuildnotificationhomeorganization `json:"homeOrg,omitempty"`

}

func (o *Architectdependencytrackingbuildnotificationuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectdependencytrackingbuildnotificationuser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		HomeOrg *Architectdependencytrackingbuildnotificationhomeorganization `json:"homeOrg,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		HomeOrg: o.HomeOrg,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectdependencytrackingbuildnotificationuser) UnmarshalJSON(b []byte) error {
	var ArchitectdependencytrackingbuildnotificationuserMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectdependencytrackingbuildnotificationuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectdependencytrackingbuildnotificationuserMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ArchitectdependencytrackingbuildnotificationuserMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if HomeOrg, ok := ArchitectdependencytrackingbuildnotificationuserMap["homeOrg"].(map[string]interface{}); ok {
		HomeOrgString, _ := json.Marshal(HomeOrg)
		json.Unmarshal(HomeOrgString, &o.HomeOrg)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectdependencytrackingbuildnotificationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
