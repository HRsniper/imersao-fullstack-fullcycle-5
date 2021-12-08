# Imersão fullstack && fullcycle 5 - Back-end das ordens de pagamento

## Descrição

Repositório do back-end das ordens de pagamento feito com Nest.js

**Importante**: A aplicação do Apache Kafka e Golang deve estar rodando primeiro.

## Rodar a aplicação

### Configurar /etc/hosts

A comunicação entre as aplicações se dá de forma direta através da rede da máquina.
Para isto é necessário configurar um endereços que todos os containers Docker consigam acessar.

Acrescente no seu /etc/hosts (para Windows o caminho é C:\Windows\system32\drivers\etc\hosts):

```
127.0.0.1 host.docker.internal
```

Em todos os sistemas operacionais é necessário abrir o programa para editar o _hosts_ como Administrator da máquina ou root.

Execute os comandos:

```bash
docker compose up
```

Acessar http://localhost:3000. Use os arquivos `accounts.http` e `orders.http` para testar a API.
