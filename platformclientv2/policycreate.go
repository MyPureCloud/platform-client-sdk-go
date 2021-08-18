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

func (u *Policycreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Policycreate

	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	CreatedDate := new(string)
	if u.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(u.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		Name: u.Name,
		
		ModifiedDate: ModifiedDate,
		
		CreatedDate: CreatedDate,
		
		Order: u.Order,
		
		Description: u.Description,
		
		Enabled: u.Enabled,
		
		MediaPolicies: u.MediaPolicies,
		
		Conditions: u.Conditions,
		
		Actions: u.Actions,
		
		PolicyErrors: u.PolicyErrors,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Policycreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
