/*
Portions Copyright (c) Microsoft Corporation.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package imagefamily

import (
	"context"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	armcomputev5 "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"sigs.k8s.io/karpenter/pkg/scheduling"
)

const (
	AKSUbuntuPublicGalleryURL     = "AKSUbuntu-38d80f77-467a-481f-a8d4-09b6d4220bd2"
	AKSAzureLinuxPublicGalleryURL = "AKSAzureLinux-f7c7cda5-1c9a-4bdc-a222-9614c968580b"

	AKSUbuntuResourceGroup     = "AKS-Ubuntu"
	AKSAzureLinuxResourceGroup = "AKS-AzureLinux"

	AKSUbuntuGalleryName     = "AKSUbuntu"
	AKSAzureLinuxGalleryName = "AKSAzureLinux"
)

// DefaultImageOutput is a stub describing our desired image with an image's name and requirements to run that image
type DefaultImageOutput struct {
	// Community Image Gallery
	PublicGalleryURL string
	// Shared Image Gallery
	GalleryResourceGroup string
	GalleryName          string
	// Common
	ImageDefinition string
	Distro          string
	Requirements    scheduling.Requirements
}

func (d *DefaultImageOutput) PopulateImageTraitsFromID(imageID string) {
	// We want to take a community image gallery image id or a shared image gallery id and populate the contents of DefaultImageOutput
	imageIDParts := strings.Split(imageID, "/")
	if imageIDParts[1] == "subscriptions" { // Shared Image Gallery
		d.GalleryResourceGroup = imageIDParts[4]
		d.GalleryName = imageIDParts[8]
		d.ImageDefinition = imageIDParts[10]
	}
	if imageIDParts[1] == "CommunityGalleries" { // Community Image Gallery
		d.PublicGalleryURL = imageIDParts[2]
		d.ImageDefinition = imageIDParts[4]
	}
}

// CommunityGalleryImageVersionsAPI is used for listing community gallery image versions.
type CommunityGalleryImageVersionsAPI interface {
	NewListPager(location string, publicGalleryName string, galleryImageName string, options *armcomputev5.CommunityGalleryImageVersionsClientListOptions) *runtime.Pager[armcomputev5.CommunityGalleryImageVersionsClientListResponse]
}

type NodeImageVersion struct {
	FullName string `json:"fullName"`
	OS       string `json:"os"`
	SKU      string `json:"sku"`
	Version  string `json:"version"`
}

type NodeImageVersionsResponse struct {
	Values []NodeImageVersion `json:"values"`
}

type NodeImageVersionsAPI interface {
	List(ctx context.Context, location, subscription string) (NodeImageVersionsResponse, error)
}
