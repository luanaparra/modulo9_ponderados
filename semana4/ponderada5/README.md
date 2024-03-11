# Ponderada 3 - Módulo 9
## Simulação de ataques usando MQTT
Atividade para subir um broker remoto/local do MQTT para conduzir cenários de análise vulnerabilidade, identificando situações onde pode ser comprometido cada um dos três pilares.

**O que acontece se você utilizar o mesmo ClientID em outra máquina ou sessão do browser? Algum pilar do CIA Triad é violado com isso?**

Utilizar um clientID em outra máquina/sessão, a conexão irá se desconectar. Dessa maneira, tal informação única pode ser vazada, tendo violado o pilar de Confidencialidade. 

**Sem autenticação (repare que a variável allow_anonymous está como true), como a parte de confidencialidade pode ser violada?**

No contexto de Integridade, é possível identificar a violação no momento em que qualquer usuário pode publicar informações e subscrever-se nos tópicos, haja vista que a informação não é restrita. 


**Com os parâmetros de resources, algum pilar do CIA Triad pode ser facilmente violado?**

O pilar de Disponibilidade pode ser violado, tendo em vista que o uso de recursos computacionais podem afetar a performance por limites operacionais restritos.


## Demonstração

Os vídeos de demonstração se encontram na pasta \demonstração.

