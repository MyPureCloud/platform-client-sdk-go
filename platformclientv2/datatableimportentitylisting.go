package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Datatableimportentitylisting
type Datatableimportentitylisting struct { 
	// Entities
	Entities *[]Datatableimportjob `json:"entities,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

func (u *Datatableimportentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Datatableimportentitylisting

	

	return json.Marshal(&struct { 
		Entities *[]Datatableimportjob `json:"entities,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		
		PageSize: u.PageSize,
		
		PageNumber: u.PageNumber,
		
		Total: u.Total,
		
		PageCount: u.PageCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Datatableimportentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
