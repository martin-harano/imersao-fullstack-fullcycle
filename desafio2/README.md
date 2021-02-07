# Imersão Full Stack && Full Cycle

Participe: https://imersao.fullcycle.com.br

# Informações do desafio

Nesse desafio você deverá criar uma simples aplicação utilizando Golang que seja capaz de publicar uma mensagem em um tópico do Apache Kafka e também consumir mensagens desse tópico.

Para deixar mais claro a separação das responsabilidades, crie uma pasta chamada producer e outra consumer. Em cada uma delas crie um arquivo main.go que será responsável por produzir e consumir as mensagens respectivamente

Utilize o mesmo Dockerfile e docker-compose.yaml utilizados no projeto do CodePix. Fique na liberdade para escolher o nome do tópico que será utilizado

## Como executar

NOTA: Apesar do desafio propor utilizar dois main.go distintos, esta implementação separa em dois comandos. Um para consumir os tópicos kafka e outro para publicar registro de usuário.


Utilizamos Docker para que todos os serviços que utilizaremos fiquem disponíveis.

- Faça o clone do projeto
- Tendo o docker instalado em sua máquina apenas execute:
`docker-compose up -d`

### Como executar a aplicação

- Acesse o container da aplicação executando: `docker exec -it desafio2_app_1 bash`

#### Produtor
- Rode `go run main.go adduser -n <nome> -e <email>`
- (A chamada é blocante, para mandar outro comando interrompa com ^C)

#### Consumidor
- Rode `go run main.go consume`

**Importante:** Esse código está sendo disponibilizado conforme o andamento das aulas, logo, o arquivo para executar o projeto talvez ainda não tenha sido criado.

### Serviços utilizados ao executar o docker-compose

- Aplicação principal
- Postgres
- PgAdmin
- Apache Kafka
- Criador dos tópicos a serem utilizados pelo Kafka
- Confluent control center
- ZooKeeper

 
