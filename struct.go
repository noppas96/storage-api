package main

type ClusterMember struct {
	Server []Server
}

type Server struct {
	ID      string
	Address string
}
