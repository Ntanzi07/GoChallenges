package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var wg sync.WaitGroup

type fileToDownload struct {
	name     string
	progress int
}

func main() {
	c := make(chan fileToDownload)
	files := []string{"fileA.zip", "fileB.zip", "fileC.zip"}

	writer(c, files)
	go receive(c, files)

	wg.Wait()
	close(c)
	fmt.Println("\n✅ Todos os downloads concluídos!")
}

func writer(cw chan<- fileToDownload, files []string) {
	for _, v := range files {
		wg.Add(1)
		go downloadFile(cw, v)
	}
}

func downloadFile(cw chan<- fileToDownload, file string) {
	for i := 0; i <= 100; i++ {
		cw <- fileToDownload{name: file, progress: i}
		time.Sleep(time.Millisecond * time.Duration(rand.IntN(1e3)))
	}
	wg.Done()
}

func receive(cr <-chan fileToDownload, files []string) {
	progressMap := make(map[string]int)

	// imprime linhas iniciais (para reservar espaço no terminal)
	for range files {
		fmt.Println()
	}

	for update := range cr {
		progressMap[update.name] = update.progress
		printAllBars(progressMap, files)
	}
}

func printAllBars(progressMap map[string]int, files []string) {
	// move o cursor pro topo das barras
	fmt.Printf("\033[%dA", len(files))

	for _, name := range files {
		progress := progressMap[name]
		printProgressBar(name, progress)
	}
}

func printProgressBar(name string, progress int) {
	const barWidth = 30
	filled := int(float64(progress) / 100 * barWidth)

	bar := ""
	for i := 0; i < filled; i++ {
		bar += "█"
	}
	for i := filled; i < barWidth; i++ {
		bar += " "
	}

	fmt.Printf("[%s] [%s] %3d%%\n", name, bar, progress)
}
