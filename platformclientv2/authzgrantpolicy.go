package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Authzgrantpolicy
type Authzgrantpolicy struct { 
	// Actions
	Actions *[]string `json:"actions,omitempty"`


	// Condition
	Condition *string `json:"condition,omitempty"`


	// Domain
	Domain *string `json:"domain,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`

}

func (o *Authzgrantpolicy) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Authzgrantpolicy
	
	return json.Marshal(&struct { 
		Actions *[]string `json:"actions,omitempty"`
		
		Condition *string `json:"condition,omitempty"`
		
		Domain *string `json:"domain,omitempty"`
		
		EntityName *string `json:"entityName,omitempty"`
		*Alias
	}{ 
		Actions: o.Actions,
		
		Condition: o.Condition,
		
		Domain: o.Domain,
		
		EntityName: o.EntityName,
		Alias:    (*Alias)(o),
	})
}

func (o *Authzgrantpolicy) UnmarshalJSON(b []byte) error {
	var AuthzgrantpolicyMap map[string]interface{}
	err := json.Unmarshal(b, &AuthzgrantpolicyMap)
	if err != nil {
		return err
	}
	
	if Actions, ok := AuthzgrantpolicyMap["actions"].([]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if Condition, ok := AuthzgrantpolicyMap["condition"].(string); ok {
		o.Condition = &Condition
	}
	
	if Domain, ok := AuthzgrantpolicyMap["domain"].(string); ok {
		o.Domain = &Domain
	}
	
	if EntityName, ok := AuthzgrantpolicyMap["entityName"].(string); ok {
		o.EntityName = &EntityName
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Authzgrantpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
