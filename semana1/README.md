# Ponderada 1 - Módulo 9
## Simulador de dispositivos IoT
Criação de um simulador de dispositivos IoT, utilizando o protocolo MQTT através do uso da biblioteca Eclipse Paho.

Dessa maneira, o simulador é capaz de enviar informações em um tópico com o formato de dados consistentes com o datasheet de valores simulados de um sensor.

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
    "name": <nome_sensor>,
    "longitude": <longitude>,
    "latitude": <latitude>,
    "rate": <taxa_transmissao>, <!-- hertz -->
    "unit": <unidade_do_sensor>
}
```

Por outro lado, temos o subscriber:

```
cd subscriber
go run subscriber.go
```

## Demonstração

O vídeo de demonstração se encontra na pasta \demonstração.
