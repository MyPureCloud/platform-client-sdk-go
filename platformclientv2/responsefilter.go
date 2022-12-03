package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Responsefilter - Used to filter response queries
type Responsefilter struct { 
	// Name - Field to filter on. Allowed values are 'name', 'libraryId', 'text.contentType', 'messagingTemplate' and 'responseType'
	Name *string `json:"name,omitempty"`


	// Operator - Filter operation: IN, EQUALS, NOTEQUALS.
	Operator *string `json:"operator,omitempty"`


	// Values - Values to filter on. If name is 'responseType' then allowed values are 'CampaignSmsTemplate', 'CampaignEmailTemplate', 'Footer' and 'Signature'
	Values *[]string `json:"values,omitempty"`

}

func (o *Responsefilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responsefilter
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Operator: o.Operator,
		
		Values: o.Values,
		Alias:    (*Alias)(o),
	})
}

func (o *Responsefilter) UnmarshalJSON(b []byte) error {
	var ResponsefilterMap map[string]interface{}
	err := json.Unmarshal(b, &ResponsefilterMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ResponsefilterMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Operator, ok := ResponsefilterMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Values, ok := ResponsefilterMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Responsefilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
