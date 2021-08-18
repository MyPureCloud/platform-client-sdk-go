package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditentityreference
type Auditentityreference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

func (u *Auditentityreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditentityreference

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Action *string `json:"action,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		SelfUri: u.SelfUri,
		
		VarType: u.VarType,
		
		Action: u.Action,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Auditentityreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
