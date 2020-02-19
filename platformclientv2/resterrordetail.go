package platformclientv2
import (
	"encoding/json"
)

// Resterrordetail
type Resterrordetail struct { 
	// VarError - name of the error
	VarError *string `json:"error,omitempty"`


	// Details - additional information regarding the error
	Details *string `json:"details,omitempty"`

}

// String returns a JSON representation of the model
func (o *Resterrordetail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
