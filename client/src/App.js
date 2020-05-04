import React, { useState, useEffect } from 'react';
import { S3Provider } from './context'
import { BrowserRouter, Route, Switch, Redirect } from 'react-router-dom'
import mockData from './mock-data.json'
import { Buckets, Bucket, Resource } from './pages'
import './App.css'

export const App = () =>
  <BrowserRouter>
    <Switch>
      <Route exact path="/s3" component={Buckets} />
      <Route exact path="/s3/:bucketName" component={Bucket} />
      <Route exact path="/s3/resource/:resource" component={Resource} />
      <Redirect to="/s3" />
    </Switch>
  </BrowserRouter>

