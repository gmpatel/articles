# Articles - article api service (golang | go)
  
**POST /articles** HTTP/1.1  
Host: {server}:{port}  
Content-Type: application/json  
{  
  "title": "latest science shows that potato chips are better for you than sugar",  
  "body" : "some text, potentially containing simple markup about how potato chips are great",  
  "tags" : ["health", "fitness", "science"]  
}  
  
**GET /articles** HTTP/1.1  
Host: {server}:{port}  
  
**GET /articles/{id}** HTTP/1.1  
Host: {server}:{port}  
  
**GET /tag/{tagName}/{date}** HTTP/1.1  
Host: {server}:{port}  

## Source code of the solution

### In the MAC terminal please run the script section below to pull the full source code of the solution

$ mkdir ~/go/src/github.com/gmpatel/  
$ cd ~/go/src/github.com/gmpatel/  
$ git clone https://github.com/gmpatel/articles.git
  
## Setup/build/run the service

### In MAC terminal please run the script below to start the service

$ cd ~/go/src/github.com/gmpatel/articles/cmd/article-api/  
$ go build  
$ ./article-api  

The service should start listening on port 8083 (expecting everything went well)  

There is a **postman.json** file included in the source code for you to hit the endpoints created in no time with all the sample URLs and JSON provided. Please install **POSTMAN** and import the **POSTMAN.JSON** file.

$ cat ~/go/src/github.com/gmpatel/articles/postman.json  

## Setup the database  

For the ease of the demonstration, the service connects to my SQL Server Database in cloud by default.  
  
**DefConnStr**: "server=mssqlserver; user id=my-user; password=mypasswd; database=prefix_articles"  

## Override connection string  

**You can override the default connection string by setting up the environment variable for the connection string.**  

1. Create new blank database **i.e Articles** into your local OR intranet sqlserver.  
2. Run the script **$ cat ~/go/src/github.com/gmpatel/articles/scripts/db/create.sql** against the blank db you just created in step 1 to create the necessary database assets for the service to store data.  
3. Setup environment variable **APP_CONN_STRING** with the value of the new connection string as per the example given below. 


**$ export APP_CONN_STRING="server=localhost; user id=sa; password=welcome123; database=Articles"**  

## Override port

**You also can override the default port service is listening on by setting the environment variable for the port to listen on.**

1. Setup the environment variable **APP_LISTEN_PORT** with the value of the new port you want service to listen on, as per the example given below.  

**$ export APP_LISTEN_PORT=8088**  

## Feedback about the coding test

1. It was really good exercies to do in golang.
2. It wasn't 3/4/5 hours exercies for me for the scale of solution I have written. It took me roughly 12/13 hours to reach at this stage where the solution is at the moment.
3. I would have done **Unit Testing** and **GODOG Acceptance Testing** as well time had permitted.
4. Even though I have not done **Unit Testing** and **GODOG Acceptance Testing**, I would say I have considerable experience around that part of the solution development.
5. I like the **/tag/{tagName}/{date}** endpoint. It was fairly complex to query data, but, really good.
6. I would have used some other database or could have kept data in memory, buy, then I was wanted to work myself on SQLServer side of Golang so, I took this as an opportunity to code GoLang with SQLServer.

## Assumptions

### Count field in response json of /tag/{tagName}/{date}
  
{  
    "tag": "health",  
    "count": 2,  
    "articles": [  
        "1",  
        "2"  
    ],  
    "related_tags": [  
        "fitness",  
        "science"  
    ]  
}  

The count field here doesn't make any sense, I have put it as a count of articles found for the given tag on given date as a part of my solution, but, then that 2 can be counted from the articles[] field of the response json as well. But if that is for the different purpose then it does not clearly reflect what should be brought over there for the count field.
