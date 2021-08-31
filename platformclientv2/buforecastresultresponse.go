package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buforecastresultresponse
type Buforecastresultresponse struct { 
	// Result - The result of the operation.  Populated whenever the result is small enough to pass through the api directly
	Result *Buforecastresult `json:"result,omitempty"`


	// DownloadUrl - The download url to fetch the result.  Only populated if the result is too large to pass through the api directly
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

func (o *Buforecastresultresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buforecastresultresponse
	
	return json.Marshal(&struct { 
		Result *Buforecastresult `json:"result,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		Result: o.Result,
		
		DownloadUrl: o.DownloadUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Buforecastresultresponse) UnmarshalJSON(b []byte) error {
	var BuforecastresultresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuforecastresultresponseMap)
	if err != nil {
		return err
	}
	
	if Result, ok := BuforecastresultresponseMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if DownloadUrl, ok := BuforecastresultresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buforecastresultresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
