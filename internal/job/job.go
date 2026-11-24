package job

type Status string
type Stage string

const (
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
)

const (
	StageIngest    Stage = "ingest"
	StageTransform Stage = "transform"
	StagePublish   Stage = "publish"
)

type Job struct {
	ID             string `json:"id"`             // 작업 식별
	Stage          Stage  `json:"stage"`          // 공정 구분
	Status         Status `json:"status"`         // 처리 상태
	RetryCount     int    `json:"retryCount"`     // 재시도 횟수
	ReleaseVersion string `json:"releaseVersion"` // 버전 식별
}

func SampleJobs(version string) []Job {
	return []Job{
		{
			ID:             "job-001",
			Stage:          StageIngest,
			Status:         StatusCompleted,
			RetryCount:     0,
			ReleaseVersion: version,
		},
		{
			ID:             "job-002",
			Stage:          StageTransform,
			Status:         StatusProcessing,
			RetryCount:     0,
			ReleaseVersion: version,
		},
		{
			ID:             "job-003",
			Stage:          StagePublish,
			Status:         StatusFailed,
			RetryCount:     1,
			ReleaseVersion: version,
		},
	}
}
