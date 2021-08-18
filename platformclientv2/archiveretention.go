package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Archiveretention
type Archiveretention struct { 
	// Days
	Days *int `json:"days,omitempty"`


	// StorageMedium
	StorageMedium *string `json:"storageMedium,omitempty"`

}

func (u *Archiveretention) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Archiveretention

	

	return json.Marshal(&struct { 
		Days *int `json:"days,omitempty"`
		
		StorageMedium *string `json:"storageMedium,omitempty"`
		*Alias
	}{ 
		Days: u.Days,
		
		StorageMedium: u.StorageMedium,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Archiveretention) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
