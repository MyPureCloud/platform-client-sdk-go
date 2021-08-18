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

func (u *Auditqueryexecutionresultsresponse) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		PageSize: u.PageSize,
		
		Cursor: u.Cursor,
		
		Entities: u.Entities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Auditqueryexecutionresultsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
