package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
