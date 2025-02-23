// Code generated by gowsdl DO NOT EDIT.

package attachment24

import (
	"context"
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type ImageType string

const (
	ImageTypeUnknown ImageType = "Unknown"

	ImageTypeWmf ImageType = "Wmf"

	ImageTypePng ImageType = "Png"

	ImageTypeTiff ImageType = "Tiff"

	ImageTypePdf ImageType = "Pdf"

	ImageTypeBmp ImageType = "Bmp"

	ImageTypeGif ImageType = "Gif"

	ImageTypeJpeg ImageType = "Jpeg"
)

type FlagType string

const (
	FlagTypeNone FlagType = "None"

	FlagTypeAssigned FlagType = "Assigned"

	FlagTypeApproved FlagType = "Approved"

	FlagTypeDeclined FlagType = "Declined"

	FlagTypeArchived FlagType = "Archived"

	FlagTypeDistributed FlagType = "Distributed"

	FlagTypePrepostedInJournal FlagType = "PrepostedInJournal"

	FlagTypePostedInJournal FlagType = "PostedInJournal"
)

type AttachmentLocation string

const (
	AttachmentLocationRetrieval AttachmentLocation = "Retrieval"

	AttachmentLocationScanning AttachmentLocation = "Scanning"

	AttachmentLocationJournal AttachmentLocation = "Journal"
)

type Create struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ Create"`

	Type_ *ImageType `xml:"type,omitempty" json:"type,omitempty"`
}

type CreateResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ CreateResponse"`

	CreateResult *ImageFile `xml:"CreateResult,omitempty" json:"CreateResult,omitempty"`
}

type AppendChunk struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ AppendChunk"`

	File *ImageFile `xml:"file,omitempty" json:"file,omitempty"`

	Buffer []byte `xml:"buffer,omitempty" json:"buffer,omitempty"`

	Offset int64 `xml:"offset,omitempty" json:"offset,omitempty"`
}

type AppendChunkResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ AppendChunkResponse"`
}

type AppendChunkByLength struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ AppendChunkByLength"`

	File *ImageFile `xml:"file,omitempty" json:"file,omitempty"`

	Buffer []byte `xml:"buffer,omitempty" json:"buffer,omitempty"`

	BufferLength int32 `xml:"bufferLength,omitempty" json:"bufferLength,omitempty"`

	Offset int64 `xml:"offset,omitempty" json:"offset,omitempty"`
}

type AppendChunkByLengthResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ AppendChunkByLengthResponse"`
}

type DownloadChunk struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ DownloadChunk"`

	File *ImageFile `xml:"file,omitempty" json:"file,omitempty"`

	Offset int64 `xml:"offset,omitempty" json:"offset,omitempty"`

	BufferSize int32 `xml:"bufferSize,omitempty" json:"bufferSize,omitempty"`
}

type DownloadChunkResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ DownloadChunkResponse"`

	DownloadChunkResult []byte `xml:"DownloadChunkResult,omitempty" json:"DownloadChunkResult,omitempty"`
}

type GetFileInfo struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetFileInfo"`

	Parameters *FileSearchParameters `xml:"parameters,omitempty" json:"parameters,omitempty"`
}

type GetFileInfoResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetFileInfoResponse"`

	GetFileInfoResult *ArrayOfImageFile `xml:"GetFileInfoResult,omitempty" json:"GetFileInfoResult,omitempty"`
}

type GetSize struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetSize"`

	File *ImageFile `xml:"file,omitempty" json:"file,omitempty"`
}

type GetSizeResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetSizeResponse"`

	GetSizeResult int64 `xml:"GetSizeResult,omitempty" json:"GetSizeResult,omitempty"`
}

type Save struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ Save"`

	File *ImageFile `xml:"file,omitempty" json:"file,omitempty"`

	Location *AttachmentLocation `xml:"location,omitempty" json:"location,omitempty"`
}

type SaveResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ SaveResponse"`
}

type GetChecksum struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetChecksum"`

	File *ImageFile `xml:"file,omitempty" json:"file,omitempty"`
}

type GetChecksumResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetChecksumResponse"`

	GetChecksumResult string `xml:"GetChecksumResult,omitempty" json:"GetChecksumResult,omitempty"`
}

type GetMaxRequestLength struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetMaxRequestLength"`
}

type GetMaxRequestLengthResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetMaxRequestLengthResponse"`

	GetMaxRequestLengthResult int64 `xml:"GetMaxRequestLengthResult,omitempty" json:"GetMaxRequestLengthResult,omitempty"`
}

