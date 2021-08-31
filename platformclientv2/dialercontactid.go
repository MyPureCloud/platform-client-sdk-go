package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactid
type Dialercontactid struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// ContactListId
	ContactListId *string `json:"contactListId,omitempty"`

}

func (o *Dialercontactid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactid
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ContactListId *string `json:"contactListId,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ContactListId: o.ContactListId,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialercontactid) UnmarshalJSON(b []byte) error {
	var DialercontactidMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactidMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DialercontactidMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if ContactListId, ok := DialercontactidMap["contactListId"].(string); ok {
		o.ContactListId = &ContactListId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
