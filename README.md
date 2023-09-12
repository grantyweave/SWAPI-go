# SWAPI-go

SWAPI-go is a simple Go application that fetches and displays a list of planets from the [Star Wars API (SWAPI)](https://swapi.dev/) and then sorts and displays them alphabetically.

## Installation

Before you can use SWAPI-go, make sure you have Go installed on your system.

1. Clone the repository:

   ```bash
   git clone https://github.com/grantyweave/SWAPI-go.git
   ```

2. Change to the project directory:

   ```bash
   cd SWAPI-go
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

## Usage

Once you've run the application, it will make an HTTP GET request to the SWAPI to fetch a list of planets and then display them both in their original order and sorted alphabetically.

Here's what the application does:

1. Sends an HTTP GET request to the SWAPI's planets endpoint.
2. Parses the JSON response into a list of planets.
3. Prints the list of planets in their original order.
4. Sorts the list of planets alphabetically by name.
5. Prints the sorted list of planets.

## Dependencies

SWAPI-go uses the following dependencies:

- `encoding/json` for JSON decoding.
- `fmt` for printing output.
- `io` for reading the HTTP response body.
- `net/http` for making HTTP requests.
- `sort` for sorting the list of planets.

## Contributing

If you would like to contribute to SWAPI-go, feel free to open a pull request or submit an issue on [GitHub](https://github.com/grantyweave/SWAPI-go).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
