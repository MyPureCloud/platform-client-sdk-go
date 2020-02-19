package platformclientv2
import (
	"time"
	"encoding/json"
)

// Contentmanagementworkspacedocumentstopiclockdata
type Contentmanagementworkspacedocumentstopiclockdata struct { 
	// LockedBy
	LockedBy *Contentmanagementworkspacedocumentstopicuserdata `json:"lockedBy,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateExpires
	DateExpires *time.Time `json:"dateExpires,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentmanagementworkspacedocumentstopiclockdata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
