package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const assetTemplate = AssetObject("H4sIAAAAAAAA/3SQsU7DMBCGZ99THJlsKWoXxIDUoUA3BBJ9Aje5RIbEF50vQgX13ZHjqOrCZv3+fd/nm3zz5XvC0YcIEMaJRdGCqU5npVSBqRoeJ6GUtv1PmHJAseE2xH578oke7nMUeBt41jBU4AD0PBHuUyJ9P31So5hUQuwBujk2aD3fXjp8Zd/u03HpWIe2tGskERaHv2Bar77GzD+I4ONuOW7e6PuDfEtii0gOXqjhm+So7WG1rXH5US49zV1HsgILzXp2zjkwobty7nYYw5D5RkhniVhVVwswFzAtdSSY7TbPAyeyDgzPOs262GfTspZNFt0Pg83dQqF/CbROX6NVsMx1dX4BF/gLAAD//zlu92K4AQAA")

func main() {
	name := flag.String("name", "MyAssetName", "name of the asset")
	flag.Parse()
	var data bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &data)
	comp := gzip.NewWriter(encoder)
	io.Copy(comp, os.Stdin)
	comp.Close()
	encoder.Close()
	template, _ := assetTemplate.LoadAsString()
	fmt.Printf(template+"\nconst %s = AssetObject(\"%s\")\n", strings.Title(*name), data.String())
}
