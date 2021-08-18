package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Historyentry
type Historyentry struct { 
	// Action - The action performed
	Action *string `json:"action,omitempty"`


	// Resource - For actions performed not on the item itself, but on a sub-item, this field identifies the sub-item by name.  For example, for actions performed on prompt resources, this will be the prompt resource name.
	Resource *string `json:"resource,omitempty"`


	// Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// User - User associated with this entry.
	User *User `json:"user,omitempty"`


	// Client - OAuth client associated with this entry.
	Client *Domainentityref `json:"client,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// Secure
	Secure *bool `json:"secure,omitempty"`

}

func (u *Historyentry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historyentry

	
	Timestamp := new(string)
	if u.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(u.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	

	return json.Marshal(&struct { 
		Action *string `json:"action,omitempty"`
		
		Resource *string `json:"resource,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Client *Domainentityref `json:"client,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		Secure *bool `json:"secure,omitempty"`
		*Alias
	}{ 
		Action: u.Action,
		
		Resource: u.Resource,
		
		Timestamp: Timestamp,
		
		User: u.User,
		
		Client: u.Client,
		
		Version: u.Version,
		
		Secure: u.Secure,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Historyentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
