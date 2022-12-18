package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmtimeoffbalancejobtopictimeoffbalancejobnotification
type Wfmtimeoffbalancejobtopictimeoffbalancejobnotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Entities
	Entities *[]Wfmtimeoffbalancejobtopictimeoffbalance `json:"entities,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// VarError
	VarError *Wfmtimeoffbalancejobtopicerrorbody `json:"error,omitempty"`

}

func (o *Wfmtimeoffbalancejobtopictimeoffbalancejobnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmtimeoffbalancejobtopictimeoffbalancejobnotification
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Entities *[]Wfmtimeoffbalancejobtopictimeoffbalance `json:"entities,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		VarError *Wfmtimeoffbalancejobtopicerrorbody `json:"error,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Entities: o.Entities,
		
		Status: o.Status,
		
		VarError: o.VarError,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmtimeoffbalancejobtopictimeoffbalancejobnotification) UnmarshalJSON(b []byte) error {
	var WfmtimeoffbalancejobtopictimeoffbalancejobnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmtimeoffbalancejobtopictimeoffbalancejobnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmtimeoffbalancejobtopictimeoffbalancejobnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Entities, ok := WfmtimeoffbalancejobtopictimeoffbalancejobnotificationMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if Status, ok := WfmtimeoffbalancejobtopictimeoffbalancejobnotificationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if VarError, ok := WfmtimeoffbalancejobtopictimeoffbalancejobnotificationMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmtimeoffbalancejobtopictimeoffbalancejobnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
