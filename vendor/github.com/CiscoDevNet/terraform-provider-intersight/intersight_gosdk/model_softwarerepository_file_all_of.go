/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-28T13:03:38Z.
 *
 * API version: 1.0.9-4267
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"time"
)

// SoftwarerepositoryFileAllOf Definition of the list of properties defined in 'softwarerepository.File', excluding properties defined in parent classes.
type SoftwarerepositoryFileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// User provided description about the file. Cisco provided description for image inventoried from a Cisco repository.
	Description *string `json:"Description,omitempty"`
	// The number of times this file has been downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache.
	DownloadCount *int64 `json:"DownloadCount,omitempty"`
	// The action to be performed on the imported file. If 'PreCache' is set, the image will be cached in Appliance. Applicable in Intersight appliance deployment. If 'Evict' is set, the cached file will be removed. Applicable in Intersight appliance deployment. If 'GeneratePreSignedUploadUrl' is set, generates pre signed URL (s) for the file to be imported into the repository. Applicable for local machine source. The URL (s) will be populated under LocalMachine file server. If 'CompleteImportProcess' is set, the ImportState is marked as 'Imported'. Applicable for local machine source. If 'Cancel' is set, the ImportState is marked as 'Failed'. Applicable for local machine source. * `None` - No action should be taken on the imported file. * `GeneratePreSignedUploadUrl` - Generate pre signed URL of file for importing into the repository. * `GeneratePreSignedDownloadUrl` - Generate pre signed URL of file in the repository to download. * `CompleteImportProcess` - Mark that the import process of the file into the repository is complete. * `MarkImportFailed` - Mark to indicate that the import process of the file into the repository failed. * `PreCache` - Cache the file into the Intersight Appliance. * `Cancel` - The cancel import process for the file into the repository. * `Extract` - The action to extract the file in the external repository. * `Evict` - Evict the cached file from the Intersight Appliance.
	ImportAction *string `json:"ImportAction,omitempty"`
	// The state  of this file in the repository or Appliance. The importState is updated during the import operation and as part of the repository monitoring process. * `ReadyForImport` - The image is ready to be imported into the repository. * `Importing` - The image is being imported into the repository. * `Imported` - The image has been extracted and imported into the repository. * `PendingExtraction` - Indicates that the image has been imported but not extracted in the repository. * `Extracting` - Indicates that the image is being extracted into the repository. * `Extracted` - Indicates that the image has been extracted into the repository. * `Failed` - The image import from an external source to the repository has failed. * `MetaOnly` - The image is present in an external repository. * `ReadyForCache` - The image is ready to be cached into the Intersight Appliance. * `Caching` - Indicates that the image is being cached into the Intersight Appliance or endpoint cache. * `Cached` - Indicates that the image has been cached into the Intersight Appliance or endpoint cache. * `CachingFailed` - Indicates that the image caching into the Intersight Appliance failed or endpoint cache. * `Corrupted` - Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached. * `Evicted` - Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space. * `Invalid` - Indicates that the corresponding distributable MO has been removed from the backend. This can be due to unpublishing of an image.
	ImportState *string `json:"ImportState,omitempty"`
	// The time at which this image or file was imported/cached into the repositry. if the 'ImportState' is 'Imported', the time at which this image or file was imported. if the 'ImportState' is 'Cached', the time at which this image or file was cached.
	ImportedTime *time.Time `json:"ImportedTime,omitempty"`
	// The time at which this file was last downloaded from the local repository. It is used by the repository monitoring process to determine the files that are to be evicted from the cache.
	LastAccessTime *time.Time `json:"LastAccessTime,omitempty"`
	// The MD5 ETag for a file that is stored in Intersight repository or in the appliance cache. Warning - MD5 is currently broken and this will be migrated to SHA shortly.
	Md5eTag *string `json:"Md5eTag,omitempty"`
	// The md5sum checksum of the file. This information is available for all Cisco distributed images and files imported to the local repository.
	Md5sum *string `json:"Md5sum,omitempty"`
	// The name of the file. It is populated as part of the image import operation.
	Name *string `json:"Name,omitempty"`
	// The date on which the file was released or distributed by its vendor.
	ReleaseDate *time.Time `json:"ReleaseDate,omitempty"`
	// The sha512sum of the file. This information is available for all Cisco distributed images and files imported to the local repository.
	Sha512sum *string `json:"Sha512sum,omitempty"`
	// The size (in bytes) of the file. This information is available for all Cisco distributed images and files imported to the local repository.
	Size *int64 `json:"Size,omitempty"`
	// The software advisory, if any, provided by the vendor for this file.
	SoftwareAdvisoryUrl *string                              `json:"SoftwareAdvisoryUrl,omitempty"`
	Source              NullableSoftwarerepositoryFileServer `json:"Source,omitempty"`
	// Vendor provided version for the file.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryFileAllOf SoftwarerepositoryFileAllOf

