package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker
type Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker struct { 
	// TimeOffRequestId
	TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`


	// ManagementUnitDate
	ManagementUnitDate *string `json:"managementUnitDate,omitempty"`


	// ActivityCodeId
	ActivityCodeId *string `json:"activityCodeId,omitempty"`


	// IsPaid
	IsPaid *bool `json:"isPaid,omitempty"`


	// LengthInMinutes
	LengthInMinutes *int `json:"lengthInMinutes,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Paid
	Paid *bool `json:"paid,omitempty"`

}

func (o *Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker
	
	return json.Marshal(&struct { 
		TimeOffRequestId *string `json:"timeOffRequestId,omitempty"`
		
		ManagementUnitDate *string `json:"managementUnitDate,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		IsPaid *bool `json:"isPaid,omitempty"`
		
		LengthInMinutes *int `json:"lengthInMinutes,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Paid *bool `json:"paid,omitempty"`
		*Alias
	}{ 
		TimeOffRequestId: o.TimeOffRequestId,
		
		ManagementUnitDate: o.ManagementUnitDate,
		
		ActivityCodeId: o.ActivityCodeId,
		
		IsPaid: o.IsPaid,
		
		LengthInMinutes: o.LengthInMinutes,
		
		Description: o.Description,
		
		Paid: o.Paid,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker) UnmarshalJSON(b []byte) error {
	var WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap map[string]interface{}
	err := json.Unmarshal(b, &WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap)
	if err != nil {
		return err
	}
	
	if TimeOffRequestId, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["timeOffRequestId"].(string); ok {
		o.TimeOffRequestId = &TimeOffRequestId
	}
    
	if ManagementUnitDate, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["managementUnitDate"].(string); ok {
		o.ManagementUnitDate = &ManagementUnitDate
	}
    
	if ActivityCodeId, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if IsPaid, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["isPaid"].(bool); ok {
		o.IsPaid = &IsPaid
	}
    
	if LengthInMinutes, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["lengthInMinutes"].(float64); ok {
		LengthInMinutesInt := int(LengthInMinutes)
		o.LengthInMinutes = &LengthInMinutesInt
	}
	
	if Description, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Paid, ok := WfmagentscheduleupdatetopicwfmfulldaytimeoffmarkerMap["paid"].(bool); ok {
		o.Paid = &Paid
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmagentscheduleupdatetopicwfmfulldaytimeoffmarker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
