package interfaces

import (
	"domain"
)

type GcpHandler interface {
	ListBuckets(flowType domain.FlowType, projectId string) (buckets []domain.Bucket, err error)
	FileExists(fileName string) (fileExists bool, err error)
	DownloadFile(fileName string) (success bool, err error)
	UploadFile(fileName string) (success bool, err error)
}

type GcpRepo struct {
	gcpHandlers map[string]GcpHandler
	gcpHandler  GcpHandler
}

type SourceBucketRepo GcpRepo
type SinkBucketRepo GcpRepo


func NewSourceBucketRepo(gcpHandlers map[string]GcpHandler) *SourceBucketRepo {
	bucketRepo := new(SourceBucketRepo)
	bucketRepo.gcpHandlers = gcpHandlers
	bucketRepo.gcpHandler = gcpHandlers["SourceBucketRepo"]
	return bucketRepo
}

func (repo *SourceBucketRepo) List(projectId string) (buckets []domain.Bucket, err error) {
	return repo.gcpHandler.ListBuckets(domain.SourceFlow, projectId)
}

func (repo *SourceBucketRepo) FileExists(fileName string) (fileExists bool, err error) {
	return repo.gcpHandler.FileExists(fileName)
}

func (repo *SourceBucketRepo) DownloadFile(fileName string) (success bool, err error) {
	return repo.gcpHandler.DownloadFile(fileName)
}
// UploadFile is not operational for a source bucket
func (repo *SourceBucketRepo) UploadFile(fileName string) (success bool, err error) {
	return false, nil
}


func NewSinkBucketRepo(gcpHandlers map[string]GcpHandler) *SinkBucketRepo {
	bucketRepo := new(SinkBucketRepo)
	bucketRepo.gcpHandlers = gcpHandlers
	bucketRepo.gcpHandler = gcpHandlers["SinkBucketRepo"]
	return bucketRepo
}

func (repo *SinkBucketRepo) List(projectId string) (buckets []domain.Bucket, err error) {
	return repo.gcpHandler.ListBuckets(domain.SinkFlow, projectId)
}

func (repo *SinkBucketRepo) FileExists(fileName string) (fileExists bool, err error) {
	return repo.gcpHandler.FileExists(fileName)
}
// DownloadFile is not operational for a sink bucket
func (repo *SinkBucketRepo) DownloadFile(fileName string) (success bool, err error) {
	return false, nil
}

func (repo *SinkBucketRepo) UploadFile(fileName string) (success bool, err error) {
	return repo.gcpHandler.UploadFile(fileName)
}
// ListFileNamesToFetch is not operational for a sink bucket
func (repo *SinkBucketRepo) ListFileNamesToFetch(fileName string) (cloudFiles domain.CloudFiles, err error) {
	return cloudFiles, err
}


