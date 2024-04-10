# Ponderada 8 - Módulo 9
## DESAFIO 1BR

O desafio consiste em escrever um programa que recupere valores de medição de temperatura de um arquivo de texto e calcule a temperatura mínima, média e máxima por estação meteorológica, com um arquivo em parquet com 1.000.000.000 de linhas.

Dessa maneira, o script feito pode ser explicado da seguinte maneira:

1. Importação de bibliotecas: As bibliotecas time, os e pandas são importadas. time é usada para medir o tempo de execução, os para obter informações sobre o arquivo, e pandas para processar os dados.
2. Função process_file: Esta função recebe o caminho do arquivo como entrada.
3. Medição do tempo de execução: Antes de começar a processar o arquivo, a função process_file registra o tempo atual para posterior comparação do tempo de execução.
4. Leitura do arquivo Parquet: O arquivo Parquet é lido usando a função pd.read_parquet. Este método do Pandas lê o arquivo Parquet e carrega-o em um DataFrame.
5. Cálculo das estatísticas de temperatura: O DataFrame é então agrupado pela coluna 'Station' (estação) e as estatísticas de temperatura (mínimo, média e máximo) são calculadas usando a função agg.
6. Ordenação dos resultados: Os resultados são ordenados em ordem alfabética pelo nome da estação usando a função sort_values.
7. Impressão dos resultados: Os resultados são impressos na tela.
8. Cálculo do tempo de execução: Após o processamento do arquivo, o tempo atual é registrado novamente e subtraído do tempo inicial para obter o tempo total de execução.
9. Função main: Esta função é responsável por configurar o caminho do arquivo Parquet e exibir o tamanho do arquivo em megabytes antes de chamar a função process_file.
10. Execução do programa: O programa é executado chamando a função main.

## Demonstração

O vídeo de demonstração se encontra na pasta \demonstração.