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

func (o *Usertokenstopictokennotification) MarshalJSON() ([]byte, error) {
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
		User: o.User,
		
		IpAddress: o.IpAddress,
		
		DateCreated: o.DateCreated,
		
		TokenExpirationDate: o.TokenExpirationDate,
		
		SessionId: o.SessionId,
		
		ClientId: o.ClientId,
		
		TokenHash: o.TokenHash,
		Alias:    (*Alias)(o),
	})
}

func (o *Usertokenstopictokennotification) UnmarshalJSON(b []byte) error {
	var UsertokenstopictokennotificationMap map[string]interface{}
	err := json.Unmarshal(b, &UsertokenstopictokennotificationMap)
	if err != nil {
		return err
	}
	
	if User, ok := UsertokenstopictokennotificationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if IpAddress, ok := UsertokenstopictokennotificationMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
	
	if DateCreated, ok := UsertokenstopictokennotificationMap["dateCreated"].(string); ok {
		o.DateCreated = &DateCreated
	}
	
	if TokenExpirationDate, ok := UsertokenstopictokennotificationMap["tokenExpirationDate"].(string); ok {
		o.TokenExpirationDate = &TokenExpirationDate
	}
	
	if SessionId, ok := UsertokenstopictokennotificationMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
	
	if ClientId, ok := UsertokenstopictokennotificationMap["clientId"].(string); ok {
		o.ClientId = &ClientId
	}
	
	if TokenHash, ok := UsertokenstopictokennotificationMap["tokenHash"].(string); ok {
		o.TokenHash = &TokenHash
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Usertokenstopictokennotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
