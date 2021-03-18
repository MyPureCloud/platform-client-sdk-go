package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Edgelogicalinterfaceschangetopicdomainlogicalinterfacechange
type Edgelogicalinterfaceschangetopicdomainlogicalinterfacechange struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ErrorInfo
	ErrorInfo *Edgelogicalinterfaceschangetopicerrorinfo `json:"errorInfo,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgelogicalinterfaceschangetopicdomainlogicalinterfacechange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
