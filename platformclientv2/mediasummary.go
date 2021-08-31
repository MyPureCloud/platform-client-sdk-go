package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediasummary
type Mediasummary struct { 
	// ContactCenter
	ContactCenter *Mediasummarydetail `json:"contactCenter,omitempty"`


	// Enterprise
	Enterprise *Mediasummarydetail `json:"enterprise,omitempty"`

}

func (o *Mediasummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediasummary
	
	return json.Marshal(&struct { 
		ContactCenter *Mediasummarydetail `json:"contactCenter,omitempty"`
		
		Enterprise *Mediasummarydetail `json:"enterprise,omitempty"`
		*Alias
	}{ 
		ContactCenter: o.ContactCenter,
		
		Enterprise: o.Enterprise,
		Alias:    (*Alias)(o),
	})
}

func (o *Mediasummary) UnmarshalJSON(b []byte) error {
	var MediasummaryMap map[string]interface{}
	err := json.Unmarshal(b, &MediasummaryMap)
	if err != nil {
		return err
	}
	
	if ContactCenter, ok := MediasummaryMap["contactCenter"].(map[string]interface{}); ok {
		ContactCenterString, _ := json.Marshal(ContactCenter)
		json.Unmarshal(ContactCenterString, &o.ContactCenter)
	}
	
	if Enterprise, ok := MediasummaryMap["enterprise"].(map[string]interface{}); ok {
		EnterpriseString, _ := json.Marshal(Enterprise)
		json.Unmarshal(EnterpriseString, &o.Enterprise)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Mediasummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
