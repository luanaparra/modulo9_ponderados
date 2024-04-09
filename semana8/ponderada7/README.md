# Ponderada 8 - Módulo 9
## Integração entre o Kafka cloud e o HiveMQ e consumer do Kafka

O exercício consiste em fazer com que o subscriber passe a consumir mensagens vindas do Kafka e as armazene em um banco de dados Mongo, além de ter uma integração com o Hive MQ

### Como executar? 

Utilizamos o MongoDB em container, para seu lançamento, sendo necessário rodar o comando:

```
docker compose up -d
```

Dessa maneira, para rodar o publisher, em seus respectivos diretórios rode:

```
cd publisher
go run .
```

Ademais, para rodar o código kafka:

```
cd pkg/kafka
go run .
```

Por fim, no `kafka.go` cria um consumer Kafka, que se conecta a um cluster Kafka utilizando autenticação SASL_SSL. Sendo que o passo a passo é:

1. Carregamento de variáveis de ambiente;
2. Configuração do consumer Kafka;
3. Criação do consumer Kafka;
4. Assinatura de tópicos;
5. Loop de consumo de mensagens;
6. Fechamento do consumer.


## Demonstração

O vídeo de demonstração se encontra na pasta \demonstração.