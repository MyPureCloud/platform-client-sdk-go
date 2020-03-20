package platformclientv2
import (
	"encoding/json"
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
	return string(j)
}
