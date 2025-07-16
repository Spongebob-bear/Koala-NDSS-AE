# Koala-NDSS-AE
*The artifact for AE of NDSS'26*

This repository contains artifacts for the paper "Consensus in the Known Participation Model with Byzantine Failures and Sleepy Replicas". The code is primarily intended for academic research, which include:
- HotStuff, a classic partially synchronous consensus protocol.
- Koala-2, the first partially synchronous consensus protocol that tolerates Byzantine failures and sleepy replicas without the stable storage assumption.

This document serves as the Artifact Evaluation guide for our paper. It contains detailed instructions to set up the environment, compile the source code, and, most importantly, reproduce the experimental results presented in the paper. All steps have been tested on an Ubuntu Server 22.04 LTS.


## System Requirements

A machine with the following minimum specifications is required:

- Operating System: Ubuntu Server 22.04 LTS (amd64)
- CPU: 4 cores
- Memory: 16 GB
- Disk: 40 GB of free space (SSD recommended)


## Prerequisites

Before you begin, you need to update your system's package list and install some essential tools like `git`, `wget`, and `python3-pip`.

Open your terminal and run the following commands:

```bash
# Update package list
sudo apt update -y

# Install essential build and utility tools
sudo apt install -y python3-pip git wget zip unzip vim curl
```

## Installation

The installation process is divided into two main steps: installing the Go programming language and then building the project itself.

### Step 1: Install Go (GoLang)

This project requires Go version `1.19.4`.

1.  **Download the Go binary archive.**
    You can download it from the official Go website. If you are in a region with slow access (like mainland China), a mirror is provided.

    *   **Official Link:**
        ```bash
        wget https://go.dev/dl/go1.19.4.linux-amd64.tar.gz
        ```
    *   **Aliyun Mirror (for users in China):**
        ```bash
        wget https://mirrors.aliyun.com/golang/go1.19.4.linux-amd64.tar.gz
        ```

2.  **Extract the archive.**
    This command will install Go into the `/usr/local` directory, which is the standard location.

    ```bash
    sudo tar -C /usr/local -xzf go1.19.4.linux-amd64.tar.gz
    ```

