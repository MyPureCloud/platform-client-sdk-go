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

func (u *Buheadcountforecastresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buheadcountforecastresponse

	

	return json.Marshal(&struct { 
		Result *Buheadcountforecast `json:"result,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		Result: u.Result,
		
		DownloadUrl: u.DownloadUrl,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buheadcountforecastresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
