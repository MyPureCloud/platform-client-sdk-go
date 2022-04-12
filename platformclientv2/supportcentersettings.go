package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportcentersettings - Settings concerning support center
type Supportcentersettings struct { 
	// Enabled - Whether or not support center is enabled
	Enabled *bool `json:"enabled,omitempty"`


	// KnowledgeBase - The knowledge base for support center
	KnowledgeBase *Addressableentityref `json:"knowledgeBase,omitempty"`


	// CustomMessages - Customizable display texts for support center
	CustomMessages *[]Supportcentercustommessage `json:"customMessages,omitempty"`


	// RouterType - Router type for support center
	RouterType *string `json:"routerType,omitempty"`


	// Screens - Available screens for the support center with its modules
	Screens *[]Supportcenterscreen `json:"screens,omitempty"`


	// EnabledCategories - Enabled article categories for support center
	EnabledCategories *[]Addressableentityref `json:"enabledCategories,omitempty"`


	// StyleSetting - Style attributes for support center
	StyleSetting *Supportcenterstylesetting `json:"styleSetting,omitempty"`

}

func (o *Supportcentersettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportcentersettings
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		KnowledgeBase *Addressableentityref `json:"knowledgeBase,omitempty"`
		
		CustomMessages *[]Supportcentercustommessage `json:"customMessages,omitempty"`
		
		RouterType *string `json:"routerType,omitempty"`
		
		Screens *[]Supportcenterscreen `json:"screens,omitempty"`
		
		EnabledCategories *[]Addressableentityref `json:"enabledCategories,omitempty"`
		
		StyleSetting *Supportcenterstylesetting `json:"styleSetting,omitempty"`
		*Alias
	}{ 
		Enabled: o.Enabled,
		
		KnowledgeBase: o.KnowledgeBase,
		
		CustomMessages: o.CustomMessages,
		
		RouterType: o.RouterType,
		
		Screens: o.Screens,
		
		EnabledCategories: o.EnabledCategories,
		
		StyleSetting: o.StyleSetting,
		Alias:    (*Alias)(o),
	})
}

func (o *Supportcentersettings) UnmarshalJSON(b []byte) error {
	var SupportcentersettingsMap map[string]interface{}
	err := json.Unmarshal(b, &SupportcentersettingsMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := SupportcentersettingsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if KnowledgeBase, ok := SupportcentersettingsMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if CustomMessages, ok := SupportcentersettingsMap["customMessages"].([]interface{}); ok {
		CustomMessagesString, _ := json.Marshal(CustomMessages)
		json.Unmarshal(CustomMessagesString, &o.CustomMessages)
	}
	
	if RouterType, ok := SupportcentersettingsMap["routerType"].(string); ok {
		o.RouterType = &RouterType
	}
	
	if Screens, ok := SupportcentersettingsMap["screens"].([]interface{}); ok {
		ScreensString, _ := json.Marshal(Screens)
		json.Unmarshal(ScreensString, &o.Screens)
	}
	
	if EnabledCategories, ok := SupportcentersettingsMap["enabledCategories"].([]interface{}); ok {
		EnabledCategoriesString, _ := json.Marshal(EnabledCategories)
		json.Unmarshal(EnabledCategoriesString, &o.EnabledCategories)
	}
	
	if StyleSetting, ok := SupportcentersettingsMap["styleSetting"].(map[string]interface{}); ok {
		StyleSettingString, _ := json.Marshal(StyleSetting)
		json.Unmarshal(StyleSettingString, &o.StyleSetting)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Supportcentersettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
