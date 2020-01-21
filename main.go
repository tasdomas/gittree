package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	gogit "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/filemode"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func main() {
	repo, err := gogit.PlainOpen(os.Args[1])
	if err != nil {
		panic(err)
	}

	fname := &os.Args[2]

	logOpts := gogit.LogOptions{
		Order:    gogit.LogOrderCommitterTime,
		FileName: fname,
	}
	commits, err := repo.Log(&logOpts)
	if err != nil {
		panic(err)
	}
	defer commits.Close()

	err = commits.ForEach(func(c *object.Commit) error {
		tree, err := repo.TreeObject(c.TreeHash)
		if err != nil {
			return err
		}
		for _, e := range tree.Entries {
			if e.Mode == filemode.Regular || e.Mode == filemode.Executable {
				if e.Name == *fname {
					// Found our file.
					f, err := tree.File(e.Name)
					if err != nil {
						return err
					}
					h := sha1.New()
					r, err := f.Reader()
					if err != nil {
						return err
					}
					_, err = io.Copy(h, r)
					if err != nil {
						return err
					}
					r.Close()
					fmt.Printf("%s %s\n", c.Hash, hex.EncodeToString(h.Sum(nil)))
				}
			}
		}
		return nil
	})
	if err != nil && err != io.EOF {
		panic(err)
	}
}
