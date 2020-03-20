package platformclientv2
import (
	"time"
	"encoding/json"
)

// Defaultgreetinglist
type Defaultgreetinglist struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Owner
	Owner *Greetingowner `json:"owner,omitempty"`


	// OwnerType
	OwnerType *string `json:"ownerType,omitempty"`


	// Greetings
	Greetings *map[string]Greeting `json:"greetings,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// CreatedBy
	CreatedBy *string `json:"createdBy,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// ModifiedBy
	ModifiedBy *string `json:"modifiedBy,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Defaultgreetinglist) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
