// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.11.2
// source: proto/music.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"` // Group name
	Song  string `protobuf:"bytes,2,opt,name=song,proto3" json:"song,omitempty"`   // Song name
}

func (x *AddSongRequest) Reset() {
	*x = AddSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSongRequest) ProtoMessage() {}

func (x *AddSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSongRequest.ProtoReflect.Descriptor instead.
func (*AddSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{0}
}

func (x *AddSongRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *AddSongRequest) GetSong() string {
	if x != nil {
		return x.Song
	}
	return ""
}

type AddSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongId string `protobuf:"bytes,1,opt,name=song_id,json=songId,proto3" json:"song_id,omitempty"` // Song ID
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`               // "success" or "failure"
}

func (x *AddSongResponse) Reset() {
	*x = AddSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSongResponse) ProtoMessage() {}

func (x *AddSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSongResponse.ProtoReflect.Descriptor instead.
func (*AddSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{1}
}

func (x *AddSongResponse) GetSongId() string {
	if x != nil {
		return x.SongId
	}
	return ""
}

func (x *AddSongResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongId string `protobuf:"bytes,1,opt,name=song_id,json=songId,proto3" json:"song_id,omitempty"`
	Group  string `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"` // Group name (update)
	Song   string `protobuf:"bytes,3,opt,name=song,proto3" json:"song,omitempty"`   // Song name (update)
}

func (x *UpdateSongRequest) Reset() {
	*x = UpdateSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSongRequest) ProtoMessage() {}

func (x *UpdateSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSongRequest.ProtoReflect.Descriptor instead.
func (*UpdateSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSongRequest) GetSongId() string {
	if x != nil {
		return x.SongId
	}
	return ""
}

func (x *UpdateSongRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *UpdateSongRequest) GetSong() string {
	if x != nil {
		return x.Song
	}
	return ""
}

type UpdateSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // Status indicating success or failure of update
}

func (x *UpdateSongResponse) Reset() {
	*x = UpdateSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSongResponse) ProtoMessage() {}

func (x *UpdateSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSongResponse.ProtoReflect.Descriptor instead.
func (*UpdateSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSongResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type DeleteSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongId string `protobuf:"bytes,1,opt,name=song_id,json=songId,proto3" json:"song_id,omitempty"`
}

func (x *DeleteSongRequest) Reset() {
	*x = DeleteSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSongRequest) ProtoMessage() {}

func (x *DeleteSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSongRequest.ProtoReflect.Descriptor instead.
func (*DeleteSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSongRequest) GetSongId() string {
	if x != nil {
		return x.SongId
	}
	return ""
}

type DeleteSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // "success" or "failure"
}

func (x *DeleteSongResponse) Reset() {
	*x = DeleteSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSongResponse) ProtoMessage() {}

func (x *DeleteSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSongResponse.ProtoReflect.Descriptor instead.
func (*DeleteSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSongResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongId string `protobuf:"bytes,1,opt,name=song_id,json=songId,proto3" json:"song_id,omitempty"`
}

func (x *GetSongRequest) Reset() {
	*x = GetSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSongRequest) ProtoMessage() {}

func (x *GetSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSongRequest.ProtoReflect.Descriptor instead.
func (*GetSongRequest) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{6}
}

func (x *GetSongRequest) GetSongId() string {
	if x != nil {
		return x.SongId
	}
	return ""
}

type GetSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongId      string `protobuf:"bytes,1,opt,name=song_id,json=songId,proto3" json:"song_id,omitempty"`
	Group       string `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Song        string `protobuf:"bytes,3,opt,name=song,proto3" json:"song,omitempty"`
	ReleaseDate string `protobuf:"bytes,4,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"` // Release date
	Link        string `protobuf:"bytes,5,opt,name=link,proto3" json:"link,omitempty"`                                  // Song link
}

func (x *GetSongResponse) Reset() {
	*x = GetSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSongResponse) ProtoMessage() {}

func (x *GetSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSongResponse.ProtoReflect.Descriptor instead.
func (*GetSongResponse) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{7}
}

func (x *GetSongResponse) GetSongId() string {
	if x != nil {
		return x.SongId
	}
	return ""
}

func (x *GetSongResponse) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GetSongResponse) GetSong() string {
	if x != nil {
		return x.Song
	}
	return ""
}

func (x *GetSongResponse) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *GetSongResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type GetSongsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group    string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`                        // Group name (optional)
	Page     int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`                         // Page number
	PageSize int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"` // Number of items per page
}

func (x *GetSongsRequest) Reset() {
	*x = GetSongsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSongsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSongsRequest) ProtoMessage() {}

func (x *GetSongsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSongsRequest.ProtoReflect.Descriptor instead.
func (*GetSongsRequest) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{8}
}

func (x *GetSongsRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GetSongsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetSongsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetSongsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Songs []*Song `protobuf:"bytes,1,rep,name=songs,proto3" json:"songs,omitempty"` // List of songs
}

func (x *GetSongsResponse) Reset() {
	*x = GetSongsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSongsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSongsResponse) ProtoMessage() {}

func (x *GetSongsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSongsResponse.ProtoReflect.Descriptor instead.
func (*GetSongsResponse) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{9}
}

func (x *GetSongsResponse) GetSongs() []*Song {
	if x != nil {
		return x.Songs
	}
	return nil
}

type GetSongLyricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongId   string `protobuf:"bytes,1,opt,name=song_id,json=songId,proto3" json:"song_id,omitempty"`        // Song ID
	Page     int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`                         // Page number
	PageSize int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"` // Number of verses per page
}

