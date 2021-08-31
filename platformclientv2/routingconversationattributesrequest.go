package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingconversationattributesrequest
type Routingconversationattributesrequest struct { 
	// Priority - Priority to be updated on in-queue conversation. Range:[-25000000, 25000000]
	Priority *int `json:"priority,omitempty"`


	// SkillIds - Skills to be updated on in-queue conversation.
	SkillIds *[]string `json:"skillIds,omitempty"`


	// LanguageId - Language required on the in-queue conversation.
	LanguageId *string `json:"languageId,omitempty"`

}

func (o *Routingconversationattributesrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingconversationattributesrequest
	
	return json.Marshal(&struct { 
		Priority *int `json:"priority,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		*Alias
	}{ 
		Priority: o.Priority,
		
		SkillIds: o.SkillIds,
		
		LanguageId: o.LanguageId,
		Alias:    (*Alias)(o),
	})
}

func (o *Routingconversationattributesrequest) UnmarshalJSON(b []byte) error {
	var RoutingconversationattributesrequestMap map[string]interface{}
	err := json.Unmarshal(b, &RoutingconversationattributesrequestMap)
	if err != nil {
		return err
	}
	
	if Priority, ok := RoutingconversationattributesrequestMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if SkillIds, ok := RoutingconversationattributesrequestMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if LanguageId, ok := RoutingconversationattributesrequestMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Routingconversationattributesrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
