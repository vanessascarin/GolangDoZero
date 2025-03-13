# Banco de Dados

Um banco de dados é um sistema de armazenamento de informações que permite a coleta, o armazenamento, a recuperação e a manipulação de dados de maneira estruturada e eficiente.

No módulo 9, vimos o exemplo da API sem um banco de dados. Assim que "matamos" a aplicação, todos os dados desaparecem.

"Ah mas é só não matar a app". Tá, mas o ponto é que problemas acontecem e a app pode morrer sem a gente querer. Pensando em aplicações críticas, que precisam dos dados sempre ali e atualizados, isso é inadmissível!

E outro ponto: uma coisa é uma app pequena dessa rodando localmente no nosso pc, como no exemplo... outra é uma grande app com um imenso volume de dados!

Existem vários tipos de banco de dados, incluindo relacionais, NoSQL e outros, cada um com suas próprias características e usos específicos.

## Tipos: Relacional e Não Relacional

Existem diversos tipos de banco de dados, mas vamos focar na diferença desses dois: relacional e não relacional.

**Relacional:** Este é o tipo mais comum de banco de dados. Os dados são organizados em tabelas com linhas e colunas, e a relação entre os dados é estabelecida por meio de chaves primárias e estrangeiras.
Exemplos populares incluem MySQL, PostgreSQL, Oracle e Microsoft SQL Server. A linguagem utilizada nesse formato é o SQL.

**Não Relacional:** Bancos de dados NoSQL (ou não relacional) são projetados para lidar com dados não estruturados, semiestruturados ou altamente variáveis.
Eles podem ser baseados em documentos (como o MongoDB), em colunas (como o Apache Cassandra), em gráficos (como o Neo4j) ou em chave-valor (como o Redis).

## ORM

Normalmente, gerenciar dados em banco de dados é algo complexo. Para facilitar essa manipulação dos dados em aplicações, são utilizados ORMs. ORM (Object Relational Mapper) é uma técnica de mapeamento objeto relacional que permite fazer uma relação dos objetos com os dados que os mesmos representam.

O ORM é usado pra abstrair a forma como fazemos consultas no banco de dados. Ao inves de fazermos várias queries usando código SQL (que nem sempre é fácil), usamos o ORM para fazer essa consulta, facilitando essa comunicação.

O ORM do Go é o GORM: https://gorm.io/index.html

E é ele que vamos utilizar para o nosso projeto de app em Go.
