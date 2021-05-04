package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Useravailabletimes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
