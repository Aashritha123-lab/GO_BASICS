carInventory/

	|--main.go
	|--config/
	|	|--config.go  			 // configuration for databases
	|--models/
	|	|--car.go     			// structure of your model (here basically we are storing the details of car)
	|--handlers/
	|	|--car_handlers.go 		// basically handler is a code where it will take the HTTP request and give the response 
	|--middleware/
	|	|--logging.go
	|	|--security.go			// the layer before the request reaching the server
	|--utils/
	|	|--response.go 				// where it will have json decoding and encoding
	