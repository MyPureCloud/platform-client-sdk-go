package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Locationcreatedefinition
type Locationcreatedefinition struct { 
	// Name - The name of the Location. Required for creates, not required for updates
	Name *string `json:"name,omitempty"`


	// Version - Current version of the location
	Version *int `json:"version,omitempty"`


	// State - Current activity status of the location.
	State *string `json:"state,omitempty"`


	// Path - A list of ancestor ids
	Path *[]string `json:"path,omitempty"`


	// Notes - Notes for the location
	Notes *string `json:"notes,omitempty"`


	// ContactUser - The user id of the location contact
	ContactUser *string `json:"contactUser,omitempty"`


	// EmergencyNumber - Emergency number for the location
	EmergencyNumber *Locationemergencynumber `json:"emergencyNumber,omitempty"`


	// Address - Address of the location
	Address *Locationaddress `json:"address,omitempty"`

}

func (o *Locationcreatedefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Locationcreatedefinition
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Path *[]string `json:"path,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		ContactUser *string `json:"contactUser,omitempty"`
		
		EmergencyNumber *Locationemergencynumber `json:"emergencyNumber,omitempty"`
		
		Address *Locationaddress `json:"address,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Version: o.Version,
		
		State: o.State,
		
		Path: o.Path,
		
		Notes: o.Notes,
		
		ContactUser: o.ContactUser,
		
		EmergencyNumber: o.EmergencyNumber,
		
		Address: o.Address,
		Alias:    (*Alias)(o),
	})
}

func (o *Locationcreatedefinition) UnmarshalJSON(b []byte) error {
	var LocationcreatedefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &LocationcreatedefinitionMap)
	if err != nil {
		return err
	}
	
	if Name, ok := LocationcreatedefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Version, ok := LocationcreatedefinitionMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if State, ok := LocationcreatedefinitionMap["state"].(string); ok {
		o.State = &State
	}
    
	if Path, ok := LocationcreatedefinitionMap["path"].([]interface{}); ok {
		PathString, _ := json.Marshal(Path)
		json.Unmarshal(PathString, &o.Path)
	}
	
	if Notes, ok := LocationcreatedefinitionMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if ContactUser, ok := LocationcreatedefinitionMap["contactUser"].(string); ok {
		o.ContactUser = &ContactUser
	}
    
	if EmergencyNumber, ok := LocationcreatedefinitionMap["emergencyNumber"].(map[string]interface{}); ok {
		EmergencyNumberString, _ := json.Marshal(EmergencyNumber)
		json.Unmarshal(EmergencyNumberString, &o.EmergencyNumber)
	}
	
	if Address, ok := LocationcreatedefinitionMap["address"].(map[string]interface{}); ok {
		AddressString, _ := json.Marshal(Address)
		json.Unmarshal(AddressString, &o.Address)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Locationcreatedefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
