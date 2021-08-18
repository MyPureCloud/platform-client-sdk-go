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

func (u *Buforecastresultresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buforecastresultresponse

	

	return json.Marshal(&struct { 
		Result *Buforecastresult `json:"result,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		Result: u.Result,
		
		DownloadUrl: u.DownloadUrl,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buforecastresultresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
