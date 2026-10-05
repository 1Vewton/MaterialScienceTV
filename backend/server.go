package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/1Vewton/MaterialScienceTV/backend/graph"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/database"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/database/redismanager"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/middleware"
	"github.com/1Vewton/MaterialScienceTV/backend/internal/user"
	"github.com/1Vewton/MaterialScienceTV/backend/pkg/config"
	"github.com/1Vewton/MaterialScienceTV/backend/pkg/logger"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

var defaultPort = config.Config.GetServerPort()

var mainLogger = logger.NewLogger(
	"Application",
	nil,
)

func init() {
	resDB, err := database.InitDataBase(
		config.Config.GetDatabaseType(),
		config.Config.GetDatabaseURL(),
		&user.User{},
	)
	database.DataBase = resDB
	if err != nil {
		panic(err)
	}
	dialTimeOut, err := config.Config.GetRedisDialTimeout()
	if err != nil {
		panic(err)
	}
	readTimeOut, err := config.Config.GetRedisReadTimeout()
	if err != nil {
		panic(err)
	}
	writeTimeOut, err := config.Config.GetRedisWriteTimeout()
	if err != nil {
		panic(err)
	}
	maxRetries, err := config.Config.GetRedisMaxRetries()
	if err != nil {
		panic(err)
	}
	maxRetryBackoff, err := config.Config.GetRedisMaxRetryBackOff()
	if err != nil {
		panic(err)
	}
	minRetryBackoff, err := config.Config.GetRedisMinRetryBackOff()
	if err != nil {
		panic(err)
	}
	redismanager.RedisClient = redismanager.NewRedisConfig(
		config.Config.GetRedisURL(),
		config.Config.GetRedisPassword(),
	).WithDialTimeout(
		dialTimeOut,
	).WithWriteTimeout(
		writeTimeOut,
	).WithReadTimeout(
		readTimeOut,
	).WithMaxRetryBackoff(
		maxRetryBackoff,
	).WithMinRetryBackoff(
		minRetryBackoff,
	).WithMaxRetries(
		maxRetries,
	).ToClient()
	mainLogger.Info(
		fmt.Sprintf(
			"successfully connected to %s",
			config.Config.GetRedisURL(),
		),
	)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// Add middleware here
	serverHandler := playground.Handler("GraphQL playground", "/query")
	serverHandler = middleware.SetCookieMiddleware(
		serverHandler,
	)

	http.Handle("/", serverHandler)
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	runServer := func() {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
	go runServer()
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	err := database.CloseDatabase(
		database.DataBase,
	)
	if err != nil {
		mainLogger.Error(
			err.Error(),
		)
		panic(err)
	}
	err = redismanager.Close(
		redismanager.RedisClient,
	)
	if err != nil {
		mainLogger.Error(
			err.Error(),
		)
		panic(err)
	}
}
