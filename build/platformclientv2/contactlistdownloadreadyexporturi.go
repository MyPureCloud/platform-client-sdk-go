package platformclientv2
import (
	"encoding/json"
)

// Contactlistdownloadreadyexporturi
type Contactlistdownloadreadyexporturi struct { 
	// Uri
	Uri *string `json:"uri,omitempty"`


	// ExportTimestamp
	ExportTimestamp *string `json:"exportTimestamp,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contactlistdownloadreadyexporturi) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
