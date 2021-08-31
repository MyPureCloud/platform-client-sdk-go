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

func (o *Historyentry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historyentry
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Action: o.Action,
		
		Resource: o.Resource,
		
		Timestamp: Timestamp,
		
		User: o.User,
		
		Client: o.Client,
		
		Version: o.Version,
		
		Secure: o.Secure,
		Alias:    (*Alias)(o),
	})
}

func (o *Historyentry) UnmarshalJSON(b []byte) error {
	var HistoryentryMap map[string]interface{}
	err := json.Unmarshal(b, &HistoryentryMap)
	if err != nil {
		return err
	}
	
	if Action, ok := HistoryentryMap["action"].(string); ok {
		o.Action = &Action
	}
	
	if Resource, ok := HistoryentryMap["resource"].(string); ok {
		o.Resource = &Resource
	}
	
	if timestampString, ok := HistoryentryMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if User, ok := HistoryentryMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Client, ok := HistoryentryMap["client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}
	
	if Version, ok := HistoryentryMap["version"].(string); ok {
		o.Version = &Version
	}
	
	if Secure, ok := HistoryentryMap["secure"].(bool); ok {
		o.Secure = &Secure
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historyentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
