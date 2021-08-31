package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Longtermforecastresultresponse
type Longtermforecastresultresponse struct { 
	// Result - The result of the operation.  Populated whenever the result is small enough to pass through the api directly
	Result *Longtermforecastresult `json:"result,omitempty"`


	// DownloadUrl - The download url to fetch the result.  Only populated if the result is too large to pass through the api directly
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

func (o *Longtermforecastresultresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Longtermforecastresultresponse
	
	return json.Marshal(&struct { 
		Result *Longtermforecastresult `json:"result,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		Result: o.Result,
		
		DownloadUrl: o.DownloadUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Longtermforecastresultresponse) UnmarshalJSON(b []byte) error {
	var LongtermforecastresultresponseMap map[string]interface{}
	err := json.Unmarshal(b, &LongtermforecastresultresponseMap)
	if err != nil {
		return err
	}
	
	if Result, ok := LongtermforecastresultresponseMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if DownloadUrl, ok := LongtermforecastresultresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Longtermforecastresultresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
