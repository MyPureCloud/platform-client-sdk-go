package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Retentionduration
type Retentionduration struct { 
	// ArchiveRetention
	ArchiveRetention *Archiveretention `json:"archiveRetention,omitempty"`


	// DeleteRetention
	DeleteRetention *Deleteretention `json:"deleteRetention,omitempty"`

}

func (u *Retentionduration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Retentionduration

	

	return json.Marshal(&struct { 
		ArchiveRetention *Archiveretention `json:"archiveRetention,omitempty"`
		
		DeleteRetention *Deleteretention `json:"deleteRetention,omitempty"`
		*Alias
	}{ 
		ArchiveRetention: u.ArchiveRetention,
		
		DeleteRetention: u.DeleteRetention,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Retentionduration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
