package kwic

import "bufio"
import "os"
import "errors"

// FileBasedStorageManager : struct que herda da interface DataStorageManager
// para leitura de arquivos
type FileBasedStorageManager struct {
	lines []string
}

// Init : Inicializa a estrutura de linhas a partir da leitura de arquivos
func (f *FileBasedStorageManager) Init() error {
	var filePath string

	// fmt.Print("Enter the path to the input file: ")
	// fmt.Scan(&filePath)

	filePath = "../resources/papers.txt"

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
func (f *FileBasedStorageManager) Line(index int) string {
	return f.lines[index]
}

// Length : Retorna a quantidade de linhas lidas do arquivo
func (f *FileBasedStorageManager) Length() int {
	return len(f.lines)
}
