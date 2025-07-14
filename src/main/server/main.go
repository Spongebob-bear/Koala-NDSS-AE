/*
Main function for servers. Start A replica.
*/

package server

import (
	"flag"
	"log"
	"os"
	"sleepy-hotstuff/src/communication/receiver"
	"sleepy-hotstuff/src/db"
)

const (
	helpText_server = `
Main function for servers. Start A replica. 
server [ReplicaID] [Membership]
[Membership] can only be 1 or 2 if membership requests are needed. 
1 stands for join request and 2 stands for leave request. 
`
)

func main() {
	log.SetFlags(log.Ltime)
	helpPtr := flag.Bool("help", false, helpText_server)
	flag.Parse()

	if *helpPtr || len(os.Args) < 2 {
		log.Printf(helpText_server)
		return
	}

	id := "0"
	if len(os.Args) > 1 {
		id = os.Args[1]
	}

	mem := "0"

	if len(os.Args) > 2 {
		mem = os.Args[2]
	}

	if mem != "0" && mem != "1" && mem != "2" {
		log.Printf(helpText_server)
		return
	}
	log.Printf("**Starting replica %s", id)

	err := db.StartDB(id)
	if err != nil {
		log.Println("error starting the local database:%v", err)
		return
	}
	log.Println("the local database has started")
	defer db.CloseDB()

	receiver.StartReceiver(id, true, mem)

}
