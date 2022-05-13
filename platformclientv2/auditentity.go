package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditentity
type Auditentity struct { 
	// VarType - The type of the entity the action of this AuditEntity targeted.
	VarType *string `json:"type,omitempty"`


	// Id - The id of the entity the action of this AuditEntity targeted.
	Id *string `json:"id,omitempty"`


	// Name - The name of the entity the action of this AuditEntity targeted.
	Name *string `json:"name,omitempty"`


	// SelfUri - The selfUri for this entity.
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Auditentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditentity
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Id: o.Id,
		
		Name: o.Name,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditentity) UnmarshalJSON(b []byte) error {
	var AuditentityMap map[string]interface{}
	err := json.Unmarshal(b, &AuditentityMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := AuditentityMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Id, ok := AuditentityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AuditentityMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SelfUri, ok := AuditentityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Auditentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
