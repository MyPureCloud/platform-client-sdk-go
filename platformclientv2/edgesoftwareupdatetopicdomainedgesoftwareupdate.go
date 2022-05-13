package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Edgesoftwareupdatetopicdomainedgesoftwareupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgesoftwareupdatetopicdomainedgesoftwareupdate
	
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
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		DownloadStartTime *string `json:"downloadStartTime,omitempty"`
		
		ExecuteStartTime *string `json:"executeStartTime,omitempty"`
		
		ExecuteStopTime *string `json:"executeStopTime,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Status: o.Status,
		
		DownloadStartTime: DownloadStartTime,
		
		ExecuteStartTime: ExecuteStartTime,
		
		ExecuteStopTime: ExecuteStopTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgesoftwareupdatetopicdomainedgesoftwareupdate) UnmarshalJSON(b []byte) error {
	var EdgesoftwareupdatetopicdomainedgesoftwareupdateMap map[string]interface{}
	err := json.Unmarshal(b, &EdgesoftwareupdatetopicdomainedgesoftwareupdateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EdgesoftwareupdatetopicdomainedgesoftwareupdateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Status, ok := EdgesoftwareupdatetopicdomainedgesoftwareupdateMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if downloadStartTimeString, ok := EdgesoftwareupdatetopicdomainedgesoftwareupdateMap["downloadStartTime"].(string); ok {
		DownloadStartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", downloadStartTimeString)
		o.DownloadStartTime = &DownloadStartTime
	}
	
	if executeStartTimeString, ok := EdgesoftwareupdatetopicdomainedgesoftwareupdateMap["executeStartTime"].(string); ok {
		ExecuteStartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", executeStartTimeString)
		o.ExecuteStartTime = &ExecuteStartTime
	}
	
	if executeStopTimeString, ok := EdgesoftwareupdatetopicdomainedgesoftwareupdateMap["executeStopTime"].(string); ok {
		ExecuteStopTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", executeStopTimeString)
		o.ExecuteStopTime = &ExecuteStopTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgesoftwareupdatetopicdomainedgesoftwareupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
