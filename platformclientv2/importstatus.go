package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Importstatus
type Importstatus struct { 
	// State - current status of the import
	State *string `json:"state,omitempty"`


	// TotalRecords - total number of records to be imported
	TotalRecords *int `json:"totalRecords,omitempty"`


	// CompletedRecords - number of records finished importing
	CompletedRecords *int `json:"completedRecords,omitempty"`


	// PercentComplete - percentage of records finished importing
	PercentComplete *int `json:"percentComplete,omitempty"`


	// FailureReason - if the import has failed, the reason for the failure
	FailureReason *string `json:"failureReason,omitempty"`

}

func (o *Importstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Importstatus
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		TotalRecords *int `json:"totalRecords,omitempty"`
		
		CompletedRecords *int `json:"completedRecords,omitempty"`
		
		PercentComplete *int `json:"percentComplete,omitempty"`
		
		FailureReason *string `json:"failureReason,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		TotalRecords: o.TotalRecords,
		
		CompletedRecords: o.CompletedRecords,
		
		PercentComplete: o.PercentComplete,
		
		FailureReason: o.FailureReason,
		Alias:    (*Alias)(o),
	})
}

func (o *Importstatus) UnmarshalJSON(b []byte) error {
	var ImportstatusMap map[string]interface{}
	err := json.Unmarshal(b, &ImportstatusMap)
	if err != nil {
		return err
	}
	
	if State, ok := ImportstatusMap["state"].(string); ok {
		o.State = &State
	}
    
	if TotalRecords, ok := ImportstatusMap["totalRecords"].(float64); ok {
		TotalRecordsInt := int(TotalRecords)
		o.TotalRecords = &TotalRecordsInt
	}
	
	if CompletedRecords, ok := ImportstatusMap["completedRecords"].(float64); ok {
		CompletedRecordsInt := int(CompletedRecords)
		o.CompletedRecords = &CompletedRecordsInt
	}
	
	if PercentComplete, ok := ImportstatusMap["percentComplete"].(float64); ok {
		PercentCompleteInt := int(PercentComplete)
		o.PercentComplete = &PercentCompleteInt
	}
	
	if FailureReason, ok := ImportstatusMap["failureReason"].(string); ok {
		o.FailureReason = &FailureReason
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Importstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
