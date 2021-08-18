package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Minerintent
type Minerintent struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Intent name.
	Name *string `json:"name,omitempty"`


	// Miner - The miner to which the intent belongs.
	Miner *Miner `json:"miner,omitempty"`


	// Utterances - The utterances that are extracted for an Intent.
	Utterances *[]Utterance `json:"utterances,omitempty"`


	// AnalyticVolumePercent - Percentage of conversations belonging to the intent.
	AnalyticVolumePercent *float64 `json:"analyticVolumePercent,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Minerintent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Minerintent

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Miner *Miner `json:"miner,omitempty"`
		
		Utterances *[]Utterance `json:"utterances,omitempty"`
		
		AnalyticVolumePercent *float64 `json:"analyticVolumePercent,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Miner: u.Miner,
		
		Utterances: u.Utterances,
		
		AnalyticVolumePercent: u.AnalyticVolumePercent,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Minerintent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
