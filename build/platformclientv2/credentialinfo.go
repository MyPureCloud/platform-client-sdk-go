package platformclientv2
import (
	"time"
	"encoding/json"
)

// Credentialinfo
type Credentialinfo struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// CreatedDate - Date the credentials were created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - Date credentials were last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// VarType - Type of the credentials.
	VarType *Credentialtype `json:"type,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Credentialinfo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
