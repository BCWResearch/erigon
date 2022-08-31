// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: remote/kv.proto

package remote

import (
	types "github.com/ledgerwatch/erigon-lib/gointerfaces/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Op int32

const (
	Op_FIRST           Op = 0
	Op_FIRST_DUP       Op = 1
	Op_SEEK            Op = 2
	Op_SEEK_BOTH       Op = 3
	Op_CURRENT         Op = 4
	Op_LAST            Op = 6
	Op_LAST_DUP        Op = 7
	Op_NEXT            Op = 8
	Op_NEXT_DUP        Op = 9
	Op_NEXT_NO_DUP     Op = 11
	Op_PREV            Op = 12
	Op_PREV_DUP        Op = 13
	Op_PREV_NO_DUP     Op = 14
	Op_SEEK_EXACT      Op = 15
	Op_SEEK_BOTH_EXACT Op = 16
	Op_OPEN            Op = 30
	Op_CLOSE           Op = 31
	Op_OPEN_DUP_SORT   Op = 32
)

// Enum value maps for Op.
var (
	Op_name = map[int32]string{
		0:  "FIRST",
		1:  "FIRST_DUP",
		2:  "SEEK",
		3:  "SEEK_BOTH",
		4:  "CURRENT",
		6:  "LAST",
		7:  "LAST_DUP",
		8:  "NEXT",
		9:  "NEXT_DUP",
		11: "NEXT_NO_DUP",
		12: "PREV",
		13: "PREV_DUP",
		14: "PREV_NO_DUP",
		15: "SEEK_EXACT",
		16: "SEEK_BOTH_EXACT",
		30: "OPEN",
		31: "CLOSE",
		32: "OPEN_DUP_SORT",
	}
	Op_value = map[string]int32{
		"FIRST":           0,
		"FIRST_DUP":       1,
		"SEEK":            2,
		"SEEK_BOTH":       3,
		"CURRENT":         4,
		"LAST":            6,
		"LAST_DUP":        7,
		"NEXT":            8,
		"NEXT_DUP":        9,
		"NEXT_NO_DUP":     11,
		"PREV":            12,
		"PREV_DUP":        13,
		"PREV_NO_DUP":     14,
		"SEEK_EXACT":      15,
		"SEEK_BOTH_EXACT": 16,
		"OPEN":            30,
		"CLOSE":           31,
		"OPEN_DUP_SORT":   32,
	}
)

func (x Op) Enum() *Op {
	p := new(Op)
	*p = x
	return p
}

func (x Op) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Op) Descriptor() protoreflect.EnumDescriptor {
	return file_remote_kv_proto_enumTypes[0].Descriptor()
}

func (Op) Type() protoreflect.EnumType {
	return &file_remote_kv_proto_enumTypes[0]
}

func (x Op) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Op.Descriptor instead.
func (Op) EnumDescriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{0}
}

type Action int32

const (
	Action_STORAGE     Action = 0 // Change only in the storage
	Action_UPSERT      Action = 1 // Change of balance or nonce (and optionally storage)
	Action_CODE        Action = 2 // Change of code (and optionally storage)
	Action_UPSERT_CODE Action = 3 // Change in (balance or nonce) and code (and optinally storage)
	Action_REMOVE      Action = 4 // Account is deleted
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "STORAGE",
		1: "UPSERT",
		2: "CODE",
		3: "UPSERT_CODE",
		4: "REMOVE",
	}
	Action_value = map[string]int32{
		"STORAGE":     0,
		"UPSERT":      1,
		"CODE":        2,
		"UPSERT_CODE": 3,
		"REMOVE":      4,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_remote_kv_proto_enumTypes[1].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_remote_kv_proto_enumTypes[1]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{1}
}

type Direction int32

const (
	Direction_FORWARD Direction = 0
	Direction_UNWIND  Direction = 1
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "FORWARD",
		1: "UNWIND",
	}
	Direction_value = map[string]int32{
		"FORWARD": 0,
		"UNWIND":  1,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_remote_kv_proto_enumTypes[2].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_remote_kv_proto_enumTypes[2]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{2}
}

type Cursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op         Op     `protobuf:"varint,1,opt,name=op,proto3,enum=remote.Op" json:"op,omitempty"`
	BucketName string `protobuf:"bytes,2,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	Cursor     uint32 `protobuf:"varint,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	K          []byte `protobuf:"bytes,4,opt,name=k,proto3" json:"k,omitempty"`
	V          []byte `protobuf:"bytes,5,opt,name=v,proto3" json:"v,omitempty"`
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_remote_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{0}
}

func (x *Cursor) GetOp() Op {
	if x != nil {
		return x.Op
	}
	return Op_FIRST
}

func (x *Cursor) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *Cursor) GetCursor() uint32 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *Cursor) GetK() []byte {
	if x != nil {
		return x.K
	}
	return nil
}

func (x *Cursor) GetV() []byte {
	if x != nil {
		return x.V
	}
	return nil
}

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	K        []byte `protobuf:"bytes,1,opt,name=k,proto3" json:"k,omitempty"`
	V        []byte `protobuf:"bytes,2,opt,name=v,proto3" json:"v,omitempty"`
	CursorID uint32 `protobuf:"varint,3,opt,name=cursorID,proto3" json:"cursorID,omitempty"` // send once after new cursor open
	TxID     uint64 `protobuf:"varint,4,opt,name=txID,proto3" json:"txID,omitempty"`         // send once after tx open. mdbx's tx.ID() - id of write transaction in db - where this changes happened
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_remote_kv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{1}
}

func (x *Pair) GetK() []byte {
	if x != nil {
		return x.K
	}
	return nil
}

func (x *Pair) GetV() []byte {
	if x != nil {
		return x.V
	}
	return nil
}

func (x *Pair) GetCursorID() uint32 {
	if x != nil {
		return x.CursorID
	}
	return 0
}

func (x *Pair) GetTxID() uint64 {
	if x != nil {
		return x.TxID
	}
	return 0
}

type StorageChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *types.H256 `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Data     []byte      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StorageChange) Reset() {
	*x = StorageChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_kv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageChange) ProtoMessage() {}

func (x *StorageChange) ProtoReflect() protoreflect.Message {
	mi := &file_remote_kv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageChange.ProtoReflect.Descriptor instead.
func (*StorageChange) Descriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{2}
}

func (x *StorageChange) GetLocation() *types.H256 {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *StorageChange) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type AccountChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address        *types.H160      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Incarnation    uint64           `protobuf:"varint,2,opt,name=incarnation,proto3" json:"incarnation,omitempty"`
	Action         Action           `protobuf:"varint,3,opt,name=action,proto3,enum=remote.Action" json:"action,omitempty"`
	Data           []byte           `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"` // nil if there is no UPSERT in action
	Code           []byte           `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"` // nil if there is no CODE in action
	StorageChanges []*StorageChange `protobuf:"bytes,6,rep,name=storageChanges,proto3" json:"storageChanges,omitempty"`
}

func (x *AccountChange) Reset() {
	*x = AccountChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_kv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountChange) ProtoMessage() {}

func (x *AccountChange) ProtoReflect() protoreflect.Message {
	mi := &file_remote_kv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountChange.ProtoReflect.Descriptor instead.
func (*AccountChange) Descriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{3}
}

func (x *AccountChange) GetAddress() *types.H160 {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *AccountChange) GetIncarnation() uint64 {
	if x != nil {
		return x.Incarnation
	}
	return 0
}

func (x *AccountChange) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_STORAGE
}

func (x *AccountChange) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *AccountChange) GetCode() []byte {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *AccountChange) GetStorageChanges() []*StorageChange {
	if x != nil {
		return x.StorageChanges
	}
	return nil
}

// StateChangeBatch - list of StateDiff done in one DB transaction
type StateChangeBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatabaseViewID      uint64         `protobuf:"varint,1,opt,name=databaseViewID,proto3" json:"databaseViewID,omitempty"` // mdbx's tx.ID() - id of write transaction in db - where this changes happened
	ChangeBatch         []*StateChange `protobuf:"bytes,2,rep,name=changeBatch,proto3" json:"changeBatch,omitempty"`
	PendingBlockBaseFee uint64         `protobuf:"varint,3,opt,name=pendingBlockBaseFee,proto3" json:"pendingBlockBaseFee,omitempty"` // BaseFee of the next block to be produced
	BlockGasLimit       uint64         `protobuf:"varint,4,opt,name=blockGasLimit,proto3" json:"blockGasLimit,omitempty"`             // GasLimit of the latest block - proxy for the gas limit of the next block to be produced
}

