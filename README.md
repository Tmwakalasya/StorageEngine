```markdown
# Key-Value Store

This project is a key-value store implemented in Go. It supports basic operations such as `SET`, `GET`, `DELETE`, and `EXISTS`, and includes functionality for logging operations and rebuilding the store from a log file. It is a distributed key-value store in its initial stages, with the goal of becoming a production-ready key-value storage solution.
## Features

- **SET**: Add or update a key-value pair.
- **GET**: Retrieve the value associated with a key.
- **DELETE**: Remove a key-value pair.
- **EXISTS**: Check if a key exists.
- **Logging**: Log all operations to a file.
- **Rebuild**: Rebuild the store from a log file.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/Tmwakalasya/StorageEngine/tree/main
   cd your-repo-name
   ```

2. Build the project:
   ```sh
   go build
   ```

### Usage

1. Run the application:
   ```sh
   ./StorageEngine
   ```

2. The application will perform some predefined operations and print the results.

### Code Structure

- `main.go`: Entry point of the application.
- `storage.go`: Contains the implementation of the key-value store and related functions.
- `logs.txt`: Log file for storing operations.
- `qodana.yaml`: Configuration file for Qodana analysis.

### Example

```go

store := NewKeyValueStorage()
store.Set("Customer 1", "$547.45")
fmt.Println(store.Get("Customer 1"))
store.Delete("Customer 1")
fmt.Println(store.Exists("Customer 1"))
store.ReBuildStore("logs.txt")
```
## Current Features Being Built
Rebuild Method: Rebuild the key-value store from a log file.
Replication Logic: Implementing replication to ensure data consistency and availability.

## Future Goals
Master-Slave Replication: Implement a master-slave replication where the master node handles all write operations and propagates changes to one or more slave nodes.
Leader Election: Implement a leader election mechanism to handle failover scenarios where a new master is elected if the current master fails.
Consistency Mechanisms: Ensure data consistency across all nodes by implementing mechanisms like write-ahead logging or two-phase commit.
Periodic Backups: Implement a feature to periodically back up the key-value store to a file or a remote storage service.
Restore from Backup: Implement a feature to restore the key-value store from a backup file.
Sharding: Implement sharding to distribute data across multiple nodes to improve scalability and performance.
Health Checks: Implement health checks to monitor the status of each node in the cluster.
Metrics Collection: Collect and expose metrics such as read/write latency, error rates, and replication lag.
Authentication and Authorization: Implement authentication and authorization mechanisms to control access to the key-value store.
Encryption: Implement encryption for data at rest and in transit to ensure data security.
Advanced Querying: Implement support for range queries and secondary indexes to support efficient querying on non-primary key attributes.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
```

