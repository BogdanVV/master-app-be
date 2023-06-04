# Decisions made throughout developing process

## 1. Using cookies-based authentication

I haven't been working with cookie-based auth yet, so it would be cool to figure out wtf this is and how to work with that. In terms of security there's no final conclusion on what approach is better - jwt-based or cookies-based auth. From what I can tell after walking through a dozen of articles on that matter it's up to specific requirements (which isn't hte case) or up to personal preferences. So why not trying cookies? :)
Since I like jwt schema more, I'm gonna combine jwt and cookies based auth - jwt are gonna be stored in cookies.

## 2. Switching to jwt-based auth

I'm creating mobile master-app using React Native. And seems like RN doesn't deal well with cookies - it stores them for some time but after a few requests or some period of time issues pop up:

- https://javascript.plainenglish.io/react-native-cookie-authentication-83ef6e84ba70
- https://reactnative.dev/docs/0.61/network#known-issues-with-fetch-and-cookie-based-authentication.

Though I'm not experiencing these issues at the moment on dev (local) environment I don't wanna face it later.
So, the decision is to move to jwt-based auth.
The good news is that figured out how to work with cookies (not a big deal actually), so I assume the mission (to figure out wft is cookies and how to work with that) is acomplished anyway.

# Running the project

## Files needed for running the project

- `config/config.yml`
- `.env`

## Running the project

```bash
go mod tidy
```

```bash
go run main.go
```

## Migrations

### Create migrations:

```bash
migrate create -ext sql -dir ./migrations -seq init
```

### Run up migration

```bash
migrate -database "postgresql://{USERNAME}:{PASSWORD}@{HOST}:{PORT}/{DBNAME}?sslmode=disable" -path ./migrations up
```

### Run down migration

```bash
migrate -database "postgresql://{USERNAME}:{PASSWORD}@{HOST}:{PORT}/{DBNAME}?sslmode=disable" -path ./migrations down
```
