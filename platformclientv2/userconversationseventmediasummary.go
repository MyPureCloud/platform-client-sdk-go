package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userconversationseventmediasummary
type Userconversationseventmediasummary struct { 
	// ContactCenter
	ContactCenter *Userconversationseventmediasummarydetail `json:"contactCenter,omitempty"`


	// Enterprise
	Enterprise *Userconversationseventmediasummarydetail `json:"enterprise,omitempty"`

}

func (o *Userconversationseventmediasummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userconversationseventmediasummary
	
	return json.Marshal(&struct { 
		ContactCenter *Userconversationseventmediasummarydetail `json:"contactCenter,omitempty"`
		
		Enterprise *Userconversationseventmediasummarydetail `json:"enterprise,omitempty"`
		*Alias
	}{ 
		ContactCenter: o.ContactCenter,
		
		Enterprise: o.Enterprise,
		Alias:    (*Alias)(o),
	})
}

func (o *Userconversationseventmediasummary) UnmarshalJSON(b []byte) error {
	var UserconversationseventmediasummaryMap map[string]interface{}
	err := json.Unmarshal(b, &UserconversationseventmediasummaryMap)
	if err != nil {
		return err
	}
	
	if ContactCenter, ok := UserconversationseventmediasummaryMap["contactCenter"].(map[string]interface{}); ok {
		ContactCenterString, _ := json.Marshal(ContactCenter)
		json.Unmarshal(ContactCenterString, &o.ContactCenter)
	}
	
	if Enterprise, ok := UserconversationseventmediasummaryMap["enterprise"].(map[string]interface{}); ok {
		EnterpriseString, _ := json.Marshal(Enterprise)
		json.Unmarshal(EnterpriseString, &o.Enterprise)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userconversationseventmediasummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
