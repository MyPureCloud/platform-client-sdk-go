package platformclientv2
import (
	"encoding/json"
)

// Updatemanagementunitrequest
type Updatemanagementunitrequest struct { 
	// Name - The new name of the management unit
	Name *string `json:"name,omitempty"`


	// DivisionId - The new division id for the management unit
	DivisionId *string `json:"divisionId,omitempty"`


	// Settings - Updated settings for the management unit
	Settings *Managementunitsettingsrequest `json:"settings,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updatemanagementunitrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
