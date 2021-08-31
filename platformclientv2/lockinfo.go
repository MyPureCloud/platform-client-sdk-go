package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lockinfo
type Lockinfo struct { 
	// LockedBy
	LockedBy *Domainentityref `json:"lockedBy,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateExpires - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateExpires *time.Time `json:"dateExpires,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

func (o *Lockinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lockinfo
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateExpires := new(string)
	if o.DateExpires != nil {
		
		*DateExpires = timeutil.Strftime(o.DateExpires, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateExpires = nil
	}
	
	return json.Marshal(&struct { 
		LockedBy *Domainentityref `json:"lockedBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateExpires *string `json:"dateExpires,omitempty"`
		
		Action *string `json:"action,omitempty"`
		*Alias
	}{ 
		LockedBy: o.LockedBy,
		
		DateCreated: DateCreated,
		
		DateExpires: DateExpires,
		
		Action: o.Action,
		Alias:    (*Alias)(o),
	})
}

func (o *Lockinfo) UnmarshalJSON(b []byte) error {
	var LockinfoMap map[string]interface{}
	err := json.Unmarshal(b, &LockinfoMap)
	if err != nil {
		return err
	}
	
	if LockedBy, ok := LockinfoMap["lockedBy"].(map[string]interface{}); ok {
		LockedByString, _ := json.Marshal(LockedBy)
		json.Unmarshal(LockedByString, &o.LockedBy)
	}
	
	if dateCreatedString, ok := LockinfoMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateExpiresString, ok := LockinfoMap["dateExpires"].(string); ok {
		DateExpires, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateExpiresString)
		o.DateExpires = &DateExpires
	}
	
	if Action, ok := LockinfoMap["action"].(string); ok {
		o.Action = &Action
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Lockinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
