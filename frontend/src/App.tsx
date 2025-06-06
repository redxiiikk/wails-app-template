import { useEffect, useState } from "react";
import healthCheckApi from "./api/healthStatus";

export default function App() {
  const [healthStatus, setHealthStatus] = useState<string | null>(null);

  useEffect(() => {
    healthCheckApi()
      .then((status) => {
        setHealthStatus(
          status.map((item) => `${item.systemName}: ${item.status}`).join(", ")
        );
      })
      .catch((error) => {
        console.error("Error fetching health status:", error);
        setHealthStatus("Error fetching health status");
      });
  }, [])

  return (
    <div className="App">
      <h1>Welcome to the Wails React App</h1>
      <p>This is a simple application built with Wails and React.</p>

      {healthStatus ? (
        <div>
          <h2>Health Status</h2>
          <p>{healthStatus}</p>
        </div>
      ) : (
        <div>
          <h2>Loading Health Status...</h2>
        </div>
      )}
    </div>
  );
}