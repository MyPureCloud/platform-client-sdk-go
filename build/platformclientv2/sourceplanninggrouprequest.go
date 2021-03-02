package platformclientv2
import (
	"encoding/json"
)

// Sourceplanninggrouprequest - Source planning group
type Sourceplanninggrouprequest struct { 
	// Id - The ID of the planning group
	Id *string `json:"id,omitempty"`


	// Metadata - Version metadata for the planning group
	Metadata *Wfmversionedentitymetadata `json:"metadata,omitempty"`

}

// String returns a JSON representation of the model
func (o *Sourceplanninggrouprequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
