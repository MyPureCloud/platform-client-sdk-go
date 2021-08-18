package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentmanagementworkspacedocumentstopiclockdata
type Contentmanagementworkspacedocumentstopiclockdata struct { 
	// LockedBy
	LockedBy *Contentmanagementworkspacedocumentstopicuserdata `json:"lockedBy,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateExpires
	DateExpires *time.Time `json:"dateExpires,omitempty"`

}

func (u *Contentmanagementworkspacedocumentstopiclockdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentmanagementworkspacedocumentstopiclockdata

	
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
		LockedBy *Contentmanagementworkspacedocumentstopicuserdata `json:"lockedBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateExpires *string `json:"dateExpires,omitempty"`
		*Alias
	}{ 
		LockedBy: u.LockedBy,
		
		DateCreated: DateCreated,
		
		DateExpires: DateExpires,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contentmanagementworkspacedocumentstopiclockdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
