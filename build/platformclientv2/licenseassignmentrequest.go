package platformclientv2
import (
	"encoding/json"
)

// Licenseassignmentrequest
type Licenseassignmentrequest struct { 
	// LicenseId - The id of the license to assign/unassign.
	LicenseId *string `json:"licenseId,omitempty"`


	// UserIdsAdd - The ids of users to assign this license to.
	UserIdsAdd *[]string `json:"userIdsAdd,omitempty"`


	// UserIdsRemove - The ids of users to unassign this license from.
	UserIdsRemove *[]string `json:"userIdsRemove,omitempty"`

}

// String returns a JSON representation of the model
func (o *Licenseassignmentrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
