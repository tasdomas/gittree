calculate sha1 hash of all versions of a file
=============================================

This is an experiment in using the `go-git` package to
calculate the sha1 sums of all versions of a file in a git repo.

To test this run:
```
$ go run main.go ./ a.txt
```

The first parameter is the location of the git repository.
The second paramter is the name of a file to calculate sums for (only
top level files supported at the moment).
