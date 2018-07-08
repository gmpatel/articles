# Articles
golang exercise - article api service

## Source code of the solution

### Please run the script section below to pull the full source code of the solution

$ mkdir ~/go/src/github.com/gmpatel/  
$ cd ~/go/src/github.com/gmpatel/  
$ git clone https://github.com/gmpatel/articles.git
  
## Setup the service

### In MAC terminal window please run the script below to start the service

$ cd ~/go/src/github.com/gmpatel/articles/cmd/article-api/  
$ go build  
$ ./article-api  

The service should start listening on port 8083 (expecting everything went well)  

There is a **postman.json** file included in the source code for you to hit the endpoints created in no time with all the sample URLs and JSON provided. Please install **POSTMAN** and import the **POSTMAN.JSON** file.

$ cat ~/go/src/github.com/gmpatel/articles/postman.json  

## Setup database  

For the ease of the demonstration, the service connects to my SQL Server Database in cloud by default.  
  
**DefConnStr**: "server=mssqlserver; user id=my-user; password=mypasswd; database=prefix_articles"  

## Overrides  

###You can override the default connection string by setting up the environment variable for the connection string.  

1. Create new blank database **i.e Articles** into your local OR intranet sqlserver.  
2. Run the script **$ cat ~/go/src/github.com/gmpatel/articles/scripts/db/create.sql** against the blank db you just created in step 1 to create the necessary database assets for the service to store data.  
3. Setup environment variable **APP_CONN_STRING** with the value of the new connection string as per the example given below.  
**export APP_CONN_STRING="server=localhost; user id=sa; password=welcome123; database=Articles"**  

###You can also override the default port service is listening on by setting the environment variable for the port to listen on.

1. Setup the environment variable **APP_LISTEN_PORT** with the value of the new port you want service to listen on, as per the example given below.  

**export APP_LISTEN_PORT=8088**  


