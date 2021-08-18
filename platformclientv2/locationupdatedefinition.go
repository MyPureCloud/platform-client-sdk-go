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

func (u *Locationupdatedefinition) MarshalJSON() ([]byte, error) {
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
		Name: u.Name,
		
		Version: u.Version,
		
		State: u.State,
		
		Path: u.Path,
		
		Notes: u.Notes,
		
		ContactUser: u.ContactUser,
		
		EmergencyNumber: u.EmergencyNumber,
		
		Address: u.Address,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Locationupdatedefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
