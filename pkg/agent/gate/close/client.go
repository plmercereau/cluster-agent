package close

import (
	"log"
	"net/rpc"

	"github.com/plmercereau/cluster-agent/pkg/agent/config"
)

func CloseRequest() {
	client, err := rpc.DialHTTP("unix", config.SOCKET_PATH)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	args := CloseArgs{}
	err = client.Call("Gate.Close", args, &reply)
	if err != nil {
		log.Fatalln("Error calling service Gate.Close", err)
	}
	log.Println(reply)

}
