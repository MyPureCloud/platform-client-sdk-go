package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectdependencytrackingbuildnotificationhomeorganization
type Architectdependencytrackingbuildnotificationhomeorganization struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ThirdPartyOrgName
	ThirdPartyOrgName *string `json:"thirdPartyOrgName,omitempty"`

}

func (o *Architectdependencytrackingbuildnotificationhomeorganization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectdependencytrackingbuildnotificationhomeorganization
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ThirdPartyOrgName *string `json:"thirdPartyOrgName,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ThirdPartyOrgName: o.ThirdPartyOrgName,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectdependencytrackingbuildnotificationhomeorganization) UnmarshalJSON(b []byte) error {
	var ArchitectdependencytrackingbuildnotificationhomeorganizationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectdependencytrackingbuildnotificationhomeorganizationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectdependencytrackingbuildnotificationhomeorganizationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ArchitectdependencytrackingbuildnotificationhomeorganizationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ThirdPartyOrgName, ok := ArchitectdependencytrackingbuildnotificationhomeorganizationMap["thirdPartyOrgName"].(string); ok {
		o.ThirdPartyOrgName = &ThirdPartyOrgName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Architectdependencytrackingbuildnotificationhomeorganization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
