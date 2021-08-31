package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmupdateagentdetailstopicwfmupdateagentdetailscomplete
type Wfmupdateagentdetailstopicwfmupdateagentdetailscomplete struct { 
	// Status
	Status *string `json:"status,omitempty"`

}

func (o *Wfmupdateagentdetailstopicwfmupdateagentdetailscomplete) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmupdateagentdetailstopicwfmupdateagentdetailscomplete
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmupdateagentdetailstopicwfmupdateagentdetailscomplete) UnmarshalJSON(b []byte) error {
	var WfmupdateagentdetailstopicwfmupdateagentdetailscompleteMap map[string]interface{}
	err := json.Unmarshal(b, &WfmupdateagentdetailstopicwfmupdateagentdetailscompleteMap)
	if err != nil {
		return err
	}
	
	if Status, ok := WfmupdateagentdetailstopicwfmupdateagentdetailscompleteMap["status"].(string); ok {
		o.Status = &Status
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmupdateagentdetailstopicwfmupdateagentdetailscomplete) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
