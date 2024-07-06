# Micro-serviços simulando um sistema de vendas

## Descrição
Este projeto é um simulador de um sistema de vendas, onde é possível realizar a compra de produtos, fazer checkout, finalizar a compra e consultar o status da compra, feito em micro-serviços, utilizando GO, RabbitMQ e PostgreSQL, com o objetivo de estudar e aprender mais sobre micro-serviços.

## Tecnologias
- GO
    - Gin Gonic
- RabbitMQ
- PostgreSQL

## Estrutura
- **products-go**: Micro-serviço responsável por gerenciar os produtos.
- **checkout-go**: Micro-serviço responsável por gerenciar o carrinho de compras.
- **order-go**: Micro-serviço responsável por gerenciar as ordens de compra.
- **payment-go**: Micro-serviço responsável por gerenciar os pagamentos.

## Como rodar
1. Clone o repositório
2. Entre na pasta do projeto
3. Execute o comando `docker-compose up --build`

caso queira rodar os micro-serviços separadamente, entre na pasta de cada micro-serviço e execute o comando `go run main.go`
ou se preferir, pode buildar cada imagem (docker build -t <nome> <dockerfile>) e rodar separadamente.

## Documentação
Detalhes de cada micro-serviço e suas rotas estão disponíveis na raiz em: `docs.yaml`,
feita com o [Swagger](https://swagger.io/)

## Autor
- [Jonas Neto](https://github.com/jonasssneto)