type GetApproverList struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetApproverList"`
}

type GetApproverListResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetApproverListResponse"`

	GetApproverListResult *ArrayOfKeyValuePair `xml:"GetApproverListResult,omitempty" json:"GetApproverListResult,omitempty"`
}

type GetStampNo struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetStampNo"`
}

type GetStampNoResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetStampNoResponse"`

	GetStampNoResult int32 `xml:"GetStampNoResult,omitempty" json:"GetStampNoResult,omitempty"`
}

type GetSeriesStampNo struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetSeriesStampNo"`

	SeriesId *Guid `xml:"SeriesId,omitempty" json:"SeriesId,omitempty"`
}

type GetSeriesStampNoResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetSeriesStampNoResponse"`

	GetSeriesStampNoResult int32 `xml:"GetSeriesStampNoResult,omitempty" json:"GetSeriesStampNoResult,omitempty"`
}

type GetSeries struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetSeries"`
}

type GetSeriesResponse struct {
	XMLName xml.Name `xml:"http://24sevenoffice.com/webservices/economy/accounting/ GetSeriesResponse"`

	GetSeriesResult *ArrayOfStampSeries `xml:"GetSeriesResult,omitempty" json:"GetSeriesResult,omitempty"`
}

type ImageFile struct {
	Id int32 `xml:"Id,omitempty" json:"Id,omitempty"`

	Type *ImageType `xml:"Type,omitempty" json:"Type,omitempty"`

	StampNo int32 `xml:"StampNo,omitempty" json:"StampNo,omitempty"`

	StampMeta *ArrayOfKeyValuePair `xml:"StampMeta,omitempty" json:"StampMeta,omitempty"`

	FrameInfo *ArrayOfImageFrameInfo `xml:"FrameInfo,omitempty" json:"FrameInfo,omitempty"`

	ContactId *ArrayOfInt `xml:"ContactId,omitempty" json:"ContactId,omitempty"`
}

type ArrayOfKeyValuePair struct {
	KeyValuePair []*KeyValuePair `xml:"KeyValuePair,omitempty" json:"KeyValuePair,omitempty"`
}

type KeyValuePair struct {
	Key string `xml:"Key,omitempty" json:"Key,omitempty"`

	Value string `xml:"Value,omitempty" json:"Value,omitempty"`
}

type ArrayOfImageFrameInfo struct {
	ImageFrameInfo []*ImageFrameInfo `xml:"ImageFrameInfo,omitempty" json:"ImageFrameInfo,omitempty"`
}

type ImageFrameInfo struct {
	Id int32 `xml:"Id,omitempty" json:"Id,omitempty"`

	Uri string `xml:"Uri,omitempty" json:"Uri,omitempty"`

	StampNo int32 `xml:"StampNo,omitempty" json:"StampNo,omitempty"`

	MetaData *ArrayOfKeyValuePair `xml:"MetaData,omitempty" json:"MetaData,omitempty"`

	Status int32 `xml:"Status,omitempty" json:"Status,omitempty"`
}

type ArrayOfInt struct {
	Int []int32 `xml:"int,omitempty" json:"int,omitempty"`
}

type FileSearchParameters struct {
	StampNo *ArrayOfInt `xml:"StampNo,omitempty" json:"StampNo,omitempty"`

	FileId *ArrayOfInt `xml:"FileId,omitempty" json:"FileId,omitempty"`

	AttachmentRegisteredAfter soap.XSDDateTime `xml:"AttachmentRegisteredAfter,omitempty" json:"AttachmentRegisteredAfter,omitempty"`

	AttachmentChangedAfter soap.XSDDateTime `xml:"AttachmentChangedAfter,omitempty" json:"AttachmentChangedAfter,omitempty"`

	HasStampNo *bool `xml:"HasStampNo,omitempty" json:"HasStampNo,omitempty"`

	FileApproved *bool `xml:"FileApproved,omitempty" json:"FileApproved,omitempty"`

	AttachmentStatus *ArrayOfFlagType `xml:"AttachmentStatus,omitempty" json:"AttachmentStatus,omitempty"`
}

type ArrayOfFlagType struct {
	FlagType []*FlagType `xml:"FlagType,omitempty" json:"FlagType,omitempty"`
}

type ArrayOfImageFile struct {
	ImageFile []*ImageFile `xml:"ImageFile,omitempty" json:"ImageFile,omitempty"`
}

