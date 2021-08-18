package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Useravailabletimes
type Useravailabletimes struct { 
	// User - User reference
	User *Userreference `json:"user,omitempty"`


	// AvailableTimes - Periods of availability to schedule coaching appointment for an user
	AvailableTimes *[]Availabletime `json:"availableTimes,omitempty"`

}

func (u *Useravailabletimes) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Useravailabletimes

	

	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		AvailableTimes *[]Availabletime `json:"availableTimes,omitempty"`
		*Alias
	}{ 
		User: u.User,
		
		AvailableTimes: u.AvailableTimes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Useravailabletimes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
