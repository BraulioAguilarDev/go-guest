# Event Go

API REST to generate events or date with many guets

## Docker

Copy .env file

```sh
$ cp .env.example .env
```

## API REST

### Create event

`POST: /api/events`

Request:

```json
{
	"name": "My event 1",
	"location": "CDMX",
	"date": "01-01-2021",
	"time": "12:00"
}
```

### Create user

`POST: /sing-up`

Request:

```json
{
  "first_name": "Braulio",
  "last_name": "Aguilar",
  "email": "simoncositas@domain.com",
  "password": "secret"
}
```
