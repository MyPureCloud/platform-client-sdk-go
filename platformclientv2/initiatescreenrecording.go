package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Initiatescreenrecording
type Initiatescreenrecording struct { 
	// RecordACW
	RecordACW *bool `json:"recordACW,omitempty"`


	// ArchiveRetention
	ArchiveRetention *Archiveretention `json:"archiveRetention,omitempty"`


	// DeleteRetention
	DeleteRetention *Deleteretention `json:"deleteRetention,omitempty"`

}

func (u *Initiatescreenrecording) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Initiatescreenrecording

	

	return json.Marshal(&struct { 
		RecordACW *bool `json:"recordACW,omitempty"`
		
		ArchiveRetention *Archiveretention `json:"archiveRetention,omitempty"`
		
		DeleteRetention *Deleteretention `json:"deleteRetention,omitempty"`
		*Alias
	}{ 
		RecordACW: u.RecordACW,
		
		ArchiveRetention: u.ArchiveRetention,
		
		DeleteRetention: u.DeleteRetention,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Initiatescreenrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
