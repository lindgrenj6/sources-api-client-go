/*
 * Sources
 *
 * Sources
 *
 * API version: 3.1.0
 * Contact: support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sourcesapi

import (
	"encoding/json"
)

// AppMetaDataCollection struct for AppMetaDataCollection
type AppMetaDataCollection struct {
	Meta *CollectionMetadata `json:"meta,omitempty"`
	Links *CollectionLinks `json:"links,omitempty"`
	Data *[]AppMetaData `json:"data,omitempty"`
}

// NewAppMetaDataCollection instantiates a new AppMetaDataCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppMetaDataCollection() *AppMetaDataCollection {
	this := AppMetaDataCollection{}
	return &this
}

// NewAppMetaDataCollectionWithDefaults instantiates a new AppMetaDataCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppMetaDataCollectionWithDefaults() *AppMetaDataCollection {
	this := AppMetaDataCollection{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppMetaDataCollection) GetMeta() CollectionMetadata {
	if o == nil || o.Meta == nil {
		var ret CollectionMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMetaDataCollection) GetMetaOk() (*CollectionMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppMetaDataCollection) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CollectionMetadata and assigns it to the Meta field.
func (o *AppMetaDataCollection) SetMeta(v CollectionMetadata) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppMetaDataCollection) GetLinks() CollectionLinks {
	if o == nil || o.Links == nil {
		var ret CollectionLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMetaDataCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppMetaDataCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CollectionLinks and assigns it to the Links field.
func (o *AppMetaDataCollection) SetLinks(v CollectionLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppMetaDataCollection) GetData() []AppMetaData {
	if o == nil || o.Data == nil {
		var ret []AppMetaData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppMetaDataCollection) GetDataOk() (*[]AppMetaData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppMetaDataCollection) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppMetaData and assigns it to the Data field.
func (o *AppMetaDataCollection) SetData(v []AppMetaData) {
	o.Data = &v
}

func (o AppMetaDataCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAppMetaDataCollection struct {
	value *AppMetaDataCollection
	isSet bool
}

func (v NullableAppMetaDataCollection) Get() *AppMetaDataCollection {
	return v.value
}

func (v *NullableAppMetaDataCollection) Set(val *AppMetaDataCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableAppMetaDataCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableAppMetaDataCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppMetaDataCollection(val *AppMetaDataCollection) *NullableAppMetaDataCollection {
	return &NullableAppMetaDataCollection{value: val, isSet: true}
}

func (v NullableAppMetaDataCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppMetaDataCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


