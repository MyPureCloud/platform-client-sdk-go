package platformclientv2
import (
	"time"
	"encoding/json"
)

// Domainedgesoftwareupdatedto
type Domainedgesoftwareupdatedto struct { 
	// Version - Version
	Version *Domainedgesoftwareversiondto `json:"version,omitempty"`


	// MaxDownloadRate
	MaxDownloadRate *int32 `json:"maxDownloadRate,omitempty"`


	// DownloadStartTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DownloadStartTime *time.Time `json:"downloadStartTime,omitempty"`


	// ExecuteStartTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ExecuteStartTime *time.Time `json:"executeStartTime,omitempty"`


	// ExecuteStopTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ExecuteStopTime *time.Time `json:"executeStopTime,omitempty"`


	// ExecuteOnIdle
	ExecuteOnIdle *bool `json:"executeOnIdle,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// EdgeUri
	EdgeUri *string `json:"edgeUri,omitempty"`


	// CallDrainingWaitTimeSeconds
	CallDrainingWaitTimeSeconds *int64 `json:"callDrainingWaitTimeSeconds,omitempty"`


	// Current
	Current *bool `json:"current,omitempty"`

}

// String returns a JSON representation of the model
func (o *Domainedgesoftwareupdatedto) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
