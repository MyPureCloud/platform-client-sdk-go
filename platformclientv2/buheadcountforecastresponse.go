package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buheadcountforecastresponse
type Buheadcountforecastresponse struct { 
	// Result - The headcount forecast, null when downloadUrl is provided
	Result *Buheadcountforecast `json:"result,omitempty"`


	// DownloadUrl - Download URL.  Null unless the response is too large to pass directly through the api
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

func (o *Buheadcountforecastresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buheadcountforecastresponse
	
	return json.Marshal(&struct { 
		Result *Buheadcountforecast `json:"result,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		Result: o.Result,
		
		DownloadUrl: o.DownloadUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Buheadcountforecastresponse) UnmarshalJSON(b []byte) error {
	var BuheadcountforecastresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuheadcountforecastresponseMap)
	if err != nil {
		return err
	}
	
	if Result, ok := BuheadcountforecastresponseMap["result"].(map[string]interface{}); ok {
		ResultString, _ := json.Marshal(Result)
		json.Unmarshal(ResultString, &o.Result)
	}
	
	if DownloadUrl, ok := BuheadcountforecastresponseMap["downloadUrl"].(string); ok {
		o.DownloadUrl = &DownloadUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Buheadcountforecastresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
