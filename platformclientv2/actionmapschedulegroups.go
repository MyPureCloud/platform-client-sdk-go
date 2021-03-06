package platformclientv2
import (
	"encoding/json"
)

// Actionmapschedulegroups
type Actionmapschedulegroups struct { 
	// ActionMapScheduleGroup - The actions map's associated schedule group.
	ActionMapScheduleGroup *Actionmapschedulegroup `json:"actionMapScheduleGroup,omitempty"`


	// EmergencyActionMapScheduleGroup - The action map's associated emergency schedule group.
	EmergencyActionMapScheduleGroup *Actionmapschedulegroup `json:"emergencyActionMapScheduleGroup,omitempty"`

}

// String returns a JSON representation of the model
func (o *Actionmapschedulegroups) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
