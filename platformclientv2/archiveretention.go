package platformclientv2
import (
	"encoding/json"
)

// Archiveretention
type Archiveretention struct { 
	// Days
	Days *int32 `json:"days,omitempty"`


	// StorageMedium
	StorageMedium *string `json:"storageMedium,omitempty"`

}

// String returns a JSON representation of the model
func (o *Archiveretention) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
