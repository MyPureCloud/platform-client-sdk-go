package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Licensedefinition
type Licensedefinition struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Permissions
	Permissions *Permissions `json:"permissions,omitempty"`


	// Prerequisites
	Prerequisites *[]Addressablelicensedefinition `json:"prerequisites,omitempty"`


	// Comprises
	Comprises *[]Licensedefinition `json:"comprises,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Licensedefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Licensedefinition
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Permissions *Permissions `json:"permissions,omitempty"`
		
		Prerequisites *[]Addressablelicensedefinition `json:"prerequisites,omitempty"`
		
		Comprises *[]Licensedefinition `json:"comprises,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Description: o.Description,
		
		Permissions: o.Permissions,
		
		Prerequisites: o.Prerequisites,
		
		Comprises: o.Comprises,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Licensedefinition) UnmarshalJSON(b []byte) error {
	var LicensedefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &LicensedefinitionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LicensedefinitionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Description, ok := LicensedefinitionMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Permissions, ok := LicensedefinitionMap["permissions"].(map[string]interface{}); ok {
		PermissionsString, _ := json.Marshal(Permissions)
		json.Unmarshal(PermissionsString, &o.Permissions)
	}
	
	if Prerequisites, ok := LicensedefinitionMap["prerequisites"].([]interface{}); ok {
		PrerequisitesString, _ := json.Marshal(Prerequisites)
		json.Unmarshal(PrerequisitesString, &o.Prerequisites)
	}
	
	if Comprises, ok := LicensedefinitionMap["comprises"].([]interface{}); ok {
		ComprisesString, _ := json.Marshal(Comprises)
		json.Unmarshal(ComprisesString, &o.Comprises)
	}
	
	if SelfUri, ok := LicensedefinitionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Licensedefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
