---
name: Bug report
about: Create a report to help us improve
title: ''
labels: ''
assignees: '' 

---

**Describe the bug**
A clear and concise description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:

**Expected behavior**
A clear and concise description of what you expected to happen.

**Screenshots**
If applicable, add screenshots to help explain your problem.

**Logs**
If applicable, add the logs generated from the container.
To know more about logs please refer here: [ServiceBus-Emulator Logs](https://learn.microsoft.com/en-us/azure/service-bus-messaging/overview-emulator#logs-for-debugging)
<img src="https://raw.githubusercontent.com/Azure/service-bus-emulator-installer/main/.github/ISSUE_TEMPLATE/EmulatorContainerLogs.png" alt="EmulatorContainerLogSample">

**Desktop:**
 - OS: [e.g. Windows]

**Image Platform-Architecture Used:**
- Linux-amd64
- Linux-arm64

**Docker Version:**
- Docker Desktop Version (If applicable): 
- Docker Engine Version:

**Arguments && Environment variables to start Emulator:**
- ACCEPT_EULA:
- CONFIG_PATH:

**Emulator Launch Configuration:**
- Include the `Config.json` file used for launching the emulator. Below is a sample of the current checked-in `Config.json`:

```json
{
  "UserConfig": {
    "Namespaces": [
      {
        "Name": "sbemulatorns",
        "Queues": [
          {
            "Name": "queue.1",
            "Properties": {
              "DeadLetteringOnMessageExpiration": false,
              "DefaultMessageTimeToLive": "PT1H",
              "DuplicateDetectionHistoryTimeWindow": "PT20S",
              "ForwardDeadLetteredMessagesTo": "",
              "ForwardTo": "",
              "LockDuration": "PT1M",
              "MaxDeliveryCount": 3,
              "RequiresDuplicateDetection": false,
              "RequiresSession": false
            }
          }
        ],

        "Topics": [
          {
            "Name": "topic.1",
            "Properties": {
              "DefaultMessageTimeToLive": "PT1H",
              "DuplicateDetectionHistoryTimeWindow": "PT20S",
              "RequiresDuplicateDetection": false
            },
            "Subscriptions": [
              {
                "Name": "subscription.1",
                "Properties": {
                  "DeadLetteringOnMessageExpiration": false,
                  "DefaultMessageTimeToLive": "PT1H",
                  "LockDuration": "PT1M",
                  "MaxDeliveryCount": 3,
                  "ForwardDeadLetteredMessagesTo": "",
                  "ForwardTo": "",
                  "RequiresSession": false
                },
                "Rules": [
                  {
                    "Name": "app-prop-filter-1",
                    "Properties": {
                      "FilterType": "Correlation",
                      "CorrelationFilter": {
                        "ContentType": "application/json"
                        // Other supported properties
                        // "CorrelationId": "id1",
                        // "Label": "subject1",
                        // "MessageId": "msgid1",
                        // "ReplyTo": "someQueue",
                        // "ReplyToSessionId": "sessionId",
                        // "SessionId": "session1",
                        // "To": "xyz"
                      }
                    }
                  }
                ]
              },
              {
                "Name": "subscription.2",
                "Properties": {
                  "DeadLetteringOnMessageExpiration": false,
                  "DefaultMessageTimeToLive": "PT1H",
                  "LockDuration": "PT1M",
                  "MaxDeliveryCount": 3,
                  "ForwardDeadLetteredMessagesTo": "",
                  "ForwardTo": "",
                  "RequiresSession": false
                },
                "Rules": [
                  {
                    "Name": "user-prop-filter-1",
                    "Properties": {
                      "FilterType": "Correlation",
                      "CorrelationFilter": {
                        "Properties": {
                          "prop1": "value1"
                        }
                      }
                    }
                  }
                ]
              },
              {
                "Name": "subscription.3",
                "Properties": {
                  "DeadLetteringOnMessageExpiration": false,
                  "DefaultMessageTimeToLive": "PT1H",
                  "LockDuration": "PT1M",
                  "MaxDeliveryCount": 3,
                  "ForwardDeadLetteredMessagesTo": "",
                  "ForwardTo": "",
                  "RequiresSession": false
                }
              }
            ]
          }
        ]
      }
    ],
    "Logging": {
      "Type": "File"
    }
  }
}
```

**Emulator Launch Method:**
 - Direct Docker compose
 - Launcher Script
 - Others (Please specify)

**Additional context:**
Add any other context about the problem here.