# go-oui
 A Simple Golang Library for OUI Lookup

 `go-oui` is a lightweight and efficient Golang library designed to query IEEE Organizationally Unique Identifiers (OUIs). With this library, you can easily map MAC addresses to their manufacturer information using the historical MAC address assignment data from the [Historical Tracking of MAC Address Assignments](https://github.com/yao560909/go-oui) repository.

## OUI Database Information
The oui.csv file utilized in this project is sourced from the official OUI (Organizationally Unique Identifier) database maintained by IEEE (Institute of Electrical and Electronics Engineers). The specific official retrieval address is:
https://standards-oui.ieee.org/oui/oui.csv

To guarantee the accuracy and timeliness of query results, we regularly fetch the latest version of oui.csv from the above official address. The most recent retrieval was completed on July 24, 2025.

If users encounter issues such as lagging or inaccurate query results during usage, they may manually retrieve the latest oui.csv file from the official link provided and replace the existing file in the project to update the database.

Note: In case of access issues (e.g., "system internal error" when accessing the official URL), please retry later or check network connectivity to ensure successful retrieval of the latest data.