package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func scan() {
	pathValue := flag.String("path", "", "Espeficica o caminho do diretorio")
	programName := flag.String("name", "", "nome do app")
	flag.Parse()
	fmt.Println(*pathValue)
	arr, err := os.ReadDir(*pathValue)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range arr {
		fmt.Println(v)
	}

	template(*programName, *pathValue+*programName, *programName+".svg")

}

func main() {

	scan()
}

func template(name, exec, icon string) {
	str := fmt.Sprintf(`[Desktop Entry]
Name=%s
Exec=%s
Icon=%s
Type=Application
Categories=Application;
`, name, exec, icon)

	appName := name + ".desktop"

	f, err := os.Create(appName)

	if err != nil {
		log.Fatal(err)
		return
	}

	f.WriteString(str)
	log.Println(name + ".desktop" + " Create with successfully")

}
