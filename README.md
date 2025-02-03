# Random Affirmations Generator

This is a simple Go application that randomly selects affirmations from a predefined list and prints them. It includes tests to ensure the functionality of the randomization and the validity of the affirmations.

## Requirements

- Go 1.x or higher
- Git

## Installation

1. Clone the repository:
   ```bash
      git clone https://github.com/yourusername/affirmations.git
   ```

2. Change to the project directory:
   ``` bash
      cd affirmations
   ```
3. (Optional) Install Go, if you don’t already have it installed. You can download Go from the [official website](https://go.dev/dl/).

## Running the Application

To run the application and print a random affirmation:
```bash
go run affirmations.go
```

You should see an output like:
```css
I am grateful for every moment in my life
```

## Running Tests

To run the tests for this application:
```bash
go test -v
```

You will see output similar to:
```diff
=== RUN   TestAffirmationsHasAtLeastFiveOptions
    affirmations_test.go:13: There are 7 possible affirmations
--- PASS: TestAffirmationsHasAtLeastFiveOptions (0.00s)
=== RUN   TestRandomAffirmationReturnsAllValidOptions
    affirmations_test.go:47: Found 7 of 7 possible affirmations in 25 retrieval attempts
--- PASS: TestRandomAffirmationReturnsAllValidOptions (0.00s)
PASS
ok    example/affirmations 0.186s
```

## Formatting the Code

To format all go files for this application:
```bash
go fmt
```
This will automatically format your Go code according to the standard Go conventions.

## Code Structure

* `affirmations.go`: Contains the logic for random affirmation generation and the main function.
* `affirmations_test.go`: Contains tests for ensuring the validity and randomization of the affirmations.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
