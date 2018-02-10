package main

import (
	"os"
	"log"
	"fmt"
	"flag"
	"strings"
	"encoding/json"

	"github.com/IBM-Bluemix/bluemix-go/session"
	"github.com/IBM-Bluemix/bluemix-go/trace"
	"github.com/IBM-Bluemix/bluemix-go/api/mccp/mccpv2"

	credreader "github.com/silveraid/ibmcloud_blockchain_info/pkg/credreader"
)

func main() {

	//var org string
	//flag.StringVar(&org, "org", "", "Bluemix Organization")
	//
	//var space string
	//flag.StringVar(&space, "space", "", "Bluemix Space")

	var instanceName string
	flag.StringVar(&instanceName, "name", "", "Bluemix Instance Name")

	flag.Parse()

	//if org == "" || space == "" || instanceName == "" {
	//	flag.Usage()
	//	os.Exit(1)
	//}

	if instanceName == "" {
		flag.Usage()
		os.Exit(1)
	}

	// for troubleshooting
	trace.Logger = trace.NewLogger("false")

	// new session
	s, err := session.New()

	if err != nil {
		log.Fatal(err)
	}

	// new client
	client, err := mccpv2.New(s)

	if err != nil {
		log.Fatal(err)
	}

	//region := s.Config.Region
	//
	//orgAPI := client.Organizations()
	//myorg, err := orgAPI.FindByName(org, region)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//_ = myorg
	//spaceAPI := client.Spaces()
	//myspace, err := spaceAPI.FindByNameInOrg(myorg.GUID, space, region)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	serviceInstanceAPI := client.ServiceInstances()

	myService, err := serviceInstanceAPI.FindByName(instanceName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service instance found ...")

	// creating new service key to retrieve data
	serviceKeys := client.ServiceKeys()

	// looking for a previous key
	mykeys, err := serviceKeys.FindByName(myService.GUID, "BlockchainInfo")

	if err != nil {

		// if the key only missing, let's create one
		if strings.Contains(err.Error(), "erviceKeyDoesNotExist") {

			//
			fmt.Println("Key does not exist, creating a new one ...")
			// creating a new service key
			_, err = serviceKeys.Create(myService.GUID, "BlockchainInfo", nil)

			if err != nil {
				log.Fatal(err)
			}

			//
			fmt.Println("Retrieving the new key ...")

			// find the key again
			mykeys, err = serviceKeys.FindByName(myService.GUID, "BlockchainInfo")

			if err != nil {
				log.Fatal(err)
			}

		} else {
			log.Fatal(err)
		}
	}

	// Printing out info
	fmt.Println("Key name:", mykeys.Name)

	// Retrieving the credential data
	credentials := mykeys.Credentials["credentials"].(map[string]interface{})

	// Parsing the credentials
	c := credreader.CredReader(credentials)

	// Printing out in JSON for validation
	jsonData, _ := json.Marshal(c)
	fmt.Println(string(jsonData))
}
