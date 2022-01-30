// import React from 'react';
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Route } from "react-router-dom";
import Login from './pages/Login';
import Register from './pages/Register';
import Home from './pages/Home';
import Nav from './components/Nav';
import ForgotPassword from './pages/Forgot';
import Reset from './pages/Reset';

function App() {
  const [user, setUser] = useState(null);
  const [login, setLogin] = useState(false)
  useEffect(() => {
    (
      async () => {
        try {

          const response = await axios.get('user');
          const user = response.data;
          setUser(user);
        } catch (error) {
          setUser(null);
        }
      }
    )();
  }, [login]);
  return (
    <div className="App">
      <BrowserRouter>
        <Nav user={user} setLogin={() => setLogin(false)}/>

        <Route path="/" exact component={() => <Home user={user} />} />
        <Route path="/login" component={() => <Login setLogin={() => setLogin(true)} />} />
        <Route path="/register" component={Register} />
        <Route path="/forgot" component={ForgotPassword} />
        <Route path="/reset/:token" component={Reset} />
      </BrowserRouter>
    </div>
  );
}

export default App;
