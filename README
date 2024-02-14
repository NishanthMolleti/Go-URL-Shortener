# GoLang URL Shortener 

##   Introduction 
This backend application works on your local efficiently . It   is a simple URL shortener written in Golang. The main purpose of this project is to create
This is a simple URL shortener written in Golang. 

## Steps to Use 
There are 2 APIs in this application namely 
1. localhost:9808/create-short-url  - This API takes the long url and a unique ID as input and returns the generated shortened URL. (This is a POST API).
2. localhost:9808/DvobfWfx - This API takes the 8 unique character value assigned by the algorithm and redirects you to the lonk URL. (This is a GET API).

## Steps to Setup in local 
* Fork my repo and clone it onto your local.
* Install Redis on you local if you dont have by using `brew install redis` , then use `brew services start redis` to start the Redis on your local.
* Make sure that port 9808 is not being utilized anywhere else. You can change the port on main.go to any other available port if needed. 
* Now use `go run main.go` to run the service.
* Now open postman and do a POST request on `http://localhost:9808/create-short-url` with JSON body as `{
    "long_url": "long_url_that_you_want_to_shorten",
    "user_id" : "put_in_a_unique_user_id_here"
}`
Your response should be something like below
`{
    "message": "short url created successfully",
    "short_url": "http://localhost:9808/9Zatkhpi"
}`
* You can use the 2nd api to redirect to the actual website with long url.

## Next Steps 
* If you like the application and would like to work on it then  please fork the repository, make changes according to your needs and create a pull request. I will review. I suggest people to raise an issue and dockerise the backend service. I have also attached the Dockerfile above. I have hosted the docker image on Azure and it works fine. 
* Try creating a simple frontend for the application as well. 

