// Code generated by protoc-gen-go.
// source: node.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateNodeRequest struct {
	// Name of the application to which the node must be added.
	ApplicationName string `protobuf:"bytes,13,opt,name=applicationName" json:"applicationName,omitempty"`
	// Hex encoded DevEUI.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// Hex encoded AppEUI.
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// Hex encoded AppKey. When isABP is set to true, this field is not needed.
	AppKey string `protobuf:"bytes,3,opt,name=appKey" json:"appKey,omitempty"`
	// RX delay.
	RxDelay uint32 `protobuf:"varint,4,opt,name=rxDelay" json:"rxDelay,omitempty"`
	// RX1 data-rate offset.
	Rx1DROffset uint32 `protobuf:"varint,5,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	// Channel-list ID used for CFlist (see LoRaWAN regional parameters if this applies to your region).
	ChannelListID int64 `protobuf:"varint,6,opt,name=channelListID" json:"channelListID,omitempty"`
	// RX window to use.
	RxWindow RXWindow `protobuf:"varint,7,opt,name=rxWindow,enum=api.RXWindow" json:"rxWindow,omitempty"`
	// Data-rate to use for RX2.
	Rx2DR uint32 `protobuf:"varint,8,opt,name=rx2DR" json:"rx2DR,omitempty"`
	// Name of the node (if left blank, it will be set to the DevEUI).
	Name string `protobuf:"bytes,9,opt,name=name" json:"name,omitempty"`
	// Relax frame-counter mode is enabled.
	RelaxFCnt bool `protobuf:"varint,10,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
	// Interval (in frames) in which the ADR engine may adapt the data-rate of the node (0 = disabled).
	AdrInterval uint32 `protobuf:"varint,11,opt,name=adrInterval" json:"adrInterval,omitempty"`
	// Installation-margin to use for ADR calculation.
	InstallationMargin float64 `protobuf:"fixed64,12,opt,name=installationMargin" json:"installationMargin,omitempty"`
	// Description of the node.
	Description string `protobuf:"bytes,14,opt,name=description" json:"description,omitempty"`
	// Node is activated by ABP.
	IsABP bool `protobuf:"varint,15,opt,name=isABP" json:"isABP,omitempty"`
}

func (m *CreateNodeRequest) Reset()                    { *m = CreateNodeRequest{} }
func (m *CreateNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeRequest) ProtoMessage()               {}
func (*CreateNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateNodeRequest) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *CreateNodeRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *CreateNodeRequest) GetAppEUI() string {
	if m != nil {
		return m.AppEUI
	}
	return ""
}

func (m *CreateNodeRequest) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *CreateNodeRequest) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *CreateNodeRequest) GetRx1DROffset() uint32 {
	if m != nil {
		return m.Rx1DROffset
	}
	return 0
}

func (m *CreateNodeRequest) GetChannelListID() int64 {
	if m != nil {
		return m.ChannelListID
	}
	return 0
}

func (m *CreateNodeRequest) GetRxWindow() RXWindow {
	if m != nil {
		return m.RxWindow
	}
	return RXWindow_RX1
}

func (m *CreateNodeRequest) GetRx2DR() uint32 {
	if m != nil {
		return m.Rx2DR
	}
	return 0
}

func (m *CreateNodeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateNodeRequest) GetRelaxFCnt() bool {
	if m != nil {
		return m.RelaxFCnt
	}
	return false
}

func (m *CreateNodeRequest) GetAdrInterval() uint32 {
	if m != nil {
		return m.AdrInterval
	}
	return 0
}

func (m *CreateNodeRequest) GetInstallationMargin() float64 {
	if m != nil {
		return m.InstallationMargin
	}
	return 0
}

func (m *CreateNodeRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateNodeRequest) GetIsABP() bool {
	if m != nil {
		return m.IsABP
	}
	return false
}

type CreateNodeResponse struct {
}

