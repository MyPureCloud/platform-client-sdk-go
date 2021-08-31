package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatedraftinput - Definition of an Action Draft to be created or updated.
type Updatedraftinput struct { 
	// Category - Category of action, Can be up to 256 characters long
	Category *string `json:"category,omitempty"`


	// Name - Name of action, Can be up to 256 characters long
	Name *string `json:"name,omitempty"`


	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`


	// Contract - Action contract
	Contract *Actioncontractinput `json:"contract,omitempty"`


	// Secure - Indication of whether or not the action is designed to accept sensitive data
	Secure *bool `json:"secure,omitempty"`


	// Version - Version of current Draft
	Version *int `json:"version,omitempty"`

}

func (o *Updatedraftinput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatedraftinput
	
	return json.Marshal(&struct { 
		Category *string `json:"category,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Config *Actionconfig `json:"config,omitempty"`
		
		Contract *Actioncontractinput `json:"contract,omitempty"`
		
		Secure *bool `json:"secure,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		Category: o.Category,
		
		Name: o.Name,
		
		Config: o.Config,
		
		Contract: o.Contract,
		
		Secure: o.Secure,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Updatedraftinput) UnmarshalJSON(b []byte) error {
	var UpdatedraftinputMap map[string]interface{}
	err := json.Unmarshal(b, &UpdatedraftinputMap)
	if err != nil {
		return err
	}
	
	if Category, ok := UpdatedraftinputMap["category"].(string); ok {
		o.Category = &Category
	}
	
	if Name, ok := UpdatedraftinputMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Config, ok := UpdatedraftinputMap["config"].(map[string]interface{}); ok {
		ConfigString, _ := json.Marshal(Config)
		json.Unmarshal(ConfigString, &o.Config)
	}
	
	if Contract, ok := UpdatedraftinputMap["contract"].(map[string]interface{}); ok {
		ContractString, _ := json.Marshal(Contract)
		json.Unmarshal(ContractString, &o.Contract)
	}
	
	if Secure, ok := UpdatedraftinputMap["secure"].(bool); ok {
		o.Secure = &Secure
	}
	
	if Version, ok := UpdatedraftinputMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updatedraftinput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
