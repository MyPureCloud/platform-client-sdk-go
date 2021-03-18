package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Buheadcountforecastresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
