# Objetivo

Atender aos objetivos do desafio multithreading disponibilizado pela Pós Go Expert, segue abaixo descrição do teste.

```
Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://brasilapi.com.br/api/cep/v1/01153000 + cep

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.
```

Com isso foi tomado a decisão de adicionar um cli passando o cep para executar essas chamadas e retornar a response pela API mais rápida.

Exemplo de retorno no terminal.

```json
API BrasilAPI
Response: {
    "cep":"01310930",
    "state":"SP",
    "city":"São Paulo",
    "neighborhood":"Bela Vista",
    "street":"Avenida Paulista, 2100",
    "service":"correios-alt"
}
```

## Instrução

Para executar o programa fast_api temos duas opções:

- Rode o comando

```bash
go run main.go cep {cep}
```

- Execute o build do programa Go e chame o mesmo

```bash
go build -o fast_api ./main.go
./fast_api cep {cep}
```
