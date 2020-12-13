package front

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/servnet/controller"
	"github.com/massarakhsh/servnet/ruler"
)

type DataRule struct {
	ruler.DataRule
}

type DataRuler interface {
	ruler.DataRuler
}

func BuildRule(page *ruler.DataPage) *DataRule {
	rule := &DataRule{}
	rule.BindPage(page)
	return rule
}

func (rule *DataRule) Execute() lik.Seter {
	rule.SeekPageSize()
	rule.execute()
	if rule.IsNeedRedraw {
		rule.ShowRedraw()
	}
	return rule.GetAllResponse()
}

func (rule *DataRule) execute() {
	if rule.IsShift("front") {
		rule.execute()
	} else if _,ctrl := rule.SeekControl(rule.Shift()); ctrl != nil {
		ctrl.(controller.Controller).Execute(rule, rule.GetPath())
	}
}

