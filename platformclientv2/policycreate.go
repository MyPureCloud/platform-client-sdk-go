package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Policycreate
type Policycreate struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The policy name.
	Name *string `json:"name,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// Order
	Order *int `json:"order,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// MediaPolicies - Conditions and actions per media type
	MediaPolicies *Mediapolicies `json:"mediaPolicies,omitempty"`


	// Conditions - Conditions
	Conditions *Policyconditions `json:"conditions,omitempty"`


	// Actions - Actions
	Actions *Policyactions `json:"actions,omitempty"`


	// PolicyErrors
	PolicyErrors *Policyerrors `json:"policyErrors,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Policycreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Policycreate
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		Order *int `json:"order,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		MediaPolicies *Mediapolicies `json:"mediaPolicies,omitempty"`
		
		Conditions *Policyconditions `json:"conditions,omitempty"`
		
		Actions *Policyactions `json:"actions,omitempty"`
		
		PolicyErrors *Policyerrors `json:"policyErrors,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ModifiedDate: ModifiedDate,
		
		CreatedDate: CreatedDate,
		
		Order: o.Order,
		
		Description: o.Description,
		
		Enabled: o.Enabled,
		
		MediaPolicies: o.MediaPolicies,
		
		Conditions: o.Conditions,
		
		Actions: o.Actions,
		
		PolicyErrors: o.PolicyErrors,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Policycreate) UnmarshalJSON(b []byte) error {
	var PolicycreateMap map[string]interface{}
	err := json.Unmarshal(b, &PolicycreateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PolicycreateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := PolicycreateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if modifiedDateString, ok := PolicycreateMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if createdDateString, ok := PolicycreateMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if Order, ok := PolicycreateMap["order"].(float64); ok {
		OrderInt := int(Order)
		o.Order = &OrderInt
	}
	
	if Description, ok := PolicycreateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Enabled, ok := PolicycreateMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if MediaPolicies, ok := PolicycreateMap["mediaPolicies"].(map[string]interface{}); ok {
		MediaPoliciesString, _ := json.Marshal(MediaPolicies)
		json.Unmarshal(MediaPoliciesString, &o.MediaPolicies)
	}
	
	if Conditions, ok := PolicycreateMap["conditions"].(map[string]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if Actions, ok := PolicycreateMap["actions"].(map[string]interface{}); ok {
		ActionsString, _ := json.Marshal(Actions)
		json.Unmarshal(ActionsString, &o.Actions)
	}
	
	if PolicyErrors, ok := PolicycreateMap["policyErrors"].(map[string]interface{}); ok {
		PolicyErrorsString, _ := json.Marshal(PolicyErrors)
		json.Unmarshal(PolicyErrorsString, &o.PolicyErrors)
	}
	
	if SelfUri, ok := PolicycreateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Policycreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
