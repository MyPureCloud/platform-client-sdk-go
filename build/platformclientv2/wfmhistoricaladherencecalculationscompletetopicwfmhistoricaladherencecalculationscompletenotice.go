package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenotice
type Wfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenotice struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`


	// DownloadUrls
	DownloadUrls *[]string `json:"downloadUrls,omitempty"`


	// QueryState
	QueryState *string `json:"queryState,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenotice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
