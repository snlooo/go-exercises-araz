package xml

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
)

//Create a program that generates an XML configuration file named config.xml with fields like
//database, username, password, and a nested options field that includes auto_backup and max_connections.
//Write the XML structure to config.xml and ensure it's formatted in a readable way.

type Config struct {
	Database string `xml:"database"`
	Username string `xml:"username"`
	Password string `xml:"password"`
	Option   Option `xml:"options"`
}

type Option struct {
	AutoBackup     bool `xml:"auto_backup"`
	MaxConnections int  `xml:"max_connections"`
}

func ReadConfigXml(db, username, password string, autoBackup bool, maxConnection int) error {

	fmt.Println("\n#================================#")

	f, err := os.OpenFile("./xml/config.xml", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	config := Config{
		Database: db,
		Username: username,
		Password: password,
		Option: Option{
			AutoBackup:     autoBackup,
			MaxConnections: maxConnection,
		},
	}

	output, _ := xml.Marshal(config)
	newWriter := bufio.NewWriter(f)
	_, writeError := newWriter.WriteString(string(output))
	if writeError != nil {
		return err
	}
	flushError := newWriter.Flush()
	if flushError != nil {
		return flushError
	}

	fmt.Println("Successfully wrote config to config.xml")

	return nil
}
