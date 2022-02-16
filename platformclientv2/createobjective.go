package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createobjective
type Createobjective struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// TemplateId - The id of this objective's base template
	TemplateId *string `json:"templateId,omitempty"`


	// Zones - Objective zone specifies min,max points and values for the associated metric
	Zones *[]Objectivezone `json:"zones,omitempty"`


	// Enabled - A flag for whether this objective is enabled for the related metric
	Enabled *bool `json:"enabled,omitempty"`


	// TopicIds - A list of topic ids for detected topic metrics
	TopicIds *[]string `json:"topicIds,omitempty"`


	// MediaTypes - A list of media types for the metric
	MediaTypes *[]string `json:"mediaTypes,omitempty"`


	// QueueIds - A list of queue ids for the metric
	QueueIds *[]string `json:"queueIds,omitempty"`


	// TopicIdsFilterType - A filter type for topic Ids. It's only used for objectives with topicIds. Default filter behavior is \"or\".
	TopicIdsFilterType *string `json:"topicIdsFilterType,omitempty"`


	// DateStart - start date of the objective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStart *time.Time `json:"dateStart,omitempty"`

}

func (o *Createobjective) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createobjective
	
	DateStart := new(string)
	if o.DateStart != nil {
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%d")
	} else {
		DateStart = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		TemplateId *string `json:"templateId,omitempty"`
		
		Zones *[]Objectivezone `json:"zones,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		TopicIds *[]string `json:"topicIds,omitempty"`
		
		MediaTypes *[]string `json:"mediaTypes,omitempty"`
		
		QueueIds *[]string `json:"queueIds,omitempty"`
		
		TopicIdsFilterType *string `json:"topicIdsFilterType,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		TemplateId: o.TemplateId,
		
		Zones: o.Zones,
		
		Enabled: o.Enabled,
		
		TopicIds: o.TopicIds,
		
		MediaTypes: o.MediaTypes,
		
		QueueIds: o.QueueIds,
		
		TopicIdsFilterType: o.TopicIdsFilterType,
		
		DateStart: DateStart,
		Alias:    (*Alias)(o),
	})
}

func (o *Createobjective) UnmarshalJSON(b []byte) error {
	var CreateobjectiveMap map[string]interface{}
	err := json.Unmarshal(b, &CreateobjectiveMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CreateobjectiveMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if TemplateId, ok := CreateobjectiveMap["templateId"].(string); ok {
		o.TemplateId = &TemplateId
	}
	
	if Zones, ok := CreateobjectiveMap["zones"].([]interface{}); ok {
		ZonesString, _ := json.Marshal(Zones)
		json.Unmarshal(ZonesString, &o.Zones)
	}
	
	if Enabled, ok := CreateobjectiveMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if TopicIds, ok := CreateobjectiveMap["topicIds"].([]interface{}); ok {
		TopicIdsString, _ := json.Marshal(TopicIds)
		json.Unmarshal(TopicIdsString, &o.TopicIds)
	}
	
	if MediaTypes, ok := CreateobjectiveMap["mediaTypes"].([]interface{}); ok {
		MediaTypesString, _ := json.Marshal(MediaTypes)
		json.Unmarshal(MediaTypesString, &o.MediaTypes)
	}
	
	if QueueIds, ok := CreateobjectiveMap["queueIds"].([]interface{}); ok {
		QueueIdsString, _ := json.Marshal(QueueIds)
		json.Unmarshal(QueueIdsString, &o.QueueIds)
	}
	
	if TopicIdsFilterType, ok := CreateobjectiveMap["topicIdsFilterType"].(string); ok {
		o.TopicIdsFilterType = &TopicIdsFilterType
	}
	
	if dateStartString, ok := CreateobjectiveMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02", dateStartString)
		o.DateStart = &DateStart
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createobjective) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
