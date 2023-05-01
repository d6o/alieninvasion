# Alien Invasion

Alien Invasion is a command-line program written in Go that simulates an alien invasion on the non-existent world of X. The program reads a map from a given input file, spawns a specified number of aliens, and simulates their movement and battles until all aliens are destroyed or each alien has moved 10,000 times.

## Installation

To build the program, run the following command:

```
go build -o alienInvasion main.go
```

## Usage

To get help for the program, run:

```
./alienInvasion --help
```

To simulate an alien invasion with 2 aliens using the `./input` file and enable verbose logs, execute:

```
./alienInvasion -aliens 2 -input ./input -v
```

### Command-line Arguments

- `-aliens int`: Defines the number of aliens to be spawned (default 3)
- `-input string`: Input file path. Defaults to stdin
- `-output string`: Output file path. Defaults to stdout
- `-v`: Verbose mode - enable logs

## Getting Started

Follow the steps below to use the Alien Invasion program:

1. Clone the repository or download the source code.

2. Navigate to the project directory and build the program with the following command:

```
go build -o alienInvasion main.go
```

3. Prepare an input file according to the challenge description. The file should contain city names and their respective connections. For example:

```
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
```

4. Run the Alien Invasion simulation with your desired parameters. For instance, to simulate an invasion with 5 aliens using the `./input` file and enable verbose logs, execute:

```
./alienInvasion -aliens 5 -input ./input -v
```

5. Observe the output, which will show alien movements, battles, and destroyed cities. The remaining state of the world will be printed in the same format as the input file.

## Tests

To ensure the reliability and correctness of the alienInvasion, it is essential to run tests. This section explains how
to execute tests for the alienInvasion.

### Running Tests

To run all tests, execute the following command:

```bash
go test -v ./...
```

This command will run all tests in the alienInvasion project and display the results, including passed and failed tests,
with verbose output.

### Running Tests with Race Condition Check

To run all tests with a race condition check, execute the following command:

```bash
go test -v -race ./...
```

This command will run all tests and check for race conditions in the alienInvasion project. The race detector is a
valuable tool for identifying data races in your code, which can lead to unexpected behavior and hard-to-debug issues.

By regularly running tests and checking for race conditions, you can ensure the quality and stability of the
alienInvasion.

## Generating Mocks

When developing and testing the alienInvasion, it's essential to use mocks to isolate components and simulate external
dependencies' behavior. To generate mocks for the alienInvasion, you need to have `gomock` installed. This section
explains how to install `gomock` and regenerate mocks.

### Installing gomock

To install `gomock`, run the following command:

```
go install github.com/golang/mock/mockgen@v1.6.0
```

This command installs the `gomock` package and the `mockgen` tool at the specified version (v1.6.0).

### Regenerating Mocks

To regenerate the mocks for the alienInvasion, run the following command:

```
go generate ./...
```

This command will traverse all directories within the project and regenerate mocks based on the source files and
interfaces defined in the project. The newly generated mocks can be used in your tests to isolate components and
simulate the behavior of external dependencies, making it easier to write effective and accurate test cases.

## Code Quality and Linting

To maintain high-quality code and ensure adherence to best practices, you can use the `golangci-lint` tool. This section
explains how to run `golangci-lint` using Docker.

### Running golangci-lint with Docker

To run `golangci-lint` with Docker, execute the following command:

```
docker run -t --rm -v $(pwd):/app -w /app golangci/golangci-lint golangci-lint run -v
```

This command will run `golangci-lint` in a Docker container, mounting the current directory (`$(pwd)`) to `/app` inside
the container, and setting the container's working directory to `/app`.

The `-v` flag enables verbose output, displaying the results of the linting process. If any issues are
found, `golangci-lint` will report them, allowing you to review and correct the problems to maintain high-quality code.

By regularly running `golangci-lint`, you can ensure that your code follows best practices and prevent potential issues
in the alienInvasin.

## Contributing

Contributions to the Alien Invasion program are welcome. To contribute, follow these steps:

1. Fork the repository.

2. Create a new branch with a descriptive name.

3. Make your changes in the new branch.

4. Commit your changes and write a clear and concise commit message.

5. Push your changes to your fork.

6. Open a pull request, describing the changes you've made and the reasons for them.

## License

The Alien Invasion program is released under the [MIT License](LICENSE).
