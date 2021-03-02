package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Auditqueryexecutionresultsresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
