
> ## 2º Desafio
# API Desafio - GoLang

## Create a GET API using Fiber or Gorilla/Mux

`/fib/n` should receive a value and return the result to the user as a JSON object:

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

### The API must have a 500 ms timeout.
If the calculation takes longer than 500 ms, it should return:

```JSON
{
    "done": false
}
```

### The calculation must be performed in parallel and stored in a `map`.
When the same number is requested again, the API should check the `map` and return the cached value instead of recalculating it.

> The API must also provide a /all endpoint that returns a JSON object containing all the queries made during the program’s execution, for example:

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