package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistimportstatusimportstatus
type Contactlistimportstatusimportstatus struct { 
	// ImportState - current status of the import
	ImportState *string `json:"importState,omitempty"`


	// TotalRecords - total number of records to be imported
	TotalRecords *int `json:"totalRecords,omitempty"`


	// CompletedRecords - number of records finished importing
	CompletedRecords *int `json:"completedRecords,omitempty"`


	// PercentageComplete - percentage of records finished importing
	PercentageComplete *int `json:"percentageComplete,omitempty"`


	// FailureReason - if the import has failed, the reason for the failure
	FailureReason *string `json:"failureReason,omitempty"`

}

func (o *Contactlistimportstatusimportstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlistimportstatusimportstatus
	
	return json.Marshal(&struct { 
		ImportState *string `json:"importState,omitempty"`
		
		TotalRecords *int `json:"totalRecords,omitempty"`
		
		CompletedRecords *int `json:"completedRecords,omitempty"`
		
		PercentageComplete *int `json:"percentageComplete,omitempty"`
		
		FailureReason *string `json:"failureReason,omitempty"`
		*Alias
	}{ 
		ImportState: o.ImportState,
		
		TotalRecords: o.TotalRecords,
		
		CompletedRecords: o.CompletedRecords,
		
		PercentageComplete: o.PercentageComplete,
		
		FailureReason: o.FailureReason,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactlistimportstatusimportstatus) UnmarshalJSON(b []byte) error {
	var ContactlistimportstatusimportstatusMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlistimportstatusimportstatusMap)
	if err != nil {
		return err
	}
	
	if ImportState, ok := ContactlistimportstatusimportstatusMap["importState"].(string); ok {
		o.ImportState = &ImportState
	}
	
	if TotalRecords, ok := ContactlistimportstatusimportstatusMap["totalRecords"].(float64); ok {
		TotalRecordsInt := int(TotalRecords)
		o.TotalRecords = &TotalRecordsInt
	}
	
	if CompletedRecords, ok := ContactlistimportstatusimportstatusMap["completedRecords"].(float64); ok {
		CompletedRecordsInt := int(CompletedRecords)
		o.CompletedRecords = &CompletedRecordsInt
	}
	
	if PercentageComplete, ok := ContactlistimportstatusimportstatusMap["percentageComplete"].(float64); ok {
		PercentageCompleteInt := int(PercentageComplete)
		o.PercentageComplete = &PercentageCompleteInt
	}
	
	if FailureReason, ok := ContactlistimportstatusimportstatusMap["failureReason"].(string); ok {
		o.FailureReason = &FailureReason
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlistimportstatusimportstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
