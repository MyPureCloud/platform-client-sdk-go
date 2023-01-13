package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaladherencebulkitem
type Wfmhistoricaladherencebulkitem struct { 
	// ManagementUnitId - The ID of the management unit to query
	ManagementUnitId *string `json:"managementUnitId,omitempty"`


	// StartDate - Beginning of the date range to query in ISO-8601 format
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - End of the date range to query in ISO-8601 format
	EndDate *time.Time `json:"endDate,omitempty"`


	// UserIds - The IDs of the users to query. If not included, will query every user in the management unit
	UserIds *[]string `json:"userIds,omitempty"`


	// IncludeExceptions - Whether user exceptions should be returned as part of the results. If not included, will default to false
	IncludeExceptions *bool `json:"includeExceptions,omitempty"`

}

func (o *Wfmhistoricaladherencebulkitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaladherencebulkitem
	
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
		ManagementUnitId *string `json:"managementUnitId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		UserIds *[]string `json:"userIds,omitempty"`
		
		IncludeExceptions *bool `json:"includeExceptions,omitempty"`
		*Alias
	}{ 
		ManagementUnitId: o.ManagementUnitId,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		UserIds: o.UserIds,
		
		IncludeExceptions: o.IncludeExceptions,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaladherencebulkitem) UnmarshalJSON(b []byte) error {
	var WfmhistoricaladherencebulkitemMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaladherencebulkitemMap)
	if err != nil {
		return err
	}
	
	if ManagementUnitId, ok := WfmhistoricaladherencebulkitemMap["managementUnitId"].(string); ok {
		o.ManagementUnitId = &ManagementUnitId
	}
    
	if startDateString, ok := WfmhistoricaladherencebulkitemMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := WfmhistoricaladherencebulkitemMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if UserIds, ok := WfmhistoricaladherencebulkitemMap["userIds"].([]interface{}); ok {
		UserIdsString, _ := json.Marshal(UserIds)
		json.Unmarshal(UserIdsString, &o.UserIds)
	}
	
	if IncludeExceptions, ok := WfmhistoricaladherencebulkitemMap["includeExceptions"].(bool); ok {
		o.IncludeExceptions = &IncludeExceptions
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherencebulkitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
