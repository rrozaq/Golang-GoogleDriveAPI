# Google Drive API
this is simple implementation for API google drive using golang

## How to use
- create service google drive api, you can read the docs for information https://developers.google.com/drive/api/quickstart/go?hl=id
- Download your credential file and rename it to 'credentials.json'. put it in the root folder of this project

#### Run project
```bash
go run cmd/api/main.go
```
or can you run using air, but make sure install library air https://github.com/cosmtrek/air
```bash
air
```

- now you can start by calling the auth api

```bash
Login using oauth2:
Method: GET
URL: {{ _.base_url }}/api/auth

GET ALL FILE:
Method: GET
URL: {{ _.base_url }}/api/file

GET ALL FILE:
Method: POST
URL: {{ _.base_url }}/api/upload
BODY FORM: file
KEY FORM: file

```

maybe can you create PR to cleancode or fixing bug and other
thanktou.