func (m *CreateNodeResponse) Reset()                    { *m = CreateNodeResponse{} }
func (m *CreateNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeResponse) ProtoMessage()               {}
func (*CreateNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type GetNodeRequest struct {
	// Name of the application owning the node.
	ApplicationName string `protobuf:"bytes,2,opt,name=applicationName" json:"applicationName,omitempty"`
	// Name of the node.
	NodeName string `protobuf:"bytes,1,opt,name=nodeName" json:"nodeName,omitempty"`
}

func (m *GetNodeRequest) Reset()                    { *m = GetNodeRequest{} }
func (m *GetNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodeRequest) ProtoMessage()               {}
func (*GetNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetNodeRequest) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *GetNodeRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type GetNodeResponse struct {
	// Hex encoded DevEUI.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// Hex encoded AppEUI.
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// Hex encoded AppKey.
	AppKey string `protobuf:"bytes,3,opt,name=appKey" json:"appKey,omitempty"`
	// RX delay.
	RxDelay uint32 `protobuf:"varint,4,opt,name=rxDelay" json:"rxDelay,omitempty"`
	// RX1 data-rate offset.
	Rx1DROffset uint32 `protobuf:"varint,5,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	// Channel-list ID used for CFlist (see LoRaWAN regional parameters if this applies to your region).
	ChannelListID int64 `protobuf:"varint,6,opt,name=channelListID" json:"channelListID,omitempty"`
	// RX window to use.
	RxWindow RXWindow `protobuf:"varint,7,opt,name=rxWindow,enum=api.RXWindow" json:"rxWindow,omitempty"`
	// Data-rate to use for RX2.
	Rx2DR uint32 `protobuf:"varint,8,opt,name=rx2DR" json:"rx2DR,omitempty"`
	// Name of the node.
	Name string `protobuf:"bytes,9,opt,name=name" json:"name,omitempty"`
	// Relax frame-counter mode is enabled.
	RelaxFCnt bool `protobuf:"varint,10,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
	// Interval (in frames) in which the ADR engine may adapt the data-rate of the node (0 = disabled).
	AdrInterval uint32 `protobuf:"varint,11,opt,name=adrInterval" json:"adrInterval,omitempty"`
	// Installation-margin to use for ADR calculation.
	InstallationMargin float64 `protobuf:"fixed64,12,opt,name=installationMargin" json:"installationMargin,omitempty"`
	// Description of the node.
	Description string `protobuf:"bytes,13,opt,name=description" json:"description,omitempty"`
	// Node is activated by ABP.
	IsABP bool `protobuf:"varint,14,opt,name=isABP" json:"isABP,omitempty"`
}

func (m *GetNodeResponse) Reset()                    { *m = GetNodeResponse{} }
func (m *GetNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNodeResponse) ProtoMessage()               {}
func (*GetNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetNodeResponse) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *GetNodeResponse) GetAppEUI() string {
	if m != nil {
		return m.AppEUI
	}
	return ""
}

func (m *GetNodeResponse) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *GetNodeResponse) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *GetNodeResponse) GetRx1DROffset() uint32 {
	if m != nil {
		return m.Rx1DROffset
	}
	return 0
}

func (m *GetNodeResponse) GetChannelListID() int64 {
	if m != nil {
		return m.ChannelListID
	}
	return 0
}

func (m *GetNodeResponse) GetRxWindow() RXWindow {
	if m != nil {
		return m.RxWindow
	}
	return RXWindow_RX1
}

func (m *GetNodeResponse) GetRx2DR() uint32 {
	if m != nil {
		return m.Rx2DR
	}
	return 0
}

func (m *GetNodeResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetNodeResponse) GetRelaxFCnt() bool {
	if m != nil {
		return m.RelaxFCnt
	}
	return false
}

func (m *GetNodeResponse) GetAdrInterval() uint32 {
	if m != nil {
		return m.AdrInterval
	}
	return 0
}

