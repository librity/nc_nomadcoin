package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/librity/nc_nomadcoin/blockchain"
)

const port string = ":4000"

func main() {
	welcome()
	blockchainDemo()
	serverDemo()
}

func welcome() {
	fmt.Println("Welcome to Nomad Coin!")
	fmt.Println("---")
}

func blockchainDemo() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("1 BTC from 04382d626ecf46e13e92754d501c4b3f72573cc7e03c377f016483c476986724 to eb6b64f149ac6a8fa99e8b244d7ffc51c3816d3557ac7896752d853c9b9c98df.")

	chain = blockchain.GetBlockchain()
	chain.AddBlock("43 BTC from 0f9bc8e2c5aaa04ae67ab93be52fabe04817e85d7ee8dc6ba68a837cbe81a6d7 to 082aa74e0b446393839d15f354110339561ea3a8d6c2c1f1dc3e8ded23aa34b2.")
	chain.AddBlock("67 BTC from 208776bda9bf749c7b33979ed0fc4335e63e0dc66c3b31e4deb55de3b7e3514a to efd3143378d80b5f0403a719fc4cfc68134e0814745f0d4269e7b02496d560c6.")
	chain.AddBlock("100 BTC from 857a87260f2d78502daa69ca4c7392d104a2f755f4192ce8ff5cf2af3aeac1a6 to 5bcfd4af383b6a935db660365071d84945645f7cbbfb15a1bfad6db065352be5.")

	chain = blockchain.GetBlockchain()
	chain.AddBlock("20 BTC from c65992d074b7de15d22677bf1b6ca6d02bfd47be7b39f017cc4aacc52dd2d40a to 2fa0890f810e5f6529b875857614d6cd2c8419ce573961984a5e4e4acc1fcfd0.")

	chain.ListBlocks()
}

func serverDemo() {
	http.HandleFunc("/", home)

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	chain := blockchain.GetBlockchain()
	data := homeData{"Welcome to Nomad Coin 1.0!", chain.GetAllBlocks()}

	tmpl.Execute(rw, data)
}
