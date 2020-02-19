package platformclientv2
import (
	"encoding/json"
)

// Contactcolumntodataactionfieldmapping
type Contactcolumntodataactionfieldmapping struct { }

// String returns a JSON representation of the model
func (o *Contactcolumntodataactionfieldmapping) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
