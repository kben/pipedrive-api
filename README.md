# Golang Pipedrive client

[![Build Status](https://travis-ci.org/kben/pipedrive-api.svg?branch=master)](https://travis-ci.org/kben/pipedrive-api)

Requires Go version 1.7 or greater.

# Supported resources

- [x] Activities
- [x] ActivityFields
- [x] ActivityTypes
- [x] Authorizations
- [x] Currencies
- [x] Deals
- [x] DealFields
- [x] Files
- [x] Filters
- [x] Goals
- [x] Notes
- [x] NoteFields
- [x] Organizations
- [x] OrganizationFields
- [x] Persons
- [x] PersonFields
- [x] Pipelines
- [x] Products
- [x] ProductFields
- [x] Recents
- [x] SearchResults
- [x] Stages
- [x] Users
- [x] User connections
- [x] User settings
- [x] Webhooks

## Installation

```go
go get -v github.com/kben/pipedrive-api/pipedrive
```

## Usage

```go
import "github.com/kben/pipedrive-api/pipedrive"
```

Construct a new Pipedrive client, then use the various services on the client to
access different parts of the API. For example:

```go
    const apiKey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

    client := pipedrive.New(&pipedrive.Config{
        APIKey: apiKey,
    })

    // Return list of all fields for note
    noteFields, _, _ := client.NoteFields.List()

    // You can then access data like this:
    fmt.Println("Success = ", noteFields.Success)
    fmt.Println("First note field: ", noteFields.Data[0].Name)
```

### Integration Tests ###

You can run integration tests from the `test` directory. See the integration tests [README](test/README.md).

## Contributions & Issues

Contributions are welcome. Please clearly explain the purpose of the PR and follow the current style.

Issues can be resolved quickest if they are descriptive and include both a reduced test case and a set of steps to reproduce.

## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.
