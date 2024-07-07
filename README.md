# Golang Distributed File System

Welcome to the Golang Distributed File System (GDFS)! This project is a high-performance, distributed file system implemented in Go, designed to provide scalable, fault-tolerant, and efficient storage solutions for modern applications.
Features

- Scalability: Easily scale out by adding more nodes to handle increasing amounts of data and traffic.

- Fault Tolerance: Automatic data replication and recovery mechanisms ensure high availability and durability.
- Performance: Optimized for high throughput and low latency, making it suitable for a variety of workloads.
- Simple API: Provides a straightforward API for file operations, making it easy to integrate into your applications.

### Getting Started
#### Prerequisites

    Go 1.18 or higher
    A Unix-like operating system (Linux or macOS) for development
    Basic understanding of distributed systems concepts

### Installation

##### 1. Clone the repository:

```bash

git clone --depth 1 https://github.com/AnurajBhaskar47/Distributed_File_System.git
cd Distributed_File_System
```

##### 2. Build the project:

```bash
go build -o gdfs ./main.go
```

###### Run the system:

Start the master node:

```bash
./gdfs master
```

Start worker nodes:

```bash
./gdfs worker
```

#### 3. Configuration

Edit the config.yaml file to configure your master and worker nodes. Refer to the Configuration Guide for detailed instructions.

##### Usage

###### Initialize the file system:

```bash
./gdfs init
```

###### Upload a file:

```bash
./gdfs put /path/to/local/file /path/in/distributed/system
```

###### Download a file:

```bash
./gdfs get /path/in/distributed/system /path/to/local/destination
```

###### List files:

```bash
    ./gdfs ls /path/in/distributed/system
```

<!-- ### Documentation

For more detailed documentation, please refer to the following resources:

    [User Guide](https://www.google.com)
    API Documentation
    Configuration Guide

Contributing

We welcome contributions from the community! If you’d like to contribute to the project, please refer to our Contributing Guidelines. -->

### License

This project is licensed under the MIT License.
Contact

For any questions or issues, please open an issue on GitHub or contact me at bhaskar25903@gmail.com.
