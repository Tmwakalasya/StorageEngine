package main

func DisplayHelp() {

}

//func CliTools(kv *KeyValue) {
//
//	SETcmd := flag.NewFlagSet("SET", flag.ExitOnError)
//	//GETcmd := flag.NewFlagSet("GET", flag.ExitOnError)
//	//DELETEcmd := flag.NewFlagSet("DELETE", flag.ExitOnError)
//	//EXISTScmd := flag.NewFlagSet("EXISTS", flag.ExitOnError)
//
//	SetKV := SETcmd.String("set", "null", "Set a key and value to the store")
//	if len(os.Args) < 2 {
//		fmt.Println("Error: No command given, available commands(GET,SET,DELETE,EXISTS)")
//	}
//	switch os.Args[1] {
//	case "SET":
//		SETcmd.Parse(os.Args[2:])
//
//	default:
//		fmt.Println("unknown argument", os.Args[1])
//
//	}
//
//
//}
