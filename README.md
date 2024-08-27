# JSON To CSV

Simple program to convert a file containing a single JSON object to a CSV file.

## Usage

Install the dependencies:

```bash
go mod tidy
```

Run the program:

```bash
go run main.go --file <path-to-json-file> --key <key> --value <value> --output <path-to-output-csv-file> --exclude-keys <comma-separated-list-of-keys-to-exclude>
```
