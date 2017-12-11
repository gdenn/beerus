package manager

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Route at backup manager
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes registered at backup manager
type Routes []Route

// NewRouter creates an instance of mux.Router and registers
// all routes (see swagger doc in /api)
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = AddContext(Logger(handler, route.Name))

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// Index the test route to validate your manager is running
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/beerus-manager/1.0.0/",
		Index,
	},

	Route{
		"CreateAgent",
		"POST",
		"/beerus-manager/1.0.0/agents",
		CreateAgent,
	},

	Route{
		"DeleteAgent",
		"DELETE",
		"/beerus-manager/1.0.0/agents/{agent_uuid}",
		DeleteAgent,
	},

	Route{
		"GetAgentForAdress",
		"GET",
		"/beerus-manager/1.0.0/agent/{address}",
		GetAgentForAdress,
	},

	Route{
		"ListAgents",
		"GET",
		"/beerus-manager/1.0.0/agents",
		ListAgents,
	},

	Route{
		"ListBrokerAgents",
		"GET",
		"/beerus-manager/1.0.0/brokers/{broker_uuid}/agents",
		ListBrokerAgents,
	},

	Route{
		"CreateAgentBackup",
		"POST",
		"/beerus-manager/1.0.0/brokers/{broker_uuid}/{agent_uuid}/backups",
		CreateAgentBackup,
	},

	Route{
		"CreateBackup",
		"POST",
		"/beerus-manager/1.0.0/backups",
		CreateBackup,
	},

	Route{
		"CreateBrokerBackup",
		"POST",
		"/beerus-manager/1.0.0/brokers/{broker_uuid}",
		CreateBrokerBackup,
	},

	Route{
		"DeleteBackup",
		"DELETE",
		"/beerus-manager/1.0.0/backups/{backup_uuid}",
		DeleteBackup,
	},

	Route{
		"ListAgentBackups",
		"GET",
		"/beerus-manager/1.0.0/agents/{agent_uuid}",
		ListAgentBackups,
	},

	Route{
		"ListBackups",
		"GET",
		"/beerus-manager/1.0.0/backups",
		ListBackups,
	},

	Route{
		"ListBrokerBackups",
		"GET",
		"/beerus-manager/1.0.0/broker/{broker_uuid}/backups",
		ListBrokerBackups,
	},

	Route{
		"ListSpecificAgentBackups",
		"GET",
		"/beerus-manager/1.0.0/brokers/{broker_uuid}/{agent_uuid}/backups",
		ListSpecificAgentBackups,
	},

	Route{
		"BrokerDetails",
		"GET",
		"/beerus-manager/1.0.0/brokers/{broker_uuid}",
		BrokerDetails,
	},

	Route{
		"CreateBroker",
		"POST",
		"/beerus-manager/1.0.0/brokers",
		CreateBroker,
	},

	Route{
		"ListBrokers",
		"GET",
		"/beerus-manager/1.0.0/brokers",
		ListBrokers,
	},

	Route{
		"RemoveBroker",
		"DELETE",
		"/beerus-manager/1.0.0/brokers/{broker_uuid}",
		RemoveBroker,
	},

	Route{
		"CreateUser",
		"POST",
		"/beerus-manager/1.0.0/users",
		AddUserTable(CreateUser),
	},

	Route{
		"DeleteUserByID",
		"DELETE",
		"/beerus-manager/1.0.0/users/{user_uuid}",
		AddUserTable(DeleteUserByID),
	},

	Route{
		"GetUser",
		"GET",
		"/beerus-manager/1.0.0/users/{user_uuid}",
		AddUserTable(GetUser),
	},

	Route{
		"GetUserByName",
		"GET",
		"/beerus-manager/1.0.0/users/username/{username}",
		AddUserTable(GetUserByName),
	},

	Route{
		"ListUser",
		"GET",
		"/beerus-manager/1.0.0/users",
		AddUserTable(ListUser),
	},
}
