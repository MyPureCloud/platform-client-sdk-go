package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userconversationseventmediasummarydetail
type Userconversationseventmediasummarydetail struct { 
	// Active
	Active *int `json:"active,omitempty"`


	// Acw
	Acw *int `json:"acw,omitempty"`

}

func (o *Userconversationseventmediasummarydetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userconversationseventmediasummarydetail
	
	return json.Marshal(&struct { 
		Active *int `json:"active,omitempty"`
		
		Acw *int `json:"acw,omitempty"`
		*Alias
	}{ 
		Active: o.Active,
		
		Acw: o.Acw,
		Alias:    (*Alias)(o),
	})
}

func (o *Userconversationseventmediasummarydetail) UnmarshalJSON(b []byte) error {
	var UserconversationseventmediasummarydetailMap map[string]interface{}
	err := json.Unmarshal(b, &UserconversationseventmediasummarydetailMap)
	if err != nil {
		return err
	}
	
	if Active, ok := UserconversationseventmediasummarydetailMap["active"].(float64); ok {
		ActiveInt := int(Active)
		o.Active = &ActiveInt
	}
	
	if Acw, ok := UserconversationseventmediasummarydetailMap["acw"].(float64); ok {
		AcwInt := int(Acw)
		o.Acw = &AcwInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userconversationseventmediasummarydetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