func (x *StateChangeBatch) Reset() {
	*x = StateChangeBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_kv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateChangeBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateChangeBatch) ProtoMessage() {}

func (x *StateChangeBatch) ProtoReflect() protoreflect.Message {
	mi := &file_remote_kv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateChangeBatch.ProtoReflect.Descriptor instead.
func (*StateChangeBatch) Descriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{4}
}

func (x *StateChangeBatch) GetDatabaseViewID() uint64 {
	if x != nil {
		return x.DatabaseViewID
	}
	return 0
}

func (x *StateChangeBatch) GetChangeBatch() []*StateChange {
	if x != nil {
		return x.ChangeBatch
	}
	return nil
}

func (x *StateChangeBatch) GetPendingBlockBaseFee() uint64 {
	if x != nil {
		return x.PendingBlockBaseFee
	}
	return 0
}

func (x *StateChangeBatch) GetBlockGasLimit() uint64 {
	if x != nil {
		return x.BlockGasLimit
	}
	return 0
}

// StateChange - changes done by 1 block or by 1 unwind
type StateChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Direction   Direction        `protobuf:"varint,1,opt,name=direction,proto3,enum=remote.Direction" json:"direction,omitempty"`
	BlockHeight uint64           `protobuf:"varint,2,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	BlockHash   *types.H256      `protobuf:"bytes,3,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	Changes     []*AccountChange `protobuf:"bytes,4,rep,name=changes,proto3" json:"changes,omitempty"`
	Txs         [][]byte         `protobuf:"bytes,5,rep,name=txs,proto3" json:"txs,omitempty"` // enable by withTransactions=true
}

func (x *StateChange) Reset() {
	*x = StateChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_kv_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateChange) ProtoMessage() {}

func (x *StateChange) ProtoReflect() protoreflect.Message {
	mi := &file_remote_kv_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateChange.ProtoReflect.Descriptor instead.
func (*StateChange) Descriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{5}
}

func (x *StateChange) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_FORWARD
}

func (x *StateChange) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *StateChange) GetBlockHash() *types.H256 {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *StateChange) GetChanges() []*AccountChange {
	if x != nil {
		return x.Changes
	}
	return nil
}

func (x *StateChange) GetTxs() [][]byte {
	if x != nil {
		return x.Txs
	}
	return nil
}

type StateChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WithStorage      bool `protobuf:"varint,1,opt,name=withStorage,proto3" json:"withStorage,omitempty"`
	WithTransactions bool `protobuf:"varint,2,opt,name=withTransactions,proto3" json:"withTransactions,omitempty"`
}

func (x *StateChangeRequest) Reset() {
	*x = StateChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_kv_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateChangeRequest) ProtoMessage() {}

func (x *StateChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_kv_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateChangeRequest.ProtoReflect.Descriptor instead.
func (*StateChangeRequest) Descriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{6}
}

func (x *StateChangeRequest) GetWithStorage() bool {
	if x != nil {
		return x.WithStorage
	}
	return false
}

func (x *StateChangeRequest) GetWithTransactions() bool {
	if x != nil {
		return x.WithTransactions
	}
	return false
}

type SnapshotsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SnapshotsRequest) Reset() {
	*x = SnapshotsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_kv_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotsRequest) ProtoMessage() {}

func (x *SnapshotsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_kv_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotsRequest.ProtoReflect.Descriptor instead.
func (*SnapshotsRequest) Descriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{7}
}

type SnapshotsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *SnapshotsReply) Reset() {
	*x = SnapshotsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_kv_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotsReply) ProtoMessage() {}

