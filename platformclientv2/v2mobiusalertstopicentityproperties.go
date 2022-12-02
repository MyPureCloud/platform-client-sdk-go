package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusalertstopicentityproperties
type V2mobiusalertstopicentityproperties struct { 
	// EntityType
	EntityType *string `json:"entityType,omitempty"`


	// UserDisplayName
	UserDisplayName *string `json:"userDisplayName,omitempty"`


	// GroupDisplayName
	GroupDisplayName *string `json:"groupDisplayName,omitempty"`


	// QueueDisplayName
	QueueDisplayName *string `json:"queueDisplayName,omitempty"`

}

func (o *V2mobiusalertstopicentityproperties) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2mobiusalertstopicentityproperties
	
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

func (o *V2mobiusalertstopicentityproperties) UnmarshalJSON(b []byte) error {
	var V2mobiusalertstopicentitypropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusalertstopicentitypropertiesMap)
	if err != nil {
		return err
	}
	
	if EntityType, ok := V2mobiusalertstopicentitypropertiesMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
    
	if UserDisplayName, ok := V2mobiusalertstopicentitypropertiesMap["userDisplayName"].(string); ok {
		o.UserDisplayName = &UserDisplayName
	}
    
	if GroupDisplayName, ok := V2mobiusalertstopicentitypropertiesMap["groupDisplayName"].(string); ok {
		o.GroupDisplayName = &GroupDisplayName
	}
    
	if QueueDisplayName, ok := V2mobiusalertstopicentitypropertiesMap["queueDisplayName"].(string); ok {
		o.QueueDisplayName = &QueueDisplayName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusalertstopicentityproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
