package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Edgelogicalinterfaceschangetopicdomainlogicalinterfacechange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgelogicalinterfaceschangetopicdomainlogicalinterfacechange

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ErrorInfo *Edgelogicalinterfaceschangetopicerrorinfo `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		ErrorInfo: u.ErrorInfo,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgelogicalinterfaceschangetopicdomainlogicalinterfacechange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
