package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuschedulesearchresulttopicbuschedulesearchresultnotification
type Wfmbuschedulesearchresulttopicbuschedulesearchresultnotification struct { 
	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// BusinessUnitId
	BusinessUnitId *string `json:"businessUnitId,omitempty"`


	// DownloadUrl
	DownloadUrl *string `json:"downloadUrl,omitempty"`

}

func (u *Wfmbuschedulesearchresulttopicbuschedulesearchresultnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuschedulesearchresulttopicbuschedulesearchresultnotification

	

	return json.Marshal(&struct { 
		OperationId *string `json:"operationId,omitempty"`
		
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		
		DownloadUrl *string `json:"downloadUrl,omitempty"`
		*Alias
	}{ 
		OperationId: u.OperationId,
		
		BusinessUnitId: u.BusinessUnitId,
		
		DownloadUrl: u.DownloadUrl,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbuschedulesearchresulttopicbuschedulesearchresultnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
