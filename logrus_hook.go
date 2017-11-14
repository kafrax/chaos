package chaos

import (
	"github.com/Sirupsen/logrus"
	"runtime"
	"strings"
	"path"
)

//line number hook
//https://github.com/sirupsen/logrus/issues/63
type LineNumberHook struct{}

func (hook LineNumberHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook LineNumberHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 3, 3)
	cnt := runtime.Callers(6, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !strings.Contains(name, "github.com/Sirupsen/logrus") {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data["file"] = path.Base(file)
			entry.Data["func"] = path.Base(name)
			entry.Data["line"] = line
			break
		}
	}
	return nil
}

type PersistenceHook struct{}

func (hook PersistenceHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook PersistenceHook) Fire(entry *logrus.Entry) error {
	return nil
}