// NewSoftwarerepositoryFileAllOf instantiates a new SoftwarerepositoryFileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryFileAllOf(classId string, objectType string) *SoftwarerepositoryFileAllOf {
	this := SoftwarerepositoryFileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var importAction string = "None"
	this.ImportAction = &importAction
	var importState string = "ReadyForImport"
	this.ImportState = &importState
	return &this
}

// NewSoftwarerepositoryFileAllOfWithDefaults instantiates a new SoftwarerepositoryFileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryFileAllOfWithDefaults() *SoftwarerepositoryFileAllOf {
	this := SoftwarerepositoryFileAllOf{}
	var importAction string = "None"
	this.ImportAction = &importAction
	var importState string = "ReadyForImport"
	this.ImportState = &importState
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryFileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryFileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryFileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryFileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SoftwarerepositoryFileAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetDownloadCount returns the DownloadCount field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetDownloadCount() int64 {
	if o == nil || o.DownloadCount == nil {
		var ret int64
		return ret
	}
	return *o.DownloadCount
}

// GetDownloadCountOk returns a tuple with the DownloadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetDownloadCountOk() (*int64, bool) {
	if o == nil || o.DownloadCount == nil {
		return nil, false
	}
	return o.DownloadCount, true
}

// HasDownloadCount returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasDownloadCount() bool {
	if o != nil && o.DownloadCount != nil {
		return true
	}

	return false
}

// SetDownloadCount gets a reference to the given int64 and assigns it to the DownloadCount field.
func (o *SoftwarerepositoryFileAllOf) SetDownloadCount(v int64) {
	o.DownloadCount = &v
}

// GetImportAction returns the ImportAction field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetImportAction() string {
	if o == nil || o.ImportAction == nil {
		var ret string
		return ret
	}
	return *o.ImportAction
}

// GetImportActionOk returns a tuple with the ImportAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetImportActionOk() (*string, bool) {
	if o == nil || o.ImportAction == nil {
		return nil, false
	}
	return o.ImportAction, true
}

// HasImportAction returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasImportAction() bool {
	if o != nil && o.ImportAction != nil {
		return true
	}

	return false
}

// SetImportAction gets a reference to the given string and assigns it to the ImportAction field.
func (o *SoftwarerepositoryFileAllOf) SetImportAction(v string) {
	o.ImportAction = &v
}

// GetImportState returns the ImportState field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetImportState() string {
	if o == nil || o.ImportState == nil {
		var ret string
		return ret
	}
	return *o.ImportState
}

// GetImportStateOk returns a tuple with the ImportState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetImportStateOk() (*string, bool) {
	if o == nil || o.ImportState == nil {
		return nil, false
	}
	return o.ImportState, true
}

// HasImportState returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasImportState() bool {
	if o != nil && o.ImportState != nil {
		return true
	}

	return false
}

// SetImportState gets a reference to the given string and assigns it to the ImportState field.
func (o *SoftwarerepositoryFileAllOf) SetImportState(v string) {
	o.ImportState = &v
}

// GetImportedTime returns the ImportedTime field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetImportedTime() time.Time {
	if o == nil || o.ImportedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ImportedTime
}

