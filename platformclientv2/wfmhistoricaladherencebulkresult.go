package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherencebulkresult
type Wfmhistoricaladherencebulkresult struct { 
	// StartDate - Beginning of the date range for this result in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End of the date range for this result in ISO-8601 format
	EndDate *time.Time `json:"endDate,omitempty"`


	// ManagementUnitId - The ID of the management unit for this result
	ManagementUnitId *string `json:"managementUnitId,omitempty"`


	// UserResults - The individual results for each user
	UserResults *[]Wfmhistoricaladherencebulkuserresult `json:"userResults,omitempty"`


	// LookupIdToSecondaryPresenceId - Map of secondary presence lookup ID to corresponding secondary presence ID
	LookupIdToSecondaryPresenceId *map[string]string `json:"lookupIdToSecondaryPresenceId,omitempty"`

}

func (o *Wfmhistoricaladherencebulkresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherencebulkresult
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		ManagementUnitId *string `json:"managementUnitId,omitempty"`
		
		UserResults *[]Wfmhistoricaladherencebulkuserresult `json:"userResults,omitempty"`
		
		LookupIdToSecondaryPresenceId *map[string]string `json:"lookupIdToSecondaryPresenceId,omitempty"`
		*Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		ManagementUnitId: o.ManagementUnitId,
		
		UserResults: o.UserResults,
		
		LookupIdToSecondaryPresenceId: o.LookupIdToSecondaryPresenceId,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaladherencebulkresult) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherencebulkresultMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherencebulkresultMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := WfmhistoricaladherencebulkresultMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := WfmhistoricaladherencebulkresultMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if ManagementUnitId, ok := WfmhistoricaladherencebulkresultMap["managementUnitId"].(string); ok {
		o.ManagementUnitId = &ManagementUnitId
	}
    
	if UserResults, ok := WfmhistoricaladherencebulkresultMap["userResults"].([]interface{}); ok {
		UserResultsString, _ := json.Marshal(UserResults)
		json.Unmarshal(UserResultsString, &o.UserResults)
	}
	
	if LookupIdToSecondaryPresenceId, ok := WfmhistoricaladherencebulkresultMap["lookupIdToSecondaryPresenceId"].(map[string]interface{}); ok {
		LookupIdToSecondaryPresenceIdString, _ := json.Marshal(LookupIdToSecondaryPresenceId)
		json.Unmarshal(LookupIdToSecondaryPresenceIdString, &o.LookupIdToSecondaryPresenceId)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
