//+build file

package kwic

import "fmt"
import "bufio"
import "os"
import "errors"

// DataStorageManager : struct que herda da interface DataStorageManager
// para leitura de arquivos
type DataStorageManager struct {
	lines []string
}

// Init : Inicializa a estrutura de linhas a partir da leitura de arquivos
func (f *DataStorageManager) Init() error {
	var filePath string

	fmt.Print("Enter the path to the input file: ")
	fmt.Scan(&filePath)

	file, err := os.Open(filePath)

	if err != nil {
		return errors.New("Não foi possível abrir o arquivo")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f.lines = append(f.lines, scanner.Text())
	}
	return nil
}

// Line : Retorna a linha baseada no indice index passado
func (f *DataStorageManager) Line(index int) string {
	return f.lines[index]
}

// Length : Retorna a quantidade de linhas lidas do arquivo
func (f *DataStorageManager) Length() int {
	return len(f.lines)
}
