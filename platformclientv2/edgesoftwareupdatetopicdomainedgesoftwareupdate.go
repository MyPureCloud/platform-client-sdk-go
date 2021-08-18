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

func (u *Edgesoftwareupdatetopicdomainedgesoftwareupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgesoftwareupdatetopicdomainedgesoftwareupdate

	
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
		Id *string `json:"id,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		DownloadStartTime *string `json:"downloadStartTime,omitempty"`
		
		ExecuteStartTime *string `json:"executeStartTime,omitempty"`
		
		ExecuteStopTime *string `json:"executeStopTime,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Status: u.Status,
		
		DownloadStartTime: DownloadStartTime,
		
		ExecuteStartTime: ExecuteStartTime,
		
		ExecuteStopTime: ExecuteStopTime,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgesoftwareupdatetopicdomainedgesoftwareupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
