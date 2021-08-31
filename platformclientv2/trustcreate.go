package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trustcreate
type Trustcreate struct { 
	// PairingId - The pairing Id created by the trustee. This is required to prove that the trustee agrees to the relationship.  Not required when creating a default pairing with Customer Care.
	PairingId *string `json:"pairingId,omitempty"`


	// Enabled - If disabled no trustee user will have access, even if they were previously added.
	Enabled *bool `json:"enabled,omitempty"`


	// Users - The list of users and their roles to which access will be granted. The users are from the trustee and the roles are from the trustor. If no users are specified, at least one group is required.
	Users *[]Trustmembercreate `json:"users,omitempty"`


	// Groups - The list of groups and their roles to which access will be granted. The groups are from the trustee and the roles are from the trustor. If no groups are specified, at least one user is required.
	Groups *[]Trustmembercreate `json:"groups,omitempty"`


	// DateExpired - The expiration date of the trust. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpired *time.Time `json:"dateExpired,omitempty"`

}

func (o *Trustcreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trustcreate
	
	DateExpired := new(string)
	if o.DateExpired != nil {
		
		*DateExpired = timeutil.Strftime(o.DateExpired, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpired = nil
	}
	
	return json.Marshal(&struct { 
		PairingId *string `json:"pairingId,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Users *[]Trustmembercreate `json:"users,omitempty"`
		
		Groups *[]Trustmembercreate `json:"groups,omitempty"`
		
		DateExpired *string `json:"dateExpired,omitempty"`
		*Alias
	}{ 
		PairingId: o.PairingId,
		
		Enabled: o.Enabled,
		
		Users: o.Users,
		
		Groups: o.Groups,
		
		DateExpired: DateExpired,
		Alias:    (*Alias)(o),
	})
}

func (o *Trustcreate) UnmarshalJSON(b []byte) error {
	var TrustcreateMap map[string]interface{}
	err := json.Unmarshal(b, &TrustcreateMap)
	if err != nil {
		return err
	}
	
	if PairingId, ok := TrustcreateMap["pairingId"].(string); ok {
		o.PairingId = &PairingId
	}
	
	if Enabled, ok := TrustcreateMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if Users, ok := TrustcreateMap["users"].([]interface{}); ok {
		UsersString, _ := json.Marshal(Users)
		json.Unmarshal(UsersString, &o.Users)
	}
	
	if Groups, ok := TrustcreateMap["groups"].([]interface{}); ok {
		GroupsString, _ := json.Marshal(Groups)
		json.Unmarshal(GroupsString, &o.Groups)
	}
	
	if dateExpiredString, ok := TrustcreateMap["dateExpired"].(string); ok {
		DateExpired, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiredString)
		o.DateExpired = &DateExpired
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trustcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
