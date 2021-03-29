![Gopher Pirate](https://img.devrant.com/devrant/rant/r_557673_5ZQ9F.jpg)

# API Battleship

API Battleship is a REST API written in Golang for playing the strategy guessing game Battleship.

## Prerequisites

Golang is needed to build the executable: https://golang.org/doc/install

## How to Run the Application

Run `go build` from the project directory, which will generate an executable called `apibattleship`, which you can then run:

```bash
$> ./apibattleship
```

The application was designed to be run in a command line interface.

## Testing

By default, the application starts a server on `localhost:8080`

Create a new game by `POST`ing to the `/games/new` endpoint, which will return a session ID that can be used in the other endpoints.

Setting up a game is done via `POST`ing to the `/games/{sessionId}/setup` endpoint.

Playing a game is done via `POST`ing to the `games/{sessionId}/play` endpoint.

Get a game's status and players by `GET`ing the `games/{sessionId}` endpoint.


## Troubleshooting

If for some reason the executable is unable to run, try deleting it and rerunning `go build` to generate a new executable.

## License
[MIT](https://choosealicense.com/licenses/mit/)