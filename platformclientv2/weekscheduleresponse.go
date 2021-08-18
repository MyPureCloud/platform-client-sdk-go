package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekscheduleresponse - Response for query for week schedule for a given week in management unit
type Weekscheduleresponse struct { 
	// Result - The result of the request. The value will be null if response is large
	Result *Weekschedule `json:"result,omitempty"`


	// DownloadUrl - The url to fetch the result for large responses. The value is null if result contains the data
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

func (u *Weekscheduleresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Weekscheduleresponse

	

	return json.Marshal(&struct { 
		Result *Weekschedule `json:"result,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		Result: u.Result,
		
		DownloadUrl: u.DownloadUrl,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Weekscheduleresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
