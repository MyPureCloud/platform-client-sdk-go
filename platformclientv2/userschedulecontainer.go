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

func (o *Userschedulecontainer) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userschedulecontainer
	
	return json.Marshal(&struct { 
		ManagementUnitTimeZone *string `json:"managementUnitTimeZone,omitempty"`
		
		PublishedSchedules *[]Weekschedulereference `json:"publishedSchedules,omitempty"`
		
		UserSchedules *map[string]Userschedule `json:"userSchedules,omitempty"`
		*Alias
	}{ 
		ManagementUnitTimeZone: o.ManagementUnitTimeZone,
		
		PublishedSchedules: o.PublishedSchedules,
		
		UserSchedules: o.UserSchedules,
		Alias:    (*Alias)(o),
	})
}

func (o *Userschedulecontainer) UnmarshalJSON(b []byte) error {
	var UserschedulecontainerMap map[string]interface{}
	err := json.Unmarshal(b, &UserschedulecontainerMap)
	if err != nil {
		return err
	}
	
	if ManagementUnitTimeZone, ok := UserschedulecontainerMap["managementUnitTimeZone"].(string); ok {
		o.ManagementUnitTimeZone = &ManagementUnitTimeZone
	}
	
	if PublishedSchedules, ok := UserschedulecontainerMap["publishedSchedules"].([]interface{}); ok {
		PublishedSchedulesString, _ := json.Marshal(PublishedSchedules)
		json.Unmarshal(PublishedSchedulesString, &o.PublishedSchedules)
	}
	
	if UserSchedules, ok := UserschedulecontainerMap["userSchedules"].(map[string]interface{}); ok {
		UserSchedulesString, _ := json.Marshal(UserSchedules)
		json.Unmarshal(UserSchedulesString, &o.UserSchedules)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userschedulecontainer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
