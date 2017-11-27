package octoprint

import (
	"encoding/json"
)

const JobTool = "/api/job"

// JobCommand retrieve information about the current job (if there is one).
type JobCommand struct{}

// JobResponse is the response from a job command.
type JobResponse struct {
	// Job contains information regarding the target of the current print job.
	Job JobInformation `json:"job"`
	// Progress contains information regarding the progress of the current job.
	Progress ProgressInformation `json:"progress"`
}

// Do sends an API request and returns the API response.
func (cmd *JobCommand) Do(c *Client) (*JobResponse, error) {
	b, err := c.doRequest("GET", JobTool, nil)
	if err != nil {
		return nil, err
	}

	r := &JobResponse{}
	if err := json.Unmarshal(b, r); err != nil {
		return nil, err
	}

	return r, err
}
