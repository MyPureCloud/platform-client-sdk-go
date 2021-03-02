package platformclientv2
import (
	"encoding/json"
)

// Usertokenstopictokennotification
type Usertokenstopictokennotification struct { 
	// User
	User *Usertokenstopicurireference `json:"user,omitempty"`


	// IpAddress
	IpAddress *string `json:"ipAddress,omitempty"`


	// DateCreated
	DateCreated *string `json:"dateCreated,omitempty"`


	// TokenExpirationDate
	TokenExpirationDate *string `json:"tokenExpirationDate,omitempty"`


	// SessionId
	SessionId *string `json:"sessionId,omitempty"`


	// ClientId
	ClientId *string `json:"clientId,omitempty"`


	// TokenHash
	TokenHash *string `json:"tokenHash,omitempty"`

}

// String returns a JSON representation of the model
func (o *Usertokenstopictokennotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