func (x *SnapshotsReply) ProtoReflect() protoreflect.Message {
	mi := &file_remote_kv_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotsReply.ProtoReflect.Descriptor instead.
func (*SnapshotsReply) Descriptor() ([]byte, []int) {
	return file_remote_kv_proto_rawDescGZIP(), []int{8}
}

func (x *SnapshotsReply) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_remote_kv_proto protoreflect.FileDescriptor

var file_remote_kv_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x06, 0x43, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0a, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x4f, 0x70, 0x52, 0x02, 0x6f, 0x70, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x01, 0x6b, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x01, 0x76, 0x22, 0x52, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x6b, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x76, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x74, 0x78, 0x49, 0x44, 0x22, 0x4c, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x48, 0x32, 0x35, 0x36, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xe7, 0x01, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x48, 0x31, 0x36, 0x30, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x69, 0x6e, 0x63, 0x61, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x69, 0x6e, 0x63, 0x61, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x26, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x3d, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x22,
	0xc9, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x56, 0x69, 0x65, 0x77, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x56, 0x69, 0x65, 0x77, 0x49, 0x44, 0x12, 0x35, 0x0a, 0x0b,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x42, 0x61, 0x73, 0x65, 0x46, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x13, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x61,
	0x73, 0x65, 0x46, 0x65, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x47, 0x61,
	0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x47, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xce, 0x01, 0x0a, 0x0b,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x29,
	0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x32, 0x35, 0x36, 0x52, 0x09,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x78,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x74, 0x78, 0x73, 0x22, 0x62, 0x0a, 0x12,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x77, 0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x77, 0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x12, 0x0a, 0x10, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x0e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2a, 0xfb, 0x01, 0x0a,
	0x02, 0x4f, 0x70, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x44, 0x55, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x53, 0x45, 0x45, 0x4b, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x45, 0x45, 0x4b, 0x5f,
	0x42, 0x4f, 0x54, 0x48, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x55, 0x52, 0x52, 0x45, 0x4e,
	0x54, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x41, 0x53, 0x54, 0x10, 0x06, 0x12, 0x0c, 0x0a,
	0x08, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x44, 0x55, 0x50, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x45, 0x58, 0x54, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x45, 0x58, 0x54, 0x5f, 0x44, 0x55,
	0x50, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x45, 0x58, 0x54, 0x5f, 0x4e, 0x4f, 0x5f, 0x44,
	0x55, 0x50, 0x10, 0x0b, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x52, 0x45, 0x56, 0x10, 0x0c, 0x12, 0x0c,
	0x0a, 0x08, 0x50, 0x52, 0x45, 0x56, 0x5f, 0x44, 0x55, 0x50, 0x10, 0x0d, 0x12, 0x0f, 0x0a, 0x0b,
	0x50, 0x52, 0x45, 0x56, 0x5f, 0x4e, 0x4f, 0x5f, 0x44, 0x55, 0x50, 0x10, 0x0e, 0x12, 0x0e, 0x0a,
	0x0a, 0x53, 0x45, 0x45, 0x4b, 0x5f, 0x45, 0x58, 0x41, 0x43, 0x54, 0x10, 0x0f, 0x12, 0x13, 0x0a,
	0x0f, 0x53, 0x45, 0x45, 0x4b, 0x5f, 0x42, 0x4f, 0x54, 0x48, 0x5f, 0x45, 0x58, 0x41, 0x43, 0x54,
	0x10, 0x10, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x1e, 0x12, 0x09, 0x0a, 0x05,
	0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x1f, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x50, 0x45, 0x4e, 0x5f,
	0x44, 0x55, 0x50, 0x5f, 0x53, 0x4f, 0x52, 0x54, 0x10, 0x20, 0x2a, 0x48, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x53, 0x45, 0x52, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x50, 0x53, 0x45, 0x52,
	0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f,
	0x56, 0x45, 0x10, 0x04, 0x2a, 0x24, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x4f, 0x52, 0x57, 0x41, 0x52, 0x44, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x55, 0x4e, 0x57, 0x49, 0x4e, 0x44, 0x10, 0x01, 0x32, 0xeb, 0x01, 0x0a, 0x02, 0x4b,
	0x56, 0x12, 0x36, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x02, 0x54, 0x78, 0x12,
	0x0e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x1a,
	0x0c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x28, 0x01, 0x30,
	0x01, 0x12, 0x46, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x30, 0x01, 0x12, 0x3d, 0x0a, 0x09, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x3b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_remote_kv_proto_rawDescOnce sync.Once
	file_remote_kv_proto_rawDescData = file_remote_kv_proto_rawDesc
)

