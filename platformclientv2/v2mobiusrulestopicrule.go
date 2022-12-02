package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusrulestopicrule
type V2mobiusrulestopicrule struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// UserId
	UserId *string `json:"userId,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Notifications
	Notifications *[]V2mobiusrulestopicalertnotification `json:"notifications,omitempty"`


	// Conditions
	Conditions *V2mobiusrulestopiccondition `json:"conditions,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// InAlarm
	InAlarm *bool `json:"inAlarm,omitempty"`


	// Action
	Action *string `json:"action,omitempty"`

}

func (o *V2mobiusrulestopicrule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2mobiusrulestopicrule
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Notifications *[]V2mobiusrulestopicalertnotification `json:"notifications,omitempty"`
		
		Conditions *V2mobiusrulestopiccondition `json:"conditions,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		InAlarm *bool `json:"inAlarm,omitempty"`
		
		Action *string `json:"action,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		UserId: o.UserId,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		Notifications: o.Notifications,
		
		Conditions: o.Conditions,
		
		Enabled: o.Enabled,
		
		InAlarm: o.InAlarm,
		
		Action: o.Action,
		Alias:    (*Alias)(o),
	})
}

func (o *V2mobiusrulestopicrule) UnmarshalJSON(b []byte) error {
	var V2mobiusrulestopicruleMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusrulestopicruleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2mobiusrulestopicruleMap["id"].(map[string]interface{}); ok {
		IdString, _ := json.Marshal(Id)
		json.Unmarshal(IdString, &o.Id)
	}
	
	if UserId, ok := V2mobiusrulestopicruleMap["userId"].(map[string]interface{}); ok {
		UserIdString, _ := json.Marshal(UserId)
		json.Unmarshal(UserIdString, &o.UserId)
	}
	
	if Name, ok := V2mobiusrulestopicruleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := V2mobiusrulestopicruleMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Notifications, ok := V2mobiusrulestopicruleMap["notifications"].([]interface{}); ok {
		NotificationsString, _ := json.Marshal(Notifications)
		json.Unmarshal(NotificationsString, &o.Notifications)
	}
	
	if Conditions, ok := V2mobiusrulestopicruleMap["conditions"].(map[string]interface{}); ok {
		ConditionsString, _ := json.Marshal(Conditions)
		json.Unmarshal(ConditionsString, &o.Conditions)
	}
	
	if Enabled, ok := V2mobiusrulestopicruleMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if InAlarm, ok := V2mobiusrulestopicruleMap["inAlarm"].(bool); ok {
		o.InAlarm = &InAlarm
	}
    
	if Action, ok := V2mobiusrulestopicruleMap["action"].(string); ok {
		o.Action = &Action
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusrulestopicrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
