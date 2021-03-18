package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Responseconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
