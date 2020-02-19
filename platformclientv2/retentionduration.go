package platformclientv2
import (
	"encoding/json"
)

// Retentionduration
type Retentionduration struct { 
	// ArchiveRetention
	ArchiveRetention *Archiveretention `json:"archiveRetention,omitempty"`


	// DeleteRetention
	DeleteRetention *Deleteretention `json:"deleteRetention,omitempty"`

}

// String returns a JSON representation of the model
func (o *Retentionduration) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
