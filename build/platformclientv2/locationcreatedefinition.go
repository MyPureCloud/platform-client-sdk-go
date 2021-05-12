package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Locationcreatedefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
