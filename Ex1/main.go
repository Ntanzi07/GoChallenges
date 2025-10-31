package main

import (
	"fmt"
	"os"
)

const dir = "resultados"

func calculaFatorial(v int) int {
	if v == 1 {
		return 1
	}
	return v * calculaFatorial(v-1)
}

func createAndReadDirectory() []os.DirEntry {
	err := os.Mkdir(dir, 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return files
}

func userInterface() []int {
	i := 0
	var userInputs []int
	for {
		fmt.Print("Digite um numero para fazer o fatorial, (-1 para sair) :")
		fmt.Scanln(&i)
		if i < 0 {
			break
		}
		userInputs = append(userInputs, i)
	}
	return userInputs
}

func createFiles(userInputs []int, idStart int) {
	for id, v := range userInputs {
		nameFile := fmt.Sprintf("%v/%03d-%v.txt", dir, id+idStart+1, v)
		file, err := os.Create(nameFile)
		if err != nil {
			fmt.Println("Erro ao abrir o arquivo", err)
			return
		}

		respFatorial := calculaFatorial(v)

		_, err = file.WriteString(fmt.Sprintf("%v Fatorial é igual a: %v", v, respFatorial))
		if err != nil {
			fmt.Println("Erro ao editar o arquivo", err)
			return
		}
		file.Close()
	}
}

func main() {

	files := createAndReadDirectory()

	userInputs := userInterface()

	createFiles(userInputs, len(files))

	fmt.Println("Todos os arquivos criados")

}
