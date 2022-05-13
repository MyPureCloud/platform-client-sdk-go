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

func (o *Edgelogicalinterfaceschangetopicdomainlogicalinterfacechange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgelogicalinterfaceschangetopicdomainlogicalinterfacechange
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ErrorInfo *Edgelogicalinterfaceschangetopicerrorinfo `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ErrorInfo: o.ErrorInfo,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgelogicalinterfaceschangetopicdomainlogicalinterfacechange) UnmarshalJSON(b []byte) error {
	var EdgelogicalinterfaceschangetopicdomainlogicalinterfacechangeMap map[string]interface{}
	err := json.Unmarshal(b, &EdgelogicalinterfaceschangetopicdomainlogicalinterfacechangeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EdgelogicalinterfaceschangetopicdomainlogicalinterfacechangeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ErrorInfo, ok := EdgelogicalinterfaceschangetopicdomainlogicalinterfacechangeMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgelogicalinterfaceschangetopicdomainlogicalinterfacechange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
