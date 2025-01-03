<!-- markdownlint-disable MD033 MD050 -->
# go-todotxt

[![go1.22+](https://img.shields.io/badge/Go-1.22+-blue?logo=go)](https://github.com/KEINOS/go-todotxt/blob/main/.github/workflows/unit-tests.yml#L81 "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-todotxt.svg)](https://pkg.go.dev/github.com/KEINOS/go-todotxt/todo "View document")
[![License](https://img.shields.io/github/license/KEINOS/go-todotxt)](https://github.com/KEINOS/go-todotxt/blob/master/LICENSE)

`github.com/KEINOS/go-todotxt` is a Go package for parsing and editing todo.txt files, a [text format for task annotations](https://github.com/todotxt/todo.txt) designed by [Gina Trapani](https://github.com/ginatrapani).

> __Note__: This package is based on the following packages with **custom user sort functionality**.
>
> - [**todotxt**](https://github.com/1set/todotxt) from [Kevin Tang](https://github.com/vt128)
> - [**go-todotxt**](https://github.com/JamesClonk/go-todotxt) from [Fabio Berchtold](https://github.com/JamesClonk)

## Usage

```go
// Download the package.
go get "github.com/KEINOS/go-todotxt"
```

```go
// Import the package.
import "github.com/KEINOS/go-todotxt/todo"
```

```go
func Example() {
    // Load tasks from a string.
    // You can also load from a file by using LoadFromFile().
    tasks, err := todo.LoadFromString(`
        (A) Call Mom @Phone +Family
        x (A) Schedule annual checkup +Health
        (C) Add cover sheets @Office +TPSReports
        Plan backyard herb garden @Home
        Pick up milk @GroceryStore
        Research self-publishing services +Novel @Computer
        x Download Todo.txt mobile app @Phone
        (A) This is a task should be due before yesterday due:2020-11-15
    `)
    if err != nil {
        log.Fatal(err)
    }

    // Add a new task.
    // Note the order of project ("+Novel") and context ("@Computer") later
    // in the output.
    newTask, err := todo.ParseTask("(B) Outline chapter 5 +Novel @Computer")
    if err != nil {
        log.Fatal(err)
    }

    tasks.AddTask(newTask) // append to the end of the list

    // Sort tasks by priority and then by context in ascending order.
    if err := tasks.Sort(todo.SortPriorityAsc, todo.SortContextAsc); err != nil {
        log.Fatal(err)
    }

    // AND filter.
    // Get tasks that have any priority AND are not completed AND are not overdue.
    filteredTasks := tasks.
        Filter(todo.FilterHasPriority).
        Filter(todo.FilterNotCompleted).
        Filter(todo.FilterNot(todo.FilterOverdue)) // NOT overdue

    // OR filter.
    // Filter the above tasks with priority "A" OR has project "Novel" OR has
    // context "Office".
    filteredTasks = filteredTasks.Filter(
        todo.FilterByPriority("A"),     // has (A)
        todo.FilterByProject("Novel"),  // has +Novel
        todo.FilterByContext("Office"), // has @Office
    )

    // Print the filtered tasks.
    for _, task := range filteredTasks {
        fmt.Println(task.String())
    }
    // Output:
    // (A) Call Mom @Phone +Family
    // (B) Outline chapter 5 +Novel @Computer
    // (C) Add cover sheets @Office +TPSReports
}
```

```go
func ExampleTaskList_CustomSort() {
    tasks, err := todo.LoadFromString(`
        Task 3
        Task 1
        Task 4
        Task 2
    `)
    if err != nil {
        log.Fatal(err)
    }

    // customFunc returns true if taskA is less than taskB.
    customFunc := func(taskA, taskB todo.Task) bool {
        // Compare strings of the text part of the task.
        return taskA.Todo < taskB.Todo
    }

    tasks.CustomSort(customFunc)

    fmt.Println(tasks[0].Todo)
    fmt.Println(tasks[1].Todo)
    fmt.Println(tasks[2].Todo)
    fmt.Println(tasks[3].Todo)
    // Output:
    // Task 1
    // Task 2
    // Task 3
    // Task 4
}
```

## Todo.txt format

![](https://raw.githubusercontent.com/todotxt/todo.txt/master/description.svg)

- Image from: [https://github.com/todotxt/todo.txt](https://github.com/todotxt/todo.txt)

## Contributing

[![go1.22+](https://img.shields.io/badge/Go-1.22+-blue?logo=go)](https://github.com/KEINOS/go-todotxt/blob/main/.github/workflows/unit-tests.yml#L81 "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/go-todotxt.svg)](https://pkg.go.dev/github.com/KEINOS/go-todotxt/todo "View document")

Any contribution for the better is welcome. We provide full code coverage of unit tests, so feel free to refactor or play around with the code.

- Branch to PR:
  - `main` ([Draft PR](https://github.blog/2019-02-14-introducing-draft-pull-requests/) is recommended)
- [Open an issue](https://github.com/KEINOS/go-todotxt/issues)
  - Please attach a simple and reproducible test code if possible. This helps us alot and to fix the issue faster.
- [CI](https://en.wikipedia.org/wiki/Continuous_integration)/[CD](https://en.wikipedia.org/wiki/Continuous_delivery):
  - The below tests will run on Push/Pull Request via GitHub Actions. You need to pass all the tests before requesting a review.
    - Unit testing on various Go versions (1.22 ... latest)
    - Unit testing on various platforms (Linux, macOS, Windows)
    - Static analysis/lint check by [golangci-lint](https://golangci-lint.run/).
      - Configuration: [.golangci.yml](./.golangci.yml)
  - To **run tests locally**, we provide a convenient [Makefile](./Makefile). Please run the below command to run all the tests (`docker` and `compose` are required).

    ```bash
    # Runs unit tests on Go 1.22 to latest and `golangci-lint` check.
    make test
    ```

> __Note__ : Please **DO NOT PR to the `original` branch** but to `main` branch. The branch `original` is simply a copy from the `master` branch of the [upstream repo](https://github.com/1set/todotxt). This is for the purpose of keeping the original code as is and contribute to the upstream.

## Statuses

[![UnitTests](https://github.com/KEINOS/go-todotxt/actions/workflows/unit-tests.yml/badge.svg)](https://github.com/KEINOS/go-todotxt/actions/workflows/unit-tests.yml)
[![golangci-lint](https://github.com/KEINOS/go-todotxt/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/KEINOS/go-todotxt/actions/workflows/golangci-lint.yml)
[![CodeQL-Analysis](https://github.com/KEINOS/go-todotxt/actions/workflows/codeQL-analysis.yml/badge.svg)](https://github.com/KEINOS/go-todotxt/actions/workflows/codeQL-analysis.yml)
[![PlatformTests](https://github.com/KEINOS/go-todotxt/actions/workflows/platform-tests.yml/badge.svg)](https://github.com/KEINOS/go-todotxt/actions/workflows/platform-tests.yml "Tests on Win, macOS and Linux")

[![codecov](https://codecov.io/gh/KEINOS/go-todotxt/branch/main/graph/badge.svg?token=JVY7WUeUFz)](https://codecov.io/gh/KEINOS/go-todotxt)
[![Go Report Card](https://goreportcard.com/badge/github.com/KEINOS/go-todotxt)](https://goreportcard.com/report/github.com/KEINOS/go-todotxt)

## License and Credits

- MIT License. Copyright (c) 2022 [KEINOS and the go-todotxt contributors](https://github.com/KEINOS/go-todotxt/graphs/contributors) with all the respect to Kevin Tang, Fabio Berchtold and Gina Trapani.
- This package is based on the below packages and ideas:
  - Mother/Upstream: [**todotxt**](https://github.com/1set/todotxt) authored by [Kevin Tang](https://github.com/vt128) @ [MIT](https://github.com/1set/todotxt/blob/master/LICENSE)
  - Grand Mother/Most upstream: [**go-todotxt**](https://github.com/JamesClonk/go-todotxt) authored by [Fabio Berchtold](https://github.com/JamesClonk) @ [MPL-2.0](https://github.com/JamesClonk/go-todotxt/blob/master/LICENSE)
  - Origin: [**todo.txt**](https://github.com/todotxt/todo.txt) is an awesome task format. Initially designed by [Gina Trapani](https://github.com/ginatrapani). @ [GPL-3.0](https://github.com/todotxt/todo.txt/blob/master/LICENSE)
