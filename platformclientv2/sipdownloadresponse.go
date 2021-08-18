package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sipdownloadresponse
type Sipdownloadresponse struct { 
	// DownloadId - unique id of the downloaded file
	DownloadId *string `json:"downloadId,omitempty"`


	// DocumentId - Document id of pcap file
	DocumentId *string `json:"documentId,omitempty"`

}

func (u *Sipdownloadresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sipdownloadresponse

	

	return json.Marshal(&struct { 
		DownloadId *string `json:"downloadId,omitempty"`
		
		DocumentId *string `json:"documentId,omitempty"`
		*Alias
	}{ 
		DownloadId: u.DownloadId,
		
		DocumentId: u.DocumentId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Sipdownloadresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
