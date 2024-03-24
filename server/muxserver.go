package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/luycaslima/virtual-pets-server/repositories"
	"github.com/luycaslima/virtual-pets-server/routes"
	"github.com/luycaslima/virtual-pets-server/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

type muxServer struct {
	router *mux.Router
	db     *mongo.Client
}

func NewMuxServer(db *mongo.Client) Server {
	return &muxServer{
		router: mux.NewRouter(),
		db:     db,
	}
}

func (s *muxServer) StartServer(port string) error {
	//TODO set a config file for this that need to be passed as argument
	headers := handlers.AllowedHeaders([]string{"Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With"})
	exposedHeader := handlers.ExposedHeaders([]string{"Origin"})
	methods := handlers.AllowedMethods([]string{"GET,POST,PUT,DELETE,OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"}) //TODO put this on production virtualpets.vercel.app but and the localhost?
	//maxAge := handlers.MaxAge(12)
	//exposedHeaders := handlers.ExposedHeaders([]string{"Content-Length"})

	s.initializeHandlers()

	fmt.Println("Starting Server at port:", port)
	return http.ListenAndServe(port, handlers.CORS(headers, exposedHeader, methods, origins)(s.router))
}

func (s *muxServer) initializeHandlers() {
	userRepository := repositories.NewUserMongoDbRepository(s.db)
	specieRepository := repositories.NewSpecieMongoDbRepository(s.db)
	petRepository := repositories.NewPetMongoDbRepository(s.db)

	userUseCase := usecase.NewUserService(&userRepository, &specieRepository, &petRepository)
	specieUseCase := usecase.NewSpecieService(&specieRepository)

	routes.UserRoutes(s.router, userUseCase)
	routes.SpecieRoutes(s.router, specieUseCase)
}
