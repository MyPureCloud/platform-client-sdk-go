package platformclientv2
import (
	"encoding/json"
)

// Burescheduleagentscheduleresult
type Burescheduleagentscheduleresult struct { 
	// ManagementUnit - The management unit to which this part of the result applies
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`


	// DownloadResult - The agent schedules.  Result will always come via the downloadUrl; however the schema is included for documentation
	DownloadResult *Murescheduleresultwrapper `json:"downloadResult,omitempty"`


	// DownloadUrl - The download URL from which to fetch the result
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

// String returns a JSON representation of the model
func (o *Burescheduleagentscheduleresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
