import React from 'react';
import { BrowserRouter, Route, Switch, Redirect } from 'react-router-dom'
import {  Resources } from './pages/Resources/Resources'
import './App.css'

export const App = () =>
  <BrowserRouter>
    <Switch>
      <Route exact path="/s3" component={Resources} />
      <Redirect to="/s3" />
    </Switch>
  </BrowserRouter>

