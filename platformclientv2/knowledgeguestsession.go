package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeguestsession
type Knowledgeguestsession struct { 
	// Id - Session ID.
	Id *string `json:"id,omitempty"`


	// App - The app where the session is started.
	App *Knowledgeguestsessionapp `json:"app,omitempty"`


	// CustomerId - An arbitrary ID for the customer starting the session. Used to track multiple sessions started by the same customer.
	CustomerId *string `json:"customerId,omitempty"`


	// PageUrl - URL of the page where the session is started.
	PageUrl *string `json:"pageUrl,omitempty"`


	// Contexts - The session contexts.
	Contexts *[]Knowledgeguestsessioncontext `json:"contexts,omitempty"`

}

func (o *Knowledgeguestsession) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeguestsession
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		App *Knowledgeguestsessionapp `json:"app,omitempty"`
		
		CustomerId *string `json:"customerId,omitempty"`
		
		PageUrl *string `json:"pageUrl,omitempty"`
		
		Contexts *[]Knowledgeguestsessioncontext `json:"contexts,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		App: o.App,
		
		CustomerId: o.CustomerId,
		
		PageUrl: o.PageUrl,
		
		Contexts: o.Contexts,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeguestsession) UnmarshalJSON(b []byte) error {
	var KnowledgeguestsessionMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeguestsessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgeguestsessionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if App, ok := KnowledgeguestsessionMap["app"].(map[string]interface{}); ok {
		AppString, _ := json.Marshal(App)
		json.Unmarshal(AppString, &o.App)
	}
	
	if CustomerId, ok := KnowledgeguestsessionMap["customerId"].(string); ok {
		o.CustomerId = &CustomerId
	}
    
	if PageUrl, ok := KnowledgeguestsessionMap["pageUrl"].(string); ok {
		o.PageUrl = &PageUrl
	}
    
	if Contexts, ok := KnowledgeguestsessionMap["contexts"].([]interface{}); ok {
		ContextsString, _ := json.Marshal(Contexts)
		json.Unmarshal(ContextsString, &o.Contexts)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeguestsession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
