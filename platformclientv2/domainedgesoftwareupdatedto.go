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

func (o *Domainedgesoftwareupdatedto) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainedgesoftwareupdatedto
	
	DownloadStartTime := new(string)
	if o.DownloadStartTime != nil {
		
		*DownloadStartTime = timeutil.Strftime(o.DownloadStartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DownloadStartTime = nil
	}
	
	ExecuteStartTime := new(string)
	if o.ExecuteStartTime != nil {
		
		*ExecuteStartTime = timeutil.Strftime(o.ExecuteStartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ExecuteStartTime = nil
	}
	
	ExecuteStopTime := new(string)
	if o.ExecuteStopTime != nil {
		
		*ExecuteStopTime = timeutil.Strftime(o.ExecuteStopTime, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Version: o.Version,
		
		MaxDownloadRate: o.MaxDownloadRate,
		
		DownloadStartTime: DownloadStartTime,
		
		ExecuteStartTime: ExecuteStartTime,
		
		ExecuteStopTime: ExecuteStopTime,
		
		ExecuteOnIdle: o.ExecuteOnIdle,
		
		Status: o.Status,
		
		EdgeUri: o.EdgeUri,
		
		CallDrainingWaitTimeSeconds: o.CallDrainingWaitTimeSeconds,
		
		Current: o.Current,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainedgesoftwareupdatedto) UnmarshalJSON(b []byte) error {
	var DomainedgesoftwareupdatedtoMap map[string]interface{}
	err := json.Unmarshal(b, &DomainedgesoftwareupdatedtoMap)
	if err != nil {
		return err
	}
	
	if Version, ok := DomainedgesoftwareupdatedtoMap["version"].(map[string]interface{}); ok {
		VersionString, _ := json.Marshal(Version)
		json.Unmarshal(VersionString, &o.Version)
	}
	
	if MaxDownloadRate, ok := DomainedgesoftwareupdatedtoMap["maxDownloadRate"].(float64); ok {
		MaxDownloadRateInt := int(MaxDownloadRate)
		o.MaxDownloadRate = &MaxDownloadRateInt
	}
	
	if downloadStartTimeString, ok := DomainedgesoftwareupdatedtoMap["downloadStartTime"].(string); ok {
		DownloadStartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", downloadStartTimeString)
		o.DownloadStartTime = &DownloadStartTime
	}
	
	if executeStartTimeString, ok := DomainedgesoftwareupdatedtoMap["executeStartTime"].(string); ok {
		ExecuteStartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", executeStartTimeString)
		o.ExecuteStartTime = &ExecuteStartTime
	}
	
	if executeStopTimeString, ok := DomainedgesoftwareupdatedtoMap["executeStopTime"].(string); ok {
		ExecuteStopTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", executeStopTimeString)
		o.ExecuteStopTime = &ExecuteStopTime
	}
	
	if ExecuteOnIdle, ok := DomainedgesoftwareupdatedtoMap["executeOnIdle"].(bool); ok {
		o.ExecuteOnIdle = &ExecuteOnIdle
	}
	
	if Status, ok := DomainedgesoftwareupdatedtoMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if EdgeUri, ok := DomainedgesoftwareupdatedtoMap["edgeUri"].(string); ok {
		o.EdgeUri = &EdgeUri
	}
	
	if CallDrainingWaitTimeSeconds, ok := DomainedgesoftwareupdatedtoMap["callDrainingWaitTimeSeconds"].(float64); ok {
		CallDrainingWaitTimeSecondsInt := int(CallDrainingWaitTimeSeconds)
		o.CallDrainingWaitTimeSeconds = &CallDrainingWaitTimeSecondsInt
	}
	
	if Current, ok := DomainedgesoftwareupdatedtoMap["current"].(bool); ok {
		o.Current = &Current
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainedgesoftwareupdatedto) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
