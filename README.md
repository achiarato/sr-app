# sr-app: SRv6 SIDs/uSID Policy Engine

HTTP Server able to receive HTTP GET Requests to calculate SIDs/uSID based on specified inputs found in the HTTP filters. The HTTP Server will transform the HTTP Request to a query to ArangoDB that contains the updated topology state of the network and that is capable of path computations based on the network model.

Launching the HTTP Server, it will keep a listening state waiting for requests on specifi port (the port used in the demo is random).

![image](https://user-images.githubusercontent.com/125906326/232450383-a12ed0fb-2d2c-4d1c-9c28-d50152a3f4bd.png)

SR-App Server start listening on :3333

Once HTTP Server receives an HTTP Request, it will query the ArangoDB based on the HTTP Request's inputs: Source Node, Destination Node, Query Type

![image](https://user-images.githubusercontent.com/125906326/232451222-befbbdc5-a6cb-41f1-b815-7c4daee1fec5.png)

In this case, the HTTP Request: http://localhost:3333/shortestpath?src=2_0_0_0000.0000.0001&dst=2_0_0_0000.0000.0013 will generate a query to the ArangoDB to calculate the Shortest Path between the 2_0_0_0000.0000.0001 and 2_0_0_0000.0000.0013 nodes of the topology model.

For more information about how the topology model is created/managed/updated, please refer to: https://github.com/cisco-open/jalapeno/blob/main/README.md

If the HTTP Server receives an incorrect HTTP Request, it will manage the HTTP Response accordingly (Bad Request, Status Code, ...)

![image](https://user-images.githubusercontent.com/125906326/232454235-4397fc75-a9cc-4c54-80b6-5ec0ee1f48de.png)

From the client point of view:

![image](https://user-images.githubusercontent.com/125906326/232454501-4c5fe277-341a-42e9-a637-8ff8b66020a6.png)
