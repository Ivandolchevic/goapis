// Package configuration provides an access to all configuration parameters
package configurationUtil

import (
	"time"

	. "github.com/Ivandolchevic/goapis/pkg/models/util"
	configurationUtil "github.com/Ivandolchevic/goapis/pkg/utils/configurationUtil"
	fileUtil "github.com/Ivandolchevic/goapis/pkg/utils/fileUtil"
)

// GetLogFilePath returns the file path to the current log file
func GetLogFilePath() string {
	filepath := fileUtil.PathBuilder([]string{configurationUtil.Get().LogFolder, "log"})
	return filepath
}

func WriteError(e *APIError) {
	fileUtil.WriteStringToDisk(GetLogFilePath(),
		"{date:\""+time.Now().Format("2006-01-02 15:04:05")+"\", level:\"error\", error:\""+e.Error.Error()+"\",message: \""+e.Message+"\"}")
}

func WriteString(value string) {
	fileUtil.WriteStringToDisk(GetLogFilePath(),
		"{date:\""+time.Now().Format("2006-01-02 15:04:05")+"\", level:\"infos\",message: \""+value+"\"}")
}
