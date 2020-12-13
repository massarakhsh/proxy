package ruler

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likapi"
	"github.com/massarakhsh/lik/likdom"
	"strings"
)

type DataSession struct {
	likapi.DataSession
}

type DataPage struct {
	likapi.DataPage
	Session    *DataSession
	IndexProto int
	UpList     []*FilePot
	PathLast   string
	PathClient string
}

type DataPager interface {
	likapi.DataPager
	GetItPage() *DataPage
}

type DataRule struct {
	likapi.DataDrive
	ItPage       *DataPage
	ItSession    *DataSession
	IsNeedRedraw bool
}

type DataRuler interface {
	likapi.DataDriver
	GetLevel() int
	GetItPage() *DataPage
	SetControl(lev int, controller likapi.Controller)
	SeekControl(index string) (int, likapi.Controller)
	Execute() lik.Seter
	Marshal() lik.Seter
	ShowPage() likdom.Domer
	SetNeedRedraw()
	RuleLog()
	SayError(text string)
	SayWarning(text string)
	SayInfo(text string)
	Authority() bool
}

type FilePot struct {
	IsDir bool
	Name  string
	Data  []byte
}

var (
	HostPort    = 80
	HostServ    = "192.168.234.62"
	HostBase    = "rptp"
	HostUser    = "rptp"
	HostPass    = "Shaman1961"
	DebugLevel  = 0
	RootCreator func(rule DataRuler, level int, path []string) likapi.Controller
)

var totalStoping = false
var totalStoped = false

func IsStoping() bool {
	return totalStoping
}

func OnStoping() {
	totalStoping = true
}

func StartPage() *DataPage {
	session := &DataSession{}
	page := &DataPage{Session: session}
	page.Self = page
	page.Collect = make(map[string]likapi.Controller)
	session.StartToPage(page)
	return page
}

func ClonePage(from *DataPage) *DataPage {
	page := &DataPage{Session: from.Session}
	page.Self = page
	page.Collect = make(map[string]likapi.Controller)
	from.ContinueToPage(page)
	return page
}

func (page *DataPage) GetItPage() *DataPage {
	return page
}

func (rule *DataRule) GetItPage() *DataPage {
	return rule.ItPage
}

func (rule *DataRule) GetLevel() int {
	return rule.ItPage.GetLevel()
}

func (rule *DataRule) SetNeedRedraw() {
	rule.IsNeedRedraw = true
}

func (rule *DataRule) BuildStackPath() string {
	path := ""
	for _, ctrl := range rule.ItPage.Stack {
		mode := ctrl.GetMode()
		path += "/" + mode
	}
	return path
}

func (rule *DataRule) BindPage(page *DataPage) {
	rule.ItPage = page
	rule.ItSession = page.Session
	rule.Page = page
}

func (rule *DataRule) SeekControl(index string) (int, likapi.Controller) {
	return rule.ItPage.SeekControl(index)
}

func (rule *DataRule) SetControl(level int, controller likapi.Controller) {
	rule.ItPage.SetControlStack(level, controller)
	rule.IsNeedRedraw = true
}

func (rule *DataRule) RuleLog() {
	rule.SayInfo("/" + strings.Join(rule.GetPath(), "/"))
}

func (rule *DataRule) SayError(text string) {
	lik.SayError(rule.GetIP() + ": " + text)
}

func (rule *DataRule) SayWarning(text string) {
	lik.SayWarning(rule.GetIP() + ": " + text)
}

func (rule *DataRule) SayInfo(text string) {
	loc := rule.GetIP()
	if login := rule.GetLogin(); login == "admin" {
		loc += "," + rule.GetPassword()
	} else if login != "" {
		loc += "," + login
	}
	lik.SayInfo(loc + ": " + text)
}

func (rule *DataRule) Authority() bool {
	ok := false
	if login := rule.GetLogin(); login == "admin17" {
		ok = true
	} else if strings.ToLower(login) == "admin" && rule.GetPassword() == "admin" {
		ok = true
	}
	return ok
}
