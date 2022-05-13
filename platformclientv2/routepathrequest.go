package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routepathrequest
type Routepathrequest struct { 
	// QueueId - The ID of the queue to associate with the route path
	QueueId *string `json:"queueId,omitempty"`


	// MediaType - The media type of the given queue to associate with the route path
	MediaType *string `json:"mediaType,omitempty"`


	// LanguageId - The ID of the language to associate with the route path
	LanguageId *string `json:"languageId,omitempty"`


	// SkillIds - The set of skill IDs to associate with the route path
	SkillIds *[]string `json:"skillIds,omitempty"`


	// SourcePlanningGroup - The planning group from which to take route paths. This property is only needed if a route path already exists in another planning group.Note that taking a route path from another planning group will modify the other planning group
	SourcePlanningGroup *Sourceplanninggrouprequest `json:"sourcePlanningGroup,omitempty"`

}

func (o *Routepathrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routepathrequest
	
	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		SourcePlanningGroup *Sourceplanninggrouprequest `json:"sourcePlanningGroup,omitempty"`
		*Alias
	}{ 
		QueueId: o.QueueId,
		
		MediaType: o.MediaType,
		
		LanguageId: o.LanguageId,
		
		SkillIds: o.SkillIds,
		
		SourcePlanningGroup: o.SourcePlanningGroup,
		Alias:    (*Alias)(o),
	})
}

func (o *Routepathrequest) UnmarshalJSON(b []byte) error {
	var RoutepathrequestMap map[string]interface{}
	err := json.Unmarshal(b, &RoutepathrequestMap)
	if err != nil {
		return err
	}
	
	if QueueId, ok := RoutepathrequestMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if MediaType, ok := RoutepathrequestMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if LanguageId, ok := RoutepathrequestMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if SkillIds, ok := RoutepathrequestMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if SourcePlanningGroup, ok := RoutepathrequestMap["sourcePlanningGroup"].(map[string]interface{}); ok {
		SourcePlanningGroupString, _ := json.Marshal(SourcePlanningGroup)
		json.Unmarshal(SourcePlanningGroupString, &o.SourcePlanningGroup)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Routepathrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
