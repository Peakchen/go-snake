// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: google/ads/googleads/v2/resources/customer_negative_criterion.proto

package resources

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A negative criterion for exclusions at the customer level.
type CustomerNegativeCriterion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the customer negative criterion.
	// Customer negative criterion resource names have the form:
	//
	// `customers/{customer_id}/customerNegativeCriteria/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the criterion.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The type of the criterion.
	Type enums.CriterionTypeEnum_CriterionType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.CriterionTypeEnum_CriterionType" json:"type,omitempty"`
	// The customer negative criterion.
	//
	// Exactly one must be set.
	//
	// Types that are assignable to Criterion:
	//	*CustomerNegativeCriterion_ContentLabel
	//	*CustomerNegativeCriterion_MobileApplication
	//	*CustomerNegativeCriterion_MobileAppCategory
	//	*CustomerNegativeCriterion_Placement
	//	*CustomerNegativeCriterion_YoutubeVideo
	//	*CustomerNegativeCriterion_YoutubeChannel
	Criterion isCustomerNegativeCriterion_Criterion `protobuf_oneof:"criterion"`
}

func (x *CustomerNegativeCriterion) Reset() {
	*x = CustomerNegativeCriterion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerNegativeCriterion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerNegativeCriterion) ProtoMessage() {}

func (x *CustomerNegativeCriterion) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerNegativeCriterion.ProtoReflect.Descriptor instead.
func (*CustomerNegativeCriterion) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerNegativeCriterion) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CustomerNegativeCriterion) GetId() *wrappers.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *CustomerNegativeCriterion) GetType() enums.CriterionTypeEnum_CriterionType {
	if x != nil {
		return x.Type
	}
	return enums.CriterionTypeEnum_UNSPECIFIED
}

func (m *CustomerNegativeCriterion) GetCriterion() isCustomerNegativeCriterion_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (x *CustomerNegativeCriterion) GetContentLabel() *common.ContentLabelInfo {
	if x, ok := x.GetCriterion().(*CustomerNegativeCriterion_ContentLabel); ok {
		return x.ContentLabel
	}
	return nil
}

func (x *CustomerNegativeCriterion) GetMobileApplication() *common.MobileApplicationInfo {
	if x, ok := x.GetCriterion().(*CustomerNegativeCriterion_MobileApplication); ok {
		return x.MobileApplication
	}
	return nil
}

func (x *CustomerNegativeCriterion) GetMobileAppCategory() *common.MobileAppCategoryInfo {
	if x, ok := x.GetCriterion().(*CustomerNegativeCriterion_MobileAppCategory); ok {
		return x.MobileAppCategory
	}
	return nil
}

func (x *CustomerNegativeCriterion) GetPlacement() *common.PlacementInfo {
	if x, ok := x.GetCriterion().(*CustomerNegativeCriterion_Placement); ok {
		return x.Placement
	}
	return nil
}

func (x *CustomerNegativeCriterion) GetYoutubeVideo() *common.YouTubeVideoInfo {
	if x, ok := x.GetCriterion().(*CustomerNegativeCriterion_YoutubeVideo); ok {
		return x.YoutubeVideo
	}
	return nil
}

func (x *CustomerNegativeCriterion) GetYoutubeChannel() *common.YouTubeChannelInfo {
	if x, ok := x.GetCriterion().(*CustomerNegativeCriterion_YoutubeChannel); ok {
		return x.YoutubeChannel
	}
	return nil
}

type isCustomerNegativeCriterion_Criterion interface {
	isCustomerNegativeCriterion_Criterion()
}

type CustomerNegativeCriterion_ContentLabel struct {
	// Immutable. ContentLabel.
	ContentLabel *common.ContentLabelInfo `protobuf:"bytes,4,opt,name=content_label,json=contentLabel,proto3,oneof"`
}

type CustomerNegativeCriterion_MobileApplication struct {
	// Immutable. MobileApplication.
	MobileApplication *common.MobileApplicationInfo `protobuf:"bytes,5,opt,name=mobile_application,json=mobileApplication,proto3,oneof"`
}

type CustomerNegativeCriterion_MobileAppCategory struct {
	// Immutable. MobileAppCategory.
	MobileAppCategory *common.MobileAppCategoryInfo `protobuf:"bytes,6,opt,name=mobile_app_category,json=mobileAppCategory,proto3,oneof"`
}

type CustomerNegativeCriterion_Placement struct {
	// Immutable. Placement.
	Placement *common.PlacementInfo `protobuf:"bytes,7,opt,name=placement,proto3,oneof"`
}

