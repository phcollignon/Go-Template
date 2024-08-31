package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/Masterminds/sprig"
	"github.com/ghodss/yaml"
)

type MultipleFileData struct {
	Files []struct {
		FileName string
		Data     interface{}
	}
}

func main() {

	dataFileName := flag.String("d", "", "json or yaml data file")
	templateFileName := flag.String("t", "", "go template file")
	outputDirectory := flag.String("o", ".", "output directory")
	multipleFiles := flag.String("m", "", "-m multi : generates one file for each File object in Json (or Yaml) data file")
	flag.Parse()

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("    gocodegen -d data-file (Json or Yaml) -t template-file [-o output-directory] [-m multiple-files] \n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *dataFileName == "" || *templateFileName == "" {
		flag.Usage()
	}

	hasMultipleFiles := (*multipleFiles != "") && (*multipleFiles == "multi")

	template, err := template.New(path.Base(*templateFileName)).Funcs(template.FuncMap(MyFuncMap)).Funcs(sprig.TxtFuncMap()).ParseFiles(*templateFileName)
	if err != nil {
		fmt.Println(err)
	}

	dataFile, err := os.Open(*dataFileName)
	if err != nil {
		fmt.Println(err)
	}

	defer dataFile.Close()

	byteValue, err := ioutil.ReadAll(dataFile)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, err = ToJSON(byteValue)

	if hasMultipleFiles {
		var multidata MultipleFileData

		err = json.Unmarshal(byteValue, &multidata)
		if err != nil {
			fmt.Println(err)
		}

		for _, mfile := range multidata.Files {
			outputFileName := path.Join(path.Dir(*dataFileName), mfile.FileName)
			mdata := mfile.Data
			generateFile(template, *outputDirectory, outputFileName, mdata)
		}

	} else {
		var data interface{}
		err = json.Unmarshal(byteValue, &data)
		if err != nil {
			fmt.Println(err)
		}
		outputFileName := strings.TrimSuffix(filepath.Base(*dataFileName), filepath.Ext(*dataFileName)) + ".generated.txt"
		generateFile(template, *outputDirectory, outputFileName, data)
	}
}

func generateFile(template *template.Template, outputDirectory string, outputFileName string, data interface{}) {
	absOutputFileName := path.Join(outputDirectory, outputFileName)
	os.MkdirAll(path.Dir(absOutputFileName), os.ModePerm)
	outputFile, err := os.Create(absOutputFileName)
	fmt.Println("Generating file : " + absOutputFileName)
	defer outputFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = template.Execute(outputFile, data)
	if err != nil {
		fmt.Println(err)
	}
}

func ToJSON(data []byte) ([]byte, error) {
	if hasJSONPrefix(data) {
		return data, nil
	}
	return yaml.YAMLToJSON(data)
}

var jsonPrefix = []byte("{")

func hasJSONPrefix(buf []byte) bool {
	trim := bytes.TrimLeftFunc(buf, unicode.IsSpace)
	return bytes.HasPrefix(trim, jsonPrefix)
}
