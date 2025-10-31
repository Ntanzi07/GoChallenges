# Desafios em GoLang
Esses desafios abordam alguns aspectivos positivos na programação como a lógica, clean code, design patterns, entre outros detalhes.
A maioria não salvou o que fez, mas caso queiram, só postar no tópico subsequente ao desafio.

> ## 1º Desafio
Um programa que fica perguntando números para o usuário e que calcula o fatorial desse número, ao fim do cálculo salva o resultado em um arquivo cujo nome contém o número id (com 3 dígitos) da operação e o valor digitado pelo usuário (001-12.txt).

O programa não pode congelar o cursor em nenhum momento. Ou seja, assim que um número for escolhido o programa deve pedir o próximo número logo em seguida. O programa encerra somente quando o usuário digitar -1. Todos os cálculos devem estar terminados antes de encerrar.

> ## 2º Desafio
# API Desafio - GoLang

## Criar uma API Get, utilizando Fiber ou Gorilla/Mux.

`/fib/n` deve receber um valor, e retornar o resultado para o usuário em um objeto JSON:

```JSON
{
    "done": true,
    "fib":{
        "input": 6,
        "result": 8,
        "duration": "3s"
    }
}
```

### A API deve ter um timeout de 500 ms, caso o cáculo leve mais que 500 ms, ela deve retornar:

```JSON
{
    "done": false
}
```

### O cálculo deve ser realizado em paralelo e armazenado em um `map`, quando for solicitado novamente o número, se já tiver sido calculado, olhar no `map` para retornar o valor.

> Também deve ter o retorno `/all` para a API, retornando um objeto JSON com todas as consultas feitas no programa, por exemplo:

```JSON
{
   "queries":[
      {
         "input":2,
         "result":1,
         "duration":"3s"
      },
      {
         "input":3,
         "result":2,
         "duration":"3s"
      },
      {
         "input":4,
         "result":3,
         "duration":"3s"
      }
   ]
}
```