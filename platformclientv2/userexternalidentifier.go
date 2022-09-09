package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userexternalidentifier - Defines a link between an External Identifier and Authority pair to a Entity Type and Entity Identifier pair. Represents the two way, one to one mapping of an External Authority or System of Record's identifier to a PureCloud entity. e.g. (ExternalId='05001',Authority='XyzCRM') to (entityType=user,entityId='8eb03b33-3acb-4bc1-a244-50b9b9f19495')
type Userexternalidentifier struct { 
	// AuthorityName - Authority or System of Record which owns the External Identifier
	AuthorityName *string `json:"authorityName,omitempty"`


	// ExternalKey - External Key
	ExternalKey *string `json:"externalKey,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Userexternalidentifier) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userexternalidentifier
	
	return json.Marshal(&struct { 
		AuthorityName *string `json:"authorityName,omitempty"`
		
		ExternalKey *string `json:"externalKey,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		AuthorityName: o.AuthorityName,
		
		ExternalKey: o.ExternalKey,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Userexternalidentifier) UnmarshalJSON(b []byte) error {
	var UserexternalidentifierMap map[string]interface{}
	err := json.Unmarshal(b, &UserexternalidentifierMap)
	if err != nil {
		return err
	}
	
	if AuthorityName, ok := UserexternalidentifierMap["authorityName"].(string); ok {
		o.AuthorityName = &AuthorityName
	}
    
	if ExternalKey, ok := UserexternalidentifierMap["externalKey"].(string); ok {
		o.ExternalKey = &ExternalKey
	}
    
	if SelfUri, ok := UserexternalidentifierMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userexternalidentifier) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
