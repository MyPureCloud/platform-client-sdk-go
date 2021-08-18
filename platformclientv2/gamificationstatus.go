package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Gamificationstatus
type Gamificationstatus struct { 
	// IsActive - Gamification status of the organization.
	IsActive *bool `json:"isActive,omitempty"`


	// DateStart - Gamification start date. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`


	// AutomaticUserAssignment - Automatic assignment of users to the default profile
	AutomaticUserAssignment *bool `json:"automaticUserAssignment,omitempty"`

}

func (u *Gamificationstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Gamificationstatus

	
	DateStart := new(string)
	if u.DateStart != nil {
		*DateStart = timeutil.Strftime(u.DateStart, "%Y-%m-%d")
	} else {
		DateStart = nil
	}
	

	return json.Marshal(&struct { 
		IsActive *bool `json:"isActive,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		AutomaticUserAssignment *bool `json:"automaticUserAssignment,omitempty"`
		*Alias
	}{ 
		IsActive: u.IsActive,
		
		DateStart: DateStart,
		
		AutomaticUserAssignment: u.AutomaticUserAssignment,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Gamificationstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
