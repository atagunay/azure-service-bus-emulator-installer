
 # <img src="https://github.com/Azure/azure-service-bus-emulator-installer/blob/main/azure-servicebus-emulator.svg" alt="Event-Hubs Logo" width="50">    Azure Service Bus Emulator Installer

This repository contains the scripts required to install and run the  [Azure Service Bus Emulator](https://learn.microsoft.com/en-us/azure/event-hubs/overview-emulator).

- [Azure Service Bus](#About-Azure-Service-Bus)
  - [Emulator Overview](#About-Azure-Service-Bus-Emulator)
  - [Prerequisites](#Prerequisites)
  - [Running Emulator](#Running-the-emulator)
    - [Using Automated Script](#Using-Automated-Script)
    - [Using Docker Compose](#Using-Docker-Compose-Linux-Container)
  - [Interacting with Emulator](#Interacting-with-emulator)
  - [Support](#Support)
  - [License](#License)

## About Azure Service Bus

Azure Service Bus is a fully managed enterprise message broker offering queues and publish-subscribe topics. It decouples applications and services, providing benefits like load-balancing across workers, safe data and control routing, and reliable transactional coordination. Read more [here](https://learn.microsoft.com/en-us/azure/service-bus-messaging/service-bus-messaging-overview).

## About Azure Service Bus Emulator 

The Azure Service Bus emulator offers a local development experience for the Service bus service. You can use the emulator to develop and test code against the service in isolation, free from cloud interference.

>[!CAUTION]
>Emulator is intended solely for development and testing scenarios.Any kind of Production use is strictly discouraged. There is no official support provided for Emulator.
> Any issues/suggestions should be reported via GitHub issues on [GitHub project](https://github.com/Azure/azure-service-bus-emulator-installer/issues).
## Run Azure Service Bus Emulator 

This section summarizes the steps to develop and test locally with Service Bus Emulator. To read more about Service Bus, read [here](event-hubs-about.md).

## Prerequisites

- Docker 
  - [Docker Desktop](https://docs.docker.com/desktop/install/windows-install/#:~:text=Install%20Docker%20Desktop%20on%20Windows%201%20Download%20the,on%20your%20choice%20of%20backend.%20...%20More%20items) 
- Minimum hardware Requirements:
  - 2 GB RAM
  - 5 GB of Disk space
- WSL Enablement (Only for Windows):
  - [Install Windows Subsystem for Linux (WSL) | Microsoft Learn](https://learn.microsoft.com/en-us/windows/wsl/install)
  -  [Configure Docker to use WSL](https://docs.docker.com/desktop/wsl/#:~:text=Turn%20on%20Docker%20Desktop%20WSL%202%201%20Download,engine%20..%20...%206%20Select%20Apply%20%26%20Restart.)

>[!NOTE]
>Before you continue with the subsequent steps, make sure Docker Engine is operational in the background.

## Running the Emulator 

This section highlights different steps to run Service Bus Emulator. Details are as follows:

#### [Using Automated Script](#tab/automated-script)

Before running automated script, clone the installation [repository](https://github.com/Azure/azure-service-bus-emulator-installer) locally.
 
### Windows
After completing the prerequisites, you can proceed with the following steps to run the Service Bus Emulator locally. 
1. Before executing the setup script, we need to allow execution of unsigned scripts. Run the below command in the PowerShell window:

`$>Start-Process powershell -Verb RunAs -ArgumentList 'Set-ExecutionPolicy Bypass –Scope CurrentUser’`

2. Execute setup script `LaunchEmulator.ps1`. Running the script would bring up two containers – Service Bus Emulator & Azure SQL Edge (dependency for Emulator)

### Linux & macOS
After completing the prerequisites, you can proceed with the following steps to run the Service Bus Emulator locally. 

1. Execute the setup script `LaunchEmulator.sh` . Running the script would  bring up two containers – Service Bus Emulator & Azure SQL Edge (dependency for Emulator)

1. Execute the same script `LaunchEmulator.sh` with the option `--compose-down=Y` to issue a `docker compose down` to terminate the containers.

```shell
LaunchEmulator.sh --compose-down=Y
```

#### [Using Docker Compose (Linux Container)](#tab/docker-linux-container)

You can also spin up Emulator using Docker Compose file from Microsoft Container Registry. Refer [here](https://mcr.microsoft.com/en-us/product/azure-messaging/servicebus-emulator/about#usage) for details. 

Once the steps are successful, Emulator compose set can be found in running in Docker.

![image](https://github.com/Azure/azure-event-hubs-emulator-installer/assets/62641016/f7c8d2ad-dea1-4fd5-84b6-8f105ce2b602)

## Interacting with Emulator

1. You can use the following connection string to connect to Azure Service Bus Emulator.
```
"Endpoint=sb://localhost;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=SAS_KEY_VALUE;UseDevelopmentEmulator=true;"
```
2. With the latest client SDK releases, you can interact with the Emulator in various programming language. For details, refer [here](https://learn.microsoft.com/en-us/azure/event-hubs/sdks#client-sdks)

To get started, refer to our GitHub Samples [here](https://github.com/Azure/azure-service-bus-emulator-installer/tree/main/Sample-Code-Snippets).

## Support

There is no official support provided for Emulator.Any issues/suggestions should be reported via GitHub issues on [installation repo](https://github.com/Azure/azure-service-bus-emulator-installer/issues).

## License

The scripts and documentation in this project are released under the MIT License.

The software (Azure Service Bus Emulator) that the scripts in this repository install is licensed under separate terms. See the [End User License Agreement](https://github.com/Azure/azure-service-bus-emulator-installer/blob/main/EMULATOR_EULA.txt) for the terms governing the software.






   


