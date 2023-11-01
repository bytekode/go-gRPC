package common

const G_RPC_PORT = "5001"

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
