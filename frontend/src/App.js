import React, { useState, useEffect } from 'react';
import "./App.css";

import {
  ComposableMap,
  Geographies,
  Geography,
  Marker,
  ZoomableGroup
} from "react-simple-maps";


const geoUrl = "https://cdn.jsdelivr.net/npm/world-atlas@2/countries-110m.json";

function App() {
  const [content, setContent] = useState("");
  const [markers, setMarkers] = useState([]);
  const [earthquakes, setEarthquakes] = useState([]);
  useEffect(() => {
    const webSocket = new WebSocket('ws://localhost:8080/ws');

    webSocket.onopen = () => console.log('WebSocket connection established.');

    webSocket.onmessage = (event) => {

      const earthquake = JSON.parse(event.data);
      console.log(earthquake);
      setMarkers(prevMarkers => [...prevMarkers, {
        markerOffset: 25,
        name: `Earthquake`,
        coordinates: [parseFloat(earthquake.Longitude), parseFloat(earthquake.Latitude)],
      }]);
    };

    const fetchEarthquakeData = () => {
      fetch('http://localhost:8080/earthquakes')
        .then(response => response.json())
        .then(data => {
          const allEarthquakes = data.map(earthquake => ({
            id: earthquake.id,
            markerOffset: 25,
            name: `Earthquake ${earthquake.id}`,
            coordinates: [parseFloat(earthquake.Longitude), parseFloat(earthquake.Latitude)],
          }));
          setEarthquakes(allEarthquakes);
        })
        .catch(error => console.error('Error fetching earthquake data:', error));
    };

    fetchEarthquakeData();

    const interval = setInterval(fetchEarthquakeData, 10000);
    return () => {
      webSocket.close();
      clearInterval(interval);
    };
  }, []);

  useEffect(() => {
    const interval = setInterval(() => {
      setMarkers(prevMarkers => {
        const markersCopy = [...prevMarkers];
        markersCopy.shift();
        return markersCopy;
      });
    }, 30000);

    return () => clearInterval(interval);
  }, []);

  return (
    <div style={{ display: 'flex', flexDirection: 'row', width: '100vw', height: '100vh' }}>
      <div style={{ flex: 3, padding: '1rem' }}>
        <h2 style={{ textAlign: 'center', margin: '0 0 1rem 0' }}>World Map</h2>
        <div style={{ height: 'calc(100% - 3rem)', border: '1px solid #ddd' }}>
          <ComposableMap data-tip="" style={{ width: '100%', height: 'auto', maxHeight: '100vh' }}>
            <ZoomableGroup zoom={1}>
              <Geographies geography={geoUrl}>
                {({ geographies }) =>
                  geographies.map(geo => (
                    <Geography key={geo.rsmKey}
                      geography={geo}
                      onMouseEnter={() => {
                        const { NAME } = geo.properties;
                        setContent(`${NAME}`);
                      }}
                      onMouseLeave={() => {
                        setContent("");
                      }}
                      style={{
                        default: {
                          fill: "#D6D6DA",
                          outline: "none"
                        },
                        hover: {
                          fill: "#F53",
                          outline: "none"
                        },
                        pressed: {
                          fill: "#E42",
                          outline: "none"
                        }
                      }}
                    />
                  ))
                }
              </Geographies>
              {markers.map(({ id, name, coordinates, markerOffset }) => (
                <Marker key={id} coordinates={coordinates}>
                  <circle r={10} fill="#F00" stroke="#fff" strokeWidth={2} />
                  <text
                    textAnchor="middle"
                    y={markerOffset}
                    style={{ fontFamily: "system-ui", fill: "#5D5A6D" }}
                  >
                    {name}
                  </text>
                </Marker>
              ))}
            </ZoomableGroup>
          </ComposableMap>
        </div>
      </div>

      <div style={{ flex: 1, padding: '1rem' }}>
        <h2 style={{ textAlign: 'center', margin: '0 0 1rem 0' }}>Previos Earthquake List</h2>
        <div style={{ overflowY: 'auto', height: 'calc(100% - 3rem)', border: '1px solid #ddd' }}> 
          {earthquakes.map(({ id, name, coordinates }) => (
            <div key={id} style={{
              marginBottom: '0.5rem',
              padding: '0.5rem',
              backgroundColor: '#fff',
              borderRadius: '0.25rem',
              boxShadow: '0 2px 4px rgba(0,0,0,0.1)',
              border: '1px solid #ddd'
            }}>
              <div>ID: {id}</div>
              <div>Name: {name}</div>
              <div>Coordinates: {coordinates.join(", ")}</div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default App;