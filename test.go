package main

func commit() {
	//save permanently
	serializer()
	//if clustered
	cluster()
}

func parser() {

}

func serializer() {

}

func api() {
	//cheack authorization
	auth()
	//parse request
	parser()
	//if committed save
	commit()
}

func auth() {

}

func config() {
}

func cluster() {

}

func main() {
	//load config
	config()
	//if clustered
	cluster()
	//start api
	api()
}
