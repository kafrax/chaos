package chaos

import (
	"os"
	"fmt"
	"path/filepath"

	"github.com/takama/daemon"
)

/*
eg.
./<appname> install  
./<appname> remove 
./<appname> start 
./<appname> stop  
./<appname> status 
*/
func StartDaemon() {
	os.Chdir(filepath.Dir(os.Args[0]))
	if len(os.Args) > 1 {
		proc := filepath.Base(os.Args[0])
		svc, err := daemon.New(proc, proc, proc+".service")
		if err == nil {
			var msg string
			switch os.Args[1] {
			case "install":
				msg, err = svc.Install()
			case "remove":
				msg, err = svc.Remove()
			case "start":
				msg, err = svc.Start()
			case "stop":
				msg, err = svc.Stop()
			case "status":
				msg, err = svc.Status()
			default:
				msg = "Usage: " + proc +
					" install | remove | start | stop | status"
			}
			fmt.Println(msg)
		}
		if err != nil {
			fmt.Println("Error:", err)
		}
		os.Exit(1)
	}
}