3.  **Add Go to your system's PATH.**
    This is a **critical step** so that your terminal can find the `go` command.

    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    source ~/.profile
    ```
    To verify the installation, open a **new terminal** or run `source ~/.profile` again, then type:
    ```bash
    go version
    # It should output: go version go1.19.4 linux/amd64
    ```

### Step 2: Clone and Build the Project

Now that Go is installed, you can clone the repository and compile the source code.

1.  **Clone the repository.**
    ```bash
    git clone https://github.com/Spongebob-bear/Koala-NDSS-AE.git
    ```

2.  **Navigate into the project directory.**
    ```bash
    cd Koala-NDSS-AE
    ```

3.  **Switch to the stable version.**
    ```bash
    git checkout ndss-ae
    ```

4.  **Run the build script.**
    This repository provides two build scripts. Choose one based on your needs.

    *   **Option A: Online Build**
        This script will automatically download all required dependencies from the internet before compiling.
        ```bash
        ./scripts/build.sh
        ```

    *   **Option B: Offline Build**
        This script uses the dependencies already included in the `vendor/` directory within the repository. It's faster and does not require an internet connection for dependencies, making it ideal for reproducible builds or environments with network restrictions.
        ```bash
        ./scripts/build_offline.sh
        ```

We **recommend** using the offline build to ensure a smooth and reliable compilation process.

After a successful build, you will find the executables (`server`, `client`, `ecdsagen`) in the root directory of the project.

## Usage


After a successful build, you can run the compiled binaries from the project's root directory.

### Running the Server

The server instances are configured in the `etc/conf.json` file. Before starting a server, you must define its ID, host, and port in this file.

1.  **Configure Server Replicas**

    Open `etc/conf.json` and edit the `replicas` array. Each object in the array represents a server instance that can be launched.

    ```json
    "replicas": [
       {
          "id": "0",          // The unique ID for this server instance
          "host": "127.0.0.1", // The IP address the server will listen on
          "port": "8000"       // The port number
       },
       {
          "id": "1",
          "host": "127.0.0.1",
          "port": "8001"
       }
    ]
    ```

2.  **Launch a Server Instance**

    To start a server, use the following command, where `[id]` must match one of the IDs you configured in `etc/conf.json`.

    ```bash
    ./server [id]
    ```

    **Example:**
    This command starts the server instance with `id` 0.
    ```bash
    ./server 0
    ```

### Running the Client

The client can be used to send requests to the running servers.

1.  **Command Syntax**

    Use the following format to run the client:

    ```bash
    ./client [client-id] [operation-type] [batch-size]
    ```

2.  **Parameters**

    | Parameter          | Type/Value             | Description                                                                                               |
    | ------------------ | ---------------------- | --------------------------------------------------------------------------------------------------------- |
    | `[client-id]`      | Integer                | A unique positive integer to identify this client. It must not conflict with any of the server replica IDs. |
    | `[operation-type]` | `0` or `1`             | Specifies the type of operation to perform: <br> • `0`: Single write operation. <br> • `1`: Batched write operation. |
    | `[batch-size]`     | Integer                | The number of operations to execute. For a single write, this is typically `1`. For a batch, it's the batch size. |

3.  **Examples**

    *   **To perform a single write operation:**
        This command starts a client with ID `100` and sends one write request.
        ```bash
        ./client 100 0 1
        ```

    *   **To perform a batched write operation:**
        This command starts a client with ID `200` and sends a batch of 500 write requests.
        ```bash
        ./client 200 1 500
        ```


## Evaluation


This section details the steps to reproduce the experiments presented in our paper. We provide scripts to automate the evaluation process. All commands should be executed from the project's root directory. We provide scaled-down experiments and implement evaluation on a single machine.

The evaluation consists of two main experiments, each designed to validate one or multiple key claims from our paper.

### Experiment 1: Performance with Different Storage Options

This experiment aims to reproduce the results shown in **Figure 4.a, 4.b and 4.c** of the paper, which evaluates the performance of HotStuff under different storage options.

This experiment runs 4 server replicas and 1 client process on the local machine. Detailed logs for each replica are stored in a `var` directory, which should be located at the same level as the project's root directory (`Koala-NDSS-AE`). Specifically, the output log for replica N can be found in a file `var/log/N/date_Eva.log`.

#### Experiment 1.1: All Parameters in Stable Storage

This sub-experiment measures the latency and throughput of HotStuff when all consensus parameters are stored in stable storage.

**Instructions**

Run the following script from the project's root directory:

```bash
./scripts/run_experiment_1.sh 1 50
```

*   The first argument `1` selects this specific sub-experiment (all parameters in stable storage).
*   The second argument `50` specifies the duration (in seconds) the script waits for the experiment to complete.

**Expected Outcome**

The script will first display messages indicating it is setting up the correct configuration, starting the 4 server processes, and then launching the client. You will see output similar to this as the system runs:

```
Evaluating performance when storing all consensus parameters

[Configuration] modify the configuration file
'./experiments/HS_storage/all/conf.json' -> './etc/conf.json'

[Start Server] start 4 servers
...
[Start Client] start the client.
...
Evaluation in progress... waiting 50 seconds.
...
[Replica] Processed 55000 (ops=1000, clockTime=37448 ms, seq=55) operations using 1230 ms. Throughput 813 tx/s. 
...
[Kill Processes] kill all server and client
```

After the specified wait time, the script will terminate all processes and calculate the average performance metrics. The key output to look for is the **final line**, which reports the overall throughput and latency:

```
[Output] Print the performance of the sleepy replica
throughput(tps):2027.004716981132, latency(ms):2114.544117647059
```

> **Note:** The exact numbers for throughput and latency may vary slightly depending on your system's hardware and current load. If the performance line of the block of `seq=55` is not printed, it may be because the experiment did not have enough time to complete. In that case, please try re-running the script with a longer wait time (e.g., `./scripts/run_experiment_1.sh 1 70`).




#### Experiment 1.2: Minimum Parameters in Stable Storage

This sub-experiment evaluates performance when only the minimum consensus parameters are stored in stable storage.

**Instructions**

Run the following command from the project's root directory:

```bash
./scripts/run_experiment_1.sh 2 30
```

*   The first argument `2` selects this sub-experiment (minimum parameters in stable storage).
*   The second argument `30` is the wait time in seconds.

**Expected Outcome**


```
Evaluating performance when storing minimum consensus parameters

[Configuration] modify the configuration file
'./experiments/HS_storage/minimum/conf.json' -> './etc/conf.json'

