package platformclientv2
import (
	"encoding/json"
)

// Wfmbuschedulequeryresulttopicbuschedulesearchresultnotification
type Wfmbuschedulequeryresulttopicbuschedulesearchresultnotification struct { 
	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// BusinessUnitId
	BusinessUnitId *string `json:"businessUnitId,omitempty"`


	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbuschedulequeryresulttopicbuschedulesearchresultnotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
