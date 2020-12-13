package diagnostic

import (
	"fmt"
	"github.com/massarakhsh/lik/likapi"
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/servnet/controller"
	"github.com/massarakhsh/servnet/ruler"
)

type DiagnosticControl struct {
	controller.DataControl
}

type Diagnosticer interface {
	controller.Controller
}

func BuildDiagnostic(rule ruler.DataRuler, level int, path []string) Diagnosticer {
	it := &DiagnosticControl{}
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
	it.ItExecute(rule, path)
	return it
}

func (it *DiagnosticControl) ShowMenu(rule ruler.DataRuler) likdom.Domer {
	if it.Mode == "" {
		it.Mode = "summary"
	}
	tbl := likdom.BuildTableClass("menu")
	row := tbl.BuildTr()
	it.MenuItemText(rule, row, "Диагностика")
	it.MenuItemText(rule, row, "|")
	it.MenuItemCmd(rule, row, "summary", "Сводка", "summary")
	it.MenuItemCmd(rule, row, "resource", "Ресурсы", "resource")
	it.MenuTools(rule, row)
	return tbl
}

func (it *DiagnosticControl) ShowInfo(rule ruler.DataRuler) likdom.Domer {
	div := likdom.BuildDivClass("grid")
	div.AppendItem(it.buildDiagnFront(rule))
	return div
}

func (it *DiagnosticControl) ItExecute(rule ruler.DataRuler, path []string) {
	if cmd := it.PopCommand(&path); cmd == "" {
	} else if cmd == "purgebase" {
		fmt.Println("purgebase")
	} else if cmd == "summary" {
		it.Mode = "summary"
		rule.SetNeedRedraw()
	} else if cmd == "resource" {
		it.Mode = "resource"
		rule.SetNeedRedraw()
	} else {
		it.ExecuteController(rule, cmd)
	}
}

func (it *DiagnosticControl) buildProc(part string) string {
	path := it.BuildPart(part)
	return fmt.Sprintf("%s('%s')", "cmd_diagnostic", path)
}

func (it *DiagnosticControl) buildDiagnFront(rule ruler.DataRuler) likdom.Domer {
	tbl := likdom.BuildTableClass("")
	row := tbl.BuildTr()
	if td := row.BuildTdClass("column"); td != nil {
		clm := td.BuildTableClass("")
		clm.BuildTrTd().AppendItem(it.buildDiagnInterface(rule))
		clm.BuildTrTd().AppendItem(it.buildDiagnServer(rule))
		clm.BuildTrTd().AppendItem(it.buildDiagnInterface(rule))
	}
	if td := row.BuildTdClass("column"); td != nil {
		clm := td.BuildTableClass("")
		clm.BuildTrTd().AppendItem(it.buildDiagnServer(rule))
		clm.BuildTrTd().AppendItem(it.buildDiagnInterface(rule))
	}
	return tbl
}

func (it *DiagnosticControl) buildDiagnInterface(rule ruler.DataRuler) likdom.Domer {
	tbl := likdom.BuildTableClass("")
	it.buildAppendRow(tbl, "Концентратор", "Ok")
	it.buildAppendRow(tbl, "Концентратор", "Ok")
	it.buildAppendRow(tbl, "Маршрутизатор", "Не отвечает")
	return it.HeadTableString("Интерфейсы", tbl)
}

func (it *DiagnosticControl) buildDiagnServer(rule ruler.DataRuler) likdom.Domer {
	tbl := likdom.BuildTableClass("")
	it.buildAppendRow(tbl, "Сервер", "Ok")
	it.buildAppendRow(tbl, "Сервер", "Не отвечает")
	return it.HeadTableString("Серверы", tbl)
}

func (it *DiagnosticControl) buildAppendRow(tbl likdom.Domer, title string, diagn string) {
	row := tbl.BuildTr()
	row.BuildTdClass("panelinfo").BuildString(title)
	row.BuildTdClass("panelinfo").BuildString(diagn)
}
