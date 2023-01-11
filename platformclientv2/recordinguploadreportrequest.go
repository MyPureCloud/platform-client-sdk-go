package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Recordinguploadreportrequest
type Recordinguploadreportrequest struct { 
	// DateSince - Report will include uploads since this date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateSince *time.Time `json:"dateSince,omitempty"`


	// UploadStatus - Report will include uploads with this status
	UploadStatus *string `json:"uploadStatus,omitempty"`

}

func (o *Recordinguploadreportrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Recordinguploadreportrequest
	
	DateSince := new(string)
	if o.DateSince != nil {
		
		*DateSince = timeutil.Strftime(o.DateSince, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateSince = nil
	}
	
	return json.Marshal(&struct { 
		DateSince *string `json:"dateSince,omitempty"`
		
		UploadStatus *string `json:"uploadStatus,omitempty"`
		*Alias
	}{ 
		DateSince: DateSince,
		
		UploadStatus: o.UploadStatus,
		Alias:    (*Alias)(o),
	})
}

func (o *Recordinguploadreportrequest) UnmarshalJSON(b []byte) error {
	var RecordinguploadreportrequestMap map[string]interface{}
	err := json.Unmarshal(b, &RecordinguploadreportrequestMap)
	if err != nil {
		return err
	}
	
	if dateSinceString, ok := RecordinguploadreportrequestMap["dateSince"].(string); ok {
		DateSince, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateSinceString)
		o.DateSince = &DateSince
	}
	
	if UploadStatus, ok := RecordinguploadreportrequestMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Recordinguploadreportrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
