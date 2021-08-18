package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dnclistimportstatusimportstatus
type Dnclistimportstatusimportstatus struct { 
	// ImportState
	ImportState *string `json:"importState,omitempty"`


	// TotalRecords
	TotalRecords *int `json:"totalRecords,omitempty"`


	// CompletedRecords
	CompletedRecords *int `json:"completedRecords,omitempty"`


	// PercentageComplete
	PercentageComplete *int `json:"percentageComplete,omitempty"`


	// FailureReason
	FailureReason *string `json:"failureReason,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dnclistimportstatusimportstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dnclistimportstatusimportstatus

	

	return json.Marshal(&struct { 
		ImportState *string `json:"importState,omitempty"`
		
		TotalRecords *int `json:"totalRecords,omitempty"`
		
		CompletedRecords *int `json:"completedRecords,omitempty"`
		
		PercentageComplete *int `json:"percentageComplete,omitempty"`
		
		FailureReason *string `json:"failureReason,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		ImportState: u.ImportState,
		
		TotalRecords: u.TotalRecords,
		
		CompletedRecords: u.CompletedRecords,
		
		PercentageComplete: u.PercentageComplete,
		
		FailureReason: u.FailureReason,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dnclistimportstatusimportstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
