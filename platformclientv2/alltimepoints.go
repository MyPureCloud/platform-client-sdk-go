package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Alltimepoints
type Alltimepoints struct { 
	// User - Queried user
	User *Userreference `json:"user,omitempty"`


	// DateEndWorkday - Queried end workday for all time points to be collected. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`


	// AllTimePoints - All time point collected bt the user
	AllTimePoints *int `json:"allTimePoints,omitempty"`

}

func (u *Alltimepoints) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Alltimepoints

	
	DateEndWorkday := new(string)
	if u.DateEndWorkday != nil {
		*DateEndWorkday = timeutil.Strftime(u.DateEndWorkday, "%Y-%m-%d")
	} else {
		DateEndWorkday = nil
	}
	

	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		DateEndWorkday *string `json:"dateEndWorkday,omitempty"`
		
		AllTimePoints *int `json:"allTimePoints,omitempty"`
		*Alias
	}{ 
		User: u.User,
		
		DateEndWorkday: DateEndWorkday,
		
		AllTimePoints: u.AllTimePoints,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Alltimepoints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
