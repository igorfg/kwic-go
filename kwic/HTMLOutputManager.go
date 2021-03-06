//+build html

package kwic

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

type OutputManager struct {
	lines []string
}

func (h *OutputManager) Format(winc []string) {
	const tableBegin = "<table>"
	const tableEnd = "</table>"
	const tbodyBegin = "<tbody>"
	const tbodyEnd = "</tbody>"
	const trBegin = "<tr>"
	const trEnd = "</tr>"
	const tdBegin = "<td>"
	const tdBeginRight = "<td align=\"right\">"
	const tdEnd = "</td>"
	const bBegin = "<b>"
	const bEnd = "</b>"

	fmt.Println("Inicio da Formatação")

	var cntPipe int
	var lineFormated string

	h.lines = append(h.lines, "<h1 style=\"text-align: center\">Key Word in Context</h1>")
	h.lines = append(h.lines, "<h3 style=\"text-align: center\">Autores: <br> Igor Figueira <br> Khalil Carsten </h3>")
	h.lines = append(h.lines, tableBegin)
	h.lines = append(h.lines, tbodyBegin)

	for _, s := range winc {
		cntPipe = 0
		lineFormated = trBegin + tdBeginRight
		for _, c := range s {
			if c == '|' && cntPipe == 0 {
				lineFormated += tdEnd + tdBegin + bBegin
				cntPipe++
			} else if c == '|' && cntPipe == 1 {
				lineFormated += bEnd
			} else {
				lineFormated += string(c)
			}
		}
		lineFormated += tdEnd + trEnd
		h.lines = append(h.lines, lineFormated)
	}

	h.lines = append(h.lines, tbodyEnd)
	h.lines = append(h.lines, tableEnd)

	fmt.Println("Fim da Formatação")
}

func (h *OutputManager) Exhibit() error {
	if len(h.lines) <= 0 {
		return errors.New("OutputManager esta vazio")
	}

	if _, err := os.Stat("./outputHTML"); os.IsNotExist(err) {
		err = os.MkdirAll("./outputHTML", 0755)
		if err != nil {
			panic(err)
		}
	}

	dirToCreate := "./outputHTML/output.html"

	file, err := os.Create(dirToCreate)
	if err != nil {
		return errors.New("Não foi possível criar o arquivo")
	}
	defer file.Close()

	fmt.Println("Arquivo Criado")

	for _, line := range h.lines {
		file.WriteString(line)
	}

	dir := "file://"
	pwd, err := os.Getwd()
	dir += pwd + dirToCreate[1:]
	fmt.Println(dir)

	openbrowser(dir)

	return nil
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
