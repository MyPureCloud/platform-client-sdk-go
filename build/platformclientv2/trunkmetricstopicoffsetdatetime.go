package platformclientv2
import (
	"encoding/json"
)

// Trunkmetricstopicoffsetdatetime
type Trunkmetricstopicoffsetdatetime struct { 
	// DateTime
	DateTime *Trunkmetricstopiclocaldatetime `json:"dateTime,omitempty"`


	// Offset
	Offset *Trunkmetricstopiczoneoffset `json:"offset,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricstopicoffsetdatetime) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
