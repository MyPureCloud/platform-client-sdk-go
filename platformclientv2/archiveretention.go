package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Archiveretention) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
