import React, { useState, useEffect } from 'react';
import { S3Provider } from './context'
import { BrowserRouter, Route, Switch, Redirect } from 'react-router-dom'
import mockData from './mock-data.json'
import { Bucket } from './pages'
import './App.css'

function App() {

  const [data, setData] = useState(mockData)
  // useEffect(() => {
  //   fetch('http://localhost:8080/data').then(res => res.json()).then(data => setData(data)).catch(err => console.log(err))
  // }, [])

  return (
    <S3Provider value={data}>
      <BrowserRouter>
        <Switch>
          <Route exact path="/s3" component={Bucket} />
          <Redirect to="/s3" />
        </Switch>
      </BrowserRouter>
    </S3Provider>
  );
}

export default App;
