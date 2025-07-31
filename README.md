# go-oui
A Simple Golang Library for OUI Lookup

`go-oui` is a lightweight and efficient Golang library designed to query IEEE Organizationally Unique Identifiers (OUIs). With this library, you can easily map MAC addresses to their manufacturer information using the historical MAC address assignment data from the [Historical Tracking of MAC Address Assignments](https://github.com/yao560909/go-oui) repository.

## Features
- Autoload built-in OUI database
- Support multiple MAC address formats
- Fast OUI lookup

## Installation
```bash
go get github.com/yao560909/go-oui/pkg/oui
```

## OUI Database Information
The oui.csv file utilized in this project is sourced from the official OUI (Organizationally Unique Identifier) database maintained by IEEE (Institute of Electrical and Electronics Engineers). The specific official retrieval address is:
https://standards-oui.ieee.org/oui/oui.csv

To guarantee the accuracy and timeliness of query results, we regularly fetch the latest version of oui.csv from the above official address. The most recent retrieval was completed on July 24, 2025.

If users encounter issues such as lagging or inaccurate query results during usage, they may manually retrieve the latest oui.csv file from the official link provided and replace the existing file in the project to update the database.

Note: In case of access issues (e.g., "system internal error" when accessing the official URL), please retry later or check network connectivity to ensure successful retrieval of the latest data.

## Usage

### Basic Lookup
```go
import "github.com/yao560909/go-oui/pkg/oui"

func main() {
    db := oui.NewDatabase()
    if err := db.Load(); err != nil {
        panic(err)
    }
    
    // Lookup by MAC address
    result, err := db.Lookup("08:EA:44:00:00:00")
    if err != nil {
        panic(err)
    }
    fmt.Println("Manufacturer:", result.Organization)
}
```
## API Documentation

### Query MAC Address Information
Provides a `GET` API to look up the organization information associated with a MAC address.

- **Endpoint**:  
  [`https://mac-lookup-ezhpqkhdyz.cn-beijing.fcapp.run/api/oui/search?mac=D4E3C5FF491E`](https://mac-lookup-ezhpqkhdyz.cn-beijing.fcapp.run/api/oui/search?mac=D4E3C5FF491E)

- **Method**:  
  `GET`
- **Request Headers**:
    - `Content-Type: application/json`

- **Query Parameters**:
    - `mac` (required): The MAC address to query. Supports colon-separated, hyphen-separated, or plain 12-character hexadecimal format.

- **Response Examples**:

#### Success (HTTP 200)

```json
{
    "code": 200,
    "data": {
        "registry": "MA-L",
        "assignment": "D4E3C5",
        "organization": "zte corporation",
        "address": "12/F.,zte R&D building ,kejinan Road,Shenzhen,P.R.China shenzhen  guangdong CN 518057 "
    },
    "msg": "Success"
}
```
