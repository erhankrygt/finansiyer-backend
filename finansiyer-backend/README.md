```bash
export MONGO_URI=mongodb+srv://usr_dev:Zrgl9iRx9HLEhumB@mongo-sample-cluster.6kb6ydo.mongodb.net/?retryWrites=true&w=majority
export MONGO_DATABASE=mongo-sample-cluster
export MONGO_CONNECT_TIMEOUT=15s
export MONGO_PING_TIMEOUT=15s
export MONGO_READ_TIMEOUT=15s
export MONGO_WRITE_TIMEOUT=10s
export MONGO_DISCONNECT_TIMEOUT=10s
```

The project sets these variables in configs/env-vars/env-vars.go using the environment package. 
You should access .env file about configuration environment values. Once you've launched the project, 
you can listen on localhost:8081. 

Swagger integration is already implemented, and you can access
the endpoints via Swagger by visiting localhost:8081/docs. Remember to select http instead of https in
the Schemes section.

You will be presented with 4 endpoints:

/commercial-paper: Use this endpoint to add commercial paper belonging to a company. It checks legal
requirements for each new application and adds it. There is a collection in the DB named "Company"
where you can manually add a new company or use an existing one. The ID of the existing
company is f9f85d58-7dd6-40aa-bafd-96e7aa2b762d.

/documents: You will see a list of existing documents, the operation processes, and the companies
they are connected to.

/operation: This endpoint is designed to change the status of a document and execute a deposit operation.
You can give a document an APPROVAL or REJECT status. When the document is first created (/commercial-paper),
it is in REVIEW status. You can access the expected DocumentId parameter via the
DocumentId field in /documents.

/health: The ping endpoint returns a "pong" response and tests whether the system is up and running.

As an example, I've written unit tests that cover 100% of the /commercial-paper endpoint.

Thank you for your interest.