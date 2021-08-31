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

func (o *Useravailabletimes) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Useravailabletimes
	
	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		AvailableTimes *[]Availabletime `json:"availableTimes,omitempty"`
		*Alias
	}{ 
		User: o.User,
		
		AvailableTimes: o.AvailableTimes,
		Alias:    (*Alias)(o),
	})
}

func (o *Useravailabletimes) UnmarshalJSON(b []byte) error {
	var UseravailabletimesMap map[string]interface{}
	err := json.Unmarshal(b, &UseravailabletimesMap)
	if err != nil {
		return err
	}
	
	if User, ok := UseravailabletimesMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if AvailableTimes, ok := UseravailabletimesMap["availableTimes"].([]interface{}); ok {
		AvailableTimesString, _ := json.Marshal(AvailableTimes)
		json.Unmarshal(AvailableTimesString, &o.AvailableTimes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Useravailabletimes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
