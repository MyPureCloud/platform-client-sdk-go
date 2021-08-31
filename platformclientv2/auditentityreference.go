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

func (o *Auditentityreference) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		SelfUri: o.SelfUri,
		
		VarType: o.VarType,
		
		Action: o.Action,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditentityreference) UnmarshalJSON(b []byte) error {
	var AuditentityreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &AuditentityreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AuditentityreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := AuditentityreferenceMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if SelfUri, ok := AuditentityreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if VarType, ok := AuditentityreferenceMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Action, ok := AuditentityreferenceMap["action"].(string); ok {
		o.Action = &Action
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditentityreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
