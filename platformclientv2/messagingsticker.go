package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingsticker
type Messagingsticker struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ProviderStickerId - The sticker Id of the sticker, assigned by the sticker provider.
	ProviderStickerId *int `json:"providerStickerId,omitempty"`


	// ProviderPackageId - The package Id of the sticker, assigned by the sticker provider.
	ProviderPackageId *int `json:"providerPackageId,omitempty"`


	// PackageName - The package name of the sticker, assigned by the sticker provider.
	PackageName *string `json:"packageName,omitempty"`


	// MessengerType - The type of the messenger provider.
	MessengerType *string `json:"messengerType,omitempty"`


	// StickerType - The type of the sticker.
	StickerType *string `json:"stickerType,omitempty"`


	// ProviderVersion - The version of the sticker, assigned by the provider.
	ProviderVersion *int `json:"providerVersion,omitempty"`


	// UriLocation
	UriLocation *string `json:"uriLocation,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Messagingsticker) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagingsticker

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ProviderStickerId *int `json:"providerStickerId,omitempty"`
		
		ProviderPackageId *int `json:"providerPackageId,omitempty"`
		
		PackageName *string `json:"packageName,omitempty"`
		
		MessengerType *string `json:"messengerType,omitempty"`
		
		StickerType *string `json:"stickerType,omitempty"`
		
		ProviderVersion *int `json:"providerVersion,omitempty"`
		
		UriLocation *string `json:"uriLocation,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ProviderStickerId: u.ProviderStickerId,
		
		ProviderPackageId: u.ProviderPackageId,
		
		PackageName: u.PackageName,
		
		MessengerType: u.MessengerType,
		
		StickerType: u.StickerType,
		
		ProviderVersion: u.ProviderVersion,
		
		UriLocation: u.UriLocation,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Messagingsticker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
