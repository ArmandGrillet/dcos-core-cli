package metronome

import (
	"time"
)

// Job represents a Job returned by the Metronome API.
type Job struct {
	ID          string            `json:"id"`
	Description string            `json:"description"`
	Labels      map[string]string `json:"labels"`
	// The run property of a Job represents the run configuration for that Job
	Run         struct {
		Args           []string          `json:"args"`
		Artifacts      []artifact        `json:"artifacts"`
		Cmd            string            `json:"cmd"`
		Cpus           float32           `json:"cpus"`
		Disk           int               `json:"disk"`
		Docker         docker            `json:"docker"`
		Env            map[string]string `json:"env"`
		MaxLaunchDelay int               `json:"maxLaunchDelay"`
		Mem            int               `json:"mem"`
		Placement      placement         `json:"placement"`
		User           string            `json:"user,omitempty"`
		Restart        restart           `json:"restart"`
		Volumes        []volume          `json:"volumes"`
	} `json:"run"`

	// These properties depend on the embed parameters when querying the /v1/jobs endpoints.
	ActiveRuns     []Run              `json:"activeRuns,omitempty"`
	HistorySummary *JobHistorySummary `json:"historySummary,omitempty"`
	History        *JobHistory        `json:"history,omitempty"`
	Schedules      []Schedule         `json:"schedules,omitempty"`
}

type artifact struct {
	URI        string `json:"uri"`
	Executable bool   `json:"executable"`
	Extract    bool   `json:"extract"`
	Cache      bool   `json:"cache"`
}

type docker struct {
	Image string `json:"image"`
}

type placement struct {
	Constraints []constraint `json:"constraints"`
}

type constraint struct {
	Attribute string `json:"attribute"`
	Operator  string `json:"operator"`
	Value     string `json:"value"`
}

type restart struct {
	Policy                string `json:"policy"`
	ActiveDeadlineSeconds int    `json:"activeDeadlineSeconds"`
}

type volume struct {
	ContainerPath string `json:"containerPath"`
	HostPath      string `json:"hostPath"`
	Mode          string `json:"mode"`
}

// Run contains information about a run of a Job.
type Run struct {
	ID          string       `json:"id"`
	JobID       string       `json:"jobId"`
	Status      string       `json:"status"`
	CreatedAt   string       `json:"createdAt"`
	CompletedAt string       `json:"completedAt"`
	Tasks       []activeTask `json:"tasks"`
}

type activeTask struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	StartedAt   string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
}

// JobHistory contains statistics and information about past runs of a Job.
type JobHistory struct {
	SuccessCount           int          `json:"successCount"`
	FailureCount           int          `json:"failureCount"`
	LastSuccessAt          string       `json:"lastSuccessAt"`
	LastFailureAt          string       `json:"lastFailureAt"`
	SuccessfulFinishedRuns []runHistory `json:"successfulFinishedRuns"`
	FailedRuns             []runHistory `json:"failedFinishedRuns"`
}

type runHistory struct {
	ID         string `json:"id"`
	CreatedAt  string `json:"createdAt"`
	FinishedAt string `json:"finishedAt"`
}

// JobHistorySummary contains statistics about past runs of a Job.
type JobHistorySummary struct {
	SuccessCount  int    `json:"successCount"`
	FailureCount  int    `json:"failureCount"`
	LastSuccessAt string `json:"lastSuccessAt"`
	LastFailureAt string `json:"lastFailureAt"`
}

// Schedule of a Job.
type Schedule struct {
	ID                      string `json:"id"`
	Cron                    string `json:"cron"`
	TimeZone                string `json:"timeZone"`
	StartingDeadlineSeconds int    `json:"startingDeadlineSeconds"`
	ConcurrencyPolicy       string `json:"concurrencyPolicy"`
	Enabled                 bool   `json:"enabled"`
	NextRunAt               string `json:"nextRunAt"`
}

// Status returns the status of the job depending on its active runs and its schedule.
func (j *Job) Status() string {
	switch {
	case j.ActiveRuns != nil:
		return "Running"
	case len(j.Schedules) == 0:
		return "Unscheduled"
	default:
		return "Scheduled"
	}
}

// LastRunStatus returns the status of the last run of this job.
func (j *Job) LastRunStatus() (string, error) {
	if j.HistorySummary.LastSuccessAt == "" && j.HistorySummary.LastFailureAt == "" {
		return "N/A", nil
	} else if j.HistorySummary.LastFailureAt == "" {
		return "Success", nil
	} else if j.HistorySummary.LastSuccessAt == "" {
		return "Failure", nil
	}

	lastSuccess, err := time.Parse(time.RFC3339, j.HistorySummary.LastSuccessAt)
	if err != nil {
		return "", err
	}

	lastFailure, err := time.Parse(time.RFC3339, j.HistorySummary.LastFailureAt)
	if err != nil {
		return "", err
	}

	if lastSuccess.After(lastFailure) {
		return "Success", nil
	}
	return "Failure", nil

}
