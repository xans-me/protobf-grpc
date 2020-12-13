## Description 

### Common Directory

1. `config` contains shared or global information, which is used by client and server applications.
2. The `models` directory contains `.proto` files.

### Client Directory

It contains one main file, which will be run as a client application. This client application will communicate with 2 server applications.

### Services Directory 

One proto file for one rpc server (service) application. Because there are two proto files, it means that clearly there are two rpc server applications, `service-user` and `service-garage`. This services folder holds the two service applications.