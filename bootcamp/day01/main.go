package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	cakes "src/Cakes"
	diff "src/Diff"
	snapshots "src/Snapshots"
)

var (
	Modes   = [3]string{"LPP", "LCR", "LCS"}
	InFile  = "fpin"
	OldFile = "lcold"
	NewFile = "lcnew"
)

func loadFlags() (map[string]*string, error) {
	res := map[string]*string{}
	res[InFile] = flag.String("f", "", "file to load and pretty-print\ne.g: -f data.json")
	res[OldFile] = flag.String("old", "", "source file to load and compare\ne.g: -old data.json")
	res[NewFile] = flag.String("new", "", "target file to load and compare\ne.g: -new data.json")

	flag.Parse()

	if len(*res[InFile]) > 0 && (len(*res[OldFile]) > 0 || len(*res[NewFile]) > 0) {
		return nil, errors.New("error [input]: can't provide flags -new or -old with -f")
	} else if (len(*res[OldFile]) > 1 && len(*res[NewFile]) > 1) && len(*res[InFile]) > 1 {
		return nil, errors.New("error [input]: can't provide flags -new or -old with -f")
	} else if len(*res[InFile]) == 0 && len(*res[OldFile]) == 0 && len(*res[NewFile]) == 0 {
		return nil, errors.New("error [input]: no enough options\ne.g: -f file.xml\n      -old file.json -new new.xml\n    -old snapshotOld.txt -new snapshotNew.txt")
	}

	txts := []bool{strings.HasSuffix(*res[OldFile], "txt"), strings.HasSuffix(*res[NewFile], "txt")}
	if txts[0] || txts[1] {
		if !txts[0] || !txts[1] {
			return nil, errors.New(
				"error [input]: snapshot files (.txt) can only be compared to other snapshots\n",
			)
		}
	}

	return res, nil
}

func scenarioHandler() error {
	files, err := loadFlags()
	if err != nil {
		return err
	}
	if len(*files[InFile]) > 0 {
		cakes := cakes.Cakes{}
		if err := cakes.Load(*files[InFile]); err != nil {
			return err
		}
		cakes.Print()
		return nil
	}

	if len(*files[OldFile]) > 0 && strings.HasSuffix(*files[OldFile], "txt") {
		oldSnap := snapshots.Snapshot{}
		newSnap := snapshots.Snapshot{}
		if err := oldSnap.Load(*files[OldFile]); err != nil {
			return err
		}
		if err := newSnap.Load(*files[NewFile]); err != nil {
			return err
		}
		res := diff.DiffSnapshots(oldSnap.Items, newSnap.Items)
		for _, line := range res {
			fmt.Println(line)
		}
		return nil
	}

	orgCakes := cakes.Cakes{}
	copyCakes := cakes.Cakes{}
	if err := orgCakes.Load(*files[OldFile]); err != nil {
		return err
	}
	if err := copyCakes.Load(*files[NewFile]); err != nil {
		return err
	}

	res := diff.DiffCakes(&orgCakes.Items, &copyCakes.Items)

	for _, line := range res {
		fmt.Println(line)
	}

	return nil
}

func main() {
	if err := scenarioHandler(); err != nil {
		fmt.Printf("FATAL: %v\n", err)
		os.Exit(1)
	}
}
