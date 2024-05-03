# Earthquake Monitoring System

This project collects, processes, and displays earthquake data in real-time, highlighting abnormal earthquakes on a map.

## Getting Started

Follow these steps to run this project in your local development environment.

### Prerequisites

- Docker

### Installation

#### 1. Clone the project to your local machine:

```bash
git clone hhttps://github.com/ouzzkp/earthquake-app
cd earthquake-app
```
#### Start the project with Docker:
```bash
docker-compose up --build
```
This command starts all the required services in Docker containers. The backend service runs on localhost:8080.


## API Usage

#### List All Earthquakes

```http
  GET /api/earthquakes

```

| Parametre | Tip     | Açıklama                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key. |

#### Get a Single Earthquake

```http
  GET /api/earthquakes/{id}
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Unique ID of the earthquake. | 
| `api_key`      | `string` | **Required**. Your API key. |

#### Add a New Earthquake

```http
  POST /api/earthquakes
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `Latitude`      | `string` | **Required**. Latitude of the earthquake. | 
| `Longitude`      | `string` | **Required**. Longitude of the earthquake. |
| `Magnitude`      | `float` | **Required**.  Magnitude of the earthquake. | 
| `Time`      | `string` | **Required**. Time of the earthquake. |
| `api_key`      | `string` | **Required**. Your API key. |

#### Update Earthquake Data

```http
  PUT /api/earthquakes/{id}
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Unique ID of the earthquake. | 
| `api_key`      | `string` | **Required**. Your API key. |

#### Delete Earthquake Data

```http
  DELETE /api/earthquakes/{id}
```

| Parametre | Tip     | Açıklama                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Unique ID of the earthquake. | 
| `api_key`      | `string` | **Required**. Your API key. |



## Technologies Used

**Backend:**

- Go (Golang): Used to build the server-side logic, handle HTTP requests, interact with the database, and employ WebSockets for real-time data transmission.
- PostgreSQL: The relational database chosen for storing and managing earthquake data.
- WebSocket: Used for real-time communication between server and client, particularly for broadcasting new earthquake information instantly.


**Frontend:** 

- React: A JavaScript library used to build the user interface, especially for displaying earthquake data on the map and updating it in real-time.
- react-simple-maps: A React component library used to render the map and place earthquake markers on it.

  
