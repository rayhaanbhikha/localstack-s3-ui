import React, { useState, useEffect } from 'react';
import { S3Context } from './context'

function App() {

  const [data, setData] = useState({})
  useEffect(() => {
    fetch('http://localhost:8080/data').then(res => res.json()).then(data => setData(data)).catch(err => console.log(err))
  }, [])

  return (
    <S3Context.Provider value={data}>
      <div className="App">
        hello world
    </div>

    </S3Context.Provider>
  );
}

export default App;