[Start Server] start 4 servers
...
[Start Client] start the client.
...
Evaluation in progress... waiting 30 seconds.
...
13:11:27 [Replica] Processed 2000 (ops=1000, clockTime=54 ms, seq=2) operations using 54 ms. Throughput 18518 tx/s. 
...
13:11:27 [!!!] Ready to output a value for height 10
...
13:11:37 [Replica] Processed 55000 (ops=1000, clockTime=10284 ms, seq=55) operations using 409 ms. Throughput 2444 tx/s.
...
[Kill Processes] kill all server and client
```

After the experiment concludes, the script will compute and display the average performance results. The key output is the **last line**:

```
[Output] Print the performance of the sleepy replica
throughput(tps):8046.622641509434, latency(ms):575.2696078431372
```


#### Experiment 1.3: No Parameters in Stable Storage

This sub-experiment evaluates performance when no consensus parameters are stored in stable storage.

**Instructions**

Run the following command from the project's root directory:

```bash
./scripts/run_experiment_1.sh 3 30
```

*   The first argument `3` selects this sub-experiment (no parameters in stable storage).
*   The second argument `30` is the wait time in seconds.

**Expected Outcome**


```
Evaluating performance when storing none of the consensus parameters

[Configuration] modify the configuration file
'./experiments/HS_storage/none/conf.json' -> './etc/conf.json'

