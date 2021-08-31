package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification
type Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// User
	User *Architectdependencytrackingbuildnotificationuser `json:"user,omitempty"`


	// Client
	Client *Architectdependencytrackingbuildnotificationclient `json:"client,omitempty"`


	// StartTime
	StartTime *time.Time `json:"startTime,omitempty"`

}

func (o *Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		User *Architectdependencytrackingbuildnotificationuser `json:"user,omitempty"`
		
		Client *Architectdependencytrackingbuildnotificationclient `json:"client,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		
		User: o.User,
		
		Client: o.Client,
		
		StartTime: StartTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification) UnmarshalJSON(b []byte) error {
	var ArchitectdependencytrackingbuildnotificationdependencytrackingbuildnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectdependencytrackingbuildnotificationdependencytrackingbuildnotificationMap)
	if err != nil {
		return err
	}
	
	if Status, ok := ArchitectdependencytrackingbuildnotificationdependencytrackingbuildnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if User, ok := ArchitectdependencytrackingbuildnotificationdependencytrackingbuildnotificationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Client, ok := ArchitectdependencytrackingbuildnotificationdependencytrackingbuildnotificationMap["client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}
	
	if startTimeString, ok := ArchitectdependencytrackingbuildnotificationdependencytrackingbuildnotificationMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectdependencytrackingbuildnotificationdependencytrackingbuildnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
