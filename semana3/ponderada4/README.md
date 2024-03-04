# Ponderada 4 - Módulo 9
## Integração simulador com HiveMQ
A partir do simulador de dispositivos IoT desenvolvido na ponderada 1 e 2, foi implementado a integração entre o simulador e um cluster configurado no HiveMQ, utilizando a camada de transporte TLS.

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
Por outro lado, temos o subscriber:

```
cd subscriber
go run subscriber.go
```
### Hive MQ
A partir dessa integração, a maneira de criação de um cliente utilizando as credenciais de autenticação definidas no cluster do Hive MQ:

``` 
func CreateClient(broker string, id string, callback_handler mqtt.MessageHandler, user string, password string) mqtt.Client {

	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tls://%s:%d", broker, 8883))
	opts.SetClientID(id)
	opts.SetUsername(user)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(callback_handler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	return mqtt.NewClient(opts)
}
```

Dessa maneira, tais informações são coletados no arquivo .env (ou definidas no client.go):
```
err := godotenv.Load("./.env")

	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	newBroker := os.Getenv("BROKER_ADDR")
	user := os.Getenv("HIVE_USER")
	pswd := os.Getenv("HIVE_PSWD")
```

## Demonstração

O vídeo de demonstração se encontra na pasta \demonstração.

