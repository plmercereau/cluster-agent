package open

import (
	"log"
	"net/rpc"

	"github.com/plmercereau/cluster-agent/pkg/config"
)

func OpenRequest() {
	// TODO close the cluster node so no new node can join anymore
	client, err := rpc.DialHTTP("unix", config.SOCKET_PATH)
	if err != nil {
		log.Fatalln("dialing:", err)
	}

	var reply string
	args := OpenArgs{}
	err = client.Call("Gate.Open", args, &reply)
	if err != nil {
		log.Fatalln("Error calling service Gate.Open", err)
	}
	log.Println(reply)

}
