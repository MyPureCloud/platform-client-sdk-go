package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Mediatype - Media type definition
type Mediatype struct { 
	// VarType - The media type string as defined by RFC 2046. You can define specific types such as 'image/jpeg', 'video/mpeg', or specify wild cards for a range of types, 'image/*', or all types '*/*'. See https://www.iana.org/assignments/media-types/media-types.xhtml for a list of registered media types.
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Mediatype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
