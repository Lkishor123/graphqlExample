package main

/*
==========================================================
GRAPHQL Server Struct GQServer
==========================================================
*/
type GQServer struct {
	ServerName string
}

/*
===========================================================
DB Object Struct:
	- DBinfo
	- ComplexType
===========================================================
*/

type DBinfo struct {
	ID      int
	Name    string
	Complex ComplexType
}
type ComplexType struct {
	C1param string
	C2param string
}