// GetImportedTimeOk returns a tuple with the ImportedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetImportedTimeOk() (*time.Time, bool) {
	if o == nil || o.ImportedTime == nil {
		return nil, false
	}
	return o.ImportedTime, true
}

// HasImportedTime returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasImportedTime() bool {
	if o != nil && o.ImportedTime != nil {
		return true
	}

	return false
}

// SetImportedTime gets a reference to the given time.Time and assigns it to the ImportedTime field.
func (o *SoftwarerepositoryFileAllOf) SetImportedTime(v time.Time) {
	o.ImportedTime = &v
}

// GetLastAccessTime returns the LastAccessTime field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetLastAccessTime() time.Time {
	if o == nil || o.LastAccessTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessTime
}

// GetLastAccessTimeOk returns a tuple with the LastAccessTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetLastAccessTimeOk() (*time.Time, bool) {
	if o == nil || o.LastAccessTime == nil {
		return nil, false
	}
	return o.LastAccessTime, true
}

// HasLastAccessTime returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasLastAccessTime() bool {
	if o != nil && o.LastAccessTime != nil {
		return true
	}

	return false
}

// SetLastAccessTime gets a reference to the given time.Time and assigns it to the LastAccessTime field.
func (o *SoftwarerepositoryFileAllOf) SetLastAccessTime(v time.Time) {
	o.LastAccessTime = &v
}

// GetMd5eTag returns the Md5eTag field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetMd5eTag() string {
	if o == nil || o.Md5eTag == nil {
		var ret string
		return ret
	}
	return *o.Md5eTag
}

// GetMd5eTagOk returns a tuple with the Md5eTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetMd5eTagOk() (*string, bool) {
	if o == nil || o.Md5eTag == nil {
		return nil, false
	}
	return o.Md5eTag, true
}

// HasMd5eTag returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasMd5eTag() bool {
	if o != nil && o.Md5eTag != nil {
		return true
	}

	return false
}

// SetMd5eTag gets a reference to the given string and assigns it to the Md5eTag field.
func (o *SoftwarerepositoryFileAllOf) SetMd5eTag(v string) {
	o.Md5eTag = &v
}

// GetMd5sum returns the Md5sum field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetMd5sum() string {
	if o == nil || o.Md5sum == nil {
		var ret string
		return ret
	}
	return *o.Md5sum
}

// GetMd5sumOk returns a tuple with the Md5sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetMd5sumOk() (*string, bool) {
	if o == nil || o.Md5sum == nil {
		return nil, false
	}
	return o.Md5sum, true
}

// HasMd5sum returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasMd5sum() bool {
	if o != nil && o.Md5sum != nil {
		return true
	}

	return false
}

// SetMd5sum gets a reference to the given string and assigns it to the Md5sum field.
func (o *SoftwarerepositoryFileAllOf) SetMd5sum(v string) {
	o.Md5sum = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SoftwarerepositoryFileAllOf) SetName(v string) {
	o.Name = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetReleaseDate() time.Time {
	if o == nil || o.ReleaseDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || o.ReleaseDate == nil {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate != nil {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *SoftwarerepositoryFileAllOf) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

// GetSha512sum returns the Sha512sum field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetSha512sum() string {
	if o == nil || o.Sha512sum == nil {
		var ret string
		return ret
	}
	return *o.Sha512sum
}

// GetSha512sumOk returns a tuple with the Sha512sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetSha512sumOk() (*string, bool) {
	if o == nil || o.Sha512sum == nil {
		return nil, false
	}
	return o.Sha512sum, true
}

// HasSha512sum returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasSha512sum() bool {
	if o != nil && o.Sha512sum != nil {
		return true
	}

	return false
}

// SetSha512sum gets a reference to the given string and assigns it to the Sha512sum field.
func (o *SoftwarerepositoryFileAllOf) SetSha512sum(v string) {
	o.Sha512sum = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *SoftwarerepositoryFileAllOf) SetSize(v int64) {
	o.Size = &v
}

// GetSoftwareAdvisoryUrl returns the SoftwareAdvisoryUrl field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetSoftwareAdvisoryUrl() string {
	if o == nil || o.SoftwareAdvisoryUrl == nil {
		var ret string
		return ret
	}
	return *o.SoftwareAdvisoryUrl
}