func (m *GetNodeResponse) GetInstallationMargin() float64 {
	if m != nil {
		return m.InstallationMargin
	}
	return 0
}

func (m *GetNodeResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetNodeResponse) GetIsABP() bool {
	if m != nil {
		return m.IsABP
	}
	return false
}

type DeleteNodeRequest struct {
	// Name of the application owning the node.
	ApplicationName string `protobuf:"bytes,2,opt,name=applicationName" json:"applicationName,omitempty"`
	// Name of the node.
	NodeName string `protobuf:"bytes,1,opt,name=nodeName" json:"nodeName,omitempty"`
}

func (m *DeleteNodeRequest) Reset()                    { *m = DeleteNodeRequest{} }
func (m *DeleteNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeRequest) ProtoMessage()               {}
func (*DeleteNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *DeleteNodeRequest) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *DeleteNodeRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type DeleteNodeResponse struct {
}

func (m *DeleteNodeResponse) Reset()                    { *m = DeleteNodeResponse{} }
func (m *DeleteNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeResponse) ProtoMessage()               {}
func (*DeleteNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type ListNodeRequest struct {
	// Name of the application owning the node.
	ApplicationName string `protobuf:"bytes,3,opt,name=applicationName" json:"applicationName,omitempty"`
	// Max number of nodes to return in the result-set.
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset of the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *ListNodeRequest) Reset()                    { *m = ListNodeRequest{} }
func (m *ListNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*ListNodeRequest) ProtoMessage()               {}
func (*ListNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ListNodeRequest) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *ListNodeRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListNodeRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListNodeResponse struct {
	// Total number of nodes available within the result-set.
	TotalCount int64 `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	// Nodes within this result-set.
	Result []*GetNodeResponse `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListNodeResponse) Reset()                    { *m = ListNodeResponse{} }
func (m *ListNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*ListNodeResponse) ProtoMessage()               {}
func (*ListNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *ListNodeResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListNodeResponse) GetResult() []*GetNodeResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateNodeRequest struct {
	// Name of the application owning the node.
	ApplicationName string `protobuf:"bytes,13,opt,name=applicationName" json:"applicationName,omitempty"`
	// Name of the node to update.
	NodeName string `protobuf:"bytes,14,opt,name=nodeName" json:"nodeName,omitempty"`
	// Hex encoded AppEUI.
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// Hex encoded AppKey.
	AppKey string `protobuf:"bytes,3,opt,name=appKey" json:"appKey,omitempty"`
	// RX delay.
	RxDelay uint32 `protobuf:"varint,4,opt,name=rxDelay" json:"rxDelay,omitempty"`
	// RX1 data-rate offset.
	Rx1DROffset uint32 `protobuf:"varint,5,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	// Channel-list ID used for CFlist (see LoRaWAN regional parameters if this applies to your region).
	ChannelListID int64 `protobuf:"varint,6,opt,name=channelListID" json:"channelListID,omitempty"`
	// RX window to use.
	RxWindow RXWindow `protobuf:"varint,7,opt,name=rxWindow,enum=api.RXWindow" json:"rxWindow,omitempty"`
	// Data-rate to use for RX2.
	Rx2DR uint32 `protobuf:"varint,8,opt,name=rx2DR" json:"rx2DR,omitempty"`
	// Name of the node (note that renaming the node affects its api endpoint)
	Name string `protobuf:"bytes,9,opt,name=name" json:"name,omitempty"`
	// Relax frame-counter mode is enabled.
	RelaxFCnt bool `protobuf:"varint,10,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
	// Interval (in frames) in which the ADR engine may adapt the data-rate of the node (0 = disabled).
	AdrInterval uint32 `protobuf:"varint,11,opt,name=adrInterval" json:"adrInterval,omitempty"`
	// Installation-margin to use for ADR calculation.
	InstallationMargin float64 `protobuf:"fixed64,12,opt,name=installationMargin" json:"installationMargin,omitempty"`
	// Description of the node.
	Description string `protobuf:"bytes,15,opt,name=description" json:"description,omitempty"`
	// Node is activated by ABP.
	IsABP bool `protobuf:"varint,16,opt,name=isABP" json:"isABP,omitempty"`
}

func (m *UpdateNodeRequest) Reset()                    { *m = UpdateNodeRequest{} }
func (m *UpdateNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateNodeRequest) ProtoMessage()               {}
func (*UpdateNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *UpdateNodeRequest) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *UpdateNodeRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *UpdateNodeRequest) GetAppEUI() string {
	if m != nil {
		return m.AppEUI
	}
	return ""
}

