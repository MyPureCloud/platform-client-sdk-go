package platformclientv2
import (
	"encoding/json"
)

// Contentmanagementworkspacedocumentstopicworkspacedata
type Contentmanagementworkspacedocumentstopicworkspacedata struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentmanagementworkspacedocumentstopicworkspacedata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
