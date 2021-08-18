package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchactionmapschedulegroups
type Patchactionmapschedulegroups struct { 
	// ActionMapScheduleGroup - The actions map's associated schedule group.
	ActionMapScheduleGroup *Actionmapschedulegroup `json:"actionMapScheduleGroup,omitempty"`


	// EmergencyActionMapScheduleGroup - The action map's associated emergency schedule group.
	EmergencyActionMapScheduleGroup *Actionmapschedulegroup `json:"emergencyActionMapScheduleGroup,omitempty"`

}

func (u *Patchactionmapschedulegroups) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchactionmapschedulegroups

	

	return json.Marshal(&struct { 
		ActionMapScheduleGroup *Actionmapschedulegroup `json:"actionMapScheduleGroup,omitempty"`
		
		EmergencyActionMapScheduleGroup *Actionmapschedulegroup `json:"emergencyActionMapScheduleGroup,omitempty"`
		*Alias
	}{ 
		ActionMapScheduleGroup: u.ActionMapScheduleGroup,
		
		EmergencyActionMapScheduleGroup: u.EmergencyActionMapScheduleGroup,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Patchactionmapschedulegroups) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
