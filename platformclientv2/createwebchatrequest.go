package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createwebchatrequest
type Createwebchatrequest struct { 
	// QueueId - The ID of the queue to use for routing the chat conversation.
	QueueId *string `json:"queueId,omitempty"`


	// Provider - The name of the provider that is sourcing the web chat.
	Provider *string `json:"provider,omitempty"`


	// SkillIds - The list of skill ID's to use for routing.
	SkillIds *[]string `json:"skillIds,omitempty"`


	// LanguageId - The ID of the langauge to use for routing.
	LanguageId *string `json:"languageId,omitempty"`


	// Priority - The priority to assign to the conversation for routing.
	Priority *int `json:"priority,omitempty"`


	// Attributes - The list of attributes to associate with the customer participant.
	Attributes *map[string]string `json:"attributes,omitempty"`


	// CustomerName - The name of the customer participating in the web chat.
	CustomerName *string `json:"customerName,omitempty"`

}

func (o *Createwebchatrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createwebchatrequest
	
	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		CustomerName *string `json:"customerName,omitempty"`
		*Alias
	}{ 
		QueueId: o.QueueId,
		
		Provider: o.Provider,
		
		SkillIds: o.SkillIds,
		
		LanguageId: o.LanguageId,
		
		Priority: o.Priority,
		
		Attributes: o.Attributes,
		
		CustomerName: o.CustomerName,
		Alias:    (*Alias)(o),
	})
}

func (o *Createwebchatrequest) UnmarshalJSON(b []byte) error {
	var CreatewebchatrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatewebchatrequestMap)
	if err != nil {
		return err
	}
	
	if QueueId, ok := CreatewebchatrequestMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if Provider, ok := CreatewebchatrequestMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if SkillIds, ok := CreatewebchatrequestMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if LanguageId, ok := CreatewebchatrequestMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if Priority, ok := CreatewebchatrequestMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Attributes, ok := CreatewebchatrequestMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if CustomerName, ok := CreatewebchatrequestMap["customerName"].(string); ok {
		o.CustomerName = &CustomerName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createwebchatrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
