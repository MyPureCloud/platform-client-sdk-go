package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Failedobject
type Failedobject struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`

}

func (o *Failedobject) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Failedobject
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Version: o.Version,
		
		Name: o.Name,
		
		ErrorCode: o.ErrorCode,
		Alias:    (*Alias)(o),
	})
}

func (o *Failedobject) UnmarshalJSON(b []byte) error {
	var FailedobjectMap map[string]interface{}
	err := json.Unmarshal(b, &FailedobjectMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FailedobjectMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Version, ok := FailedobjectMap["version"].(string); ok {
		o.Version = &Version
	}
	
	if Name, ok := FailedobjectMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ErrorCode, ok := FailedobjectMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Failedobject) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
