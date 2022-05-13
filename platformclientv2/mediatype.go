package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediatype - Media type definition
type Mediatype struct { 
	// VarType - The media type string as defined by RFC 2046. You can define specific types such as 'image/jpeg', 'video/mpeg', or specify wild cards for a range of types, 'image/_*', or all types '*_/_*'. See https://www.iana.org/assignments/media-types/media-types.xhtml for a list of registered media types.
	VarType *string `json:"type,omitempty"`

}

func (o *Mediatype) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediatype
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Mediatype) UnmarshalJSON(b []byte) error {
	var MediatypeMap map[string]interface{}
	err := json.Unmarshal(b, &MediatypeMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := MediatypeMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Mediatype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
