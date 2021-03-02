package platformclientv2
import (
	"time"
	"encoding/json"
)

// Adhocrecordingtopiclockdata
type Adhocrecordingtopiclockdata struct { 
	// LockedBy
	LockedBy *Adhocrecordingtopicuserdata `json:"lockedBy,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateExpires
	DateExpires *time.Time `json:"dateExpires,omitempty"`

}

// String returns a JSON representation of the model
func (o *Adhocrecordingtopiclockdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