[Start Server] start 4 servers
...
[Start Client] start the client.
...
Evaluation in progress... waiting 30 seconds.
...
13:11:37 [Replica] Processed 55000 (ops=1000, clockTime=10284 ms, seq=55) operations using 56 ms. Throughput 17857 tx/s.
...
[Kill Processes] kill all server and client
```

After the experiment concludes, the script will compute and display the average performance results. The key output is the **last line**:

```
[Output] Print the performance of the sleepy replica
throughput(tps):17784.372641509435, latency(ms):168.97549019607843
```

### Experiment 2: Double-Spending Attack

This experiment implement a demo of a decentralized payment system to test the resilience of our protocols against double-spending attacks. We configure the consensus layer using three different protocols: HotStuff without stable storage, HotStuff with minimum consensus parameters stored (i.e., HotStuff-mSS), and
Koala-2. 


#### Experiment 2.1: Attack on HotStuff without Stable Storage

This sub-experiment aims to reproduce the result shown in **Figure 5** of the paper, showing that a standard HotStuff implementation that stores no consensus parameters in stable storage is vulnerable to a double-spending attack.

**Instructions**

Run the following script from the project's root directory:

```bash
./scripts/run_experiment_2_1.sh
```

This script will:
1.  Configure the system to use HotStuff without stable storage.
2.  Start 4 replicas. The output of each replica is redirected to a separate log file for analysis (e.g., `experiments/double_spending/HS-nSS/output/server_0.log`).
3.  Simulate a client sending two conflicting transactions (from account `0` to `1`, and from account `0` to `2`).
4.  Cause one replica (replica 2) to "sleep" and "wake up" to simulate a restart.

**Expected Outcome**

The script will first show the setup and client transaction submissions. The crucial part of the output is the log from the sleepy replica (replica 2), which will be printed to the console.

You should observe the following sequence of events, which confirms a **successful double-spending attack**:

1.  **First, the replica commits the first transaction (`0 -> 1`) within a block of height 1 before going to sleep.**
    Look for the output showing the block containing the transaction `{"From":"0","To":"1","Value":40}` followed by the "Falling asleep" message.

    ```
    13:39:41 [!!!] Ready to output a value for height 2
    ...
    13:39:41 {"View":0,"Height":1,"TXS":[...,{"ID":100,"TX":{"From":"0","To":"1","Value":40"},,"Timestamp":1752586781315}]}
    ...
    13:39:41 Falling asleep in sequence 5...
    ```

2.  **After restarting, the replica forgets its previous state and commits a conflicting transaction (`0 -> 2`) within a block of height 1.**
    Look for the "Wake up" message, followed by a new block being committed that contains the transaction `{"From":"0","To":"2","Value":40}`.

    ```
    13:39:41 sleepTime: 3000 ms
    13:39:44 Wake up...
    13:39:44 Start the recovery process.
    13:39:44 recover to READY
    13:39:46 [!!!] Ready to output a value for height 1
    ...
    13:39:46 {"View":0,"Height":1,"TXS":[...,{"ID":100,"TX":{"From":"0","To":"2","Value":40"},"Timestamp":1752586786370}]}
    ```

This sequence demonstrates that the sleepy replica has committed two conflicting blocks at the same height, confirming the double-spending vulnerability of HotStuff without stable storage. 



#### Experiment 2.2: Attack on HotStuff-mSS

This sub-experiment aims to reproduce the result shown in **Figure 6** of the paper, demonstrating that our HotStuff-mSS protocol can defend against the double-spending attack.

**Instructions**

Run the following script from the project's root directory:

```bash
./scripts/run_experiment_2_2.sh
```

This script will:
1.  Configure the system to use HotStuff-mSS.
2.  Start 4 replicas. The output of each replica is redirected to a separate log file for analysis (e.g., `experiments/double_spending/HS-mSS/output/server_0.log`).
3.  Simulate a client sending two conflicting transactions (from account `0` to `1`, and from account `0` to `2`).
4.  Cause one replica (replica 2) to "sleep" and "wake up" to simulate a restart.

**Expected Outcome**

1.  **Before sleep:** The replica commits the first transaction (`0 -> 1`).
    ```
    13:58:35 [!!!] Ready to output a value for height 2
    ...
    13:58:35 {"View":0,"Height":1,"TXS":[...,{"ID":100,"TX":{"From":"0","To":"1","Value":40"},"Timestamp":1752587915411}}]}
    ...
    13:58:35 Falling asleep in sequence 5...
    ```

2.  **After restart:** The replica recovers safely based on the parameters in stable storage. It knows it was in View 0 and initiates a view change to View 1. 
    ```
    13:58:38 Wake up...
    13:58:38 Start the recovery process.
    13:58:38 recover to the view 1
    13:58:38 Starting view change to view 1
    ```

3.  **Attack fails:** The log after recovery shows that the blocks from View 0, including the one with the first transaction (0 -> 1), are preserved as part of the committed ledger. 
    
    ```
    13:58:45 [!!!] Ready to output a value for height 3
    13:58:45 {"View":0,"Height":0,"TXS":[{"ID":0,"TX":{"From":"","To":"0","Value":50},"Timestamp":0}]}
    13:58:45 {"View":0,"Height":1,"TXS":[...,{"ID":100,"TX":{"From":"0","To":"1","Value":40},"Timestamp":1752587915411}]}
    ...
    ```

#### Experiment 2.3: Attack on Koala-2

This sub-experiment aims to reproduce the result shown in **Figure 7** of the paper, demonstrateing that our Koala-2 protocol can defend against the double-spending attack.

**Instructions**

Run the following script from the project's root directory:

```bash
./scripts/run_experiment_2_3.sh
```

This script will:
1.  Configure the system to use Koala-2.
2.  Start 6 replicas. The output of each replica is redirected to a separate log file for analysis (e.g., `experiments/double_spending/Koala-2/output/server_0.log`).
3.  Simulate a client sending two conflicting transactions (from account `0` to `1`, and from account `0` to `2`).
4.  Cause one replica (replica 3) to "sleep" and "wake up" to simulate a restart.

**Expected Outcome**

1.  **Before sleep:** The replica commits the first transaction (`0 -> 1`).
    ```
    14:17:50 [!!!] Ready to output a value for height 2
    14:17:50 {"View":0,"Height":0,"TXS":[{"ID":0,"TX":{"From":"","To":"0","Value":50},"Timestamp":0}]}
    14:17:50 {"View":0,"Height":1,"TXS":[...,{"ID":100,"TX":  {"From":"0","To":"1","Value":40},"Timestamp":1752589070118}]}
    ...
    14:17:50 Falling asleep in sequence 5...
    ```

2.  **After restart:** The replica wakes up and initiates the Koala-2 recovery protocol. You will see specific log messages related to this protocol, such as ECHO1 and ECHO2 messages.
    ```
    14:17:54 Wake up...
    14:17:54 Start the recovery process.
    14:17:54 receive a ECHO1 msg from replica 0
    ...
    14:18:00 receive a TQC msg from replica 1 for view 0
    14:18:00 Starting view change to view 1
    ...
    14:18:10 receive a TQC msg from replica 5 for view 1
    14:18:10 Starting view change to view 2
    ...
    14:18:10 receive a ECHO2 msg from replica 4
    ...
    14:18:10 recover to READY
    ...
    ```

3.  **Attack fails:** The log after recovery shows that the blocks from View 0, including the one with the first transaction (0 -> 1), are preserved as part of the committed ledger. 
    
    ```
    14:18:10 [!!!] Ready to output a value for height 198
    14:18:10 {"View":0,"Height":0,"TXS":[{"ID":0,"TX":{"From":"","To":"0","Value":50},"Timestamp":0}]}
    14:18:10 {"View":0,"Height":1,"TXS":[...,{"ID":100,"TX":{"From":"0","To":"1","Value":40},"Timestamp":1752589070118}]}
    ...
    ```





