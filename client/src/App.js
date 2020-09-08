import "./App.css";

import React from "react";

function App() {
  return (
    <div className="App">
      <div className="App-header">Link Shortener</div>
      <div class="App-Container">
        <input placeholder="Link" className="App-input"></input>
        <button className="App-button">Go</button>
      </div>
    </div>
  );
}

export default App;
