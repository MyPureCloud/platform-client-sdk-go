package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusrulestopicentityproperties
type V2mobiusrulestopicentityproperties struct { 
	// EntityType
	EntityType *string `json:"entityType,omitempty"`


	// UserDisplayName
	UserDisplayName *string `json:"userDisplayName,omitempty"`


	// GroupDisplayName
	GroupDisplayName *string `json:"groupDisplayName,omitempty"`


	// QueueDisplayName
	QueueDisplayName *string `json:"queueDisplayName,omitempty"`

}

func (o *V2mobiusrulestopicentityproperties) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2mobiusrulestopicentityproperties
	
	return json.Marshal(&struct { 
		EntityType *string `json:"entityType,omitempty"`
		
		UserDisplayName *string `json:"userDisplayName,omitempty"`
		
		GroupDisplayName *string `json:"groupDisplayName,omitempty"`
		
		QueueDisplayName *string `json:"queueDisplayName,omitempty"`
		*Alias
	}{ 
		EntityType: o.EntityType,
		
		UserDisplayName: o.UserDisplayName,
		
		GroupDisplayName: o.GroupDisplayName,
		
		QueueDisplayName: o.QueueDisplayName,
		Alias:    (*Alias)(o),
	})
}

func (o *V2mobiusrulestopicentityproperties) UnmarshalJSON(b []byte) error {
	var V2mobiusrulestopicentitypropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusrulestopicentitypropertiesMap)
	if err != nil {
		return err
	}
	
	if EntityType, ok := V2mobiusrulestopicentitypropertiesMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
    
	if UserDisplayName, ok := V2mobiusrulestopicentitypropertiesMap["userDisplayName"].(string); ok {
		o.UserDisplayName = &UserDisplayName
	}
    
	if GroupDisplayName, ok := V2mobiusrulestopicentitypropertiesMap["groupDisplayName"].(string); ok {
		o.GroupDisplayName = &GroupDisplayName
	}
    
	if QueueDisplayName, ok := V2mobiusrulestopicentitypropertiesMap["queueDisplayName"].(string); ok {
		o.QueueDisplayName = &QueueDisplayName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusrulestopicentityproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
