# GoLogGuardian

O Log Guardian é uma biblioteca desenvolvida para padronizar e gerenciar logs de maneira eficiente e organizada em sistemas. Com a flexibilidade de lidar com diferentes tipos de logs, esta biblioteca proporciona uma estrutura consistente para a geração e gerenciamento de logs gerais e específicos, permitindo a integração com diferentes sistemas.

## Funcionalidades

- Padronização de Logs: O Log Guardian oferece uma estrutura unificada para diferentes tipos de logs, desde logs de função, operações em banco de dados até logs de requisições.

- Configuração Flexível: Permite a fácil integração com diferentes sistemas, possibilitando a customização e configuração dos logs de acordo com as necessidades específicas.

- Conexão com Banco de Dados: Além da gestão dos logs, o Log Guardian pode se integrar a diferentes tipos de banco de dados, como PostgreSQL, MySQL, SQLite e MongoDB. A configuração é simples, utilizando variáveis de ambiente para especificar os detalhes de conexão.

- Inserção Automática de Logs: Quando configurado corretamente com variáveis de ambiente, o Log Guardian é capaz de inserir automaticamente os logs no banco de dados especificado.

# Uso
O Log Guardian é flexível e se adapta à configuração do ambiente em que é executado. Se as variáveis de ambiente necessárias não estiverem configuradas, o Log Guardian ainda poderá retornar o modelo do log específico para inserção manual no banco de dados.

Caso as variáveis de ambiente estejam configuradas corretamente com os detalhes do banco de dados desejado, o Log Guardian é capaz de conectar automaticamente ao banco de dados especificado e inserir os logs diretamente na tabela correspondente. Ele retorna um valor booleano indicando o sucesso ou falha na inserção dos dados no banco.

Essa flexibilidade permite uma fácil integração e uso do Log Guardian em diferentes cenários de configuração, seja para apenas fornecer os modelos de log para inserção manual ou para realizar inserções automáticas no banco de dados configurado.

É recomendável consultar a seção de Configuração para detalhes sobre as variáveis de ambiente necessárias para uma configuração completa do Log Guardian.

## Configuração

O Log Guardian utiliza variáveis de ambiente para configurar suas operações de banco, incluindo definições de conexão com banco de dados e outras configurações essenciais. Aqui está a lista das variáveis de ambiente disponíveis:

### Banco de Dados Relacional

#### PostgreSQL
```
POSTGRES_HOST: Define o endereço do host para o PostgreSQL.
POSTGRES_EXTERNAL_PORT: Especifica a porta externa para o PostgreSQL.
POSTGRES_USER: Nome de usuário para autenticação no PostgreSQL.
POSTGRES_PASSWORD: Senha para autenticação no PostgreSQL.
POSTGRES_DB: Nome do banco de dados PostgreSQL a ser utilizado.
```
#### MySQL
```
MYSQL_HOST: Define o endereço do host para o MySQL.
MYSQL_PORT: Especifica a porta para o MySQL.
MYSQL_USER: Nome de usuário para autenticação no MySQL.
MYSQL_PASSWORD: Senha para autenticação no MySQL.
MYSQL_DBNAME: Nome do banco de dados MySQL a ser utilizado.
```
#### SQLite
```
SQLITE_PATH: Caminho do arquivo SQLite, se for o banco de dados escolhido.
```
### Banco de Dados NoSQL
#### MongoDB
```
MONGODB_URI: Define o URI de conexão para o MongoDB.
MONGODB_DBNAME: Nome do banco de dados MongoDB a ser utilizado.
```
#### Configuração Geral
```
LOG_GUARDIAN_DATABASE_TYPE: Especifica o tipo de banco de dados a ser utilizado pelo Log Guardian (Valores: sqlite, postgres, mysql, mongodb).
```

#### Observações:

Para utilizar as funções automáticas do GoLogGuardian, é obrigatório o uso da variável LOG_GUARDIAN_DATABASE_TYPE, pois algumas validações são executadas com base nesta variável antes de chamar as rotinas de inserção.

Certifique-se de fornecer valores válidos e corretos para cada uma dessas variáveis de ambiente. Isso garante uma conexão adequada e o funcionamento correto do Log Guardian com o banco de dados desejado.