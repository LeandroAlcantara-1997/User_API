

# **User API com Scaffold**

# **O que foi utilizado?**

* Persistencia de dados com biblioteca ORM GORM;
* Persistencia de dados no PostgreSQL.
* Arquitetura de pastas com o GoScaffold;
* Utilização de Docker para criação de containers;
* Desenvolvimento com Remote-Containers.

# **Como executar este projeto?**

[Tutorial de como criar o ambiente para o GoScaffold](https://github.com/LeandroAlcantara-1997/Tech-Doc/blob/master/Go-Scaffold/goscaffold.md)

Copie o conteudo do arquivo [application.env.sample](env/application.env.sample) e crie um novo arquivo application.env com o conteudo do arquivo copiado.

Agora só basta executar um dos seguintes comandos:

~~~make 
make api
~~~

ou 

~~~
make hot
~~~

# **Como acessar ao banco?**

Depois de baixar a imagem e subir o container via remote-containers, acesse o PGAdmin http://localhost:16543 (isso pode demorar um pouco), logue com as credenciais no [docker-compose.yml](build/docker-compose.yaml).

Com o PGAdmin logado, crie uma nova conexão clicando em Server/Register/Server... coloque as credenciais do PostgreSQL informadas no [docker-compose.yml](build/docker-compose.yaml) e assim terá acesso a tabela user criada na applicação.