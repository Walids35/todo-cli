# Todo App

## Goal
Create a cli application for managing tasks in the terminal.

```
$ todo-cli
```

## Requirements
Should be able to perform crud operations via a cli on a data file of tasks. The operations Should be as follows:

```
$ todo-cli add Gym
$ todo-cli list
$ todo-cli mark
$ todo-cli delete
```

## Notable Packages used
- **encoding/csv** for writing out as a csv file
- **strconv** for turning types into strings and visa versa
- **text/tabwriter** for writing out tab aligned output
- **os** for opening and reading files
- **github.com/spf13/cobra** for the command line interface
- **github.com/mergestat/timediff** for displaying relative friendly time differences (1 hour ago, 10 minutes ago, etc)
