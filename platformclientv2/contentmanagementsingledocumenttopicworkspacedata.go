package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentmanagementsingledocumenttopicworkspacedata
type Contentmanagementsingledocumenttopicworkspacedata struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Contentmanagementsingledocumenttopicworkspacedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentmanagementsingledocumenttopicworkspacedata
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentmanagementsingledocumenttopicworkspacedata) UnmarshalJSON(b []byte) error {
	var ContentmanagementsingledocumenttopicworkspacedataMap map[string]interface{}
	err := json.Unmarshal(b, &ContentmanagementsingledocumenttopicworkspacedataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContentmanagementsingledocumenttopicworkspacedataMap["id"].(string); ok {
		o.Id = &Id
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contentmanagementsingledocumenttopicworkspacedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
