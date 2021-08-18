package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userschedulecontainer - Container object to hold a map of user schedules
type Userschedulecontainer struct { 
	// ManagementUnitTimeZone - The reference time zone used for the management unit
	ManagementUnitTimeZone *string `json:"managementUnitTimeZone,omitempty"`


	// PublishedSchedules - References to all published week schedules overlapping the start/end date query parameters
	PublishedSchedules *[]Weekschedulereference `json:"publishedSchedules,omitempty"`


	// UserSchedules - Map of user id to user schedule
	UserSchedules *map[string]Userschedule `json:"userSchedules,omitempty"`

}

func (u *Userschedulecontainer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userschedulecontainer

	

	return json.Marshal(&struct { 
		ManagementUnitTimeZone *string `json:"managementUnitTimeZone,omitempty"`
		
		PublishedSchedules *[]Weekschedulereference `json:"publishedSchedules,omitempty"`
		
		UserSchedules *map[string]Userschedule `json:"userSchedules,omitempty"`
		*Alias
	}{ 
		ManagementUnitTimeZone: u.ManagementUnitTimeZone,
		
		PublishedSchedules: u.PublishedSchedules,
		
		UserSchedules: u.UserSchedules,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userschedulecontainer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