func (m *UpdateNodeRequest) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *UpdateNodeRequest) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *UpdateNodeRequest) GetRx1DROffset() uint32 {
	if m != nil {
		return m.Rx1DROffset
	}
	return 0
}

func (m *UpdateNodeRequest) GetChannelListID() int64 {
	if m != nil {
		return m.ChannelListID
	}
	return 0
}

func (m *UpdateNodeRequest) GetRxWindow() RXWindow {
	if m != nil {
		return m.RxWindow
	}
	return RXWindow_RX1
}

func (m *UpdateNodeRequest) GetRx2DR() uint32 {
	if m != nil {
		return m.Rx2DR
	}
	return 0
}

func (m *UpdateNodeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateNodeRequest) GetRelaxFCnt() bool {
	if m != nil {
		return m.RelaxFCnt
	}
	return false
}

func (m *UpdateNodeRequest) GetAdrInterval() uint32 {
	if m != nil {
		return m.AdrInterval
	}
	return 0
}

func (m *UpdateNodeRequest) GetInstallationMargin() float64 {
	if m != nil {
		return m.InstallationMargin
	}
	return 0
}

func (m *UpdateNodeRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateNodeRequest) GetIsABP() bool {
	if m != nil {
		return m.IsABP
	}
	return false
}

type UpdateNodeResponse struct {
}

func (m *UpdateNodeResponse) Reset()                    { *m = UpdateNodeResponse{} }
func (m *UpdateNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateNodeResponse) ProtoMessage()               {}
func (*UpdateNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

type ActivateNodeRequest struct {
	// Name of the application owning the node.
	ApplicationName string `protobuf:"bytes,1,opt,name=applicationName" json:"applicationName,omitempty"`
	// Name of the node to update.
	NodeName string `protobuf:"bytes,2,opt,name=nodeName" json:"nodeName,omitempty"`
	// Hex encoded DevAddr.
	DevAddr string `protobuf:"bytes,3,opt,name=devAddr" json:"devAddr,omitempty"`
	// Hex encoded AppSKey.
	AppSKey string `protobuf:"bytes,4,opt,name=appSKey" json:"appSKey,omitempty"`
	// Hex encoded NwkSKey.
	NwkSKey string `protobuf:"bytes,5,opt,name=nwkSKey" json:"nwkSKey,omitempty"`
	// Uplink frame-counter.
	FCntUp uint32 `protobuf:"varint,6,opt,name=fCntUp" json:"fCntUp,omitempty"`
	// Downlink frame-counter.
	FCntDown uint32 `protobuf:"varint,7,opt,name=fCntDown" json:"fCntDown,omitempty"`
}

func (m *ActivateNodeRequest) Reset()                    { *m = ActivateNodeRequest{} }
func (m *ActivateNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*ActivateNodeRequest) ProtoMessage()               {}
func (*ActivateNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *ActivateNodeRequest) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *ActivateNodeRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *ActivateNodeRequest) GetDevAddr() string {
	if m != nil {
		return m.DevAddr
	}
	return ""
}

func (m *ActivateNodeRequest) GetAppSKey() string {
	if m != nil {
		return m.AppSKey
	}
	return ""
}

