import React, { useState, useEffect } from 'react';
import "./App.css";

import {
  ComposableMap,
  Geographies,
  Geography,
  Marker,
  ZoomableGroup
} from "react-simple-maps";

import { Tooltip as ReactTooltip } from 'react-tooltip'

const geoUrl = "https://cdn.jsdelivr.net/npm/world-atlas@2/countries-110m.json";

function App() {
  const [content, setContent] = useState("");
  const [markers, setMarkers] = useState([]);

  useEffect(() => {
    const fetchEarthquakeData = async () => {
      try {
        const response = await fetch('http://localhost:8080/earthquakes');
        console.log(response);
        const data = await response.json();
        const newMarkers = data.map(earthquake => ({
          markerOffset: 25, 
          name: `Earthquake ${earthquake.id}`,
          coordinates: [parseFloat(earthquake.Longitude), parseFloat(earthquake.Latitude)], 
        }));
        setMarkers(prevMarkers => [...prevMarkers, ...newMarkers]);
      } catch (error) {
        console.error("Error fetching earthquake data:", error);
      }
    };
  
    fetchEarthquakeData();
    const interval = setInterval(fetchEarthquakeData, 60000); 
  
    return () => clearInterval(interval); 
  }, []);

  return (
    <div className="App" style={{
      display: "flex",
      justifyContent: "center",
      alignItems: "center",
      height: "100vh",
      width: "100vw",
      flexDirection: "column"
    }}>
      <h1>React Simple Maps</h1>
      <ReactTooltip>{content}</ReactTooltip>
      <div style={{ width: "70%", borderStyle: "double" }}>
        <ComposableMap data-tip="">
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
            {markers.map(({ name, coordinates, markerOffset }) => (
              <Marker key={name} coordinates={coordinates}>
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
  );
}

export default App;
