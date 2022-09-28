package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Vipbackup
type Vipbackup struct { 
	// Id - Unique ID for the selected VIP Backup option. For QUEUE this is the queueId and for USER this is the userId.
	Id *string `json:"id,omitempty"`


	// VarType - The type of VIP Backup to use.
	VarType *string `json:"type,omitempty"`

}

func (o *Vipbackup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Vipbackup
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Vipbackup) UnmarshalJSON(b []byte) error {
	var VipbackupMap map[string]interface{}
	err := json.Unmarshal(b, &VipbackupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := VipbackupMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if VarType, ok := VipbackupMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Vipbackup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
