package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createuser
type Createuser struct { 
	// Name - User's full name
	Name *string `json:"name,omitempty"`


	// Department
	Department *string `json:"department,omitempty"`


	// Email - User's email and username
	Email *string `json:"email,omitempty"`


	// Addresses - Email addresses and phone numbers for this user
	Addresses *[]Contact `json:"addresses,omitempty"`


	// Title
	Title *string `json:"title,omitempty"`


	// Password - User's password
	Password *string `json:"password,omitempty"`


	// DivisionId - The division to which this user will belong
	DivisionId *string `json:"divisionId,omitempty"`


	// State - Optional initialized state of the user. If not specified, state will be Active if invites are sent, otherwise Inactive.
	State *string `json:"state,omitempty"`

}

func (o *Createuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createuser
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Department *string `json:"department,omitempty"`
		
		Email *string `json:"email,omitempty"`
		
		Addresses *[]Contact `json:"addresses,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Password *string `json:"password,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		State *string `json:"state,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Department: o.Department,
		
		Email: o.Email,
		
		Addresses: o.Addresses,
		
		Title: o.Title,
		
		Password: o.Password,
		
		DivisionId: o.DivisionId,
		
		State: o.State,
		Alias:    (*Alias)(o),
	})
}

func (o *Createuser) UnmarshalJSON(b []byte) error {
	var CreateuserMap map[string]interface{}
	err := json.Unmarshal(b, &CreateuserMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreateuserMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Department, ok := CreateuserMap["department"].(string); ok {
		o.Department = &Department
	}
    
	if Email, ok := CreateuserMap["email"].(string); ok {
		o.Email = &Email
	}
    
	if Addresses, ok := CreateuserMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if Title, ok := CreateuserMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Password, ok := CreateuserMap["password"].(string); ok {
		o.Password = &Password
	}
    
	if DivisionId, ok := CreateuserMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if State, ok := CreateuserMap["state"].(string); ok {
		o.State = &State
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
