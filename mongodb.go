package chaos

import (
	"bytes"
	"os"
	"os/exec"
	"stream/fs"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/tomb.v2"
	"github.com/kafrax/logx"
	"fmt"
)

var mgoSession *mgo.Session

//5mode：
//primary Perform all read operations on the master node
//primaryPreferred Priority on the main node to read, if the primary node is not available, and then from the slave operation
//secondary All read operations are performed on the slave node
//secondaryPreferred Priority to read from the slave node, if all slave nodes are unavailable, and then from the master node operation。
//nearest According to the network delay time, the nearest read operation, regardless of the node type。
//default is strong mode its named primary. eg. mgoSession.SetMode(mgo.Strong)
//des:https://segmentfault.com/a/1190000000460489
var V_MGO_DIALINFO=&mgo.DialInfo{}
func newMgoSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.DialWithInfo(V_MGO_DIALINFO)
		if err != nil {
			logx.Info("newMgoSession |message=%v",err)
		}

		mgoSession.SetMode(mgo.PrimaryPreferred, false)
	}
	return mgoSession
}

func Execute(dbname, colName string, q func(*mgo.Collection) error) error {
	s := newMgoSession().Clone()
	defer s.Close()
	c := s.DB(dbname).C(colName)
	return q(c)
}

func ExecuteBulk(dbname, colName string, q func(*mgo.Bulk) error) error {
	s := newMgoSession().Clone()
	defer s.Close()
	bulk := s.DB(dbname).C(colName).Bulk()
	return q(bulk)
}

func DropDatabase(dbname string) error {
	defer func() {
		if mgoSession != nil {
			checkSessions()
			if mgoSession != nil {
				mgoSession.Close()
				mgoSession = nil
			}
		}
	}()
	return newMgoSession().DB(viper.GetString("databasename")).DropDatabase()
}

func DropCollection(dbname string, names ...string) error {
	defer func() {
		if mgoSession != nil {
			checkSessions()
			if mgoSession != nil {
				mgoSession.Close()
				mgoSession = nil
			}
		}
	}()
	for _, v := range names {
		return newMgoSession().DB(dbname).C(v).DropCollection()
	}
	return nil
}

func AllCollection(dbname string) []string {
	ret, err := newMgoSession().DB(dbname).CollectionNames()
	if err != nil {
		return nil
	}

	return ret
}

type MongoImport struct {
	output      bytes.Buffer
	server      *exec.Cmd
	dbname      string
	dir         string
	currentFile string
	tomb        tomb.Tomb
}

var mongodImport *MongoImport
//regPath
//eg. /readingmate_user_action
func NewMongoImport(dbname, regPath string) *MongoImport {
	if mongodImport == nil {
		mongodImport = new(MongoImport)
		mongodImport.dbname = dbname
		mongodImport.dir = viper.GetString("filedir") + regPath
	}
	return mongodImport
}

// ls -1 *.log | while read jsonfile; do mongoimport -d test -c tmp --file $jsonfile ; done
func (m *MongoImport) MongoImportAllPre() {
	os.Remove(m.dir + "/all.json")
	m.tomb = tomb.Tomb{}
	m.server = exec.Command("cat", "*.json ", ">> all.json")
	m.server.Dir = m.dir
	m.server.Stdout = &m.output
	m.server.Stderr = &m.output
	m.server.Start()
	m.tomb.Go(m.monitor)
}

//MongoImport1File
func (m *MongoImport) MongoImport1File(col string, file string) {
	args := []string{
		"-d", m.dbname,
		"-c", col,
		"--file", file,
		"--host", viper.GetString("mgohost"),
	}
	m.server = exec.Command("mongoimport", args...)
	m.server.Dir = m.dir
	m.server.Stdout = &m.output
	m.currentFile = file
	m.server.Stderr = &m.output
	m.server.Run()
	//m.tomb.Go(m.monitor)
}

//MongoImportAllFile
func (m *MongoImport) MongoImportAll(col string, files ...string) {
	for _, v := range files {
		file := v
		args := []string{
			"-d", m.dbname,
			"-c", col,
			"--file", file,
			"--host", viper.GetString("mgohost"),
		}
		m.server = exec.Command("mongoimport", args...)
		m.server.Dir = m.dir
		m.currentFile = file
		m.server.Stdout = &m.output
		m.server.Stderr = &m.output
		m.server.Start()
		m.tomb.Go(m.monitor)
	}
}

func (m *MongoImport) MongoImportAllByPath(col, sub string) {
	files := fs.ReadDirNoLink(m.dir)
	if len(files) < 1 {
		logx.Debug("MongoImportAllByPath |message=no dir")
	}
	for _, v := range files {
		if !strings.Contains(v, sub) {
			continue
		}
		file := v
		args := []string{
			"-d", m.dbname,
			"-c", col,
			"--file", m.dir + "/" + file,
			"--host", viper.GetString("mgohost"),
		}

		m.server = exec.Command("mongoimport", args...)
		m.server.Dir = m.dir
		m.server.Stdout = &m.output
		m.currentFile = file
		m.server.Stderr = &m.output
		m.server.Run()
		logx.Debugf("mongoimport |message=%v |file=%v", "done", m.currentFile)
		//m.tomb.Go(m.monitor)
	}
	logx.Debug("MongoImportAllByPath |导入所有文件到mongodb完毕")
}

func (m *MongoImport) monitor() error {
	m.server.Process.Wait()
	if m.tomb.Alive() {
		cmd := exec.Command("/bin/sh", "-c", "ps auxw | grep mongoimport")
		cmd.Stdout = os.Stderr
		cmd.Stderr = os.Stderr
		b, _ := cmd.Output()
		fmt.Println(string(b))
		logx.Debugf("mongoimport |message=%v |file=%v", "done", m.currentFile)
	}
	return nil
}

func (m *MongoImport) Stop() {
	if mgoSession != nil {
		checkSessions()
		if mgoSession != nil {
			mgoSession.Close()
			mgoSession = nil
		}
	}
	if m.server != nil {
		m.tomb.Kill(nil)
		m.server.Process.Signal(os.Interrupt)
		select {
		case <-m.tomb.Dead():
		case <-time.After(5 * time.Second):
			panic("timeout waiting for mongod process to die")
		}
		m.server = nil
	}
}

func checkSessions() {
	if mgoSession == nil {
		return
	}
	mgoSession.Close()
	mgoSession = nil
}