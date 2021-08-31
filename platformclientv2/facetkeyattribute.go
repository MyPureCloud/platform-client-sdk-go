package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Facetkeyattribute
type Facetkeyattribute struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Count
	Count *int `json:"count,omitempty"`

}

func (o *Facetkeyattribute) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facetkeyattribute
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Count *int `json:"count,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Count: o.Count,
		Alias:    (*Alias)(o),
	})
}

func (o *Facetkeyattribute) UnmarshalJSON(b []byte) error {
	var FacetkeyattributeMap map[string]interface{}
	err := json.Unmarshal(b, &FacetkeyattributeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FacetkeyattributeMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := FacetkeyattributeMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Count, ok := FacetkeyattributeMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Facetkeyattribute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
