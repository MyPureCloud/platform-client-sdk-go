package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Wfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenotice) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenotice

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		
		DownloadUrls *[]string `json:"downloadUrls,omitempty"`
		
		QueryState *string `json:"queryState,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		DownloadUrl: u.DownloadUrl,
		
		DownloadUrls: u.DownloadUrls,
		
		QueryState: u.QueryState,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencecalculationscompletetopicwfmhistoricaladherencecalculationscompletenotice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
