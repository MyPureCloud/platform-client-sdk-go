package platformclientv2
import (
	"time"
	"encoding/json"
)

// Edgesoftwareupdatetopicdomainedgesoftwareupdate
type Edgesoftwareupdatetopicdomainedgesoftwareupdate struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// DownloadStartTime
	DownloadStartTime *time.Time `json:"downloadStartTime,omitempty"`


	// ExecuteStartTime
	ExecuteStartTime *time.Time `json:"executeStartTime,omitempty"`


	// ExecuteStopTime
	ExecuteStopTime *time.Time `json:"executeStopTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgesoftwareupdatetopicdomainedgesoftwareupdate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
