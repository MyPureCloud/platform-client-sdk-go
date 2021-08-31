package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatenotificationresponse
type Updatenotificationresponse struct { 
	// MutableGroupId - The mutableGroupId of the notification
	MutableGroupId *string `json:"mutableGroupId,omitempty"`


	// Id - The id of the notification for mapping the potentially new mutableGroupId
	Id *string `json:"id,omitempty"`

}

func (o *Updatenotificationresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatenotificationresponse
	
	return json.Marshal(&struct { 
		MutableGroupId *string `json:"mutableGroupId,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		MutableGroupId: o.MutableGroupId,
		
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Updatenotificationresponse) UnmarshalJSON(b []byte) error {
	var UpdatenotificationresponseMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatenotificationresponseMap)
	if err != nil {
		return err
	}
	
	if MutableGroupId, ok := UpdatenotificationresponseMap["mutableGroupId"].(string); ok {
		o.MutableGroupId = &MutableGroupId
	}
	
	if Id, ok := UpdatenotificationresponseMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatenotificationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
