package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditqueryexecutionresultsresponse
type Auditqueryexecutionresultsresponse struct { 
	// Id - Id of the audit query execution request.
	Id *string `json:"id,omitempty"`


	// PageSize - Number of results in a page.
	PageSize *int `json:"pageSize,omitempty"`


	// Cursor - Optional cursor to indicate where to resume the results.
	Cursor *string `json:"cursor,omitempty"`


	// Entities - List of audit messages.
	Entities *[]Auditlogmessage `json:"entities,omitempty"`

}

func (o *Auditqueryexecutionresultsresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditqueryexecutionresultsresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		Cursor *string `json:"cursor,omitempty"`
		
		Entities *[]Auditlogmessage `json:"entities,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		PageSize: o.PageSize,
		
		Cursor: o.Cursor,
		
		Entities: o.Entities,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditqueryexecutionresultsresponse) UnmarshalJSON(b []byte) error {
	var AuditqueryexecutionresultsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &AuditqueryexecutionresultsresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AuditqueryexecutionresultsresponseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if PageSize, ok := AuditqueryexecutionresultsresponseMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if Cursor, ok := AuditqueryexecutionresultsresponseMap["cursor"].(string); ok {
		o.Cursor = &Cursor
	}
	
	if Entities, ok := AuditqueryexecutionresultsresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditqueryexecutionresultsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