func (m *ActivateNodeRequest) GetNwkSKey() string {
	if m != nil {
		return m.NwkSKey
	}
	return ""
}

func (m *ActivateNodeRequest) GetFCntUp() uint32 {
	if m != nil {
		return m.FCntUp
	}
	return 0
}

func (m *ActivateNodeRequest) GetFCntDown() uint32 {
	if m != nil {
		return m.FCntDown
	}
	return 0
}

type ActivateNodeResponse struct {
}

func (m *ActivateNodeResponse) Reset()                    { *m = ActivateNodeResponse{} }
func (m *ActivateNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*ActivateNodeResponse) ProtoMessage()               {}
func (*ActivateNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func init() {
	proto.RegisterType((*CreateNodeRequest)(nil), "api.CreateNodeRequest")
	proto.RegisterType((*CreateNodeResponse)(nil), "api.CreateNodeResponse")
	proto.RegisterType((*GetNodeRequest)(nil), "api.GetNodeRequest")
	proto.RegisterType((*GetNodeResponse)(nil), "api.GetNodeResponse")
	proto.RegisterType((*DeleteNodeRequest)(nil), "api.DeleteNodeRequest")
	proto.RegisterType((*DeleteNodeResponse)(nil), "api.DeleteNodeResponse")
	proto.RegisterType((*ListNodeRequest)(nil), "api.ListNodeRequest")
	proto.RegisterType((*ListNodeResponse)(nil), "api.ListNodeResponse")
	proto.RegisterType((*UpdateNodeRequest)(nil), "api.UpdateNodeRequest")
	proto.RegisterType((*UpdateNodeResponse)(nil), "api.UpdateNodeResponse")
	proto.RegisterType((*ActivateNodeRequest)(nil), "api.ActivateNodeRequest")
	proto.RegisterType((*ActivateNodeResponse)(nil), "api.ActivateNodeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	// Create creates the given node.
	Create(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*CreateNodeResponse, error)
	// Get returns the node for the requested name.
	Get(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error)
	// Delete deletes the node matching the given name.
	Delete(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*DeleteNodeResponse, error)
	// List lists the nodes.
	List(ctx context.Context, in *ListNodeRequest, opts ...grpc.CallOption) (*ListNodeResponse, error)
	// Update updates the node matching the given name.
	Update(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*UpdateNodeResponse, error)
	// Activate (re)activates the node (only when ABP is set to true).
	Activate(ctx context.Context, in *ActivateNodeRequest, opts ...grpc.CallOption) (*ActivateNodeResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Create(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*CreateNodeResponse, error) {
	out := new(CreateNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Get(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error) {
	out := new(GetNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Delete(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*DeleteNodeResponse, error) {
	out := new(DeleteNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) List(ctx context.Context, in *ListNodeRequest, opts ...grpc.CallOption) (*ListNodeResponse, error) {
	out := new(ListNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Update(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*UpdateNodeResponse, error) {
	out := new(UpdateNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Activate(ctx context.Context, in *ActivateNodeRequest, opts ...grpc.CallOption) (*ActivateNodeResponse, error) {
	out := new(ActivateNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Activate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	// Create creates the given node.
	Create(context.Context, *CreateNodeRequest) (*CreateNodeResponse, error)
	// Get returns the node for the requested name.
	Get(context.Context, *GetNodeRequest) (*GetNodeResponse, error)
	// Delete deletes the node matching the given name.
	Delete(context.Context, *DeleteNodeRequest) (*DeleteNodeResponse, error)
	// List lists the nodes.
	List(context.Context, *ListNodeRequest) (*ListNodeResponse, error)
	// Update updates the node matching the given name.
	Update(context.Context, *UpdateNodeRequest) (*UpdateNodeResponse, error)
	// Activate (re)activates the node (only when ABP is set to true).
	Activate(context.Context, *ActivateNodeRequest) (*ActivateNodeResponse, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Create(ctx, req.(*CreateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Get(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Delete(ctx, req.(*DeleteNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).List(ctx, req.(*ListNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Update(ctx, req.(*UpdateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Activate(ctx, req.(*ActivateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Node_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Node_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Node_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Node_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Node_Update_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _Node_Activate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

func init() { proto.RegisterFile("node.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 817 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xec, 0x56, 0x4f, 0x4f, 0x3b, 0x45,
	0x18, 0xce, 0x76, 0xdb, 0x52, 0x5e, 0x68, 0x0b, 0x43, 0xc5, 0xb1, 0x21, 0xa6, 0xd9, 0x78, 0x58,
	0xd0, 0xb4, 0xa1, 0x12, 0x13, 0x89, 0xc6, 0x40, 0xab, 0x48, 0x50, 0x34, 0x6b, 0xf0, 0xcf, 0xcd,
	0xb1, 0x3b, 0xc5, 0x8d, 0xdb, 0x99, 0x75, 0x77, 0x68, 0x8b, 0x84, 0x8b, 0x37, 0xcf, 0x7e, 0x19,
	0x3f, 0x86, 0x89, 0x67, 0x6f, 0x1c, 0xfc, 0x18, 0x66, 0xfe, 0xb4, 0x6c, 0xdb, 0x0d, 0x25, 0x3d,
	0xfc, 0x7e, 0x17, 0x6e, 0x7d, 0x9e, 0x77, 0x98, 0xe7, 0xfd, 0xf3, 0xcc, 0xcb, 0x02, 0x30, 0xee,
	0xd3, 0x66, 0x14, 0x73, 0xc1, 0x91, 0x4d, 0xa2, 0xa0, 0xbe, 0x77, 0xcd, 0xf9, 0x75, 0x48, 0x5b,
	0x24, 0x0a, 0x5a, 0x84, 0x31, 0x2e, 0x88, 0x08, 0x38, 0x4b, 0xf4, 0x91, 0xfa, 0x66, 0x8f, 0x0f,
	0x06, 0x9c, 0x69, 0xe4, 0x3c, 0xd8, 0xb0, 0xdd, 0x89, 0x29, 0x11, 0xf4, 0x92, 0xfb, 0xd4, 0xa3,
	0xbf, 0xde, 0xd0, 0x44, 0x20, 0x17, 0xaa, 0x24, 0x8a, 0xc2, 0xa0, 0xa7, 0xfe, 0xf2, 0x92, 0x0c,
	0x28, 0x2e, 0x37, 0x2c, 0x77, 0xdd, 0x9b, 0xa7, 0xd1, 0x2e, 0x14, 0x7d, 0x3a, 0xfc, 0xf4, 0xea,
	0x1c, 0x5b, 0xea, 0x80, 0x41, 0x92, 0x27, 0x51, 0x24, 0xf9, 0x9c, 0xe6, 0x35, 0x32, 0xfc, 0x05,
	0xbd, 0xc5, 0xf6, 0x94, 0xbf, 0xa0, 0xb7, 0x08, 0xc3, 0x5a, 0x3c, 0xee, 0xd2, 0x90, 0xdc, 0xe2,
	0x7c, 0xc3, 0x72, 0xcb, 0xde, 0x04, 0xa2, 0x06, 0x6c, 0xc4, 0xe3, 0xc3, 0xae, 0xf7, 0x55, 0xbf,
	0x9f, 0x50, 0x81, 0x0b, 0x2a, 0x9a, 0xa6, 0xd0, 0x3b, 0x50, 0xee, 0xfd, 0x4c, 0x18, 0xa3, 0xe1,
	0x17, 0x41, 0x22, 0xce, 0xbb, 0xb8, 0xd8, 0xb0, 0x5c, 0xdb, 0x9b, 0x25, 0xd1, 0x3e, 0x94, 0xe2,
	0xf1, 0x77, 0x01, 0xf3, 0xf9, 0x08, 0xaf, 0x35, 0x2c, 0xb7, 0xd2, 0x2e, 0x37, 0x49, 0x14, 0x34,
	0xbd, 0xef, 0x35, 0xe9, 0x4d, 0xc3, 0xa8, 0x06, 0x85, 0x78, 0xdc, 0xee, 0x7a, 0xb8, 0xa4, 0xc4,
	0x34, 0x40, 0x08, 0xf2, 0x4c, 0x76, 0x62, 0x5d, 0x25, 0xae, 0x7e, 0xa3, 0x3d, 0x58, 0x8f, 0x69,
	0x48, 0xc6, 0x9f, 0x75, 0x98, 0xc0, 0xd0, 0xb0, 0xdc, 0x92, 0xf7, 0x48, 0xc8, 0xd4, 0x89, 0x1f,
	0x9f, 0x33, 0x41, 0xe3, 0x21, 0x09, 0xf1, 0x86, 0x4e, 0x3d, 0x45, 0xa1, 0x26, 0xa0, 0x80, 0x25,
	0x82, 0x84, 0xa1, 0x6a, 0xe9, 0x97, 0x24, 0xbe, 0x0e, 0x18, 0xde, 0x6c, 0x58, 0xae, 0xe5, 0x65,
	0x44, 0xe4, 0x8d, 0x3e, 0x4d, 0x7a, 0x71, 0x10, 0x49, 0x12, 0x57, 0x54, 0x2a, 0x69, 0x4a, 0xe6,
	0x1e, 0x24, 0x27, 0xa7, 0x5f, 0xe3, 0xaa, 0xca, 0x46, 0x03, 0xa7, 0x06, 0x28, 0x3d, 0xe5, 0x24,
	0xe2, 0x2c, 0xa1, 0xce, 0xb7, 0x50, 0x39, 0xa3, 0x62, 0xc9, 0xe0, 0x73, 0xd9, 0x83, 0xaf, 0x43,
	0x49, 0xfa, 0x4e, 0x1d, 0xd1, 0xa3, 0x9f, 0x62, 0xe7, 0x2f, 0x1b, 0xaa, 0xd3, 0x8b, 0xb5, 0xd6,
	0x8b, 0x51, 0x5e, 0xab, 0x51, 0xca, 0x4f, 0x18, 0xa5, 0x92, 0x36, 0xca, 0x0f, 0xb0, 0xdd, 0xa5,
	0x21, 0x5d, 0xba, 0x0e, 0x56, 0x70, 0x45, 0x0d, 0x50, 0xfa, 0x6a, 0xe3, 0xc1, 0x00, 0xaa, 0xb2,
	0xef, 0x4b, 0xe4, 0xec, 0x6c, 0xb9, 0x1a, 0x14, 0xc2, 0x60, 0x10, 0x08, 0xa5, 0x65, 0x7b, 0x1a,
	0x48, 0xeb, 0x70, 0xed, 0x81, 0x9c, 0xa2, 0x0d, 0x72, 0x7e, 0x84, 0xad, 0x47, 0x29, 0x63, 0xcb,
	0xb7, 0x01, 0x04, 0x17, 0x24, 0xec, 0xf0, 0x1b, 0x36, 0xb9, 0x26, 0xc5, 0xa0, 0xf7, 0xa0, 0x18,
	0xd3, 0xe4, 0x26, 0x94, 0x77, 0xd9, 0xee, 0x46, 0xbb, 0xa6, 0xac, 0x30, 0x67, 0x6e, 0xcf, 0x9c,
	0x71, 0xfe, 0xb3, 0x61, 0xfb, 0x2a, 0xf2, 0x57, 0xde, 0xa6, 0xe9, 0xf6, 0x55, 0x66, 0xdb, 0xf7,
	0xf2, 0x50, 0x5e, 0xc9, 0x43, 0xa9, 0x3e, 0xf1, 0x50, 0xb6, 0xe6, 0x36, 0x6a, 0x7a, 0xd2, 0xc6,
	0xcd, 0xff, 0x5a, 0xb0, 0x73, 0xd2, 0x13, 0xc1, 0x70, 0xb9, 0x05, 0xac, 0xe5, 0x16, 0xc8, 0xcd,
	0x59, 0x00, 0xc3, 0x9a, 0x4f, 0x87, 0x27, 0xbe, 0x1f, 0x9b, 0x59, 0x4f, 0xa0, 0x8c, 0x90, 0x28,
	0xfa, 0x46, 0xba, 0x20, 0xaf, 0x23, 0x06, 0xca, 0x08, 0x1b, 0xfd, 0xa2, 0x22, 0x05, 0x1d, 0x31,
	0x50, 0x1a, 0xa7, 0xdf, 0x61, 0xe2, 0x2a, 0x52, 0xd3, 0x2d, 0x7b, 0x06, 0xc9, 0x0c, 0xe4, 0xaf,
	0x2e, 0x1f, 0x31, 0x35, 0xd6, 0xb2, 0x37, 0xc5, 0xce, 0x2e, 0xd4, 0x66, 0xcb, 0xd3, 0x75, 0xb7,
	0xff, 0x2e, 0x40, 0x5e, 0x12, 0x88, 0x43, 0x51, 0xff, 0xa3, 0x41, 0xbb, 0xca, 0x0b, 0x0b, 0xdf,
	0x16, 0xf5, 0x37, 0x17, 0x78, 0xd3, 0xbb, 0xa3, 0xdf, 0xff, 0x79, 0xf8, 0x33, 0xd7, 0x74, 0xf6,
	0xf5, 0x87, 0xcb, 0x63, 0x5f, 0x92, 0xd6, 0xdd, 0x5c, 0x97, 0xee, 0x5b, 0xb2, 0x27, 0xc9, 0xb1,
	0x75, 0x80, 0x18, 0xd8, 0x67, 0x54, 0xa0, 0x9d, 0xd9, 0x77, 0xa9, 0xa5, 0x32, 0x1f, 0xab, 0xf3,
	0x91, 0xd2, 0xf9, 0x00, 0x1d, 0x3d, 0x5b, 0xa7, 0x75, 0x37, 0x19, 0xc1, 0x3d, 0x1a, 0x41, 0x51,
	0x6f, 0x31, 0x53, 0xe0, 0xc2, 0xb6, 0x34, 0x05, 0x66, 0xac, 0x3a, 0x23, 0x7c, 0xb0, 0x9a, 0x70,
	0x1f, 0xf2, 0xf2, 0xdd, 0x21, 0x5d, 0xd4, 0xdc, 0xce, 0xac, 0xbf, 0x31, 0xc7, 0x1a, 0xc9, 0x43,
	0x25, 0xf9, 0x2e, 0x7a, 0x7e, 0x4f, 0xd1, 0x6f, 0x50, 0xd4, 0xc6, 0x36, 0x05, 0x2e, 0xec, 0x33,
	0x53, 0x60, 0x86, 0xfb, 0x3f, 0x51, 0x6a, 0x1f, 0xd6, 0x57, 0x2a, 0x50, 0x0e, 0xf3, 0x0f, 0x0b,
	0x4a, 0x13, 0x7f, 0x21, 0xac, 0x64, 0x32, 0x5e, 0x53, 0xfd, 0xad, 0x8c, 0x88, 0x49, 0xe1, 0x73,
	0x95, 0xc2, 0xa9, 0xf3, 0xf1, 0x2a, 0x29, 0xb4, 0x88, 0xb9, 0xf2, 0xd8, 0x3a, 0xf8, 0xa9, 0xa8,
	0x3e, 0x90, 0xdf, 0xff, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xf4, 0x69, 0x97, 0x5f, 0x0b, 0x00,
	0x00,
}
