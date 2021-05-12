package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Auditrealtimequeryresultsresponse
type Auditrealtimequeryresultsresponse struct { 
	// Entities
	Entities *[]Auditlogmessage `json:"entities,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditrealtimequeryresultsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
