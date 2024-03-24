package server

type Server interface {
	StartServer(port string) error
}
