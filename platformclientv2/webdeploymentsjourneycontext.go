package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentsjourneycontext
type Webdeploymentsjourneycontext struct { 
	// Customer - Journey customer information. Used for linking the authenticated customer with the journey. 
	Customer *Journeycustomer `json:"customer,omitempty"`


	// CustomerSession - Contains the Journey System's customer session details.
	CustomerSession *Journeycustomersession `json:"customerSession,omitempty"`

}

func (o *Webdeploymentsjourneycontext) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentsjourneycontext
	
	return json.Marshal(&struct { 
		Customer *Journeycustomer `json:"customer,omitempty"`
		
		CustomerSession *Journeycustomersession `json:"customerSession,omitempty"`
		*Alias
	}{ 
		Customer: o.Customer,
		
		CustomerSession: o.CustomerSession,
		Alias:    (*Alias)(o),
	})
}

func (o *Webdeploymentsjourneycontext) UnmarshalJSON(b []byte) error {
	var WebdeploymentsjourneycontextMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentsjourneycontextMap)
	if err != nil {
		return err
	}
	
	if Customer, ok := WebdeploymentsjourneycontextMap["customer"].(map[string]interface{}); ok {
		CustomerString, _ := json.Marshal(Customer)
		json.Unmarshal(CustomerString, &o.Customer)
	}
	
	if CustomerSession, ok := WebdeploymentsjourneycontextMap["customerSession"].(map[string]interface{}); ok {
		CustomerSessionString, _ := json.Marshal(CustomerSession)
		json.Unmarshal(CustomerSessionString, &o.CustomerSession)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentsjourneycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
