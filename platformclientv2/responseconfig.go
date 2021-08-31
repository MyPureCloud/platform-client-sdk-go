package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Responseconfig - Defines response components of the Action Request.
type Responseconfig struct { 
	// TranslationMap - Map 'attribute name' and 'JSON path' pairs used to extract data from REST response.
	TranslationMap *map[string]string `json:"translationMap,omitempty"`


	// TranslationMapDefaults - Map 'attribute name' and 'default value' pairs used as fallback values if JSON path extraction fails for specified key.
	TranslationMapDefaults *map[string]string `json:"translationMapDefaults,omitempty"`


	// SuccessTemplate - Velocity template to build response to return from Action.
	SuccessTemplate *string `json:"successTemplate,omitempty"`


	// SuccessTemplateUri - URI to retrieve success template.
	SuccessTemplateUri *string `json:"successTemplateUri,omitempty"`

}

func (o *Responseconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responseconfig
	
	return json.Marshal(&struct { 
		TranslationMap *map[string]string `json:"translationMap,omitempty"`
		
		TranslationMapDefaults *map[string]string `json:"translationMapDefaults,omitempty"`
		
		SuccessTemplate *string `json:"successTemplate,omitempty"`
		
		SuccessTemplateUri *string `json:"successTemplateUri,omitempty"`
		*Alias
	}{ 
		TranslationMap: o.TranslationMap,
		
		TranslationMapDefaults: o.TranslationMapDefaults,
		
		SuccessTemplate: o.SuccessTemplate,
		
		SuccessTemplateUri: o.SuccessTemplateUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Responseconfig) UnmarshalJSON(b []byte) error {
	var ResponseconfigMap map[string]interface{}
	err := json.Unmarshal(b, &ResponseconfigMap)
	if err != nil {
		return err
	}
	
	if TranslationMap, ok := ResponseconfigMap["translationMap"].(map[string]interface{}); ok {
		TranslationMapString, _ := json.Marshal(TranslationMap)
		json.Unmarshal(TranslationMapString, &o.TranslationMap)
	}
	
	if TranslationMapDefaults, ok := ResponseconfigMap["translationMapDefaults"].(map[string]interface{}); ok {
		TranslationMapDefaultsString, _ := json.Marshal(TranslationMapDefaults)
		json.Unmarshal(TranslationMapDefaultsString, &o.TranslationMapDefaults)
	}
	
	if SuccessTemplate, ok := ResponseconfigMap["successTemplate"].(string); ok {
		o.SuccessTemplate = &SuccessTemplate
	}
	
	if SuccessTemplateUri, ok := ResponseconfigMap["successTemplateUri"].(string); ok {
		o.SuccessTemplateUri = &SuccessTemplateUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Responseconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
