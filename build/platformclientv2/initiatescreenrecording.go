package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Initiatescreenrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
