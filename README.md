# IBM Blockchain Credential information parser

Converts the structured but untyped data from the IBM API into strcts with proper types.

Instructions to run the demo

### 1. Install the SDK using the following command

```bash
go get github.com/IBM-Bluemix/bluemix-go
```

### 2. Export the following environment variables

* IBMID - This is the IBM ID
* IBMID_PASSWORD - This is the password for the above ID

or 

* BM_API_KEY/BLUEMIX_API_KEY - This is the Bluemix API Key

### 3. Run the app with the -name parameter

```
go run main.go -name your_blockchain_service_name
```