func (x *GetSongLyricsRequest) Reset() {
	*x = GetSongLyricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSongLyricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSongLyricsRequest) ProtoMessage() {}

func (x *GetSongLyricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSongLyricsRequest.ProtoReflect.Descriptor instead.
func (*GetSongLyricsRequest) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{10}
}

func (x *GetSongLyricsRequest) GetSongId() string {
	if x != nil {
		return x.SongId
	}
	return ""
}

func (x *GetSongLyricsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetSongLyricsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetSongLyricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lyrics []*SongLyrics `protobuf:"bytes,1,rep,name=lyrics,proto3" json:"lyrics,omitempty"` // List of verses
}

func (x *GetSongLyricsResponse) Reset() {
	*x = GetSongLyricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSongLyricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSongLyricsResponse) ProtoMessage() {}

func (x *GetSongLyricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSongLyricsResponse.ProtoReflect.Descriptor instead.
func (*GetSongLyricsResponse) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{11}
}

func (x *GetSongLyricsResponse) GetLyrics() []*SongLyrics {
	if x != nil {
		return x.Lyrics
	}
	return nil
}

type Song struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongId      string `protobuf:"bytes,1,opt,name=song_id,json=songId,proto3" json:"song_id,omitempty"`
	Group       string `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Song        string `protobuf:"bytes,3,opt,name=song,proto3" json:"song,omitempty"`
	ReleaseDate string `protobuf:"bytes,4,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	Link        string `protobuf:"bytes,5,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *Song) Reset() {
	*x = Song{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Song) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Song) ProtoMessage() {}

func (x *Song) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Song.ProtoReflect.Descriptor instead.
func (*Song) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{12}
}

func (x *Song) GetSongId() string {
	if x != nil {
		return x.SongId
	}
	return ""
}

func (x *Song) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Song) GetSong() string {
	if x != nil {
		return x.Song
	}
	return ""
}

func (x *Song) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *Song) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type SongLyrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerseNumber int32  `protobuf:"varint,1,opt,name=verse_number,json=verseNumber,proto3" json:"verse_number,omitempty"` // Verse number
	VerseText   string `protobuf:"bytes,2,opt,name=verse_text,json=verseText,proto3" json:"verse_text,omitempty"`        // Verse text
}

func (x *SongLyrics) Reset() {
	*x = SongLyrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_music_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SongLyrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongLyrics) ProtoMessage() {}

func (x *SongLyrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_music_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongLyrics.ProtoReflect.Descriptor instead.
func (*SongLyrics) Descriptor() ([]byte, []int) {
	return file_proto_music_proto_rawDescGZIP(), []int{13}
}

func (x *SongLyrics) GetVerseNumber() int32 {
	if x != nil {
		return x.VerseNumber
	}
	return 0
}

func (x *SongLyrics) GetVerseText() string {
	if x != nil {
		return x.VerseText
	}
	return ""
}

var File_proto_music_proto protoreflect.FileDescriptor

var file_proto_music_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x22,
	0x42, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x56, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x6e, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x22, 0x2c, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64,
	0x22, 0x8b, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x58,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53,
	0x6f, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05,
	0x73, 0x6f, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x53, 0x6f,
	0x6e, 0x67, 0x52, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x22, 0x60, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x53, 0x6f, 0x6e, 0x67, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x3c, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x6c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x4c, 0x79, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x06, 0x6c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x04, 0x53, 0x6f,
	0x6e, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6f, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x4e, 0x0a, 0x0a,
	0x53, 0x6f, 0x6e, 0x67, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x65, 0x54, 0x65, 0x78, 0x74, 0x32, 0xc8, 0x02, 0x0a,
	0x0b, 0x53, 0x6f, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x6f,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53,
	0x6f, 0x6e, 0x67, 0x73, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x53, 0x6f, 0x6e, 0x67, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x6f, 0x6e, 0x67, 0x4c, 0x79, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x4c, 0x79, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x12,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_music_proto_rawDescOnce sync.Once
	file_proto_music_proto_rawDescData = file_proto_music_proto_rawDesc
)

