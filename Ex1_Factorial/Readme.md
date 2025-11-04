> ## 1º Desafio
# Fatorial Concorrente - GoLang

## 🇺🇸 Description
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

## 🇧🇷 Descrição
Crie um programa em Go que:
- Fique **pedindo números ao usuário** continuamente.  
- Para cada número digitado, **calcule o fatorial** desse número.  
- Após o cálculo, **salve o resultado em um arquivo** no formato:

resultados/001-12.txt

yaml
Copiar código

Onde:
- `001` é o **ID da operação** (3 dígitos, com zeros à esquerda).  
- `12` é o **número digitado** pelo usuário.

---

### Regras
- O programa **não pode travar o cursor**. Assim que um número for informado, ele deve **pedir o próximo imediatamente**.  
- O programa **só termina** quando o usuário digitar `-1`.  
- Antes de encerrar, **todos os cálculos devem estar concluídos**.

---

### Dicas
- Utilize **goroutines** e **WaitGroups** para realizar cálculos em paralelo.  
- Armazene os resultados em arquivos usando o pacote `os`.  
- Organize o código com boas práticas de **concorrência e sincronização**.

---

## Exemplo de extrutura
```
/fatorial
├── main.go
├── resultados/
│ ├── 001-5.txt
│ ├── 002-7.txt
│ └── 003-10.txt
```