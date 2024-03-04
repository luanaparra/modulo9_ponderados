# Ponderada 2 - Módulo 9
## Testes de um simulador de dispositivos IoT
A partir do simulador de dispositivos IoT desenvolvido na ponderada 1, foi implementado testes automatizados para validar o simulador, utilizando conceitos de TDD.

Dessa maneira, os testes abordam:
1. O recebimento, garantia que os dados enviados pelo simulador são recebidos pelo broker;
2. Validação dos dados, garantia que os dados enviados pelo simulador cheguem sem alternações;
3. Confirmação da taxa de disparo, garantia que o simulador atende às especificações de taxa de disparo de mensagens dentro de uma margem de erro razoável.

## Como executar?
Para a execução dos módulos, é necessário clonar este repositório:

```
git clone https://github.com/luanaparra/modulo9_ponderados
```

Caso precise, instale o mosquitto com o arquivo de configuração e dependências do GO.

```
mosquitto -c mosquitto.conf
go mod tidy
```

Dessa maneira, para rodar o publisher é preciso realizar esses comandos:

```
cd publisher
go run publisher.go
```
```
{   
    "sensor": <nome_sensor>,
    "transmission_rate": <taxa_transmissao>, <!-- hertz -->
    "unit": <unidade>, <!--unidade do sensor -->
    "longitude": <longitude>,
    "latitude": <latitude>,
    "qos": <qos> <!-- qualidade para comunicação MQTT -->
}
```

Por outro lado, temos o subscriber:

```
cd subscriber
go run subscriber.go
```
### Testes unitários
Para rodar os testes, diferentes comandos são necessários:

**Publisher**
```
cd publisher
go test -v
```

**Subscriber**
```
cd subscriber
go run subscriber.go
```

Desse modo, os testes foram adicionados para a avaliação do recebimento das mensagens pelo broker usando o mecanismo de QoS, além da confirmação da taxa de disparo dos sensores. Ademais, os testes asseguram que as mensagens enviadas pelo publisher foram recebidas corretamente pelo subscriber. 

## Demonstração

O vídeo de demonstração se encontra na pasta \demonstração.

