package workflow

import (
	"os"
	"strconv"
)

// https://www.alfredapp.com/help/workflows/script-environment-variables/
var (
	AlfredVersion           string
	AlfredVersionBuild      string
	AlfredWorkflowBundledId string
	AlfredWorkflowCacheDir  string
	AlfredWorkflowDataDir   string
	AlfredWorkflowName      string
	AlfredWorkflowUid       string
	AlfredDebug             int
)

func init() {
	AlfredVersion = os.Getenv("alfred_version")
	AlfredVersionBuild = os.Getenv("alfred_version_build")
	AlfredWorkflowBundledId = os.Getenv("alfred_workflow_bundleid")
	AlfredWorkflowCacheDir = os.Getenv("alfred_workflow_cache")
	AlfredWorkflowDataDir = os.Getenv("alfred_workflow_data")
	AlfredWorkflowName = os.Getenv("alfred_workflow_name")
	AlfredWorkflowUid = os.Getenv("alfred_workflow_uid")
	AlfredDebug, _ = strconv.Atoi(os.Getenv("alfred_workflow_cache"))
}
