package root

import (
	"github.com/massarakhsh/lik/likapi"
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/servnet/controller"
	"github.com/massarakhsh/servnet/controller/diagnostic"
	"github.com/massarakhsh/servnet/controller/gate"
	"github.com/massarakhsh/servnet/controller/unit"
	"github.com/massarakhsh/servnet/ruler"
)

type RootControl struct {
	controller.DataControl
}

type Rooter interface {
	controller.Controller
}

func BuildRoot(rule ruler.DataRuler, level int, path []string) likapi.Controller {
	it := &RootControl{}
	it.IfExecute = func (drive likapi.DataDriver, path []string) {
		it.ItExecute(drive.(ruler.DataRuler), path)
	}
	it.IfShow = func (drive likapi.DataDriver, style string) likdom.Domer {
		if style == "menu" {
			return it.ShowMenu(drive.(ruler.DataRuler))
		} else {
			return it.ShowInfo(drive.(ruler.DataRuler))
		}
	}
	rule.SetControl(level, it)
	it.Execute(rule, path)
	return it
}

func (it *RootControl) ShowMenu(rule ruler.DataRuler) likdom.Domer {
	tbl := it.MenuPrepare(rule, false)
	row := tbl.BuildTr()
	it.MenuItemCmd(rule, row, "", "АО РПТП", "seek")
	//it.MenuItemText(rule, row, base.Version)
	it.MenuItemText(rule, row, "|")
	it.MenuItemCmd(rule, row, "unit", "Устройства", "unit")
	it.MenuItemCmd(rule, row, "gate", "Шлюз", "gate")
	it.MenuItemCmd(rule, row, "diagnostic", "Диагностика", "diagnostic")
	it.MenuItemCmd(rule, row, "file", "Файлы", "file")
	it.MenuTools(rule, row)
	return tbl
}

func (it *RootControl) ShowInfo(rule ruler.DataRuler) likdom.Domer {
	div := likdom.BuildDivClass("grid")
	return div
}

func (it *RootControl) ItExecute(rule ruler.DataRuler, path []string) {
	if cmd := it.PopCommand(&path); cmd == "" {
	} else if cmd == "unit" {
		it.Mode = "unit"
		unit.BuildUnit(rule, it.GetLevel(rule)+1, path)
	} else if cmd == "file" {
		it.Mode = "file"
		controller.BuildFile(rule, it.GetLevel(rule)+1, path)
	} else if cmd == "gate" {
		it.Mode = "gate"
		gate.BuildGate(rule, it.GetLevel(rule)+1, path)
	} else if cmd == "diagnostic" {
		it.Mode = "diagnostic"
		diagnostic.BuildDiagnostic(rule, it.GetLevel(rule)+1, path)
	} else {
		it.ExecuteController(rule, cmd)
	}
}

