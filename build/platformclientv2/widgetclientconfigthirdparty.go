package platformclientv2
import (
	"encoding/json"
)

// Widgetclientconfigthirdparty
type Widgetclientconfigthirdparty struct { }

// String returns a JSON representation of the model
func (o *Widgetclientconfigthirdparty) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
