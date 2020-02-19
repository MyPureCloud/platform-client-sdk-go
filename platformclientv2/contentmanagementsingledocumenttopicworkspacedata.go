package platformclientv2
import (
	"encoding/json"
)

// Contentmanagementsingledocumenttopicworkspacedata
type Contentmanagementsingledocumenttopicworkspacedata struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contentmanagementsingledocumenttopicworkspacedata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