func file_proto_music_proto_rawDescGZIP() []byte {
	file_proto_music_proto_rawDescOnce.Do(func() {
		file_proto_music_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_music_proto_rawDescData)
	})
	return file_proto_music_proto_rawDescData
}

var file_proto_music_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_music_proto_goTypes = []any{
	(*AddSongRequest)(nil),        // 0: AddSongRequest
	(*AddSongResponse)(nil),       // 1: AddSongResponse
	(*UpdateSongRequest)(nil),     // 2: UpdateSongRequest
	(*UpdateSongResponse)(nil),    // 3: UpdateSongResponse
	(*DeleteSongRequest)(nil),     // 4: DeleteSongRequest
	(*DeleteSongResponse)(nil),    // 5: DeleteSongResponse
	(*GetSongRequest)(nil),        // 6: GetSongRequest
	(*GetSongResponse)(nil),       // 7: GetSongResponse
	(*GetSongsRequest)(nil),       // 8: GetSongsRequest
	(*GetSongsResponse)(nil),      // 9: GetSongsResponse
	(*GetSongLyricsRequest)(nil),  // 10: GetSongLyricsRequest
	(*GetSongLyricsResponse)(nil), // 11: GetSongLyricsResponse
	(*Song)(nil),                  // 12: Song
	(*SongLyrics)(nil),            // 13: SongLyrics
}
var file_proto_music_proto_depIdxs = []int32{
	12, // 0: GetSongsResponse.songs:type_name -> Song
	13, // 1: GetSongLyricsResponse.lyrics:type_name -> SongLyrics
	0,  // 2: SongService.AddSong:input_type -> AddSongRequest
	6,  // 3: SongService.GetSong:input_type -> GetSongRequest
	8,  // 4: SongService.GetSongs:input_type -> GetSongsRequest
	10, // 5: SongService.GetSongLyrics:input_type -> GetSongLyricsRequest
	2,  // 6: SongService.UpdateSong:input_type -> UpdateSongRequest
	4,  // 7: SongService.DeleteSong:input_type -> DeleteSongRequest
	1,  // 8: SongService.AddSong:output_type -> AddSongResponse
	7,  // 9: SongService.GetSong:output_type -> GetSongResponse
	9,  // 10: SongService.GetSongs:output_type -> GetSongsResponse
	11, // 11: SongService.GetSongLyrics:output_type -> GetSongLyricsResponse
	3,  // 12: SongService.UpdateSong:output_type -> UpdateSongResponse
	5,  // 13: SongService.DeleteSong:output_type -> DeleteSongResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_music_proto_init() }
func file_proto_music_proto_init() {
	if File_proto_music_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_music_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddSongRequest); i {
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
		file_proto_music_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddSongResponse); i {
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
		file_proto_music_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateSongRequest); i {
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
		file_proto_music_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateSongResponse); i {
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
		file_proto_music_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteSongRequest); i {
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
		file_proto_music_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteSongResponse); i {
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
		file_proto_music_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetSongRequest); i {
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
		file_proto_music_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetSongResponse); i {
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
		file_proto_music_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetSongsRequest); i {
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
		file_proto_music_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetSongsResponse); i {
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
		file_proto_music_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*GetSongLyricsRequest); i {
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
		file_proto_music_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*GetSongLyricsResponse); i {
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
		file_proto_music_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*Song); i {
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
		file_proto_music_proto_msgTypes[13].Exporter = func(v any, i int) any {
			switch v := v.(*SongLyrics); i {
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
			RawDescriptor: file_proto_music_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_music_proto_goTypes,
		DependencyIndexes: file_proto_music_proto_depIdxs,
		MessageInfos:      file_proto_music_proto_msgTypes,
	}.Build()
	File_proto_music_proto = out.File
	file_proto_music_proto_rawDesc = nil
	file_proto_music_proto_goTypes = nil
	file_proto_music_proto_depIdxs = nil
}