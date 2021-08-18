package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkresponseresultexternalorganizationentity
type Bulkresponseresultexternalorganizationentity struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Externalorganization `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrorentity `json:"error,omitempty"`

}

func (u *Bulkresponseresultexternalorganizationentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkresponseresultexternalorganizationentity

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Success *bool `json:"success,omitempty"`
		
		Entity *Externalorganization `json:"entity,omitempty"`
		
		VarError *Bulkerrorentity `json:"error,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Success: u.Success,
		
		Entity: u.Entity,
		
		VarError: u.VarError,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Bulkresponseresultexternalorganizationentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
