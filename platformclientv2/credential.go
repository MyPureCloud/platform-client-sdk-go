package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Credential
type Credential struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType - The type of credential.
	VarType *Credentialtype `json:"type,omitempty"`


	// CredentialFields
	CredentialFields *map[string]string `json:"credentialFields,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Credential) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Credential
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *Credentialtype `json:"type,omitempty"`
		
		CredentialFields *map[string]string `json:"credentialFields,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		CredentialFields: o.CredentialFields,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Credential) UnmarshalJSON(b []byte) error {
	var CredentialMap map[string]interface{}
	err := json.Unmarshal(b, &CredentialMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CredentialMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CredentialMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := CredentialMap["type"].(map[string]interface{}); ok {
		VarTypeString, _ := json.Marshal(VarType)
		json.Unmarshal(VarTypeString, &o.VarType)
	}
	
	if CredentialFields, ok := CredentialMap["credentialFields"].(map[string]interface{}); ok {
		CredentialFieldsString, _ := json.Marshal(CredentialFields)
		json.Unmarshal(CredentialFieldsString, &o.CredentialFields)
	}
	
	if SelfUri, ok := CredentialMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Credential) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
