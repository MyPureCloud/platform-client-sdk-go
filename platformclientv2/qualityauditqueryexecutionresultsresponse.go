package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Qualityauditqueryexecutionresultsresponse
type Qualityauditqueryexecutionresultsresponse struct { 
	// Id - Id of the audit query execution request.
	Id *string `json:"id,omitempty"`


	// PageSize - Number of results in a page.
	PageSize *int `json:"pageSize,omitempty"`


	// Cursor - Optional cursor to indicate where to resume the results.
	Cursor *string `json:"cursor,omitempty"`


	// Entities - List of audit messages.
	Entities *[]Qualityauditlogmessage `json:"entities,omitempty"`

}

func (o *Qualityauditqueryexecutionresultsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Qualityauditqueryexecutionresultsresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		Cursor *string `json:"cursor,omitempty"`
		
		Entities *[]Qualityauditlogmessage `json:"entities,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		PageSize: o.PageSize,
		
		Cursor: o.Cursor,
		
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Qualityauditqueryexecutionresultsresponse) UnmarshalJSON(b []byte) error {
	var QualityauditqueryexecutionresultsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &QualityauditqueryexecutionresultsresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QualityauditqueryexecutionresultsresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if PageSize, ok := QualityauditqueryexecutionresultsresponseMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if Cursor, ok := QualityauditqueryexecutionresultsresponseMap["cursor"].(string); ok {
		o.Cursor = &Cursor
	}
    
	if Entities, ok := QualityauditqueryexecutionresultsresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Qualityauditqueryexecutionresultsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
