package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherenceresponse - Response for Historical Adherence Query, intended to tell the client what to listen for on a notification topic
type Wfmhistoricaladherenceresponse struct { 
	// Id - The query ID to listen for
	Id *string `json:"id,omitempty"`


	// DownloadUrl - Deprecated. Use downloadUrls instead.
	DownloadUrl *string `json:"downloadUrl,omitempty"`


	// DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
	DownloadResult *Wfmhistoricaladherenceresultwrapper `json:"downloadResult,omitempty"`


	// DownloadUrls - The uri list to GET the results of the Historical Adherence query. For notification purposes only
	DownloadUrls *[]string `json:"downloadUrls,omitempty"`


	// QueryState - The state of the adherence query
	QueryState *string `json:"queryState,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherenceresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
