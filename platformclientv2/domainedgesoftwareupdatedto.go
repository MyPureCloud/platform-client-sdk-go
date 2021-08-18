package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainedgesoftwareupdatedto
type Domainedgesoftwareupdatedto struct { 
	// Version - Version
	Version *Domainedgesoftwareversiondto `json:"version,omitempty"`


	// MaxDownloadRate
	MaxDownloadRate *int `json:"maxDownloadRate,omitempty"`


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
	CallDrainingWaitTimeSeconds *int `json:"callDrainingWaitTimeSeconds,omitempty"`


	// Current
	Current *bool `json:"current,omitempty"`

}

func (u *Domainedgesoftwareupdatedto) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainedgesoftwareupdatedto

	
	DownloadStartTime := new(string)
	if u.DownloadStartTime != nil {
		
		*DownloadStartTime = timeutil.Strftime(u.DownloadStartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DownloadStartTime = nil
	}
	
	ExecuteStartTime := new(string)
	if u.ExecuteStartTime != nil {
		
		*ExecuteStartTime = timeutil.Strftime(u.ExecuteStartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExecuteStartTime = nil
	}
	
	ExecuteStopTime := new(string)
	if u.ExecuteStopTime != nil {
		
		*ExecuteStopTime = timeutil.Strftime(u.ExecuteStopTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExecuteStopTime = nil
	}
	

	return json.Marshal(&struct { 
		Version *Domainedgesoftwareversiondto `json:"version,omitempty"`
		
		MaxDownloadRate *int `json:"maxDownloadRate,omitempty"`
		
		DownloadStartTime *string `json:"downloadStartTime,omitempty"`
		
		ExecuteStartTime *string `json:"executeStartTime,omitempty"`
		
		ExecuteStopTime *string `json:"executeStopTime,omitempty"`
		
		ExecuteOnIdle *bool `json:"executeOnIdle,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		EdgeUri *string `json:"edgeUri,omitempty"`
		
		CallDrainingWaitTimeSeconds *int `json:"callDrainingWaitTimeSeconds,omitempty"`
		
		Current *bool `json:"current,omitempty"`
		*Alias
	}{ 
		Version: u.Version,
		
		MaxDownloadRate: u.MaxDownloadRate,
		
		DownloadStartTime: DownloadStartTime,
		
		ExecuteStartTime: ExecuteStartTime,
		
		ExecuteStopTime: ExecuteStopTime,
		
		ExecuteOnIdle: u.ExecuteOnIdle,
		
		Status: u.Status,
		
		EdgeUri: u.EdgeUri,
		
		CallDrainingWaitTimeSeconds: u.CallDrainingWaitTimeSeconds,
		
		Current: u.Current,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Domainedgesoftwareupdatedto) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
