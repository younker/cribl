# Thoughtful AI Technical Screen

## Getting Started

The application read in from STDIN and prints the desired stack to STDOUT with errors being placed on STDERR.

Due to time constraints, I was not able to dockerize this in order to facilitate running it locally nor was I able to upload it to `replit.com`.

There are 2 ways to run the program.

1. If you are on a arm based Mac, there is a binary compiled in `./bin` which you can run. See below for how it was built and how to run it:

```shell
➜  thoughtful_ai git:(main) ✗ GOARCH=arm64 GOOS= go build -o bin/stack_dispatch
➜  thoughtful_ai git:(main) ✗ ./bin/stack_dispatch 2>stdout.txt
19,20,20,20
STANDARD
20,20,20,20
SPECIAL
19,1000,1000,1
SPECIAL
20,1000,1000,1
REJECTED
bogus,input
^C
➜  thoughtful_ai git:(main) ✗ head stdout.txt
unable to parse input 'bogus,input' due to: inputs should contain 4 values 'mass,height,width,depth' in that order. We received invalid input 'bogus,input'
➜  thoughtful_ai git:(main) ✗
```

If you have go `1.23.1` installed, you can also run it using the following:

```
$ go run main.go 2> stdout.txt
```

You can send the kill signal to the process by running `ctrl+c`.

Note that `STDERR` was redirected to `stdout.txt` so any errors would not be ingested by the downstream consumer though in hindsight it would seem that printing `REJECTED` to `STDOUT`, in addition to logging the error to `STDERR`, would be prudent for such a outcome.

## Running Tests

You can run the tests using the following:

```shell
➜  thoughtful_ai git:(main) ✗ pwd
~/projects/thoughtful_ai
➜  thoughtful_ai git:(main) ✗ go test ./...
?   	github.com/younker/thoughtful_ai	[no test files]
?   	github.com/younker/thoughtful_ai/adapter	[no test files]
ok  	github.com/younker/thoughtful_ai/model	(cached)
```

## Methodology

I read the [Platform Technical Screen documentation](https://thoughtfulautomation.notion.site/Platform-Technical-Screen-b61b6f6980714c198dc49b91dd23d695) 3 days ago. Since then, I have considered some of the technical aspects such that when I sat down to work on the problem today, I had already mapped out most of the domain models, functionality, inputs ane outpus. This allowed me to do more in order to provide a more realistic demonstration of the kind of code I produce and the standards and practices I advocate for.

However, despite this approach I still went over the requested allocated time by 23 minutes. As a result, there are missing tests and changes I would would make should this be shipped to production (eg more exhaustive tests, dockerization, ci/cd, etc).
