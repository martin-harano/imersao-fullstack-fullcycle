# Imersão Full Stack && Full Cycle

Participe: https://imersao.fullcycle.com.br

# Informações do desafio

Nesse desafio você deverá montar seu ambiente de desenvolvimento com Docker utilizando o Dockerfile e docker-compose de nosso projeto prático.

Crie uma struct de User com ID (uuid), Name e Email. Além disso, implemente uma função NewUser que seja capaz de validar que ID, Name e Email são obrigatórios, caso o contrário, ela deve retornar um erro.

## Como executar

Utilizamos Docker para que todos os serviços que utilizaremos fiquem disponíveis.

- Faça o clone do projeto
- Tendo o docker instalado em sua máquina apenas execute:
`docker-compose up -d`

### Como executar a aplicação
- Acesse o container da aplicação executando: `docker exec -it desafio1_app_1 bash`
- Rode `go run main.go adduser -n <nome> -e <email>`

**Importante:** Esse código está sendo disponibilizado conforme o andamento das aulas, logo, o arquivo para executar o projeto talvez ainda não tenha sido criado.

### Serviços utilizados ao executar o docker-compose

- Aplicação principal
- Postgres
- PgAdmin
- Apache Kafka
- Criador dos tópicos a serem utilizados pelo Kafka
- Confluent control center
- ZooKeeper

 
