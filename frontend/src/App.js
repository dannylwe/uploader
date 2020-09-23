import React from 'react';
import './App.css';
import "bootstrap/dist/css/bootstrap.min.css";

import FileUpload from "./services/FileUpload";
import Hello from "./components/hello";
import Dashboard from "./components/Dashboard";

function App() {
  return (
    <div className="container" style={{ width: "600px" }}>
      <div className="my-3">
        <h4>React Hooks File Upload</h4>
      </div>

      <Hello />
    </div>
  );
}

export default App;
