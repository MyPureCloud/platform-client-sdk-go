package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Usertokenstopictokennotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Usertokenstopictokennotification

	

	return json.Marshal(&struct { 
		User *Usertokenstopicurireference `json:"user,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		TokenExpirationDate *string `json:"tokenExpirationDate,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		ClientId *string `json:"clientId,omitempty"`
		
		TokenHash *string `json:"tokenHash,omitempty"`
		*Alias
	}{ 
		User: u.User,
		
		IpAddress: u.IpAddress,
		
		DateCreated: u.DateCreated,
		
		TokenExpirationDate: u.TokenExpirationDate,
		
		SessionId: u.SessionId,
		
		ClientId: u.ClientId,
		
		TokenHash: u.TokenHash,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Usertokenstopictokennotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
