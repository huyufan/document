package main

import (
	"bytes"
	"context"
	"exec/go/exec/gostorages"
	"exec/go/exec/gostorages/fs"
	"os"
)

func main() {
	dir := os.TempDir()
	store := fs.NewStorage(fs.Config{Root: dir})

	s := gostorages.NewNoop(store)

	ctx := context.Background()

	s.Save(ctx, bytes.NewReader([]byte("woehowehhwe")), "cccc")
	//store.Delete(ctx, "aaaa")
}
