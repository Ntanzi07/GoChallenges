> ## 1º Desafio
# Fatorial Concorrente - GoLang

## Description
Create a Go program that:
- Continuously **asks the user for numbers**.  
- For each entered number, **calculates its factorial**.  
- After the calculation, **saves the result into a file** named as follows:

results/001-12.txt

Where:
- `001` is the **operation ID** (3 digits, zero-padded).  
- `12` is the **user-entered number**.

---

### Rules
- The program **must not freeze the cursor** — it should immediately prompt for the next number after one is entered.  
- The program **ends only** when the user enters `-1`.  
- Before terminating, **all ongoing calculations must be completed**.

---

### Tips
- Use **goroutines** and **WaitGroups** to perform calculations concurrently.  
- Save results using Go’s `os` package.  
- Apply **clean code and concurrency best practices**.

---

## Example File Structure
```
/fatorial
├── main.go
├── resultados/
│ ├── 001-5.txt
│ ├── 002-7.txt
│ └── 003-10.txt
```