type CustomerNegativeCriterion_YoutubeVideo struct {
	// Immutable. YouTube Video.
	YoutubeVideo *common.YouTubeVideoInfo `protobuf:"bytes,8,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

type CustomerNegativeCriterion_YoutubeChannel struct {
	// Immutable. YouTube Channel.
	YoutubeChannel *common.YouTubeChannelInfo `protobuf:"bytes,9,opt,name=youtube_channel,json=youtubeChannel,proto3,oneof"`
}

func (*CustomerNegativeCriterion_ContentLabel) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_MobileApplication) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_MobileAppCategory) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_Placement) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_YoutubeVideo) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_YoutubeChannel) isCustomerNegativeCriterion_Criterion() {}

var File_google_ads_googleads_v2_resources_customer_negative_criterion_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x65, 0x67,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x07, 0x0a, 0x19, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x6f, 0x6e, 0x12, 0x5f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a, 0xe0, 0x41, 0x05, 0xfa,
	0x41, 0x34, 0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x57, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x5c, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x6b, 0x0a,
	0x12, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x11, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6c, 0x0a, 0x13, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41,
	0x70, 0x70, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03,
	0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x11, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x52, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48,
	0x00, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5c, 0x0a, 0x0d,
	0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x59, 0x6f, 0x75, 0x54, 0x75, 0x62, 0x65, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x0c, 0x79, 0x6f,
	0x75, 0x74, 0x75, 0x62, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x62, 0x0a, 0x0f, 0x79, 0x6f,
	0x75, 0x74, 0x75, 0x62, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x59, 0x6f, 0x75, 0x54, 0x75, 0x62, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x0e,
	0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x3a, 0x85,
	0x01, 0xea, 0x41, 0x81, 0x01, 0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x7d, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x6f, 0x6e, 0x7d, 0x42, 0x0b, 0x0a, 0x09, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x6f, 0x6e, 0x42, 0x8b, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x1e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x32, 0x5c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_rawDescData = file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_rawDesc
)

func file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_rawDescData
}

var file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_goTypes = []interface{}{
	(*CustomerNegativeCriterion)(nil),          // 0: google.ads.googleads.v2.resources.CustomerNegativeCriterion
	(*wrappers.Int64Value)(nil),                // 1: google.protobuf.Int64Value
	(enums.CriterionTypeEnum_CriterionType)(0), // 2: google.ads.googleads.v2.enums.CriterionTypeEnum.CriterionType
	(*common.ContentLabelInfo)(nil),            // 3: google.ads.googleads.v2.common.ContentLabelInfo
	(*common.MobileApplicationInfo)(nil),       // 4: google.ads.googleads.v2.common.MobileApplicationInfo
	(*common.MobileAppCategoryInfo)(nil),       // 5: google.ads.googleads.v2.common.MobileAppCategoryInfo
	(*common.PlacementInfo)(nil),               // 6: google.ads.googleads.v2.common.PlacementInfo
	(*common.YouTubeVideoInfo)(nil),            // 7: google.ads.googleads.v2.common.YouTubeVideoInfo
	(*common.YouTubeChannelInfo)(nil),          // 8: google.ads.googleads.v2.common.YouTubeChannelInfo
}
var file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v2.resources.CustomerNegativeCriterion.id:type_name -> google.protobuf.Int64Value
	2, // 1: google.ads.googleads.v2.resources.CustomerNegativeCriterion.type:type_name -> google.ads.googleads.v2.enums.CriterionTypeEnum.CriterionType
	3, // 2: google.ads.googleads.v2.resources.CustomerNegativeCriterion.content_label:type_name -> google.ads.googleads.v2.common.ContentLabelInfo
	4, // 3: google.ads.googleads.v2.resources.CustomerNegativeCriterion.mobile_application:type_name -> google.ads.googleads.v2.common.MobileApplicationInfo
	5, // 4: google.ads.googleads.v2.resources.CustomerNegativeCriterion.mobile_app_category:type_name -> google.ads.googleads.v2.common.MobileAppCategoryInfo
	6, // 5: google.ads.googleads.v2.resources.CustomerNegativeCriterion.placement:type_name -> google.ads.googleads.v2.common.PlacementInfo
	7, // 6: google.ads.googleads.v2.resources.CustomerNegativeCriterion.youtube_video:type_name -> google.ads.googleads.v2.common.YouTubeVideoInfo
	8, // 7: google.ads.googleads.v2.resources.CustomerNegativeCriterion.youtube_channel:type_name -> google.ads.googleads.v2.common.YouTubeChannelInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_init() }
func file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_init() {
	if File_google_ads_googleads_v2_resources_customer_negative_criterion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerNegativeCriterion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CustomerNegativeCriterion_ContentLabel)(nil),
		(*CustomerNegativeCriterion_MobileApplication)(nil),
		(*CustomerNegativeCriterion_MobileAppCategory)(nil),
		(*CustomerNegativeCriterion_Placement)(nil),
		(*CustomerNegativeCriterion_YoutubeVideo)(nil),
		(*CustomerNegativeCriterion_YoutubeChannel)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_resources_customer_negative_criterion_proto = out.File
	file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_rawDesc = nil
	file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_goTypes = nil
	file_google_ads_googleads_v2_resources_customer_negative_criterion_proto_depIdxs = nil
}