func file_remote_kv_proto_rawDescGZIP() []byte {
	file_remote_kv_proto_rawDescOnce.Do(func() {
		file_remote_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_kv_proto_rawDescData)
	})
	return file_remote_kv_proto_rawDescData
}

var file_remote_kv_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_remote_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_remote_kv_proto_goTypes = []interface{}{
	(Op)(0),                    // 0: remote.Op
	(Action)(0),                // 1: remote.Action
	(Direction)(0),             // 2: remote.Direction
	(*Cursor)(nil),             // 3: remote.Cursor
	(*Pair)(nil),               // 4: remote.Pair
	(*StorageChange)(nil),      // 5: remote.StorageChange
	(*AccountChange)(nil),      // 6: remote.AccountChange
	(*StateChangeBatch)(nil),   // 7: remote.StateChangeBatch
	(*StateChange)(nil),        // 8: remote.StateChange
	(*StateChangeRequest)(nil), // 9: remote.StateChangeRequest
	(*SnapshotsRequest)(nil),   // 10: remote.SnapshotsRequest
	(*SnapshotsReply)(nil),     // 11: remote.SnapshotsReply
	(*types.H256)(nil),         // 12: types.H256
	(*types.H160)(nil),         // 13: types.H160
	(*emptypb.Empty)(nil),      // 14: google.protobuf.Empty
	(*types.VersionReply)(nil), // 15: types.VersionReply
}
var file_remote_kv_proto_depIdxs = []int32{
	0,  // 0: remote.Cursor.op:type_name -> remote.Op
	12, // 1: remote.StorageChange.location:type_name -> types.H256
	13, // 2: remote.AccountChange.address:type_name -> types.H160
	1,  // 3: remote.AccountChange.action:type_name -> remote.Action
	5,  // 4: remote.AccountChange.storageChanges:type_name -> remote.StorageChange
	8,  // 5: remote.StateChangeBatch.changeBatch:type_name -> remote.StateChange
	2,  // 6: remote.StateChange.direction:type_name -> remote.Direction
	12, // 7: remote.StateChange.blockHash:type_name -> types.H256
	6,  // 8: remote.StateChange.changes:type_name -> remote.AccountChange
	14, // 9: remote.KV.Version:input_type -> google.protobuf.Empty
	3,  // 10: remote.KV.Tx:input_type -> remote.Cursor
	9,  // 11: remote.KV.StateChanges:input_type -> remote.StateChangeRequest
	10, // 12: remote.KV.Snapshots:input_type -> remote.SnapshotsRequest
	15, // 13: remote.KV.Version:output_type -> types.VersionReply
	4,  // 14: remote.KV.Tx:output_type -> remote.Pair
	7,  // 15: remote.KV.StateChanges:output_type -> remote.StateChangeBatch
	11, // 16: remote.KV.Snapshots:output_type -> remote.SnapshotsReply
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_remote_kv_proto_init() }
func file_remote_kv_proto_init() {
	if File_remote_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remote_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cursor); i {
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
		file_remote_kv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
		file_remote_kv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageChange); i {
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
		file_remote_kv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountChange); i {
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
		file_remote_kv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateChangeBatch); i {
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
		file_remote_kv_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateChange); i {
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
		file_remote_kv_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateChangeRequest); i {
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
		file_remote_kv_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotsRequest); i {
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
		file_remote_kv_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotsReply); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_remote_kv_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_remote_kv_proto_goTypes,
		DependencyIndexes: file_remote_kv_proto_depIdxs,
		EnumInfos:         file_remote_kv_proto_enumTypes,
		MessageInfos:      file_remote_kv_proto_msgTypes,
	}.Build()
	File_remote_kv_proto = out.File
	file_remote_kv_proto_rawDesc = nil
	file_remote_kv_proto_goTypes = nil
	file_remote_kv_proto_depIdxs = nil
}
