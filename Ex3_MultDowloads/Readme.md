> ## 3º Desafio
# Downloads race - GoLang

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