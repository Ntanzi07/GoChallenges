> ## 3º Desafio
# Corrida de Downloads - GoLang

## English

## Objective

Simulate a system that downloads files in parallel, each with different download times, and displays real-time progress.

## Description

Create a Go program that:

Has a list of files to “download” (simulated).
Example:
```
files := []string{"fileA.zip", "fileB.zip", "fileC.zip"}
```

For each file, the program should:

- Launch a goroutine that simulates the download (e.g., time.Sleep(time.Second * n)).

- Send progress updates (in %) through a channel.

- Notify when the download is complete.

- Meanwhile, the main function:

- Receives messages from the channels.

It should display something like:
```
[fileA.zip] 45%
[fileB.zip] 72%
[fileC.zip] 18%
```

When all downloads are finished, the program should print:
```
✅ All downloads completed!
```
## Rules

- Each “download” must run in a separate goroutine.

- Use a WaitGroup to ensure all downloads finish before exiting the program.

- Use select with time.After to simulate variable download speeds.

- Channels must be properly closed when finished.

##Extras (optional)

Add an animated progress bar in the terminal (using fmt.Printf("\r...")).

Allow configuring the number of simultaneous downloads (like a “thread limit”).

Save a log file (.txt) with all completed downloads.

## Example Output
```
[fileA.zip] 0%
[fileB.zip] 0%
[fileC.zip] 0%
[fileA.zip] 30%
[fileB.zip] 40%
[fileC.zip] 20%
[fileA.zip] 100% ✅
[fileB.zip] 100% ✅
[fileC.zip] 100% ✅
✅ All downloads completed!
```


## Portuguese 
## Objetivo

Simular um sistema que baixa arquivos em paralelo, cada um com tempos de download diferentes, e exibe o progresso em tempo real.

## Descrição

Crie um programa em Go que:

Possua uma lista de arquivos para “baixar” (simulados).
Exemplo:

```
files := []string{"fileA.zip", "fileB.zip", "fileC.zip"}
```

Para cada arquivo, o programa deve:

- Criar uma goroutine que simula o download (ex: time.Sleep(time.Second * n)).
- Enviar mensagens de progresso (em %) por um canal.
- Informar quando o download terminar.
- Enquanto isso, a função principal:
- Recebe as mensagens dos canais.

Mostra no terminal algo como:
```
[fileA.zip] 45%
[fileB.zip] 72%
[fileC.zip] 18%
```

Quando todos os downloads terminarem, o programa exibe:
```
✅ Todos os downloads concluídos!
```
## Regras

- Cada “download” deve rodar em uma goroutine separada.
- Use um WaitGroup para garantir que todos terminem antes de encerrar o programa.
- Use select com time.After` para simular variações de tempo.
- Os canais devem ser fechados corretamente ao final.

## Extras (opcionais)

Adicione uma barra de progresso animada no terminal (com fmt.Printf("\r...")).

Permita configurar a quantidade de downloads simultâneos (tipo um “limite de threads”).

Salve o log dos downloads finalizados em um arquivo .txt.

## Exemplo de saída
```
[fileA.zip] 0%
[fileB.zip] 0%
[fileC.zip] 0%
[fileA.zip] 30%
[fileB.zip] 40%
[fileC.zip] 20%
[fileA.zip] 100% ✅
[fileB.zip] 100% ✅
[fileC.zip] 100% ✅
✅ Todos os downloads concluídos!
```
