package front

import (
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/servnet/controller"
	"github.com/massarakhsh/servnet/ruler"
)

func (rule *DataRule) ShowRedraw() {
	rule.StoreItem(rule.showMainGen())
}

func (rule *DataRule) showMainGen() likdom.Domer {
	div := likdom.BuildDivClassId("main_page", "page")
	if rule.GetLevel() == 0 {
		if ruler.RootCreator != nil {
			ruler.RootCreator(rule, 0, rule.GetPath())
		}
		//root.BuildRoot(rule, 0, rule.GetPath())
	}
	rule.ItPage.PathLast = rule.BuildStackPath()
	dat := div.BuildDivClass("main_data fill")
	if rule.GetLevel() > 0 {
		dat.AppendItem(rule.showControlGen(rule.ItPage.Stack[0].(controller.Controller)))
	}
	return div
}

func (rule *DataRule) showControlGen(ctrl controller.Controller) likdom.Domer {
	tbl := likdom.BuildTableClass("main_data")
	if menu := ctrl.ShowMenu(rule); menu != nil {
		tbl.BuildTrTdClass("main_data").AppendItem(menu)
	}
	tbl.BuildTrTdClass("main_space")
	dat := tbl.BuildTrTdClass("main_info")
	if lev,_ := rule.SeekControl(ctrl.GetIndex()); lev+1 < rule.GetLevel() {
		dat.AppendItem(rule.showControlGen(rule.ItPage.Stack[lev+1].(controller.Controller)))
	} else {
		dat.AppendItem(ctrl.ShowInfo(rule))
	}
	return tbl
}
