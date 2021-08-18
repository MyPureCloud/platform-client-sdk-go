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

func (u *Lockinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lockinfo

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateExpires := new(string)
	if u.DateExpires != nil {
		
		*DateExpires = timeutil.Strftime(u.DateExpires, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		LockedBy: u.LockedBy,
		
		DateCreated: DateCreated,
		
		DateExpires: DateExpires,
		
		Action: u.Action,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Lockinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
