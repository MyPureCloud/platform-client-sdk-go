package platformclientv2
import (
	"time"
	"encoding/json"
)

// Commandstatus
type Commandstatus struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Expiration - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Expiration *time.Time `json:"expiration,omitempty"`


	// UserId
	UserId *string `json:"userId,omitempty"`


	// StatusCode
	StatusCode *string `json:"statusCode,omitempty"`


	// CommandType
	CommandType *string `json:"commandType,omitempty"`


	// Document
	Document *Document `json:"document,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Commandstatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
