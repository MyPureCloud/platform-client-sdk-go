package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentmanagementworkspacedocumentstopicworkspacedata
type Contentmanagementworkspacedocumentstopicworkspacedata struct { 
	// Id
	Id *string `json:"id,omitempty"`

}

func (o *Contentmanagementworkspacedocumentstopicworkspacedata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentmanagementworkspacedocumentstopicworkspacedata
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Contentmanagementworkspacedocumentstopicworkspacedata) UnmarshalJSON(b []byte) error {
	var ContentmanagementworkspacedocumentstopicworkspacedataMap map[string]interface{}
	err := json.Unmarshal(b, &ContentmanagementworkspacedocumentstopicworkspacedataMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContentmanagementworkspacedocumentstopicworkspacedataMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentmanagementworkspacedocumentstopicworkspacedata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
