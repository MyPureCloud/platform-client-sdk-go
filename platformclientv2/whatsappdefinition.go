package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Whatsappdefinition - A WhatsApp messaging template definition as defined in the WhatsApp Business Manager
type Whatsappdefinition struct { 
	// Name - The messaging template name.
	Name *string `json:"name,omitempty"`


	// Namespace - The messaging template namespace.
	Namespace *string `json:"namespace,omitempty"`


	// Language - The messaging template language configured for this template. This is a WhatsApp specific value. For example, 'en_US'
	Language *string `json:"language,omitempty"`

}

func (o *Whatsappdefinition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Whatsappdefinition
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Namespace *string `json:"namespace,omitempty"`
		
		Language *string `json:"language,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Namespace: o.Namespace,
		
		Language: o.Language,
		Alias:    (*Alias)(o),
	})
}

func (o *Whatsappdefinition) UnmarshalJSON(b []byte) error {
	var WhatsappdefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &WhatsappdefinitionMap)
	if err != nil {
		return err
	}
	
	if Name, ok := WhatsappdefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Namespace, ok := WhatsappdefinitionMap["namespace"].(string); ok {
		o.Namespace = &Namespace
	}
    
	if Language, ok := WhatsappdefinitionMap["language"].(string); ok {
		o.Language = &Language
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Whatsappdefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