// GetSoftwareAdvisoryUrlOk returns a tuple with the SoftwareAdvisoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetSoftwareAdvisoryUrlOk() (*string, bool) {
	if o == nil || o.SoftwareAdvisoryUrl == nil {
		return nil, false
	}
	return o.SoftwareAdvisoryUrl, true
}

// HasSoftwareAdvisoryUrl returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasSoftwareAdvisoryUrl() bool {
	if o != nil && o.SoftwareAdvisoryUrl != nil {
		return true
	}

	return false
}

// SetSoftwareAdvisoryUrl gets a reference to the given string and assigns it to the SoftwareAdvisoryUrl field.
func (o *SoftwarerepositoryFileAllOf) SetSoftwareAdvisoryUrl(v string) {
	o.SoftwareAdvisoryUrl = &v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryFileAllOf) GetSource() SoftwarerepositoryFileServer {
	if o == nil || o.Source.Get() == nil {
		var ret SoftwarerepositoryFileServer
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryFileAllOf) GetSourceOk() (*SoftwarerepositoryFileServer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableSoftwarerepositoryFileServer and assigns it to the Source field.
func (o *SoftwarerepositoryFileAllOf) SetSource(v SoftwarerepositoryFileServer) {
	o.Source.Set(&v)
}

// SetSourceNil sets the value for Source to be an explicit nil
func (o *SoftwarerepositoryFileAllOf) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *SoftwarerepositoryFileAllOf) UnsetSource() {
	o.Source.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SoftwarerepositoryFileAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryFileAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SoftwarerepositoryFileAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SoftwarerepositoryFileAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o SoftwarerepositoryFileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DownloadCount != nil {
		toSerialize["DownloadCount"] = o.DownloadCount
	}
	if o.ImportAction != nil {
		toSerialize["ImportAction"] = o.ImportAction
	}
	if o.ImportState != nil {
		toSerialize["ImportState"] = o.ImportState
	}
	if o.ImportedTime != nil {
		toSerialize["ImportedTime"] = o.ImportedTime
	}
	if o.LastAccessTime != nil {
		toSerialize["LastAccessTime"] = o.LastAccessTime
	}
	if o.Md5eTag != nil {
		toSerialize["Md5eTag"] = o.Md5eTag
	}
	if o.Md5sum != nil {
		toSerialize["Md5sum"] = o.Md5sum
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ReleaseDate != nil {
		toSerialize["ReleaseDate"] = o.ReleaseDate
	}
	if o.Sha512sum != nil {
		toSerialize["Sha512sum"] = o.Sha512sum
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.SoftwareAdvisoryUrl != nil {
		toSerialize["SoftwareAdvisoryUrl"] = o.SoftwareAdvisoryUrl
	}
	if o.Source.IsSet() {
		toSerialize["Source"] = o.Source.Get()
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryFileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwarerepositoryFileAllOf := _SoftwarerepositoryFileAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwarerepositoryFileAllOf); err == nil {
		*o = SoftwarerepositoryFileAllOf(varSoftwarerepositoryFileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "DownloadCount")
		delete(additionalProperties, "ImportAction")
		delete(additionalProperties, "ImportState")
		delete(additionalProperties, "ImportedTime")
		delete(additionalProperties, "LastAccessTime")
		delete(additionalProperties, "Md5eTag")
		delete(additionalProperties, "Md5sum")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ReleaseDate")
		delete(additionalProperties, "Sha512sum")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "SoftwareAdvisoryUrl")
		delete(additionalProperties, "Source")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryFileAllOf struct {
	value *SoftwarerepositoryFileAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryFileAllOf) Get() *SoftwarerepositoryFileAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryFileAllOf) Set(val *SoftwarerepositoryFileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryFileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryFileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryFileAllOf(val *SoftwarerepositoryFileAllOf) *NullableSoftwarerepositoryFileAllOf {
	return &NullableSoftwarerepositoryFileAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryFileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryFileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
