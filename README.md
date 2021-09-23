# newsapi-go
This project has the intent to provide a valid interface for newsapi (https://newsapi.org/).

# usage
The project is still under mantainance and will be released as soon as possible, actually the interface is quite trivial but will be improved.
To run the example main.go you should modify the config-file.json to customize news generation.

The command `go run main.go` in your terminal should execute the example program.

# architecture
The project is based on a simple architecture composed by a client and a server running in a container inside an server instance. The server is kept secret because it contains the API Key
![general architecture](docs/general_architecture.jpg)

# todo list
- [ ] write a decent readme and a documentation
- [x] import correctly go module
- [x] write an example source code 
- [x] Include OAuth to generate API Key (proxy server)
