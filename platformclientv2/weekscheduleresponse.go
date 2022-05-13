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

func (o *Weekscheduleresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Weekscheduleresponse
	
	return json.Marshal(&struct { 
		Result *Weekschedule `json:"result,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		Result: o.Result,
		
		DownloadUrl: o.DownloadUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Weekscheduleresponse) UnmarshalJSON(b []byte) error {
	var WeekscheduleresponseMap map[string]interface{}
	err := json.Unmarshal(b, &WeekscheduleresponseMap)
	if err != nil {
		return err
	}
	
	if Result, ok := WeekscheduleresponseMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if DownloadUrl, ok := WeekscheduleresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Weekscheduleresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
