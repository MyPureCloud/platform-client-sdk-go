package platformclientv2
import (
	"encoding/json"
)

// Widgetclientconfigv2
type Widgetclientconfigv2 struct { }

// String returns a JSON representation of the model
func (o *Widgetclientconfigv2) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
