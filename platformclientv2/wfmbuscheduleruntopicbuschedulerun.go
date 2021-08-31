package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduleruntopicbuschedulerun
type Wfmbuscheduleruntopicbuschedulerun struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// PercentComplete
	PercentComplete *float32 `json:"percentComplete,omitempty"`


	// IntradayRescheduling
	IntradayRescheduling *bool `json:"intradayRescheduling,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// WeekCount
	WeekCount *int `json:"weekCount,omitempty"`


	// Schedule
	Schedule *Wfmbuscheduleruntopicbuschedulereference `json:"schedule,omitempty"`


	// SchedulingCanceledBy
	SchedulingCanceledBy *Wfmbuscheduleruntopicuserreference `json:"schedulingCanceledBy,omitempty"`


	// SchedulingCompletedTime
	SchedulingCompletedTime *string `json:"schedulingCompletedTime,omitempty"`


	// MessageCount
	MessageCount *int `json:"messageCount,omitempty"`

}

func (o *Wfmbuscheduleruntopicbuschedulerun) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduleruntopicbuschedulerun
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		PercentComplete *float32 `json:"percentComplete,omitempty"`
		
		IntradayRescheduling *bool `json:"intradayRescheduling,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		Schedule *Wfmbuscheduleruntopicbuschedulereference `json:"schedule,omitempty"`
		
		SchedulingCanceledBy *Wfmbuscheduleruntopicuserreference `json:"schedulingCanceledBy,omitempty"`
		
		SchedulingCompletedTime *string `json:"schedulingCompletedTime,omitempty"`
		
		MessageCount *int `json:"messageCount,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		PercentComplete: o.PercentComplete,
		
		IntradayRescheduling: o.IntradayRescheduling,
		
		State: o.State,
		
		WeekCount: o.WeekCount,
		
		Schedule: o.Schedule,
		
		SchedulingCanceledBy: o.SchedulingCanceledBy,
		
		SchedulingCompletedTime: o.SchedulingCompletedTime,
		
		MessageCount: o.MessageCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuscheduleruntopicbuschedulerun) UnmarshalJSON(b []byte) error {
	var WfmbuscheduleruntopicbuschedulerunMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuscheduleruntopicbuschedulerunMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmbuscheduleruntopicbuschedulerunMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if PercentComplete, ok := WfmbuscheduleruntopicbuschedulerunMap["percentComplete"].(float64); ok {
		PercentCompleteFloat32 := float32(PercentComplete)
		o.PercentComplete = &PercentCompleteFloat32
	}
	
	if IntradayRescheduling, ok := WfmbuscheduleruntopicbuschedulerunMap["intradayRescheduling"].(bool); ok {
		o.IntradayRescheduling = &IntradayRescheduling
	}
	
	if State, ok := WfmbuscheduleruntopicbuschedulerunMap["state"].(string); ok {
		o.State = &State
	}
	
	if WeekCount, ok := WfmbuscheduleruntopicbuschedulerunMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if Schedule, ok := WfmbuscheduleruntopicbuschedulerunMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if SchedulingCanceledBy, ok := WfmbuscheduleruntopicbuschedulerunMap["schedulingCanceledBy"].(map[string]interface{}); ok {
		SchedulingCanceledByString, _ := json.Marshal(SchedulingCanceledBy)
		json.Unmarshal(SchedulingCanceledByString, &o.SchedulingCanceledBy)
	}
	
	if SchedulingCompletedTime, ok := WfmbuscheduleruntopicbuschedulerunMap["schedulingCompletedTime"].(string); ok {
		o.SchedulingCompletedTime = &SchedulingCompletedTime
	}
	
	if MessageCount, ok := WfmbuscheduleruntopicbuschedulerunMap["messageCount"].(float64); ok {
		MessageCountInt := int(MessageCount)
		o.MessageCount = &MessageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduleruntopicbuschedulerun) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
