package 

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/gdenn4/beerus-manager/1.0.0/",
		Index,
	},

	Route{
		"CreateAgent",
		"POST",
		"/gdenn4/beerus-manager/1.0.0/agents",
		CreateAgent,
	},

	Route{
		"DeleteAgent",
		"DELETE",
		"/gdenn4/beerus-manager/1.0.0/agents/{agent_uuid}",
		DeleteAgent,
	},

	Route{
		"GetAgentForAdress",
		"GET",
		"/gdenn4/beerus-manager/1.0.0/agent/{address}",
		GetAgentForAdress,
	},

	Route{
		"ListAgents",
		"GET",
		"/gdenn4/beerus-manager/1.0.0/agents",
		ListAgents,
	},

	Route{
		"ListBrokerAgents",
		"GET",
		"/gdenn4/beerus-manager/1.0.0/brokers/{broker_uuid}/agents",
		ListBrokerAgents,
	},

	Route{
		"CreateAgentBackup",
		"POST",
		"/gdenn4/beerus-manager/1.0.0/brokers/{broker_uuid}/{agent_uuid}/backups",
		CreateAgentBackup,
	},

	Route{
		"CreateBackup",
		"POST",
		"/gdenn4/beerus-manager/1.0.0/backups",
		CreateBackup,
	},

	Route{
		"CreateBrokerBackup",
		"POST",
		"/gdenn4/beerus-manager/1.0.0/brokers/{broker_uuid}",
		CreateBrokerBackup,
	},

	Route{
		"DeleteBackup",
		"DELETE",
		"/gdenn4/beerus-manager/1.0.0/backups/{backup_uuid}",
		DeleteBackup,
	},

	Route{
		"ListAgentBackups",
		"GET",
		"/gdenn4/beerus-manager/1.0.0/agents/{agent_uuid}",
		ListAgentBackups,
	},

	Route{
		"ListBackups",
		"GET",
		"/gdenn4/beerus-manager/1.0.0/backups",
		ListBackups,
	},

	Route{
		"ListBrokerBackups",
		"GET",
		"/gdenn4/beerus-manager/1.0.0/broker/{broker_uuid}/backups",
		ListBrokerBackups,
	},

	Route{
		"ListSpecificAgentBackups",
		"GET",
		"/gdenn4/beerus-manager/1.0.0/brokers/{broker_uuid}/{agent_uuid}/backups",
		ListSpecificAgentBackups,
	},

	Route{
		"BrokerDetails",
		"GET",
		"/gdenn4/beerus-manager/1.0.0/brokers/{broker_uuid}",
		BrokerDetails,
	},

	Route{
		"CreateBroker",
		"POST",
		"/gdenn4/beerus-manager/1.0.0/brokers",
		CreateBroker,
	},

	Route{
		"ListBrokers",
		"GET",
		"/gdenn4/beerus-manager/1.0.0/brokers",
		ListBrokers,
	},

	Route{
		"RemoveBroker",
		"DELETE",
		"/gdenn4/beerus-manager/1.0.0/brokers/{broker_uuid}",
		RemoveBroker,
	},

	Route{
		"CreateUser",
		"POST",
		"/gdenn4/beerus-manager/1.0.0/users",
		CreateUser,
	},

	Route{
		"DeleteUserByID",
		"DELETE",
		"/gdenn4/beerus-manager/1.0.0/users/{user_uuid}",
		DeleteUserByID,
	},

	Route{
		"GetUser",
		"GET",
		"/gdenn4/beerus-manager/1.0.0/users/{user_uuid}",
		GetUser,
	},

	Route{
		"GetUserByName",
		"GET",
		"/gdenn4/beerus-manager/1.0.0/users/username/{username}",
		GetUserByName,
	},

	Route{
		"ListUser",
		"GET",
		"/gdenn4/beerus-manager/1.0.0/users",
		ListUser,
	},

}
