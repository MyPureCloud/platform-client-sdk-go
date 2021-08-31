package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdate
type Wfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdate struct { 
	// Status
	Status *string `json:"status,omitempty"`

}

func (o *Wfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdate
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		*Alias
	}{ 
		Status: o.Status,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdate) UnmarshalJSON(b []byte) error {
	var WfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdateMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdateMap)
	if err != nil {
		return err
	}
	
	if Status, ok := WfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdateMap["status"].(string); ok {
		o.Status = &Status
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
