# Part 2

This is a quick service in Go using Gin library to achieve what is requested. Note that this won't survive restart (request storage is held in runtime), to go beyond we'd obviously use external SQL/NoSQL storage.

To run server:
```
PORT=8080 go run main.go
```
To test:
```
❯ curl -i http://localhost:8080/api/req

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Wed, 26 Jul 2023 01:10:10 GMT
Content-Length: 55

{"generated_id":"7658f9da-ec83-4d53-b6f9-b3ee18717cb2"}%                        
~ on ☁️  (ap-southeast-2) 
❯ curl -i http://localhost:8080/api/list_ids

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Wed, 26 Jul 2023 01:10:18 GMT
Content-Length: 361

{"ids":{"176e128f-e856-4d4c-8e7b-db09f5c21abe":true,"4376b08e-d14e-44c4-b712-593e0e7b6b76":true,"5d3dd6d7-82a2-4490-91b0-ea728d973ea4":true,"6289affd-d938-4a41-8a0d-8e266807735c":true,"7658f9da-ec83-4d53-b6f9-b3ee18717cb2":true,"c443f9db-c19b-4f3b-9ff8-60165592bc8d":true,"d3157764-4d60-4a70-a83a-a615e58ad46a":true,"e8f304eb-fc0f-41f4-b5cd-e60b7f2903c4":true}}%
```