type ArrayOfStampSeries struct {
	StampSeries []*StampSeries `xml:"StampSeries,omitempty" json:"StampSeries,omitempty"`
}

type StampSeries struct {
	Id *Guid `xml:"Id,omitempty" json:"Id,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Start int32 `xml:"Start,omitempty" json:"Start,omitempty"`

	End int32 `xml:"End,omitempty" json:"End,omitempty"`
}

type Guid string

type AttachmentServiceSoap interface {

	/* Creates a new ghost file that may be uploaded with UploadChunk. This file will only be a valid image in the client after using the 'Save' method. */
	Create(request *Create) (*CreateResponse, error)

	CreateContext(ctx context.Context, request *Create) (*CreateResponse, error)

	AppendChunk(request *AppendChunk) (*AppendChunkResponse, error)

	AppendChunkContext(ctx context.Context, request *AppendChunk) (*AppendChunkResponse, error)

	AppendChunkByLength(request *AppendChunkByLength) (*AppendChunkByLengthResponse, error)

	AppendChunkByLengthContext(ctx context.Context, request *AppendChunkByLength) (*AppendChunkByLengthResponse, error)

	DownloadChunk(request *DownloadChunk) (*DownloadChunkResponse, error)

	DownloadChunkContext(ctx context.Context, request *DownloadChunk) (*DownloadChunkResponse, error)

	/* Gets a list of fileIds from the given parameters */
	GetFileInfo(request *GetFileInfo) (*GetFileInfoResponse, error)

	GetFileInfoContext(ctx context.Context, request *GetFileInfo) (*GetFileInfoResponse, error)

	/* Returns the size of the given file in bytes. */
	GetSize(request *GetSize) (*GetSizeResponse, error)

	GetSizeContext(ctx context.Context, request *GetSize) (*GetSizeResponse, error)

	/* Saving the ghost file making it a valid image file. */
	Save(request *Save) (*SaveResponse, error)

	SaveContext(ctx context.Context, request *Save) (*SaveResponse, error)

	/* Calculates the md5 hash of the content to the given file. This may be used as a checksum to compare local and server file to see if they are equal and complete. */
	GetChecksum(request *GetChecksum) (*GetChecksumResponse, error)

	GetChecksumContext(ctx context.Context, request *GetChecksum) (*GetChecksumResponse, error)

	/* The max size of chunk in kB (with overhead) the server will accept. Do not exceed this limit when using the methods AppendChunk or DownloadChunk. */
	GetMaxRequestLength(request *GetMaxRequestLength) (*GetMaxRequestLengthResponse, error)

	GetMaxRequestLengthContext(ctx context.Context, request *GetMaxRequestLength) (*GetMaxRequestLengthResponse, error)

	/* Get a list of approvers */
	GetApproverList(request *GetApproverList) (*GetApproverListResponse, error)

	GetApproverListContext(ctx context.Context, request *GetApproverList) (*GetApproverListResponse, error)

	/* Get next StampNo */
	GetStampNo(request *GetStampNo) (*GetStampNoResponse, error)

	GetStampNoContext(ctx context.Context, request *GetStampNo) (*GetStampNoResponse, error)

	/* Get next StampNo for a specified series */
	GetSeriesStampNo(request *GetSeriesStampNo) (*GetSeriesStampNoResponse, error)

	GetSeriesStampNoContext(ctx context.Context, request *GetSeriesStampNo) (*GetSeriesStampNoResponse, error)

	/* Get Stamp number series */
	GetSeries(request *GetSeries) (*GetSeriesResponse, error)

	GetSeriesContext(ctx context.Context, request *GetSeries) (*GetSeriesResponse, error)
}

type attachmentServiceSoap struct {
	client *soap.Client
}

func NewAttachmentServiceSoap(client *soap.Client) AttachmentServiceSoap {
	return &attachmentServiceSoap{
		client: client,
	}
}

func (service *attachmentServiceSoap) CreateContext(ctx context.Context, request *Create) (*CreateResponse, error) {
	response := new(CreateResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/Create", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *attachmentServiceSoap) Create(request *Create) (*CreateResponse, error) {
	return service.CreateContext(
		context.Background(),
		request,
	)
}

func (service *attachmentServiceSoap) AppendChunkContext(ctx context.Context, request *AppendChunk) (*AppendChunkResponse, error) {
	response := new(AppendChunkResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/AppendChunk", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *attachmentServiceSoap) AppendChunk(request *AppendChunk) (*AppendChunkResponse, error) {
	return service.AppendChunkContext(
		context.Background(),
		request,
	)
}

func (service *attachmentServiceSoap) AppendChunkByLengthContext(ctx context.Context, request *AppendChunkByLength) (*AppendChunkByLengthResponse, error) {
	response := new(AppendChunkByLengthResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/AppendChunkByLength", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *attachmentServiceSoap) AppendChunkByLength(request *AppendChunkByLength) (*AppendChunkByLengthResponse, error) {
	return service.AppendChunkByLengthContext(
		context.Background(),
		request,
	)
}

func (service *attachmentServiceSoap) DownloadChunkContext(ctx context.Context, request *DownloadChunk) (*DownloadChunkResponse, error) {
	response := new(DownloadChunkResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/DownloadChunk", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *attachmentServiceSoap) DownloadChunk(request *DownloadChunk) (*DownloadChunkResponse, error) {
	return service.DownloadChunkContext(
		context.Background(),
		request,
	)
}

func (service *attachmentServiceSoap) GetFileInfoContext(ctx context.Context, request *GetFileInfo) (*GetFileInfoResponse, error) {
	response := new(GetFileInfoResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetFileInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *attachmentServiceSoap) GetFileInfo(request *GetFileInfo) (*GetFileInfoResponse, error) {
	return service.GetFileInfoContext(
		context.Background(),
		request,
	)
}

func (service *attachmentServiceSoap) GetSizeContext(ctx context.Context, request *GetSize) (*GetSizeResponse, error) {
	response := new(GetSizeResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetSize", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *attachmentServiceSoap) GetSize(request *GetSize) (*GetSizeResponse, error) {
	return service.GetSizeContext(
		context.Background(),
		request,
	)
}

func (service *attachmentServiceSoap) SaveContext(ctx context.Context, request *Save) (*SaveResponse, error) {
	response := new(SaveResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/Save", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *attachmentServiceSoap) Save(request *Save) (*SaveResponse, error) {
	return service.SaveContext(
		context.Background(),
		request,
	)
}

func (service *attachmentServiceSoap) GetChecksumContext(ctx context.Context, request *GetChecksum) (*GetChecksumResponse, error) {
	response := new(GetChecksumResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetChecksum", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *attachmentServiceSoap) GetChecksum(request *GetChecksum) (*GetChecksumResponse, error) {
	return service.GetChecksumContext(
		context.Background(),
		request,
	)
}

func (service *attachmentServiceSoap) GetMaxRequestLengthContext(ctx context.Context, request *GetMaxRequestLength) (*GetMaxRequestLengthResponse, error) {
	response := new(GetMaxRequestLengthResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetMaxRequestLength", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *attachmentServiceSoap) GetMaxRequestLength(request *GetMaxRequestLength) (*GetMaxRequestLengthResponse, error) {
	return service.GetMaxRequestLengthContext(
		context.Background(),
		request,
	)
}

func (service *attachmentServiceSoap) GetApproverListContext(ctx context.Context, request *GetApproverList) (*GetApproverListResponse, error) {
	response := new(GetApproverListResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetApproverList", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *attachmentServiceSoap) GetApproverList(request *GetApproverList) (*GetApproverListResponse, error) {
	return service.GetApproverListContext(
		context.Background(),
		request,
	)
}

func (service *attachmentServiceSoap) GetStampNoContext(ctx context.Context, request *GetStampNo) (*GetStampNoResponse, error) {
	response := new(GetStampNoResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetStampNo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *attachmentServiceSoap) GetStampNo(request *GetStampNo) (*GetStampNoResponse, error) {
	return service.GetStampNoContext(
		context.Background(),
		request,
	)
}

func (service *attachmentServiceSoap) GetSeriesStampNoContext(ctx context.Context, request *GetSeriesStampNo) (*GetSeriesStampNoResponse, error) {
	response := new(GetSeriesStampNoResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetSeriesStampNo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *attachmentServiceSoap) GetSeriesStampNo(request *GetSeriesStampNo) (*GetSeriesStampNoResponse, error) {
	return service.GetSeriesStampNoContext(
		context.Background(),
		request,
	)
}

func (service *attachmentServiceSoap) GetSeriesContext(ctx context.Context, request *GetSeries) (*GetSeriesResponse, error) {
	response := new(GetSeriesResponse)
	err := service.client.CallContext(ctx, "http://24sevenoffice.com/webservices/economy/accounting/GetSeries", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *attachmentServiceSoap) GetSeries(request *GetSeries) (*GetSeriesResponse, error) {
	return service.GetSeriesContext(
		context.Background(),
		request,
	)
}
