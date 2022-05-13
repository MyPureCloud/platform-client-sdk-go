package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Locationupdatedefinition
type Locationupdatedefinition struct { 
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

func (o *Locationupdatedefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Locationupdatedefinition
	
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

func (o *Locationupdatedefinition) UnmarshalJSON(b []byte) error {
	var LocationupdatedefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &LocationupdatedefinitionMap)
	if err != nil {
		return err
	}
	
	if Name, ok := LocationupdatedefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Version, ok := LocationupdatedefinitionMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if State, ok := LocationupdatedefinitionMap["state"].(string); ok {
		o.State = &State
	}
    
	if Path, ok := LocationupdatedefinitionMap["path"].([]interface{}); ok {
		PathString, _ := json.Marshal(Path)
		json.Unmarshal(PathString, &o.Path)
	}
	
	if Notes, ok := LocationupdatedefinitionMap["notes"].(string); ok {
		o.Notes = &Notes
	}
    
	if ContactUser, ok := LocationupdatedefinitionMap["contactUser"].(string); ok {
		o.ContactUser = &ContactUser
	}
    
	if EmergencyNumber, ok := LocationupdatedefinitionMap["emergencyNumber"].(map[string]interface{}); ok {
		EmergencyNumberString, _ := json.Marshal(EmergencyNumber)
		json.Unmarshal(EmergencyNumberString, &o.EmergencyNumber)
	}
	
	if Address, ok := LocationupdatedefinitionMap["address"].(map[string]interface{}); ok {
		AddressString, _ := json.Marshal(Address)
		json.Unmarshal(AddressString, &o.Address)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Locationupdatedefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
