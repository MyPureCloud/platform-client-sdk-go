package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentsconfigtopicwebmessagingconfigchangeeventbody
type Webdeploymentsconfigtopicwebmessagingconfigchangeeventbody struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`

}

func (o *Webdeploymentsconfigtopicwebmessagingconfigchangeeventbody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentsconfigtopicwebmessagingconfigchangeeventbody
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Version: o.Version,
		
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Webdeploymentsconfigtopicwebmessagingconfigchangeeventbody) UnmarshalJSON(b []byte) error {
	var WebdeploymentsconfigtopicwebmessagingconfigchangeeventbodyMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentsconfigtopicwebmessagingconfigchangeeventbodyMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebdeploymentsconfigtopicwebmessagingconfigchangeeventbodyMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Version, ok := WebdeploymentsconfigtopicwebmessagingconfigchangeeventbodyMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if Status, ok := WebdeploymentsconfigtopicwebmessagingconfigchangeeventbodyMap["status"].(string); ok {
		o.Status = &Status
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentsconfigtopicwebmessagingconfigchangeeventbody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
